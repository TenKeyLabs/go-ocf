package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/martian/log"
	"github.com/tenkeylabs/go-ocf/utils"
)

// TODO: Provide CLI to specify folder path for testing?
func main() {
	folderPath := "samples/tkl-realistic"
	ocfResources, err := utils.ParseOcfResources(folderPath)

	if err != nil {
		log.Errorf("error: %w", err)
	} else {
		ocfResourcesJson, _ := json.MarshalIndent(ocfResources, "", "    ")
		fmt.Printf("OCF Resources: +%v", string(ocfResourcesJson))
	}
}
