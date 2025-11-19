package main

import "embed"

// 'go:embed' is a compiler directive allowing programs to include
// arbitrary files and folders in the Go binary at build time.

// Embed directives accept paths relative to the directory containing
// the Go source file. The directive embeds the contents of the file 
// the string variable immediately following it.

//go:embed folder/single_file.txt
var fileString string
	
//go:embed folder/single_file.txt
var fileByte []byte

// You can embed multiple files (or folders) with wildcards.

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
    print(fileString)
    print(string(fileByte))
    
    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))

    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))
}