package sgen

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

type (
	Config struct{}
	Option struct{}
)

// Generate writes Go source code for models, json files for JSON schemas, and embeds the json files as binary data.
func Generate(schemaPath string, cfg *Config, options ...Option) (err error) {
	domain, err := loadDomain(schemaPath)
	if err != nil {
		return err
	}
	if err := domain.generate(); err != nil {
		return err
	}

	return nil
}

type Domain struct {
	Schemas []*Schema `json:"models,omitempty"`
}

// generate generates the files for the domain.
func (d *Domain) generate() error {
	b, err := json.Marshal(d.Schemas)
	if err != nil {
		return err
	}
	var dst bytes.Buffer
	err = json.Indent(&dst, b, "", "  ")
	if err != nil {
		return err
	}
	// err = ioutil.WriteFile("domain.json", dst.Bytes(), 0644)
	// if err != nil {
	// 	return err
	// }
	for _, s := range d.Schemas {
		j := NewJSONSchema(s)
		jsonb, err := json.Marshal(j)
		if err != nil {
			return err
		}
		var dst bytes.Buffer
		err = json.Indent(&dst, jsonb, "", "    ")
		err = ioutil.WriteFile("json/"+snake(s.Name)+".json", dst.Bytes(), 0644)
		if err != nil {
			return err
		}
	}
	_, err = bindata("internal/binjson/binjson.go", "binjson", "./json")
	if err != nil {
		return err
	}

	w := bytes.NewBuffer(nil)
	for _, s := range d.Schemas {
		err = templates.ExecuteTemplate(w, "model", s)
		if err != nil {
			return err
		}
		path := snake(s.Name) + ".go"
		err = ioutil.WriteFile(path, w.Bytes(), 0644)
		if err != nil {
			return err
		}
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read file %s: %v", path, err)
		}
		src, err := imports.Process(path, buf, nil)
		if err != nil {
			return fmt.Errorf("format file %s: %v", path, err)
		}
		if err := ioutil.WriteFile(path, src, 0644); err != nil {
			return fmt.Errorf("write file %s: %v", path, err)
		}
	}

	return nil
}

func loadDomain(schemaPath string) (*Domain, error) {
	// TODO: generate main.go from main.tmpl then run it
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadSyntax}, schemaPath)
	if err != nil {
		return nil, err
	}
	if len(pkgs) != 1 {
		return nil, errors.New("cannot resolve package from schema path")
	}
	pkgPath := pkgs[0].PkgPath
	filePath, err := filepath.Abs(schemaPath)
	if err != nil {
		return nil, err
	}
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, f := range files {
		names = append(names, pascal(strings.Trim(f.Name(), ".go")))
	}
	b := bytes.NewBuffer(nil)
	err = templates.ExecuteTemplate(b, "main", struct {
		Pkg   string
		Names []string
	}{pkgPath, names})
	if err != nil {
		return nil, fmt.Errorf("execute template: %v", err)
	}
	buf, err := format.Source(b.Bytes())
	if err != nil {
		return nil, fmt.Errorf("format template: %v", err)
	}
	if err := os.MkdirAll(".sgen", os.ModePerm); err != nil {
		return nil, err
	}
	target := fmt.Sprintf(".sgen/%s.go", filename(filePath))
	if err := ioutil.WriteFile(target, buf, 0644); err != nil {
		return nil, fmt.Errorf("write file %s: %v", target, err)
	}
	defer os.RemoveAll(".sgen")
	out, err := run(target)
	if err != nil {
		return nil, err
	}
	d := &Domain{}
	for _, line := range strings.Split(out, "\n") {
		s, err := UnmarshalSchema([]byte(line))
		if err != nil {
			return nil, err
		}
		d.Schemas = append(d.Schemas, s)
	}
	return d, nil
}

func run(target string) (string, error) {
	cmd := exec.Command("go", "run", target)
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s", stderr)
	}
	return stdout.String(), nil
}

func bindata(o, pkg, dir string) (string, error) {
	cmd := exec.Command("go", "run", "github.com/go-bindata/go-bindata/go-bindata", "-o="+o, "-pkg="+pkg, dir)
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("bindata: %s", stderr)
	}
	return stdout.String(), nil
}

func filename(pkg string) string {
	name := strings.Replace(pkg, "/", "_", -1)
	return fmt.Sprintf("entc_%s_%d", name, time.Now().Unix())
}
