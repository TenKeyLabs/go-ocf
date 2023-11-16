package files

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/objects"
	"github.com/tenkeylabs/go-ocf/types"
)

// Top-level schema describing the OCF Manifest, which holds issuer information and
// references ocf files containing transactions, stakeholders, stock classes, etc.
//
// Abstract file to be extended by all other files
type OCFManifestFile struct {
	// The point-in-time represented by this OCF Package
	AsOf string `json:"as_of"`
	// Unstructured text comments related to and stored for the cap table
	Comments []string `json:"comments,omitempty"`
	// File type field (used to select proper schema for validation)
	FileType enums.FileType `json:"file_type"`
	// Timestamp of when the package was generated. Useful when determining which set of data is
	// most up-to-date, if presented with two packages that have the same `as_of` date, but
	// different cap table data.
	GeneratedAt string `json:"generated_at"`
	// Issuer for the cap table
	Issuer objects.Issuer `json:"issuer"`
	// OCF Version Identifier -- the current semantic version
	// (https://semver.org/spec/v2.0.0.html)
	OcfVersion string `json:"ocf_version"`
	// List of files containing lists of issuer stakeholders, indexed from the file containing
	// the first such object created to the file containing the last (See separate
	// /schema/files/stakeholders_file schema to validate loaded files)
	StakeholdersFiles []types.File `json:"stakeholders_files"`
	// List of files containing lists of issuer stock classes, indexed from the file containing
	// the first such object created to the file containing the last (See separate
	// /schema/files/stock_classes_file schema to validate loaded files)
	StockClassesFiles []types.File `json:"stock_classes_files"`
	// List of files containing lists of issuer stock legend templates, indexed from the file
	// containing the first such object created to the file containing the last (See separate
	// /schema/files/stock_legend_templates_file schema to validate loaded files)
	StockLegendTemplatesFiles []types.File `json:"stock_legend_templates_files"`
	// List of files containing lists of issuer stock plans, indexed from the file containing
	// the first such object created to the file containing the last (See separate
	// /schema/files/stock_plans_file schema to validate loaded files)
	StockPlansFiles []types.File `json:"stock_plans_files"`
	// List of files containing lists of issuer transactions, indexed from the file containing
	// the first such object created to the file containing the last (See separate
	// /schema/files/transactions_file schema to validate loaded files)
	TransactionsFiles []types.File `json:"transactions_files"`
	// List of files containing lists of issuer valuations, indexed from the file containing the
	// first such object created to the file containing the last (See separate
	// /schema/files/valuations_file schema to validate loaded files)
	ValuationsFiles []types.File `json:"valuations_files"`
	// List of files containing lists of issuer vesting terms, indexed from the file containing
	// the first such object created to the file containing the last (See separate
	// /schema/files/vesting_terms_file schema to validate loaded files)
	VestingTermsFiles []types.File `json:"vesting_terms_files"`
}
