package thunk

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{thunk}
var thunk = &typewriter.Template{
	Name: "Thunk",
	Text: `

// {{.UName}}ThunkCtx ...
type {{.UName}}ThunkCtx struct {
	run    *sync.WaitGroup
	result <-chan {{.Name}}
	cancel chan<- interface{}
}

// Cancel ...
func (th {{.UName}}ThunkCtx) Cancel() {
	close(th.cancel)
	th.run.Done()
}

// Run ...
func (th {{.UName}}ThunkCtx) Run() <-chan {{.Name}} {
	th.run.Done()
	return th.result
}

// Force ...
func (th {{.UName}}ThunkCtx) Force() {{.Name}} {
	return <-th.Run()
}

// New ...
func New{{.UName}}Thunk(fn func() {{.Name}}) {{.UName}}ThunkCtx {
	result := make(chan {{.Name}})
	cancel := make(chan interface{}, 1)
	run := &sync.WaitGroup{}
	run.Add(1)

	go func() {
		run.Wait()

		select {
		case <-cancel:
			return
		default:
		}

		x := fn()

		select {
		case <-cancel:
		case result <- x:
		}
	}()

	return {{.UName}}ThunkCtx{run, result, cancel}
}

// ErrThunkTimeout ...
var ErrThunkTimeout = errors.New("Thunk Timeout")

// Timeout ...
func (th {{.UName}}ThunkCtx) Timeout(wait time.Duration) {{.UName}}ThunkCtx {
	result := make(chan {{.Name}})
	go func() {
		select {
		case <-time.After(wait):
			close(th.cancel)
			close(result)
		case result <- (<-th.result):
		}
	}()
	return {{.UName}}ThunkCtx{th.run, result, th.cancel}
}
`,
}
