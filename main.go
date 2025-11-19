package main

import (
    "fmt"
    "os"
    "path/filepath"
)

// Throughout program execution, we often want to create data that isn't needed
// after the program exits. Temporary files are directories are useful for this
// purpose since they do not pollute the file system over time.

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Temporary files can be created with os.CreateTemp as illustrated below.
    // The first arugment is the location, so "" is the current directory.
    f, err := os.CreateTemp("", "sample")
    check(err)

    // Temporary files are assigned a name based on the second argument provided
    // for os.CreateTemp and the rest is chosen automatically to ensure that
    // concurrent calls will always create different file names.
    fmt.Println("Temp file name:", f.Name())

    // Cleaning up the temporary file once we're done with it.
    defer os.Remove(f.Name())

    // Writing some data to the temporary file.
    _, err = f.Write([]byte{1,2,3,4})
    check(err)

    // We can also create a temporary directory.
    dname, err := os.MkdirTemp("", "sampledir")
    check(err)
    fmt.Println("Temp dir name:", dname)

    defer os.RemoveAll(dname)

    // Now we can synthesize temporary file names by prefixing them with the
    // temporary directory.
    fname := filepath.Join(dname, "file")
    err = os.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}