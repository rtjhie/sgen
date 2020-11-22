package sgen

type (
	Interface interface {
		Fields() []Field
	}

	Field interface {
		Descriptor() *Descriptor
	}
)
