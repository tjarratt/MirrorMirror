package main

import (
	"fmt"
	"go/ast"
	"strings"
)

type Interface struct {
	name string
	methods []Method
}

func createInterfaceFromNameAndNode(name string, anInterface *ast.InterfaceType) (i Interface) {
	i.name = name

	for _, m := range anInterface.Methods.List {
		method := newMethod(m)
		i.methods = append(i.methods, method)
	}

	return
}

func (i Interface) fakeStructDeclaration() string {
	returnsSlice := []string{}
	receivedSlice := []string{}

	for _, method := range i.methods {
		for index, p := range method.ParamSlice {
			var concatName string
			if p.name == "" {
				concatName = method.Name + "Param" + fmt.Sprintf("%d", index + 1)
			} else {
				concatName = method.Name + p.UpperName()
			}

			receivedSlice = append(receivedSlice, concatName + " " + p.paramType)
		}

		for index, r := range method.ReturnSlice {
			var concatName string
			if r.name == "" {
				concatName = method.Name + "Return" + fmt.Sprintf("%d", index + 1)
			} else {
				concatName = method.Name + r.UpperName()
			}

			returnsSlice = append(returnsSlice, concatName + " " + r.returnType)
		}
	}

	return fmt.Sprintf("type Fake%s struct {\n\t%s\n\t%s\n}\n",
		i.name,
		fmt.Sprintf("Returns struct {\n\t\t%s\n\t}\n", strings.Join(returnsSlice, "\n\t\t")),
		fmt.Sprintf("Received struct {\n\t\t%s\n\t}\n", strings.Join(receivedSlice, "\n\t\t")),
	)
}

func (i Interface) StubbedMethods() string {
	var results []string
	formatString := "func (fake *%s) %s(%s) (%s) {\n%s\n}"

	for _, m := range i.methods {
		params := m.ParamSlice.String()
		returns := m.ReturnSlice.String()
		body := "\treturn"
		results = append(results, fmt.Sprintf(formatString, i.name, m.Name, params, returns, body))
	}

	return strings.Join(results, "\n\n")
}
