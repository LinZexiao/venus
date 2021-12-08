package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/filecoin-project/venus/app/client/funcrule"
	"golang.org/x/xerrors"
)

// Rule[perm:read,ignore:true]
var rulePattern = `Rule\[(?P<rule>.*)\]`

type ruleKey = string

const (
	rkPerm   ruleKey = "perm"
	rkIgnore ruleKey = "ignore"
)

var defaultPerm = []string{"perm", "read"}

var regRule, _ = regexp.Compile(rulePattern)

func parseRule(comment string) (*funcrule.Rule, map[string][]string) {
	rule := new(funcrule.Rule)
	match := regRule.FindStringSubmatch(comment)
	tags := map[string][]string{}
	if len(match) == 2 {
		pairs := strings.Split(match[1], ",")
		for _, v := range pairs {
			pair := strings.Split(v, ":")
			if len(pair) != 2 {
				continue
			}
			switch pair[0] {
			case rkPerm:
				tags[rkPerm] = pair
				rule.Perm = pair[1]
			case rkIgnore:
				ig, err := strconv.ParseBool(pair[1])
				if err != nil {
					panic("the rule tag is invalid format")
				}
				rule.Ignore = ig
			}
		}
	} else {
		rule.Perm = "read"
		tags[rkPerm] = defaultPerm
	}
	return rule, tags
}

type methodMeta struct {
	node  ast.Node
	ftype *ast.FuncType
}

type Visitor struct {
	Methods map[string]map[string]*methodMeta
	Include map[string][]string
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	st, ok := node.(*ast.TypeSpec)
	if !ok {
		return v
	}

	iface, ok := st.Type.(*ast.InterfaceType)
	if !ok {
		return v
	}
	if v.Methods[st.Name.Name] == nil {
		v.Methods[st.Name.Name] = map[string]*methodMeta{}
	}
	for _, m := range iface.Methods.List {
		switch ft := m.Type.(type) {
		case *ast.Ident:
			v.Include[st.Name.Name] = append(v.Include[st.Name.Name], ft.Name)
		case *ast.FuncType:
			v.Methods[st.Name.Name][m.Names[0].Name] = &methodMeta{
				node:  m,
				ftype: ft,
			}
		}
	}

	return v
}

func main() {
	var arg string
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	apiPath := "../lotus/api"
	rootPath := path.Join(os.Getenv("GOPATH"), "pkg/mod/github.com/filecoin-project")
	dirs, err := os.ReadDir(rootPath)
	if err == nil {
		// Select the latest version of Lotus
		for _, dir := range dirs {
			if strings.Contains(dir.Name(), "lotus") {
				fmt.Println(dir.Name())
				apiPath = path.Join(rootPath, dir.Name(), "/api")
			}
		}
	}
	fmt.Println("lotus api path:", apiPath)

	bmp, err := benchmarkMethodPerm(apiPath)
	checkError(err)
	//outputWithJSON(bmp, "benchmarkMethodPerm: ")

	mm, err := methodMetaFromInterface("./app/client", "apiface", "client")
	checkError(err)

	smi := check(bmp, mm)
	data, err := json.MarshalIndent(smi, "", "\t")
	checkError(err)
	err = ioutil.WriteFile("./tools/gen/api/stable_method_info.json", data, 0666)
	checkError(err)
	outputWithJSON(smi, "StableMethodInfo: ")

	if arg != "compare" {
		outfile := "./app/client/full.go"
		checkError(doTemplate(outfile, mm, templ))
	}
}

func benchmarkMethodPerm(rootPath string) (map[string]string, error) {
	fileNames := []string{"api_full.go", "api_common.go", "api_net.go"}
	fset := token.NewFileSet()
	files := make([]*ast.File, 0, len(fileNames))
	visitor := &Visitor{make(map[string]map[string]*methodMeta), map[string][]string{}}

	for _, fname := range fileNames {
		f, err := parser.ParseFile(fset, path.Join(rootPath, fname), nil, parser.AllErrors|parser.ParseComments)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
		ast.Walk(visitor, f)
	}

	perms := make(map[string]string)
	for _, f := range files {
		cmap := ast.NewCommentMap(fset, f, f.Comments)
		for _, methods := range visitor.Methods {
			for mname, node := range methods {
				filteredComments := cmap.Filter(node.node).Comments()
				if len(filteredComments) > 0 {
					cmt := filteredComments[len(filteredComments)-1].List[0].Text
					if !strings.Contains(cmt, "perm:") {
						fmt.Println("lotus method not found perm: ", mname)
						continue
					}
					pairs := strings.Split(cmt, ":")
					if len(pairs) != 2 {
						continue
					}
					perms[mname] = pairs[1]
				}
			}
		}
	}

	return perms, nil
}

