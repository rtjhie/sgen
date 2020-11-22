// Code generated for package bintmpl by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.go
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

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x4d\x8f\xdb\x20\x10\x3d\xc3\xaf\x98\xfa\xb0\x82\x08\xb9\xf7\x54\xb9\x55\x95\xf6\xd0\x0f\x29\xed\x69\xb5\xea\x12\x7b\xc8\xd2\x1a\x6c\x01\x3e\x44\x51\xfe\x7b\xc5\x80\x13\xa7\xbb\x9c\x60\x3e\xde\x7b\xf3\x98\x49\x77\x7f\xf5\x11\x21\x1e\xd1\x73\x6e\xdd\x34\x86\x04\x82\xb3\x06\x7d\x37\xf6\xd6\x1f\x3f\xfe\x89\xa3\x6f\x38\x6b\x02\x9a\x01\xbb\xd4\x70\xc9\x79\x3a\x4d\x98\xab\xf6\xdd\x2b\x3a\x0d\x31\x85\xb9\x4b\x70\xe6\x8c\x7d\xd3\x0e\x01\x72\xc4\xfa\x23\x94\xf3\x92\x21\xb6\x8d\xd7\x0e\x9b\x17\xce\xd8\x17\x8b\x43\x1f\xe1\xe9\x79\x43\xb7\x47\x6f\xc6\xa5\xc6\x50\x2a\x57\x5d\x38\xbb\x65\xdf\x12\xac\x38\xfe\x87\xff\x99\xc5\xd1\x59\x6e\xb5\x22\xab\xa6\x8a\xfd\x7c\x88\x45\xf9\xa6\x4e\x50\x2b\xe2\x92\x50\xa3\xb3\x09\xdd\x94\x4e\x45\x8b\xe4\xdc\xcc\xbe\x83\xaf\x3a\xc4\x57\x3d\x94\x2e\x51\x41\x1e\x7d\xc2\x60\x74\x87\x12\xc4\xd3\xf3\xe1\x94\x50\x01\x86\x30\x06\x99\x15\x47\xd8\xee\xe0\x87\x0e\x11\xef\xba\x24\x67\x01\xd3\x1c\x3c\x64\xea\xb6\x02\x8b\x28\xf9\xa5\x72\xfd\xf2\xee\x8e\xed\x00\x05\x5c\x82\xa8\xb2\xd7\x34\x2e\xd3\x3c\x94\xf8\xf9\xc2\x99\x35\x39\x99\x83\x84\x7f\x05\x13\x07\x05\x4e\x7e\xa2\xe4\x87\x1d\x78\x3b\x90\xad\x55\x8b\xb7\x03\x81\x92\xff\x35\xe6\x54\x0e\x5f\x65\xbd\x1d\x65\x6d\xc0\x62\xe8\x99\x14\xd4\xfc\xee\x3d\x1e\xa2\x20\x73\xaa\xe8\xfa\xb9\x5b\xa8\x9b\xd6\xe6\xef\xfb\x6e\x16\xbf\xda\x9c\x14\x52\x51\x9f\x19\x03\xfc\x56\x60\x72\x7b\xd0\x3e\x2f\x30\x55\xb5\x65\xb7\x04\x59\xc2\x62\x7d\xc2\x0e\xf4\x34\xa1\xef\xc5\x12\x51\x65\x0c\x7a\x08\xd3\x7e\xc6\xd8\x05\x3b\xa5\x31\x08\x29\xe5\x7a\xf8\x87\x78\x3f\x78\xe9\xe8\x61\x73\x6b\x91\xb0\xda\xe3\xf3\xad\xf3\x1a\xbc\x4d\x46\xa7\xa7\x49\x54\xdd\xd4\x6b\x30\x3f\xd4\x7a\x39\xb7\x77\x56\xf7\xed\x7e\x3e\x94\x7b\xf1\xe0\xc2\xff\x05\x00\x00\xff\xff\xed\xe0\x73\x7b\xbc\x03\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 956, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5f\x6b\xdb\x3e\x14\x7d\xb6\x3e\xc5\xf9\x99\xfe\xa8\xdd\xa5\x4a\xdb\xb7\x0d\xf2\x50\xda\x0c\x32\xb6\x76\x90\xc2\x1e\xba\x52\x14\xfb\x3a\x11\x75\x24\xef\x4a\x29\x0b\x42\xdf\x7d\x48\x4e\xc2\xf6\x64\x4b\xe7\xdc\xf3\x47\x37\x84\xe9\x85\xb8\xb3\xc3\x9e\xf5\x7a\xe3\x71\x73\x75\xfd\xf1\x72\x60\x72\x64\x3c\x3e\xab\x86\x56\xd6\xbe\x61\x61\x1a\x89\xdb\xbe\x47\x26\x39\x24\x9c\xdf\xa9\x95\xe2\x69\xa3\x1d\x9c\xdd\x71\x43\x68\x6c\x4b\xd0\x0e\xbd\x6e\xc8\x38\x6a\xb1\x33\x2d\x31\xfc\x86\x70\x3b\xa8\x66\x43\xb8\x91\x57\x47\x14\x9d\xdd\x99\x56\x68\x93\xf1\xaf\x8b\xbb\xf9\xc3\x72\x8e\x4e\xf7\x84\xc3\x1d\x5b\xeb\xd1\x6a\xa6\xc6\x5b\xde\xc3\x76\xf0\x7f\x99\x79\x26\x92\xe2\x62\x1a\xa3\x10\x21\xa0\xa5\x4e\x1b\x42\xb9\x55\xda\x94\x88\x51\x4c\xa7\xb8\x4b\x79\xd6\x64\x88\x95\xa7\x16\xab\x3d\xce\xc9\xf8\xe6\x74\x75\x2e\x71\xff\x88\x87\xc7\x27\xcc\xef\x17\x4f\x52\x0c\xaa\x79\x53\x6b\x42\xd2\x10\x42\x6f\x07\xcb\x1e\x95\x28\x4a\xeb\x4a\x51\x94\xab\xbd\xa7\xf4\x13\x02\x3c\x6d\x87\x5e\x79\x42\x39\xb2\x5c\xb6\xcc\xd0\xc0\xda\xf8\x0e\xe5\xff\xbf\x4a\xc8\xef\x07\xc5\x18\x45\x9d\x63\x9e\xad\x94\x23\x7c\x9a\x21\x7f\x8f\x78\x9a\x7d\x57\x0c\xd7\x6c\x68\xab\x1c\x66\x78\x7e\x21\xe3\xe5\xc2\x78\xe2\x4e\x35\x14\xb2\x34\x2b\xb3\x26\x9c\xbd\x4e\x70\x66\xd4\x36\xcb\xc8\x07\xb5\x25\x97\xf4\x8b\x22\x84\xcb\x83\x7e\x8c\x32\x1d\x4e\x51\x5c\x88\xe5\x61\x26\xc6\x49\xd6\x22\xd3\xe2\x32\x46\x11\x85\xe8\x76\xa6\xc9\x9d\xab\x1a\x41\x14\x29\x48\xaf\x0d\x39\x3c\xbf\x3c\xbf\xa4\xd2\xa2\xe8\x2c\xe3\x75\x72\xc8\x97\x7c\xc7\x28\xc7\xbc\x41\x14\xc5\x6a\x02\x62\x4e\xd8\x37\xc5\x6e\xa3\xfa\x65\x06\xab\x91\x53\x8b\xa2\xd0\x5d\x66\xfc\x37\x83\xd1\x7d\x9e\x29\x3a\xa5\xfb\x8a\x98\x13\x9c\x2a\x8c\xbe\x33\xa8\x61\x20\xd3\x56\xf9\x38\xc1\xaa\x16\x09\xb5\x4e\x2e\x7d\x6b\x77\x5e\xfe\x60\xed\xa9\xca\xfb\x90\x5f\xac\x36\x47\xe2\x18\xb7\x2a\x7f\x9a\xb2\xae\xeb\x53\xb7\xa3\x4b\xb2\xb7\x9c\x4b\x8e\x5a\xc4\x3c\x6a\x2d\x3d\x6b\xb3\x4e\x1c\x39\x4f\x9c\xaa\xfe\x90\x45\x32\x71\xfe\x5b\xfb\xea\x3a\xcb\xfd\xb3\xfa\xb1\xd9\xb8\xf9\xc3\x8b\xc6\x28\xfe\x04\x00\x00\xff\xff\x95\x06\x0f\xa4\x50\x03\x00\x00")

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

	info := bindataFileInfo{name: "template/main.tmpl", size: 848, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateModelTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8f\xb1\xae\x83\x30\x0c\x45\x77\x7f\xc5\x55\x94\xf1\x3d\x3e\x00\xa9\x5b\xd5\xaa\x4b\x97\x76\x2f\x11\x98\x88\x02\x01\x11\x3a\x54\x96\xff\xbd\x32\x45\xcd\x76\x72\xe5\x73\x6d\x12\x41\xc3\x6d\x97\x18\x6e\x9c\x1a\x1e\x1c\x54\x89\xe6\x50\xf7\x21\x32\x72\xe4\x44\xb4\xbe\x67\x86\x08\x7c\x71\x0d\x23\x43\x15\x79\x5d\x5e\xf5\x0a\x21\x00\xb8\x1c\x8d\xbb\x14\x37\x12\xf9\xc7\x12\x52\x64\xf8\xc7\x1f\x7c\x8b\xf2\x00\x5f\x9c\x3a\x1e\x9a\x6c\x6e\xec\xcf\x7c\x6d\x71\xdb\x44\x5b\x6a\xde\xef\xe7\x79\xba\x5b\xa3\x2a\xaa\x67\x9e\x52\xe9\x44\x90\x53\xe8\xd9\xc2\x7d\x05\x57\xfd\xda\x38\xd9\x2c\x29\xd9\x31\x3b\x7c\x02\x00\x00\xff\xff\x54\x80\x95\x80\xd8\x00\x00\x00")

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

	info := bindataFileInfo{name: "template/model.tmpl", size: 216, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"schema.go":           schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
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
