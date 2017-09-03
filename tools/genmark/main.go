package main

import (
  "github.com/BurntSushi/toml"
	"os"
	"fmt"
	"text/template"
)

// Item describes each places or events.
type Item struct {
	Description string
}

type Items map[string]Item

type TmplData struct {
	Title string
	Items Items
}

const tmplText = `# {{.Title}}
{{range $name, $info := .Items}}
## {{$name}}
{{$info.Description}}
{{end}}
`

func main() {
	primitives := map[string]toml.Primitive{}

	meta, err := toml.DecodeReader(os.Stdin, &primitives)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var category string
	meta.PrimitiveDecode(primitives["category"], &category)
	delete(primitives, "category")

	items := Items{}
	for name, prim := range primitives {
		item := Item{}
		err = meta.PrimitiveDecode(prim, &item)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		items[name] = item
	}

	data := TmplData{
		Title: category,
		Items: items,
	}

	tmpl, err := template.New("markdown").Parse(tmplText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
