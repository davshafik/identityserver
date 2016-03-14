// Code generated by go-bindata.
// sources:
// api/companies.raml
// api/contracts.raml
// api/organizations.raml
// api/securitySchemes/oauth_2_0.raml
// api/userorganizations.raml
// api/users.raml
// DO NOT EDIT!

package specifications

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _companiesRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\xdf\x4f\xe3\x38\x10\x7e\xcf\x5f\x31\xa7\xd3\xbd\xd1\x34\x29\x3f\x54\xe5\xe9\xb8\x1e\x07\xbd\x03\x71\x82\xe5\xa1\x5a\xad\x54\x37\x9e\x06\x2f\x8e\x6d\x6c\xa7\x6d\x58\xed\xff\xbe\x4e\xd2\x90\xa4\x04\xe8\xb2\x96\x2a\x55\x1e\xcf\xcc\xf7\x7d\xf3\x23\xbf\xff\x71\x73\x7a\x75\x09\xa1\x1f\x78\x96\x59\x8e\x11\x4c\xad\x99\xc9\xec\x5a\x70\x26\xd0\x5b\xa1\x36\x4c\x8a\x08\x02\x3f\xf4\x16\xc4\xe0\x9d\x66\x11\x0c\x3d\x9b\x2b\x34\x91\x07\x30\x91\xa9\x22\x22\x2f\xfe\x02\x28\x2d\x15\x6a\xcb\x2a\x53\x75\x12\x2e\x17\x84\x33\xda\xdc\x14\xa7\xf0\x8f\xc0\x58\xcd\x44\xd2\x31\xa4\x4c\x5c\xa2\x48\xec\x7d\x04\x87\x5d\x03\xd9\xd4\x86\xf0\x38\x78\x36\xa9\x6c\xc1\x59\xfc\x1f\xe6\xe6\xf5\x04\x9f\xbf\xec\x46\x9a\x5a\x4c\x4d\x04\xa3\x26\x0e\x6e\x14\xd3\xd8\x17\x83\x12\x8b\x9d\x6b\x8d\x8f\x99\x7b\x4b\x23\x58\x12\x6e\x1a\x9b\xd4\x09\x11\xec\x89\x58\x27\xd8\x47\xd0\x84\x41\xb0\x57\x1e\x26\x96\xf2\xd7\xc8\xbe\x15\xdd\x92\x8d\xd0\x7b\x16\xab\x2f\x06\x6e\x48\xaa\x38\xf6\x34\x00\x4c\x2f\xd9\x03\x4e\x88\xa0\xf9\xbb\xe5\x1b\x40\x38\xbe\x3d\xbf\x98\xdd\xb2\xf1\xbf\x93\xd5\xe9\xea\xd3\xdd\xc9\x2c\x5d\x85\x17\x37\x69\x9a\x8d\x4f\xae\x1e\x47\x33\xf5\xb4\x5b\x3b\xc7\x30\x1c\x0f\xc2\x60\xd0\x62\xba\xab\xd5\xa0\xbc\xf9\x93\x71\x87\x24\x2e\x90\xf8\xb1\x4c\x77\xb8\xc3\x5f\x67\xe1\xe8\xf0\xa8\xfd\xf3\xdc\x93\xb8\xea\xf4\x15\xc3\x75\x13\xb1\x92\x66\x3b\x04\x9e\x37\xac\x1e\x6d\x07\x40\x49\x63\xab\xa7\x14\x4d\xac\x99\xb2\xe5\x28\xdd\x60\xc2\x8c\x45\x0d\x04\x04\xae\xeb\xb8\xe5\xbb\x85\xa4\x79\x13\x9c\x28\xe5\xc4\x29\xfb\x69\xf8\xd5\x38\xd7\x9e\xaa\x4c\x9e\xbd\x87\xdf\x2a\xa5\xa7\xf4\xfb\x76\x1a\x33\x5b\x7b\x74\xf2\xdf\xa9\xa2\xa5\x9d\x6a\x0e\x85\xab\x68\x0d\xc0\xaf\x0c\xc5\xcd\x7c\x5e\x87\x9a\xcf\x81\x19\x10\xd2\x02\xe1\x5c\xae\x91\xfa\x5e\x5d\x7a\xa3\x5c\x9b\xb7\x27\x7d\x14\x04\x6d\x84\x5d\x2e\xef\xf3\x79\xc9\xa8\x3a\x47\xc1\x61\xfb\xe1\x4b\x2a\x05\xe2\x1a\x6f\x81\xf6\x1f\xa9\x17\x8c\x52\x14\xbf\x95\x6e\xc3\x76\x0f\x24\x68\x9b\x60\x3d\x1c\x7a\x78\xbc\xc6\x66\x5f\x56\xbb\xfc\x5a\x7d\x54\xe1\x5b\x15\xd3\xe1\x2a\xd2\x8b\xf1\x31\x43\x9d\xff\x4f\x34\x49\xd1\xb5\xcc\xee\x5a\x91\x0f\xd8\x93\xaf\x33\xaf\x55\x8e\x58\x0a\xab\x49\x6c\x4d\x6f\x92\x8e\xa6\xe7\x68\xc1\xde\x23\x3c\xbb\xc0\xfa\x1e\x35\x96\x77\xed\x15\x57\x48\x1d\x82\x5c\x96\x06\x45\xca\xb5\xef\xc3\xb5\xa6\xae\xb3\x8b\x80\x28\x68\x51\x99\x45\x5e\x6e\x50\x7f\x1f\x46\x4c\xc4\x3c\xa3\x78\x56\xce\x33\xed\x6f\x8f\x85\x94\x1c\x89\xd8\xb1\x75\x18\x4c\xab\x30\x25\xb0\x6a\x37\xd0\x86\xcd\x41\x89\x08\x97\x24\xe3\x16\xa4\xe0\x79\xf9\xcc\x59\xd8\xaa\xcd\x99\x38\xc6\x1a\x6d\xa6\x45\xd3\xf3\xf5\x79\x6d\x75\x82\xd3\xdc\xe9\xd0\x8f\x9b\x09\x8b\x09\xea\xb7\x70\xdf\x16\xde\x4e\xd1\xa5\x41\x7b\x00\x99\xc1\x65\xc6\x61\x29\xb5\x53\x37\x71\x52\xfa\xf0\xf7\x16\xb6\x53\x7e\x1e\xcc\xf7\x87\xe5\x3e\x02\x1f\x06\x75\x45\x36\x45\x7e\x04\xc3\x9e\xf0\x5d\x54\xc7\x3f\x03\x6b\x0b\x8d\xa5\x59\xea\x96\xb7\xfb\xa8\xff\x08\x00\x00\xff\xff\x69\x83\xe2\xb8\x8d\x08\x00\x00")

