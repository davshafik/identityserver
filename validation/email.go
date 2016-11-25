package validation

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/itsyouonline/identityserver/credentials/password"
	"github.com/itsyouonline/identityserver/db/validation"
	"github.com/itsyouonline/identityserver/tools"
	"net/http"
	"net/url"
)

const (
	emailWithButtonTemplateName = "emailwithbutton.html"
)

//EmailService is the interface for an email communication channel, should be used by the IYOEmailAddressValidationService
type EmailService interface {
	Send(receipients []string, subject string, message string) (err error)
}

//IYOEmailAddressValidationService is the itsyou.online implementation of a EmailAddressValidationService
type IYOEmailAddressValidationService struct {
	EmailService EmailService
}

type translations struct {
	Title      string
	Text       string
	Buttontext string
	Reason     string
	Subject    string
	Urlcaption string
}

//RequestValidation validates the email address by sending an EMail
func (service *IYOEmailAddressValidationService) RequestValidation(request *http.Request, username string, email string, confirmationurl string, langKey string) (key string, err error) {
	valMngr := validation.NewManager(request)
	info, err := valMngr.NewEmailAddressValidationInformation(username, email)
	if err != nil {
		log.Error(err)
		return
	}
	err = valMngr.SaveEmailAddressValidationInformation(info)
	if err != nil {
		log.Error(err)
		return
	}

	translationFile, err := tools.LoadTranslations(langKey)
	if err != nil {
		log.Error("Error while loading translations: ", err)
		return
	}

	translations := struct {
		Emailvalidation translations
	}{}

	r := bytes.NewReader(translationFile)
	if err = json.NewDecoder(r).Decode(&translations); err != nil {
		log.Error("Error while decoding translations: ", err)
		return
	}

	validationurl := fmt.Sprintf("%s?c=%s&k=%s", confirmationurl, url.QueryEscape(info.Secret), url.QueryEscape(info.Key))
	templateParameters := struct {
		UrlCaption string
		Url        string
		Username   string
		Title      string
		Text       string
		ButtonText string
		Reason     string
	}{
		UrlCaption: translations.Emailvalidation.Urlcaption,
		Url:        validationurl,
		Username:   username,
		Title:      translations.Emailvalidation.Title,
		Text:       fmt.Sprintf(translations.Emailvalidation.Text, email),
		ButtonText: translations.Emailvalidation.Buttontext,
		Reason:     translations.Emailvalidation.Reason,
	}
	message, err := tools.RenderTemplate(emailWithButtonTemplateName, templateParameters)
	if err != nil {
		return
	}

	go service.EmailService.Send([]string{email}, translations.Emailvalidation.Subject, message)
	key = info.Key
	return
}

//RequestPasswordReset Request a password reset
func (service *IYOEmailAddressValidationService) RequestPasswordReset(request *http.Request, username string, emails []string, langKey string) (key string, err error) {
	pwdMngr := password.NewManager(request)
	token, err := pwdMngr.NewResetToken(username)
	if err != nil {
		return
	}
	if err = pwdMngr.SaveResetToken(token); err != nil {
		return
	}

	translationFile, err := tools.LoadTranslations(langKey)
	if err != nil {
		log.Error("Error while loading translations: ", err)
		return
	}

	translations := struct {
		Passwordreset translations
	}{}

	r := bytes.NewReader(translationFile)
	if err = json.NewDecoder(r).Decode(&translations); err != nil {
		log.Error("Error while decoding translations: ", err)
		return
	}

	passwordreseturl := fmt.Sprintf("https://%s/login#/resetpassword/%s", request.Host, url.QueryEscape(token.Token))
	templateParameters := struct {
		UrlCaption string
		Url        string
		Username   string
		Title      string
		Text       string
		ButtonText string
		Reason     string
	}{
		UrlCaption: translations.Passwordreset.Urlcaption,
		Url:        passwordreseturl,
		Username:   username,
		Title:      translations.Passwordreset.Title,
		Text:       translations.Passwordreset.Text,
		ButtonText: translations.Passwordreset.Buttontext,
		Reason:     translations.Passwordreset.Reason,
	}
	message, err := tools.RenderTemplate(emailWithButtonTemplateName, templateParameters)
	if err != nil {
		return
	}
	go service.EmailService.Send(emails, translations.Passwordreset.Subject, message)
	key = token.Token
	return
}

//ExpireValidation removes a pending validation
func (service *IYOEmailAddressValidationService) ExpireValidation(request *http.Request, key string) (err error) {
	if key == "" {
		return
	}
	valMngr := validation.NewManager(request)
	err = valMngr.RemoveEmailAddressValidationInformation(key)
	return
}

func (service *IYOEmailAddressValidationService) getEmailAddressValidationInformation(request *http.Request, key string) (info *validation.EmailAddressValidationInformation, err error) {
	if key == "" {
		return
	}
	valMngr := validation.NewManager(request)
	info, err = valMngr.GetByKeyEmailAddressValidationInformation(key)
	return
}

//IsConfirmed checks wether a validation request is already confirmed
func (service *IYOEmailAddressValidationService) IsConfirmed(request *http.Request, key string) (confirmed bool, err error) {
	info, err := service.getEmailAddressValidationInformation(request, key)
	if err != nil {
		return
	}
	if info == nil {
		err = ErrInvalidOrExpiredKey
		return
	}
	confirmed = info.Confirmed
	return
}

//ConfirmValidation checks if the supplied code matches the username and key
func (service *IYOEmailAddressValidationService) ConfirmValidation(request *http.Request, key, secret string) (err error) {
	info, err := service.getEmailAddressValidationInformation(request, key)
	if err != nil {
		return
	}
	if info == nil {
		err = ErrInvalidOrExpiredKey
		return
	}
	if info.Secret != secret {
		err = ErrInvalidCode
		return
	}
	valMngr := validation.NewManager(request)
	p := valMngr.NewValidatedEmailAddress(info.Username, info.EmailAddress)
	err = valMngr.SaveValidatedEmailAddress(p)
	if err != nil {
		return
	}
	err = valMngr.UpdateEmailAddressValidationInformation(key, true)
	if err != nil {
		return
	}
	return
}
