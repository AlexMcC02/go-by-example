package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Reading and writing files are basic tasks needed for many Go programs.

// Reading files requires checking most calls for errors. This helper function
// will streamline the process.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// The most basic file reading task involves dumping a file's entire contents
	// into memory.
	path := filepath.Join(os.TempDir(), "dat")
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))

	// You'll often want more control over how what parts of a file are read. For 
	// these tasks, start by Openin a file to obtain an os.File value.
	f, err := os.Open(path)
	check(err)

	// Here we read some bytes from the beginning of the file. Allow up to 5 to be
	// read but also note how many actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// You can also Seek to a known location in the file and Read from there.
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Other methods of seeking are relative to the current cursor position.
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// And relative to the end of the file.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// The io package provides some functions that may be helpful for file reading.
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind, but Seeking with an offset of 0 accomplishes this.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// The bufio package implements a buffered reader that may be useful for its efficiency
	// with many small reads and because of the additional reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Always close a file when you're done (ideally implement this with a defer).
	f.Close()
}