package hashstructure

import (
	"encoding/hex"
	"fmt"
)

func ExampleHash() {
	type ComplexStruct struct {
		Name     string
		Age      uint
		Metadata map[string]interface{}
	}

	v := ComplexStruct{
		Name: "mitchellh",
		Age:  64,
		Metadata: map[string]interface{}{
			"car":      true,
			"location": "California",
			"siblings": []string{"Bob", "John"},
		},
	}

	hash, err := Hash(v, FormatV2, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(hash))
	// Output:
	// 3bd3d1d3941c96a9c4cbdfac7c132c443db354e0c34bc22540b8a9969211d625
}

func ExampleHash_v1() {
	type ComplexStruct struct {
		Name     string
		Age      uint
		Metadata map[string]interface{}
	}

	v := ComplexStruct{
		Name: "mitchellh",
		Age:  64,
		Metadata: map[string]interface{}{
			"car":      true,
			"location": "California",
			"siblings": []string{"Bob", "John"},
		},
	}

	hash, err := Hash(v, FormatV1, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(hash))
	// Output:
	// 62aace8c92f6ea9c4699a62da03c08c1458ebbf0331210a8a0e6b43536c59c68
}
