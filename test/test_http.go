package test

// HTTPTest struct used to aux http table-test-driven
type HTTPTest struct {
	Description  string
	Route        string
	ExpectedCode int
	ExpectedBody string
}
