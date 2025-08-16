// Package twofer contains a function about people sharing cookies
package twofer

import "fmt"

// ShareWith gets an optional name as input and returns a string that
// says who are you sharing a cookie with, if you know the name
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }
	return fmt.Sprintf("One for %s, one for me.", name)
}
