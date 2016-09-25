// Code generated by go-bindata.
// sources:
// template/default.tmpl
// DO NOT EDIT!

package deftmpl

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

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3b\x79\x6f\xdb\xc6\xf2\xff\xf3\x53\x4c\x59\xfc\xd0\x18\x10\x45\xdb\x39\x50\x1f\xf2\x0f\x8a\x4c\xc7\xc2\x93\x25\x43\x92\x93\x06\x45\x11\x50\xe4\x4a\xda\x84\xe4\xb2\xdc\xa5\x65\xb7\xaf\xdf\xfd\xcd\x2c\xa9\x83\x12\x65\xcb\x46\x6a\xeb\xbd\x3a\xe9\xa1\x1d\xee\xdc\xb3\x33\xb3\xdc\xe5\x9f\x7f\x82\xcf\x86\x3c\x62\x60\x7e\xf9\xe2\x06\x2c\x51\xa1\x1b\xb9\x23\x96\x98\xf0\xd7\x5f\x75\x1a\x5f\x64\x63\x9c\xc8\x22\x1f\x81\xc6\x5a\x94\xab\x6e\x8b\xb0\xf0\x79\xd5\xb9\x51\x2c\x89\xdc\x00\x41\x08\xb1\x7f\xb4\xf5\x3c\xf9\xff\x09\xf3\x18\xbf\x66\x49\x8d\x26\x75\xf3\x41\x86\x93\x53\x2f\x92\x97\xe9\xe0\x2b\xf3\x14\x91\xfd\x95\x50\x7a\xca\x55\xa9\x84\x7f\x83\x12\x57\x71\x3c\x45\xe5\x43\x60\xbf\xcf\x1e\x9a\x43\x9e\xf0\x68\x44\x38\x87\x84\xa3\xb5\x90\xd5\x33\x0d\x45\xd4\x80\x45\x8b\x1c\x7f\x03\x9a\xf4\x21\x11\x69\xdc\x72\x07\x2c\x90\xd5\x9e\x48\x14\xf3\x2f\x5d\x9e\xc8\xea\x47\x37\x48\x19\x31\xfc\x2a\x78\x04\x26\x10\x55\xc8\x58\x8e\x14\xbc\x22\x5a\xd5\x86\x08\x43\x11\x65\xc8\x3b\x39\x6c\x81\xde\x0e\xa2\xbc\x42\x94\x09\x57\xe3\xe2\x64\xb4\x40\x28\xae\x59\x91\x7b\xdb\x0d\x91\x61\x66\xc6\x32\xee\x33\xc1\x77\x66\xbf\xd6\xf8\xc6\x67\xd2\x4b\x78\xac\xb8\x88\xcc\x3b\x6c\xac\xd8\x8d\xca\xfc\xf8\x25\xe0\x52\xe5\x53\x13\x37\x1a\xa1\x64\x38\xc8\xe4\x3a\x34\xe6\xc0\x55\x3b\x91\x55\x2c\x6d\x48\x12\x9f\x46\x35\x98\x29\x90\x0b\x96\x31\xaf\x47\x91\x40\x3f\xa1\x4c\x05\x92\x0b\xe0\xc7\xd1\xed\x89\x34\xf1\xd8\x61\xe6\x4c\x16\xb1\xc4\x55\x22\xc9\xc2\xcf\x28\x31\x54\xc1\x06\x32\x70\xbd\x6f\x55\x1c\xb9\x69\xa0\xaa\x8a\xab\x80\xe5\x56\x50\x2c\x8c\x03\x57\x15\x63\xb1\xba\xce\xe4\x45\x3a\xa9\xa4\x25\x10\x96\x91\x2a\x2e\xb4\x0d\xe9\x0d\xdd\x20\x18\x20\x60\x85\x5e\xa9\xf8\x44\x14\x03\xe7\xbe\x89\x01\x8f\xbe\x6d\x2c\x41\x9c\x30\x0a\x16\x73\xb3\xd9\x0b\xf4\xef\x34\x80\x4e\x1b\x1b\x4a\xc0\x3d\x11\xe1\x9a\xf9\xca\x37\x94\x81\xe6\xa7\x49\xb0\xa9\xc4\x2b\xca\x15\xc2\x64\xcc\x63\x6f\xec\xaa\xb9\x43\x12\x11\x3e\xde\xb9\xcb\xd4\x70\xd5\x4b\x44\xd9\x3c\xf0\x0a\xb2\xc5\xc4\xcd\x4f\xd5\xed\x8c\xde\xea\xea\x7f\x58\x30\xaf\x52\xf4\x02\xce\x22\xf5\x78\x8d\xd7\x51\x9c\xd7\x8d\xc7\x85\xc8\x2a\x5d\x1e\x49\xe5\x46\x1e\x93\x25\x74\x57\xd2\xdd\x1d\x56\x15\xb1\x1c\xb1\x88\xb3\xc7\x3b\xe9\x2e\x62\xab\x1e\xca\xab\xc3\x9a\x64\x58\x5a\x0e\x8c\xa5\x62\x54\xa8\x76\x3b\xb0\x0b\x16\xce\xc9\x80\x90\x01\x75\xda\xbd\xdb\x22\xc5\x92\xa9\x99\x58\x0b\x1a\x95\xf0\xeb\x32\x29\x82\x6b\xe6\x2f\x71\x9c\x82\x37\xe7\x39\xc5\x58\xe1\x6a\x6d\x62\x52\xa9\xab\xc0\xc3\xa3\xa9\xe0\xf5\x6b\xee\x61\xed\x40\xda\x0f\x75\xfb\x72\xbe\x7d\x48\x10\xaf\x32\x7d\x44\x7a\x29\xa8\xc1\x13\xef\x3b\x25\x17\x16\xba\x3c\x98\xdb\x78\xde\x93\x3d\x78\x0d\x14\x29\x8d\x55\xa8\x53\xb3\x71\xfc\xc3\x69\xa7\xd1\xff\x7c\xe9\x00\x81\xe0\xf2\xea\x7d\xab\xd9\x00\xd3\xb2\xed\x4f\xaf\x1b\xb6\x7d\xda\x3f\x85\x5f\xce\xfb\x17\x2d\xd8\xab\xee\x42\x1f\x5b\x06\xc9\x69\x75\xb8\x81\x6d\x3b\x6d\x5c\x07\x63\xa5\xe2\x43\xdb\x9e\x4c\x26\xd5\xc9\xeb\xaa\x48\x46\x76\xbf\x6b\xdf\x10\xad\x3d\x42\xce\x7f\x5a\x6a\x01\xb3\xea\x2b\xdf\x3c\x41\xce\x96\x65\xf4\xd4\x6d\xc0\xc0\x45\x69\x35\x13\x9f\x25\x9c\x22\x90\x1c\x00\x44\x5a\x22\xed\x11\x76\x6f\xe9\xa0\xea\x89\xd0\x26\x1d\x46\x69\x64\x6b\x72\xae\x97\xd1\xb3\xb4\x6a\xd6\xd4\x1c\x12\x2d\xd8\x1f\x33\xb8\x68\xf6\xa1\xc5\x3d\x16\x49\x06\xaf\x70\xb0\x63\x18\x0d\x11\xdf\x26\x7c\x34\xc6\x15\xe4\xed\xc0\xfe\xee\xde\x1b\xb8\xc8\x28\x1a\xc6\x25\x4b\x42\x2e\x25\x52\x04\x2e\x61\xcc\x12\x36\xb8\x85\x11\xf2\xc1\x1c\x50\x41\x81\x18\x03\x31\x04\xac\x1b\xc9\x88\x55\xb0\x09\x46\xa1\x6f\x01\xfb\x60\x89\x08\x62\xa0\x5c\x1e\xd1\x82\x75\xc1\x43\x1e\x06\xce\x54\x63\x24\x23\xc5\x50\x4d\xdc\x24\xd3\xd0\x95\x52\x78\x1c\x25\xf4\xc1\x17\x5e\x1a\x62\xfe\xd5\x99\x06\x86\x3c\xc0\xdc\xf2\x4a\xa1\xd0\x66\x2f\xc7\x30\x77\x34\x13\x9f\xb9\x81\x81\x19\x87\x9e\x4d\x1f\xe9\x76\x56\xa4\x0a\x12\x26\x55\xc2\xb5\x15\x2a\xc0\x23\x2f\x48\x7d\x92\x61\xfa\x38\xe0\x21\xcf\x39\x10\xba\x56\x5c\x1a\x48\x14\xdb\xa3\x8a\x96\xb3\x02\xa1\xf0\xf9\x90\xfe\xcf\xb4\x5a\x71\x3a\xc0\x9c\x30\xae\x80\xcf\x89\xf4\x20\x55\x08\x94\x04\xd4\x76\xac\x90\x1e\xb6\x48\x40\xb2\x20\x30\x90\x02\x47\xb9\xb5\xae\x73\xe9\xf4\x1c\x12\x3d\x26\x83\xaa\xdc\x44\x92\x20\x93\x31\x7a\xb5\xa0\x09\x97\xc6\x30\x4d\x22\x64\xc9\x34\x8e\x2f\xd0\x64\x9a\x23\x45\x33\x41\x68\xfa\x50\x04\x81\x98\x90\x6a\xd8\x53\xf8\x3c\xef\x60\xb5\x93\xdd\x01\x75\xf1\xde\xcc\xaf\x98\xbd\x51\xd4\x4c\x04\x72\x40\x3c\xf7\x6a\xfe\x48\x8e\xb1\x99\x83\x01\xcb\x0d\x86\x7c\xd1\xbc\xee\x82\x3a\x09\xb1\xa7\x22\xa6\xb8\x1b\x40\x8c\x45\x80\xf8\x2d\xab\x59\x45\xfe\xe7\x0e\xf4\x3a\x67\xfd\x4f\xf5\xae\x03\xcd\x1e\x5c\x76\x3b\x1f\x9b\xa7\xce\x29\x98\xf5\x1e\x8e\xcd\x0a\x7c\x6a\xf6\xcf\x3b\x57\x7d\xc0\x19\xdd\x7a\xbb\xff\x19\x3a\x67\x50\x6f\x7f\x86\x7f\x35\xdb\xa7\x15\x70\x7e\xb9\xec\x3a\xbd\x1e\x74\xba\x46\xf3\xe2\xb2\xd5\x74\x10\xd6\x6c\x37\x5a\x57\xa7\xcd\xf6\x07\x78\x8f\x78\xed\x0e\x86\x70\x13\x63\x17\x89\xf6\x3b\x40\x0c\x73\x52\x4d\xa7\x47\xc4\x2e\x9c\x6e\xe3\x1c\x87\xf5\xf7\xcd\x56\xb3\xff\xb9\x62\x9c\x35\xfb\x6d\xa2\x79\xd6\xe9\x42\x1d\x2e\xeb\xdd\x7e\xb3\x71\xd5\xaa\x77\x71\x61\x77\x2f\x3b\x3d\x07\xd9\x9f\x22\xd9\x76\xb3\x7d\xd6\x45\x2e\xce\x85\xd3\xee\x57\x91\x2b\xc2\xc0\xf9\x88\x03\xe8\x9d\xd7\x5b\x2d\x62\x65\xd4\xaf\x50\xfa\x2e\xc9\x07\x8d\xce\xe5\xe7\x6e\xf3\xc3\x79\x1f\xce\x3b\xad\x53\x07\x81\xef\x1d\x94\xac\xfe\xbe\xe5\x64\xac\x50\xa9\x46\xab\xde\xbc\xa8\xc0\x69\xfd\xa2\xfe\xc1\xd1\x58\x1d\xa4\xd2\x35\x68\x5a\x26\x1d\x7c\x3a\x77\x08\x44\xfc\xea\xf8\x4f\xa3\xdf\xec\xb4\x49\x8d\x46\xa7\xdd\xef\xe2\xb0\x82\x5a\x76\xfb\x33\xd4\x4f\xcd\x9e\x53\x81\x7a\xb7\xd9\x23\x83\x9c\x75\x3b\x17\x15\x83\xcc\x89\x18\x1d\x4d\x04\xf1\xda\x4e\x46\x85\x4c\x0d\x05\x8f\xe0\x14\x1a\x5f\xf5\x9c\x19\x41\x38\x75\xea\x2d\xa4\xd5\x23\x64\x52\x71\x3a\xb9\x6a\x58\x16\x66\x24\x9d\x02\x6f\xc2\x20\x92\xb5\x92\xc4\xb6\x77\x70\x70\x90\xe5\x33\x73\xb3\x49\x92\x92\x5b\xcd\x1c\x8a\x48\x59\x43\x37\xe4\xc1\xed\x21\xfc\x74\xce\xb0\xc6\x62\x24\xba\xd0\x66\x29\xfb\xa9\x02\x33\x00\xaa\x9a\x60\xc8\x61\xf8\x63\x72\xb3\x70\x0b\xc3\x87\x47\x30\x10\x37\x96\xe4\x7f\x50\xf3\x80\xbf\x13\x4c\x90\x16\x82\x8e\x40\x13\xc5\x07\xb8\xef\xda\x7b\x13\x23\x20\xc4\xc4\xc4\xa3\x43\xd8\x3d\xa2\xdc\x3a\x66\xae\xff\x9c\xfc\x43\xa6\x5c\xa0\x2d\x58\x0d\xcb\x2b\x9b\xd0\x2a\x32\x69\xf5\x2a\x4c\x7a\x35\x73\xc2\x7d\x35\xae\xf9\x0c\x2b\x2f\xb3\xf4\xe0\xf9\x8c\x05\xf6\x54\x5c\x72\xa6\xc5\x7e\x4f\xf9\x75\xcd\x6c\x64\xa2\x5a\xfd\xdb\x98\x2d\x08\x4e\xbd\x93\x4d\xce\x3d\xd2\x95\x40\x32\x55\xbb\xea\x9f\x59\x3f\x3f\xb3\xf8\x7a\xbf\xf7\x7c\xee\xbe\xab\x17\x39\xb6\xb5\x70\x27\x86\x71\x6c\x53\x50\xd2\x8f\x81\xf0\x6f\x81\x23\x8a\xc4\x9c\x8b\x12\x9b\x7a\xa0\x6e\xe9\x77\xbe\xa2\xa4\x37\xc6\xaa\xae\x57\x94\x43\xd5\xfd\x62\xda\x45\x3d\xa9\x92\xd6\x84\x0d\xbe\x71\x64\xa4\x1f\x84\x42\x60\x4d\x21\xa4\xac\x36\x70\x57\x32\x7f\x3e\x89\x62\x43\x63\x5b\xae\xff\x35\x95\xea\x10\x2b\x4e\xc4\x8e\xb0\x95\xa0\xca\x84\x24\x77\x77\xff\xef\x08\x8b\x72\xc4\xac\x19\xa8\xfa\x8e\x85\x47\xa0\x57\x40\x36\x01\x7e\xe0\x21\x2d\x16\xe4\x80\x72\xe2\x0e\x79\x94\x88\x34\xf2\x2d\x4f\x04\x22\x39\x84\x1f\x87\xef\xe8\xef\xa2\xf9\x21\x76\x7d\x5f\x4b\x45\xd1\x30\x18\xe9\x99\x35\x33\x9f\x69\x92\xbd\x95\x3b\x78\xea\xf0\x58\x50\x69\x43\x3d\x4a\x65\x07\x38\x56\xc9\x33\xe6\x31\x00\x92\xe0\x89\x33\xe9\x35\xee\x32\x90\x48\x60\x61\x88\x8d\x50\x12\x25\xe2\xa2\xa1\xae\xf5\x03\xcc\x46\x22\x36\x4f\x70\x81\xf9\x73\x41\xb3\xcc\x6a\xbe\xdb\xdd\x7d\xe2\xa5\x52\x2a\x34\x76\x91\x98\x15\x90\xed\x20\x10\xde\xb7\x42\x6c\x87\xee\x8d\x95\x07\x09\x0a\x1b\xdf\x14\x1e\x7a\x01\x73\x13\x62\xa8\xc6\x05\xf8\xba\x85\x32\x33\x0e\xb8\xa9\x12\x4b\x4b\xa2\x60\x2d\x6d\x28\x34\x95\xcf\xaf\x9f\x3a\xac\x8a\xfa\x2e\x1b\xe7\x6e\x25\xa6\x72\x93\x93\xf5\x62\xce\xfd\x4c\x96\xc0\xf2\x84\xdd\x78\x3e\xbb\x66\xee\x66\x63\x19\xbb\xde\x74\xfc\xa4\x8a\xe6\x0f\x13\xd7\xe7\xa9\x3c\x84\xd7\x1a\x56\x92\x00\x86\xc3\x42\x16\xcb\xd0\x90\x08\x86\x82\x14\x01\xf7\xe1\x47\x76\x40\x7f\x8b\x89\x61\x38\x5c\xb0\xc5\x36\x64\x87\xb9\x24\x4f\x97\x25\xde\xad\x5d\x70\x05\xeb\x6a\x94\x49\x5e\x6a\xde\xee\xa2\x91\x75\x89\xca\xe7\xe3\x86\x4e\xb1\xa4\xcc\x5f\xfa\xdf\x5d\xed\x94\x55\xbf\x39\xef\xde\xee\xef\x37\xca\x0b\xd0\x3e\xc5\xb5\x09\xf9\x7a\xcb\x18\x2c\x7a\x2f\xc3\x2d\x5f\x91\xd3\x3f\xf3\x63\xa3\xd9\x79\x11\xe8\xf7\x2e\xa5\x2f\xbf\x76\x60\x0f\x27\xc8\xd9\x0b\x0f\xd4\x39\x81\xf9\xd1\xc6\x9a\xa3\x25\x7a\xef\x01\xb0\xca\x37\x3f\xe8\xa8\x15\x8e\x39\x56\xa6\xe5\xaf\x56\x0a\xce\x9f\xe5\xe0\xd9\x38\x79\x09\xd3\x4d\x8a\xd9\x3c\x78\xf6\xb2\xe0\xb9\x2b\x36\xb6\x3e\xf7\xad\x35\xfb\x76\x05\xc1\xb6\x87\x02\xe6\x9e\x69\x2e\xb9\x2b\x1c\x72\x35\x70\xe3\x96\xb0\x61\xcd\xdc\xe4\xfd\xef\x13\xc7\xc3\x34\x69\x9e\x9d\x9d\xe5\xc9\xd7\x67\x9e\x48\xf4\x3b\xb9\xe9\xf6\xa0\xb0\x21\xd8\xa7\xed\x40\x21\x6f\x0f\x44\xe0\x97\x27\x6e\x2f\x4d\x24\x51\x8f\x05\xcf\x00\xb3\x86\x82\x47\x9a\x68\xde\x57\x2c\x25\xf8\xb7\x24\x98\xa6\xa7\x5f\xa2\x62\xc2\x0c\x91\xa6\x1b\x73\x85\xf4\xff\x60\xa5\x49\xff\xf5\x9b\x9f\x99\xef\x96\xd4\xeb\x95\x19\x39\x58\x5b\xf9\x30\x2b\xe4\x33\xe0\xac\x7b\xc3\xf2\x92\xb9\xf7\xe4\x23\x67\x13\x7a\xff\x76\xef\x8b\xf6\x63\xdb\x2d\x8d\xe1\xa5\xc4\x5b\x9e\x7e\xb3\x3f\xf7\x9d\xd6\x94\x14\x85\x97\x25\xfb\xf7\x2c\x59\xa9\x12\x11\x8d\x9e\xcf\xb4\xbf\xae\xbf\x9c\xf2\x5b\x7e\x54\x77\x6c\x67\x42\x7e\x87\xa8\x2b\x69\x18\xf2\x27\xd3\x1b\x18\xcb\x67\x7e\x2f\x71\xf8\xcf\x88\xc3\xac\x35\x9d\x85\xda\xf1\xe0\xf9\xdc\x4c\xef\x11\xcb\x6c\x74\xcf\xd5\xa3\xf5\xf7\x83\x9e\x59\x99\xf5\xeb\x2e\xd7\xaa\x50\x0b\xe6\xa7\xfe\x59\x25\x78\xf6\xc8\x58\x90\x68\x5b\xc2\xe3\x5e\x8b\xde\x7b\x9f\xec\xbf\x34\x58\x16\x3b\xcc\xe5\x0b\x6e\xcf\xd4\x50\x4e\xdb\xad\x95\x9e\x12\xbb\x36\x96\x50\xf7\x57\x0c\xa7\xec\x8a\x1e\x35\x51\xdb\x97\x63\x1e\x57\x4d\x37\x6c\xef\x16\x2f\xc7\x94\xba\xf7\xa5\x2b\xdc\x9a\x6a\xbc\x75\x91\x89\x32\x8d\xb7\x50\xa6\xad\xb3\xd3\x43\x56\xf0\x5d\x1d\xf1\xcb\xc2\xfa\xdf\x6c\x73\x17\xb7\x5b\xb3\x4b\x86\xf3\x0d\xd7\x14\xf4\x0c\x5b\xae\xc5\x2b\x8f\x2f\xd1\xf8\xcf\x88\xc6\x97\x4d\xd7\xcb\xa6\xeb\x65\xd3\xb5\xed\xc1\xf2\xb2\xe9\xda\x9a\x96\x6d\x9d\xa3\x70\x36\x9d\xc7\x9d\x3c\xe0\x28\x74\x86\x32\x87\x3c\xf9\x4d\x8c\xc2\xd5\xa4\x85\x9b\x26\x73\x47\x1f\x1c\x1c\xdc\x75\xc0\x5d\x3c\xd9\x5d\x3d\x92\xdc\x8e\xa6\x61\x9b\xda\x97\xa7\x6c\x5d\xf6\xd7\xb6\x2e\xa5\x87\x68\xf7\xb9\x7c\xa1\xb7\x59\xba\xd7\x50\xbc\x85\xb5\x98\xae\x8a\x9f\xe0\x3e\x5d\x40\xec\x2f\x66\x2b\xad\xd1\xc6\xa9\x0a\x75\x82\xc1\xed\x66\xe7\x70\xab\xb9\x63\xe5\xbe\xc3\x72\x66\x38\xb6\x71\x99\x9f\x64\xff\x35\x8a\x69\x62\xdb\xda\xda\x35\xd7\xeb\x32\x15\xe7\xf9\xeb\xd8\xa6\x5b\xac\x04\xa1\xeb\xc0\x27\x86\x51\xfe\x8d\x6f\x9c\xca\xb1\x40\x8e\xdf\xe1\x13\xd7\x15\x52\xc5\x2f\x8b\xfe\x8e\x0f\xd8\xbe\xcf\xf7\x6b\x9b\x7f\xbe\xf6\xfd\xbe\x5e\x5b\xe0\xb9\x81\x25\xe7\xdf\xa9\x3e\xe0\x0b\xb2\xff\x04\x00\x00\xff\xff\xec\x0a\x83\x85\xc0\x3f\x00\x00")

func templateDefaultTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDefaultTmpl,
		"template/default.tmpl",
	)
}

func templateDefaultTmpl() (*asset, error) {
	bytes, err := templateDefaultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/default.tmpl", size: 16320, mode: os.FileMode(420), modTime: time.Unix(1474807852, 0)}
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
	"template/default.tmpl": templateDefaultTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"default.tmpl": &bintree{templateDefaultTmpl, map[string]*bintree{}},
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

