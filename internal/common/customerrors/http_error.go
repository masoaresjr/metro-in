package customerrors

//TODO repensar nome e estratégia

// HttpError interface used to specify some errors as BadRequestErrors
type HttpError interface {
	error
	GetStatusCode() int
}
