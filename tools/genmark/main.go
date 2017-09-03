package main

import (
  "github.com/BurntSushi/toml"
	"os"
	"fmt"
	"text/template"
	"path/filepath"
)

// Restaurant describes a restaurant.
type Restaurant struct {
	Description string
}

const tmplText = `# Restaurants
{{range $name, $info := .}}
# {{$name}}
{{$info.Description}}
{{end}}
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "fatal: repository base dir is needed as argument")
		os.Exit(1)
	}
	repoPath := os.Args[1]

	restaurantsPath := filepath.Join(repoPath, "src/restaurants.toml")
	res := map[string]Restaurant{}
	fin, err := os.Open(restaurantsPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputPath := filepath.Join(repoPath, "target/restaurants.md")
	fout, err := os.Create(outputPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = toml.DecodeReader(fin, &res)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tmpl, err := template.New("markdown").Parse(tmplText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = tmpl.Execute(fout, res)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
