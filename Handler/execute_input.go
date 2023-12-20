package Handler

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sugarBox/Controller"
	"sugarBox/Payload"
)

func HandleAndExecuteInput() {
	var eventCreationReq []Payload.EventCreationReq

	inputFile := flag.String("i", "", "Input JSON file")
	outputFile := flag.String("o", "", "Output JSON file")

	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Usage: ./aggregate_events -i input.json -o output.json")
		return
	}

	file, _ := os.Open(*inputFile)
	var fileContent string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent = fileContent + scanner.Text()
	}

	err := json.Unmarshal([]byte(fileContent), &eventCreationReq)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for _, req := range eventCreationReq {
		err := Controller.ProcessInput(req)
		if err != nil {
			fmt.Println("err", err)
		}
	}

	HandleAndExecuteOutput()

}
