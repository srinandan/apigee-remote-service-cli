// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/istio-1.6/apigee-envoy-adapter.yaml
// templates/istio-1.6/envoyfilter-sidecar.yaml
// templates/istio-1.6/httpbin.yaml
// templates/istio-1.6/request-authentication.yaml
// templates/istio-1.7/apigee-envoy-adapter.yaml
// templates/istio-1.7/envoyfilter-sidecar.yaml
// templates/istio-1.7/httpbin.yaml
// templates/istio-1.7/request-authentication.yaml
// templates/native/envoy-config.yaml
package templates

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _istio16ApigeeEnvoyAdapterYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4b\x6f\xe3\x36\x10\xbe\xfb\x57\x10\xb9\xd3\x8f\xb6\x7b\x88\x80\x3d\x04\xdb\x45\x1a\xa0\xce\x1a\x4d\xda\x3b\x43\x8d\x25\xa2\x14\x87\x1d\x8e\xd4\xaa\x82\xfe\x7b\x21\x51\xb6\xa5\xc4\xaa\xdd\x20\xa8\x2f\xb2\xe6\xf1\xcd\xc7\x79\x88\xa3\xbc\xf9\x0d\x28\x18\x74\x89\x50\xde\x87\x55\xb5\x59\xfc\x6e\x5c\x9a\x88\x1f\xc1\x5b\xac\x0b\x70\xbc\x28\x80\x55\xaa\x58\x25\x0b\x21\x9c\x2a\xa0\x33\x35\x19\x80\x24\x28\x90\x41\x06\xa0\xca\x68\x90\xe0\x2a\xac\x07\x9b\xe0\x95\x86\x44\x34\xcd\xf2\xf1\xf0\xd6\xb6\x8b\xe0\x41\x77\x28\x04\xde\x1a\xad\x42\x22\x36\x0b\x21\x02\x58\xd0\x8c\xd4\x69\x84\x28\x14\xeb\xfc\x67\xf5\x02\x36\x44\x81\xe8\x98\x5d\x8a\xc9\x50\x78\xab\x18\x06\x8c\x11\xe3\x1e\xc0\x39\x64\xc5\x06\xdd\x11\x53\x88\x60\x52\xd0\x8a\x96\x26\xb0\xc1\xa5\xc1\x15\xc1\x9f\x64\x18\xee\xbc\xff\xe9\xf9\x79\xb7\x23\x7c\x01\x0a\x89\xb8\x61\x2a\xe1\xe6\xe8\xe6\x09\x0b\xe0\x1c\xca\xd0\xf9\x78\xc5\x79\x22\x56\x05\x30\x19\x1d\xe6\x8c\x90\x38\x11\x37\x9f\xd6\xeb\xcd\x1c\x4e\xd0\x39\x74\x99\xcd\x99\xfd\x1c\x4c\xd0\xa4\x3c\x5c\x20\xc4\x75\x67\x72\x92\x75\x29\x1f\x4c\xed\x24\xa9\x57\xa5\x35\xfe\xaa\x43\x8b\x54\x9b\x5e\x76\x28\x63\xff\x1f\x74\x49\x86\xeb\x2f\xe8\x18\xfe\xe2\x44\x88\xa3\x1b\x95\xee\x2e\xfc\x1a\x80\x12\x71\x7b\x7b\x3b\x15\xdf\x13\x96\xfe\x8c\xfc\x11\xdd\x2f\x88\x9c\x88\xee\x8c\x83\x4a\xa3\x63\x65\x5c\x57\x8c\x41\x22\xaf\xea\xc3\xf8\x33\x85\xca\xba\xa4\x65\x88\x99\x85\xd5\xe0\xd2\xdb\x48\x95\x2a\xcf\x40\x49\xd3\x2c\x1f\x3a\xb3\x67\x95\xb5\xed\xcd\xd4\x75\x57\x5a\xbb\x43\x6b\x74\x9d\x88\x87\xfd\x23\xf2\x8e\x20\x74\x63\x71\xcc\x3f\x12\x8f\xb2\x2a\x4f\x7c\x77\x7d\xdd\x3f\xad\xd7\xeb\xa3\xd6\x9a\x0a\x1c\x84\xd0\x77\xd7\xc9\x49\xf4\x85\xbf\x07\x1e\x8b\x84\x18\xba\x2b\x07\x65\x39\xff\x7b\xaa\x3a\x60\x6f\x46\xe2\xbd\x32\xb6\x24\x78\xce\x09\x42\x8e\x36\x8d\xf3\x75\x74\x01\x32\x98\x3e\x81\x46\x97\x76\xa3\x77\x62\x45\xa0\x52\xf3\x7f\xd2\xfa\x7e\x7d\x15\x2f\x45\x59\x18\x87\x96\x42\x4a\x8b\x99\xb4\x50\x81\xfd\x9c\xc2\x4b\x99\xbd\xd2\x6a\x74\x7b\x93\x7d\x5e\xc5\xe7\xf0\x58\xd6\xaa\xb0\xa3\xc3\x06\x2c\x49\xc3\x04\xd9\x9a\xc2\x70\x98\x1e\x53\xfb\xb2\x63\xb3\x2e\x26\xd2\x02\x0a\xa4\xba\x57\x6c\xcd\x48\x43\xf0\x47\x09\x61\x06\xe3\x1a\x88\x0a\x6d\x59\xc0\x16\x4b\x37\x6d\xa7\xa2\x93\xec\x62\xca\xe3\x79\x46\x60\xd7\xcf\x41\xac\xf2\x37\x67\xeb\x61\xb8\x9a\x46\x0a\xb3\x17\xcb\x87\x70\xff\x65\xb7\x55\x4e\x65\x90\xb6\xed\x42\x88\x83\xe2\xce\x29\x5b\xb3\xd1\xe1\x09\x34\x01\xf7\xba\x73\xa4\xd4\xc1\x4e\x86\xde\xf0\x2d\xbd\x79\x83\x57\x9c\x62\x70\xb0\x01\x7a\x06\x5f\x9d\xc6\x14\xd2\xee\x0a\x99\x8d\x8e\x9e\x87\xa1\x5e\xb1\x0d\x6f\x62\xb3\x0d\x32\x66\xf6\x62\x54\x97\xce\x06\xf1\xfd\x07\x60\xee\x7c\x73\xda\x33\x09\x1f\x07\x89\xb4\xde\xf5\x51\x8b\x7d\xb0\x55\x3e\xf9\x4f\xbd\xf0\xee\x92\xcb\x4b\x95\x8c\xaf\x63\x36\x29\xec\x55\x69\x79\x8b\x29\x24\xe2\x87\xef\xc6\x03\x1f\x8d\x1f\x7b\xc4\xa6\x59\x7e\xa3\xac\x6d\x65\xd3\x2c\xbf\xba\xaa\x6d\xe5\x99\x18\x17\x9b\x42\xce\x57\xfb\xdd\xcc\x86\x44\x96\xa9\x56\x91\xdc\x28\xac\x8c\xad\xf6\xba\xa2\xf2\x5f\x1b\xe2\x03\x73\x34\x0d\x70\xa2\x21\xa5\x5c\x8c\x77\xba\xe3\x3a\xf7\x14\xfb\xe0\x03\x77\xb9\xe9\x3e\x71\x71\x97\x38\xec\x0c\xc7\xdb\x52\x9e\xee\x8a\x78\xee\x48\x27\x23\xaf\xdf\xec\x84\x97\xd0\xff\x09\x00\x00\xff\xff\xf7\xa1\x44\x17\xc7\x0a\x00\x00"

func istio16ApigeeEnvoyAdapterYamlBytes() ([]byte, error) {
	return bindataRead(
		_istio16ApigeeEnvoyAdapterYaml,
		"istio-1.6/apigee-envoy-adapter.yaml",
	)
}

