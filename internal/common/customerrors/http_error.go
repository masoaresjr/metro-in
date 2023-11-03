package customerrors

//TODO repensar nome e estratégia

// HTTPError interface used to specify some errors as BadRequestErrors
type HTTPError interface {
	error
	GetStatusCode() int
}
