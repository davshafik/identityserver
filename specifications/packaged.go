// Code generated by go-bindata.
// sources:
// specifications/api/companies.raml
// specifications/api/contracts.raml
// specifications/api/organizations.raml
// specifications/api/securitySchemes/oauth_2_0.raml
// specifications/api/userorganizations.raml
// specifications/api/users.raml
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

var _companiesRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x4b\x73\xdb\x36\x10\xbe\xf3\x57\x6c\xdc\xe9\xcd\xa6\x48\xd9\xc9\x78\x78\xaa\xe3\xa6\x89\x5a\x3b\xe9\xd8\xf5\xc1\x93\xf1\x44\x10\xb1\x92\xd0\x90\x00\x03\x80\xb2\x98\x34\xff\xbd\x4b\x82\x12\x1f\xa6\x62\xd9\x3d\x94\x33\x3a\x08\xfb\xc0\xb7\xdf\x3e\xb0\x3f\xfd\x7c\x75\x76\x79\x01\xa1\x1f\x78\x56\xd8\x04\x23\x98\x58\x73\xab\xf2\x0f\x32\x11\x12\xbd\x15\x6a\x23\x94\x8c\x20\xf0\x43\x6f\xc6\x0c\xde\x68\x11\xc1\xc8\x33\x18\xe7\x5a\xd8\xe2\x3a\x5e\x62\x8a\x26\xf2\x00\x8e\x40\xb1\xdc\x2e\x3f\x8d\x3f\x05\xe5\x5f\xf7\xbd\x10\x32\x4e\x72\x8e\xd0\x33\x18\x6d\x75\x7d\xcd\xd2\xc4\xb3\x45\xe6\xbc\x9c\xab\x34\x63\xb2\x70\x1e\x32\xad\x32\xd4\x56\x38\x91\xfb\x16\x89\x9a\xb1\x44\xf0\xe6\xa4\xfc\x4a\xfb\x08\x8c\xd5\x42\x2e\x3a\x82\x54\xc8\x0b\x94\x0b\xbb\x8c\xe0\xb8\x2b\x60\xeb\x8d\x20\x7c\x19\x6c\x45\x59\x3e\x4b\x44\xfc\x07\x16\x66\xf7\x05\x1f\xef\xfa\x9e\x26\x16\x53\x13\xc1\xb8\xf1\x83\xeb\x4c\x68\x1c\xf2\xc1\x99\xc5\xce\xb1\xc6\x2f\x39\xe9\xf2\x08\xe6\x2c\x31\x8d\x4c\xe9\x05\x93\xe2\x2b\xb3\x94\x80\xe7\xa0\x09\x83\x60\xaf\x7b\x84\x9c\xab\xff\x16\xec\x8f\xbc\x5b\xb6\x96\x7a\xcf\x64\x0d\xf9\xc0\x35\x4b\xb3\x04\x07\x0a\x00\x26\x17\xe2\x33\x9e\x33\xc9\x8b\x47\xd3\x77\x04\xe1\xe9\xf5\xdb\x77\xb7\xd7\xe2\xf4\xf7\xf3\xd5\xd9\xea\xaf\x9b\x57\xb7\xe9\x2a\x7c\x77\x95\xa6\xf9\xe9\xab\xcb\x2f\xe3\xdb\xec\x6b\x3f\x77\x14\x61\x78\x7a\x14\x06\x47\xad\x48\xfb\x5c\x1d\x55\x27\xbf\x88\x84\x90\xc4\x25\x12\x3f\x56\x69\x2f\x76\x78\xfd\x26\x1c\x1f\x9f\xb4\x7f\x1e\xa9\xc4\xae\xd2\x57\x02\xef\x1b\x8f\x8e\x9a\xba\x09\x3c\xd7\x66\xc8\x5f\x17\x11\x7c\x6c\xda\x0b\xee\xbc\x91\xb3\xae\x3b\x23\x53\xc6\x3a\x1f\x1c\x4d\xac\x45\x66\xab\x9e\xbd\xc2\x85\x30\x16\x35\x30\x90\x78\xbf\xb9\xb0\xd2\x9b\x29\x5e\x34\xb7\xb2\x2c\x23\xd6\xaa\x42\x1b\xfd\x6d\xc8\x74\x20\x5d\xe7\x5b\xeb\x05\x6e\x2e\x13\x26\x4b\x58\xf1\x9e\xa5\x24\x7f\x8b\xb6\x56\xb9\xa0\x4b\x1f\x82\xf9\xa7\xe5\x93\x74\x61\x1b\x80\x0f\x67\x14\x97\xd2\x75\xa5\x43\x22\x52\x61\x0d\x30\x8d\x0e\x17\x72\xb0\xaa\x2a\x0d\x34\x96\x2a\x06\x72\x83\xda\xf7\x5c\xbd\x98\x8c\x7a\x03\x3b\xc9\x1e\x07\x41\x17\x7f\x37\xd6\x7d\x22\x7e\x10\x75\xd5\x04\xa3\x6f\xae\xf4\x26\xfc\xbb\x53\xdf\x12\x01\xd0\x4e\x54\x33\x06\xe1\x1b\x98\x58\x95\x93\x8d\xd2\x77\x50\xf3\x1f\x69\x64\xfc\xe0\xb0\xf9\xcf\x38\x0d\xa9\x03\xb8\x83\xef\xb0\xe9\xb5\x0e\x73\x25\x5b\xed\x59\x50\x95\x5c\xad\x38\xc0\x40\x2f\xfe\x87\xd1\x3f\x16\x7b\x3f\xdf\xee\x3b\x09\xc2\xb6\x62\x07\xe1\x8d\x64\x75\x0a\x91\xb7\x0c\x4e\x76\x1a\xbc\x57\x16\xe6\x2a\x97\x4e\x3b\xcb\x9f\xc5\xe3\x63\xbc\xdd\x64\xe5\xa8\xa5\x6e\x16\xae\x6e\x6a\x3b\xdf\x09\xca\x93\xe9\x74\x93\xd1\xe9\x14\x84\x01\x49\xb0\x58\x92\xa8\x7b\xe4\xfe\xff\x40\xf0\xf1\x6e\x82\x37\x88\x37\x78\x4b\xb4\xbf\x29\x3d\x13\x9c\xa3\x7c\x51\x99\x8d\xda\xb3\xe9\x29\x3c\x96\x76\x5d\x1a\x5b\x85\xbd\xa3\xc9\x06\xda\x6c\x98\x8b\x7d\x39\xe9\xb3\xd3\x9a\x8e\x2e\xba\x55\x39\xf3\x29\x9f\xd1\x10\x46\x9a\x0d\xba\xf8\x93\xd1\x2a\x81\x34\xef\xfa\x8f\xa5\xfa\x8c\x3b\x1b\xbc\x7e\x85\xdc\x1d\xb1\x92\x56\xb3\xd8\x9a\xc1\x4b\x9e\x5e\x9b\xad\x26\x6f\x5c\x57\xed\xdf\xa1\x7b\xa0\xdf\xed\x12\x61\x6b\x02\xf7\x4b\xa4\x69\x58\x9e\x75\xe7\x80\x81\x10\xd4\xbc\x12\x64\xac\xda\x93\x7c\xf8\xa0\x39\x4d\xfc\xd2\x21\x4a\x5e\x96\xcc\xac\xa8\x56\x0e\x7f\x1f\xb2\xea\x6d\xed\x4d\xf5\x00\xf2\xe1\xba\x9d\x29\x95\x20\x93\x3d\x59\x27\x82\x49\xbd\xf4\x95\xc0\xdc\x63\xca\x9b\x68\x0e\x2b\x44\x38\x67\x79\x42\x73\x4d\x26\x45\xa5\x46\x12\xb1\x6a\xc7\x5c\xce\x7f\x8d\x36\xd7\xb2\x69\xc6\xcd\xb7\x6b\xd7\xa0\x1c\x59\xe2\x61\x18\xb7\x90\x16\x17\xa8\x7f\x84\xfb\xba\xb4\x26\x46\xe7\x06\xed\x61\xf9\xce\xcc\xf3\x84\xe6\x94\x26\x76\x17\x44\xa5\x0f\xbf\xd6\xb0\x89\xf9\x69\x30\xdd\x1f\x16\x6d\x4d\xcf\x06\x75\xc9\xd6\xe5\xfd\xb4\x42\xd3\x78\x7d\x14\xd5\xcb\xa7\xc0\xaa\xa1\x89\x34\x4f\x69\xdb\xa1\x2d\xf8\xdf\x00\x00\x00\xff\xff\x7b\x0f\xa7\x77\x0e\x0c\x00\x00")

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

	info := bindataFileInfo{name: "companies.raml", size: 3086, mode: os.FileMode(420), modTime: time.Unix(1458554786, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _contractsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x55\x5d\x6f\xeb\x44\x10\x7d\xf7\xaf\x18\x2e\x42\x80\x54\x12\x27\xea\xbd\xba\xf2\x0b\x2a\xa1\x85\x88\x16\x50\x3f\x1e\x50\x55\x55\x9b\xf5\x24\x5e\xea\xec\x9a\xfd\x68\x63\x4a\xff\x3b\xb3\x6b\x7b\x6d\x27\x94\x0a\x90\xc8\x43\x55\x8f\x77\xe6\x9c\x39\x73\x76\xfc\xe9\x67\x97\x27\x17\xe7\x30\x9b\xa4\x89\x15\xb6\xc4\x0c\x96\xd6\xfc\xa2\xdc\x4f\xb2\x14\x12\x93\x47\xd4\x46\x28\x99\x41\x3a\x99\x25\x2b\x66\xf0\x46\x8b\x0c\xa6\x89\x41\xee\xb4\xb0\xf5\x15\x2f\x70\x8b\x26\x4b\x00\xbe\x02\xc5\x9c\x2d\xee\xe7\xf7\xa9\x7f\x6c\x7e\x9f\x08\xc9\x4b\x97\x23\xec\x25\x4c\xe3\xd9\x89\x66\xdb\x32\x49\x6c\x5d\x35\x65\xae\xc4\x46\x32\xeb\x34\x36\x45\x2a\xad\x2a\xd4\x56\x34\x2f\xfd\xcf\xd0\x01\xcc\xbf\xa9\x33\x30\x56\x0b\xb9\x69\xc3\x39\xb3\x44\xde\xff\x6d\x03\x95\x5b\x95\x82\xff\x80\x75\xcf\xc6\x83\xec\xa5\x51\x22\x1a\xae\x45\x65\x43\x9b\xd7\x05\xb6\x89\xf0\x80\x35\x38\x83\xf9\x00\x76\xc0\xeb\xa0\x1c\x45\x17\x4a\x5a\xcd\xb8\x7d\x8d\x79\xc5\x46\x8f\xe3\x02\xb7\x77\x31\xbc\x15\x72\x69\x71\x6b\x32\x98\xf7\x31\xb6\xeb\x62\x69\x0c\x3a\x29\x7e\x73\xd8\xc6\xad\x76\x5d\xeb\x9c\x88\xa0\xb4\x7b\xad\xf2\x96\xde\xb5\x47\xfd\x7b\x4d\x08\xed\x1c\xe5\xc6\x16\x19\x1c\xa7\x7f\xad\xd4\x99\x46\x84\xb5\xc0\x32\x07\xab\x60\x85\x41\x2b\x58\x2b\x0d\x96\x6d\x0c\xac\x6a\xb0\x5e\x4b\x6a\xb9\x06\xae\x91\x59\xaa\x0f\x2c\x92\x68\x8b\x0a\xf9\xc8\x4a\xe1\xa7\x66\xbe\x7e\x5b\x97\xa8\xc1\xec\x15\x52\xa5\x30\x16\xd4\x3a\xa2\x2c\xf3\xcf\x0d\xf1\x10\x26\x46\x86\x88\x93\xb6\x08\xee\x48\xad\xfc\x7f\x23\x60\x80\xc9\x06\xd3\x5f\x2d\x70\x95\x92\x91\x48\x25\x34\x19\x64\x68\xe3\xbe\xd2\x3f\xf1\xf1\x1f\x31\x0c\xc1\xd3\x7d\x15\x8f\xef\x27\x53\x30\x53\x78\xa6\xdd\x94\xc8\x98\x47\xe1\xa1\x35\x4f\xff\xd0\x79\xe6\x08\x88\x78\x1e\xc2\x81\x28\xf3\x50\x81\xea\x64\x88\xa6\x80\xb3\x92\xbb\x92\xe2\x11\x88\x8a\xb1\x87\xe6\xf1\x57\x43\x49\x1a\x2b\xea\x93\x50\x9a\x1a\x2d\x8d\x28\xd1\x93\xb0\x85\x72\x36\x04\x45\x1e\x50\xe3\xf5\xf3\x2c\xb5\xd8\x86\x77\x12\x9f\xfc\x96\x32\x40\xae\x33\x15\xe3\xf4\xdf\x17\xb8\xe3\x58\x51\x89\x02\xa5\x3f\x53\x03\xab\x2a\x64\x9a\xe6\x1e\x52\x68\xf8\x0e\xcd\x97\xa1\x26\x4a\xae\x68\x35\xd1\x1b\x70\x76\xfd\x71\xd8\xc5\x99\x56\xdb\x66\x6e\x84\xe8\x4a\xdb\x36\xc0\xe0\xea\xfb\x93\xf9\xfb\x0f\xe4\xf3\xb2\x54\x4f\x64\x78\xf2\x39\x83\xcb\xe5\xcf\xa7\x17\xdf\xce\x3e\xa4\x47\x44\x69\x67\xc1\x2f\xca\xf7\x1f\x17\x05\xf2\x87\xd3\x06\xc3\x43\x37\x95\x42\x73\x94\x93\xee\xd2\x94\x76\x04\xae\xc5\x6e\x92\xec\xef\x18\xb2\x40\xdc\x83\xe4\xbf\x66\xdd\x36\x4b\xef\xb6\x5f\xb3\x70\x97\x4c\x3b\xcd\xc2\x56\xa9\x94\x69\x97\xcf\xc8\x0c\x0b\x7f\xfd\x3c\x79\xd2\x2b\x8a\xdc\x74\x4b\x50\x64\x3f\xd3\x2f\xa5\x79\x3a\xeb\x7d\xb6\x52\xf9\x60\x7b\x82\x97\x92\x36\x63\x98\xd9\xd4\xcf\x71\xf8\xae\x73\xe5\x62\x7c\xc1\x8f\x87\xf5\x46\xac\x6e\xa4\xef\x43\x69\xf1\x3b\xed\x58\x3a\x32\x7d\xee\x4d\xfa\xd2\xe4\x6c\xd0\xc6\xad\x3f\x50\xa0\xff\xce\xc0\x33\x18\xae\xfc\x97\x83\x74\x79\xd7\xe5\x67\xd4\x6f\xfe\x0e\xee\xe0\x05\xba\xbb\x3b\x42\xfe\x0e\xed\xe1\x26\x3a\x50\xc2\x6b\x91\x0e\x3b\xdc\x57\xe3\x2d\x3d\x5e\x51\xc4\x6b\x72\x3c\x3c\x39\xe2\xf6\xa3\xb2\x64\x2e\x27\x9b\xcf\xce\x74\xe0\x88\x98\xd1\x4f\xf9\x5f\x48\x13\xae\x3a\x17\x15\x93\x76\xac\xd0\x01\x13\x6f\xc0\x43\x99\xfe\x8b\x10\xd1\xd2\xc9\x9f\x01\x00\x00\xff\xff\x1f\x14\xe6\x48\x76\x08\x00\x00")

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

	info := bindataFileInfo{name: "contracts.raml", size: 2166, mode: os.FileMode(420), modTime: time.Unix(1458557466, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _organizationsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5a\xdf\x6f\xdb\x38\x12\x7e\xf7\x5f\xc1\xcb\xe1\xde\x12\xff\x48\xdd\xee\xc5\x2f\x87\x5c\xb3\x28\x82\x4b\xd3\xa2\x69\xb7\xb8\x2d\x8a\x2d\x25\xd1\x36\x37\x34\xa9\x92\x54\x6c\x6f\xd0\xff\xfd\x86\xa4\x64\x49\x14\x25\x2b\x89\xb7\xb7\xa8\x5f\x02\x5b\xe2\x70\xf8\xcd\xc7\x6f\x86\xc3\xfc\xfd\x1f\xef\xce\x5f\x5f\xa1\xc9\x70\x3c\xd0\x54\x33\x32\x43\x97\x5a\xfd\x57\x64\x6f\x38\xa3\x9c\x0c\xee\x88\x54\x54\xf0\x19\x1a\x0f\x27\x83\x08\x2b\xf2\x41\xd2\x19\x1a\x0d\x14\x89\x33\x49\xf5\xf6\x26\x5e\x92\x15\x51\xb3\x01\x42\x27\x48\xe0\x4c\x2f\x7f\x3b\xfd\x6d\x6c\xbe\xba\xcf\xdf\x28\x8f\x59\x96\x10\xe4\x0d\x18\xed\xde\x1d\x4a\xbc\x62\x83\x81\xde\xa6\xce\xcc\x1b\xb9\xc0\x9c\xfe\x81\xb5\x99\xd6\xda\x49\xa5\x48\x89\xd4\xd4\x3d\x37\x9f\x05\x13\x11\x66\x34\x29\xe7\x31\xc3\x67\x48\x69\x49\xf9\x62\xf7\xe3\x8a\xf2\x2b\xc2\x17\x7a\x39\x43\xcf\xca\x1f\xf1\xa6\xf8\x71\xf2\x7c\x9c\xff\x9c\x66\x11\xa3\xf1\x7f\xc8\x56\x85\x4d\x7e\xfa\x5c\x1d\x7f\xa9\xc9\x4a\xcd\xd0\x69\x31\x3a\xe1\x0f\x19\x36\x19\x8f\x77\xbf\x26\x44\xc5\x92\xa6\x76\xad\xe8\x8a\x2a\x8d\xc4\x1c\x89\x0a\x02\xe8\xe2\xfa\x66\x98\xbf\x2e\xd6\x1c\xa2\xf1\x18\x07\x5b\x26\xca\x14\xd8\x43\x5f\xcc\x1f\x8e\x21\x28\x5f\x8a\x99\x56\x64\x15\x3d\x74\xaa\x7d\xab\x72\x93\xad\x97\x02\x61\x49\xf2\x19\x10\xe5\x48\x2f\xa9\xaa\x2d\xb9\x70\x22\x67\xce\x41\xa0\x75\x84\xb9\x4c\x8c\x23\x2a\x8b\x6a\xf3\xa9\x81\x1d\x43\x36\x78\x95\x02\xff\x7d\x8a\xa1\x85\x24\x84\x53\x6d\x7e\x20\x1d\x74\x39\x41\x93\xf3\xb3\x8f\x1f\x97\x2f\xe8\xf9\xcf\xd3\x77\xbf\xbe\xba\xfe\x09\x6f\xb6\x2f\x36\xbf\x46\xd9\xc7\xec\xea\x96\x7f\xfd\x78\x25\x7f\x09\xf1\xe5\xa4\x36\xc1\x30\x16\xab\x96\x70\x9f\x58\x08\x27\xde\xf7\x53\xef\xfb\xb3\xb6\x10\xba\xc7\xd3\x56\x70\x4f\x10\x4e\x60\xc3\x0c\x1b\xde\x78\x7b\xf2\x3d\x3c\x37\x70\x3f\x69\x6f\xc6\x4b\xca\x12\x49\x20\x32\x21\xd3\x10\xdb\x41\xb1\x82\xb6\x69\x0a\xd2\xee\x91\x80\x1a\x0b\x3e\x28\x92\x00\x03\x09\x47\x58\x29\xba\xe0\xf0\x22\xc2\x05\x15\x35\x10\x93\x7b\x44\x0c\x11\x63\x37\x31\x8a\x44\x64\xde\xb8\xe4\x77\x54\x77\xca\x95\x1b\xd4\xf0\x4d\x0a\x56\x71\xbf\x65\x01\x30\x3f\xcf\x56\x33\xf4\xc9\xd2\xe1\x38\xf7\xb6\x24\x7f\x2c\x09\xd6\x04\x68\x9a\xc0\x9f\x56\x87\x9d\xb3\x95\x69\x1d\xbb\x06\x9e\x8d\x9b\x8c\x1f\xa3\xf1\x0b\x74\x2d\xee\xd0\xe4\xec\x6c\x8a\xc6\xff\x9c\x4d\xcf\x66\xcf\x7e\x42\xaf\x5e\xbf\x37\xc6\xcf\xdf\x5e\x02\xeb\xdb\xd6\xc9\x70\x44\xd8\x3e\x4d\x2e\xe5\xf7\xf9\x38\xa4\xd4\x05\xa1\x63\xcc\x58\x84\xe3\xdb\x0f\xef\xae\xfe\xd5\xdf\xe6\x69\xd8\xe8\xf3\xc2\x28\xa3\x84\xeb\x97\x92\x24\xf0\x87\x62\xa6\x5e\x49\xcc\xf5\x7b\x30\x5a\x99\xa3\xc6\x99\x4b\x9e\xd0\x18\xd0\x51\x88\xce\x9d\x56\xdd\x92\x2d\xcc\xb8\x45\x11\x31\xd0\x26\x46\xc3\x70\x6e\xd8\x20\x59\x58\x76\xd9\xf0\x14\xcd\x99\x58\x0f\x3d\xff\x23\x01\x31\xc0\xbc\x32\xe3\x1c\x67\x4c\xcf\xd0\x1c\x06\x16\x22\x03\x19\x53\x12\xfd\xc0\xb5\xbb\xc4\x4c\x92\x7f\x6f\x81\x32\x65\x42\x46\x9f\x07\xa3\x9a\xe6\x19\xab\xa9\x50\xda\x59\x4f\xa8\x4a\x19\xde\x5e\x5b\x5a\xbf\xb4\x6c\xb8\x26\xeb\xea\xd6\x1c\x34\x80\x71\xaf\xc1\xd2\x39\x59\xd7\xb7\x0d\x9a\x58\xce\x21\xb5\x14\x19\x4b\x0c\x4e\x56\xe6\x49\xae\x68\x88\x41\x46\x18\xa2\x5f\x8c\x48\xb8\x1c\x07\xa0\x02\x93\xe6\x42\xae\x00\x4e\xd8\x87\x50\x22\xc4\xb7\x0e\x6f\xbf\x70\x40\xc0\x0a\xb1\x56\xb0\x6c\x8e\x17\xf0\x1d\x30\x17\x79\x0e\x31\x73\x3a\xa0\x23\x91\x6c\x0b\xd8\x70\x9a\x32\x13\x40\x98\x67\xf4\xbb\x2a\xf6\x68\x09\x67\x63\x91\x92\xa8\x14\x00\x2a\x59\x7d\x3a\x9e\x94\x83\xaa\xa6\xbb\xcd\x77\x4c\x81\xd0\xb4\x6a\xb3\xae\x51\xdc\xc4\x4c\x48\xfa\x07\x49\xe0\x8d\xd1\x7d\xa1\xa7\xdf\xdc\x80\x05\xd1\xb3\x92\x20\xbb\x50\x97\xa5\x17\xba\x47\x2a\x16\xa6\x98\x02\x02\x1c\x55\x03\x33\x73\xda\x71\x74\xec\xfd\x6c\xc3\x72\x84\x3e\xa3\x6f\xa8\xd0\x95\x9a\x4b\xaf\x88\xae\xd7\x24\x94\xcf\x45\x21\x26\x3e\x5a\x06\xaf\x71\x15\x05\x1f\xb1\x7d\x98\x75\xa0\xe6\xe1\xd6\x8d\x5c\x31\x60\xda\x3a\xe0\x5a\x68\x34\x17\x19\x77\x6f\xa7\xd9\xe3\x81\xdd\x87\xe0\x87\xd4\xa8\x73\x2b\x88\x75\x88\xba\xe0\x69\x85\xe6\x07\x8a\xc3\x4e\x95\x0e\x14\x88\xa0\xb6\xdd\x64\x51\x60\x61\xed\x02\x07\x45\x63\xa8\x46\x3d\x90\x1c\x04\xa2\xe7\xe9\x4e\x38\x82\xfb\x63\xd8\x19\xc5\x46\x1c\xf7\x45\x72\xe4\x55\x95\xd5\x58\x1d\x26\x5a\x9e\x07\xe7\xb6\x4e\xab\x15\x69\xa1\x30\x3c\x2e\x10\xce\x66\x59\x8e\xf5\x0a\x42\xcd\xbd\xd7\xce\x2b\x57\x4d\x42\xee\x52\x59\x1c\x13\xa5\xe6\x19\x63\xdb\x83\x85\xce\x73\xf3\xc1\x41\x2b\x06\x4d\x3b\x06\x95\x5b\x30\x7f\x67\x74\x5f\xd4\xb9\xdf\xaa\x89\x8a\x11\x5d\x2b\x59\x0f\x13\x72\xcf\x99\x77\x64\x25\xee\x48\x19\xf4\xb9\x14\xab\x5a\xd8\x2b\x03\x83\x41\x33\x61\x9b\xfa\x70\x86\x02\xe7\x16\xd4\x11\xb7\x06\xd4\xfd\xc0\x0e\xc0\xed\x0d\x7c\xbf\x24\xae\x48\x12\xd2\x95\x46\xd5\xd4\x90\x08\x28\x37\x39\x84\x84\x6c\x4c\xa9\xe4\x19\x3e\xeb\x67\x18\x33\xd0\xaf\x64\x6b\x8a\xab\x02\xc9\xfc\x40\x33\xaa\x9f\x2c\xff\xf4\x4d\x6c\x0f\x48\x26\x9e\xd6\x2f\xd8\xc2\x11\x81\x73\x65\x5e\x0e\x9a\x13\x79\xe3\xdc\x55\x98\x79\x6c\xce\xea\xb3\xb3\x03\x7b\x3b\xec\x77\x7e\x3a\x6a\x27\x49\x78\x7b\xf7\xd9\xe0\xad\x5b\xfc\x87\x61\x1e\x77\x51\x1e\x76\xea\x4a\x48\x59\x0e\xa7\x2d\x2d\xea\x92\x7b\xd6\xa9\x2e\xad\xd4\x09\x2a\x8c\x37\xd1\x1b\x6b\x5e\xda\xe9\x3a\xd9\x13\x08\x75\xdf\x60\x07\xc3\xdd\xad\xed\xa3\x58\x70\x2d\x71\xac\xd5\x13\xab\x2c\xff\x18\x51\xda\x35\xe1\xaf\xc7\xa1\x72\x6a\xf1\x2a\x32\x38\x5c\xbc\x2c\x06\x86\xe5\xc3\x1c\x3f\x0c\x51\x77\xf6\x4d\x0b\x47\x92\x26\x79\x81\x6f\x13\x23\x27\xe6\x41\x8a\x6d\x67\x62\x08\x85\x4f\x62\x85\x5e\xc5\x04\x4e\xf2\x7c\x81\xa2\xad\x6d\x95\x94\xdc\xfe\x9a\x11\xb9\x7d\x8b\x25\xb8\xa3\x6b\x1d\xb3\x5d\x97\xec\xe7\x4d\x4a\x01\xa0\xb0\xd6\xf8\x67\xf9\xc0\x0a\x2e\xf3\x1e\xb8\x71\x8c\x38\x5b\xe5\x6a\x8e\xad\x47\xae\x01\x00\x67\x59\xb6\xb5\xaf\xc1\x13\x7a\x57\x5d\xb3\x69\x9a\x4a\xa2\x33\x09\x75\xc6\xd0\x9b\x4c\x92\xaf\x99\x75\xb0\xd6\x3f\xb0\xa1\xd5\x80\x43\xd8\x6f\xca\x35\x59\x78\xa2\xe3\xf9\x7d\x63\x46\x03\xa2\x73\x45\xf4\xb1\xd9\xdb\xc0\x5e\xa0\x92\x04\x74\x17\x00\xe5\x10\x5d\xe4\x6e\x03\xf2\x5f\xc6\x5f\xfa\xbb\xb5\xc2\x9b\x47\x3b\xf5\x1a\x6f\xcc\xfc\x04\x29\xd8\x0d\x7b\xbd\x7a\xfe\x10\xb7\x72\xd7\xe8\xca\x34\xdc\x6c\x37\xc5\xfc\x32\xa2\xbb\x06\xdf\x53\xb7\x4c\xff\x3d\xf1\xd6\xb1\xb5\xec\x2d\xee\xd9\x1c\x2c\xef\xb2\xa7\x39\xcb\x2b\x3e\x5b\x6c\x5c\xf7\x1d\x72\xef\xef\xa2\xab\xe7\xde\x91\x2b\xc7\x7e\xc4\x9e\x9a\xf4\xca\xb5\xb9\x7e\xaf\xc5\xba\x5f\xd9\xe9\xe1\xe5\x24\xbd\x01\x59\x7b\x0e\x78\x89\x79\x4c\x18\x54\x24\x4d\xb8\x86\xfd\xe4\xbf\x25\x01\x84\x2a\x08\x27\x4e\xb1\x9d\x92\x91\x42\x85\x71\x4a\x6f\x2b\xf7\x07\xf5\x53\xcf\xdb\x4b\xd3\x65\x74\x9b\xde\xca\x9c\x6b\x25\xe6\x4d\x46\xd7\x18\x34\xc9\x35\xd9\x75\x49\x21\x6b\x13\x92\xb8\xf6\x19\xb6\x99\xc6\xc9\x48\x4a\x87\xdf\x8b\xb2\xae\x37\x7c\x65\x7a\xc0\x3d\xc9\x9a\x8b\x1c\x38\x69\x97\xfb\x7d\x49\xd8\x72\x9d\x54\xaf\x86\xc3\xcd\x03\xb7\xd2\xf0\x1a\x6b\x5d\x03\x13\x48\x78\xf3\x18\xbe\xba\xa0\x21\xaa\x15\x61\xf3\xa2\x35\x6a\x2a\xad\x88\x98\x4e\xfa\x1d\x85\xe0\x1d\xc3\x63\xb4\xa6\x8c\x99\x1f\x17\x04\xe0\x77\x85\x27\x91\xf6\x1a\x36\x21\x4f\x3b\xef\x7a\x6e\xf7\xaf\x8a\x9f\x0e\xb2\x37\xb5\xf9\x04\xca\x49\xff\xe6\xd0\x50\xc9\x56\x91\x79\x41\x69\xda\xed\xc3\x9d\x52\xd8\xdb\x86\x8a\x4c\xd4\xc8\xd9\x46\xcf\x3e\x47\xc7\x06\xc9\xda\x10\xe8\x8f\x42\x17\x0c\xfd\xeb\xb7\xe2\x71\xa5\x55\xd9\x58\xa9\x6b\x35\x36\x26\x0a\xb4\x23\x9d\x42\x58\x14\xcd\x59\x40\xc0\x37\x59\xb9\xd4\xb1\x1b\xd4\xec\xcb\xaa\x22\x3e\xfc\x40\x16\xba\x0d\xdb\x03\x4a\x7f\x62\x06\xd7\xe5\x1f\x87\xf6\xd1\x0c\x36\x74\x8e\x82\x47\xb5\x8e\x04\x54\x83\xfc\xc2\x3e\xee\x86\xdc\xa5\x28\x7b\x20\xca\xd5\x7d\xff\x82\xfb\x1c\x32\x72\x63\xc5\x31\xc3\xe5\x96\xca\xfd\xb2\x57\x3a\x99\x4b\x13\x38\xe4\xad\x30\x14\x01\xf6\xca\x1f\xb4\x98\xdf\xe6\x69\x23\x7c\x0a\x3f\x44\xde\x18\xdd\x83\x4b\x5e\x56\xaf\x0b\x6d\x50\x6a\x2f\x2a\x55\x4f\x50\x66\x55\xae\xb3\x17\xd7\x37\x76\x39\xa6\x13\x27\x62\x6a\x55\x73\x4d\xf5\xd2\x5f\xd3\x9f\xc6\xe4\xfa\x4d\x74\xf9\x69\xb9\xd2\x2d\x3e\x95\x1b\xca\x69\xe8\x71\xf8\x5e\xb3\x95\x30\xfb\xb9\xbe\x03\xaa\x45\x55\xdd\xe7\xc0\x29\xa0\x0b\xb8\xfc\xe3\x2e\xd4\x3d\x9c\xf6\xcb\x1c\x10\xc4\x7c\xdd\xa7\x73\xc0\x02\xdb\xcf\x30\x95\xde\xff\x9d\x2a\x82\x25\xdf\x99\x2d\x39\xc4\x64\xfd\xd7\x60\x69\x8b\x92\x37\x04\xd3\x78\xdb\xec\x76\x04\x18\x1e\xd0\xf3\x36\x9e\x37\xcc\xed\xd3\x57\x63\xc8\x50\xcc\x96\x6a\xf5\x2c\xdc\x2b\x2b\xc0\xf8\x2e\x86\xee\xd2\xc2\xce\xe3\xc3\xa4\x85\xdd\xfa\xab\x79\xe1\x89\xeb\x76\xb9\x45\x4b\x52\xfe\x8f\x54\xd7\x71\xc0\xff\xc7\x9e\xee\xba\xb3\x51\x75\x3d\xee\xc6\xa4\x5b\x6a\x8e\x3e\x7d\x3e\x0a\xff\xc7\xd1\xe0\x7f\x01\x00\x00\xff\xff\x01\x64\xf5\x31\x07\x29\x00\x00")

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

	info := bindataFileInfo{name: "organizations.raml", size: 10503, mode: os.FileMode(420), modTime: time.Unix(1465211324, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _securityschemesOauth_2_0Raml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x92\x41\x8f\xd3\x40\x0c\x85\xef\xf9\x15\x56\x2e\x9c\x48\x96\xd5\x9e\x72\x5b\x84\xc4\x11\x0e\xf4\x84\x10\x9a\x9d\x31\xcd\x88\x64\x3c\xd8\x4e\x4b\x10\x3f\x1e\x4f\xd3\x54\x6d\xd5\x4a\x9c\xb6\xa7\xea\x79\xde\xf3\x67\xc7\x01\xc5\x73\xcc\x1a\x29\x75\xf0\xb7\x02\xf8\xf4\x3c\x69\xff\x08\x51\xc0\x41\x66\x52\xf2\x34\x80\xf6\x4e\x61\x40\x15\xc0\xdf\x8a\x9c\xdc\x00\x2e\x67\x01\xc6\x5f\x13\x8a\x82\x33\x0f\x71\xfc\xe3\x4a\x0e\x28\x99\x33\xee\x9c\xa2\xe5\x05\x54\x17\x07\x81\x98\x2c\x70\x12\xe4\x37\x96\xec\x3d\x4d\x49\x61\x1f\xcd\x36\x29\x6c\x51\x35\xa6\xad\xb5\xc1\xc8\x90\x9d\xc8\x9e\x38\x34\x95\xce\x19\xbb\x85\x08\x1e\x9b\x87\x6a\x81\x7d\xc1\xf0\x7e\xee\x2c\xba\x47\x17\x90\xa5\xfc\x05\x78\x3e\x47\x58\xa4\xd2\xfc\x6a\xba\xe5\xb7\x11\x0c\x85\x52\x30\x05\xa3\xda\xb9\x21\x86\xb5\x4d\x81\x43\x11\x2b\xff\xc4\xd4\xc0\x07\x82\x44\x5a\xc0\x4d\x31\xce\x1e\xf9\x80\x7d\xca\x32\x05\xea\xc5\xf3\xfd\xe0\xa9\xc1\x76\xc2\x33\x88\x72\x99\x29\x3b\x76\x23\xda\xd2\x1a\xb3\x1c\x2a\x9f\x57\xe5\x48\x7e\x6e\x7e\x6d\xf0\x8b\xad\xd5\xc7\x8d\xda\x13\x46\xc9\x94\x04\x8f\x88\x4f\x0f\x4f\x37\xc9\x36\x69\xfd\xf2\x18\x2a\x59\xbe\xe2\xc1\x72\x71\x10\x1b\x8e\x1d\xf4\xaa\x59\xba\xb6\x8d\x2a\x33\x4d\x0d\xa5\x21\x26\x6c\x77\xef\x5a\x2a\x6f\xdb\x53\x4e\xb5\x2e\xe4\x4b\x99\xe3\xbf\xac\x67\xfb\xbb\x6e\xfd\x91\x5d\x52\xe9\xe0\x2b\x78\x0a\x08\xdf\xac\x2e\x9e\xf2\x3a\xd7\x5b\xa8\xcb\x49\x76\x2e\x8c\x31\xd5\x17\x52\x4c\x3f\xe8\xa4\x10\x6f\x5d\x5a\x6f\x8b\xf6\x09\xf9\x76\x69\xc4\xf1\xe5\x5e\xed\x7e\xa0\xa7\xa4\xec\xbc\x71\xb2\xed\xff\xf4\xc6\xd3\x98\x5d\x9a\xaf\xe0\x56\xf5\xe6\xd3\x8b\x1e\xab\x78\x37\x7e\x91\xef\xa8\x76\xb7\x1a\x7d\xb4\x04\xad\xab\x7f\x01\x00\x00\xff\xff\x62\x8e\x61\x93\x26\x04\x00\x00")

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

	info := bindataFileInfo{name: "securitySchemes/oauth_2_0.raml", size: 1062, mode: os.FileMode(420), modTime: time.Unix(1459339155, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userorganizationsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\xc1\x6e\xdb\x30\x0c\x86\xef\x7a\x0a\xae\xc3\x6e\x5b\x9c\x14\x3b\xf9\xd6\x5d\x86\x0e\x1b\x0a\xb4\xd8\x61\x28\x82\x42\x91\xd9\x84\x85\x2c\x09\xa2\xdc\xc1\xcb\xf2\xee\x93\xac\xa4\x96\xb3\x34\x1b\x2f\x4e\x4c\xf2\xff\xc9\x8f\x7e\xfb\xee\xf6\xea\xdb\x57\x58\xcc\xe6\x22\x50\xd0\x58\xc3\x75\xe0\x1f\xb6\xbb\x31\x9a\x0c\x8a\x67\xf4\x4c\xd6\xd4\x30\x9f\x2d\xc4\x4a\x32\x7e\xf7\x54\x43\x25\x18\x55\xe7\x29\xf4\x77\x6a\x83\x2d\x72\x2d\x00\x3e\x80\x95\x5d\xd8\x3c\x5c\x3e\xcc\xd3\xdf\x1c\x6f\xc8\x28\xdd\x35\x08\x47\x0d\xd5\x4b\xed\xcc\xcb\x56\x8b\xd0\xbb\xa4\xb2\xef\xfb\x62\xc9\xdc\xf8\xb5\x34\xf4\x4b\x86\x68\x7f\x6d\x9e\x29\x0c\xbf\x46\x65\xe7\xad\x43\x1f\x28\x9b\x8f\x61\x8b\xbe\x1a\x38\x78\x32\xeb\x49\x41\xc7\xe8\x4f\x26\xbc\x8d\xeb\x4f\xde\xa4\x48\x93\x9d\x2c\x4f\x81\xa6\x6b\x6b\xb8\xb7\x3f\x0d\xfa\xf7\xd0\x62\xbb\x42\xbf\x14\x99\x0e\x36\x9f\xfa\x98\x1b\xa9\xc0\x52\x54\xc9\x9c\xab\x6d\x7a\x18\xd9\xe2\xae\x2a\xc7\x1d\x36\x29\x7b\x47\xa0\xb0\x05\x56\x36\x31\x8a\x8a\x17\xc3\x0a\xb2\x69\xc9\x5c\xc0\x12\x76\x51\x18\x60\x8d\x21\x0f\xdf\x20\x2b\x4f\x2e\xef\xff\x19\x03\x84\x0d\x82\x26\x0e\x13\x34\x0c\x72\x20\x01\xc4\x30\x8c\x1f\xb3\xfb\x05\xc0\x3e\x0e\x42\x1e\xd9\xc5\xc2\x11\xf0\xe5\xbc\xb8\xec\xca\x36\x7d\x49\x4b\x3a\xa7\x49\x0d\xda\xd5\x13\x97\x97\x3a\x77\x2d\xc8\xe6\x07\xc0\xf7\xcb\xa3\x6c\x9e\xa8\x48\xc7\x7c\xb5\x5d\x6b\xbb\x92\x9a\x9a\x5d\x95\x6e\x16\x79\xa6\xc7\x2e\x2b\x3b\xcb\xe1\xe0\x31\x41\x71\xa5\x14\xba\xb0\x57\xe4\x0d\x39\x20\x33\x41\x22\x4e\x2d\x76\x6e\xad\xfc\x69\xbc\xfe\xb1\xee\x4b\xff\xe2\x98\x48\x2e\x4a\xa1\x63\x96\xff\xa2\xf9\x9f\xd6\x0d\x6a\x0c\x78\x12\xc6\xef\x17\xbd\x5b\x7c\x42\x75\x84\xe5\xa0\x92\x08\xc9\x29\xa4\xd9\xb9\xa5\x3e\x96\x63\xbe\xe2\x97\xe2\xae\x8b\xb7\xe0\xc7\x4e\xeb\x3e\x0a\x25\x7f\x6c\x0a\xdb\x99\xf8\x13\x00\x00\xff\xff\xc7\x4d\xbb\x01\x96\x04\x00\x00")

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

	info := bindataFileInfo{name: "userorganizations.raml", size: 1174, mode: os.FileMode(420), modTime: time.Unix(1458554765, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _usersRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1c\x5d\x73\xdb\xb8\xf1\x5d\xbf\x02\x75\xdb\xe9\xc3\xd9\xfa\x8a\x7d\xbd\xd3\x4b\xce\xc9\x25\xa9\x73\x76\xe2\x89\x9d\x74\x3a\x6e\x9a\x83\x48\x48\x42\x4d\x11\x3c\x10\xb4\xad\x64\xf2\xdf\xbb\x00\x41\x11\x20\x01\x92\xb2\x14\x27\x99\x29\xa7\xbd\x58\xe4\x62\xb1\x5f\xd8\x2f\x80\xfc\xf3\x5f\xdf\x1c\x9f\x9d\xa2\x51\x7f\xd8\x13\x54\x44\x64\x82\x4e\x44\xfa\x2f\x96\xbd\x8e\x23\x1a\x93\xde\x0d\xe1\x29\x65\xf1\x04\x0d\xfb\xa3\xde\x14\xa7\xe4\x2d\xa7\x13\x34\xe8\xa5\x24\xc8\x38\x15\xab\x8b\x60\x41\x96\x24\x9d\xf4\x10\x3a\x40\x0c\x67\x62\xf1\x61\xfc\x61\x28\x7f\xe6\xd7\x9f\x68\x1c\x44\x59\x48\x50\x65\xc0\x60\x0d\xdb\xe7\x78\x19\xf5\xc4\x2a\xc9\xb1\x9c\xe2\x29\x89\xf2\xf1\xf2\xde\x04\xa5\x82\xd3\x78\xae\x6e\x2c\xf1\xdd\x29\x89\xe7\x62\x31\x41\x47\xc3\xfc\x0e\x8d\x8b\x3b\xe3\x1e\xdc\x39\x3e\x3f\xb9\x26\xab\x7c\x78\x48\xd2\x80\xd3\x44\x28\xf2\xdf\xa6\x84\xa3\x34\x21\x01\x9d\xd1\x40\x82\x21\x80\x53\x60\x09\x67\x09\xe1\x82\xe6\xd3\xe7\x57\x06\xd0\x31\x5e\xda\xb3\xcb\x0b\x27\x54\xe2\xaf\xdf\x4e\x22\x1a\x60\x39\x15\x0d\x6b\x4f\xd3\x80\x49\xe6\xf4\xed\xab\xf7\xeb\x07\x91\xe2\xb5\x00\x97\xe4\x83\x4c\x18\xa7\x1f\x15\x26\x07\x17\xcf\x19\x47\x38\x46\xe4\x2e\x89\x70\xac\x80\x10\x9e\xb2\x4c\xe8\x29\xe0\x59\x98\xff\xb9\x04\x8a\x00\xe7\x3e\x88\x9d\xa0\x85\x10\x49\x3a\x19\x0c\xe6\x54\x2c\xb2\x69\x3f\x60\xcb\x01\x15\xe9\x8a\x65\x4c\xe9\x78\x40\x43\x12\x83\xee\x57\xc0\x34\xa8\x7b\x30\x8d\xd8\x74\xb0\xc4\xa9\x80\xbf\x43\x16\x68\x55\x8d\x07\xf9\x1c\xfd\x65\xb8\xb9\xd8\xe6\x1c\xc7\x82\x84\x97\xac\x84\x75\xa8\x37\xbf\x2c\x86\x2f\x17\x04\x31\x3e\xc7\xb1\x96\x09\x12\x0b\x2c\x10\x4d\x0b\x8c\x88\x13\x1c\x22\x1c\x04\x24\x4d\x91\x60\x28\x65\x4b\x18\x30\x03\x30\xa2\xa8\xf9\x5b\x8a\x68\x3c\x63\x7c\xa9\x46\x97\xfa\x0a\x43\x0e\x23\x1e\x4f\x7c\xf3\x9e\xe5\xf2\x2b\x50\x29\x45\xa5\xc5\xaf\x62\x6e\x16\x13\x35\xa9\x01\x21\xaa\xf4\x72\xf2\x47\x46\x40\x94\x61\xdf\x98\xca\x25\x3b\x79\xed\x5d\xbd\xdf\xb3\xef\x78\xa4\x94\x64\x53\xb0\xb7\xdf\xc8\x0a\x58\xa8\x9b\x15\x68\x9f\x46\xdf\x2f\x6f\x0b\x98\xfb\xbb\xa5\x7e\x8a\xe3\xeb\xef\x96\xf8\x19\x0e\xc8\x94\x31\x60\x00\xc1\x3f\x11\xc1\xe5\x8a\xc9\x7d\x87\xe3\x81\x49\x56\xea\x5b\xde\x86\x71\x56\x24\x72\x4a\x53\x21\x05\x60\xa1\x51\xfc\x6a\x16\x95\xb8\x4c\xd6\x03\x70\x80\xd2\xab\x81\xff\xe2\x68\x49\x96\x53\x88\x50\x0b\x9a\x00\x92\xbe\x74\xa1\x4f\x40\x01\xc7\x41\xc0\xb2\x58\x4c\xbc\x8e\x8a\x82\x9a\x3a\xb8\x22\x23\xde\x3c\x1a\x96\x2a\xa6\xc1\x66\x63\x47\xe5\x58\x45\x17\x5f\x6d\x36\xfe\x70\x28\x19\x3b\x97\xeb\x22\xce\x24\xbf\x9a\x31\x2c\xc0\x45\x83\x0c\xff\xf3\xef\x1f\x1e\x5f\x0d\x0f\x7e\x7e\xff\xc3\x5f\x54\x10\xc9\x9d\x9b\x9f\xfb\x00\x7c\xbd\x6d\x0f\x1e\x1a\x7c\x12\x00\x40\x42\xc4\xe6\x28\x8e\x4a\x14\x31\xdf\x7c\xb8\x21\x47\x06\x16\xc2\x1f\x6f\xc5\x84\x43\x15\xdd\x70\x1c\x96\x38\x12\x96\x0a\x1c\x05\x2c\x24\x9b\xa3\x19\x4b\x34\xcf\xf5\x8a\x6b\xb1\x58\x48\x2a\x7a\x36\x6e\x0a\xae\x62\x4e\xb8\xbe\xab\x82\x6e\xaf\x61\xf6\x84\x06\x22\xe3\xcd\x30\x90\x0b\x5c\x7b\x00\xe0\xee\x0b\xe5\x00\x5a\xe8\x8c\xd8\x9c\xc6\x8d\x93\xb4\x71\x82\x6f\xb0\xc0\xfc\x43\xc6\xa3\x46\x34\x0b\xb1\x8c\x5a\x81\x1a\xa4\x02\x77\x65\x46\xe8\x5f\x23\xeb\x54\x66\x1b\x13\x53\x4f\xcc\x0c\xb5\xb8\x39\xa3\x3c\x15\xce\x4c\x29\xc2\x9e\x07\x65\xc8\x77\x45\xfc\xbb\x84\x72\x08\x9a\x28\xc4\x82\x34\xe4\x01\x5f\x20\x38\x6f\x8e\xd2\x70\x64\x8d\x09\xd9\xe6\x98\xb5\xe7\xf3\x07\xe3\xcd\x51\x1a\xe1\xc4\x11\x26\x1d\xc6\x51\x59\xd1\xb5\x00\xea\x18\x62\x2d\xad\x9e\x02\x20\x77\x78\x99\x44\xc4\x95\x5a\x4f\xd9\xd4\x65\x14\x16\xde\x03\x34\x7a\xf1\xee\xe2\xef\x47\x87\x87\xe2\xe6\x9f\x67\xe3\xcb\xb3\xd1\xa3\x9b\x57\xff\xf8\x29\x9d\x32\xbc\x7c\xf9\xf1\xf4\xc7\xe5\xcb\x47\x3f\x56\x8c\x47\x3a\xa4\xd1\x4f\x07\xa3\xe1\xc1\x78\x68\x5b\x90\x8d\xfa\x96\xf1\x6b\x45\xc4\x2f\x50\x47\x24\x38\x5e\xc9\x7a\xc2\x82\x58\x30\x4d\xe6\x2f\x9a\x8b\x1a\xc4\x5e\x88\xe3\x80\x20\x28\x0c\xa7\x7b\x39\xa4\xac\x43\xd4\x4d\x79\xaf\x5f\x8c\x83\xb8\x6f\x5b\x9e\x49\xca\x14\x26\xf9\xe1\xd1\x78\x34\x7e\x74\xa8\xff\x6f\x3c\x24\x73\x78\x08\x0c\x39\x1e\x62\x33\x3c\x5a\x34\x57\x0c\x40\x07\x4a\x74\x91\xc8\x25\x30\xa3\x24\x0a\x6b\x10\x3a\x12\x42\x62\x47\x63\xfd\xa3\x06\x03\xa1\x0e\x1d\x1e\x3d\xa9\x63\xd7\x01\x08\x5d\xb2\x25\xe3\x9c\xdd\x42\x4d\x57\x9f\xc1\x08\x31\xe8\xf2\xe0\xe7\xc3\xa3\x61\x5d\x1d\x1e\xc2\x9f\x42\xb5\x0a\x63\xd1\x53\xf8\xe5\xa5\xbc\x00\x6a\x20\x7e\xb4\x13\xd2\x47\xc3\xe1\xd0\x5a\x99\x36\xd9\xd7\xd3\xa0\xce\x87\xca\xd3\xd0\xe5\xa9\xd4\xe1\x51\xf9\x9f\x1a\x9c\x4c\xc9\xd0\xf1\x93\xa7\xbf\x3e\x7b\xde\x91\xd6\x5e\xbe\xa4\x6e\x28\xb9\xdd\x49\x10\xd8\xb4\x22\xfb\xbf\x27\xfe\xf6\x3d\x71\xf1\xcc\x2e\x71\x4a\xed\x02\xc0\x4b\x46\xe3\xd7\xc6\xe3\x93\xf8\x06\x56\x53\xd9\xbd\x71\xb1\x69\xa2\xab\xa9\x55\x9a\x5c\xed\x26\x67\x51\xd7\x44\x44\xaa\x73\x82\xae\xd8\x6d\x4c\xf8\xbe\x2e\x8f\x14\xa5\x4f\x19\xac\x02\x1c\x88\x0b\x3a\x8f\x61\xd0\x9b\xbc\xba\x6a\x28\x12\x34\xfc\x49\xbd\x9f\x95\x60\x2e\xca\x1e\x58\xde\x0a\x24\xe1\x13\xb8\x75\x55\xb6\x00\xd1\xfb\xde\x40\x32\xa3\x50\x4a\x4f\xe0\x68\x67\x3d\xe5\x04\xf2\x15\x84\x51\x4c\x6e\x15\xe7\x0a\x64\xca\x42\x23\x2b\x37\x5a\x6b\x83\xff\xa6\x2c\x76\xc9\x41\xa6\x72\x92\xc7\xc1\xa7\x62\xc5\x7e\xce\xc1\xe6\x65\x95\x62\x52\x59\xf6\x29\xd1\xa7\x75\x77\xee\x0a\xed\x29\xe1\xe3\x10\x52\xb6\x3d\xf4\x1e\x7d\x46\xc5\x02\x06\x53\x4e\x40\xf5\x55\x5b\x1d\x0f\x87\x75\x97\x65\x53\x6f\x5e\xcd\x9c\x78\xb8\x92\x3f\x07\xa6\x0f\xda\x86\x8f\x24\x33\x8a\xb6\x90\xa6\x49\x84\x57\xaf\x54\x52\xf1\x36\x91\x99\xa3\xfc\xbb\x7c\x6e\xb5\x4f\xd5\xf3\x75\x5b\x0d\x2d\x68\x5a\x26\xb1\xaa\xed\x58\x24\xae\x3d\x9f\x20\x9a\xd9\xf7\xf9\x83\xa6\x5c\x39\xbf\xbc\x19\xb3\x53\x6b\xe3\xe1\xa1\x3d\x81\xc5\xe6\x45\xa6\x3a\x89\xb3\x2c\x8a\x56\x28\x53\x3c\x87\x6d\x7c\x22\x28\x08\x9b\x70\xaa\xce\x73\xcc\x04\x9a\x81\x53\x09\xf3\xf4\x6e\x90\xe0\x34\x85\xd0\x1d\xee\x5c\xab\x4e\xbd\x9e\xeb\xd9\x7a\x3e\x12\x1d\xda\x4d\xea\x63\xea\xa6\xdd\x66\xd2\x7e\xad\x82\x83\xc9\x38\x27\xb1\x58\x4b\xc2\xad\x5d\x24\x7d\x43\x03\x8c\x67\x69\x3a\x14\x5d\xe3\xfa\x15\x5b\x8f\x96\x75\x13\xae\x40\x1f\x8e\xc7\xae\x15\x6a\xa1\x00\x77\x8f\x23\x1a\xae\x85\x85\xf2\x56\x33\xb8\x78\x28\x6b\xe1\x91\x7c\x56\xe1\xd3\x81\x72\x7b\x97\xd1\x24\xe7\xe2\x22\x90\xfb\xd8\xa1\x65\xa0\xd2\x12\x1d\xf0\xcb\xa1\x5b\xd9\xe2\xda\xd1\x6b\x69\x99\xc6\xf8\x86\xcc\xa9\xdc\x55\x78\x45\x6e\x9f\xc9\x99\xab\x49\x42\x45\xb8\x05\xb8\x8e\x0f\x8a\xd8\x22\x3d\xf9\x62\x56\xa9\xb7\x64\xd4\x2e\x54\xed\xa1\x29\xaf\xcd\x6c\x71\xd4\x62\x8b\x05\xaf\xe0\x6e\x9a\xb9\xf5\xf1\xdc\x8d\xf3\x2e\x12\xe8\x26\x8a\xf5\xd5\x22\x13\x79\x1d\x0e\x7f\x6e\x5d\x49\x6a\x16\xb9\xb1\x83\x23\xb9\xa5\xb3\x92\x8e\x68\xdd\xdc\x1e\x7c\x52\xb4\x7c\x9e\x18\xe9\x76\xa5\xfd\xe8\xf0\x7a\x1e\x1b\x73\xfa\xbe\xb4\xec\xb6\x4b\x2f\x3f\x80\x05\x0c\xcb\x37\x53\x7b\x49\x72\xc3\xcd\xab\x0d\x97\x26\xda\xb5\xd0\x2c\xff\x16\xc1\xb7\x4a\x7c\x33\x3b\x74\xca\xa3\xea\xa6\xba\xa8\x50\xee\xd5\x49\xd3\x8d\x5c\xaa\x34\xd2\x8a\x88\x08\xd2\xa0\xbc\x5f\x15\x40\x37\xe5\xbd\x21\x4b\x76\xa3\xf6\x3d\x1b\x54\xe4\x15\x87\x27\x48\x54\xe6\x78\x66\x22\x06\x64\x72\x46\x6b\xdb\x65\x33\x09\xc9\xfc\xc1\x26\x56\xed\x68\xc8\x04\x61\x4a\x4a\xf4\x6b\x54\x03\x15\x61\xb0\x2d\xb2\x46\x37\xfb\x4e\x0f\xe8\xe2\x63\x2f\x48\x1c\xa6\x48\x4f\x21\x77\x57\x72\xca\x04\xf3\xca\xd3\x9b\x0b\xd7\x25\x69\xcd\xf4\xae\x3a\xc7\x02\xa7\xc0\x31\x91\x9b\x39\xeb\xb4\x28\xdf\x5e\xff\xf2\x91\x08\xa4\x72\x9c\xd0\xdf\xc8\xca\x27\x17\x00\x50\x56\x75\x7c\x7e\x02\x50\xc5\x6e\xdc\xba\x44\xc9\xaf\x07\x89\x3c\x5b\x04\x15\x60\x42\xc5\x93\x46\x26\x7c\xac\x74\x63\xa8\xb8\x74\xd5\xaf\x26\xda\xb5\xf7\x9f\x13\xbf\x22\xe5\xf6\x61\x3e\xab\xd7\xc4\x25\x48\xee\xde\xf5\xa9\x8f\x0e\xd6\xec\xa9\xed\x9c\x3b\x97\x0e\xac\xf9\xb5\x4b\x99\xae\x7b\x49\xf7\x0b\x85\xf9\xb1\x98\x4d\x82\x20\x24\xb2\xea\x17\xac\x48\x54\x1d\xfa\xa0\x31\xef\x9b\x0e\x66\xf3\xea\x2e\xa8\x25\xfb\x17\x44\xb4\x09\x1e\x40\xb4\x9b\x91\x52\x46\xd3\x55\x3e\x61\x57\x09\x38\x6c\xb4\xea\x03\x8c\x93\x4e\xf6\xb5\x7d\xd5\x61\x2d\xfa\x4d\x03\x7c\x9b\x64\x8c\xd0\xee\xe2\x61\xcb\xa0\x5e\x08\xdc\x8e\xb7\xfa\x84\x54\x31\xba\xca\x89\x83\x0b\x77\xc3\xd0\xb6\x41\x79\xc8\xea\x5a\xb7\x16\x51\x15\xb4\x7b\x38\xb5\xe9\xcf\xd1\xa0\x4c\x21\x27\x45\xfc\x2c\xfa\xa1\x3e\x0e\xb6\x09\xa8\x4e\xfe\x7d\x3d\x56\x8b\xd8\x1c\x34\x77\x27\x69\xca\x02\x9a\xb7\x55\xf4\x50\x79\x7a\xeb\x9e\x22\x71\xce\xe3\xc0\xdc\xb3\x5a\x83\x03\x79\x22\xec\xfe\xfd\x41\x39\x7a\x6f\xdf\x2f\xa7\xaa\x0b\x90\xfd\x9f\x93\xda\x19\xb4\xaf\xd3\x52\x2c\xb6\x3c\x2a\x02\xb1\x53\x4c\x43\x28\x7f\x64\x84\xaf\xce\x31\x07\x28\xa1\x5b\xb9\x06\x4a\x76\x4d\x1c\x33\x56\xf7\xd9\xcd\x89\x2a\xbd\x86\xfb\x9b\xa3\x41\xe3\x43\x4b\xb2\x53\xed\xec\xdc\xc7\x70\xc9\xc9\xcc\xcf\xcd\x94\xd5\xd7\x38\xb1\xf3\xf9\xa6\x7e\x89\x9d\xba\x6f\xd2\x54\xbf\x77\x7f\x64\x5d\x94\xda\x54\x3a\x34\x54\x09\xde\xcd\xbd\x90\x5d\xf4\x7c\xb6\xac\xb9\x3d\x9c\xd5\x73\x89\xae\x69\x6d\x35\x8f\xb3\x32\x89\x87\x35\x69\xd7\x8e\x5d\x7b\x4f\xb9\xa5\xb0\x34\x5a\xca\xfe\xae\xca\x1d\xcd\xcf\xf7\x69\xe9\xf6\x1f\xba\xb6\xd2\x03\x8d\x83\x75\x35\xfd\x36\x28\xa3\xad\xfc\x72\x25\x9f\xce\xd4\x73\xe3\xc4\xb3\x9e\x62\xb9\x12\xac\xb6\xee\xea\x3a\xbf\xea\x5c\xe4\x77\x09\xc2\xc7\xd5\x46\x49\x25\x04\x18\xd2\xde\x65\x14\xf0\x46\xaa\x22\xb6\x85\x26\xdd\x5d\x8e\x9e\x33\xf5\x2f\x8e\x2a\xa8\x65\x19\x3d\xa3\x91\xfc\x8b\xc5\xf0\xbf\x68\x55\x4e\x61\x9a\x52\xa3\xf3\xb3\xd6\xf1\xee\x9d\x99\x3f\xf8\xb8\xf7\xfd\xbb\x04\x9e\xfa\x59\x81\xa6\xe0\x93\xd4\xa0\xbb\x07\x20\x1f\x6f\xde\x45\x6c\x2d\xe0\x3a\x9d\x9e\x7d\xc0\xca\xf2\x6d\x0e\x41\x6a\x0a\x54\x39\x2a\x51\x67\xab\x9d\xb9\x36\x26\xdb\xb9\xed\xca\xb6\xbc\x1c\xee\xe6\x3b\x0c\x50\xfe\xe3\x2a\x6d\x1b\xda\x2e\xa9\xdc\x3f\x4c\x19\x02\x2f\x43\xd5\xae\xb6\xb9\x1b\xf4\xdd\xaa\xe8\x6e\x16\xde\x1a\x9e\xda\xac\xe5\x3e\xa1\xc9\x11\x98\x5a\x95\xb2\x0e\x4d\x0e\x37\x72\xef\x8a\xf9\xdc\x58\xc1\xd5\x2e\xfe\x00\x07\x82\xde\x58\x3d\xf6\x4e\x1d\x76\xf7\xa2\x6b\x6e\xb0\x0b\x72\x27\x64\x0c\x31\x3d\x4a\xa7\x74\xa3\xb6\xb4\x7c\xed\x75\xcd\xa3\x6a\xab\xdb\x23\xdc\x4b\xb1\xcb\x32\x2c\x19\x70\xbd\xfe\xd6\x94\x2d\xbe\x23\x9c\xce\x56\x4a\x52\xaf\x1a\x25\xa5\x00\x69\xa9\xf6\xba\xb7\xdd\x6d\x4e\x98\x2e\xd3\xfc\x68\xa2\xe7\xdc\x41\x33\xcb\xbb\xb1\xc3\x9b\x9c\xe7\x4a\x8e\x58\x3f\x78\x60\x61\xa0\xfa\xd0\x81\x45\xa0\x3c\x72\x20\xd9\xd9\x99\xc6\xed\xe3\x02\x95\x24\x4e\x1e\xe7\xdb\x3a\x7b\xb3\xb2\x0e\xc7\xb1\x30\xe9\x6a\xe4\x44\x95\x06\xd1\x6e\x2a\x59\x77\x8a\xe4\x3b\x69\xd8\x5e\xc1\x3e\x7c\x12\x57\xa5\xf2\xab\x06\xe8\xdd\x36\x45\x6c\xd6\x6a\x86\x67\xa7\x26\x3b\x29\x1f\xbe\x56\x4a\x53\x55\xa2\x9d\xd0\x38\x32\x15\x33\x1f\x91\xc2\xd0\x6b\x43\x1f\x4d\x03\xa9\xf8\x33\x93\x2e\xd4\xb9\xa9\xaa\x05\x75\x47\x77\x17\x37\x68\x2d\x66\x42\xbe\x76\x6d\xbe\x23\xb8\x23\xad\xd5\x36\x53\x54\x26\xa7\x77\xe6\xac\x69\xf7\xe5\xa3\x14\xe8\xe4\x04\x25\x10\x1a\xa5\x00\xe9\xfa\xb4\x6e\x2a\xfd\x27\x88\x87\x33\xf0\xaa\xe9\x37\xbe\x8c\x90\x49\xf8\xa4\xe1\x0c\xb2\xf5\xda\xa5\x93\xae\x9c\xdf\x2d\x70\x14\xc7\x84\xf5\x79\x62\x40\xe5\x3e\x68\x9c\x1f\x97\x36\x8d\xa2\x18\xb9\x4b\x83\x70\x9a\xc3\x7a\x22\x74\xbb\x80\x92\xae\x3c\xe6\x08\x19\xec\xa8\x78\xf9\x56\x9e\x69\x06\xb1\xf7\xd1\x6b\x1e\xc2\x23\x89\x48\x1b\xc9\x74\xa5\xde\xaa\x2a\xd6\x95\xb7\xd1\xa0\x3f\x74\xf0\x4c\xbd\x4d\x13\x56\x5a\xe5\x6a\x59\x55\x5f\x9d\xad\x51\x7c\xa2\xbf\x95\x20\x09\xca\xdf\xca\x09\x4b\xea\xf7\x15\x25\x64\x86\xb3\x48\xe4\x7d\x07\xb5\xa1\x22\x33\x57\x93\x47\x69\xdf\x9c\x88\x8c\xc7\xd5\x53\x2a\xf2\x85\x5a\x45\x1a\x9a\x81\xce\xcb\x4c\x21\x15\xc0\xbb\x8b\x5e\xfb\xcd\x3c\x07\xbd\x17\x72\x24\x48\x70\x96\x12\xb1\x2f\x65\x3a\xcb\xf2\x9d\xe3\x04\xcf\x41\x74\x7d\xf0\x0d\x39\xb9\x20\xe9\xdf\x87\xbf\x77\x23\x67\x89\xef\xee\x45\xcc\x19\xbe\x93\xf3\x12\x94\xd2\x8f\xa4\x95\x9a\xa3\xae\xe4\x68\x92\xe8\x52\x9e\xab\x1f\x1f\x0d\x6b\x5b\x1a\xe6\x27\x1b\x76\xec\xdc\xaa\x9b\xc9\x51\x64\x7d\x20\x22\xed\xe6\x06\x6d\x0a\xfb\x5f\x39\xe2\x59\x1c\x68\xe7\x32\xf8\xb4\xfe\x40\x84\xbb\xdf\x50\x93\x84\x89\xa4\x59\x0a\x16\xf7\xca\x1c\x70\xf9\x1d\x10\xf3\x55\x8c\x6f\x4a\x30\xdd\xab\xe9\x0e\xa2\xc8\xeb\x69\x9f\x34\x62\x4b\x0a\xfb\xf6\xa7\x08\xcc\x17\xef\x6f\x69\x14\x41\x4c\x45\x11\x8b\xe7\xf2\x90\x38\xbe\x21\xeb\x2f\x6f\x38\x3f\xb7\x51\x48\xb4\xad\x4b\xd3\x81\x85\x33\x16\x42\x15\x09\x0e\x9c\x06\x0b\x73\x8a\x2a\xf9\xaa\x2b\x31\x8d\x88\xfa\x16\x08\x21\xbb\x48\x86\x6c\xea\xfe\x17\x00\x00\xff\xff\x95\xb9\x35\x0e\x2f\x47\x00\x00")

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

	info := bindataFileInfo{name: "users.raml", size: 18223, mode: os.FileMode(420), modTime: time.Unix(1465383637, 0)}
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

