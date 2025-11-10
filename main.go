package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

// Go offers built-in support for JSON encoding and decoding, including to
// and from built-in and custom data types.

// We'll use these two structs to demonstrate encoding and decoding of custom
// types below.
type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

	// Below are some examples of encoding basic data types to JSON.
	// Note: fields must start with capital letters to be exported.
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

	// Here are some examples for slices are maps, which encode to
	// JSON arrays and objects as you'd expect.
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

	// The JSON package can automatically encode your custom data types.
	// It will only include exported fields in the encoded output and will
	// by default use those names as the JSON keys.
    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

	// You can use tags on struct field declarations to custgomize the encoded
	// JSON key names.
    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

	// Here is an example of a generic data structure.
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Here is the variable that will be used to store the decoded
	// JSON data. This syntax will hold a map of strings to arbitray data types.
    var dat map[string]interface{}

	// The decoding proper, with an appropriate error check.
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

	// Type conversion.
    num := dat["num"].(float64)
    fmt.Println(num)

	// Accessing nested data requires a series of conversions.
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

	// JSON can be decoded into custom data types.
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

	// Whereas in previous examples bytes and stgrings were used as intermediates
	// between the data and JSON representation on standard out. We can also stream
	// JSON encodings directly to os.Writers like os.Stdout or even HTTP response
	// bodies.
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)

	// Streaming reads from os.Readers like os.Stdin or HTTP request bodies is done with
	// json.Decoder.
    dec := json.NewDecoder(strings.NewReader(str))
    res1 := response2{}
    dec.Decode(&res1)
    fmt.Println(res1)
}