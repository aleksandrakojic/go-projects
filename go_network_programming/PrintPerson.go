package main

import (
    "fmt"
    "text/template"
    "os"
)

type Person struct {
    Name string
    Age  int
	Email []string
	Jobs  []*Job
}

type Job struct {
    Employer string
    Role string
}

const templ = `The name is {{ .Name }}
	The age is {{.Age}}.
	{{range .Email}}
		An email is {{.}}
	{{end}}

	{{with .Jobs}}
		{{range .}}
			An employer is {{.Employer}}
			and the role is {{.Role}}
		{{end}}
	{{end}}
`

func main() {
	job1 := Job{Employer: "Box Hill Institute", Role: "Director, Commerce and ICT"}
	job2 := Job{Employer: "Canberra University", Role: "Adjunct Professor"}

	person := Person{
		Name: "jan",
		Age: 66,
		Email: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
		Jobs: []*Job{&job1, &job2},
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}