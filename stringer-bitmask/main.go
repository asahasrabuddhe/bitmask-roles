package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"path"
)

var tpl = `
package {{ .Package }}

import "strings"

var rolesMap = map[string]{{.Type}} {
	{{ range .Constants -}}
	"{{ . }}": {{ . }},
	{{ end }}
}

func (r *{{ .Type }}) String() string {
	var roles []string

	for key, value := range rolesMap {
		if r.IsRole(value) {
			roles = append(roles, key)
		}
	}

	return strings.Join(roles, ", ")
}`

type Data struct {
	Package   string
	Type      string
	Constants []string
}

func main() {
	var fileName, typeName string

	if len(os.Args) == 3 {
		fileName = os.Args[1]
		typeName = os.Args[2]
	} else {
		log.Fatal("usage: stringer-bitmask <filename> <type>")
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, path.Join(dir, fileName), nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	data := &Data{Type: typeName}

	ast.Inspect(file, func(n ast.Node) bool {
		if val, ok := n.(*ast.File); ok {
			data.Package = val.Name.Name
		}
		if val, ok := n.(*ast.GenDecl); ok {
			if val.Tok == token.CONST {
				for _, spec := range val.Specs {
					if val, ok := spec.(*ast.ValueSpec); ok {
						data.Constants = append(data.Constants, val.Names[0].Name)
					}
				}
			}
		}
		return true
	})

	var buf bytes.Buffer

	t := template.Must(template.New("").Parse(tpl))
	err = t.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(fmt.Sprintf("%s_string.go", typeName))
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	err = outFile.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}

	_, err = outFile.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = outFile.Write(p)
}
