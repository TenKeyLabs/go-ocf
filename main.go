package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/martian/log"
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/files"
	"github.com/tenkeylabs/go-ocf/objects"
	"github.com/tenkeylabs/go-ocf/types"
)

type OcfResources struct {
	Stakeholders         []objects.Stakeholder
	StockClasses         []objects.StockClass
	StockLegendTemplates []objects.StockLegendTemplate
	StockPlans           []objects.StockPlan
	Transactions         []objects.Transaction
	Valuations           []objects.Valuation
	VestingTerms         []objects.VestingTerms
}

type ocfFileContent[T any] struct {
	FileType enums.FileType `json:"file_type"`
	Items    []T            `json:"items"`
}

func parseOcfResources() (OcfResources, error) {
	// Read the manifest file
	folderPath := "samples/tkl-realistic"
	ocfManifestFileJson, err := os.ReadFile(fmt.Sprintf("%v/manifest.ocf.json", folderPath))
	if err != nil {
		return OcfResources{}, fmt.Errorf("unable to parse ocf manifest: %w", err)
	}
	var ocfManifest files.OCFManifestFile
	json.Unmarshal(ocfManifestFileJson, &ocfManifest)

	stakeholders, err := parseOcfResource[objects.Stakeholder](folderPath, ocfManifest.StakeholdersFiles)
	if err != nil {
		return OcfResources{}, fmt.Errorf("unable to parse ocf stakeholders: %w", err)
	}

	stockClasses, err := parseOcfResource[objects.StockClass](folderPath, ocfManifest.StockClassesFiles)
	if err != nil {
		return OcfResources{}, fmt.Errorf("unable to parse ocf stockClasses: %w", err)
	}

	stockLegendTemplates, err := parseOcfResource[objects.StockLegendTemplate](folderPath, ocfManifest.StockLegendTemplatesFiles)
	if err != nil {
		return OcfResources{}, fmt.Errorf("unable to parse ocf stockLegendTemplates: %w", err)
	}

	stockPlans, err := parseOcfResource[objects.StockPlan](folderPath, ocfManifest.StockPlansFiles)
	if err != nil {
		return OcfResources{}, fmt.Errorf("unable to parse ocf stockPlans: %w", err)
	}

	transactions, err := parseOcfResource[objects.Transaction](folderPath, ocfManifest.TransactionsFiles)
	if err != nil {
		return OcfResources{}, fmt.Errorf("unable to parse ocf transactions: %w", err)
	}

	valuations, err := parseOcfResource[objects.Valuation](folderPath, ocfManifest.ValuationsFiles)
	if err != nil {
		return OcfResources{}, fmt.Errorf("unable to parse ocf valuations: %w", err)
	}

	vestingTerms, err := parseOcfResource[objects.VestingTerms](folderPath, ocfManifest.VestingTermsFiles)
	if err != nil {
		return OcfResources{}, fmt.Errorf("unable to parse ocf vestingTerms: %w", err)
	}

	return OcfResources{
		Stakeholders:         stakeholders,
		StockClasses:         stockClasses,
		StockLegendTemplates: stockLegendTemplates,
		StockPlans:           stockPlans,
		Transactions:         transactions,
		Valuations:           valuations,
		VestingTerms:         vestingTerms,
	}, nil
}

func parseOcfResource[T any](folderPath string, filesInfo []types.File) ([]T, error) {
	var items []T
	for _, fileInfo := range filesInfo {
		file, err := decodeOcfFile[ocfFileContent[T]](fmt.Sprintf("%v/%v", folderPath, fileInfo.Filepath))
		if err != nil {
			return nil, fmt.Errorf("failed to decode file %q: %w", fileInfo.Filepath, err)
		}
		items = append(items, file.Items...)
	}

	return items, nil
}

func decodeOcfFile[T any](fileName string) (T, error) {
	var response T
	fileToDecode, err := os.Open(fileName)
	if err != nil {
		return response, fmt.Errorf("failed to open file %q: %w", fileName, err)
	}
	defer fileToDecode.Close()

	// Do not attempt to read fileToDecode for debugging since that will read through all the buffer data and the json decoding will produce an empty object
	decoder := json.NewDecoder(fileToDecode)
	err = decoder.Decode(&response)
	if err != nil {
		return response, fmt.Errorf("unable to decode file %q: %w", fileName, err)
	}

	return response, nil
}

func main() {
	ocfResources, err := parseOcfResources()

	if err != nil {
		log.Errorf("error: %w", err)
	} else {
		// var vestingTerms []objects.VestingTerms
		fmt.Printf("Parsed Vesting Terms: %v", ocfResources.VestingTerms[0])
		vestingTerms, _ := json.MarshalIndent(ocfResources.VestingTerms, "", "    ")
		log.Infof("%v", string(vestingTerms))
	}

	// fmt.Printf("\n\n")
	// vestingTermsFileJson, err := os.ReadFile(fmt.Sprintf("%v/%v", manifestFolder, ocfManifest.VestingTermsFiles[0].Filepath))
	// var vestingTerms files.VestingTermsFile
	// json.Unmarshal(vestingTermsFileJson, &vestingTerms)
	// vestingTermsJson, _ := json.MarshalIndent(vestingTerms, "", "    ")
	// fmt.Printf("Vesting Terms: %v", string(vestingTermsJson))

	// zipReader, err := zip.NewReader(file, ocfFileHeader.Size)
	// if err != nil {
	// 	return OcfImportSummaryInfo{}, fmt.Errorf("failed to open ocf zip file: %w", err)
	// }

	// Parse the manifest file
	// ocfManifest, err := decodeOcfFile[ocfFiles.OCFManifestFile](zipReader, ocfManifestFileName)
	// if err != nil {
	// 	log.Error("unable to parse ocf manifest: %w", err)
	// 	return
	// }

	// ocfResources, err := parseOcfResources(zipReader, ocfManifest)
	// if err != nil {
	// 	log.Error("unable to parse ocf resources: %w", err)
	// 	return
	// }
	// log.Info("OCF Resources", zap.Any("ocf resources", ocfResources))

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

	// stakeholdersFileJson, _ := os.ReadFile("samples/Stakeholders.ocf.json")
	// stockClassesFileJson, _ := os.ReadFile("samples/StockClasses.ocf.json")
	// stockLegendsFileJson, _ := os.ReadFile("samples/StockLegends.ocf.json")
	// stockPlansFileJson, _ := os.ReadFile("samples/StockPlans.ocf.json")
	// transactionsFileJson, _ := os.ReadFile("samples/Transactions.ocf.json")
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

	// var transactionsFile files.TransactionsFile
	// json.Unmarshal(transactionsFileJson, &transactionsFile)
	// transactionsJson, _ := json.MarshalIndent(transactionsFile, "", "    ")

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
	// fmt.Printf("%v", string(transactionsJson))
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
