package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

// The filepath prackage provides functions to parse and construct file paths
// in a way that is portable between OSes.

func main() {

	// Join should be used to construct paths is a portable way. It takes any
	// number of arguments and constructs a hierarchial path from them.
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

	// You should always use Join instead of concatenating /s or \s manually. In
	// addition to the portability benefits, Join will normalise paths by removing
	// superfluous separators and directly changes.
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir and Base can be used split a path to the directory and the file. Alternatively
	//, Split will return will both in the same call.
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

	// Here we can check if a filepath is absolute.
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

	// Some file names have extensions following a dot. We can split the extension out of
	// such names with Ext.
    ext := filepath.Ext(filename)
    fmt.Println(ext)

	// To find the file's name with the extension removed, use strings.TrimSuffix.
    fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and a target. It returns an error if the
	// target cannot be made relative to base.
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}