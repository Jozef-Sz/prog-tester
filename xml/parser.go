package xml

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

const ERR_HELP_MSG = `Most likely you are trying to use reserved xml tokens.
Try replace one of these:
"   &quot;
'   &apos;
<   &lt;
>   &gt;
&   &amp;`

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
		log.Fatalf("[ERROR]: %s.\n%s", err, ERR_HELP_MSG)
	}
	return testcases.All
}
