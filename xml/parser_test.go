package xml

import (
	"testing"
)

func TestParsingCorrectXml(t *testing.T) {
	rawXmlStructure := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<tester>
		<testcase>
			<args>-r priatel nepriatel</args>
			<input>The ladies of Longbourn soon waited on those of Netherfield  The</input>
			<expect>Mrs  Hurst and Miss Bingley  and</expect>
			<exitcode>1</exitcode>
		</testcase>
	
		<testcase>
			<args>-r "Godzilla vs. Kong" zadanie_4</args>
			<input>Velmi sa tesim na zadanie_4. Len ci sa take zadanie da vobec pripravit.</input>
			<expect>Len ci sa take zadanie da vobec pripravit.</expect>
			<exitcode>0</exitcode>
		</testcase>
	</tester>`)

	firstTestCase := TestCase{
		"-r priatel nepriatel",
		"The ladies of Longbourn soon waited on those of Netherfield  The",
		"Mrs  Hurst and Miss Bingley  and",
		1,
	}
	secondTestCase := TestCase{
		"-r \"Godzilla vs. Kong\" zadanie_4",
		"Velmi sa tesim na zadanie_4. Len ci sa take zadanie da vobec pripravit.",
		"Len ci sa take zadanie da vobec pripravit.",
		1,
	}

	testCases, err := getTestCasesFromXml(rawXmlStructure)

	if !testCases[0].equals(firstTestCase) && !testCases[1].equals(secondTestCase) && err != nil {
		t.Error("Test Cases does not match with expected")
	}
}

func TestIncorrectXmlSyntax(t *testing.T) {
	rawXmlStructure := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<tester>
		<testcase>
			<arguss>-r priatel nepriatel</arrrgs>
			<input>The ladies of Longbourn soon waited on those of Netherfield  The</input>
			<expect>Mrs  Hurst and Miss Bingley  and</expect>
			<exitcode>1</exitcode>
		</testcase>
	</tester>`)

	testCases, err := getTestCasesFromXml(rawXmlStructure)

	if testCases != nil && err == nil {
		t.Error("Test Cases should be nil and an error is expected")
	}
}

func TestEmptyTestCase(t *testing.T) {
	rawXmlStructure := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<tester>
		<testcase>
		</testcase>
	</tester>`)

	expected := TestCase{
		"",
		"",
		"",
		0,
	}

	testCases, err := getTestCasesFromXml(rawXmlStructure)

	if !testCases[0].equals(expected) && err != nil {
		t.Errorf("got: %+v, expected: %+v", testCases[0], expected)
	}
}

func TestOmittedInputAndExpect(t *testing.T) {
	rawXmlStructure := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<tester>
		<testcase>
			<args>-r hello</args>
			<exitcode>15</exitcode>
		</testcase>
	</tester>`)

	expected := TestCase{
		"-r hello",
		"",
		"",
		15,
	}

	testCases, err := getTestCasesFromXml(rawXmlStructure)

	if !testCases[0].equals(expected) && err != nil {
		t.Errorf("got: %+v, expected: %+v", testCases[0], expected)
	}
}
