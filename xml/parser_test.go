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

func TestMissingInputInTestcase(t *testing.T) {
	rawXmlStructure := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<tester>
		<testcase>
			<args>-r priatel nepriatel</args>
			<input></input>
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

	testCases, err := getTestCasesFromXml(rawXmlStructure)

	if len(testCases) != 1 && err != nil {
		t.Errorf("Test Cases got: %d, expected: 1", len(testCases))
	}
}

func TestMissingExpectInTestcase(t *testing.T) {
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
			<expect></expect>
			<exitcode>0</exitcode>
		</testcase>
	</tester>`)

	testCases, err := getTestCasesFromXml(rawXmlStructure)

	if len(testCases) != 1 && err != nil {
		t.Errorf("Test Cases got: %d, expected: 1", len(testCases))
	}
}
