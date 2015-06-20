// Code generated by go-bindata.
// sources:
// ui/healthd.css
// ui/healthd.js
// ui/index.html
// DO NOT EDIT!

package healthd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
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

var _uiHealthdCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x52\xa8\xe6\xe2\x4c\xce\xcf\xc9\x2f\xb2\x52\x28\x4a\x4d\xb1\xe6\xaa\x05\x04\x00\x00\xff\xff\xb2\xc2\xbe\x83\x12\x00\x00\x00")

func uiHealthdCssBytes() ([]byte, error) {
	return bindataRead(
		_uiHealthdCss,
		"ui/healthd.css",
	)
}

func uiHealthdCss() (*asset, error) {
	bytes, err := uiHealthdCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/healthd.css", size: 18, mode: os.FileMode(420), modTime: time.Unix(1434836906, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiHealthdJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xcc\x49\x2d\x2a\xd1\x50\x2f\x2e\x2d\x28\x48\x2d\x2a\x56\xd7\x04\x04\x00\x00\xff\xff\x5b\x8e\xed\x8c\x10\x00\x00\x00")

func uiHealthdJsBytes() ([]byte, error) {
	return bindataRead(
		_uiHealthdJs,
		"ui/healthd.js",
	)
}

func uiHealthdJs() (*asset, error) {
	bytes, err := uiHealthdJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/healthd.js", size: 16, mode: os.FileMode(420), modTime: time.Unix(1434836922, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xc9\x28\xc9\xcd\xb1\xe3\x2a\x2e\x2d\xe0\xb2\xd1\x07\xb3\x01\x01\x00\x00\xff\xff\xb6\x22\xd6\x13\x12\x00\x00\x00")

func uiIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_uiIndexHtml,
		"ui/index.html",
	)
}

func uiIndexHtml() (*asset, error) {
	bytes, err := uiIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/index.html", size: 18, mode: os.FileMode(420), modTime: time.Unix(1434836885, 0)}
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
	"ui/healthd.css": uiHealthdCss,
	"ui/healthd.js":  uiHealthdJs,
	"ui/index.html":  uiIndexHtml,
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
	"ui": &bintree{nil, map[string]*bintree{
		"healthd.css": &bintree{uiHealthdCss, map[string]*bintree{}},
		"healthd.js":  &bintree{uiHealthdJs, map[string]*bintree{}},
		"index.html":  &bintree{uiIndexHtml, map[string]*bintree{}},
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
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
		err = RestoreAssets(dir, path.Join(name, child))
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