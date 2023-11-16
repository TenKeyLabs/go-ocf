package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tenkeylabs/go-ocf/files"
	"github.com/tenkeylabs/go-ocf/objects"
)

type CapTable struct {
	Manifest             files.OCFManifestFile
	StockClasses         []objects.StockClass
	StockPlans           []objects.StockPlan
	StockLegendTemplates []objects.StockLegendTemplate
	VestingTerms         []objects.VestingTerms
	Valuations           []objects.Valuation
	// Transactions         []objects.Transactions
	Stakeholders []objects.Stakeholder
}

func main() {
	// openedFile, _ := ocfZip.Open()
	// defer openedFile.Close()

	// zipFile, _ := os.Create("some.zip")
	// defer zipFile.Close()
	// zipFileSize, _ := io.Copy(zipFile, openedFile)
	// openedFile.Close() //Optimistically close file early?

	// zipReader, _ := zip.NewReader(zipFile, zipFileSize)

	// zipReader.File[0].FileHeader
	// ocfManifestFile, _ := zipReader.Open("Manifest.ocf.json")

	// ocfManifestFile, _ := os.Open("Manifest.ocf.json")

	// ocfManifestFileJson, _ := os.ReadFile("samples/Manifest.ocf.json")
	// stakeholdersFileJson, _ := os.ReadFile("samples/Stakeholders.ocf.json")
	// stockClassesFileJson, _ := os.ReadFile("samples/StockClasses.ocf.json")
	// stockLegendsFileJson, _ := os.ReadFile("samples/StockLegends.ocf.json")
	// stockPlansFileJson, _ := os.ReadFile("samples/StockPlans.ocf.json")
	transactionsFileJson, _ := os.ReadFile("samples/Transactions.ocf.json")
	// valuationsFileJson, _ := os.ReadFile("samples/Valuations.ocf.json")
	// vestingTerms1FileJson, _ := os.ReadFile("samples/VestingTerms.ocf.json")
	// vestingTerms2FileJson, _ := os.ReadFile("samples/VestingTerms.example1.ocf.json")
	// vestingTerms3FileJson, _ := os.ReadFile("samples/VestingTerms.example2.ocf.json")
	// vestingTransactionsFileJson, _ := os.ReadFile("samples/VestingTransactions.examples.ocf.json")

	// var ocfManifest files.OCFManifestFile
	// json.Unmarshal(ocfManifestFileJson, &ocfManifest)
	// ocfManifestJson, _ := json.MarshalIndent(ocfManifest, "", "    ")

	// var stakeholdersFile files.StakeholdersFile
	// json.Unmarshal(stakeholdersFileJson, &stakeholdersFile)
	// stakeholdersJson, _ := json.MarshalIndent(stakeholdersFile, "", "    ")

	// var stockClasses files.StockClassesFile
	// json.Unmarshal(stockClassesFileJson, &stockClasses)
	// stockClassesJson, _ := json.MarshalIndent(stockClasses, "", "    ")

	// var stockLegendsFile files.StockLegendTemplatesFile
	// json.Unmarshal(stockLegendsFileJson, &stockLegendsFile)
	// stockLegendsJson, _ := json.MarshalIndent(stockLegendsFile, "", "    ")

	// var stockPlans files.StockPlansFile
	// json.Unmarshal(stockPlansFileJson, &stockPlans)
	// stockPlansJson, _ := json.MarshalIndent(stockPlans, "", "    ")

	var transactionsFile files.TransactionsFile
	json.Unmarshal(transactionsFileJson, &transactionsFile)
	transactionsJson, _ := json.MarshalIndent(transactionsFile, "", "    ")

	// var valuations files.ValuationsFile
	// json.Unmarshal(valuationsFileJson, &valuations)
	// valuationsJson, _ := json.MarshalIndent(valuations, "", "    ")

	// var vestingTerms1File files.VestingTermsFile
	// json.Unmarshal(vestingTerms1FileJson, &vestingTerms1File)
	// vestingTerms1Json, _ := json.MarshalIndent(vestingTerms1File, "", "    ")

	// var vestingTerms2File files.VestingTermsFile
	// json.Unmarshal(vestingTerms2FileJson, &vestingTerms2File)
	// vestingTerms2Json, _ := json.MarshalIndent(vestingTerms2File, "", "    ")

	// var vestingTerms3File files.VestingTermsFile
	// json.Unmarshal(vestingTerms3FileJson, &vestingTerms3File)
	// vestingTerms3Json, _ := json.MarshalIndent(vestingTerms3File, "", "    ")

	// var vestingTransactionsFile files.TransactionsFile
	// json.Unmarshal(vestingTransactionsFileJson, &vestingTransactionsFile)
	// vestingTransactionsJson, _ := json.MarshalIndent(vestingTransactionsFile, "", "    ")

	// fmt.Printf("Manifest: +%v", string(ocfManifestJson))
	// fmt.Println("\n\n\n\n")
	// fmt.Printf("Stakeholders: +%v", string(stakeholdersJson))
	// fmt.Println("\n\n\n\n")
	// fmt.Printf("StockClasses: +%v", string(stockClassesJson))
	// fmt.Println("\n\n\n\n")
	// fmt.Printf("StockLegends: +%v", string(stockLegendsJson))
	// fmt.Println("\n\n\n\n")
	// fmt.Printf("StockPlans: +%v", string(stockPlansJson))
	// fmt.Println("\n\n\n\n")
	fmt.Printf("%v", string(transactionsJson))
	// fmt.Println("\n\n\n\n")
	// fmt.Printf("Valuations: +%v", string(valuationsJson))
	// fmt.Println("\n\n\n\n")
	// fmt.Printf("VestingTerms1: +%v", string(vestingTerms1Json))
	// fmt.Println("\n\n\n\n")
	// fmt.Printf("VestingTerms2: +%v", string(vestingTerms2Json))
	// fmt.Println("\n\n\n\n")
	// fmt.Printf("VestingTerms3: +%v", string(vestingTerms3Json))
	// fmt.Println("\n\n\n\n")
	// fmt.Printf("VestingTransactions: +%v", string(vestingTransactionsJson))
}
