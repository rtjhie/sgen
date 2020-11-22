// Code generated for package bintmpl by go-bindata DO NOT EDIT. (@generated)
// sources:
// template/main.tmpl
// template/model.tmpl
package bintmpl

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

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\x41\xab\xd4\x30\x10\xc7\xef\xf9\x14\x7f\xcb\x3b\xb4\x58\xb3\x78\x15\x16\x0f\xf2\x0e\x0a\x8a\xb0\x07\x0f\xeb\xf2\x48\xdb\x69\x37\xbe\x36\x29\x93\xac\xb8\x84\x7c\x77\x49\xb2\xad\xf6\x34\x09\xbf\xfc\xe6\x3f\xd3\x10\x30\xd0\xa8\x0d\xa1\x5a\x94\x36\x15\x62\x14\x62\x55\xfd\xab\x9a\x08\xe9\x46\x08\xbd\xac\x96\x3d\x6a\x01\x00\x55\x77\xf7\xe4\xaa\x52\x5b\x57\x89\x52\x4d\xda\x5f\x6f\x9d\xec\xed\x72\x60\xff\xeb\xaa\xe9\xe0\x26\x32\x0f\x2c\x04\xc8\xef\xaf\x13\x62\xac\x44\x23\xc4\xe1\x80\x4f\x76\x9e\xa9\xf7\x70\xfd\x95\x16\xe5\xd0\x29\x47\x03\xac\xc1\xa8\xd9\x79\xf8\xfb\x4a\xd0\xe9\x34\xd3\x47\xf1\x5b\xf1\x0e\x1e\x71\xbe\x24\xb3\xfc\x6c\x3c\xf1\xa8\x7a\x0a\xb9\x47\x08\xef\xc0\xca\x4c\x84\xa7\x97\x16\x4f\x46\x2d\x84\x0f\x47\xc8\x6f\x6a\x21\x97\x66\x4a\x50\x91\xc8\x10\x1e\x40\x8c\x21\xb6\xfb\x73\x32\x43\x02\xa3\x10\xe3\xcd\xf4\x79\xf6\xba\x41\xd1\xa7\x08\xb3\x36\xe4\x70\xbe\x9c\x2f\x69\x05\xf9\x7a\xb4\x8c\x97\xf6\xe1\x4d\xfd\x4a\x84\x2d\x6c\x79\x9b\xbe\xae\x05\x31\x27\x22\x87\xff\xaa\xd8\x5d\xd5\x7c\xca\x5c\x5d\xf0\x66\x87\xf5\x98\xe1\x37\x47\x18\x3d\xff\x27\xc9\x1d\x95\x9e\x6b\x62\xfe\x47\xc7\xbd\x2a\x01\x8f\x50\xeb\x4a\x66\xa8\xf3\xb1\x45\x57\xd0\x82\x59\x27\x4f\x7e\xb0\x37\x2f\x7f\xb0\xf6\x54\xe7\x9f\x29\xbf\x58\x6d\x36\xbc\x4c\x57\x57\x3f\x4d\xd5\x34\xcd\xbe\x8d\xad\x6f\x4a\x66\x79\x5b\x4b\xd1\x11\x73\xd1\x9d\x3c\x6b\x33\x25\x4c\x3e\x27\xac\x6e\xf0\x16\xd9\xb4\xd1\xcf\x7f\xb4\xaf\xdf\x67\x6d\x08\xdb\xc6\xff\x06\x00\x00\xff\xff\xe2\xa0\x4d\x8f\x83\x02\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 643, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateModelTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x41\x4b\xc3\x40\x10\x85\xef\xfb\x2b\x1e\x4b\x0e\x2d\xe8\x7a\x2f\x78\x10\x44\x50\xb0\x1e\xe2\x4d\xa4\xdd\x24\x93\x98\x36\xdd\xc6\xcc\xe6\x50\x96\xfc\x77\x99\x36\x8d\xa1\x54\xa8\xd0\x5b\x66\x79\xf3\xde\xcb\x97\xac\x0a\x01\x19\xe5\xa5\x23\xe8\xcd\x36\xa3\x4a\xa3\xeb\x94\xaa\x6d\xba\xb6\x05\x81\x0b\x72\x4a\xf9\x5d\x4d\x08\x01\x91\x99\xdb\x0d\xa1\xeb\xc0\xbe\x69\x53\x8f\xa0\x00\xe0\xf9\x51\xe6\xd2\x15\x58\xae\x78\xeb\x66\x7a\x51\x66\x7a\xa9\x42\xb8\x45\x63\x5d\x41\x88\x16\x37\x88\x72\xcc\xee\x11\x99\xa7\x92\xaa\x8c\x25\x44\x56\x45\x53\xe6\xa0\x6f\x44\xb9\x79\xdf\xd5\x64\xe2\x83\x93\xde\x26\x2b\x4a\xbd\x3e\x0a\x0f\x62\xd4\x96\x53\x5b\x89\xf8\xd8\x44\x6a\xe5\x26\x6e\x13\x4e\xbf\x68\x63\x87\xf3\xbe\x4a\x08\x60\x67\xd7\x34\x5a\xd1\xcb\x21\x9a\x2a\xa6\x0b\x13\xc6\xe5\x2e\xb6\x77\x99\xb8\x8f\x1e\x3b\xa5\x42\x18\x41\xe1\x03\x94\xa1\xfe\x00\x66\x40\xce\xe7\x99\x9f\x85\xcb\x27\x74\xff\x4d\xf8\xfa\x94\xff\x22\x7d\x65\xda\x27\xc4\xcf\x8c\x7b\xf2\xfd\x9c\xb7\x2e\xc5\x64\xfc\x43\x4f\xf1\x12\xbf\xcd\xe3\xfd\xcb\x4d\xa6\xf8\xf8\x4c\x76\x9e\x7a\xd4\x0d\xf9\xb6\x71\x48\x4a\x27\x2d\xcc\x6b\xcb\xfe\x81\x99\xfc\x44\xcb\x7c\xf7\x5b\xea\x68\x66\xe4\x5c\x4f\xe5\x63\xcb\xf5\xea\x43\x7f\x02\x00\x00\xff\xff\x7f\xc8\x2f\x37\x6a\x03\x00\x00")

func templateModelTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateModelTmpl,
		"template/model.tmpl",
	)
}

func templateModelTmpl() (*asset, error) {
	bytes, err := templateModelTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/model.tmpl", size: 874, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/main.tmpl":  templateMainTmpl,
	"template/model.tmpl": templateModelTmpl,
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
		"main.tmpl":  &bintree{templateMainTmpl, map[string]*bintree{}},
		"model.tmpl": &bintree{templateModelTmpl, map[string]*bintree{}},
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
