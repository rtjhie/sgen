package main

//TODO: This should be a generated file.

import (
	"bytes"
	"os"

	"github.com/rtjhie/sgen"
	"github.com/rtjhie/sgen/examples/basic/sgen/schema"
)

// Collect schemas based on first type in file?
var schemas = []sgen.Interface{
	schema.User{},
}

func main() {
	var lines [][]byte
	for _, schema := range schemas {
		b, err := sgen.MarshalSchema(schema)
		if err != nil {
			fail(err)
		}
		lines = append(lines, b)
	}
	os.Stdout.Write(bytes.Join(lines, []byte("\n")))
}

func fail(err error) {
	os.Stderr.WriteString(err.Error() + "\n")
	os.Exit(1)
}
