package Handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sugarBox/Controller"
)

func HandleAndExecuteOutput() {
	eventOutput := Controller.ProcessOutput()
	outputFile, err := os.OpenFile("output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Println("Error In Closing Output File")
		}
	}(outputFile)

	writer := bufio.NewWriter(outputFile)
	var jsonData []byte

	jsonData, _ = json.MarshalIndent(eventOutput, "", "  ")
	_, err = writer.Write(jsonData)

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error In Writing The Content In Output File")
		return
	}
}
