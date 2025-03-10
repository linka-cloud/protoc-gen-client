// Copyright 2022 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/alta/protopatch/patch/gopb"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func Module() *client {
	return &client{
		ModuleBase: &pgs.ModuleBase{},
		imports:    make(map[string]struct{}),
	}
}

type client struct {
	*pgs.ModuleBase
	ctx     pgsgo.Context
	imports map[string]struct{}
	tpl     *template.Template
}

func (p *client) Name() string {
	return "client"
}

func (p *client) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("client").Funcs(map[string]interface{}{
		"package": p.ctx.PackageName,
		"name": func(m pgs.Message) string {
			if _, ok := p.imports[p.ctx.ImportPath(m).String()]; ok {
				return fmt.Sprintf("%s.%s", p.ctx.PackageName(m), p.ctx.Name(m).String())
			}
			return p.ctx.Name(m).String()
		},
		"serverName":   p.ctx.ServerName,
		"clientName":   p.ctx.ClientName,
		"serverStream": p.ctx.ServerStream,
		"params": func(m pgs.Message) string {
			params := []string{"ctx context.Context"}
			oneofs := make(map[string]struct{})
			for _, v := range m.Fields() {
				t := p.ctx.Type(v).String()
				var o *gopb.Options
				if ok, _ := v.Extension(gopb.E_Field, &o); ok && o.GetType() != "" {
					if v.InOneOf() && v.HasOptionalKeyword() {
						t = "*" + o.GetType()
					} else {
						t = o.GetType()
					}
				}
				if _, ok := p.imports[p.ctx.ImportPath(m).String()]; ok && o.GetType() != "" {
					t = fmt.Sprintf("%s.%s", p.ctx.PackageName(m), t)
				}
				switch {
				case v.InRealOneOf():
					if _, ok := oneofs[v.OneOf().Name().String()]; ok {
						continue
					}
					oneofs[v.OneOf().Name().String()] = struct{}{}
					params = append(params, fmt.Sprintf("%s is%s_%s", p.ctx.Name(v.OneOf()).LowerCamelCase(), p.ctx.Name(m), p.ctx.Name(v.OneOf())))
				default:
					params = append(params, fmt.Sprintf("%s %s", p.ctx.Name(v).LowerCamelCase(), t))
				}
			}
			return strings.Join(params, ", ")
		},
		"returns": func(m pgs.Message) string {
			var returns []string
			oneofs := make(map[string]struct{})
			for _, v := range m.Fields() {
				t := p.ctx.Type(v).String()
				var o *gopb.Options
				if ok, _ := v.Extension(gopb.E_Field, &o); ok && o.GetType() != "" {
					if v.InOneOf() && v.HasOptionalKeyword() {
						t = "*" + o.GetType()
					} else {
						t = o.GetType()
					}
				}
				if _, ok := p.imports[p.ctx.ImportPath(m).String()]; ok && o.GetType() != "" {
					t = fmt.Sprintf("%s.%s", p.ctx.PackageName(m), t)
				}
				switch {
				case v.InRealOneOf():
					if _, ok := oneofs[v.OneOf().Name().String()]; ok {
						continue
					}
					oneofs[v.OneOf().Name().String()] = struct{}{}
					returns = append(returns, fmt.Sprintf("%s is%s_%s", p.ctx.Name(v.OneOf()), p.ctx.Name(m), p.ctx.Name(v.OneOf())))
				default:
					returns = append(returns, fmt.Sprintf("%s %s", p.ctx.Name(v), t))
				}
			}
			return strings.Join(append(returns, "err error"), ", ")
		},
		"empty": func(m pgs.Message) bool {
			return len(m.Fields()) == 0
		},
		"requestParams": func(m pgs.Message) string {
			var fields []string
			oneofs := make(map[string]struct{})
			typeName := p.ctx.Name(m).UpperCamelCase().String()
			if _, ok := p.imports[p.ctx.ImportPath(m).String()]; ok {
				typeName = fmt.Sprintf("%s.%s", p.ctx.PackageName(m), typeName)
			}
			for _, v := range m.Fields() {
				var o *gopb.Options
				v.Extension(gopb.E_Field, &o)
				switch {
				case v.InRealOneOf():
					if _, ok := oneofs[v.OneOf().Name().String()]; ok {
						continue
					}
					oneofs[v.OneOf().Name().String()] = struct{}{}
					name := p.ctx.Name(v.OneOf()).String()
					if n := o.GetName(); n != "" {
						name = n
					}
					fields = append(fields, fmt.Sprintf("%s: %s", name, p.ctx.Name(v.OneOf()).LowerCamelCase()))
				default:
					name := p.ctx.Name(v).String()
					if n := o.GetName(); n != "" {
						name = n
					}
					fields = append(fields, fmt.Sprintf("%s: %s", name, p.ctx.Name(v).LowerCamelCase()))
				}
			}
			return fmt.Sprintf("ctx, &%s{%s}, opts...", typeName, strings.Join(fields, ", "))
		},
		"response": func(m pgs.Message) string {
			var fields []string
			oneofs := make(map[string]struct{})
			for _, v := range m.Fields() {
				switch {
				case v.InRealOneOf():
					if _, ok := oneofs[v.OneOf().Name().String()]; ok {
						continue
					}
					oneofs[v.OneOf().Name().String()] = struct{}{}
					var o *gopb.Options
					v.Extension(gopb.E_Field, &o)
					name := p.ctx.Name(v.OneOf()).String()
					if n := o.GetName(); n != "" {
						name = n
					}
					fields = append(fields, fmt.Sprintf("res.%s", name))
				default:
					var o *gopb.Options
					v.Extension(gopb.E_Field, &o)
					name := p.ctx.Name(v).String()
					if n := o.GetName(); n != "" {
						name = n
					}
					fields = append(fields, fmt.Sprintf("res.%s", name))
				}
			}
			return strings.Join(append(fields, "nil"), ", ")
		},
		"comment": func(s string) string {
			var out string
			parts := strings.Split(s, "\n")
			for i, v := range parts {
				if i == len(parts)-1 && v == "" {
					return out
				}
				out += "//" + v + "\n"
			}
			return out
		},
		"imports": func() string {
			var imports string
			for v := range p.imports {
				imports += fmt.Sprintf("\"%s\"\n", v)
			}
			return imports
		},
	})
	p.tpl = template.Must(tpl.Parse(fieldsTpl))
}

