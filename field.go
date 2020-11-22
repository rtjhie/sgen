package sgen

type Type uint8

const (
	TypeString Type = iota
	TypeObject
	TypeArray
	endTypes
)

var (
	typeNames = [...]string{
		TypeString: "string",
		TypeObject: "object",
		TypeArray:  "array",
	}
)

func (t Type) String() string {
	if t >= endTypes {
		return ""
	}
	return typeNames[t]
}

type Descriptor struct {
	Name      string
	Type      Type
	SubSchema Interface
}

func String(name string) *stringBuilder {
	return &stringBuilder{&Descriptor{
		Name: name,
		Type: TypeString,
	}}
}

type stringBuilder struct {
	desc *Descriptor
}

func (b *stringBuilder) Descriptor() *Descriptor {
	return b.desc
}

func Object(name string) *objectBuilder {
	return &objectBuilder{&Descriptor{
		Name: name,
		Type: TypeObject,
	}}
}

type objectBuilder struct {
	desc *Descriptor
}

func (b *objectBuilder) Descriptor() *Descriptor {
	return b.desc
}

func (b *objectBuilder) Schema(s Interface) *objectBuilder {
	b.desc.SubSchema = s
	return b
}
