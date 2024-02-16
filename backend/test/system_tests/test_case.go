package system_tests

type testCase struct {
	Method         string
	Endpoint       string
	Payload        interface{}
	ExpectedStatus int
}
