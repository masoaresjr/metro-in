package test

// HttpTest struct used to aux http table-test-driven
type HttpTest struct {
	Description 	string
	Route			string
	ExpectedCode	int
	ExpectedBody	string
}