type methodInfo struct {
	Name                                     string
	node                                     ast.Node
	Tags                                     map[string][]string
	NamedParams, ParamNames, Results, DefRes string
}
type strinfo struct {
	Name    string
	Methods map[string]*methodInfo
	Include []string
}
type meta struct {
	Infos   map[string]*strinfo
	Imports map[string]string
	OutPkg  string
}

func methodMetaFromInterface(rootPath string, pkg, outpkg string) (*meta, error) {
	fset := token.NewFileSet()
	apiDir, err := filepath.Abs(rootPath)
	if err != nil {
		return nil, err
	}

	visitor := &Visitor{make(map[string]map[string]*methodMeta), map[string][]string{}}
	m := &meta{
		OutPkg:  outpkg,
		Infos:   map[string]*strinfo{},
		Imports: map[string]string{},
	}
	//filter := isGoFile
	pkgs, err := parser.ParseDir(fset, path.Join(apiDir, pkg), nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return nil, err
	}
	ap := pkgs[pkg]

	ast.Walk(visitor, ap)
	ignoreMethods := map[string][]string{}
	for _, f := range ap.Files {
		cmap := ast.NewCommentMap(fset, f, f.Comments)
		for _, im := range f.Imports {
			m.Imports[im.Path.Value] = im.Path.Value
			if im.Name != nil {
				m.Imports[im.Path.Value] = im.Name.Name + " " + m.Imports[im.Path.Value]
			}
		}

		for ifname, methods := range visitor.Methods {
			if _, ok := m.Infos[ifname]; !ok {
				m.Infos[ifname] = &strinfo{
					Name:    ifname,
					Methods: map[string]*methodInfo{},
					Include: visitor.Include[ifname],
				}
			}
			info := m.Infos[ifname]
			for mname, node := range methods {
				filteredComments := cmap.Filter(node.node).Comments()
				if _, ok := info.Methods[mname]; !ok {
					var params, pnames []string
					for _, param := range node.ftype.Params.List {
						pstr, err := typeName(param.Type, outpkg)
						if err != nil {
							return nil, err
						}

						c := len(param.Names)
						if c == 0 {
							c = 1
						}

						for i := 0; i < c; i++ {
							pname := fmt.Sprintf("p%d", len(params))
							pnames = append(pnames, pname)
							params = append(params, pname+" "+pstr)
						}
					}

					var results []string
					for _, result := range node.ftype.Results.List {
						rs, err := typeName(result.Type, outpkg)
						if err != nil {
							return nil, err
						}
						results = append(results, rs)
					}

					defRes := ""
					if len(results) > 1 {
						defRes = results[0]
						switch {
						case defRes[0] == '*' || defRes[0] == '<', defRes == "interface{}":
							defRes = "nil"
						case defRes == "bool":
							defRes = "false"
						case defRes == "string":
							defRes = `""`
						case defRes == "int", defRes == "int64", defRes == "uint64", defRes == "uint":
							defRes = "0"
						default:
							defRes = "*new(" + defRes + ")"
						}
						defRes += ", "
					}

					info.Methods[mname] = &methodInfo{
						Name:        mname,
						node:        node.node,
						Tags:        map[string][]string{},
						NamedParams: strings.Join(params, ", "),
						ParamNames:  strings.Join(pnames, ", "),
						Results:     strings.Join(results, ", "),
						DefRes:      defRes,
					}
				}

				// try to parse tag info
				if len(filteredComments) > 0 {
					cmt := filteredComments[0].List[len(filteredComments[0].List)-1].Text
					rule, tags := parseRule(cmt)
					info.Methods[mname].Tags[rkPerm] = tags[rkPerm]
					// remove ignore method
					if rule.Ignore {
						ignoreMethods[ifname] = append(ignoreMethods[ifname], mname)
					}
				}
			}
		}
	}
	for ifname, mnames := range ignoreMethods {
		for _, mname := range mnames {
			delete(m.Infos[ifname].Methods, mname)
		}
	}

	return m, nil
}

