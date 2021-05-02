package tester

import (
	"fmt"
	"os"
	"tester/testcase"
)

const (
	OUTPUT_HEADER   = "===================== OUTPUT =====================\n"
	EXPECTED_HEADER = "==================== EXPECTED ====================\n"
	DIFF_HEADER     = "====================== DIFF ======================\n"
)

const FOLDER_NAME = "outputs"

func saveTestResults(results []testcase.Result, doSave bool) {
	if doSave {
		os.Mkdir(FOLDER_NAME, 0777)
		for _, result := range results {
			saveResultToFile(result)
		}
	}
}

func saveResultToFile(result testcase.Result) {
	output := OUTPUT_HEADER + result.Output + "\n"
	expect := EXPECTED_HEADER + result.TestCase.Expect
	file := createFile(fmt.Sprintf("%s/output%d.txt", FOLDER_NAME, result.TestCase.ID))
	defer file.Close()
	if file != nil {
		writeString(file, output)
		writeString(file, expect)
	}
}

func createFile(filePath string) *os.File {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("[WARNING]: failed to create file %s\n", filePath)
		return nil
	}
	return file
}

func writeString(file *os.File, content string) {
	_, err := file.WriteString(content)
	if err != nil {
		fmt.Println("[WARNING]: couldn't write to output text file")
	}
}
