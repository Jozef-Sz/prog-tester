package main

import (
	"fmt"
	"tester/xml"
)

func main() {
	// cli.Execute()
	testCases := xml.ParseXmlSchema("test_schema.xml")
	fmt.Println(testCases[0].Args)
}
