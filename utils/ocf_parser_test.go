package utils

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var _ = Describe("Safes", func() {
	folderPath := "../samples/ocf"
	format.MaxLength = 100000

	It("should parse json into go-ocf models", func() {
		ocfResources, err := ParseOcfResources(folderPath)

		Expect(err).To(BeNil())
		Expect(ocfResources).ToNot(BeNil())
	})

	It("should marshal go-ocf models into json that matches the original json", func() {
		ocfResources, err := ParseOcfResources(folderPath)
		Expect(err).To(BeNil())

		originalStockPlansJson, err := getFileItemsJson(folderPath, "stock_plans.ocf.json")
		Expect(err).To(BeNil())
		goOCFStockPlansJson, err := json.Marshal(ocfResources.StockPlans)
		Expect(err).To(BeNil())

		originalStockLegendTemplatesJson, err := getFileItemsJson(folderPath, "stock_legends.ocf.json")
		Expect(err).To(BeNil())
		goOCFStockLegendTemplatesJson, err := json.Marshal(ocfResources.StockLegendTemplates)
		Expect(err).To(BeNil())

		originalStockClassesJson, err := getFileItemsJson(folderPath, "stock_classes.ocf.json")
		Expect(err).To(BeNil())
		goOCFStockClassesJson, err := json.Marshal(ocfResources.StockClasses)
		Expect(err).To(BeNil())

		originalTransactionsJson, err := getFileItemsJson(folderPath, "transactions.ocf.json")
		Expect(err).To(BeNil())
		goOCFTransactionsJson, err := json.MarshalIndent(ocfResources.Transactions, "", "    ")
		Expect(err).To(BeNil())

		originalStakeholdersJson, err := getFileItemsJson(folderPath, "stakeholders.ocf.json")
		Expect(err).To(BeNil())
		goOCFStakeholdersJson, err := json.Marshal(ocfResources.Stakeholders)
		Expect(err).To(BeNil())

		originalVestingTermsJson, err := getFileItemsJson(folderPath, "vesting_terms.ocf.json")
		Expect(err).To(BeNil())
		goOCFVestingTermsJson, err := json.MarshalIndent(ocfResources.VestingTerms, "", "    ")
		Expect(err).To(BeNil())

		originalValuationsJson, err := getFileItemsJson(folderPath, "valuations.ocf.json")
		Expect(err).To(BeNil())
		goOCFValuationsJson, err := json.Marshal(ocfResources.Valuations)
		Expect(err).To(BeNil())

		Expect(goOCFStockPlansJson).To(MatchJSON(originalStockPlansJson))
		Expect(goOCFStockLegendTemplatesJson).To(MatchJSON(originalStockLegendTemplatesJson))
		Expect(goOCFStockClassesJson).To(MatchJSON(originalStockClassesJson))
		Expect(goOCFTransactionsJson).To(MatchJSON(originalTransactionsJson))
		Expect(goOCFStakeholdersJson).To(MatchJSON(originalStakeholdersJson))
		Expect(goOCFVestingTermsJson).To(MatchJSON(originalVestingTermsJson))
		Expect(goOCFValuationsJson).To(MatchJSON(originalValuationsJson))

		Expect(err).To(BeNil())
	})
})

func getFileItemsJson(folderPath string, filePath string) ([]byte, error) {
	fileJson, err := os.ReadFile(fmt.Sprintf("%v/%v", folderPath, filePath))
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}

	var fileContent ocfFileContent[map[string]interface{}]
	err = json.Unmarshal(fileJson, &fileContent)
	if err != nil {
		return nil, fmt.Errorf("unable unmarshal into generic file: %w", err)
	}

	items, err := json.Marshal(fileContent.Items)
	if err != nil {
		return nil, fmt.Errorf("unable marshal file items: %w", err)
	}
	return items, nil
}
