package main

import (
	"os"
	"fmt"
	"go/ast"
	"go/token"
	"go/parser"
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

	fmt.Printf("inside package %s, interface is %#v\n", packageName, interfaceNode)
}
