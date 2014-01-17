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
	var fakeStructDeclaration string

	for _, method := range i.methods {
		params := method.ParamSlice.String()
		returns := method.ReturnSlice.String()
		fakeStructDeclaration = fmt.Sprintf("%s\t%s (%s) (%s)\n", fakeStructDeclaration, method.Name, params, returns)
	}

	return fmt.Sprintf("type Fake%s struct {\n%s}", i.name, fakeStructDeclaration)
}

func (i Interface) StubbedMethods() string {
	var results []string
	formatString := "func %s (fake %s) (%s) (%s) {\n%s\n}"

	for _, m := range i.methods {
		params := m.ParamSlice.String()
		returns := m.ReturnSlice.String()
		body := ""
		results = append(results, fmt.Sprintf(formatString, m.Name, i.name, params, returns, body))
	}

	return strings.Join(results, "\n\n")
}