func companiesRamlBytes() ([]byte, error) {
	return bindataRead(
		_companiesRaml,
		"companies.raml",
	)
}

func companiesRaml() (*asset, error) {
	bytes, err := companiesRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "companies.raml", size: 2189, mode: os.FileMode(420), modTime: time.Unix(1457705563, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _contractsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x55\x4f\x6f\x13\x3f\x10\xbd\xef\xa7\x18\xe9\xa7\x9f\x00\xa9\x4a\x36\x55\x5b\x55\xb9\xa0\x12\x5a\x88\x68\x01\xf5\xcf\x01\x21\x0e\xce\xee\x24\x31\xdd\xd8\xc6\x1e\xb7\x09\x7f\xbe\x3b\x63\xef\xae\xb3\x49\x28\x15\x1c\xe8\xa1\xaa\x67\x3d\xf3\xde\x9b\x79\x9e\xfe\xf7\xff\xe5\xc9\xc5\x39\x0c\x7a\x79\x46\x92\x2a\x1c\xc2\x98\xdc\x07\xed\xdf\xa9\x4a\x2a\xcc\xee\xd0\x3a\xa9\xd5\x10\xf2\xde\x20\x9b\x08\x87\x37\x56\x0e\xa1\x9f\x65\xb4\x32\xe8\x86\x19\xc0\x95\x9c\x29\x41\xde\x62\x38\x00\x18\xab\x0d\x5a\x92\xf5\xc7\xf0\xe3\xf8\x02\x96\x2f\x56\x43\x70\x64\xa5\x9a\x35\xe1\x52\x10\x83\x85\xdf\x4d\xc0\xf8\x49\x25\x8b\x37\xb8\x6a\x13\x01\x02\xc8\x56\x1a\x27\xa2\x2b\xac\x34\x14\x69\x5d\xcf\xb1\x49\x84\x5b\x5c\x81\x77\x58\x76\x60\x3b\xbc\x76\xca\x71\x74\xa4\x15\x59\x51\xd0\x43\xcc\x8d\xd8\x38\x6e\x16\xf8\xf8\x29\x85\x17\x52\x8d\x09\x17\x6e\x08\xfb\xeb\x98\x58\xb6\xb1\x3c\x05\xbd\x92\x5f\x3c\x36\x71\xb2\xbe\x95\x5e\x30\x11\x54\xb4\x25\xb5\x68\xe8\x5d\x07\xd4\xdf\xf7\x84\xd1\xce\x51\xcd\x68\x3e\x84\x83\xfc\xd7\x9d\x3a\xb3\x88\x30\x95\x58\x95\x40\x1a\x26\x18\x7b\x05\x53\x6d\x81\xc4\xcc\xc1\x64\x05\x14\x7a\xc9\x92\x57\x50\x58\x14\xc4\xf5\x41\x24\x12\x4d\x51\xa9\xee\x44\x25\xc3\xd4\xdc\xf3\xc7\xfb\x92\x7a\x30\x78\x80\x54\x25\x1d\x81\x9e\x26\x94\x71\xf9\xc4\x31\x0f\xe9\x52\xa4\x8b\xd8\x6b\x8a\xe0\x92\xbb\x55\xfe\x33\x02\x0e\x84\xaa\x31\xc3\x53\x00\x6f\xb4\x4a\x44\x8c\xb4\x6c\x90\xae\x8d\xd7\x95\xfe\xc4\xc7\xdf\x53\x18\xa2\xa7\xd7\x55\x02\x7e\x98\xcc\x5c\xb8\x79\x60\xda\x4e\x89\x8d\xb9\x17\x0f\x8d\x79\xd6\x87\xd6\x33\x7b\xc0\xc4\xcb\x18\x8e\x44\x45\x80\x8a\x54\x7b\x5d\x34\x0d\x85\xa8\x0a\x5f\x71\x3c\x01\x71\x31\x71\x5b\x1f\x3f\x3b\x4e\xb2\x68\x58\x27\xa3\xd4\x35\x1a\x1a\xa9\x45\xf7\x92\xe6\xda\x53\x0c\xca\x32\xa2\xa6\xe7\x17\x58\x5a\xb9\x88\xdf\x14\xde\x87\xad\xe2\x80\x5d\xe7\x8c\x28\xf8\xaf\xa7\xb8\x2c\xd0\x70\x89\x39\xaa\x70\x67\x05\xc2\x18\x14\x96\xe7\x1e\x53\x78\xf8\x1e\xdd\xb3\x58\x13\x55\xa1\x4b\x0c\x5f\xc0\xd3\xf4\xb8\xab\xe2\xcc\xea\x45\x3d\x37\x46\xf4\x15\x35\x02\x04\x5c\xbd\x3e\xd9\x3f\x3c\x62\x9f\x57\x95\xbe\x67\xc3\xb3\xcf\x05\x5c\x8e\xdf\x9f\x5e\xbc\x1c\x1c\xe5\x7b\x4c\x69\x49\x10\x16\xdb\xe1\xf1\x68\x8e\xc5\xed\x69\x8d\x11\xa0\xeb\x4a\x51\x1c\xe7\xe4\xcb\x3c\xe7\x1d\x81\x53\xb9\xec\x65\xdb\x3b\x86\x2d\x90\xf6\x20\xfb\x2f\xeb\xb7\xad\x89\xcb\xc3\x68\xd7\xec\x98\x8d\x99\x8f\xc2\x2b\x0b\x1c\xb9\x2d\xa9\x97\xb5\x28\xae\xc8\x2e\x73\xeb\xdd\xb3\x9f\x0f\xd6\x76\x9a\xe8\xb2\xb3\x24\x21\x74\x8c\x17\x60\x1c\x4d\x3f\x8c\xab\xfb\xad\x35\xdf\x68\xf3\x1d\x1f\x74\xeb\x6d\xb0\xba\x51\xc2\xf3\x34\xad\xfc\xca\xab\x94\xaf\xf4\xbf\xad\xbd\xf8\xa3\xce\x99\x21\xb5\xc9\x1b\xa9\xaf\x90\x76\x37\xc6\x8e\x94\x20\x26\xef\x52\xdc\x96\xf3\x98\xa0\x07\x24\x05\x51\x07\xdd\x9b\x1b\xdc\xde\x6a\x62\x13\x78\x55\x6e\x4b\xea\x77\x86\x58\xff\x1f\x48\xd3\xda\x29\x12\x66\xbc\xab\xf0\xef\x06\x92\xfc\x92\xfd\x0c\x00\x00\xff\xff\x50\xa8\x00\x0d\x83\x07\x00\x00")

func contractsRamlBytes() ([]byte, error) {
	return bindataRead(
		_contractsRaml,
		"contracts.raml",
	)
}

func contractsRaml() (*asset, error) {
	bytes, err := contractsRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts.raml", size: 1923, mode: os.FileMode(420), modTime: time.Unix(1457614661, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _organizationsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x57\x51\x6f\xdb\x36\x10\x7e\xd7\xaf\x20\x30\xec\xad\x51\xec\x36\xe9\x30\xbf\x19\x6b\x51\x04\x73\x92\x21\x41\x1a\xac\xc3\x80\x50\xd2\x49\xe6\x2a\x91\x2a\x49\x25\x76\xb7\xfe\xf7\x1d\x49\xc9\xb2\x24\x4a\x76\x82\x62\x0f\xcb\x4b\x60\x89\xc7\xfb\xee\xfb\xbe\x23\x4f\x3f\xfc\x78\xb3\xbc\x5c\x91\x79\x38\x0b\x34\xd3\x39\x2c\xc8\x85\x56\xbf\x8b\xea\x9a\xe7\x8c\x43\xf0\x08\x52\x31\xc1\x17\x64\x16\xce\x83\x88\x2a\xb8\x93\x6c\x41\x4e\x83\x40\x6f\x4b\x50\x8b\x80\x90\x6b\x99\x51\xce\xbe\x52\x6d\x96\xe1\x6f\x42\x4a\x29\x4a\x90\x9a\xb9\xf7\xe6\x2f\xcb\x45\x44\x73\x96\x34\xbf\x09\x31\xe1\x0b\xa2\xb4\x64\x3c\xdb\x3d\x2c\x18\x5f\x01\xcf\xf4\x7a\x41\xde\xb4\x0f\xe9\xa6\x79\x38\x3f\x9f\xd5\x8f\xcb\x2a\xca\x59\xfc\x2b\x6c\x95\x7f\xcb\x3f\xfe\xdc\x8f\xbf\xd0\x50\xa8\x05\x79\xdd\x44\x27\xfc\x39\x61\xf3\xd9\x6c\xf7\x34\x01\x15\x4b\x56\xda\x5a\xc9\x8a\x29\x4d\x44\x4a\xc4\x1e\x03\xe4\xdd\xd5\x6d\x58\x2f\x17\x4f\x1c\xd9\x7b\x09\xc0\x91\x44\x95\xc2\xfd\xc8\x83\xf9\xc7\x69\x01\xea\xa1\xc9\x54\x40\x11\x3d\x37\xd5\xa1\xaa\x5c\xb2\xa7\xb5\x20\x54\x42\x9d\x81\x30\x4e\xf4\x9a\xa9\x4e\xc9\x0d\x08\xc6\xe3\xbc\x4a\xe0\xbb\x50\xeb\x0c\x73\x91\x18\x20\xaa\x8a\x3a\xf9\x54\x60\x63\x60\x43\x8b\x12\xfd\xda\xb7\x18\xc9\x24\x00\x67\xda\x3c\x80\x09\xbb\x9c\x90\xf9\xf2\xe7\xfb\xfb\xf5\x5b\xb6\x7c\x7f\x76\xf3\xe9\xc3\xd5\x4f\x74\xb3\x7d\xbb\xf9\x14\x55\xf7\xd5\xea\x33\xff\x72\xbf\x92\x1f\x7d\x7e\x39\xe9\x24\x08\x63\x51\x8c\xc8\x7d\x62\x29\x9c\xf7\x7e\xbf\xee\xfd\x7e\x33\x26\xa1\x7b\x7d\x36\x4a\xee\x09\xa1\x09\x36\x4c\x38\x40\xd3\xeb\xc9\x3a\xe4\xa0\x73\x55\xd0\x4a\xb6\x1f\x8f\xc2\x05\x0d\xbc\xb1\xfe\x6e\x1c\x79\xa0\xbf\x3b\x18\xee\x14\x24\x68\x2f\xe0\x84\x2a\xc5\x32\x8e\x0b\x09\x6d\x7c\xa6\xd1\x75\xbc\xe7\x32\x9f\xea\xbb\xc4\x24\x12\x51\x10\x9c\x8a\x7e\xe1\xa5\x50\xda\x43\xc0\x2f\x12\xa8\x06\xcc\xc7\xe1\xa9\x9b\x86\xcc\xed\xa6\x44\xad\x45\x95\x27\x24\x02\xe7\x79\xa8\xe5\x25\x39\x52\x17\x92\x8f\xc6\x6a\xae\xe1\xb1\x1b\x90\x8d\x54\xc8\x02\xeb\x41\xdc\xf1\x1a\xe2\xcf\x84\xa5\x36\x48\x41\x5c\x49\xa6\xb7\xb7\xf8\xb4\xc0\x7c\x79\x2e\x9e\x14\x36\x00\xa7\x19\xfe\xe6\x28\x42\xdd\x50\x26\xa7\x6b\xa4\x48\x24\xdb\xa6\x3c\x5a\x96\xe8\x5b\x9b\xe7\xf4\x2f\xd5\x9c\xaf\x7e\x9d\xec\x2b\x09\xaa\xc4\xca\x5b\x65\x5e\xcf\xe6\x6d\xd0\xfe\xd6\xd3\xdb\xfb\x53\xa8\xfa\xfd\xd9\xfe\xa6\x5d\x51\x39\xad\xf4\x5a\x48\xf6\x15\x12\x5c\x91\x81\x8f\xfc\x7f\xea\xd8\x0f\xa0\xbb\x0e\x0c\xc9\xb2\x8e\x76\xd4\xe6\xac\x60\x5a\xd9\xf3\xc7\x42\x75\x04\x4b\xf8\x52\x81\xd2\xc6\x2f\x2d\x6b\x9e\xc2\x67\xdf\xb3\x70\x5c\x70\xfa\x77\x73\xc6\x7c\x73\xeb\x77\xd5\xf5\xea\xeb\xd7\x85\x0e\x4a\x45\xbd\x70\x80\xb3\x87\x74\x88\xf5\x10\xda\x51\x2f\x0c\x94\x9a\xd6\xaa\x09\x38\x1b\x0d\xb8\x12\x9a\xa4\xa2\xe2\x6e\x75\x59\xf9\xab\xbf\x2b\x13\xd3\x5a\x63\x04\x74\xcb\x9b\x2a\x6d\xb4\xac\xff\x07\x87\x76\xc5\x69\xef\xd8\x6f\x8f\xab\x41\xe4\xd2\x9e\x92\x9d\x23\xd2\x77\x0b\xbf\xcc\xec\x6e\xcf\xdd\x0b\x0f\xc3\xbd\xa3\x64\x00\xef\xd2\xa1\x72\x67\x39\x36\xaa\xaa\xe2\x18\x94\x4a\xab\x3c\xdf\x76\xa2\x86\xda\x1c\x56\x67\x04\xe6\x40\x9b\x63\xd4\x19\xe8\x33\xad\x90\xe9\xfb\xe6\x96\xf9\xb6\x2f\x4d\x0e\x1a\x46\x65\xbe\x81\x42\x3c\x42\x2b\x56\x2a\x45\xd1\x91\x6b\x2f\xd0\x4b\xb6\xa1\xfb\xac\x4f\x83\x8f\x70\x07\x64\x82\xef\x01\x45\xc7\x91\xe4\xa1\x69\x9c\x28\xe7\xe5\xee\x00\xd4\xb5\x72\x2f\xf6\x82\x3f\x32\x7b\xff\xda\xbb\x16\xad\x1c\x01\x8e\x2e\xf5\x25\x6b\x26\x93\xc1\xed\xdf\x6e\xf4\xd2\xe6\x1e\x98\x67\x94\xf8\x69\xc2\x6a\xec\xb1\x1d\x21\x26\x88\xf7\x5b\xfd\x18\xb3\x8f\x22\xfe\x6f\xd4\x74\x6f\xbd\xb6\xf7\x19\x7f\xc4\xfa\xbc\x16\x73\xca\xfa\xa3\x1a\x78\xed\xdf\x4b\x74\x6d\xb7\x97\x36\xdd\xa4\x0c\x1e\xce\x8e\x65\xcd\xcb\xdb\xa1\x3e\x88\x05\xd7\x92\xc6\xba\xfd\xfc\x85\xb1\x43\xdd\x8c\x0a\x66\x4e\xdc\x85\x98\x89\x18\xc7\x1d\x3b\x70\x76\x6e\x50\x85\x83\xa9\x70\x43\x65\x49\xed\xf0\x1d\xe2\x75\x95\xd8\x13\x40\xc5\xc0\x13\x33\x0f\x45\x5b\x62\x2e\xdf\xb6\x5b\x70\x52\x92\xdb\xdf\xa8\x44\x15\x75\xe7\xeb\x62\xf7\x45\xf1\x7e\x53\x32\x09\x89\xbf\x65\x22\x21\x72\xa0\x7c\xba\x1b\xec\x36\x16\x18\xb8\xbd\xda\x6a\x5e\x59\x44\x90\xd2\x2a\x37\xa3\x6e\xbe\xb5\xcb\xf0\x0d\x7b\xdc\xaf\xd9\x0c\x78\x12\x74\x25\xf1\xe2\x08\x7b\xc9\xcc\xb4\x67\x01\x92\x94\xe6\x0a\xf6\xde\x2a\x8d\x3c\xf8\x71\x33\xae\x21\xeb\x75\x4e\x0f\xf7\xad\x89\x46\x46\x53\x05\xfa\x95\x39\x86\xd0\x39\x28\xa3\x44\x76\x33\xa4\x32\x24\xef\x6a\xd8\xc8\xfc\xc3\xec\xe1\x78\x58\xf8\x75\xfb\x62\x50\x97\x74\x63\xf2\xe3\x77\x03\x3a\xf1\x20\xaa\xf3\xe7\xc0\xaa\xa1\xb1\xa2\x2a\xf0\xeb\xff\x7c\x16\xfc\x1b\x00\x00\xff\xff\x4e\xf8\x72\x1c\xf9\x11\x00\x00")

func organizationsRamlBytes() ([]byte, error) {
	return bindataRead(
		_organizationsRaml,
		"organizations.raml",
	)
}

func organizationsRaml() (*asset, error) {
	bytes, err := organizationsRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "organizations.raml", size: 4601, mode: os.FileMode(420), modTime: time.Unix(1457705524, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _securityschemesOauth_2_0Raml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x92\x3f\x6f\xdb\x40\x0c\xc5\x77\x7d\x0a\x42\x4b\xa7\x4a\x86\xe1\x49\x9b\x8b\x02\x1d\xdb\xa1\x9e\x8a\xa2\xb8\xea\x08\xeb\x10\xe5\x78\x21\x29\x3b\x0a\xf2\xe1\xc3\xd3\x1f\xc3\x36\x32\x64\x8a\xa6\x03\xa9\xf7\xf8\xe3\xbb\xf3\x28\x2d\x87\xa4\x81\x62\x03\xaf\x05\xc0\xcf\xfd\xa0\xdd\x16\x82\x80\x83\xc4\xa4\xd4\x52\x0f\xda\x39\x85\x1e\x55\x00\x9f\x15\x39\xba\x1e\x5c\x4a\x02\x8c\x4f\x03\x8a\x82\x33\x0d\x71\x78\x71\xd9\x07\x94\x4c\x19\x4e\x4e\xd1\xfc\x3c\xaa\x0b\xbd\x40\x88\x66\x38\x08\xf2\x17\x73\x6e\x5b\x1a\xa2\xc2\x39\x98\x6c\x50\x38\xa2\x6a\x88\x47\x1b\x83\x81\x21\x39\x91\x33\xb1\xaf\x0a\x1d\x13\x36\x33\x11\x6c\xab\x4d\x31\xc3\xfe\x47\xff\x6d\x6c\xcc\xba\x43\xe7\x91\x25\x1f\x01\xf6\xd7\x08\x73\x29\x0f\xbf\xdb\x6e\xfe\x0e\x82\x3e\x53\x0a\x46\x6f\x54\x27\xd7\x07\xbf\x8e\xc9\x70\x28\x62\xed\x07\x8c\x15\x7c\x27\x88\xa4\x19\xdc\x2a\xc6\xd9\x21\x4f\xd8\x17\x2f\xab\x40\x39\x6b\xfe\x4d\x9a\x12\x2c\x13\x1e\x41\x94\xf3\x4e\xc9\xb1\x7b\x44\x0b\xad\x32\xc9\xd4\xf9\xb5\x56\x16\xf2\x6b\xf1\x67\x83\xdf\xa4\x56\x2e\x89\xda\x2f\x8c\x92\x28\x0a\x2e\x88\xbb\xcd\xee\x5d\xb2\x43\x5c\x6f\x1e\x7d\x21\xf3\x2d\x4e\x92\x9b\x07\x71\xe0\xd0\x40\xa7\x9a\xa4\xa9\xeb\xa0\x32\xd2\x50\x51\xec\x43\xc4\xba\xa7\x63\x88\x35\xe5\xdf\xeb\x8b\x55\xb1\x66\xf2\x3b\xaf\xf2\x51\xf5\x55\x8a\xf7\x00\x3f\xd8\x45\x95\x06\xfe\x40\x4b\x1e\xe1\xaf\xf5\xa5\xa5\xb4\x6e\xf7\x75\x7a\x97\xcb\xb1\xcc\xe7\x86\x2d\x88\xb2\x78\x0b\x00\x00\xff\xff\x88\xbd\xe6\x45\x1d\x03\x00\x00")

func securityschemesOauth_2_0RamlBytes() ([]byte, error) {
	return bindataRead(
		_securityschemesOauth_2_0Raml,
		"securitySchemes/oauth_2_0.raml",
	)
}

func securityschemesOauth_2_0Raml() (*asset, error) {
	bytes, err := securityschemesOauth_2_0RamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "securitySchemes/oauth_2_0.raml", size: 797, mode: os.FileMode(420), modTime: time.Unix(1454406267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userorganizationsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x41\x8f\xd3\x30\x10\x85\xef\xfe\x15\x23\x21\x6e\x28\x49\x11\xa7\xdc\xf6\x84\x16\x81\x56\x5a\xc4\x01\x55\x3d\x38\xc9\x34\x9d\xca\xb1\x2d\x8f\x53\x54\x4a\xff\x3b\x76\x92\x52\xa7\xa4\x85\xb9\xb8\xea\x8c\xdf\x7b\xf3\x39\x6f\xde\xbe\x3e\x7d\xf9\x0c\xab\xac\x10\x9e\xbc\xc2\x12\x9e\x3d\x7f\x37\xfd\x8b\x56\xa4\x51\x1c\xd0\x31\x19\x5d\x42\x91\xad\x44\x25\x19\xbf\x39\x2a\x21\x17\xfe\x68\x91\x4b\x21\x60\xa8\x4f\x86\xf4\x8b\x6b\xa5\xa6\x9f\xd2\x87\xf1\x67\x7d\x20\x3f\xfc\x2a\xa7\x09\x00\xeb\x8c\x45\xe7\x29\x5e\x83\xa4\x4c\x72\xaf\x04\xf6\x8e\x74\x3b\x1b\xe8\x19\xdd\x62\xc3\x99\x10\x77\xf6\x4f\xac\x98\x6c\x71\x3c\x16\xea\xbe\x2b\x61\x6d\x7e\x68\x74\xef\xa0\xc3\xae\x42\xb7\x11\xa1\xf2\xe8\xc2\xf9\x29\x1e\x5a\x76\x78\xce\xd3\x5c\x43\xe4\x16\xfd\xe8\xd6\x20\xd7\x8e\xec\x18\xf8\x23\x7a\xf0\x3b\x04\x45\xec\x67\xbb\x30\xc8\x21\x3a\x10\xc3\xe0\x17\xba\x93\x23\x98\xed\x20\xe4\x90\x6d\x18\xbc\x12\x79\x5f\x14\xd7\x85\x2a\xd3\x1c\xd3\xf5\xa4\xb5\x8a\xea\x41\x3b\xdf\x73\x8a\xf6\x11\x5e\x18\xcd\x2f\x44\xd6\x9b\x9b\xee\x98\x28\x69\x87\x7e\x7e\x6a\x95\xa9\xa4\xa2\xe6\x9c\x47\xc8\x81\x4b\x3c\xce\xa3\xb2\x35\xec\x2f\x1e\x33\x14\x4f\x75\x8d\xd6\x4f\x8a\xbc\x23\x0b\xa4\x67\x48\xc4\xd2\x62\x8f\xd6\x1a\xdf\xf2\xfe\xd7\x35\x8d\xfe\xc5\x31\x92\x5c\xa5\x42\xb7\x2c\xff\x45\xf3\x3f\xad\x1b\x54\xe8\x71\x11\xc6\xaf\x3f\x7a\xaf\xb8\xc7\xfa\x06\xcb\x45\x25\x12\x92\x73\x48\xd9\xa3\xa5\x3e\xa4\x31\xef\xf8\xc5\xfa\xda\x87\xb7\xe0\x6d\xaf\xd4\x31\x08\x45\x7f\x6c\x12\xdb\x4c\xfc\x0e\x00\x00\xff\xff\x2b\x39\x6a\x40\xf7\x03\x00\x00")

func userorganizationsRamlBytes() ([]byte, error) {
	return bindataRead(
		_userorganizationsRaml,
		"userorganizations.raml",
	)
}

func userorganizationsRaml() (*asset, error) {
	bytes, err := userorganizationsRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "userorganizations.raml", size: 1015, mode: os.FileMode(420), modTime: time.Unix(1457621431, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _usersRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x7b\x73\xdb\x36\x12\xff\x5f\x9f\x02\xe7\xbb\x9b\xbb\x4e\x6c\xbd\x2c\xa7\x89\xfe\x49\x1c\xe7\x51\xa7\x76\x92\x89\xed\x76\x3a\x6e\x3a\x06\x49\x48\x42\x0d\x02\x2c\x00\x4a\x56\x32\xfd\xee\x5d\x80\xa4\x04\x92\xa0\x2d\xf9\x51\x3b\x99\x6a\x26\xb1\x44\x2c\x16\xbf\x7d\x62\x17\xe0\xbf\xff\xfb\x71\xf7\xf0\x00\xf5\xda\xdd\x96\xa6\x9a\x91\x21\xda\xd7\xea\x17\x91\xbe\xe7\x8c\x72\xd2\x9a\x12\xa9\xa8\xe0\x43\xd4\x6d\xf7\x5a\x01\x56\xe4\x44\xd2\x21\xea\xb4\xf4\x3c\x21\x6a\xd8\x42\xe8\x28\x14\x09\x31\x5f\x10\x8a\x88\x0a\x25\x4d\xb4\xa5\x7f\x2d\x24\xc2\x1c\x91\x8b\x84\x61\x8e\xcd\x33\x84\x03\x91\x6a\xa4\xcc\x04\x05\x63\x51\xf6\x35\xc6\x49\x42\xf9\x78\x13\x29\x42\xd0\x44\xeb\x44\x0d\x3b\x9d\x31\xd5\x93\x34\x68\x87\x22\xee\x50\xad\xe6\x22\x15\x16\x4e\x87\x46\x84\x03\xcc\xb9\x22\x12\x90\x75\x02\x26\x82\x4e\x8c\x95\x86\xef\x91\x08\x55\x47\xe0\x54\x4f\xfa\x9d\x6c\x8d\x76\x1c\x59\x5c\x89\x84\x5f\x52\xd3\x0c\x70\xf6\x49\x81\x03\xc7\x31\x88\xab\xb4\x84\xe5\x17\x03\x63\x89\xb9\x26\xd1\xb1\x58\xd2\x22\x64\xa4\xad\x51\x56\x04\x3e\x9e\x10\x24\xe4\x18\x73\xfa\x39\x13\x57\x4f\xb0\x46\x54\x15\x1c\x91\x24\x38\x42\x38\x0c\x89\x52\x48\x0b\xa4\x44\x0c\x13\x46\x40\x46\x2c\x9a\xff\x29\x44\xf9\x48\xc8\xd8\xce\x5e\xac\x82\xa3\x48\xc2\x8c\x67\xc3\xa6\x75\x0f\x33\xfd\x15\xac\x18\x0e\x08\x53\xc5\xaf\x62\x6d\xc1\x89\x5d\xd4\xa1\xd0\x55\xbc\x92\xfc\x91\x12\x50\x65\xd4\x76\x96\xf2\xe9\xce\x7c\x36\x4e\x3f\x6d\x94\x9f\x34\x68\x29\x49\x03\x46\xc3\x1f\xc9\x1c\x44\xc8\xc7\x4e\x3f\x2d\x46\xc1\xfa\x94\x7d\xbd\xb2\x4d\x60\xed\xaf\x16\x7d\x80\xf9\xf9\x57\x0b\x7e\x84\x43\x12\x08\x01\x02\x20\xf8\xc3\x08\x5e\x46\x4c\x96\x3b\x3c\x03\x2e\x2c\xd5\x14\xde\x8e\x73\x56\x34\x72\x40\x95\x36\x0a\x28\xb1\xb1\xf2\xe6\x22\x5a\x75\xb9\xa2\x87\x90\x00\x4d\x56\x83\xfc\x25\x51\x4c\xe2\x00\x92\xe9\x84\x26\xc0\xa4\xdd\x82\x55\x5e\x80\x01\x76\xc3\x50\xa4\x5c\x0f\x1b\x13\x15\x05\x33\xad\x90\x8a\x62\x7c\x71\x40\xf8\x58\x4f\x86\x68\xbb\xbb\x34\x31\x0d\xd7\x9b\xdb\x5b\xce\xb5\xb8\xe4\x7c\xbd\xf9\x83\xae\x11\xec\x83\x89\x0b\x9e\x1a\x79\x73\xc1\xb0\x86\x14\x0d\x3a\xfc\xed\xd7\x47\xa7\xdd\xad\xa7\x9f\x1e\xfd\xc7\xd0\xed\x66\xb9\xad\x59\xf8\x10\x52\x7d\xd9\x1d\x1a\x20\x34\x29\x00\x08\x09\xd1\xeb\xb3\xd8\x59\xb2\xe0\x72\xfd\xe9\x8e\x1a\x05\x38\x88\x7c\x76\x23\x21\x3c\x96\x58\x8d\xc7\x60\xc9\x23\x11\x4a\x63\x16\x8a\x88\xac\xcf\xa6\x6f\xd8\x9c\xa8\x85\x31\x2f\xdb\x4e\x6f\x22\xa7\x1d\xa1\x7c\xb1\xac\x67\x0f\xf1\x6d\x21\x17\x09\x95\x90\x85\x51\x84\x35\xb9\x64\x63\xb9\x83\x6c\xbf\x3e\x4b\x27\x32\x2e\xdd\xe1\xd7\xe7\x9c\xc7\x52\x73\x76\x5f\x9f\xa5\x93\x9f\x3c\x79\xf7\x26\xde\x58\xa4\xe8\xeb\xc4\xa5\x1d\x22\x17\x38\x4e\x18\xf1\xd5\x73\x81\x08\x7c\x8e\x53\xe2\xb8\x85\x7a\x6f\x7e\x3a\xfa\x7e\x67\x30\xd0\xd3\x9f\x0f\xfb\xc7\x87\xbd\xed\xe9\xbb\x1f\x9e\xa8\x40\xe0\xf8\xed\xe7\x83\xc7\xf1\xdb\xed\xc7\x15\x07\x33\x61\xd0\x7b\xb2\xd5\xeb\x6e\xf5\xbb\x65\x2f\x2b\xb3\x9e\x09\x79\x6e\x41\x3c\x87\xe2\x35\xc1\x7c\x6e\x8a\xd8\x12\xc5\x44\xe4\x30\x9f\xe7\x52\xd4\x28\x36\x22\xcc\x43\x82\x42\x96\x06\x1b\x19\xa5\x29\x7e\xed\x43\xf3\xac\x5d\xcc\x83\xcd\xa6\xec\x9d\x2e\x94\x00\x16\x79\xb4\xdd\xef\xf5\xb7\x07\xf9\x3f\x67\x90\x8c\x61\x10\x04\xf2\x0c\x62\x37\x29\x97\x30\x57\x9c\x24\x4f\xcf\xe8\x28\x31\x16\x1b\x51\xc2\xa2\x1a\x45\x9e\x7f\xa1\x9a\xa0\x3c\xff\x51\xa3\x81\x04\x8b\x06\x3b\x2f\xea\xdc\xf3\xb4\x87\x8e\x45\x2c\xa4\x14\x33\x68\x24\xea\x2b\x38\x89\x0d\x1d\x6f\x3d\x1d\xec\x74\xeb\xe6\x68\x00\xbe\x87\x13\x0a\x73\xd1\x1e\xfc\x6a\x44\x5e\x10\x5d\x02\xbe\x77\x2b\xd0\x7b\xdd\x6e\xb7\x14\xbd\x65\xd8\xe7\x41\x58\x97\xc3\x16\x07\xe8\xf8\xc0\xd8\x70\x67\xf9\x5f\x8d\xce\xd4\x01\x68\xf7\xc5\xde\xcb\x57\xaf\x57\xc4\xda\xca\x42\x6a\x4a\xc9\xec\x56\xb2\xfe\xba\x6d\xc0\x3f\xd9\xfa\xdb\xc8\xd6\xc5\xd3\x72\xed\xbd\xf4\x00\x20\x78\x2b\x28\x7f\xef\x0c\xef\xf3\x29\x44\x9c\xad\xb9\x1b\x3d\xcf\x65\x57\x43\x61\xdc\xb2\xf6\x50\x0a\xb6\x6a\x75\x62\x4c\x3e\x44\xa7\x62\xc6\x89\xdc\xcc\xeb\x76\x8b\x74\x4f\x40\xa4\xe0\x50\x1f\xd1\x31\x87\x49\x1f\xb3\xb2\xff\x92\xf2\x35\xa7\xdf\x8f\xea\xee\x89\xa5\x49\x41\xf9\xd3\x56\xc7\x60\xb6\x33\x4d\x52\xf0\x1c\xa7\xec\x49\x02\xe5\x0d\xc2\x88\x93\x99\x15\xd0\x92\x04\x22\x72\xca\x42\xe8\xd8\x20\xc6\xac\x4e\x3a\xbf\x2b\xc1\x7d\xe2\x9e\x64\x53\x3b\x5f\x8a\xd8\xfd\x33\xa3\x1a\xbb\x55\x32\xb8\x66\x02\x66\xaa\xfa\x5e\xbf\xdb\xad\xa7\xa0\x32\x04\xf7\x73\x39\x9c\x46\x68\x26\x57\x2c\xb0\x94\x74\x70\x92\x44\x56\x07\xe6\x54\x89\x66\xed\x96\x91\xa2\x9d\x0d\x98\x9f\x67\x67\x85\x58\x67\x67\xe6\xe8\x85\x0b\x8d\x30\x63\x62\x06\xdd\xa7\x3d\xa0\xb1\x4d\x2a\x83\xce\x95\x01\x03\x65\x1b\x34\x43\x13\x10\x94\x5a\xe6\x11\xf4\x73\x30\x6f\x86\xe7\x9b\xb6\xb3\xe3\xe6\x3c\x86\x39\xc6\xb5\x53\xfe\xcf\xe8\x39\xc9\xc3\xc2\x1e\x64\x15\x71\x66\x8e\x77\x4c\x00\x7e\x57\xf4\xba\x1e\x4d\x56\xb4\x58\xd7\xdf\x55\x7a\xab\x68\xcb\x7c\x06\xdd\xed\xc6\x46\x7e\xa1\x9c\x42\x35\x46\x31\xaf\x85\x0c\x68\x14\x11\xfe\xaf\xb2\x33\x74\xcc\x11\x94\xc7\x23\x22\xaa\x12\x86\xe7\xef\x6c\x85\xf5\x86\x68\xb3\xfa\xbe\xe7\xb4\xea\x7e\x3c\xa7\xd8\xa9\x2a\xb2\x4c\x31\xa3\xc6\xa8\x1e\x79\x20\x70\xe5\xfc\x03\x96\x40\xa7\xf3\xc0\x73\x98\x8a\x73\xe2\x59\xb3\x92\x2f\xdc\x95\xf2\x4d\xa1\x10\xfb\x5e\x83\xa9\x69\xe7\x28\x7f\x7c\xfb\x88\x4f\x5c\x77\xa3\x5a\xa6\x26\x7f\x72\x32\xa9\x09\x97\x36\xb6\x75\xf2\x53\x33\x6e\x3f\x56\xdf\x46\xea\x0d\xb7\xde\xcd\xc2\xed\x2a\x8d\x36\x6b\xb2\x8a\xd0\xeb\x31\x9d\x2f\x36\x27\xdd\x7f\x1a\xf6\xd8\x3b\x75\x33\x80\x27\x13\x0b\x89\xc2\x7c\x5f\x72\x72\x72\x2e\xda\xf2\xb8\xef\x3a\x3a\xaf\xa3\x89\x08\x23\x9a\x34\x00\x7a\x69\x07\x0d\x0c\xec\x55\x77\xb2\xac\xe4\xea\x31\x7a\x17\x49\xfa\xa6\x5e\x53\x2d\x3d\x57\x89\xbd\xa4\x56\xae\xde\x7d\xfc\xd5\x4b\xe4\xbb\x88\x41\xdf\x5a\x4d\xe6\x7d\x30\x01\xd5\x68\xca\xeb\x85\x95\x23\xe3\x6d\x84\x56\x15\xd5\x4a\xe1\x55\x71\x31\xd7\x04\xa6\x83\x29\x4e\x72\x57\xf0\x55\x43\x5e\x54\x4b\x7f\x9b\xb3\xd6\x5b\xa4\x87\xb8\x61\x54\x51\x7e\x43\xd5\x44\x59\xb4\x9a\xfb\x3c\xb8\xd0\xad\x9a\xe2\x9a\xa1\x6b\x84\xcb\x7d\xfd\x36\x42\xb7\x8a\x6a\xb5\xd0\x6d\xd6\x3d\x74\x40\x74\x94\x2f\xec\xab\x60\x4b\xfc\xa0\xf8\xcf\xae\xfb\xf2\xdb\xb0\xd2\x64\xdb\x34\x29\x58\x4d\x12\x94\x10\x1e\x19\xe1\xe9\xa2\x9d\x57\x46\x3d\x20\xa4\x14\x50\x99\xab\x07\xee\xd2\xc8\x05\x3e\xbc\xe4\x90\xa2\x74\x61\xe8\xc5\x95\xc9\x7b\x03\x1e\xc5\x39\x42\x7e\xe0\x00\xac\xfc\x27\x11\xd9\x79\x8a\x6b\xda\x62\x66\xdd\xac\x5e\xa3\x2e\xc8\xd1\x6c\x42\xc0\x86\xc5\xbb\x09\xa6\x65\xec\x15\x97\xbf\xe6\xe8\x02\x94\xd7\x46\xef\x65\x04\x43\x86\x51\x6e\xea\x60\x6e\x2f\x61\x0a\x1f\x6f\x6c\xb2\x28\x0f\x59\x1a\x91\x57\xf6\x60\x3d\xf2\x1d\x53\x54\xaf\x6e\x6b\x88\xf7\x33\x16\x16\x50\x76\x40\x1f\x2d\xd1\x6f\x5a\x24\x64\x84\x53\x06\x3e\xca\xd9\xdc\x92\xc1\x08\x9d\xba\x32\x1a\x2f\x95\x44\xa7\x92\x97\x2f\xa7\x91\xbd\xd0\xb5\xd0\xa0\xb3\x67\x6a\x79\xab\xa4\x34\xc8\xee\xc3\x4b\xb9\x26\x63\xa7\x21\xaf\xe1\x3d\x32\x33\x41\x83\x23\x45\xf4\xa6\xd1\xe9\x28\x65\x08\xfa\x67\xd0\xe6\x18\x54\xd7\x86\x38\xcd\xe0\x82\xa6\xcf\xba\x67\xab\xc1\x89\xf1\xc5\xb5\xc0\x1c\xe2\x0b\xb3\x2e\x41\x8a\x7e\x26\x57\xa2\xd9\x59\x15\x4e\x0e\x89\xc6\xe6\xf8\xac\x9f\x5d\xd5\xb8\xce\x98\xbd\x95\xb3\x6e\x82\x31\x6f\xf4\x08\x59\x5c\xa9\x67\x3c\xee\x79\x5b\xb0\xef\x3c\xd9\x98\xad\x8b\xd7\xf9\xb2\x78\x83\xc8\xb7\x8f\xd5\x44\xc5\x48\x25\x24\x34\x29\xb4\x2c\xe9\x43\x10\xf1\xea\xdd\xe5\x23\x89\xc5\xd4\xec\x2e\x96\x7e\xb3\xfc\x72\x88\xfb\x2a\xc4\x8c\x32\x06\x7b\x05\x62\x82\x83\x6b\xa2\x09\x36\xb3\xf2\x77\xa1\xbc\x2f\x40\xb5\x57\xdb\x70\xf3\x95\x5b\x7f\x05\x00\x00\xff\xff\x72\x44\xd0\x0d\xc7\x26\x00\x00")

func usersRamlBytes() ([]byte, error) {
	return bindataRead(
		_usersRaml,
		"users.raml",
	)
}

func usersRaml() (*asset, error) {
	bytes, err := usersRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "users.raml", size: 9927, mode: os.FileMode(420), modTime: time.Unix(1457621611, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"companies.raml": companiesRaml,
	"contracts.raml": contractsRaml,
	"organizations.raml": organizationsRaml,
	"securitySchemes/oauth_2_0.raml": securityschemesOauth_2_0Raml,
	"userorganizations.raml": userorganizationsRaml,
	"users.raml": usersRaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"companies.raml": &bintree{companiesRaml, map[string]*bintree{}},
	"contracts.raml": &bintree{contractsRaml, map[string]*bintree{}},
	"organizations.raml": &bintree{organizationsRaml, map[string]*bintree{}},
	"securitySchemes": &bintree{nil, map[string]*bintree{
		"oauth_2_0.raml": &bintree{securityschemesOauth_2_0Raml, map[string]*bintree{}},
	}},
	"userorganizations.raml": &bintree{userorganizationsRaml, map[string]*bintree{}},
	"users.raml": &bintree{usersRaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
