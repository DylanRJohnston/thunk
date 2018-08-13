package thunk

import (
	"io"

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

func (th *Thunk) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(th)
	if !found {
		return nil
	}

	tmpl, err := templates.ByTag(t, tag)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, t)
	if err != nil {
		return err
	}

	return nil
}
