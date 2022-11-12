package main

import (
	"fmt"

	"github.com/adhocore/jsonc"
)

func main() {
	v := make(map[string]interface{})
	j := jsonc.New()

	// strip and unmarshal from file: j.UnmarshalFile(file, &v)
	j.UnmarshalFile("./examples/test.json5", &v)
	fmt.Printf("%+v\n---\n", v)

	// strip and unmarshal from byte array: j.Unmarshal(b, &v)
	b := []byte(`{
		// comment
		"a'b": "apple'ball",
		"cat": ["dog",],
	}`)
	v1 := make(map[string]interface{})
	j.Unmarshal(b, &v1)
	fmt.Printf("%+v\n---\n", v1)

	// strip and unmarshal from string: j.Unmarshal(s, &v)
	s := `{
	"a'b": "apple'ball",
	"cat": ["fish"] // comment
	/* comment */` + "\n}"
	v2 := make(map[string]interface{})
	j.Unmarshal([]byte(s), &v2)
	fmt.Printf("%+v\n---\n", v2)

	// strip only from byte array: j.Strip(b)
	fmt.Printf("%+v\n---\n", j.Strip(b))

	// strip only from string: j.StripS(s)
	fmt.Printf("%+v\n---\n", j.StripS(s))
}
