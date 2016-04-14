package oauthservice

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

type authorizationRequest struct {
	AuthorizationCode string
	Username          string
	RedirectURL       string
	ClientID          string
	State             string
	Scope             string
	CreatedAt         time.Time
}

func (ar *authorizationRequest) IsExpiredAt(testtime time.Time) bool {
	return testtime.After(ar.CreatedAt.Add(time.Second * 10))
}

func newAuthorizationRequest(username, clientID, state string) *authorizationRequest {
	var ar authorizationRequest
	randombytes := make([]byte, 21) //Multiple of 3 to make sure no padding is added
	rand.Read(randombytes)
	ar.AuthorizationCode = base64.URLEncoding.EncodeToString(randombytes)
	ar.CreatedAt = time.Now()
	ar.Username = username
	ar.ClientID = clientID
	ar.State = state

	return &ar
}

func (service *Service) validateRedirectURI(redirectURI, clientID string) {
	//TODO: validate redirectURI
	//TODO: seperate case for "itsyouonline"
}

func redirecToLoginPage(w http.ResponseWriter, r *http.Request) {
	queryvalues := r.URL.Query()
	queryvalues.Add("endpoint", r.URL.EscapedPath())
	//TODO: redirect according the the received http method
	http.Redirect(w, r, "/login?"+queryvalues.Encode(), http.StatusFound)
}

func redirectToScopeRequestPage(w http.ResponseWriter, r *http.Request) {
	queryvalues := r.URL.Query()
	queryvalues.Add("endpoint", r.URL.EscapedPath())
	//TODO: redirect according the the received http method
	http.Redirect(w, r, "/authorize?"+queryvalues.Encode(), http.StatusFound)
}

func (service *Service) validAuthorizationForScopes(r *http.Request, username, clientID, requestedScopes string) (valid bool, err error) {

	if clientID == "itsyouonline" {
		valid = true
		return
	}
	valid, err = service.identityService.ValidAuthorizationForScopes(r, username, clientID, requestedScopes)

	//TODO: check if the client still has a valid authorization for the requested scope
	//TODO: how to request explicit confirmation?

	return
}

//AuthorizeHandler is the handler of the /v1/oauth/authorize endpoint
func (service *Service) AuthorizeHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Debug("ERROR parsing form")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Check if the requested authorization grant type is supported
	requestedResponseType := r.Form.Get("response_type")
	if requestedResponseType != AuthorizationGrantCodeType && requestedResponseType != ImplicitGrantCodeType {
		log.Debug("Invalid authorization grant type requested")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Check if the user is already authenticated, if not, redirect to the login page before returning here
	username, err := service.GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if username == "" {
		redirecToLoginPage(w, r)
		return
	}

	//Validate client and redirect_uri
	redirectURI, err := url.QueryUnescape(r.Form.Get("redirect_uri"))
	if err != nil {
		log.Debug("Unparsable redirect_uri")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	clientID := r.Form.Get("client_id")
	service.validateRedirectURI(redirectURI, clientID)

	requestedScopes := r.Form.Get("scope")
	validAuthorization, err := service.validAuthorizationForScopes(r, username, clientID, requestedScopes)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !validAuthorization {
		token, err := service.createItsYouOnlineAdminToken(username, r)
		if err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		service.sessionService.SetAPIAccessToken(w, token)
		redirectToScopeRequestPage(w, r)
		return
	}

	switch requestedResponseType {
	case AuthorizationGrantCodeType:
		redirectURI, err = handleAuthorizationGrantCodeType(r, username, clientID, redirectURI)
	case ImplicitGrantCodeType:
		redirectURI, err = handleImplicitGrantCodeType(r, username, clientID, redirectURI)
	}

	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, redirectURI, http.StatusFound)

}

func handleAuthorizationGrantCodeType(r *http.Request, username, clientID, redirectURI string) (correctedRedirectURI string, err error) {
	correctedRedirectURI = redirectURI

	clientState := r.Form.Get("state")
	//TODO: validate state (length and stuff)

	ar := newAuthorizationRequest(username, clientID, clientState)
	mgr := NewManager(r)
	err = mgr.saveAuthorizationRequest(ar)
	if err != nil {
		return
	}

	parameters := make(url.Values)
	parameters.Add("code", ar.AuthorizationCode)
	parameters.Add("state", clientState)

	//Don't parse the redirect url, can only give errors while we don't gain much
	if !strings.Contains(correctedRedirectURI, "?") {
		correctedRedirectURI += "?"
	} else {
		if !strings.HasSuffix(correctedRedirectURI, "&") {
			correctedRedirectURI += "&"
		}
	}
	correctedRedirectURI += parameters.Encode()
	return
}

func handleImplicitGrantCodeType(r *http.Request, username, clientID, redirectURI string) (correctedRedirectURI string, err error) {

	scopes := ""
	if clientID == "itsyouonline" {
		scopes = "admin"
		//hardcoded override the redirect_uri to prevent spoofing
		redirectURI = "/"
	}
	//TODO: scope mapping for other clients

	mgr := NewManager(r)

	at := newAccessToken(username, clientID, scopes)
	err = mgr.saveAccessToken(at)
	if err != nil {
		return
	}

	correctedRedirectURI = redirectURI
	parameters := make(url.Values)
	parameters.Add("token", at.AccessToken)
	//Don't parse the redirect url, can only give errors while we don't gain much
	if !strings.Contains(correctedRedirectURI, "#") {
		correctedRedirectURI += "#"
	}
	correctedRedirectURI += parameters.Encode()
	return
}
