// Code generated by go-bindata. DO NOT EDIT.
// sources:
// api/protobuf-spec/mmlogic.swagger.json (12.457kB)

package mmlogic

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

var _mmlogicSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x6f\xdb\x48\xf2\xbf\xfb\x53\x14\xf4\xff\x03\x4e\x16\xb2\x3c\x3b\xbb\xd8\x83\x6f\x1e\xc7\x9e\xd1\x22\x8e\x0d\xd9\x3e\x2c\x86\x41\xd0\x22\x8b\x62\x8f\xc9\x6e\xa6\x1f\x72\x88\xc0\xdf\x7d\x51\xdd\xcd\xa7\x29\x5b\x7e\x4e\x06\x58\x9d\x44\xb2\xbb\x58\x55\x5d\xf5\xab\x17\xbf\xef\x00\x4c\xf4\x0d\x5b\xad\x50\x4d\x0e\x60\xf2\xf3\xec\xa7\xc9\x94\xee\x71\x91\xca\xc9\x01\xd0\x73\x80\x89\xe1\x26\x47\x7a\xce\x4a\xbe\x5f\x2a\x69\xe4\xd2\xa6\x7b\xba\xc4\x78\xbf\x28\x72\xb9\xe2\xf1\xcc\xdd\x75\x7b\x01\x26\x6b\x54\x9a\x4b\x41\x3b\xc2\x5f\x10\xd2\x80\x46\x33\xd9\x01\xb8\x75\x6f\xd0\x71\x86\x05\xea\xc9\x01\xfc\xee\x37\x65\xc6\x94\x35\x01\xfa\xaf\x69\xed\x67\xb7\x36\x96\x42\xdb\xde\x62\x56\x96\x39\x8f\x99\xe1\x52\xec\xff\xa1\xa5\x68\xd7\x96\x4a\x26\x36\xde\x72\x2d\x33\x99\x6e\xc5\xdc\x5f\xff\x7d\xdf\x49\xb3\x5f\x30\x13\x67\x7b\xa5\x92\x29\xcf\x51\xef\x7f\x77\xd7\x33\x9e\xdc\x36\x8b\x01\x26\x2b\x34\x9d\x4b\x12\xc9\x16\x05\x53\x15\x89\x7d\x81\x22\x81\x5f\xd1\x9c\x7b\x12\xc0\xc0\x91\x00\xb9\xfc\x03\x63\x03\x37\xdc\x64\x60\x32\x84\xf9\x07\x48\x39\xe6\x09\x94\xb2\xb4\x39\x33\x98\x4c\x81\xd3\xf3\x3c\x07\x85\xc6\x2a\x01\x2c\x12\xb0\x9b\xf2\x3c\xc7\x64\x17\xa4\xc0\x59\x24\xe0\x93\x34\x78\x00\x29\xcf\x0d\x2a\x0d\x4c\x21\x30\x4d\x1a\x4a\xc0\x48\xc8\xd8\x1a\x61\x89\x28\x20\xce\x30\xbe\xc6\x04\x52\xa9\x60\xcd\x72\x9e\x70\x53\xc1\xb2\xa2\x37\x47\x02\x96\x2c\xbe\x46\x91\xb0\x92\x03\xdc\x64\x28\x80\xc5\x31\x96\x86\x8b\x15\x30\x08\xb2\x87\x03\x71\xe2\xc9\x12\x95\xd3\xe2\x3c\x21\x11\x5b\xe9\xba\x8b\x14\xea\x52\x0a\x8d\xba\xa7\x1a\x80\xc9\xcf\x3f\xfd\x34\xb8\x05\x30\x49\x50\xc7\x8a\x97\x26\x18\xcb\x21\x68\x1b\xc7\xa8\x75\x6a\x49\x7c\x4f\x69\xd6\x21\xef\xd5\x4c\x96\xc3\xee\x10\x03\x98\xfc\xbf\xc2\x94\xe8\xfc\xdf\x7e\x82\x29\x17\x9c\xe8\xea\x7d\x56\xf2\x96\xd7\x45\xa0\x3a\xe9\xed\xbd\xdd\x19\xfb\x7f\xdb\x91\xab\x64\x8a\x15\x48\xea\x6e\x0c\xcb\xff\x06\x12\x09\x56\x38\x4f\xa9\x2d\x66\xc8\x3c\x77\x82\x92\xe1\x0d\x9f\x28\xfc\x6a\xb9\x42\xd2\xad\x51\x16\x07\x4f\x4d\x55\x3a\xba\xda\x28\x2e\x56\x5d\xee\x6f\xa7\xdb\x72\x53\x2a\x3a\x42\xc3\x51\x8f\x73\xf5\xd5\xa2\xaa\xee\x61\x2b\x65\xb9\x7e\x0d\xbe\x50\x29\xa9\x7e\x2c\x96\xb4\x61\xc6\xbe\x8d\x9a\x9a\xff\x9f\x3b\xe6\x66\xd8\x6a\x68\x68\x93\xd3\xe2\x23\x81\x53\xbb\xf9\xf3\x4e\x97\x44\x10\x6e\x0c\xc5\x4a\xa9\x59\xde\xf5\xc8\x49\x69\x37\x63\xd7\x91\x42\x66\xf0\x3c\x6c\x03\xae\x21\x66\x84\x3e\x84\x1c\xa7\xa7\x27\x1a\x4c\xc6\x08\xa2\x74\x46\x68\x73\xa3\xb8\x41\x42\x14\xae\xc8\x65\x6d\x6e\x34\x18\x19\x09\x07\x21\xa5\xd4\x98\xc0\x29\xf1\x71\xe6\x90\x6f\xea\x37\xc7\x4c\xc0\x12\x41\xa3\x30\x20\xad\x71\x50\xf8\x8b\x87\x23\x38\x3c\x9f\x83\x14\x31\x12\x0e\x66\x4c\x47\xc2\x41\x19\x2b\x4b\x25\xd7\x98\xc0\xbb\x65\x05\x09\xa6\xcc\xe6\x66\x1a\xb0\x0c\x70\xcd\x72\xcb\x8c\x54\xf4\x4e\xc2\x8f\xf7\x04\x92\x7b\xc0\x92\x44\x03\xcb\x73\x28\x73\x56\x11\x52\x72\xe1\x2e\x17\x52\x3b\xe0\x34\xd2\x6d\x6f\x18\xf5\xcb\x80\xaf\x84\x54\x08\x39\xd7\xc6\x91\x71\x22\xea\xfe\x52\x0f\xe7\x2d\x81\x35\x4f\x30\x81\x6b\xac\xda\x17\x3b\x49\xaf\xb1\xaa\x57\x11\x3d\x90\x29\x34\x07\x42\x0f\x96\x08\x14\xdd\x78\x82\x0a\x93\x48\xcc\x3f\x9d\x5f\x5d\x1e\x40\x24\xe0\x6f\x70\x79\x06\x8b\xe3\xcb\xab\xc5\x27\x38\x84\xd3\xc3\xcb\xa3\xdf\xce\x7e\xf9\xf7\xf1\xd1\x25\x1c\x9e\x5c\x1e\x2f\xe0\x10\x2e\xae\x8e\x8e\x8e\x2f\x2e\x4e\xae\x3e\xd2\xb1\xc0\xe2\xea\x53\x24\x00\xe0\x52\x42\xec\x4e\xb0\x09\x3a\x1d\xfd\x43\x81\x5a\xb3\x15\x36\x11\x48\xa3\x0f\x40\xba\x8d\x40\x07\x8e\x0c\xc0\x1e\xf0\x64\x4a\x41\xbb\x16\x80\xb4\x8c\x24\x01\x5d\x9c\x9e\x9e\x7c\x39\x5f\x9c\x9d\x9f\x5d\x1c\x7e\xfc\x32\xff\x00\x28\xd6\xb0\x66\xaa\xd9\xdb\x82\x4d\x73\xcb\xf9\xf9\x0c\xe0\x3f\xd2\x42\x61\xb5\x01\xfc\x46\x91\x99\x9b\xbc\xf2\xaf\xc9\xb8\xd3\x09\x13\x80\x45\x69\x2a\xf0\xde\x02\x3c\x85\x4a\x5a\x45\xaf\x6c\x68\x29\x77\x84\xd3\x36\x92\xfa\xb3\x9b\x7f\xd0\xe0\x63\x25\x9d\x35\xdd\xdf\x0d\x67\xbf\x0b\x0a\x4b\x52\x4b\xe2\x05\x9e\x41\x43\xeb\x77\xe9\x22\x10\xcb\x3f\x43\x29\x65\xae\x7b\x42\x4b\x6b\x4a\x6b\x20\x55\xb2\xf0\xf4\x28\x96\x38\x92\xe7\xb4\x76\xd7\x39\xc7\x34\xd0\x02\x1f\xb9\x6b\x55\x7a\xc6\x68\x99\xe7\x93\x50\x45\x03\x5b\x92\xc9\x67\xf2\x06\x0a\x26\xaa\xc6\x36\x69\x6d\x88\xe8\x0d\x35\x77\x7e\x98\x00\x13\x89\xdb\x90\x4b\xb1\xea\x2e\x04\x23\xe5\x35\x71\xaa\xac\x98\xc2\x4d\xc6\xe3\xcc\x73\xd0\x71\xae\x86\x18\xed\x0b\x31\x1f\x28\xe8\x33\x47\xcd\x31\xe6\xd4\xeb\x8d\x25\xf8\xf0\xac\x95\x88\xe9\x4e\x58\x9e\x0d\x4d\xf3\x13\x1c\x2f\x16\x67\x0b\xa8\x6d\x4f\x61\x29\x95\x01\x06\x29\xe3\xb9\x55\x08\x52\xf9\x83\x27\xad\xd2\x9b\x1f\x30\x47\x47\xe7\xb9\x86\xe9\x58\xea\x58\x65\xe7\xb0\x1b\x5e\xdc\x6e\x56\xdb\x98\x27\xe1\xf3\x91\x25\x0f\x5a\x76\x6b\x1b\xdb\x03\x14\xb1\xb4\xc2\x90\xa7\xce\x46\x8d\xa7\x31\xf9\x29\x30\x51\x99\x8c\xe8\x54\xd2\x02\x19\x50\x86\x0a\x09\x4b\x7d\x56\xe7\x33\x35\x9f\x87\xf9\x5f\x73\x32\x83\x53\xf1\x4e\x33\xf6\x3a\xef\x02\x6f\xf3\xae\xff\xf9\x45\xeb\x17\x67\x57\x97\x0e\xa4\x19\x2c\xdc\xbd\xbe\x11\x33\x58\x4a\x99\x23\x13\xb5\xd3\x04\xd3\x22\x56\x09\xd8\x9c\x49\x79\xa3\x8b\x04\x4f\xdb\x7b\xe4\x67\x1d\x0b\xbb\x27\xf3\xee\x07\xe8\x1f\x3f\xfb\xee\xf3\xfb\x56\x19\xf8\x52\x26\x77\xb2\x34\x9f\xc0\x8d\x3d\xb9\x3f\xfb\x7e\x01\xa1\xbf\x5a\xd4\x66\x1b\x99\x5f\x3c\x0d\xf4\x8e\xb4\xe7\xfc\x6d\xff\xbb\xbf\xfa\x42\x57\x33\x52\xd5\xd6\x25\xad\x77\x6e\x97\xc5\xb8\x12\x51\x24\xc1\xdb\xe8\x2a\xb5\x22\x76\xa2\x77\xad\x71\x60\x5f\x0b\x34\x8a\xe3\x1a\x5b\x94\x80\x15\x1a\xdd\xcf\x8d\x1a\xa7\x67\x26\x38\x1e\xae\x51\x55\x70\xe2\x5e\x15\x82\x7a\x24\x5a\x1a\x53\x98\xe1\xb7\x38\xb7\x09\xb1\xd1\xcd\xf3\x44\x45\x99\x55\xca\x57\x56\x51\x36\xd0\xe6\x74\x7a\x06\x30\x37\x91\x88\x65\xb1\xe4\x22\x24\x76\xc1\xbd\xa7\x4e\x2e\x8f\x9b\xdd\x07\x2d\x71\x87\x5b\xb3\x07\xea\xe2\x86\xbb\x57\x76\xce\x77\xda\x28\x64\x05\x71\xd7\x50\x7f\xff\x14\x8f\xfd\xb6\xe7\x29\xed\x8d\x14\xce\x8d\x30\x6f\xe5\xb9\x43\x13\xfd\x31\x6a\xe8\x2e\x57\x3e\xf4\xde\xc3\xdc\xdb\xd7\xad\x5d\xf6\x5c\x3c\x9d\xb9\x50\xf2\x6a\xec\x0d\x9e\xa6\x52\x15\x8c\xc0\x63\xc2\x85\xf9\xd7\x3f\x9f\xcf\x3d\xe6\xac\xd4\xb8\xa1\x7f\xf2\x1c\xfe\x85\x2d\x96\x78\xa7\xdd\xd0\xf2\x9f\x48\xbb\xcc\xf1\x75\x4a\xf4\xa6\xfb\xd9\x71\xb3\xb6\xff\xe8\x3d\xed\xd0\x18\xc5\x97\xd6\x60\x17\x98\x6b\xe6\x7d\xff\xb0\x61\x7e\xd2\x69\xe9\xf4\x70\x3b\xe8\xb5\x07\x33\x9b\xec\xab\xeb\xaa\x2e\x57\x79\x60\x63\x57\x73\x9b\xcf\xfd\x76\x34\x2a\x6d\x0c\x8c\xcf\x90\xd5\x85\x89\x21\xcf\xe3\x71\x39\x64\x6a\xba\x53\x82\x3c\x81\xe3\x80\x82\x9b\x59\x1e\x12\xf8\x15\xcd\x61\x9e\xcf\x5d\x08\x4a\xfc\x29\xeb\x2d\xc8\x3c\x28\xb9\x87\xa1\xc7\x89\xee\x7b\x1f\x5b\x48\x3d\x0e\xfd\xcf\xe0\xb6\xe3\xe3\x8f\x63\xb9\x13\x4e\xb7\x63\x7b\xd0\xea\xfd\x91\x6d\xeb\x23\xd7\xe6\xaf\x64\x18\xf5\x6a\x9f\x8f\xfd\x19\x10\xc5\x46\xf0\x71\xeb\xcd\x05\xfb\xb6\x7e\x19\x78\xeb\x12\xe5\xe2\xe5\x89\xba\x28\xf8\xb8\x13\xbc\x70\x5b\xee\x1e\x60\x73\x14\x77\x72\xca\xdd\x8c\xa9\x64\x37\xe4\xf1\xae\x1f\x52\x96\x79\xd3\xb0\xec\x25\xbc\xa3\x56\x30\xcf\xe7\xa2\xdf\x4a\xde\x68\x06\x83\x77\x87\x9d\xc0\x75\xdb\xee\xab\xeb\x68\x85\x1a\xd5\x3a\xcc\xad\x52\x6b\xac\x42\xb0\x54\x8a\x8e\xf2\xd0\xf5\xb7\x67\x98\x23\x4f\x9e\x62\x4f\x1b\xa8\x6d\xb9\xdb\x0f\x3e\x9e\xb0\x31\xb4\x7e\x36\x6d\x65\x4a\xb1\x7e\x7a\x34\xe1\x06\x8b\xe1\xfa\x47\x43\xc2\xc6\x04\x9f\x4a\xcb\xd7\x66\x66\x2c\x0a\x6c\x62\x28\x8c\x6f\xb6\xd5\xec\x43\x8e\x72\x56\xa2\xf0\x6d\xcb\x5d\xaa\x2d\x0d\x2a\xc1\xa8\x0a\x2b\xc9\x52\x85\x71\xc5\x9f\x2b\x1c\x6f\xb8\x72\xd3\x00\x23\x63\x99\x83\xf7\x73\x67\xc5\xd1\xa4\x63\xa6\x3a\x9a\xcc\x22\x31\x17\x20\x55\xe2\xdd\x4e\xf9\x2c\xa8\xe9\xd8\x5b\x5d\xb7\x20\x3b\x53\x91\xa9\x6f\x47\xd5\xad\xaa\x58\x26\x08\x3a\x93\x36\x4f\x60\x85\x82\x6a\x50\x8c\x04\x03\x81\x37\xbd\x16\xab\xef\x4a\x09\x98\x7f\x70\x2c\xb6\x16\xdb\x69\x94\xbf\x23\x1e\x0b\x2a\x90\x13\x34\x8c\xe7\x75\x33\xce\x0d\x8d\x23\xe1\xb2\x42\xed\x04\xe9\x76\x66\xa7\xa0\xd1\xf7\xf2\x12\x19\xdb\xa2\xd6\xc4\xfb\x19\x40\xab\x31\x7a\x2e\x5c\xeb\xd7\xd7\xd3\xee\x66\x24\x02\x73\x26\x53\xd2\xae\xdc\x20\xa5\xed\xb5\x15\xec\xba\xdb\x53\x98\xc2\x8d\x6b\x68\x56\xd2\x02\x4b\x92\xb6\x47\x20\x61\x37\x38\xc2\x2e\x89\x16\x09\x6d\x48\x04\xaa\xfc\x7d\xcd\x9b\xa3\xd6\x90\x30\xc3\xdc\x5e\x37\xb1\x6a\xa6\x02\x8d\x1a\x76\xeb\x79\x00\x5c\xd6\xec\x79\xe6\x22\xc1\x75\xc3\xbd\x99\x76\x47\xf4\x7e\x66\x15\x58\x1f\xcc\xae\x6a\x51\xba\xe7\x34\x8b\x44\x24\xba\x16\x00\xb1\x14\x86\x71\x01\x0c\x7c\x3d\x02\x32\x6d\xb4\xba\xb4\xc6\xb7\x40\x57\x8b\xf3\x23\xd7\x47\x0d\x0d\x11\xc3\xae\x11\x58\x8f\x12\x30\xb2\x48\x42\x52\x29\xf2\x0a\x42\x0d\x04\x0c\x52\xbc\x09\x1d\xf1\x22\xcc\x99\x9a\xe3\x9e\x01\x1c\x65\x18\x5f\xfb\x56\x8a\x7b\x49\xad\x6b\xd2\x8f\x33\x45\xfa\x3f\x34\x8a\x0d\x00\xec\x1d\xf3\xaf\x86\xbd\x63\x69\xe8\xa3\xd2\x8f\x57\x03\xbc\x61\x19\xf8\x20\xda\x31\xad\xf9\x4a\x90\xff\x3d\x45\xa0\x47\x62\xe5\x93\xc3\xd7\x5b\x80\x6c\x48\xa1\xef\xe0\x2b\x12\xb9\x1e\xb8\xb4\x08\x7b\xa2\xa4\x30\x7d\x88\x8d\x73\x8e\xc2\x6c\x44\x58\x37\xb0\xd5\x86\x96\xbc\xd3\xac\xa8\x7b\x84\x8e\x07\x64\x71\x56\xef\xf7\x9d\x4b\xc3\x0b\x07\x92\x15\xe4\xcc\x8a\x38\x7b\x3f\x40\xe4\x48\x3c\x1e\x92\x61\x13\x22\x47\x62\x03\x24\xd7\x2d\xd3\x87\x91\x87\x36\x6f\x02\x9e\xd0\x03\xfe\x51\x30\xe7\xbc\xef\xc3\x6f\x56\x82\x84\x29\xd3\x6b\x67\x3c\xa1\xc6\x7a\xd0\xff\x5f\xa4\xda\x7b\xdb\xd2\xa3\x33\x29\x74\x9f\xaf\x39\x9a\xfe\x13\x13\xe6\xa6\x8c\x32\xed\x17\x27\xa1\x33\x1f\x3e\x1c\x69\xcc\x2a\x12\x83\x0f\x00\x7a\xa3\x03\x93\x49\xdd\xcc\x04\x5d\x0c\xee\xbe\xf7\x59\x21\xb8\x33\xc4\xf8\xf3\xbd\x61\x31\x34\x80\x37\xf3\x84\xa0\xf4\xb7\xc9\xfd\x37\x78\xc2\x43\xb6\xf6\x81\x12\x40\x6d\x94\x8d\x5d\x25\x69\x24\x64\x32\x4f\x80\xdd\x19\x3a\x39\x5b\xf0\xdf\x7f\x8d\xeb\xf9\x62\xe0\x18\x8f\x56\xb3\x6f\xc9\xbf\x74\xab\xa0\x6e\x95\x6f\xa0\x3b\xd2\xf0\xbe\xa7\xdd\x7d\x57\x9d\xcd\xd7\xbf\xbf\xc9\x3c\xd1\x6e\x58\xcf\xb5\xe1\xb1\xee\x2b\xa9\xfe\x2e\xf8\x50\x54\xcf\x51\x11\xad\xff\x62\xd5\x93\x72\xb3\x67\x74\xaf\x97\x95\x19\x53\x42\x4f\x42\x65\x05\x05\xf3\x0b\x37\x29\x3b\x1e\x24\x3f\x8f\x16\x74\xa5\xca\xf8\x0b\xa5\x18\x9b\x18\xa6\x1c\x68\xb5\xf9\xe0\xb8\x30\xff\xf8\x79\x5c\x0f\x99\x31\xe5\x2b\x91\x0e\xae\xf0\x94\xd3\x71\x5c\x3d\x3d\xdb\x0c\x10\xf8\x5a\x50\xd3\x35\xdf\xfb\x71\x66\x07\x7a\xa3\x9c\xb1\xd1\x69\x3b\xd3\x79\x8d\x46\xba\x4f\x37\xb7\x0b\xd0\x0f\xcf\x70\x1f\xcc\xe8\xc7\x09\x8f\x38\xc3\x36\x18\xe2\xd7\xd7\x19\xb3\x4c\xe1\x7e\x06\x9d\x9a\x77\x6e\x77\xfe\x1b\x00\x00\xff\xff\xa0\x1b\x67\x45\xa9\x30\x00\x00")

func mmlogicSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_mmlogicSwaggerJson,
		"mmlogic.swagger.json",
	)
}

func mmlogicSwaggerJson() (*asset, error) {
	bytes, err := mmlogicSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mmlogic.swagger.json", size: 12457, mode: os.FileMode(0644), modTime: time.Unix(1556139164, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4, 0xa4, 0x7c, 0x5c, 0x59, 0xa7, 0x2, 0xc6, 0x5, 0x49, 0xfd, 0xc2, 0x89, 0x8c, 0x79, 0x25, 0x1b, 0xb0, 0xb3, 0xa0, 0x2, 0x40, 0x50, 0xb2, 0x3e, 0x2, 0xce, 0x46, 0x93, 0x88, 0x7d, 0x6a}}
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
	"mmlogic.swagger.json": mmlogicSwaggerJson,
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
	"mmlogic.swagger.json": &bintree{mmlogicSwaggerJson, map[string]*bintree{}},
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
