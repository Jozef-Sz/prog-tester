package main

import "tester/cli"

var VERSION = "0.4"

func main() {
	cli.Execute(VERSION)
	// testCases := xml.ParseXmlSchema("test_schema.xml")
	// fmt.Println(testCases[0].Args)
}
