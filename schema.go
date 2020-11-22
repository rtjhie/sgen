package sgen

import (
	"encoding/json"
	"reflect"
)

type (
	Schema struct {
		Name   string       `json:"name"`
		Fields []*FieldInfo `json:"fields"`
	}
	FieldInfo struct {
		Name      string  `json:"name"`
		Type      Type    `json:"type"`
		Subschema *Schema `json:"subschema,omitempty"`
	}
)

func MarshalSchema(schema Interface) ([]byte, error) {
	s := ParseSchema(schema)
	return json.Marshal(s)
}

func UnmarshalSchema(b []byte) (*Schema, error) {
	m := &Schema{}
	if err := json.Unmarshal(b, m); err != nil {
		return nil, err
	}
	return m, nil
}

func ParseSchema(schema Interface) *Schema {
	if schema == nil {
		return nil
	}
	s := Schema{
		Name: reflect.TypeOf(schema).Name(),
	}
	for _, f := range schema.Fields() {
		s.Fields = append(s.Fields, ParseField(f.Descriptor()))
	}
	return &s
}

func ParseField(d *Descriptor) *FieldInfo {
	return &FieldInfo{
		Name:      d.Name,
		Type:      d.Type,
		Subschema: ParseSchema(d.SubSchema),
	}
}
