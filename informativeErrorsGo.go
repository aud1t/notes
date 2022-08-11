type Component int
type Code int

//go:generate stringer -type=Component
//go:generate stringer -type=Code
const (
    ErrOne Code = iota
    ErrTwo

    ComponentOne Component = iota
    ComponentTwo
)



type CompBaseError interface {
    error
    Code() Code
    Component() Component
}

type ErrorDescription struct {
    component Component
    code Code
    message string
}

func (e *ErrorDescription) Error() string {
    return fmt.Sprintf("%s | %s: %s", e.component, e.code, e.message)
}

func (e *ErrorDescription) Code() Code {
    return e.code
}

func (e *ErrorDescription) Component() Component {
    return e.component
}



func foo(val int) error {
    if val == 42 {
        return ErrorDescription{ComponentOne, ErrTwo, "42 is not allowed"}
    }
    // normal workflow
    return nil
}



...
err := foo(1)
if errCmpBase, ok := err.(CompBaseError); ok && errCmpBase.Component() == ComponentOne && errCmpBase.Code() == ErrTwo {
    // do some error handling for error from ComponentOne with code ErrTwo
}
// normal workflow