func (p *client) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		if len(f.Services()) == 0 {
			continue
		}
		var gen bool
		for _, v := range f.Services() {
			for _, v := range v.Methods() {
				if v.ClientStreaming() {
					continue
				}
				gen = true
				break
			}
		}
		if gen {
			p.generate(f)
		}
	}
	return p.Artifacts()
}

func (p *client) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}
	name := p.ctx.OutputPath(f).SetExt(".client.go")
	p.imports = make(map[string]struct{})
	for _, s := range f.Services() {
		for _, m := range s.Methods() {
			// TODO(adphi): support client, server and bidirectional streams
			if m.ServerStreaming() || m.ClientStreaming() {
				continue
			}
			current := p.ctx.ImportPath(s).String()
			if i := p.ctx.ImportPath(m.Input()).String(); i != current {
				p.imports[i] = struct{}{}
			}
			for _, f := range m.Input().Fields() {
				if !f.Type().IsEmbed() {
					continue
				}
				if i := p.ctx.ImportPath(f.Type().Embed()).String(); i != current {
					p.imports[i] = struct{}{}
				}
			}
			for _, f := range m.Output().Fields() {
				if !f.Type().IsEmbed() {
					continue
				}
				if i := p.ctx.ImportPath(f.Type().Embed()).String(); i != current {
					p.imports[i] = struct{}{}
				}
			}
		}
	}
	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}

const fieldsTpl = `{{ comment .SyntaxSourceCodeInfo.LeadingComments }}
{{ range .SyntaxSourceCodeInfo.LeadingDetachedComments }}
{{ comment . }}
{{ end }}
// Code generated by protoc-gen-defaults. DO NOT EDIT.

package {{ package . }}

import (
    "context"
    "errors"
    "io"

    "google.golang.org/grpc"
    "google.golang.org/grpc/status"
    {{ imports }}
)

{{ range .Services }}
{{ $name := .Name }}

var _ {{ $name }} = (*client{{ $name }})(nil)

type {{ $name }} interface {
{{- range .Methods }}
{{- if and .ClientStreaming .ServerStreaming }}

{{- else if .ClientStreaming }}

{{- else if .ServerStreaming }}
	{{ .Name }}({{ params .Input}}, opts ...grpc.CallOption) (<-chan *{{ .Name }}Msg, error)
{{ else }}
	{{ .Name }}({{ params .Input}}, opts ...grpc.CallOption) ({{ returns .Output }})
{{- end }}
{{- end }}
}

func New{{ $name }}(cc grpc.ClientConnInterface) {{ $name }} {
	return &client{{ $name }}{c: New{{ clientName . }}(cc)}
}

type client{{ $name }} struct {
	c {{ $name }}Client
}

{{ range .Methods }}
{{- if and .ClientStreaming .ServerStreaming }}

{{- else if .ClientStreaming }}

{{- else if .ServerStreaming }}

type {{ .Name }}Msg struct {
	Msg *{{ name .Output }}
	Err error
}

func (x *client{{ $name }}) {{ .Name }}({{ params .Input}}, opts ...grpc.CallOption) (cn <-chan *{{ .Name }}Msg, err error) {
	ss, err := x.c.{{ .Name }}({{ requestParams .Input }})
	if err != nil {
		return nil, err
	}
	ch := make(chan *{{ .Name }}Msg, 1)
	go func() {
		defer close(ch)
		for {
			res, err := ss.Recv()
			if errors.Is(io.EOF, err) {
				return
			}
			if err != nil {
				ch <- &{{ .Name }}Msg{Err: err}
				break
			}
			ch <- &{{ .Name }}Msg{Msg: res}
		}
	}()
	return ch, nil
}

{{ else }}
// {{ .Name }} ...
func (x *client{{ $name }}) {{ .Name }}({{ params .Input}}, opts ...grpc.CallOption) ({{ returns .Output }}){
    {{- if not (empty .Output) }}var res *{{ name .Output }}{{ end }}
    {{ if not (empty .Output) }}res{{ else }}_{{ end }}, err = x.c.{{ .Name }}({{ requestParams .Input }})
    err = x.unwrap(err)
    if err != nil {
        return
    }
    return {{ response .Output }}
}
{{ end }}
{{ end }}
// unwrap convert grpc status error to go error
func (x *client{{ $name }}) unwrap(err error) error {
	s, ok := status.FromError(err)
	if !ok {
		return err
	}
	if s != nil {
		return errors.New(s.Message())
	}
	return nil
}
{{ end }}
`
