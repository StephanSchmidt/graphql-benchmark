package main

import (
	"bytes"
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/labstack/echo/v4"
	"io"
	"reflect"
	"strings"
)

type Map map[string]interface{}

type Options struct {
	Loader          jet.Loader
	Directory       string
	DevelopmentMode bool
}

type Renderer struct {
	templates *jet.Set
}

// Jet examples:
// https://github.com/CloudyKit/jet/blob/master/examples/todos/main.go

func New(o Options) *Renderer {
	r := &Renderer{}

	if o.Loader != nil {
		r.templates = jet.NewHTMLSetLoader(o.Loader)
	} else {
		r.templates = jet.NewHTMLSet(o.Directory)
		r.templates.SetDevelopmentMode(o.DevelopmentMode)
	}
	r.templates.AddGlobalFunc("dash", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("dash", 1, 1)
		buffer := bytes.NewBuffer(nil)
		fmt.Fprint(buffer, a.Get(0))
		value := string(buffer.Bytes())
		if len(strings.TrimSpace(value)) == 0 {
			return reflect.ValueOf("-")
		} else {
			return reflect.ValueOf(value)
		}
	})
	return r
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t, err := r.templates.GetTemplate(name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// convert Map into jet.VarMap
	vars := make(jet.VarMap)

	// Add global methods if data is a map
	if datamap, ok := data.(map[string]interface{}); ok {
		for k := range datamap {
			vars.Set(k, datamap[k])
		}
	}
	if datamap, ok := data.(Map); ok {
		for k := range datamap {
			vars.Set(k, datamap[k])
		}
	}

	// render template
	if err = t.Execute(w, vars, data); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func configRenderer(e *echo.Echo) {
	e.Renderer = New(Options{
		Directory:       "web/template/", // Path from current working dir
		DevelopmentMode: false,
	})
}
