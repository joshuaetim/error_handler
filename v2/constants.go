package error_handler

type Color string

const (
	ErrNilPointer   = "nil pointer dereference"
	ErrOutOfRange   = "index out of range"
	ErrDivideByZero = "divide by zero"
)

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorReset  Color = "\u001b[0m"
)
