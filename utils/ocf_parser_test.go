package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOcfParser(t *testing.T) {
	folderPath := "../samples/tkl-realistic"

	originalStockPlansJson, err := getFileItemsJson(folderPath, "stock_plans.ocf.json")
	if err != nil {
		t.Errorf("error: %v: ", err)
	}

	originalStockLegendTemplatesJson, err := getFileItemsJson(folderPath, "stock_legends.ocf.json")
	if err != nil {
		t.Errorf("error: %v: ", err)
	}

	originalStockClassesJson, err := getFileItemsJson(folderPath, "stock_classes.ocf.json")
	if err != nil {
		t.Errorf("error: %v: ", err)
	}

	originalTransactionsJson, err := getFileItemsJson(folderPath, "transactions.ocf.json")
	if err != nil {
		t.Errorf("error: %v: ", err)
	}

	originalStakeholdersJson, err := getFileItemsJson(folderPath, "stakeholders.ocf.json")
	if err != nil {
		t.Errorf("error: %v: ", err)
	}

	originalVestingTermsJson, err := getFileItemsJson(folderPath, "vesting_terms.ocf.json")
	if err != nil {
		t.Errorf("error: %v: ", err)
	}

	originalValuationsJson, err := getFileItemsJson(folderPath, "valuations.ocf.json")
	if err != nil {
		t.Errorf("error: %v: ", err)
	}

	ocfResources, err := ParseOcfResources(folderPath)
	if err != nil {
		t.Errorf("error: %v: ", err)
	}

	ocfStockPlansJson, _ := json.Marshal(ocfResources.StockPlans)
	ocfStockLegendTemplatesJson, _ := json.Marshal(ocfResources.StockLegendTemplates)
	ocfStockClassesJson, _ := json.Marshal(ocfResources.StockClasses)
	ocfTransactionsJson, _ := json.Marshal(ocfResources.Transactions)
	ocfStakeholdersJson, _ := json.Marshal(ocfResources.Stakeholders)
	ocfVestingTermsJson, _ := json.Marshal(ocfResources.VestingTerms)
	ocfValuationsJson, _ := json.Marshal(ocfResources.Valuations)

  jsondiff.CompareJSON(originalStockPlansJson, ocfStockPlansJson)
	stockPlansDiff := cmp.Diff(string(originalStockPlansJson), string(ocfStockPlansJson))
	fmt.Printf("Original Stock Plans: %v\n\n", string(originalStockPlansJson))
	fmt.Printf("OCF Stock Plans: %v\n\n", string(ocfStockPlansJson))
	stockLegendTemplatesDiff := cmp.Diff(string(originalStockLegendTemplatesJson), string(ocfStockLegendTemplatesJson))
	stockClassesDiff := cmp.Diff(string(originalStockClassesJson), string(ocfStockClassesJson))
	transactionsDiff := cmp.Diff(string(originalTransactionsJson), string(ocfTransactionsJson))
	stakeholdersDiff := cmp.Diff(string(originalStakeholdersJson), string(ocfStakeholdersJson))
	vestingTermsDiff := cmp.Diff(string(originalVestingTermsJson), string(ocfVestingTermsJson))
	valuationsDiff := cmp.Diff(string(originalValuationsJson), string(ocfValuationsJson))

	// t.Logf("stockPlansDiff: %s", stockPlansDiff)
	// t.Logf("stockLegendTemplatesDiff: %s", stockLegendTemplatesDiff)
	// t.Logf("stockClassesDiff: %s", stockClassesDiff)
	// t.Logf("transactionsDiff: %s", transactionsDiff)
	// t.Logf("stakeholdersDiff: %s", stakeholdersDiff)
	// t.Logf("vestingTermsDiff: %s", vestingTermsDiff)
	// t.Logf("valuationDiff: %s", valuationsDiff)

	if stockPlansDiff != "" {
		t.Errorf("stockPlansDiff: %s", stockPlansDiff)
		return
	}
	if stockLegendTemplatesDiff != "" {
		t.Errorf("stockLegendTemplatesDiff: %s", stockLegendTemplatesDiff)
		return
	}
	if stockClassesDiff != "" {
		t.Errorf("stockClassesDiff: %s", stockClassesDiff)
		return
	}
	if transactionsDiff != "" {
		t.Errorf("transactionsDiff: %s", transactionsDiff)
		return
	}
	if stakeholdersDiff != "" {
		t.Errorf("stakeholdersDiff: %s", stakeholdersDiff)
		return
	}
	if vestingTermsDiff != "" {
		t.Errorf("vestingTermsDiff: %s", vestingTermsDiff)
		return
	}
	if valuationsDiff != "" {
		t.Errorf("valuationDiff: %s", valuationsDiff)
		return
	}
}

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

	pretty, _ := json.MarshalIndent(fileContent.Items, "", "    ")
	fmt.Printf("Pretty File: %v\n\n", string(pretty))

	items, err := json.Marshal(fileContent.Items)
	if err != nil {
		return nil, fmt.Errorf("unable marshal file items: %w", err)
	}
	return items, nil
}
