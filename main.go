package main

import (
	"os"
	"text/template"
)

// Go offers built-in support for creating dynamic content or showing customised output to the user
// with the text/tempalte package.

func main() {

	// Here a new template is created and its body is parsed from a string.
	// Templates consist of a mix of a static text and actions enclosed {{...}}
	// that can be used to dynamically insert content.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// The template.Must function can be used to panic in case Parse returns an error.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// By executing the template we generatge its text with specific values for its actions.
	// Here the {{.}} action is replaced by the value passed as a parameter to Execute.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string {
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Helper function that will be used later.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	
	// If the data is a struct, the {{.FieldName}} action is used to access it fields.
	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// A similar approach is used for maps.
	t2.Execute(os.Stdout, map[string]string {
		"Name": "Mickey Mouse",
	})

	// If/else provide conditional execution for templates.
	// A value is considered 'false' if it's the default value for a type.
	t3 := Create("t3",
        "{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty")
    t3.Execute(os.Stdout, "")

	// Range blocks can be used for looping.
	t4 := Create("t4",
        "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout,
        []string{
            "Go",
            "Rust",
            "C++",
            "C#",
        })
}