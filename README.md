# sgen
Describe your schemas using Go to generate your Go structs and JSON schemas. Heavily inspired by `github.com/facebook/ent` and `github.com/alecthomas/jsonschema`.

## Getting Started

Follow along to generate your very first model.

Get sgen!
```
go get github.com/rtjhie/sgen/cmd/sgen
```

Go to your working directory and run:
```
sgen init User
```

This creates a directory structure like:
```
sgen/
|-- json/
|-- schema/
|   |-- user.go
|-- generate.go
```

Where `user.go` will look like:
```go
package schema

import "github.com/rtjhie/sgen"

type User struct{}

func (User) Fields() []sgen.Field {
  return []sgen.Field{
    // ** Define fields here **
  }
}
```

Next you'll want to populate your User schema with some fields:
```go
func (User) Fields() []sgen.Field {
  return []sgen.Field{
    sgen.String("first_name"),
    sgen.String("last_name"),
    sgen.String("email"),
  }
}
```

Then, when you call `go generate $PROJECT_ROOT/sgen` you'll get a couple neat things:

**Go struct**: `.sgen/user.go`
```go
package sgen

import "example.com/example/sgen/internal/binjson"

type User struct {
  ID        string `json:"_id"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Email     string `json:"email"`
}

func (User) JSONSchema() []byte {
  return binjson.MustAsset("json/user.json")
}

```

**JSON schema**: `./sgen/json/user.json`
```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "_id": {
      "type": "string"
    },
    "first_name": {
      "type": "string"
    },
    "last_name": {
      "type": "string"
    },
    "email": {
      "type": "string"
    }
  },
  "additionalProperties": false,
  "type": "object"
}
```

Now just import your new User type into your packages!
