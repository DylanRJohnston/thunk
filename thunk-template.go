package thunk

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{thunk}
var thunk = &typewriter.Template{
	Name: "Thunk",
	Text: ``,
}

var text = `
// {{.Name}}ThunkCtx ...
type {{.Name}}ThunkCtx struct {
	run    *sync.WaitGroup
	result <-chan {{.Name}}
	cancel chan<- interface{}
}

// Cancel ...
func (th {{.Name}}ThunkCtx) Cancel() {
	close(th.cancel)
	th.run.Done()
}

// Run ...
func (th {{.Name}}ThunkCtx) Run() <-chan {{.Name}} {
	th.run.Done()
	return th.result
}

// Force ...
func (th {{.Name}}ThunkCtx) Force() {{.Name}} {
	return <-th.Run()
}

// New ...
func New{{.Name}}Thunk(fn func() {{.Name}}) {{.Name}}ThunkCtx {
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

	return {{.Name}}ThunkCtx{run, result, cancel}
}

// ErrThunkTimeout ...
var ErrThunkTimeout = errors.New("Thunk Timeout")

// Timeout ...
func (th {{.Name}}ThunkCtx) Timeout(wait time.Duration) {{.Name}}ThunkCtx {
	result := make(chan {{.Name}})
	go func() {
		select {
		case <-time.After(wait):
			close(th.cancel)
			close(result)
		case result <- (<-th.result):
		}
	}()
	return {{.Name}}ThunkCtx{th.run, result, th.cancel}
}
`
