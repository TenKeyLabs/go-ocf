package enums

// Distinguish individuals from institutions
//
// Enumeration of stakeholder types - individual (human) or institution (entity)
type StakeholderType string

const (
	Individual  StakeholderType = "INDIVIDUAL"
	Institution StakeholderType = "INSTITUTION"
)
