package test

type testHelper interface {
	Helper()
}

type testFailer interface {
	Fatalf(string, ...interface{})
}

type tester interface {
	testHelper
	testFailer
}
