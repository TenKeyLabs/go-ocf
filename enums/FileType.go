package enums

// File type field (used to select proper schema for validation)
//
// Enumeration of different OCF file types which are used to load proper schemas for
// validation
type FileType string

const (
	OcfManifestFile             FileType = "OCF_MANIFEST_FILE"
	OcfStakeholdersFile         FileType = "OCF_STAKEHOLDERS_FILE"
	OcfStockClassesFile         FileType = "OCF_STOCK_CLASSES_FILE"
	OcfStockLegendTemplatesFile FileType = "OCF_STOCK_LEGEND_TEMPLATES_FILE"
	OcfStockPlansFile           FileType = "OCF_STOCK_PLANS_FILE"
	OcfTransactionsFile         FileType = "OCF_TRANSACTIONS_FILE"
	OcfValuationsFile           FileType = "OCF_VALUATIONS_FILE"
	OcfVestingTermsFile         FileType = "OCF_VESTING_TERMS_FILE"
)