func typeName(e ast.Expr, pkg string) (string, error) {
	switch t := e.(type) {
	case *ast.SelectorExpr:
		return t.X.(*ast.Ident).Name + "." + t.Sel.Name, nil
	case *ast.Ident:
		pstr := t.Name
		if !unicode.IsLower(rune(pstr[0])) && pkg != "client" {
			pstr = "client." + pstr // todo src pkg name
		}
		return pstr, nil
	case *ast.ArrayType:
		subt, err := typeName(t.Elt, pkg)
		if err != nil {
			return "", err
		}
		return "[]" + subt, nil
	case *ast.StarExpr:
		subt, err := typeName(t.X, pkg)
		if err != nil {
			return "", err
		}
		return "*" + subt, nil
	case *ast.MapType:
		k, err := typeName(t.Key, pkg)
		if err != nil {
			return "", err
		}
		v, err := typeName(t.Value, pkg)
		if err != nil {
			return "", err
		}
		return "map[" + k + "]" + v, nil
	case *ast.StructType:
		if len(t.Fields.List) != 0 {
			return "", xerrors.Errorf("can't struct")
		}
		return "struct{}", nil
	case *ast.InterfaceType:
		if len(t.Methods.List) != 0 {
			return "", xerrors.Errorf("can't interface")
		}
		return "interface{}", nil
	case *ast.ChanType:
		subt, err := typeName(t.Value, pkg)
		if err != nil {
			return "", err
		}
		if t.Dir == ast.SEND {
			subt = "->chan " + subt
		} else {
			subt = "<-chan " + subt
		}
		return subt, nil
	default:
		return "", xerrors.Errorf("unknown type")
	}
}

type stableMethodInfo struct {
	// Lotus and Venus both have functions and the same permissions
	Common map[string]string
	// Venus has functions that Lotus does not
	Extend map[string]string
	// Lotus has functions that Venus does not
	Loss map[string]string
	// Both Lotus and venus has functions but the permissions are different
	Gap map[string]string
}

func newStableMethodInfo() *stableMethodInfo {
	return &stableMethodInfo{
		Common: make(map[string]string),
		Extend: make(map[string]string),
		Loss:   make(map[string]string),
		Gap:    make(map[string]string),
	}
}

func check(bmp map[string]string, m *meta) *stableMethodInfo {
	smi := newStableMethodInfo()
	vMethodPerms := make(map[string]string)
	for _, info := range m.Infos {
		for _, one := range info.Methods {
			mperm := one.Tags[rkPerm][1]
			vMethodPerms[one.Name] = mperm
			if perm, ok := bmp[one.Name]; ok {
				if mperm != perm {
					smi.Gap[one.Name] = fmt.Sprintf("venus:%s lotus:%s", mperm, perm)
					continue
				}
				smi.Common[one.Name] = mperm
			} else {
				smi.Extend[one.Name] = mperm
			}
		}
	}
	for m, p := range bmp {
		if _, ok := vMethodPerms[m]; !ok {
			smi.Loss[m] = p
		}
	}

	return smi
}

func doTemplate(outfile string, info interface{}, templ string) error {
	w, err := os.OpenFile(outfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	t := template.Must(template.New("").
		Funcs(template.FuncMap{}).Parse(templ))

	return t.Execute(w, info)
}

var templ = `// Code generated by github.com/filecoin-project/tools/gen/api. DO NOT EDIT.

package {{.OutPkg}}

import (
{{range .Imports}}	{{.}}
{{end}}
)

{{range .Infos}}

{{$name := .Name}}
type {{.Name}}Struct struct {
{{range .Include}}	{{.}}Struct
{{end}}

{{ if gt (len .Methods) 0 }}
  Internal struct {
    {{range .Methods}}	{{.Name}} func({{.NamedParams}}) ({{.Results}}) ` + "`" + `{{range .Tags}}{{index . 0}}:"{{index . 1}}"{{end}}` + "`" + `
    {{end}}
  }
{{ end }}
}

{{range .Methods}}  func(s *{{$name}}Struct) {{.Name}} ({{.NamedParams}}) ({{.Results}}){
  return s.Internal.{{.Name}}({{.ParamNames}})
}

{{end}}

{{end}}
`

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func outputWithJSON(obj interface{}, comment string) {
	b, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		fmt.Println("json marshal error: ", err)
	}
	fmt.Println(comment, "\n", string(b))
}
