package sgen

import (
	"encoding/json"

	"github.com/iancoleman/orderedmap"
)

func NewJSONSchema(s *Schema) *JSONSchema {
	j := &JSONSchema{
		Definition: &Definition{
			Version:              Version,
			Type:                 "object",
			Properties:           orderedmap.New(),
			AdditionalProperties: []byte("false"),
		},
	}
	for _, f := range s.Fields {
		property := GetProperty(f)
		j.Properties.Set(f.Name, property)
	}
	return j
}

func GetProperty(f *FieldInfo) *Definition {
	ret := &Definition{
		Type: f.Type.String(),
	}
	if f.Type == TypeObject {
		ret.Properties = orderedmap.New()
		for _, ff := range f.Subschema.Fields {
			property := GetProperty(ff)
			ret.Properties.Set(ff.Name, property)
		}
	}
	return ret
}

// Version is the JSON Schema version.
// If extending JSON Schema with custom values use a custom URI.
// RFC draft-wright-json-schema-00, section 6
var Version = "http://json-schema.org/draft-04/schema#"

// JSONSchema is the root schema.
// RFC draft-wright-json-schema-00, section 4.5
type JSONSchema struct {
	*Definition
	Definitions Definitions
}

// Definition represents a JSON Schema object definition.
type Definition struct {
	// RFC draft-wright-json-schema-00
	Version string `json:"$schema,omitempty"` // section 6.1
	Ref     string `json:"$ref,omitempty"`    // section 7
	// RFC draft-wright-json-schema-validation-00, section 5
	MultipleOf           int                    `json:"multipleOf,omitempty"`           // section 5.1
	Maximum              int                    `json:"maximum,omitempty"`              // section 5.2
	ExclusiveMaximum     bool                   `json:"exclusiveMaximum,omitempty"`     // section 5.3
	Minimum              int                    `json:"minimum,omitempty"`              // section 5.4
	ExclusiveMinimum     bool                   `json:"exclusiveMinimum,omitempty"`     // section 5.5
	MaxLength            int                    `json:"maxLength,omitempty"`            // section 5.6
	MinLength            int                    `json:"minLength,omitempty"`            // section 5.7
	Pattern              string                 `json:"pattern,omitempty"`              // section 5.8
	AdditionalItems      *Definition            `json:"additionalItems,omitempty"`      // section 5.9
	Items                *Definition            `json:"items,omitempty"`                // section 5.9
	MaxItems             int                    `json:"maxItems,omitempty"`             // section 5.10
	MinItems             int                    `json:"minItems,omitempty"`             // section 5.11
	UniqueItems          bool                   `json:"uniqueItems,omitempty"`          // section 5.12
	MaxProperties        int                    `json:"maxProperties,omitempty"`        // section 5.13
	MinProperties        int                    `json:"minProperties,omitempty"`        // section 5.14
	Required             []string               `json:"required,omitempty"`             // section 5.15
	Properties           *orderedmap.OrderedMap `json:"properties,omitempty"`           // section 5.16
	PatternProperties    map[string]*Definition `json:"patternProperties,omitempty"`    // section 5.17
	AdditionalProperties json.RawMessage        `json:"additionalProperties,omitempty"` // section 5.18
	Dependencies         map[string]*Definition `json:"dependencies,omitempty"`         // section 5.19
	Enum                 []interface{}          `json:"enum,omitempty"`                 // section 5.20
	Type                 string                 `json:"type,omitempty"`                 // section 5.21
	AllOf                []*Definition          `json:"allOf,omitempty"`                // section 5.22
	AnyOf                []*Definition          `json:"anyOf,omitempty"`                // section 5.23
	OneOf                []*Definition          `json:"oneOf,omitempty"`                // section 5.24
	Not                  *Definition            `json:"not,omitempty"`                  // section 5.25
	Definitions          Definitions            `json:"definitions,omitempty"`          // section 5.26
	// RFC draft-wright-json-schema-validation-00, section 6, 7
	Title       string        `json:"title,omitempty"`       // section 6.1
	Description string        `json:"description,omitempty"` // section 6.1
	Default     interface{}   `json:"default,omitempty"`     // section 6.2
	Format      string        `json:"format,omitempty"`      // section 7
	Examples    []interface{} `json:"examples,omitempty"`    // section 7.4
	// RFC draft-wright-json-schema-hyperschema-00, section 4
	Media          *Definition `json:"media,omitempty"`          // section 4.3
	BinaryEncoding string      `json:"binaryEncoding,omitempty"` // section 4.3

	Extras map[string]interface{} `json:"-"`
}

// Definitions hold schema definitions.
// http://json-schema.org/latest/json-schema-validation.html#rfc.section.5.26
// RFC draft-wright-json-schema-validation-00, section 5.26
type Definitions map[string]*Definition

func (s *JSONSchema) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(s.Definition)
	if err != nil {
		return nil, err
	}
	if s.Definitions == nil || len(s.Definitions) == 0 {
		return b, nil
	}
	d, err := json.Marshal(struct {
		Definitions Definitions `json:"definitions,omitempty"`
	}{s.Definitions})
	if err != nil {
		return nil, err
	}
	if len(b) == 2 {
		return d, nil
	} else {
		b[len(b)-1] = ','
		return append(b, d[1:]...), nil
	}
}

func (t *Definition) MarshalJSON() ([]byte, error) {
	type Type_ Definition
	b, err := json.Marshal((*Type_)(t))
	if err != nil {
		return nil, err
	}
	if t.Extras == nil || len(t.Extras) == 0 {
		return b, nil
	}
	m, err := json.Marshal(t.Extras)
	if err != nil {
		return nil, err
	}
	if len(b) == 2 {
		return m, nil
	} else {
		b[len(b)-1] = ','
		return append(b, m[1:]...), nil
	}
}
