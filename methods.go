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

func (rs Returns) String() string {
	pairs := []string{}
	for _, r := range rs {
		rStr := r.returnType
		if r.name != "" {
			rStr = strings.Join([]string{r.name, r.returnType}, " ")
		}

		pairs = append(pairs, rStr)
	}

	return strings.Join(pairs, ", ")
}

type Params []Param
type Param struct {
	name string
	paramType string
}

func (ps Params) String() string {
	pairs := []string{}
	for _, p := range ps {
		pStr := p.paramType
		if p.name != "" {
			pStr = strings.Join([]string{p.name, p.paramType}, " ")
		}

		pairs = append(pairs, pStr)
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
