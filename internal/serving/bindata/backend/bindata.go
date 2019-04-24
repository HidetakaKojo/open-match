// Code generated by go-bindata. DO NOT EDIT.
// sources:
// api/protobuf-spec/backend.swagger.json (15.868kB)

package backend

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _backendSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5f\x6f\xdb\x38\x12\x7f\xcf\xa7\x18\xf8\x0e\x48\x0a\xb8\xce\x5e\xef\x70\x0f\x7d\xcb\xa5\xed\x5e\x80\xcd\x35\x48\xb2\x38\x1c\x56\xc5\x82\x96\x46\x16\x37\x14\xa9\xf2\x4f\x52\x63\xd1\xef\x7e\xe0\x90\xb2\xfe\x58\x4a\x6c\x27\xf1\xa6\xe8\xf6\x25\xaa\x44\x8e\x66\x86\x33\x3f\xfe\x34\x1c\xff\x7e\x00\x30\x31\x77\x6c\xb1\x40\x3d\x79\x0b\x93\x37\xb3\x1f\x26\x53\x7f\x8f\xcb\x5c\x4d\xde\x82\x7f\x0e\x30\xb1\xdc\x0a\xf4\xcf\x59\xc5\x8f\x2b\xad\xac\x9a\xbb\xfc\xb5\xa9\x30\x3d\x9e\xb3\xf4\x06\x65\x36\xa3\xbb\x34\x17\x60\x72\x8b\xda\x70\x25\xfd\x8c\x78\x09\x52\x59\x30\x68\x27\x07\x00\x5f\xe9\x0d\x26\x2d\xb0\x44\x33\x79\x0b\xbf\x84\x49\x85\xb5\x55\x2d\xc0\x5f\x1b\x3f\xf6\x13\x8d\x4d\x95\x34\xae\x33\x98\x55\x95\xe0\x29\xb3\x5c\xc9\xe3\xdf\x8c\x92\xcd\xd8\x4a\xab\xcc\xa5\x1b\x8e\x65\xb6\x30\x8d\x99\xc7\xb7\x7f\xab\xed\x39\x66\xc6\xf0\x85\x2c\x51\xda\x66\x00\xc0\x24\x43\x81\x16\x5b\x77\xbc\x25\xae\x2c\x99\x5e\x7a\x6b\x2f\xb1\x54\xb7\x08\xef\x7e\xbc\x82\x54\x49\x89\xa9\x7f\x29\x78\x67\x42\xae\x55\x09\xc6\x32\x8b\x60\xac\xd2\x6c\x81\x90\x2b\x0d\x95\x60\x4b\xd4\x66\x96\xc8\xb3\xff\x5c\xfc\x7c\xfd\x16\x2e\x95\xb1\xa8\xa1\x44\x63\xfc\x98\x3b\x6e\x0b\xb0\x05\xc2\x61\x1c\x79\x08\x39\x47\x91\x41\xa5\x2a\x27\x98\xc5\x6c\x96\x48\x00\xb8\x2e\x10\x94\x14\xcb\xf8\x94\x4b\x9a\x14\x84\x1d\x1a\xb8\xa0\xc9\xb5\x54\x03\xce\x60\x06\xf3\x25\x4d\x7d\x47\x36\x9d\x34\x06\x03\x37\xe1\x95\x3c\x8b\x6f\x9b\x01\x9c\x08\x01\xca\x16\xa8\x0d\x30\x8d\x60\xb8\x40\x69\xc5\x12\xf8\x42\x2a\x8d\x7e\xc4\x59\x4e\xe2\x96\xca\x81\x44\xcc\xc0\x2a\x08\xee\x82\xd2\x09\xcb\x2b\x81\xa0\x49\x1f\x33\x85\x92\xdd\xb4\x6e\xa7\x4c\x08\x33\x8b\x8b\x4f\x3e\x55\x15\x6a\x5a\xb1\xb3\xcc\xfb\x75\x4d\xc5\x37\xed\xc1\x1a\x4d\xa5\xa4\x41\xd3\x59\x17\x80\xc9\x9b\x1f\x7e\xe8\xdd\xa2\x25\x34\xa9\xe6\x95\x8d\x01\x7a\x02\xc6\xa5\x29\x1a\x93\x3b\x01\xb5\xa4\xb6\x2e\x61\x8d\x7d\xb4\xb2\x35\x61\x00\x93\xbf\x6a\xcc\xbd\x9c\xbf\x1c\x67\x98\x73\xc9\xbd\x5c\x73\xcc\x2a\xbe\xa6\xf3\x65\x14\x3e\xe9\x88\xf8\x7a\x30\x74\xfd\xb5\x65\x5e\xc5\x34\x2b\xd1\xfb\x6d\x15\xd3\xe1\x5f\xcf\x30\xc9\x4a\x4a\xd2\xe0\xe4\x19\xfd\xb7\x67\x06\x27\x93\x3f\x3b\xd4\xcb\xfe\x23\x8d\x9f\x1d\xd7\xe8\xdd\x9d\x33\x61\xb0\xf7\xd8\x2e\x2b\x12\x6e\xac\xe6\x72\x31\x19\x54\xfa\x53\x4b\x69\xcb\x16\x7d\x75\x27\xff\x0a\xb9\xd5\x4c\xfe\x74\xd0\xb3\x76\x52\x39\x3b\x9a\x5c\xff\xd5\xdc\x22\x05\xe6\x5a\x6e\x29\x4d\xf7\x05\x37\x16\x54\x5e\x27\x55\xcc\x82\x44\xb6\x16\x61\x56\xa7\xc0\x2c\xe4\x86\xf1\x61\xda\x49\xcb\x59\x48\xa6\x0f\x5a\x49\x8b\x32\x83\x93\x8b\x33\xe0\x26\x91\x31\x36\xf8\x5c\x84\xcc\x35\x28\x33\x2e\x17\xc0\xe4\xd2\x16\xfe\xc2\xa0\xb4\x50\xa0\x46\x2f\xd2\xab\xb3\x60\x25\x42\x2a\x38\xbd\x37\x91\x57\xf5\x84\xa8\x5f\x18\xc6\x0d\xe4\x4e\x06\x63\x6e\x78\x7a\x63\x40\xe5\xb9\x1f\xa3\x95\x8f\x4a\xb0\x05\xb3\xa0\x09\x57\x4c\x22\xbd\xd8\x38\x9b\x00\x25\x77\xd6\x69\x84\x92\xd9\xb4\x28\xd9\x8d\x17\x5f\x4b\x33\x30\x5f\x02\xcb\xe8\x95\xb6\xc0\x32\x6a\x95\xc8\xc3\x0c\xb9\xcc\xf0\x0b\x66\x87\xb5\x2c\x72\x1c\x93\x99\x1f\x20\x43\xd6\xc6\x69\xbc\x86\x28\x38\x7b\x37\x80\x61\x89\x0c\xa2\x5a\x00\xd6\x06\x92\x3e\x8a\x19\x0c\x78\x62\x1a\xf8\x7a\x9b\x48\x78\x0d\x0d\xdc\x4e\x1b\x87\x7a\x24\xb9\x0b\x8b\x1e\x3d\x15\xa2\xcf\xe3\x13\x39\xdb\xaa\xce\x32\x91\xa4\x08\x31\x33\xf8\x9f\x72\x90\x32\x49\xeb\xe4\x65\x82\x74\xe5\x1c\xb5\x8f\x8f\x15\x0c\xa5\x4a\x5a\xc6\x65\x5c\xc6\x66\x04\x01\x59\x6d\xf8\x2a\x62\x08\x03\xeb\xd8\x22\x5f\x30\x21\x6a\x61\x70\xc7\x85\x80\x39\x46\xcd\x0a\x6c\xd9\x14\x20\x7a\x47\x8c\x86\x53\x8d\xac\x8f\xcf\x41\xa0\x9f\xce\xb3\x75\x88\xae\x9d\x1c\xdf\xd0\x17\x3c\x88\xdf\xf7\xa0\xef\x9a\x02\x2f\x1f\x7c\xd7\x54\xde\x17\xf8\xce\x55\xb6\x06\xad\x01\x75\x87\x9e\xb4\x40\xd7\x6a\xd7\xc7\xdc\xa7\xb1\xfb\xb3\x43\x63\x37\x31\x7b\x77\xf8\x3e\x68\x79\xad\xc3\xa2\x08\x97\x70\x1b\x06\x15\x76\x4d\x60\x70\xee\xa7\x7e\x9c\xff\x86\xa9\x1d\xe2\x4e\x25\x93\x8e\x09\xb1\x9c\xc1\x51\x6b\x24\x85\x3c\x0d\x4c\x64\x3d\x92\xf2\x92\x09\xa3\x80\x39\xab\x4a\x66\xb9\xa7\x1b\x4b\xc0\x2f\x15\xd7\x08\x2c\xf7\x74\x8b\xc1\x5d\xc1\x05\x4e\x81\xfc\x89\xab\xe4\x4c\x95\xcc\xf9\xe2\xd5\x0a\xdc\xda\x5a\xad\x53\xb4\x15\x5f\xea\xb0\xb3\xa3\xb5\xbc\xf4\x09\x58\xe7\xdd\xab\x07\x69\x0f\xbd\xf3\xe5\xa7\x5c\x4b\xd9\xef\x24\xd9\x3a\x16\xef\x21\xcd\x36\x62\x49\x97\x4e\xc2\xf9\xf9\x07\x50\x32\xf5\x54\xe6\x12\xad\xd3\x12\x58\xa0\x08\x2a\x44\x2e\x71\x8a\x9c\x5b\x13\x76\xd5\x4a\xab\x9c\x0b\x9c\x6d\x14\xe5\xa3\x5b\x38\xcf\xe8\x4f\xa5\x7d\x10\x5b\xee\xd9\x0a\xbc\x86\x5f\x14\x45\x1e\x13\x9f\xe2\x46\x49\xfb\x7b\x2d\xc2\xef\xf0\x39\xe5\xa7\xcf\xc4\x5b\xc6\x05\xf3\xf4\xca\x2a\xff\x44\x7b\x3b\x66\x7d\x29\x95\x52\xc2\x6c\x2f\xe4\xe3\xcf\xd7\x4f\x6d\x1a\x6a\xad\xf4\x0c\xde\x97\x95\x5d\x02\xcf\x41\xaa\x70\x0b\xee\x98\x01\x94\xa9\x72\xd2\xa2\xc6\xac\x4d\x4b\xa6\x7e\xa0\x57\x38\x2d\x94\x32\xa4\x24\xa9\x4e\x04\x8d\xcb\x46\x63\x38\xba\xc4\x54\x95\x25\xca\x0c\xb3\x57\xe1\xf5\xc1\xf2\x28\x80\xd8\x81\x47\x9d\xf3\xf3\x9f\xd4\x82\xa7\x44\x53\xc7\x04\x4c\x89\xdc\x55\x5a\xdd\xf2\x0c\x0d\x21\xa4\xb9\x0f\x78\xc2\x36\xf2\x8d\x00\x4f\x4b\xd9\xef\x04\x78\x3a\x16\xbf\x80\xfd\xfd\xf8\x77\xba\x98\xf1\xec\xeb\x9f\x5b\xfd\xa3\xb7\xfa\x6f\xa6\xb6\xb1\xd7\x94\xab\x23\x6c\x38\xed\x2a\x66\x8b\xed\xd2\xee\x9e\x82\xc6\x74\x53\x6d\x9a\x0d\x61\x2f\x85\x96\x8d\xf5\xa2\x6d\xe8\x65\xa9\xe4\x33\xd5\xed\xc7\x4d\xfb\x05\xbc\xfa\xb2\x89\x85\x0e\x06\x2e\x70\x9c\xaa\x9d\x2a\x69\x79\xc0\x36\xd0\x91\xb6\xf9\x7d\xda\x58\x8d\xac\x84\x0e\xd8\xd5\x94\xad\xc3\xd8\xc0\x49\xcb\x45\xa8\x0a\x45\x0d\x63\xbd\x09\x52\xa1\x0c\x9a\x5e\xb5\x6c\x06\x70\xc5\x4a\x04\x2e\x2b\x67\xcd\xb1\x72\xd6\xff\x05\x66\xa0\xb5\xa3\xdc\x57\x09\xf8\x89\x1b\x7b\x1e\xbf\xe9\x9e\x17\xa4\x8e\x82\x13\xb8\x5c\xac\xee\x75\x18\x0b\x6c\x8a\x5c\x5f\x5e\x07\x49\xaf\x7b\x10\xd6\x32\xe5\x4f\x08\x1b\x87\xb0\x3f\x46\xab\xef\x1b\xc0\x36\xd1\xa9\xcc\xd3\x7c\xb1\xbf\x1a\xff\xe6\x2a\x15\xca\xd8\x17\xa6\x52\xa5\xf4\x33\xa8\xc4\xa5\xc5\x05\xae\x45\x69\xae\x74\xc9\x6c\x1c\xf0\xf7\x37\x3b\xaa\x4c\x2f\x79\x2e\x2f\xf6\x9e\xa2\x74\x65\x0f\xc0\xe8\xfe\x8f\x97\x17\xa7\xbd\xb1\x00\x93\xcb\xf7\x57\xd7\x5d\xa0\xfc\x34\xed\xe3\x7b\xce\x9c\x20\x17\x90\x88\x67\xd9\x9b\x57\xe7\xc8\x2d\x58\x6f\x4e\x72\x43\xc5\xfb\xc4\x5a\xcd\xe7\xae\xf3\x09\xb2\xf2\x45\x28\x85\x4c\x9a\x9a\x4a\x03\x84\x9d\xfd\x3a\x2e\x4c\x67\x5b\x1b\x8b\xca\xf6\xe6\x70\xcb\x84\x7b\x68\x62\xdb\x71\xdd\xc0\xf9\xe7\x3f\x5a\x52\x07\x59\xc9\x7d\xc5\xde\x47\x98\xdb\x1c\x5a\xf4\x75\x1f\xfe\x18\xa8\x0f\x15\xda\xa7\x03\x3b\x69\x1e\xb7\xe0\x71\xd5\x87\x65\x74\x3e\x82\x1f\x61\x37\x81\xfc\x76\x26\xb7\xf8\xd9\x70\x08\x84\x5c\xde\x4c\x28\xab\xf8\x79\x99\x9f\xd2\x37\xe8\xc6\x0e\xec\x7e\x80\xbd\x04\xeb\xc7\xf4\x1d\x38\x0c\x7f\xf4\x92\x85\x92\xda\x76\x5a\x87\xc3\xb6\xdd\x14\xde\x3a\x42\x07\xea\xc3\x2f\x7f\x8d\x36\x8d\xa9\xfe\xec\x21\x3e\xfd\x92\xad\x6d\xb2\xed\x0f\xd8\x1c\x88\x24\xed\x30\x8f\x98\xcc\xc8\xbc\x21\x3e\x32\xce\x46\xda\x52\xa3\x84\x2d\x51\xea\xda\xcf\xda\xc2\xcf\xd7\xdd\xb7\x8c\xed\x85\xeb\x74\xa4\x4f\x44\x3a\x14\x64\xc5\x23\x86\x69\x47\xad\xc8\xd0\x36\xf5\x68\xe8\x59\xfb\xdc\xad\xc5\x30\xad\x59\x97\xa3\x4d\xb8\xc5\xb2\x3f\x7e\x6b\xb4\x1a\xfb\x10\x1d\xdf\xb4\x47\x83\x6a\x70\xb9\xea\xf7\x7e\xe0\xa2\x8b\xac\x7b\x4b\x0c\x36\x40\xd9\x36\x9e\x5c\xb2\x2f\xb7\x4f\xc3\xb8\xda\x42\xb9\x7c\x7a\xa1\x74\x00\xb3\x1d\xba\x5d\xd1\x94\xf5\x05\x6c\xc5\x7f\xaf\xac\x72\x58\x30\x4d\x55\x6a\xbf\x98\x60\x15\xb0\xaa\x12\xcb\xba\x39\x29\x36\xb7\x54\x4a\x89\xd9\x70\xae\xb4\xf1\xf4\x11\xa1\xc0\xb3\x9d\x90\x6e\x58\xda\x86\xb3\x43\xfd\x60\x87\x89\x2f\x2c\xb1\xe9\xc8\xef\xb9\x95\x09\x1f\x4b\x17\x4a\x89\x87\x15\x8a\x55\x90\x6d\x51\x66\x2c\x48\x3f\x56\x28\x43\x9d\xf3\xd0\x80\xdf\xc0\xb4\x64\x02\x34\x56\x1a\x0d\x4a\x4b\x95\x47\x2a\x89\xde\x71\x8d\x40\x3d\xc6\xa9\x12\x10\x72\x8c\xba\xf0\x92\x49\xbb\x4c\x9a\x4c\x66\x89\x3c\x93\xa0\x74\x16\x42\x5e\x07\xde\x55\x9f\x7f\x83\x33\xb1\xaf\x0d\xe2\x07\x26\x9c\x5c\x9c\x4d\xc3\xa9\xe9\xaa\x82\xaa\x32\x04\x53\x28\x27\x32\x58\xa0\x44\x4d\x67\x4c\x0c\x24\xde\x75\x0e\x85\xe8\x30\x88\x49\x38\x7b\x57\x9f\xae\xc6\x88\xa5\x23\xdd\x70\xb0\x74\xe4\x75\x2c\x95\x46\xc8\xd0\x32\x2e\x0c\xb0\xb9\x72\x16\x6e\x99\xe0\x59\x22\xe9\x23\xd1\xd4\xfd\x8b\xab\xf3\xe7\x29\x18\x0c\x9d\x8e\x99\x4a\x5d\x59\x7b\xe2\xd5\x0c\xa0\xf1\x58\xe8\xd5\x33\x28\xb3\x50\xe6\xa5\x9b\x89\xfc\x58\x1f\xf1\x6b\xe5\x16\xc5\xea\x14\x7c\xa8\x43\x70\x0a\x77\xd4\xae\xb8\x54\x0e\x58\x96\xad\x5a\xdb\xac\x82\xc3\x98\x08\x87\xde\xb4\x70\xc0\x86\x74\xe6\x1e\x4a\xae\x02\x8d\x81\x8c\x59\x16\x7a\xf4\xb8\x29\xea\x53\xb4\xc3\xc6\x0d\x4d\xef\xf0\x75\xad\x5e\x50\x2e\x91\xa1\xc1\x98\xb4\xb7\xd3\xe6\x04\x6d\x0a\xde\x37\x2b\xd5\xbb\xab\xb4\x32\xa5\xbd\x4e\xb3\x44\x26\xb2\x53\x28\x8f\x2d\x7d\xc0\x5a\xfd\x7e\xb5\x57\xe7\xce\x42\xe9\xad\x58\x5c\x5e\x9c\x86\xae\xe3\x50\x58\xb7\xec\x06\x81\x75\x24\x01\x33\xa1\x4e\x1e\x5a\xf6\x62\x85\x05\x18\xe4\x78\xe7\x65\xd6\x2d\x95\x73\x6c\x96\x7b\x06\x70\x5a\x60\x7a\x13\xfa\x2c\xe9\x25\xab\xde\x4e\x2e\x81\x42\xd1\x5f\xf7\x83\x62\x04\x80\x43\x62\x7e\x6b\xd8\xeb\x31\xeb\x51\x5b\xff\xb3\x01\x5e\xbf\x2a\xf4\x1c\xbc\x6a\x77\xac\xdc\x79\xfb\xda\x07\xc8\x06\xd7\xad\xe3\x2b\x4a\xfa\xb1\x42\x0b\x5c\x1a\x84\x6d\xb7\xe6\x46\x88\xad\x0f\xa7\x46\x10\x36\x55\xd2\x70\x63\xfd\x90\x23\xc3\x4a\x04\x8d\xc6\x89\xa0\x03\xb2\xb4\xa8\xe7\xe3\x2d\xea\x25\x58\x5e\x12\x48\x2e\x41\x30\x27\xd3\xe2\x55\x0f\x91\x13\xb9\x3d\x24\xc3\x18\x22\x27\x72\x04\x92\x13\x19\x7d\xf3\x20\xf2\x50\x4b\xf8\x08\xf0\xc4\x16\xdd\x97\x82\x39\x17\xdd\x1c\xde\x1b\xfd\x0f\x5c\xf5\xd9\x19\x4f\xfc\xbe\x79\x30\xff\x9f\xa4\xc2\xb4\x5f\xda\xdf\xac\x5f\xe8\x3d\xa9\x9b\x5c\xe6\x4b\x60\x60\x90\x7e\x1e\xd1\xf9\x30\x30\xa1\x37\x2c\x65\xb2\x13\x56\x89\x5c\xf5\xba\xac\x68\x81\x0f\xd8\x32\x52\x0f\x45\xa9\x41\x02\x68\x0f\x6e\xbf\xf7\x51\x5b\x70\x23\xe8\x05\x64\xc3\x65\x3f\x00\xf6\x96\x09\xd1\xe9\xfb\xe1\xfe\x23\x99\xf0\x50\xac\xbd\xf3\x04\xd0\x58\xed\x52\xfa\x11\x8a\x55\x50\x28\x91\x01\x1b\xfa\x1d\x4e\x24\xe0\x23\x7e\xbe\xea\x25\xc6\xd6\x6e\xa6\x66\xcb\x27\xff\x4c\x47\xc1\x2a\x83\xa3\x1c\x2a\x44\xf7\x98\xdc\x4c\xb9\xb9\x18\xaa\x90\xad\x8c\x59\xfd\xac\xf2\xdf\x4a\x64\xa1\x2b\x93\x1b\xcb\x53\xd3\x75\x52\xfd\x83\xcb\x13\xb9\x7c\x8c\x8b\xfc\xf8\x5f\x9d\xde\x89\x9b\x3d\xe2\x30\x6b\xbe\xb4\x0f\x96\x09\xb5\x93\x7e\x33\xbf\xa2\x46\x8d\xf7\x3d\xf2\xb3\xb5\xa1\x0b\x5d\xa5\xbf\x7a\x8a\xf1\xf4\x95\xd2\xc2\xda\xea\x99\x44\xc7\x54\xd8\xa9\x9a\xec\xb5\xda\x9d\x6d\x46\x08\x7c\x2e\xa8\x69\x87\xef\xfd\x38\x73\x00\x9d\x93\xdd\xa1\xce\x9d\xe6\x88\xf7\xe9\x8f\x1c\x02\xd9\xdc\xb8\x10\x7e\x6f\x03\xd1\x83\x6c\x7e\x58\xec\x40\x22\x6c\x82\x1f\x61\x7c\xcd\x96\x55\x0e\xf7\xa9\x47\x0e\x3e\xf8\x7a\xf0\xff\x00\x00\x00\xff\xff\x56\x03\xc9\x7f\xfc\x3d\x00\x00")

func backendSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_backendSwaggerJson,
		"backend.swagger.json",
	)
}

func backendSwaggerJson() (*asset, error) {
	bytes, err := backendSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "backend.swagger.json", size: 15868, mode: os.FileMode(0644), modTime: time.Unix(1556139164, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf8, 0x7c, 0xff, 0x2b, 0xa4, 0xe0, 0xae, 0x2c, 0xc3, 0x61, 0xc9, 0x60, 0x17, 0xb8, 0xf, 0x87, 0xe6, 0xa3, 0xfd, 0xe8, 0x8f, 0xc1, 0x3c, 0x15, 0xf6, 0x5e, 0x3d, 0x1e, 0xe9, 0xe1, 0xf7, 0x8e}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"backend.swagger.json": backendSwaggerJson,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"backend.swagger.json": &bintree{backendSwaggerJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
