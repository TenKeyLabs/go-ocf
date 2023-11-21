package utils

import (
	"encoding/json"
	"fmt"
	"os"

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

// TODO: Take in file
func ParseOcfResources(folderPath string) (OcfResources, error) {
	// Read the manifest file
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