func istio16ApigeeEnvoyAdapterYaml() (*asset, error) {
	bytes, err := istio16ApigeeEnvoyAdapterYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "istio-1.6/apigee-envoy-adapter.yaml", size: 2759, mode: os.FileMode(416), modTime: time.Unix(1600282230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _istio16EnvoyfilterSidecarYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x56\x4d\x6f\xe3\x36\x10\xbd\xeb\x57\x0c\xe4\x6b\xad\x4d\x5a\xb4\x07\x9d\x9a\xcd\x3a\x5d\xa3\x59\x67\xe1\x78\xbb\xbd\x09\x34\x35\x96\xd8\x50\x24\x4b\x8e\x9c\x68\x83\xfc\xf7\x82\x22\x95\x58\xfe\x42\x8b\x1e\x0a\x34\x87\x58\xe0\x7c\xf0\xcd\x9b\x37\x23\x4d\x60\xae\x1c\x31\x29\x1d\x30\x05\x73\x47\x42\xc3\x4c\x6d\x75\x77\x23\x24\xa1\x05\xa1\x80\x6a\x84\x12\x37\xac\x95\x04\x8a\x35\xe8\x0c\xe3\x98\x25\x13\x98\x13\x30\x63\xa4\x40\x07\xa4\x81\x49\x09\x0e\xed\x56\x70\x74\x43\xd8\x9b\x3b\xcc\x37\xd0\xe9\x16\x1e\x85\xab\xbf\xf3\x4f\xc9\x04\x1a\xd6\x01\x31\x21\xb5\x85\x75\x07\xce\x20\x17\x9b\x4e\xa8\x0a\xd2\x47\x6d\x1f\xa4\x66\xe5\x3d\x4a\xe4\xa4\x6d\x0a\x1b\x6d\xa3\x87\xe0\x40\xcc\x56\x48\xce\x63\xf8\xe2\xd0\x41\xca\x8c\xa8\x10\xa7\x16\x1b\x4d\x38\x8d\x28\xa6\xe8\xeb\xc8\x82\x2d\xff\xf1\xe2\xe2\x22\x05\xe6\x62\x70\x96\x24\x13\x58\xd5\xc2\x81\x70\xf0\xa0\xf4\xa3\xf2\x35\xf8\x7b\x81\x6a\xab\xdb\xaa\x8e\x64\x5c\x66\x3f\x65\x49\xc2\x8c\xf8\x0d\xad\x13\x5a\xe5\xa0\x90\xbc\x9f\x50\x55\x26\xbc\x4b\x26\xf4\xbb\xed\x25\x93\xa6\x66\x3f\x24\x0f\x42\x95\xf9\x2e\x83\x49\x83\xc4\x4a\x46\x2c\x4f\xa0\xe7\x23\x87\x31\xda\xe7\xe7\x6c\xd5\x43\xba\x0f\xb0\xb3\x05\x6b\xf0\xe5\x25\x7a\xf7\xec\xe5\x03\xff\x89\xa7\xc0\x27\xda\x27\xc8\x9f\x01\x48\xb6\x46\xe9\xc2\x33\x40\xc3\x14\xab\xb0\x9c\xae\xbb\xe1\xca\x04\x80\x6b\xb5\x11\xd5\x67\x46\xbc\x46\x97\x27\x09\xc0\xb4\x6f\x62\xb7\xd2\x39\x7c\x5c\xad\x3e\x17\x37\xf3\xdb\xd5\x6c\x99\x84\x0c\xc4\xeb\x21\x1d\xd7\x8a\xf0\x89\x72\xb8\x9f\x7f\x98\x5d\x5f\x2d\x8b\xf9\xe2\xfd\xdd\x97\xc5\x87\x68\x96\xc2\x11\x2a\xb4\x83\x3b\xc0\xa6\xaf\xff\xba\x66\x42\xbd\x1d\x0e\xc7\xbb\x27\x03\x2f\xa1\x5f\x35\x91\x29\xb8\x56\x0a\x39\x09\xad\x8a\x50\x86\x1d\xf9\xbb\x76\x7d\x73\x24\xcd\x38\x91\xd5\x2d\xc5\xb8\xfe\x9f\xd9\xad\x46\x1b\xb4\x8c\xfa\x7e\xce\x17\xf7\xb3\xe5\xaa\x78\x3f\xbb\xb9\x5b\xce\xa2\x79\xcb\x64\x8b\x6f\xc9\x77\xd3\xe2\x13\x15\xac\xa5\xfa\xdb\xab\x35\x70\xba\x0b\xa5\xb2\x86\x17\x51\x86\x63\x88\x95\xd6\x95\xc4\xc2\x3b\xec\x63\x0f\xc2\x2c\x5a\x2b\xf6\x25\x72\x52\xd0\x7b\x19\x1c\x31\x2a\x8c\xc5\x8d\x78\x3a\x91\x62\x14\x40\xa2\x41\xdd\x52\x0e\x97\x6e\xe7\x7c\xd0\x6b\x11\x1b\x5e\xbc\xca\xd0\x8d\x11\x4f\x23\x21\xa1\xa1\xae\x6f\x5c\xf6\xc7\x63\x60\x47\xfd\xaf\xa5\x35\xd6\xc0\xdf\x51\xd7\xd5\xcd\x50\xf9\x59\x71\x8d\xb8\xb4\x6b\xc6\xcf\x88\xcc\xb6\x72\xbf\x25\x8c\x87\x4b\xaf\x6e\x6f\xef\xbe\x8e\x2c\x46\x4b\xc1\xc5\xbe\x3f\x0c\x2a\x89\xa4\xe8\x83\xaa\x01\x8c\x15\x8a\x0b\xc3\xe4\x41\x6c\xdf\x60\xd5\xe5\x40\xb6\xc5\xc3\x38\xb4\x8d\x70\x7e\x65\x1e\x0d\xac\x91\x95\x68\x73\x78\x86\xd4\x13\x90\xe6\x90\xfe\x3e\xbd\x0a\x68\xae\x5a\xaa\xb5\x15\xdf\xb0\x4c\xe1\x65\x4f\x48\x8b\xd9\xea\xeb\xdd\xf2\xd7\xff\x50\x4b\xe9\x59\x31\xa5\x67\xd5\xf0\x69\xb6\xfc\xe5\xc4\x8e\xa1\xce\x60\x59\x1c\xb6\x39\xfd\xd9\x5b\x3c\x3f\xfe\x37\x0b\x0b\x84\x19\xe1\x32\xae\x9b\x77\x01\x4b\x88\x8a\xea\xc9\xe2\x0b\xea\x14\xc2\x6c\xfb\x7d\xf6\x91\xc8\x5c\xbf\x5a\x3e\xed\x42\x1f\x94\xc4\xd1\xb9\x42\xea\x11\x96\xe9\xe1\x38\xf5\xab\xee\xcd\x7b\xc4\xd7\x61\x31\xe1\xb4\x69\xb4\x3a\x52\x69\xec\x90\xae\x8a\x63\xef\xc9\x63\x1b\x0c\xce\xee\xda\x68\x3f\xbd\x71\x23\xf1\xff\x7a\xef\x86\xbf\x7f\xb8\x7d\x01\x58\x59\x0a\x4f\x3f\x93\x85\xc5\x3f\x5b\x74\x54\x84\xa9\x70\x05\xe9\x7d\xea\x03\xfd\x39\x0b\xa3\x41\x1d\x4c\x5e\x3f\xca\x42\x01\x71\xa2\xf6\x42\x26\xc3\x40\x44\xb3\x3b\x48\xf9\x34\x8d\x58\x43\x13\x49\x3f\xa0\x3a\xe3\x64\xc4\x59\xa3\xb1\xba\x6c\x39\x9d\xbb\xc6\x7f\x32\xf2\x7e\x1e\x4e\x3b\x71\x29\x50\x91\x28\x4f\x7b\x94\xb8\x45\xe9\x27\x0b\x1b\x26\xe4\x5f\x01\x00\x00\xff\xff\x1b\x28\x64\x85\xc7\x0a\x00\x00"

func istio16EnvoyfilterSidecarYamlBytes() ([]byte, error) {
	return bindataRead(
		_istio16EnvoyfilterSidecarYaml,
		"istio-1.6/envoyfilter-sidecar.yaml",
	)
}

func istio16EnvoyfilterSidecarYaml() (*asset, error) {
	bytes, err := istio16EnvoyfilterSidecarYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "istio-1.6/envoyfilter-sidecar.yaml", size: 2759, mode: os.FileMode(416), modTime: time.Unix(1600282230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _istio16HttpbinYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x4f\x6b\xe3\x30\x10\xc5\xef\xfe\x14\x03\x7b\x76\xb2\xb9\x2d\xba\x2d\xec\x65\xa1\x14\x43\xa1\xf7\x89\xfc\xe2\x88\xe8\x1f\xd2\x24\x34\xfd\xf4\x45\x76\xe4\x38\x6e\x2e\x9d\x9b\x66\x1e\x3f\xbd\x37\xf3\x8b\xfe\x7a\x3a\x8a\xc4\xbd\xf1\x24\x9c\x06\x08\xe1\x83\x5d\xb4\xa0\x7f\x88\x36\x5c\x1d\xbc\x10\xfb\x9e\xde\x90\x2e\x46\x63\xd3\x70\x34\xef\x48\xd9\x04\xaf\xe8\xb2\x6b\x4e\xc6\xf7\xaa\x4e\x1b\x07\xe1\x9e\x85\x55\x43\xe4\xd9\x41\x55\xfa\xed\x9d\x23\x6b\x28\xea\x71\xe0\xb3\x95\x86\xc8\xf2\x1e\x36\x17\x39\x11\xc7\x78\xd7\xe7\x08\x5d\xda\x31\x24\x19\xe7\xed\x02\x38\xca\xcb\x44\xd1\x9f\xdf\xe3\x63\x32\xdf\xdd\x5b\x19\x16\x5a\x42\x7a\x82\x6e\xdb\xf6\x21\x05\xc7\x98\xb7\x73\x94\x7b\xee\x9f\xa6\xa9\x96\x13\xa2\x35\x9a\xb3\xa2\xdd\x37\x23\x8e\x45\x1f\x5f\x16\xa1\x57\xde\xa6\xd6\x65\xb1\x60\x22\x81\x8b\x96\x05\x37\xc2\xc2\x54\x29\xfb\x00\x7b\x8a\x5b\x03\xa7\x72\xec\x79\x40\xdf\xee\xaf\x65\x05\x66\x00\xc6\x51\x4d\x51\x4a\x07\x2f\x6c\x3c\xd2\xcc\x6f\xc9\x38\x1e\x4a\xe8\xa0\x4f\x48\x1b\x13\xb6\x27\x78\x0f\x39\x26\x18\xf9\xdc\xae\x3f\x1e\xd5\xdd\xd9\xda\x2e\x58\xa3\xaf\x8a\xfe\x1f\x5e\x83\x74\x09\xb9\x2c\xb8\xaa\xd6\xcb\x9d\x6a\x3e\x7d\xfd\x7a\xb6\x53\xef\xfc\x15\x00\x00\xff\xff\x35\x03\x1d\x9f\xc0\x02\x00\x00"

func istio16HttpbinYamlBytes() ([]byte, error) {
	return bindataRead(
		_istio16HttpbinYaml,
		"istio-1.6/httpbin.yaml",
	)
}

func istio16HttpbinYaml() (*asset, error) {
	bytes, err := istio16HttpbinYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "istio-1.6/httpbin.yaml", size: 704, mode: os.FileMode(416), modTime: time.Unix(1600282230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _istio16RequestAuthenticationYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xbd\x6e\xdc\x3a\x10\x85\x7b\x3e\xc5\x81\xd9\x5d\x5c\x69\xe1\x56\x9d\x53\xc5\x4e\x10\x18\x86\x13\x17\x41\x0a\x8a\x1a\xaf\xc6\x4b\x91\x32\x67\xb8\x82\x62\xf8\xdd\x03\x69\xb3\x71\xb0\xd8\x00\x66\x35\x7f\x3c\x73\xf0\x8d\xc5\x97\xa4\xd4\xe0\x2a\xe2\xaa\x68\x9f\x32\xff\x74\xca\x29\xde\xa6\xc0\x7e\x86\xb8\x61\x0c\x04\x16\x48\x9f\xa6\x88\x96\x42\x9a\xe0\x62\x07\x9f\x86\x81\xa2\x52\x87\x54\xb4\x36\x16\x1f\x66\xc4\xa4\xe8\xdd\x9e\xe3\x16\xee\xac\xde\xff\x20\xd6\x9e\x32\xae\x6e\xaf\xf1\x89\x66\x41\xca\xb8\x79\xb8\x17\x78\xb7\x88\x1b\x8b\x22\xd4\xa1\x9d\xe1\x03\x53\x54\xa9\x71\xfd\x88\x39\x15\x50\x74\x6d\x20\x68\x4f\xe7\x85\x6f\x1e\xee\xa1\x69\x47\x51\x30\x71\x08\xc6\xa2\x25\x64\x7a\x2e\x9c\xa9\x5b\x1d\xff\xd9\xb9\xf4\x57\xaf\x53\xca\xbb\x65\x97\xf6\x34\x08\x85\x3d\x49\x6d\x8c\x1b\xf9\x1b\x65\xe1\x14\x1b\x08\xf9\x92\x59\xe7\x9a\x45\x39\xd5\x9c\x36\xfb\xcb\x96\xd4\x5d\x9a\x1d\xc7\xae\xc1\x1d\x3d\x17\x12\x5d\x0c\x51\x54\xf6\xab\x23\x33\x90\xba\xce\xa9\x6b\x0c\x10\xdd\x40\x0d\xdc\xc8\x5b\xa2\xdf\xa9\x8c\xce\x53\x83\x8e\x1e\x5d\x09\x6a\x64\x24\xbf\x4c\x0a\x05\xf2\x9a\xf2\x12\x03\x83\x53\xdf\x7f\x76\x2d\x05\x39\x14\x96\x52\x74\x5b\xea\xaa\x76\xfe\x4b\xf0\x69\xd2\xbb\x12\x68\x1d\xaa\xc0\x22\x85\x72\x83\x5e\x75\x94\x66\xb3\x79\x79\xa9\xef\x4a\x54\x1e\xe8\x63\x12\x7d\x7d\xdd\x64\x1a\x92\x52\x25\x94\xf7\xec\x69\xb3\x02\x5b\xe5\x9f\xa6\x9d\x7c\xcd\xfc\xfe\xaf\x9e\xb2\x8a\xb1\xa8\xaa\xca\x58\xbc\x0f\x9a\xc5\x01\xdb\x99\x03\x1a\x8b\x37\x6c\x16\x27\xe0\x8e\x85\x13\x74\x16\x07\x78\x4b\xfb\x0d\x9f\x3d\xc5\x67\xff\x85\x6f\x69\xe4\x03\xbd\x25\xac\xf0\x98\xd3\x70\x9c\xaf\x20\xa9\x64\x4f\xc7\x7c\x79\xf9\x70\xef\xdb\xcc\xd1\xf3\xe8\x82\x34\xf8\x7e\xf1\xdf\xc5\x8f\x5f\x01\x00\x00\xff\xff\x03\x51\x27\x2f\x46\x03\x00\x00"

func istio16RequestAuthenticationYamlBytes() ([]byte, error) {
	return bindataRead(
		_istio16RequestAuthenticationYaml,
		"istio-1.6/request-authentication.yaml",
	)
}

func istio16RequestAuthenticationYaml() (*asset, error) {
	bytes, err := istio16RequestAuthenticationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "istio-1.6/request-authentication.yaml", size: 838, mode: os.FileMode(416), modTime: time.Unix(1600282230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _istio17ApigeeEnvoyAdapterYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4b\x6f\xe3\x36\x10\xbe\xfb\x57\x10\xb9\xd3\x8f\xb6\x7b\x88\x80\x3d\x04\xdb\x45\x1a\xa0\xce\x1a\x4d\xda\x3b\x43\x8d\x25\xa2\x14\x87\x1d\x8e\xd4\xaa\x82\xfe\x7b\x21\x51\xb6\xa5\xc4\xaa\xdd\x20\xa8\x2f\xb2\xe6\xf1\xcd\xc7\x79\x88\xa3\xbc\xf9\x0d\x28\x18\x74\x89\x50\xde\x87\x55\xb5\x59\xfc\x6e\x5c\x9a\x88\x1f\xc1\x5b\xac\x0b\x70\xbc\x28\x80\x55\xaa\x58\x25\x0b\x21\x9c\x2a\xa0\x33\x35\x19\x80\x24\x28\x90\x41\x06\xa0\xca\x68\x90\xe0\x2a\xac\x07\x9b\xe0\x95\x86\x44\x34\xcd\xf2\xf1\xf0\xd6\xb6\x8b\xe0\x41\x77\x28\x04\xde\x1a\xad\x42\x22\x36\x0b\x21\x02\x58\xd0\x8c\xd4\x69\x84\x28\x14\xeb\xfc\x67\xf5\x02\x36\x44\x81\xe8\x98\x5d\x8a\xc9\x50\x78\xab\x18\x06\x8c\x11\xe3\x1e\xc0\x39\x64\xc5\x06\xdd\x11\x53\x88\x60\x52\xd0\x8a\x96\x26\xb0\xc1\xa5\xc1\x15\xc1\x9f\x64\x18\xee\xbc\xff\xe9\xf9\x79\xb7\x23\x7c\x01\x0a\x89\xb8\x61\x2a\xe1\xe6\xe8\xe6\x09\x0b\xe0\x1c\xca\xd0\xf9\x78\xc5\x79\x22\x56\x05\x30\x19\x1d\xe6\x8c\x90\x38\x11\x37\x9f\xd6\xeb\xcd\x1c\x4e\xd0\x39\x74\x99\xcd\x99\xfd\x1c\x4c\xd0\xa4\x3c\x5c\x20\xc4\x75\x67\x72\x92\x75\x29\x1f\x4c\xed\x24\xa9\x57\xa5\x35\xfe\xaa\x43\x8b\x54\x9b\x5e\x76\x28\x63\xff\x1f\x74\x49\x86\xeb\x2f\xe8\x18\xfe\xe2\x44\x88\xa3\x1b\x95\xee\x2e\xfc\x1a\x80\x12\x71\x7b\x7b\x3b\x15\xdf\x13\x96\xfe\x8c\xfc\x11\xdd\x2f\x88\x9c\x88\xee\x8c\x83\x4a\xa3\x63\x65\x5c\x57\x8c\x41\x22\xaf\xea\xc3\xf8\x33\x85\xca\xba\xa4\x65\x88\x99\x85\xd5\xe0\xd2\xdb\x48\x95\x2a\xcf\x40\x49\xd3\x2c\x1f\x3a\xb3\x67\x95\xb5\xed\xcd\xd4\x75\x57\x5a\xbb\x43\x6b\x74\x9d\x88\x87\xfd\x23\xf2\x8e\x20\x74\x63\x71\xcc\x3f\x12\x8f\xb2\x2a\x4f\x7c\x77\x7d\xdd\x3f\xad\xd7\xeb\xa3\xd6\x9a\x0a\x1c\x84\xd0\x77\xd7\xc9\x49\xf4\x85\xbf\x07\x1e\x8b\x84\x18\xba\x2b\x07\x65\x39\xff\x7b\xaa\x3a\x60\x6f\x46\xe2\xbd\x32\xb6\x24\x78\xce\x09\x42\x8e\x36\x8d\xf3\x75\x74\x01\x32\x98\x3e\x81\x46\x97\x76\xa3\x77\x62\x45\xa0\x52\xf3\x7f\xd2\xfa\x7e\x7d\x15\x2f\x45\x59\x18\x87\x96\x42\x4a\x8b\x99\xb4\x50\x81\xfd\x9c\xc2\x4b\x99\xbd\xd2\x6a\x74\x7b\x93\x7d\x5e\xc5\xe7\xf0\x58\xd6\xaa\xb0\xa3\xc3\x06\x2c\x49\xc3\x04\xd9\x9a\xc2\x70\x98\x1e\x53\xfb\xb2\x63\xb3\x2e\x26\xd2\x02\x0a\xa4\xba\x57\x6c\xcd\x48\x43\xf0\x47\x09\x61\x06\xe3\x1a\x88\x0a\x6d\x59\xc0\x16\x4b\x37\x6d\xa7\xa2\x93\xec\x62\xca\xe3\x79\x46\x60\xd7\xcf\x41\xac\xf2\x37\x67\xeb\x61\xb8\x9a\x46\x0a\xb3\x17\xcb\x87\x70\xff\x65\xb7\x55\x4e\x65\x90\xb6\xed\x42\x88\x83\xe2\xce\x29\x5b\xb3\xd1\xe1\x09\x34\x01\xf7\xba\x73\xa4\xd4\xc1\x4e\x86\xde\xf0\x2d\xbd\x79\x83\x57\x9c\x62\x70\xb0\x01\x7a\x06\x5f\x9d\xc6\x14\xd2\xee\x0a\x99\x8d\x8e\x9e\x87\xa1\x5e\xb1\x0d\x6f\x62\xb3\x0d\x32\x66\xf6\x62\x54\x97\xce\x06\xf1\xfd\x07\x60\xee\x7c\x73\xda\x33\x09\x1f\x07\x89\xb4\xde\xf5\x51\x8b\x7d\xb0\x55\x3e\xf9\x4f\xbd\xf0\xee\x92\xcb\x4b\x95\x8c\xaf\x63\x36\x29\xec\x55\x69\x79\x8b\x29\x24\xe2\x87\xef\xc6\x03\x1f\x8d\x1f\x7b\xc4\xa6\x59\x7e\xa3\xac\x6d\x65\xd3\x2c\xbf\xba\xaa\x6d\xe5\x99\x18\x17\x9b\x42\xce\x57\xfb\xdd\xcc\x86\x44\x96\xa9\x56\x91\xdc\x28\xac\x8c\xad\xf6\xba\xa2\xf2\x5f\x1b\xe2\x03\x73\x34\x0d\x70\xa2\x21\xa5\x5c\x8c\x77\xba\xe3\x3a\xf7\x14\xfb\xe0\x03\x77\xb9\xe9\x3e\x71\x71\x97\x38\xec\x0c\xc7\xdb\x52\x9e\xee\x8a\x78\xee\x48\x27\x23\xaf\xdf\xec\x84\x97\xd0\xff\x09\x00\x00\xff\xff\xf7\xa1\x44\x17\xc7\x0a\x00\x00"

func istio17ApigeeEnvoyAdapterYamlBytes() ([]byte, error) {
	return bindataRead(
		_istio17ApigeeEnvoyAdapterYaml,
		"istio-1.7/apigee-envoy-adapter.yaml",
	)
}

func istio17ApigeeEnvoyAdapterYaml() (*asset, error) {
	bytes, err := istio17ApigeeEnvoyAdapterYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "istio-1.7/apigee-envoy-adapter.yaml", size: 2759, mode: os.FileMode(416), modTime: time.Unix(1600282230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _istio17EnvoyfilterSidecarYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x56\x41\x6f\xe3\x36\x13\xbd\xeb\x57\x0c\xec\xe3\xf7\x59\x9b\x74\x51\x14\xd0\xa9\x4e\xd6\xd9\x35\x9a\x4d\x16\x8e\xb7\xdb\x9b\xc2\x50\x63\x89\x0d\x45\xb2\xe4\xc8\xb1\x36\xc8\x7f\x2f\x28\x4a\x89\xa5\x58\xee\x2e\x8a\xa2\x40\x73\x08\x0c\xce\xcc\xe3\x9b\xc7\x37\xa4\xa6\xb0\x54\x8e\x98\x94\x0e\x98\x82\xa5\x23\xa1\x61\xa1\xb6\xba\xbe\x10\x92\xd0\x82\x50\x40\x05\x42\x86\x1b\x56\x49\x02\xc5\x4a\x74\x86\x71\x8c\xa3\x29\x2c\x09\x98\x31\x52\xa0\x03\xd2\xc0\xa4\x04\x87\x76\x2b\x38\xba\xae\xec\x25\x1d\x96\x1b\xa8\x75\x05\x0f\xc2\x15\xff\xf7\xbf\xa2\x29\x94\xac\x06\x62\x42\x6a\x0b\x77\x35\x38\x83\x5c\x6c\x6a\xa1\x72\xb8\x7d\xd0\xf6\x5e\x6a\x96\xdd\xa0\x44\x4e\xda\xde\xc2\x46\xdb\x36\x43\x70\x20\x66\x73\x24\xe7\x39\x7c\x76\xe8\xe0\x96\x19\x91\x23\xce\x2c\x96\x9a\x70\xd6\xb2\x98\xa1\xef\x23\x0e\xb1\xe4\xc7\x93\x93\x93\x5b\x60\xae\x2d\x8e\xa3\x68\x0a\xeb\x42\x38\x10\xae\x01\x0f\xad\x9f\xc6\x3f\xfd\x2f\x8e\x22\x66\xc4\xaf\x68\x9d\xd0\x2a\x01\x85\xe4\xe9\x08\x95\xc7\xc2\xe7\xc4\x42\xbf\xd9\x9e\x32\x69\x0a\xf6\x36\xba\x17\x2a\x4b\xf6\x05\x8b\x4a\x24\x96\x31\x62\x49\x04\x4d\xfb\x09\xf4\xc9\x3d\x3e\xc6\xeb\x86\xc1\x4d\x60\x19\x5f\xb1\x12\x9f\x9e\xda\xec\x46\xac\xa4\x93\x3b\xf2\x1d\x7b\xa0\xa1\x1e\x7e\x0d\x40\xb2\x3b\x94\x2e\xfc\x06\x28\x99\x62\x39\x66\xb3\xbb\xba\xdb\x32\x02\xe0\x5a\x6d\x44\xfe\x89\x11\x2f\xd0\x25\x51\x04\x30\x6b\xce\xac\x5e\xeb\x04\x3e\xac\xd7\x9f\xd2\x8b\xe5\xe5\x7a\xb1\x8a\x02\x02\xf1\xa2\x83\xe3\x5a\x11\xee\x28\x81\x9b\xe5\xbb\xc5\xf9\x7c\x95\x2e\xaf\xce\xae\x3f\x5f\xbd\x6b\xc3\x52\x38\x42\x85\xb6\x4b\x07\xd8\x34\xfd\x9f\x17\x4c\xa8\x97\xc5\x6e\x79\x7f\xa5\xd3\x25\x1c\x4f\x41\x64\x52\xae\x95\x42\x4e\x42\xab\x34\xb4\x61\x7b\xf9\xae\xba\xbb\x38\x00\xd3\x07\xb2\xba\xa2\xb6\xae\xf9\x67\xf6\xbb\xd1\x06\x2d\xa3\xe6\x3c\x97\x57\x37\x8b\xd5\x3a\x3d\x5b\x5c\x5c\xaf\x16\x6d\x78\xcb\x64\x85\x2f\xe0\xfb\xb0\x81\xbf\x6b\x78\xc6\xb8\xa3\x94\x55\x54\x7c\x7d\x4e\xa5\xda\x60\x96\x06\x99\xf7\xd9\x4d\x7e\xf6\x91\x49\xd2\x24\xc4\xb9\xd6\xb9\x44\x66\x84\x8b\xb9\x2e\xdf\x04\xe4\x50\xd4\x6e\x30\xc0\x8f\xb7\x3f\xc4\x8b\x1d\xcd\x7b\x7b\x01\xe4\xd6\xf0\xb4\xb5\x77\x5f\x8b\xb0\x43\xea\x13\x86\x22\x05\xc3\xa7\x95\x15\x43\x2f\x8e\x0e\xca\x00\xc1\x11\xa3\xd4\x58\xdc\x88\xdd\x08\x44\xaf\x80\x44\x89\xba\xa2\x04\x4e\xdd\xde\x7a\x37\x18\x69\xeb\xac\xf4\xd9\xef\xae\xcf\x78\x76\x48\xf9\xdf\x1f\x82\x32\xea\x3f\xed\xe1\x51\xb3\x7d\x8b\xa3\xe7\x17\x9d\x08\xdf\x6e\x68\x7b\xc7\xf8\x3f\xe4\x65\x0f\xed\x6d\xbc\x3a\x9b\x9f\xef\x81\xd9\x4a\x0e\x0f\x9c\xf1\xd0\xc7\xfc\xf2\xf2\xfa\x4b\x2f\x62\xb4\x14\x5c\x0c\xf3\xa1\xf3\x60\x2b\xb9\x7e\xa5\x29\x80\xb1\x42\x71\x61\x98\x7c\x55\xdb\xd8\x47\xd5\x09\x90\xad\xf0\x75\x1d\xda\x52\x38\x7f\xf3\x1f\x2c\x2c\x90\x65\x68\x13\x78\x84\x89\xd7\x74\x92\xc0\xe4\xb7\xd9\x3c\xb0\xf1\xc3\xaa\xad\xf8\x8a\xd9\x04\x9e\x06\x36\xbd\x5a\xac\xbf\x5c\xaf\x7e\xf9\x17\x9d\x3a\x39\x6a\xd5\xc9\x51\x83\x7d\x5c\xac\xde\x8f\x5c\x95\x7f\xe9\x99\xc9\xb8\x69\x70\x47\xa8\x1a\xa9\x9f\x4d\xd9\x3e\xb6\x63\x34\xe3\xed\xdb\xf8\x03\x91\x39\x7f\x8e\x7c\xdc\xe7\xdf\xd9\x89\xa3\x73\xa9\xd4\x3d\x42\xb3\xde\x10\xbc\xe4\xe4\xdd\x2c\x34\x77\x67\xff\x1e\x1b\xe9\x0d\xbe\x6b\x26\xc2\x56\x52\xe7\x7e\x18\x3c\xf9\xf7\xd6\xf0\x79\xb3\x78\xa9\xf3\xf3\x26\x69\x00\xce\x75\x59\x6a\x35\xb2\x33\x80\xd4\x79\x7a\xe8\xd3\xe2\xd0\x5d\x0c\x47\x5f\x8d\x36\x3e\xfe\x76\xb4\x42\xfc\xed\x17\x24\xfc\x7d\xe7\x3b\x02\xc0\xb2\x4c\xf8\x53\x66\x32\xb5\xf8\x47\x85\x8e\xd2\x30\x81\x2e\x25\x3d\x3c\xe1\x70\xca\x09\x0b\x63\x48\x35\x4c\x9f\x3f\x5b\x43\x03\xed\xf4\x0e\x4a\xa6\xdd\xf0\xb5\x61\xf7\x0a\x72\x37\x6b\xb9\x86\xa3\x24\x7d\x8f\xea\x48\x92\x11\x47\x83\xc6\xea\xac\xe2\x74\x6c\x1b\xff\x51\xcd\x9b\xd9\x1b\x4f\xe2\x52\xa0\x22\x91\x8d\x67\x64\xb8\x45\xe9\xa7\x18\x4b\x26\xe4\x9f\x01\x00\x00\xff\xff\xc6\x6f\x50\xc8\xe9\x0b\x00\x00"

func istio17EnvoyfilterSidecarYamlBytes() ([]byte, error) {
	return bindataRead(
		_istio17EnvoyfilterSidecarYaml,
		"istio-1.7/envoyfilter-sidecar.yaml",
	)
}

func istio17EnvoyfilterSidecarYaml() (*asset, error) {
	bytes, err := istio17EnvoyfilterSidecarYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "istio-1.7/envoyfilter-sidecar.yaml", size: 3049, mode: os.FileMode(416), modTime: time.Unix(1600282230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _istio17HttpbinYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x4f\x6b\xe3\x30\x10\xc5\xef\xfe\x14\x03\x7b\x76\xb2\xb9\x2d\xba\x2d\xec\x65\xa1\x14\x43\xa1\xf7\x89\xfc\xe2\x88\xe8\x1f\xd2\x24\x34\xfd\xf4\x45\x76\xe4\x38\x6e\x2e\x9d\x9b\x66\x1e\x3f\xbd\x37\xf3\x8b\xfe\x7a\x3a\x8a\xc4\xbd\xf1\x24\x9c\x06\x08\xe1\x83\x5d\xb4\xa0\x7f\x88\x36\x5c\x1d\xbc\x10\xfb\x9e\xde\x90\x2e\x46\x63\xd3\x70\x34\xef\x48\xd9\x04\xaf\xe8\xb2\x6b\x4e\xc6\xf7\xaa\x4e\x1b\x07\xe1\x9e\x85\x55\x43\xe4\xd9\x41\x55\xfa\xed\x9d\x23\x6b\x28\xea\x71\xe0\xb3\x95\x86\xc8\xf2\x1e\x36\x17\x39\x11\xc7\x78\xd7\xe7\x08\x5d\xda\x31\x24\x19\xe7\xed\x02\x38\xca\xcb\x44\xd1\x9f\xdf\xe3\x63\x32\xdf\xdd\x5b\x19\x16\x5a\x42\x7a\x82\x6e\xdb\xf6\x21\x05\xc7\x98\xb7\x73\x94\x7b\xee\x9f\xa6\xa9\x96\x13\xa2\x35\x9a\xb3\xa2\xdd\x37\x23\x8e\x45\x1f\x5f\x16\xa1\x57\xde\xa6\xd6\x65\xb1\x60\x22\x81\x8b\x96\x05\x37\xc2\xc2\x54\x29\xfb\x00\x7b\x8a\x5b\x03\xa7\x72\xec\x79\x40\xdf\xee\xaf\x65\x05\x66\x00\xc6\x51\x4d\x51\x4a\x07\x2f\x6c\x3c\xd2\xcc\x6f\xc9\x38\x1e\x4a\xe8\xa0\x4f\x48\x1b\x13\xb6\x27\x78\x0f\x39\x26\x18\xf9\xdc\xae\x3f\x1e\xd5\xdd\xd9\xda\x2e\x58\xa3\xaf\x8a\xfe\x1f\x5e\x83\x74\x09\xb9\x2c\xb8\xaa\xd6\xcb\x9d\x6a\x3e\x7d\xfd\x7a\xb6\x53\xef\xfc\x15\x00\x00\xff\xff\x35\x03\x1d\x9f\xc0\x02\x00\x00"

func istio17HttpbinYamlBytes() ([]byte, error) {
	return bindataRead(
		_istio17HttpbinYaml,
		"istio-1.7/httpbin.yaml",
	)
}

func istio17HttpbinYaml() (*asset, error) {
	bytes, err := istio17HttpbinYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "istio-1.7/httpbin.yaml", size: 704, mode: os.FileMode(416), modTime: time.Unix(1600282230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _istio17RequestAuthenticationYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xbd\x6e\xdc\x3a\x10\x85\x7b\x3e\xc5\x81\xd9\x5d\x5c\x69\xe1\x56\x9d\x53\xc5\x4e\x10\x18\x86\x13\x17\x41\x0a\x8a\x1a\xaf\xc6\x4b\x91\x32\x67\xb8\x82\x62\xf8\xdd\x03\x69\xb3\x71\xb0\xd8\x00\x66\x35\x7f\x3c\x73\xf0\x8d\xc5\x97\xa4\xd4\xe0\x2a\xe2\xaa\x68\x9f\x32\xff\x74\xca\x29\xde\xa6\xc0\x7e\x86\xb8\x61\x0c\x04\x16\x48\x9f\xa6\x88\x96\x42\x9a\xe0\x62\x07\x9f\x86\x81\xa2\x52\x87\x54\xb4\x36\x16\x1f\x66\xc4\xa4\xe8\xdd\x9e\xe3\x16\xee\xac\xde\xff\x20\xd6\x9e\x32\xae\x6e\xaf\xf1\x89\x66\x41\xca\xb8\x79\xb8\x17\x78\xb7\x88\x1b\x8b\x22\xd4\xa1\x9d\xe1\x03\x53\x54\xa9\x71\xfd\x88\x39\x15\x50\x74\x6d\x20\x68\x4f\xe7\x85\x6f\x1e\xee\xa1\x69\x47\x51\x30\x71\x08\xc6\xa2\x25\x64\x7a\x2e\x9c\xa9\x5b\x1d\xff\xd9\xb9\xf4\x57\xaf\x53\xca\xbb\x65\x97\xf6\x34\x08\x85\x3d\x49\x6d\x8c\x1b\xf9\x1b\x65\xe1\x14\x1b\x08\xf9\x92\x59\xe7\x9a\x45\x39\xd5\x9c\x36\xfb\xcb\x96\xd4\x5d\x9a\x1d\xc7\xae\xc1\x1d\x3d\x17\x12\x5d\x0c\x51\x54\xf6\xab\x23\x33\x90\xba\xce\xa9\x6b\x0c\x10\xdd\x40\x0d\xdc\xc8\x5b\xa2\xdf\xa9\x8c\xce\x53\x83\x8e\x1e\x5d\x09\x6a\x64\x24\xbf\x4c\x0a\x05\xf2\x9a\xf2\x12\x03\x83\x53\xdf\x7f\x76\x2d\x05\x39\x14\x96\x52\x74\x5b\xea\xaa\x76\xfe\x4b\xf0\x69\xd2\xbb\x12\x68\x1d\xaa\xc0\x22\x85\x72\x83\x5e\x75\x94\x66\xb3\x79\x79\xa9\xef\x4a\x54\x1e\xe8\x63\x12\x7d\x7d\xdd\x64\x1a\x92\x52\x25\x94\xf7\xec\x69\xb3\x02\x5b\xe5\x9f\xa6\x9d\x7c\xcd\xfc\xfe\xaf\x9e\xb2\x8a\xb1\xa8\xaa\xca\x58\xbc\x0f\x9a\xc5\x01\xdb\x99\x03\x1a\x8b\x37\x6c\x16\x27\xe0\x8e\x85\x13\x74\x16\x07\x78\x4b\xfb\x0d\x9f\x3d\xc5\x67\xff\x85\x6f\x69\xe4\x03\xbd\x25\xac\xf0\x98\xd3\x70\x9c\xaf\x20\xa9\x64\x4f\xc7\x7c\x79\xf9\x70\xef\xdb\xcc\xd1\xf3\xe8\x82\x34\xf8\x7e\xf1\xdf\xc5\x8f\x5f\x01\x00\x00\xff\xff\x03\x51\x27\x2f\x46\x03\x00\x00"

func istio17RequestAuthenticationYamlBytes() ([]byte, error) {
	return bindataRead(
		_istio17RequestAuthenticationYaml,
		"istio-1.7/request-authentication.yaml",
	)
}

func istio17RequestAuthenticationYaml() (*asset, error) {
	bytes, err := istio17RequestAuthenticationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "istio-1.7/request-authentication.yaml", size: 838, mode: os.FileMode(416), modTime: time.Unix(1600282230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nativeEnvoyConfigYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x5f\x93\xdb\xb6\x11\x7f\xd7\xa7\xd8\x91\x5f\xda\x4e\xc4\x93\xcf\xce\x34\xe5\x4c\x66\xaa\x9c\xdb\xc4\xb6\x7a\x97\xb9\x73\xea\xf6\x09\x03\x83\x2b\x09\x77\x20\x80\x02\xa0\xee\x64\x8d\xbe\x7b\x07\x00\x29\x91\x14\xa8\x53\xda\x78\xd2\x4e\xf8\x46\xec\x3f\x60\xff\xfc\x76\x01\x5a\x94\x5c\xe6\x23\x00\xca\x18\x5a\x4b\x84\x5a\x12\x4d\xdd\x2a\x87\x0b\x57\xea\x0b\x94\x6b\xb5\x21\x81\x29\x13\x6a\xe9\xf9\x8a\xc2\xa0\xb5\x5e\x04\xc0\x2a\xf6\x80\x8e\x74\xd6\x0e\x2c\xf0\xf2\xf2\x8f\xd9\x34\x9b\x66\x2f\x6b\x82\x56\xc6\x91\x35\x15\x15\xe6\xf0\xa7\xe9\x74\x3a\x02\x18\x59\x47\x1d\x67\xc4\xa0\x55\x95\x61\x18\x94\x08\x6e\x1d\x4a\x34\xe1\x67\x02\x3d\xf5\x3d\xa3\xb0\x3d\x18\x9c\x06\x73\xd3\xaf\x3a\x96\xbe\x99\x7e\x33\x85\x9d\xb7\xe5\xa5\x17\x5c\x38\x34\x84\xad\x28\x97\xb5\xca\x49\xbd\xb8\xb7\x30\x01\x49\x4b\xcc\x21\x9c\x3e\xab\x89\x99\x44\xf7\xa8\xcc\x43\xb6\x72\x4e\x13\xa6\xa4\x44\xe6\xb8\x92\xa4\xa4\x92\x2e\xd1\xd4\xb2\x00\x6e\xa3\xb1\xf0\x0c\x0b\xbe\xcc\x6b\xab\xf1\x1b\xff\xd9\xd3\xc6\x79\x60\xc9\x96\x4a\x2d\x05\x52\xcd\x6d\xc6\x54\x19\x5d\x9d\xe1\x93\x43\x69\xb9\x92\xf6\x5c\xbb\xd9\xfa\x55\xf6\x83\x73\xfa\x6a\x4f\xf9\x5b\x6f\x43\x00\xde\xc9\x44\x1b\x5c\xf0\xa7\x1c\xb8\x5c\x7a\x6f\x11\xaf\xaf\xc5\x63\x54\xe5\xb0\xd9\x76\x6b\x1d\x60\xcd\x8d\xab\xa8\x20\x2b\x65\x9d\xed\x92\x1a\x4f\x15\xb8\xa0\x95\x70\x1d\x1a\x40\xa1\xca\xe0\x65\x18\xff\x61\xdc\x23\x05\x6b\x3d\x65\x5e\x5d\x49\x1d\x5b\xf9\x98\x36\xbb\xbd\x08\xa1\x4b\x08\xf7\x65\xfd\xc7\x44\x65\x1d\x9a\x1c\xb6\xdb\xec\x03\x35\x4b\x74\x77\x68\xd6\x9c\x61\x76\x4d\x4b\xdc\x1d\x6b\x8a\xb1\xd2\x68\x48\x93\x17\x89\xf3\xc7\xaf\x9b\x0c\xde\x79\x59\xb1\x91\xb4\xe4\x8c\x2c\x94\x79\xa4\xa6\x20\xda\xa8\xa7\x4d\x4a\xf6\x3f\x0d\xfd\xb0\x15\x1f\xf5\x1f\xd1\xdc\x7a\x4f\x5c\x85\x3d\x27\xcd\xfa\x90\x11\x83\x8f\x86\x3b\x24\x82\x3b\x34\x54\x24\xbc\xf3\x83\xb2\x6e\xb7\x03\x78\x01\x35\xaf\x85\x9c\x56\x6e\xa5\x0c\x77\x1b\x58\x28\x03\x33\xcd\x97\x88\xa0\x8d\x2a\x2a\xe6\x80\xad\x90\x3d\x8c\x3a\xc9\x1d\xd2\xb3\x55\x48\xa3\x76\x58\x53\xf5\x34\x7c\xb8\xd1\x71\x88\xd2\x71\x39\xc3\xab\x51\xb0\x36\x7a\xd2\xa1\x97\x54\xe8\x15\xcd\xfe\x1a\x38\x93\x3e\x2d\xa4\x25\x8c\xb2\x55\xba\x4c\xfc\x57\x57\x43\xca\x40\x47\xf2\x48\xd0\xab\x16\x4a\x3d\x54\x9a\x2c\x68\xc9\xc5\x26\x87\xbf\xbf\x26\x37\xd7\xf3\x7f\x76\x5d\xf9\x02\xd0\xc3\x1a\x75\x08\xef\x3e\x7e\x00\xa7\x1e\x50\xda\xaf\x80\x0a\xa1\x1e\x49\xc9\xad\xe5\x72\x19\xff\x2c\xcc\x7e\x7c\x0b\xef\x71\x03\x54\x58\x75\x5e\x34\xee\x1f\x1d\xf1\x71\x97\xc3\x11\x80\x5f\x24\x04\x7b\x43\x7b\xb7\xbf\x7b\x74\xb3\xca\xad\x50\x3a\xce\xa8\x47\xb1\x9e\x1d\x6d\xd4\x9a\x17\x2d\x90\x3e\x7c\x34\xe4\x66\xaa\xee\xb8\xb5\x95\x87\x03\x6f\xd3\xe6\x17\x17\xdb\x6d\x76\x5b\x49\xc7\x4b\x8c\x19\x7f\x61\xb0\x54\x0e\x27\x36\xd6\xc1\x45\x70\x67\x42\x0f\xad\x0a\x8e\x92\x1d\x03\x56\xf4\x67\x57\xcb\x84\x09\x8e\xb2\x8f\x85\xfe\x8b\x7c\xe4\xfe\xf1\x21\xa9\xa8\xae\xa1\xca\xf0\x34\x15\xc0\x93\xce\x3e\x0c\x43\xe3\xec\x80\xa2\x3d\x50\x46\xdf\x4d\x7c\x2c\x1a\xc1\x01\x11\x6f\x48\x55\x2e\x87\xaf\xd3\x4a\x63\x7a\x17\x95\x09\xd1\x1b\x3a\x81\x45\xa6\x64\x61\x73\x78\x15\xda\x7f\xff\xd3\x74\x23\x14\x2d\x08\x97\xa4\x44\x47\x0b\xea\x68\xb3\xc7\x7e\x03\xa9\xc4\x70\xff\x48\x69\x6e\xfa\xc9\x71\x37\xc1\x7f\x55\xdc\xa4\x63\xdb\xd0\x08\x95\x03\xc0\x5e\x73\x94\x28\xfb\xbd\xf1\xb0\xa9\x26\x77\x49\x2c\xbc\xe4\x79\x1a\xde\x4e\x25\xe7\xb0\xdd\x0d\xd6\x7f\x0d\xc9\xc1\x11\xe7\x95\x37\x3e\xc5\xaa\xfb\xfc\x85\x01\x76\x6f\x27\x5b\x5f\x66\x7f\x79\x0a\x85\xfd\xb9\xa7\x79\x69\x34\x23\x75\xc6\x1d\xfb\x2d\xce\x9d\x9e\xe7\x44\x9b\xef\xb8\x73\xd2\x2b\xc2\xa0\xe1\xb8\xdf\x37\x49\xfc\xd2\xf6\x49\x3d\xde\x26\xfd\xbc\x63\x9c\x3f\x90\x37\x66\x35\x4d\x80\xc0\xe4\x39\x28\x4d\xc7\xef\xf6\xbb\xd9\x15\xfc\x4e\xa2\x9f\xbb\xa9\xe9\xb4\xd9\x06\x6b\x0d\xde\x23\x73\xa4\x92\x75\x33\xfe\x8c\x05\x7c\xfb\x2d\x2c\xa8\xb0\xf8\xfb\xf3\x62\x6e\x3e\x51\xf6\x85\xc3\xed\x4d\xf8\x48\xfb\x03\x9d\x53\xa5\xfe\xb2\x11\x30\x02\x66\xf3\xf9\xcd\xc7\x23\xaa\x56\x82\x33\x9e\x2e\xc8\x3a\xd8\xf5\x18\xac\x4c\xba\xe4\xb4\xe1\x92\x71\x4d\xc5\x60\x49\xfa\x7a\x06\x67\xaa\x74\x19\x6a\x34\xa1\x02\x95\x1c\x54\xb0\x42\x5a\x84\x49\x13\xc6\xde\xf7\xe3\x1c\xc6\xff\x98\xc4\xf0\x4d\x66\xfb\x68\x8d\x61\x77\xf6\x28\x14\xa6\x5a\xf3\x4b\xc6\x6a\x68\xa2\x8c\x96\xfc\x08\x79\xdb\xd8\x6c\x29\x3f\xdc\x04\x43\xb3\xef\x24\x2f\x53\x42\x20\x73\x4d\x9e\x52\x49\xc5\xc6\x71\x36\x8c\x3f\x07\x65\xcb\xc6\x7c\x28\xeb\x2f\x90\x92\xd1\x94\x50\x4b\x9f\x8b\xfe\x46\xf4\xbd\xd1\x6c\x16\x16\xe7\x6a\x99\x1c\xe8\x98\x2a\x4b\x25\x07\xa7\xb9\xd3\x18\xf5\x1c\x4a\xfd\x37\x38\xe5\xaf\xe1\x3f\x57\x8a\x16\x05\xf7\x65\x45\x05\xf1\x4d\x09\xad\x23\x31\x49\x2d\x71\x2a\x84\xf3\x08\xb7\x5a\x43\xfe\x8b\xe6\x0e\x07\x2e\x5c\x0d\xea\x04\xef\x89\xf8\x0c\x08\x70\x58\x93\xfb\xa3\xc0\x04\x9e\x26\xcd\x48\x11\x3c\x9f\x9a\xa8\xda\x4c\x9a\x9f\x24\xd6\xf7\x8d\x53\x66\xb4\x16\xe9\x89\xb1\xc5\x14\x27\x32\x5e\x0c\x73\x14\xb8\x46\xa1\x34\x1a\x2c\x29\x17\x31\xed\xeb\xf0\xd5\x10\x30\x82\xe8\x23\x2e\xb1\xa1\x04\xc8\x1e\xba\x4c\x45\x37\x8e\x0e\x05\x71\xea\x4e\x5a\xe3\x19\xd9\xf7\xa7\xcb\x78\x64\x9f\xf4\x39\xcc\x6f\xbe\x7f\x7b\x35\x9b\x93\x37\xd7\x77\x61\xf5\xe4\x65\x01\x40\x7c\x22\x01\x42\x37\x39\xdc\xde\xfc\x74\xfd\x86\xdc\xde\x7c\xf7\xf6\x3a\x92\xfc\x70\x45\xad\xe5\x4b\xe9\x27\x96\x26\x23\xba\x89\x7a\xfa\xf2\x8c\xb2\xd0\x8a\xb7\xa6\x9d\x89\x37\x78\xb4\x1a\xdb\x62\x5c\xec\x26\x5e\xef\x39\xa7\xf9\xd2\x6f\x49\xc7\x62\x43\x0e\x4f\xb4\x91\xc3\x1b\xd0\xeb\xd7\xaf\xa2\x3f\x0d\x95\x36\x10\xa2\xb9\xc6\x4c\x1b\xb2\xfa\x3c\x36\x73\xa2\x49\xc0\x34\x52\xfd\x3c\x28\x4e\xea\xf7\x60\xfc\x93\xb6\xce\x20\x2d\x3f\x08\x7b\x15\xab\x6c\x6f\xc1\x4a\x7e\xea\xe0\x43\xb9\xd9\x8c\x88\x01\x3c\xe0\x30\xdf\x37\x29\xf9\x1c\xb8\xc4\xf4\x8b\x4f\x74\x61\xc1\xe3\xf7\xa5\xbf\xcf\x3a\xc5\x94\x20\x4a\xbb\xd0\x23\xe3\x98\x7a\x66\x7a\x3d\x8f\x68\xbf\x7a\x8a\xf5\x1f\x2e\xd3\x49\xf5\xf5\xb4\xbe\xc3\xd4\x6d\x44\x7c\xea\x65\xc6\x0a\xa9\x70\xab\x0d\xd1\x54\x72\x46\xdc\xca\xa0\x5d\x29\x51\x1c\x0c\xef\x15\x65\x51\x51\x14\x20\xe1\x75\xa5\x75\xf8\xd6\xd4\xba\x17\xe5\xd2\xa1\x59\x53\xd1\xb9\x8f\x35\x8b\xe4\x9e\xbb\x70\xbd\x6b\x09\x48\x45\x9c\xa1\x8b\x05\x67\x24\x29\x5b\xc9\x66\xbb\x87\x8d\xc2\xc1\x01\x09\xe2\xab\x3d\x31\x34\xca\xf6\xe6\xf7\x19\x71\x04\x6c\xd3\xec\xb2\x36\xba\xdd\xf2\x05\x64\x1f\xe6\x77\xd9\x1b\x6e\x76\xbb\x17\x21\x63\x59\x65\x9d\x2a\xe1\xee\x6e\x0e\x87\xa7\x4e\x70\xaa\x77\xc7\x4e\x17\x33\xfc\x9f\x54\x73\x9d\x2f\x4e\xd8\xe6\x76\xd1\xce\xc5\xb0\x8c\xc6\xf1\x85\xef\x6b\x68\xf3\xee\xfc\xd8\x22\xc5\x57\xec\x08\x0c\xf3\xbb\xec\xca\xf4\x61\x50\x1b\xbe\xf6\x7c\x0f\xb8\xd9\x73\xbd\xc7\xcd\x6e\xb7\xdd\xa2\x2c\xce\x81\x8e\x77\x1f\xdf\xdf\x41\xf3\x92\xd0\x83\x8d\xa3\x47\x83\xff\xa9\x2e\x36\xfc\xb2\xf1\xab\xc3\x4b\xff\xe5\xe6\x37\xd1\xba\x3a\x27\xfe\x77\x00\x00\x00\xff\xff\x3a\xa8\x4c\xa3\x77\x1a\x00\x00"

func nativeEnvoyConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_nativeEnvoyConfigYaml,
		"native/envoy-config.yaml",
	)
}

func nativeEnvoyConfigYaml() (*asset, error) {
	bytes, err := nativeEnvoyConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "native/envoy-config.yaml", size: 6775, mode: os.FileMode(416), modTime: time.Unix(1600282230, 0)}
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
	"istio-1.6/apigee-envoy-adapter.yaml":   istio16ApigeeEnvoyAdapterYaml,
	"istio-1.6/envoyfilter-sidecar.yaml":    istio16EnvoyfilterSidecarYaml,
	"istio-1.6/httpbin.yaml":                istio16HttpbinYaml,
	"istio-1.6/request-authentication.yaml": istio16RequestAuthenticationYaml,
	"istio-1.7/apigee-envoy-adapter.yaml":   istio17ApigeeEnvoyAdapterYaml,
	"istio-1.7/envoyfilter-sidecar.yaml":    istio17EnvoyfilterSidecarYaml,
	"istio-1.7/httpbin.yaml":                istio17HttpbinYaml,
	"istio-1.7/request-authentication.yaml": istio17RequestAuthenticationYaml,
	"native/envoy-config.yaml":              nativeEnvoyConfigYaml,
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
	"istio-1.6": &bintree{nil, map[string]*bintree{
		"apigee-envoy-adapter.yaml":   &bintree{istio16ApigeeEnvoyAdapterYaml, map[string]*bintree{}},
		"envoyfilter-sidecar.yaml":    &bintree{istio16EnvoyfilterSidecarYaml, map[string]*bintree{}},
		"httpbin.yaml":                &bintree{istio16HttpbinYaml, map[string]*bintree{}},
		"request-authentication.yaml": &bintree{istio16RequestAuthenticationYaml, map[string]*bintree{}},
	}},
	"istio-1.7": &bintree{nil, map[string]*bintree{
		"apigee-envoy-adapter.yaml":   &bintree{istio17ApigeeEnvoyAdapterYaml, map[string]*bintree{}},
		"envoyfilter-sidecar.yaml":    &bintree{istio17EnvoyfilterSidecarYaml, map[string]*bintree{}},
		"httpbin.yaml":                &bintree{istio17HttpbinYaml, map[string]*bintree{}},
		"request-authentication.yaml": &bintree{istio17RequestAuthenticationYaml, map[string]*bintree{}},
	}},
	"native": &bintree{nil, map[string]*bintree{
		"envoy-config.yaml": &bintree{nativeEnvoyConfigYaml, map[string]*bintree{}},
	}},
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
