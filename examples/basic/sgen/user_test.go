package sgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/alecthomas/jsonschema"
)

func TestUser_JSONSchema(t *testing.T) {
	t.Run("JSONSchema", func(t *testing.T) {
		u := User{}
		schema := &jsonschema.Schema{}
		json.Unmarshal(u.JSONSchema(), schema)
		b, _ := json.Marshal(schema)
		out := bytes.Buffer{}
		json.Indent(&out, b, "", "  ")
		fmt.Println(out.String())
	})
}
