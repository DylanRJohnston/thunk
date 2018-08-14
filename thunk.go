package thunk

import (
	"io"
	"strings"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewThunk())
	if err != nil {
		panic(err)
	}
}

type Thunk struct{}

func NewThunk() *Thunk {
	return &Thunk{}
}

func (th *Thunk) Name() string {
	return "thunk"
}

func (th *Thunk) Imports(t typewriter.Type) []typewriter.ImportSpec {
	return []typewriter.ImportSpec{}
}

type params struct {
	Name  string
	UName string
}

func (th *Thunk) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(th)
	if !found {
		return nil
	}

	// For generating templates against builtin types
	if tag.Values[0].Name == "UnderlyingType" {
		t.Name = t.Underlying().String()
	}

	tmpl, err := templates.ByTag(t, tag)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, params{Name: t.Name, UName: strings.Title(t.Name)})
	if err != nil {
		return err
	}

	return nil
}
