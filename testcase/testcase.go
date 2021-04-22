package testcase

import "tester/xml"

type TestCase struct {
	id       int
	args     string
	input    string
	expect   string
	exitcode int
}

type Result struct {
	output   string
	exitCode int
	errorMsg string
	testCase TestCase
}

func GetTestCasesFromSchema(schemaName string) []TestCase {
	parsedTestcases := xml.ParseXmlSchema(schemaName)
	return serializeTestcases(parsedTestcases)
}

func serializeTestcases(rawTestcases []xml.TestcaseXmlTag) []TestCase {
	var testCases []TestCase
	for i, testcase := range rawTestcases {
		testCases[i] = TestCase{
			i,
			testcase.Args,
			testcase.Input,
			testcase.Expect,
			testcase.ExitCode}
	}
	return testCases
}
