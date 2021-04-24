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
			appendNewlines(testcase.Input.Value, testcase.Input.Newline),
			appendNewlines(testcase.Expect.Value, testcase.Expect.Newline),
			testcase.ExitCode})
	}
	return testCases
}

func appendNewlines(text string, count int) string {
	for i := 0; i < count; i++ {
		text += "\n"
	}
	return text
}
