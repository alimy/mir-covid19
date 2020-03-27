// Code generated by go-bindata. DO NOT EDIT.
// sources:
// public/favicon.ico (4.158kB)
// public/index.html (2.575kB)

package web

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
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

var _faviconIco = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x97\x5b\x4c\x93\x67\x18\xc7\x5f\x4e\x82\x56\xa0\x90\x69\x32\x97\x4d\x77\x70\xb2\xcd\xab\xed\x6a\x73\x17\x26\xdb\x8d\xc9\xa2\x82\x6e\xb8\x79\x62\x73\x51\x23\x3a\xc8\x12\x8d\x13\x75\x55\x14\x15\x15\xb1\x03\x8a\x1a\x75\xea\x30\x73\x4a\xa5\x8a\x03\x51\x54\x14\x84\x01\x72\x5a\x41\x7b\x02\x6c\x41\x68\x69\x4b\x0f\xd0\xf3\xd7\xff\xf2\x7d\xc5\x06\xd3\xb7\x96\x6d\x17\xde\xf8\x24\x4f\x9a\x7c\xe1\xff\xfe\x9e\xef\xfd\x9e\xf7\xff\x3e\x10\x12\x46\xc2\xc8\xcc\x59\xec\xef\x2c\x92\xc4\x27\x64\x3a\x21\x24\x89\x10\x32\x93\x10\x32\x9f\xf8\x9e\xb3\x21\x88\x27\x24\x91\xe7\xcb\xa7\x81\x97\xf1\xc2\xc3\xeb\x71\xc1\xae\x6c\xc6\xa8\xf4\x36\x6c\xd2\x3b\xbe\xec\xba\x0b\xd7\x13\x39\xbc\x6e\xe7\x73\xb5\x2e\x5d\x2f\x6c\xe3\x75\xd2\x3b\x70\xf4\xb4\x81\x19\x35\x4d\x88\xcd\x8c\x0c\x43\xbd\x6d\x1e\xe4\x5f\x4d\x82\x7c\xd9\x64\x2e\x15\x5f\x4f\x81\xea\xfb\x19\x18\x3a\xbb\x05\x8c\xdd\x12\x54\x6b\x2c\xcb\x7d\x46\xc7\xa6\x72\x55\x22\xfa\x72\xbe\x80\x6b\x40\x31\x01\xb8\x07\x4f\x0e\x2f\xc3\xa3\x45\x04\xb2\x94\x71\xb9\x98\x40\xb6\x34\x0a\xe6\x3b\x67\x83\x4a\xad\x0d\x62\xc8\xbf\x8c\x82\x2c\x79\x9c\x2e\x99\x70\x6b\x0d\x1c\x5d\x05\xaf\xc7\x1d\x12\x6f\xb8\xb4\xd7\xc7\x4a\x79\x36\xb9\x35\x7e\x49\x03\xbc\x0c\x55\xe7\x54\x4b\xa1\xfc\x76\xfa\xb3\xec\xb1\xba\x7b\xd2\xdf\x85\x5b\xdf\x17\x92\x6d\xaa\x3a\x16\xa8\x1f\x5b\xe3\xf1\xe6\x8f\xe0\xb1\x1a\xa9\x3a\xf7\x90\x1a\xdd\xeb\x67\x05\x6a\x93\x09\x14\xdf\x4c\x85\xad\xb3\x26\x24\x7b\xa4\xb5\x12\xf2\xd4\x98\x40\x76\x32\x81\x32\x6d\x1a\x1c\xbd\x1d\x54\x1d\x33\x6a\x86\x7a\xeb\xc7\x78\x44\xd9\x33\x36\x87\x2b\x8a\x42\xb2\x5d\x03\x4a\xa8\xd6\xbe\x4e\x7d\x77\xb6\x7f\xd8\x5e\xa6\x86\xd7\x8b\xc1\xc2\x35\x81\xbd\x32\xf6\xbd\x74\xa7\x32\x43\xb2\x19\xc7\x08\x34\x3b\xe6\xd3\xeb\x5f\x1a\x01\xcb\xbd\xf3\x41\xb5\xc3\xd7\x84\xd4\x9a\xb9\x5e\x11\xae\xe6\xea\x0b\x15\x83\x05\xdf\xd1\xeb\x5f\x4c\x30\x74\x6e\x6b\x50\x9d\xef\x7b\x45\x53\x75\xea\x9f\x3e\x01\x63\x33\x87\x64\xeb\xce\x6c\x0e\xba\x77\x7d\x7b\x16\xc0\xeb\x72\x50\x75\x76\x59\x3d\x14\xcb\x63\xa9\xfd\xa6\x5a\xf3\x2a\x9c\x7d\x0f\x43\xb2\xcd\xb7\x4e\x43\xb6\x24\x8c\xda\xeb\x3d\x1b\xe7\xc0\x6d\x7c\x42\xd5\xb9\xf5\x1a\x74\x6f\x78\x87\x7a\x46\x59\xdf\xb1\x36\x4a\x42\xb2\xed\x8a\x46\x28\x57\xf0\xb9\x7a\xe5\xe3\x92\x5d\xb3\x7b\xcd\x6b\x5c\x3f\xd2\x82\xdd\x8f\xbe\xec\x05\xd4\x3d\x63\xd7\x32\x5d\x17\xf9\xff\xd6\xc1\x8c\xc0\xce\x58\x61\x67\x2c\xb0\x31\x66\xd8\x18\x13\xf7\x6b\x31\xc8\xf0\x70\x4b\x12\xda\x97\x13\xb4\xaf\x0c\x47\xfb\xaa\x70\xb4\x71\x19\x86\xf6\xb5\x53\xa1\x95\x8a\x61\x76\x0f\xc0\xec\xd4\xc0\xe4\x54\xc3\xe4\x1a\x4b\xb7\x06\x9a\xd3\x6b\xa9\xef\xcd\x3e\xd3\x5f\x10\xf8\xd9\x4d\x96\xcb\xa8\x35\xfd\x86\x9a\xe1\x53\xb8\x65\x2c\xc6\x0d\x83\x10\x95\xfa\x83\x28\x1f\xdc\x85\x92\x2b\x73\x71\xfa\x4c\x3c\x4e\x96\xf0\x71\xe2\xf7\x78\x1c\xbb\x18\x07\x91\x38\x16\x85\x92\x58\x14\xdd\x7b\x13\x45\x0f\x3e\x40\x41\xdb\x7b\x10\x76\xcc\x41\x7e\xe7\x6c\xe4\x3d\x7a\x1b\x87\x94\x6f\x41\xfc\xeb\x34\xdf\x1e\x51\x7a\xa5\x3f\x37\xc5\xef\xad\xcd\x16\x09\x6e\x18\x45\xa8\x34\xe4\xa3\x7c\xe8\x00\x24\xba\xdd\x28\xd5\x66\xe1\x8f\x81\x1f\x71\xf2\xf6\x87\x28\x2a\x4d\x40\x81\x24\x11\xc2\x6b\x09\x38\x52\xc9\xc7\xe1\x1b\xf1\x38\x78\x3b\x0e\xfb\x6b\x78\xc8\xa9\x9d\x8c\x3d\x0d\x31\xd8\xdd\x34\x09\x82\x96\x28\xec\xec\x88\x40\xd6\xc3\x70\x9c\xbd\xc0\xe3\xce\x22\xed\xbd\x7b\x33\xe7\xc2\x63\xd2\x72\x6c\xe9\x48\x35\xc7\xbe\x6e\x38\x8a\x6b\xfa\x5c\x5c\xd1\x65\x43\xac\xdd\x8e\x8b\xda\xcd\x38\xd3\xf8\x19\x44\x97\x12\x50\x50\x96\x08\x61\x79\x02\xf2\x2b\xf8\xc8\x63\xd9\xb7\xe2\x70\xa0\x26\x16\xfb\xea\x78\xd8\xdb\x30\x05\xd9\x4d\x31\xd8\xd5\x1a\x8d\x9f\x3b\xa2\xb0\xbd\x2b\x02\x45\x37\xe3\xd0\xb5\x22\x9a\xee\xad\x2b\x13\x60\x57\x35\x73\x6c\xa5\xad\x11\x37\x8d\xc5\x14\xf6\x16\x9c\xeb\x5c\x02\x91\xf8\x95\x7f\xc5\xde\xd1\x19\x81\x03\x4d\x3c\xb4\xa4\xf3\x20\xa7\xf5\x7a\x6a\x0c\x46\x3b\xaa\x39\xb6\xd5\xa3\x47\x9d\xa9\x04\x15\x86\x23\x1c\x5b\xc2\xb1\xb3\xb8\xf7\x2e\x51\xad\x46\xf1\xd5\x19\x28\x28\x4b\xf0\xb3\x0f\xd3\xd8\xcd\x3e\xf6\xce\x31\xf6\x2e\x69\x34\x6a\x76\xc6\x52\xd9\xec\xb9\x35\x57\x9f\xf2\xf7\xdb\xa0\x53\xc1\xf1\xaf\x73\xdf\x7c\x1f\x24\x3a\x01\xc4\xda\x6d\xb8\xf0\x78\x3d\x4e\x5c\x79\x03\xa2\x52\x3e\x0a\xcb\xf8\x10\x5e\xe5\x23\xbf\x22\x1e\x79\x55\x71\x38\x54\x3d\x15\xb9\x35\x3c\xec\xaf\x9d\x82\x9c\xfa\x18\xec\x69\x9c\x84\xdd\x2d\x51\x10\xb4\x46\x40\xd0\x1a\x86\x3f\x73\x22\xa1\x08\x72\x0f\xeb\x4e\x66\x8c\x3f\x95\x70\x30\xa3\x30\xba\xfb\xa1\x77\xf5\x62\xc8\xd5\x03\x9d\x4b\x85\x41\x6b\x07\xda\xf3\xe6\xa1\x61\x5d\x18\xea\x37\x44\xa2\x3e\x3d\x12\xf7\x37\x46\xe2\xfe\xa6\x08\xd4\xfd\xe0\xcb\xda\x8c\x70\xd4\x66\x86\xe3\x5e\x66\x18\xee\x66\x10\xd4\xed\xe0\xa3\xb5\x34\x15\xaa\xe3\xa9\x90\xa7\x04\x7a\x13\xe7\x8b\xd9\xac\x2f\xda\x9f\x6f\x30\xec\xbd\x24\x4c\x83\x6c\xe1\x98\xb7\x2c\xa6\xe4\x22\x02\xf9\x42\x82\xee\xd5\xd3\xa1\x15\xa6\xc1\xd6\x52\x05\x78\x3c\x3e\x5f\x5f\x16\x78\x0f\xfb\x7d\x7d\x02\x33\x9c\x41\xbc\x3f\xa8\x4f\x70\x1e\xb7\x6e\x26\x37\xc7\xb1\x3e\xce\xce\x5b\x4f\xc3\xd9\x2f\xe3\xfc\x9b\xd6\xeb\xac\x66\x22\xbe\x1e\x30\xc3\xb0\xcc\x25\x11\xe8\xd9\x94\x04\xc3\xc5\x6c\x38\x1e\xff\x4d\x9d\xa1\xdc\x86\x7e\x74\xa7\xcf\xa6\xd6\xcd\xce\x9e\xd6\xbf\x2e\x87\x64\x3f\x9d\x61\xb8\xbb\x7c\x69\x24\x37\x33\x99\xaa\x8e\x73\xf3\x30\xdb\x27\xc1\xc2\xcb\xb8\x31\x58\xbc\x8e\xee\xeb\x29\x04\x43\xe7\xb7\x87\x64\xb3\x77\x46\x4f\xc6\xfb\x50\x67\x7d\xca\xcd\xa8\xc1\xee\x2f\x5a\xb0\xfb\xae\x11\x7c\xee\xbb\xcf\x93\x89\x3f\xd9\x7a\xd8\x59\x38\x54\xb0\xff\x2b\x38\x35\x5d\x41\xe7\xc3\x50\xe1\xb1\xe8\x31\xf2\xa0\x1c\xc3\x95\x22\x0c\x57\x16\xf9\xb2\xa2\x10\xa3\x6d\x55\xff\x69\xbd\x97\xf1\xff\x82\xbc\xe0\xf8\x27\x00\x00\xff\xff\x08\xed\x9e\xf4\x3e\x10\x00\x00"

func faviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_faviconIco,
		"favicon.ico",
	)
}

func faviconIco() (*asset, error) {
	bytes, err := faviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "favicon.ico", size: 4158, mode: os.FileMode(0644), modTime: time.Unix(1585310817, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x56, 0x2, 0xfa, 0x79, 0xb6, 0xe7, 0x11, 0x13, 0x59, 0xf8, 0xcb, 0xfa, 0xce, 0xdb, 0x7d, 0x4b, 0x60, 0x2b, 0xe7, 0xad, 0xeb, 0x7c, 0x6, 0xa2, 0x63, 0x40, 0x82, 0xde, 0x29, 0x53, 0x83, 0xc0}}
	return a, nil
}

var _indexHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x5d\x6f\xdc\x44\x17\xbe\xef\xaf\x98\x58\xaa\xbc\x91\xd6\xf6\xa6\x6d\xde\xbe\xdd\xac\x17\x25\x25\x12\x48\xb4\xa0\x92\x4a\xa0\xaa\xaa\x66\x3d\x67\x77\x67\x63\xcf\x38\x33\xc7\xf6\x2e\x55\x24\x10\x2d\xa4\x04\x51\x24\x82\x40\x01\x24\x90\xfa\xc5\x4d\x12\x09\x24\xaa\x36\xe4\xd7\xc4\x9b\xdc\xe5\x2f\x20\xdb\xf9\xf2\x26\xad\x9a\x3b\x2e\xf0\xc5\xda\x67\xce\x79\x9e\xe3\x73\x9e\x33\xe3\x6d\x8c\xbd\xfd\xfe\xd5\xb9\x8f\x3f\x98\x25\x5d\x0c\xfc\xe6\xb9\x46\x76\x23\x3e\x15\x1d\xd7\x00\x61\x10\x8d\x03\x1f\x5c\x23\xe1\x0c\xbb\x75\x32\x51\xab\xc5\xc9\x94\xd1\x3c\x47\x08\x21\x8d\x2e\x50\x56\x3c\xe6\x66\x00\x48\x89\xd7\xa5\x4a\x03\xba\x46\x84\x6d\xeb\xff\x06\x71\x46\x03\xba\x88\xa1\x05\x0b\x11\x8f\x5d\xe3\x23\xeb\xe6\xb4\x75\x55\x06\x21\x45\xde\xf2\xc1\x20\x9e\x14\x08\x02\x5d\xe3\xdd\x59\x17\x58\x07\xca\xf8\x31\xcb\x22\xf9\x8b\xec\x6d\xae\xa6\x6b\x7f\xa7\x2f\x9e\x6c\x3f\x7f\xb1\xbb\xb6\x95\x3e\xfa\xb2\x30\x09\x17\x1c\x39\xf5\x2d\xed\x51\x1f\xea\x24\x5d\xfa\x25\x7d\xba\xbc\xb3\x7a\x6f\x67\xf3\xf7\xe1\xca\xd6\x70\x7d\x65\x7b\x6b\x99\xec\x6d\x2e\xed\x7e\xfd\x79\xfa\xd3\x9f\xdb\x2f\xbf\x69\xd6\xd2\xa5\x0d\x32\x51\x23\x7b\x9b\x0f\x48\xc0\x05\x0f\xa2\xe0\x10\x7d\xff\xb3\xdd\xb5\xe7\x3b\x2b\xcf\x86\x4b\x7f\x15\x04\xe9\xd2\xc6\xce\xea\xbd\xe1\xcf\x9f\xa6\x1b\x0f\xf7\xc9\x02\xda\x7f\x53\xcc\xa3\xa7\xfb\x98\x48\x83\xca\x01\xb4\x95\x61\x8a\xe8\xe1\x8f\xeb\xe9\xb7\x4f\xd2\x87\xeb\xdb\x2f\x1f\x0f\x1f\x2c\xa7\x5f\x3d\x2b\xf0\xc4\xb2\x46\x5a\x78\x68\x65\x97\xa0\x01\xb8\x46\xcc\x21\x09\xa5\x42\xa3\xe4\x3b\xec\x66\xde\x34\x97\x41\xcc\x3d\xb0\x72\xa3\x5a\x6e\x95\x3b\x61\xd7\xaa\xa5\x52\x8a\x95\xa2\x21\xf9\x73\xe9\xa5\x5d\x21\xab\x07\x39\xad\x36\x47\xd7\x93\x31\xa8\xa3\xec\xc7\x55\xf3\xb9\x98\x27\x0a\x7c\xd7\xe0\x9e\x14\x06\xe9\x2a\x68\xbb\x46\xe3\xbc\x4b\x66\xa6\x3f\x9c\xbd\x73\xf3\xc6\x7b\xe4\x7c\xd3\x69\xd3\x38\x73\xdb\xdc\x93\x65\xd1\x91\xa3\x0f\xcd\x86\x53\xdc\xcb\xc3\xd0\x8a\xb8\xcf\xea\x8c\x22\xbd\x46\xc3\x72\xa3\x32\x37\x08\x96\x47\x94\x3d\xda\x53\x3c\x44\x82\x83\x10\x5c\x03\xa1\x8f\x4e\x8f\xc6\xb4\x58\x35\x88\x56\x9e\x6b\x64\x33\xaa\xeb\x8e\xa3\x40\xdb\x49\xdf\x5e\x58\xb0\x3d\x19\x38\x32\x04\xe1\xf4\xb4\xd3\x4b\x80\xf7\xb9\xb0\x26\xec\x8b\xf6\x05\xbb\xa7\x8d\x66\xc3\x29\xe0\x27\xb3\x94\xe8\x02\x8a\x13\x76\x07\x79\xd0\xc9\xf9\x92\x24\x71\x14\x50\xc6\x45\xc7\xa1\x82\x29\xc9\x59\x4e\xaf\x69\xc8\x8b\xdf\x53\xc8\x1b\xce\xd1\xae\x6b\xb4\x24\x1b\x10\x29\x50\x46\x5e\x57\x23\x55\x78\xb0\x5d\x33\x35\xda\xbe\x4c\xac\x7e\x9d\x74\x39\x63\x20\x0e\x36\x6d\x8e\x63\x3c\x26\x9c\xb9\x06\x0d\xc3\x8c\x9f\xf1\x78\xa4\x73\x5c\x6a\xb2\xbb\xf5\x5d\x7a\xff\xf1\xf0\xb7\x2f\x86\x3f\xfc\x3a\xfc\xe3\xfb\xc2\x1c\x69\x65\x96\xad\x59\x9a\xb9\x5b\xfb\x43\x07\x8c\x63\x3e\x29\x26\xaa\x08\xcc\xdb\xd5\x52\x14\x17\x61\x84\xe5\xa5\x4c\x09\xaa\x80\x92\xbb\xa5\xe5\xec\xb2\x12\x68\xcd\x73\xb4\x8a\x11\x04\x1f\x3c\xac\x13\x1a\xa1\x24\x63\x3c\xc8\x46\x90\x0a\x9c\x3a\x89\x9a\xcf\x8e\xb2\xb3\x82\x02\xf9\xc9\x99\x21\xfa\xac\x08\x79\x46\xc0\x9b\x47\x2f\x1e\x89\xe3\x8c\xa8\xd3\x18\x9d\xd1\x5c\x88\x36\xa9\xdc\x80\xce\x6c\x3f\xac\x98\x6d\xe0\x03\x2a\xec\x04\x3c\x8e\x83\xfd\xa1\x37\xc7\x6d\x04\x8d\x95\x84\x0b\x26\x13\xdb\x97\x1e\x45\x2e\x85\x2d\x15\xef\x70\x31\x3e\x7e\x8a\x5a\x31\x55\xe4\x4e\x80\xd4\x23\x2e\xb9\xbb\x78\xc2\x3d\x55\x69\x47\xc2\xcb\x48\x48\xe5\x34\xf8\x01\x45\x80\x94\xb8\x84\x49\x2f\x0a\x40\xa0\xed\x29\xa0\x08\xb3\x3e\x64\x56\xc5\x2c\x6a\x31\xc7\x4f\x85\x07\x48\x6d\xad\xb2\xfc\xa6\xe3\x84\x5c\x74\x7a\xfa\x60\x0f\x77\x27\x1d\x8d\x14\xb5\xdd\xd3\x6f\xc5\x17\xec\x9a\x7d\xc9\x7c\x35\x05\xe0\x34\xa2\xe2\xad\x08\xa1\x62\x66\x07\xad\x59\x25\xe6\xb5\xb9\xe9\x77\x26\x5f\x97\xb8\x84\xd2\x9c\x65\xa0\xc9\x5a\xed\x72\xed\xca\x95\xcb\xff\x7b\x05\x30\x2b\x58\x1f\x2f\xb7\x03\xb8\x5f\xab\x9e\x19\xcc\xd1\xce\x75\x1a\xc0\x51\xd5\xb7\x6a\xb7\x4f\xa5\xd1\x76\x48\x15\x08\xbc\x2e\x19\xd8\x5c\x68\x50\x38\x03\x6d\xa9\xa0\x12\x20\xad\x12\x7d\x32\xf9\xe2\x78\xa5\xbc\xb8\x48\xc0\xd7\xf0\x9f\xac\xaf\x47\x95\x65\xbd\x74\xf1\xdf\x2f\xeb\xf1\x73\xa1\xf4\x39\xc9\xbe\x21\xcd\x73\x0d\xa7\xf8\xfb\xf7\x4f\x00\x00\x00\xff\xff\xd0\x71\xcf\xcc\x0f\x0a\x00\x00"

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 2575, mode: os.FileMode(0644), modTime: time.Unix(1585310817, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x35, 0x2d, 0x45, 0x84, 0xcb, 0x48, 0xd6, 0x74, 0x17, 0x7a, 0x7d, 0x9f, 0x4d, 0xe6, 0x58, 0xe7, 0xda, 0x49, 0xf1, 0x5e, 0x8, 0x32, 0xd0, 0x8b, 0xea, 0x80, 0xbd, 0x70, 0x93, 0xf4, 0xf1, 0xda}}
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
	"favicon.ico": faviconIco,
	"index.html":  indexHtml,
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
	"favicon.ico": {faviconIco, map[string]*bintree{}},
	"index.html":  {indexHtml, map[string]*bintree{}},
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