package testInit

type TestCase struct {
	Method         string
	Endpoint       string
	Payload        interface{}
	ExpectedStatus int
}
