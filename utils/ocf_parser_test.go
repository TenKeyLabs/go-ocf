package utils

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/benjamintf1/unmarshalledmatchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Safes", func() {
	It("should parse json into go-ocf models", func() {
		folderPath := "../samples/tkl-realistic"
		ocfResources, err := ParseOcfResources(folderPath)

		Expect(err).To(BeNil())
		Expect(ocfResources).ToNot(BeNil())
	})

	It("should marshal go-ocf models into json that matches the original json", func() {
		folderPath := "../samples/tkl-realistic"
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

		Expect(goOCFStockPlansJson).To(MatchUnorderedJSON(originalStockPlansJson))
		Expect(goOCFStockLegendTemplatesJson).To(MatchUnorderedJSON(originalStockLegendTemplatesJson))
		Expect(goOCFStockClassesJson).To(MatchUnorderedJSON(originalStockClassesJson))
		Expect(goOCFTransactionsJson).To(MatchUnorderedJSON(originalTransactionsJson))
		Expect(goOCFStakeholdersJson).To(MatchUnorderedJSON(originalStakeholdersJson))
		Expect(goOCFVestingTermsJson).To(MatchUnorderedJSON(originalVestingTermsJson))
		Expect(goOCFValuationsJson).To(MatchUnorderedJSON(originalValuationsJson))

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
