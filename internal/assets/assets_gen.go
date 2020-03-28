// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../assets/config/app.toml (670B)

package assets

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

var _configAppToml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\x41\x6f\xdb\x30\x0c\x85\xef\xfa\x15\x84\x7a\xed\x3c\xbb\x01\x3a\xb4\x58\x86\xed\xb2\xd3\x76\xdb\x2d\x08\x02\xda\x62\x6a\xa2\x92\xa5\x90\x52\x32\xff\xfb\x41\xf6\x92\xe6\x62\x40\x4f\xf4\xf7\x9e\x1e\x77\x98\x92\xe7\x01\x33\xc7\x69\x6f\x26\x0c\x04\x5b\xb0\x43\x3c\xb3\xb3\xe6\x4c\xa2\x1c\xa7\xaa\xb4\x4d\xd7\xb4\xd6\x60\xc9\x63\x14\x85\x2d\xec\xec\x6f\x1e\x46\x24\x0f\xbf\x18\xbe\xa2\xe7\x30\x7f\x7f\x63\xcf\x79\x6e\x26\xca\xdf\xec\xde\x38\xd2\x41\x38\xe5\xff\x84\x85\xd9\xbd\x00\x4f\xc7\x28\x61\x31\x04\x25\x39\xf3\x40\xd6\x3c\xc0\x8f\x94\x40\xca\x04\x21\x3a\x7a\x04\xe2\x3c\x92\x80\x75\x74\x26\x1f\x93\x7d\x04\x9b\x24\xba\x32\x64\x6b\xa4\x4c\x87\x3a\x55\xa1\x37\xd1\xec\x2a\x8b\xf6\x06\x9d\x93\x7a\xf3\xda\xb5\xdd\xf3\x4b\xbd\x70\x98\xb1\x47\xa5\xbd\x79\x80\x3f\x23\xc1\xf5\x0c\x3d\x0e\xef\x34\xb9\x0f\xb7\x14\x35\xbf\x09\x69\xb5\x0b\xb3\x9e\xbc\x05\xab\x27\xcf\x99\x36\x16\xa2\x80\x0d\x5a\xc5\xc6\xe4\x39\x2d\xf6\xeb\x90\x19\xa3\xe6\x7a\xec\x9e\xbe\x34\x6d\xd3\x36\xdd\xeb\x66\xd3\x3e\xdb\x5b\x9d\xae\x3f\x50\x62\x47\x81\x07\x6b\x8a\x92\xdc\xf5\x61\x4d\x42\xd5\x4b\x14\x07\xdb\x55\x7b\x6a\x57\x55\x30\xd4\xb9\x5a\xce\xcf\x78\x9f\x0e\xe2\xe4\xe7\xbb\x8e\x58\xb1\xf7\x54\x43\x0b\x9d\x0a\x0b\xad\x61\xcf\x24\x7c\x9c\x3f\x1d\x8b\xaf\x91\x55\xfd\xad\xb5\xeb\x1f\x57\xf2\xc7\x1b\x17\x70\xc0\x77\x02\x2d\x42\x90\x23\x14\x25\xc0\x5e\xa3\x2f\x99\x20\x61\x1e\x1b\x53\xbf\x0b\x06\x33\x7e\xbe\x3e\xac\x71\xfd\x52\xf6\xba\xb0\xda\xf5\x65\xa4\x25\x60\x25\x1c\x2b\x92\x43\xf2\x14\x68\xca\xeb\xf2\xe3\x11\xfe\x46\x09\x75\xe3\xc5\x93\xa9\xb4\x0b\x0a\x1d\x96\xd9\x2d\x64\x29\xf4\x2f\x00\x00\xff\xff\x88\x2b\xc1\x86\x9e\x02\x00\x00"

func configAppTomlBytes() ([]byte, error) {
	return bindataRead(
		_configAppToml,
		"config/app.toml",
	)
}

func configAppToml() (*asset, error) {
	bytes, err := configAppTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/app.toml", size: 670, mode: os.FileMode(0644), modTime: time.Unix(1585411209, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc6, 0x8d, 0xc8, 0x23, 0xcd, 0xc2, 0x2, 0xec, 0xdd, 0x7d, 0xf2, 0x79, 0x3a, 0x4e, 0x13, 0x19, 0x11, 0x3b, 0xbb, 0xc5, 0x3d, 0xec, 0xb0, 0x8d, 0x33, 0xa0, 0xe9, 0x8a, 0xf6, 0x13, 0x6b, 0x90}}
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
	"config/app.toml": configAppToml,
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
	"config": {nil, map[string]*bintree{
		"app.toml": {configAppToml, map[string]*bintree{}},
	}},
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
