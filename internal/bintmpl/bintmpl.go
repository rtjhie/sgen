// Code generated for package bintmpl by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.go
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

var _templateModelTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x90\x41\x4b\xc3\x40\x10\x85\xef\xfb\x2b\x1e\x4b\x0e\x09\x68\xbc\x17\x3c\x08\x22\x28\x58\x0f\x3d\x8a\xd8\xed\xee\x24\xae\x4d\xa6\x21\xbb\x39\x94\x61\xff\xbb\x0c\x44\x9a\xe3\x1b\x66\xbe\xef\x31\x46\x04\x81\xba\xc8\x04\x3b\x5e\x02\x0d\x16\xa5\x18\x33\x39\x7f\x76\x3d\x21\xf5\xc4\xc6\xe4\xeb\x44\x10\x41\xd5\xee\xdd\x48\x28\x05\x29\xcf\x8b\xcf\x10\x03\x00\xaf\xcf\x9a\x23\xf7\x38\xfe\xa6\x0b\xef\x6c\x0c\xf6\x68\x44\xee\x31\x3b\xee\x09\xd5\xf7\x1d\xaa\x0e\xbb\x47\x54\xed\x4b\xa4\x21\x24\x75\xe8\xa5\x08\x26\x97\xbc\x1b\x50\x75\x5b\xf6\x86\x25\x82\xc4\xee\x4c\x9b\x8d\x15\x4e\x1c\x94\x53\x8c\xe9\x16\xf6\xa8\xb7\x05\x1b\xbc\x1d\x3e\xf6\x07\xff\x43\xa3\xab\x1b\x7c\x7e\x9d\xae\x99\xd6\xba\x33\xe5\x65\x66\x9c\x22\xab\xa1\x7d\x5f\x52\x7e\x4a\x89\x72\x6d\x35\x3f\xdc\x84\xff\xb0\x56\xe7\xb6\x51\x93\xbe\x6b\xf5\xfe\x05\x00\x00\xff\xff\xe7\xd5\xac\x42\x3a\x01\x00\x00")

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

	info := bindataFileInfo{name: "template/model.tmpl", size: 314, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
