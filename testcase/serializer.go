package testcase

import "tester/xml"

func GetTestCasesFromSchema(schemaName string) []TestCase {
	parsedTestcases := xml.ParseXmlSchema(schemaName)
	return serializeTestcases(parsedTestcases)
}

func serializeTestcases(rawTestcases []xml.TestcaseXmlTag) []TestCase {
	var testCases []TestCase
	for i, testcase := range rawTestcases {
		testCases = append(testCases, TestCase{
			i + 1,
			testcase.Args,
			testcase.Input,
			testcase.Expect,
			testcase.ExitCode})
	}
	return testCases
}
