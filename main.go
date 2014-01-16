package main

import (
	"os"
	"fmt"
	"go/ast"
	"go/token"
	"go/parser"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}

	sprinkleSugarOn(os.Args[1], os.Args[2])
}

func usage() {
	println("Error: Not enough args. Expected path to file and interface to mock")
	println("usage: sugar /path/to/some/file InterfaceToMock")
}

func sprinkleSugarOn(pathToFile, interfaceToMock string) {
	if _, err := os.Stat(pathToFile); err != nil {
		dir, _ := os.Getwd()
		fmt.Printf("Couldn't find file from dir %s\n", dir)
		fmt.Printf("Error: given file '%s' does not exist\n", pathToFile)
		return
	}

	fileSet := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fileSet, pathToFile, nil, 0)
	if err != nil {
		fmt.Printf("Error parsing '%s':\n%s\n", pathToFile, err)
		return
	}

	packageName := ""
	expectingNamedInterface := false
	var interfaceNode *ast.InterfaceType

	ast.Inspect(parsedFile, func(n ast.Node) bool {
		if n == nil {
			return true
		}

	 	switch n := n.(type) {
	 	case *ast.Ident:
	 		if n.Obj == nil && packageName == "" {
	 			packageName = n.Name
	 		} else if n.Name == interfaceToMock {
	 			expectingNamedInterface = true
	 		}
		case *ast.InterfaceType:
			if expectingNamedInterface {
				interfaceNode = n
			}
	 	default:
	 		expectingNamedInterface = false
		}

		return true
	})

	fmt.Printf("\nmocking out %s.%s Interface\n", packageName, interfaceToMock)

	outputFileName := filepath.Base(pathToFile)
	outputDir := filepath.Dir(pathToFile)
	outputFilePath := filepath.Join(outputDir, "fake_" + outputFileName)

	fmt.Printf("Writing a mock interface to %s\n\n", outputFilePath)

	fakeMock := DeclareFakeImplementingInterface(interfaceToMock, interfaceNode)
	fmt.Printf("fake implementation looks like:\n**********\n%s\n**********\n", fakeMock)
}

func DeclareFakeImplementingInterface(name string, anInterface *ast.InterfaceType) string {
	methods := []Method{}
	var fakeStructDeclaration string

	for _, m := range anInterface.Methods.List {
		method := newMethod(m)
		methods = append(methods, method)

		params := method.ParamSlice.String()
		returns := method.ReturnSlice.String()
		fakeStructDeclaration = fmt.Sprintf("%s\t%s (%s) (%s)\n", fakeStructDeclaration, method.Name, params, returns)
	}


	return fmt.Sprintf("type Fake%s struct {\n%s}", name, fakeStructDeclaration)
}

func StubbedMethodsForInterface(anInterface *ast.InterfaceType) string {
  formatString := "func (fake %s) %s (%s) (%s) {\n%s\n}"
	args := []interface{}{}
	return fmt.Sprintf(formatString, args...)
}
