package main

import (
	"fmt"
)

type Section struct {
	SID  int
	Name string
}

func ListSections() []Section {
	// We only allow some sections here. So I am just going to hard code these.
	sections := make([]Section, 0)
	sections = append(sections, Section{6, "General Discussion"})
	sections = append(sections, Section{60, "Fast Threads"})
	sections = append(sections, Section{396, "Sensationalist Headlines"})
	sections = append(sections, Section{110, "General Games Discussion"})
	sections = append(sections, Section{240, "Programming"})
	return sections
}

func GetSectionsString() string {
	output := ""
	Sections := ListSections()
	for k, v := range Sections {
		output += fmt.Sprintf("Press %d for %s... ", k, v.Name)
	}
	return output
}
