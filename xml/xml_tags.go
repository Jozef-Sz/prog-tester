package xml

import (
	"encoding/xml"
)

type TesterXmlTag struct {
	XMLName xml.Name         `xml:"tester"`
	All     []TestcaseXmlTag `xml:"testcase"`
}

type TestcaseXmlTag struct {
	Args     string       `xml:"args"`
	Input    InputXmlTag  `xml:"input"`
	Expect   ExpectXmlTag `xml:"expect"`
	ExitCode int          `xml:"exitcode"`
}

type InputXmlTag struct {
	Value   string `xml:",chardata"`
	Newline int    `xml:"newline,attr"`
}

type ExpectXmlTag struct {
	Value   string `xml:",chardata"`
	Newline int    `xml:"newline,attr"`
}
