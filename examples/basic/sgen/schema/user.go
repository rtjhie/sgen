package schema

import (
	"github.com/rtjhie/sgen"
)

// User holds the schema definition for the User entity.
type User struct{}

// Fields of the User.
func (User) Fields() []sgen.Field {
	return []sgen.Field{
		sgen.String("name"),
		sgen.Object("address").
			Schema(Address{}),
	}
}

type Address struct{}

func (Address) Fields() []sgen.Field {
	return []sgen.Field{
		sgen.String("street"),
		sgen.String("city"),
		sgen.String("province"),
		sgen.String("postalCode"),
	}
}
