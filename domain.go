package sgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type (
	Config struct{}
	Option struct{}
)

// Generate writes Go source code for models, json files for JSON schemas, and embeds the json files as binary data.
func Generate(schemaPath string, cfg *Config, options ...Option) (err error) {
	// prepare env
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
	err = ioutil.WriteFile("domain.json", dst.Bytes(), 0644)
	if err != nil {
		return err
	}
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
	_, err = bindata(".")
	if err != nil {
		return err
	}
	// go code
	return nil
}

func loadDomain(schemaPath string) (*Domain, error) {
	// TODO: generate main.go from main.tmpl then run it
	out, err := run("./intermediate/main.go")
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
	fmt.Println(target)
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

func bindata(path string) (string, error) {
	o := fmt.Sprintf("-o=%s/internal/bindata/bindata.go", path)
	dir := fmt.Sprintf("%s/json/...", path)
	cmd := exec.Command("go", "run", "github.com/go-bindata/go-bindata/go-bindata", o, "-pkg=bindata", dir)
	fmt.Println(cmd.Args)
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("bindata: %s", stderr)
	}
	return stdout.String(), nil
}
