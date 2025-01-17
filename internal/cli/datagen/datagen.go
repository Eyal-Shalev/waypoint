// Code generated by go-bindata.
// sources:
// data/init.tpl.hcl
// DO NOT EDIT!

package datagen

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

var _initTplHcl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xbd\x6e\xdc\x30\x10\x84\x7b\x3d\xc5\x40\x2a\xdc\x38\x0a\xdc\x06\x70\xe1\xc4\x8d\x81\x74\xf9\xab\xf7\xa8\xbd\x3b\xc6\x24\x97\xe1\x4f\x2e\xca\x41\xef\x1e\xac\x78\x16\xec\x6b\xa2\x46\x24\x17\xf3\xed\xee\xcc\x80\xaf\x47\x46\x20\xcf\x90\x3d\x66\xa9\x09\x31\xc9\x4f\x36\x65\xc4\xc3\xcb\x11\x65\x8e\xd6\x90\x73\x33\x3c\xc5\x8c\xbb\x0f\x77\x28\x02\xc2\xf7\x4f\x5f\x90\x38\x4a\xb6\x45\xd2\x3c\x76\x4a\xb3\xb9\xe1\x7c\xcd\x05\x3b\x46\x0d\xf6\x57\x65\xec\x25\x35\xfc\x0f\x9a\xa3\xd8\x50\x90\x39\xfd\xe6\x34\xe2\x69\xed\x7b\x93\x18\xa9\x86\x60\xc3\x01\x36\x74\x03\x9c\x18\x72\xf0\x32\xf1\x2d\x8a\x52\xaf\x80\x45\x1a\xcf\x93\x39\xda\xc0\x63\xf7\x32\xec\x3d\x7a\x3f\xbf\xbb\xdc\xfa\xae\x1b\xf0\x99\x76\xec\x32\x0c\x05\xd5\xe7\xc8\xc6\xee\x2d\x4f\xeb\x4c\x92\x0e\x14\xec\x5f\x2a\x56\x02\x39\xc4\x9a\xa2\x64\xce\xba\x8b\x6b\xb2\x7b\x9c\xd1\xef\x45\x7a\x25\xef\x28\xf5\x58\x14\xfa\x10\x40\x31\x3a\x6b\x56\xa9\x8e\x33\x71\x74\x32\x8f\x1d\xc5\x88\xfe\xc4\xbb\x1e\xe7\x0e\x00\x06\x7c\xac\xd6\x4d\x5b\xe3\x8c\xa3\x9c\x40\x6f\xf5\xf9\x28\xd5\x4d\x3a\x5f\xc3\xf0\x34\xe2\x29\xb4\xcd\x0d\x65\xbe\xbd\xa0\x4e\x7c\xe3\x1c\x76\x2b\xb0\x66\x75\x8b\xf0\x28\xe6\x99\xd3\xde\x3a\x06\x85\x09\xcf\xcc\x71\xb5\xb1\xc0\x06\xd0\xc5\xc9\xc4\x07\x9b\x8b\xa6\xa4\xa0\x06\x68\xf3\xe9\x57\x33\xa3\x9f\x56\x4e\x8f\xf3\xb2\xbd\x6f\x87\x01\xdf\x82\x11\xef\x39\x68\x08\x4e\x4e\xba\xb1\xaa\x08\x89\xbd\x14\x46\x53\x6f\x7d\xb4\x1e\x6b\x3e\xb6\x94\xb4\x5f\x81\xf5\x74\x50\x6f\x37\xe8\x2b\xfc\xa6\x3b\xbf\x7a\xbc\x1e\xec\x4d\x09\x8d\xa7\xb1\x6c\xcb\xf1\x1f\xf2\xd1\xf1\x68\xc4\xbf\x5f\xab\xfd\x95\xa4\xd0\x01\x50\x89\xa3\xc2\xb9\xbc\x2d\x2f\xaf\x6e\x4b\xb7\x5e\x2e\xbf\x01\x8f\x6b\x2c\xba\x55\xb3\x7b\x7d\x6e\x59\xfd\xc7\xc7\xa5\x5b\xba\x7f\x01\x00\x00\xff\xff\xd8\x32\x99\x2d\x68\x03\x00\x00"

func initTplHclBytes() ([]byte, error) {
	return bindataRead(
		_initTplHcl,
		"init.tpl.hcl",
	)
}

func initTplHcl() (*asset, error) {
	bytes, err := initTplHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init.tpl.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"init.tpl.hcl": initTplHcl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"init.tpl.hcl": {initTplHcl, map[string]*bintree{}},
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
