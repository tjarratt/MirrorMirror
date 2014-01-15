package main

import (
	"os"
	"fmt"
	"strings"
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
	fakeMethods := ""
	for _, method := range anInterface.Methods.List {
		var params, returns string
		var funcNode *ast.FuncType
		funcNode = method.Type.(*ast.FuncType)

		if funcNode.Params != nil && funcNode.Params.List != nil&& len(funcNode.Params.List) > 0 {
			paramsSlice := []string{}
			for _, p := range funcNode.Params.List {
				var paramName, paramPair string
				paramType := string(p.Type.(*ast.Ident).Name)

				if len(p.Names) > 0 {
					paramName = fmt.Sprintf("%s", p.Names[0])
					paramPair = strings.Join([]string{paramName, paramType}, " ")
				} else {
					paramPair = paramType
				}

				paramsSlice = append(paramsSlice, paramPair)
			}

			params = strings.Join(paramsSlice, ", ")
		}

		if funcNode.Results != nil && funcNode.Results.List != nil && len(funcNode.Results.List) > 0 {
			resultsSlice := []string{}
			for _, r := range funcNode.Results.List {
				var returnName, returnPair string
				returnType := string(r.Type.(*ast.Ident).Name)

				if len(r.Names) > 0 {
					returnName = fmt.Sprintf("%s", r.Names[0])
					returnPair = strings.Join([]string{returnName, returnType}, " ")
				} else {
					returnPair = returnType
				}

				resultsSlice = append(resultsSlice, returnPair)
			}

			returns = strings.Join(resultsSlice, ", ")
		}

		fakeMethods = fmt.Sprintf("%s\t%s (%s) (%s)\n", fakeMethods, method.Names[0].Name, params, returns)
	}

	return fmt.Sprintf("type Fake%s struct {\n%s}", name, fakeMethods)
}

func StubbedMethodsForInterface(anInterface *ast.InterfaceType) string {
  formatString := "func (fake %s) %s (%s) (%s) {\n%s\n}"
	args := []interface{}{}
	return fmt.Sprintf(formatString, args...)
}
