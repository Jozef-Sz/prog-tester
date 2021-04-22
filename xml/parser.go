package xml

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type TesterXmlTag struct {
	XMLName xml.Name         `xml:"tester"`
	All     []TestcaseXmlTag `xml:"testcase"`
}

type TestcaseXmlTag struct {
	Args     string `xml:"args"`
	Input    string `xml:"input"`
	Expect   string `xml:"expect"`
	ExitCode int    `xml:"exitcode"`
}

func ParseXmlSchema(fileName string) []TestcaseXmlTag {
	xmlFile := openXmlFile(fileName)
	defer xmlFile.Close()

	xmlFileByteValue, _ := ioutil.ReadAll(xmlFile)

	testCases := getTestcasesFromXml(xmlFileByteValue)
	return testCases
}

func openXmlFile(filePath string) *os.File {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("[ERROR]: Couldn't open %s. Make sure that the name is corrent.\n", filePath)
	}
	return xmlFile
}

func getTestcasesFromXml(xmlBytes []byte) []TestcaseXmlTag {
	var testcases TesterXmlTag
	err := xml.Unmarshal(xmlBytes, &testcases)
	if err != nil {
		log.Fatalf("[ERROR]: %s.\n", err)
	}
	return testcases.All
}
