package testInit

type TestCase struct {
	Method         string
	Endpoint       string
	Headers        interface{}
	Payload        interface{}
	ExpectedStatus int
}
