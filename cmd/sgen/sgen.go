package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rtjhie/sgen"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("no args")
	}
	cmd := os.Args[1]
	switch cmd {
	case "init":
		var names []string
		if len(os.Args) > 2 {
			names = os.Args[2:]
		}
		return initEnv("sgen", names)
	case "generate":
		path := os.Args[2]
		return sgen.Generate(path, nil)
	}
	return nil
}

var tmpl = template.Must(template.New("schema").
	Parse(`package schema

import "github.com/rtjhie/sgen"

type {{ . }} struct{}

func ({{ . }}) Fields() []sgen.Field {
	return []sgen.Field{
		// ** Define fields here **
	}
}

`))

func initEnv(target string, names []string) error {
	if err := createDir(target); err != nil {
		return err
	}
	for _, name := range names {
		b := bytes.NewBuffer(nil)
		if err := tmpl.Execute(b, name); err != nil {
			log.Fatalln(err)
		}
		path := filepath.Join(target+"/schema", strings.ToLower(name+".go"))
		if err := ioutil.WriteFile(path, b.Bytes(), 0644); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func createDir(target string) error {
	_, err := os.Stat(target)
	if err == nil || !os.IsNotExist(err) {
		return err
	}
	if err := os.MkdirAll(target+"/json", os.ModePerm); err != nil {
		return fmt.Errorf("creating json directory: %w", err)
	}
	if err := os.MkdirAll(target+"/schema", os.ModePerm); err != nil {
		return fmt.Errorf("creating schema directory: %w", err)
	}
	if err := ioutil.WriteFile("sgen/generate.go", []byte("package sgen\n\n//go:generate go run github.com/rtjhie/sgen/cmd/sgen generate ./schema\n"), 0644); err != nil {
		return fmt.Errorf("creating generate.go file: %w", err)
	}
	return nil
}
