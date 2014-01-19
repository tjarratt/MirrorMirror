package main

import (
	"fmt"
	"strings"
	"go/ast"
)

type Returns []Return
type Return struct {
	name string
	returnType string
}

func (r Return) UpperName() string {
	return strings.ToUpper(string(r.name[0])) + r.name[1 : len(r.name)]
}

func (rs Returns) String() string {
	pairs := []string{}
	for index, r := range rs {
		if r.name != "" {
			pairs = append(pairs, strings.Join([]string{r.name, r.returnType}, " "))
		} else {
			pairs = append(pairs, strings.Join([]string{"return" + fmt.Sprintf("%d", index + 1), r.returnType}, " "))
		}
	}

	return strings.Join(pairs, ", ")
}

type Params []Param
type Param struct {
	name string
	paramType string
}

func (p Param) UpperName() string {
	return strings.ToUpper(string(p.name[0])) + p.name[1 : len(p.name)]
}

func (ps Params) String() string {
	pairs := []string{}
	for index, p := range ps {
		if p.name != "" {
			pairs = append(pairs, strings.Join([]string{p.name, p.paramType}, " "))
		} else {
			pairs = append(pairs, strings.Join([]string{fmt.Sprintf("param%d", index + 1), p.paramType}, " "))
		}
	}

	return strings.Join(pairs, ", ")
}

type Method struct {
	Name string
	ParamSlice Params
	ReturnSlice Returns
}

func newMethod(node *ast.Field) (m Method) {
	m.Name = node.Names[0].Name
	m.ParamSlice = paramsForNode(node.Type.(*ast.FuncType))
	m.ReturnSlice = returnsForNode(node.Type.(*ast.FuncType))
	return
}

func paramsForNode(node *ast.FuncType) (ps Params) {
	ps = []Param{}
	if node.Params != nil && node.Params.List != nil {
		for _, p := range node.Params.List {
			var param Param
			param.paramType = string(p.Type.(*ast.Ident).Name)

			if len(p.Names) > 0 {
				param.name = fmt.Sprintf("%s", p.Names[0])
			}

			ps = append(ps, param)
		}
	}

	return
}

func returnsForNode(node *ast.FuncType) (rs Returns) {
	rs = []Return{}
	if node.Results != nil && node.Results.List != nil {
		for _, r := range node.Results.List {
			var ret Return
			ret.returnType = string(r.Type.(*ast.Ident).Name)

			if len(r.Names) > 0 {
				ret.name = fmt.Sprintf("%s", r.Names[0])
			}

			rs = append(rs, ret)
		}
	}
	return
}
