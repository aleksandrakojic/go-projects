package main

import "fmt"

type Folder struct {
	component []Component
	name     string
}

func (f *Folder) search(keyword string) {
    fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.component {
        composite.search(keyword)
    }
}

func (f *Folder) add(c Component) {
	f.component = append(f.component, c)
}