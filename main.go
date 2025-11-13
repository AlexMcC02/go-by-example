package main

import (
	"fmt"
	"net"
	"net/url"
)

// URLs provide a uniform way to locate resources.

func main() {
	
	// We'll use this example URL for parsing.
	// It includes a scheme, auth info, host, port, path, query params and fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parsing the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// A simple dot syntax allows us to access parts of the scheme.
	fmt.Println(u.Scheme)

	// User contains all authentication info, here we call Username
	// and Password on this for individual values.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// The Host contains both the hostname and the port, if present
	// Use SplitHostPort to extract them.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Here we extract the path and the fragment after the #.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// To get query params in the string of k=v format, use RawQuery.
	// You can also parse query params into a map. The parsed query
	// maps are from string to slices of strings, so index into [0]
	// if you want only the first value.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}