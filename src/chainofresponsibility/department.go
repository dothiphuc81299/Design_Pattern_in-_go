package chainofresponsibility

// Department ...
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
