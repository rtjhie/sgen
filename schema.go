package sgen

import (
	"encoding/json"
	"reflect"
)

type (
	Schema struct {
		Name       string       `json:"name"`
		Fields     []*FieldInfo `json:"fields"`
		Subschemas []*Subschema `json:"subschemas,omitempty"`
	}
	Subschema struct {
		Name   string       `json:"name"`
		Fields []*FieldInfo `json:"fields"`
	}
	FieldInfo struct {
		Name      string     `json:"name"`
		Type      Type       `json:"type"`
		Subschema *Subschema `json:"subschema,omitempty"`
	}
)

func MarshalSchema(schema Interface) ([]byte, error) {
	s := (&Parser{}).ParseSchema(schema)
	return json.Marshal(s)
}

func UnmarshalSchema(b []byte) (*Schema, error) {
	m := &Schema{}
	if err := json.Unmarshal(b, m); err != nil {
		return nil, err
	}
	return m, nil
}

type Parser struct {
	schema     *Schema
	subschemas []*Subschema
}

func (p *Parser) ParseSchema(schema Interface) *Schema {
	if schema == nil {
		return nil
	}
	s := &Schema{
		Name: reflect.TypeOf(schema).Name(),
	}
	p.schema = s
	for _, f := range schema.Fields() {
		s.Fields = append(s.Fields, p.ParseField(f.Descriptor()))
	}
	p.schema.Subschemas = p.subschemas
	return s
}

func (p *Parser) ParseSubschema(schema Interface) *Subschema {
	if schema == nil {
		return nil
	}
	s := &Subschema{
		Name: reflect.TypeOf(schema).Name(),
	}
	p.subschemas = append(p.subschemas, s)
	for _, f := range schema.Fields() {
		s.Fields = append(s.Fields, p.ParseField(f.Descriptor()))
	}
	return s
}

func (p *Parser) ParseField(d *Descriptor) *FieldInfo {
	f := &FieldInfo{
		Name:      d.Name,
		Type:      d.Type,
		Subschema: p.ParseSubschema(d.SubSchema),
	}
	return f
}
