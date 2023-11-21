package enums

// If the plan security is compensation, what kind?
//
// Enumeration of equity compensation types (there are some things around the margins like
// RSAs that don't currently fit under the EquityCompensation umbrella but might arguably
// fall under this. If you want to create an RSA, create a stock issuance with vesting - you
// can link it to a plan as well, if you want).
//
// **The enums stand for:**
// 1. OPTION_ISO (qualified)
// 2. OPTION_NSO (non-qualified)
// 3. OPTION (not NSO or ISO)
// 4. RSU (restricted share units)
// 5. CSAR(cash-settled stock appreciation rights)
// 6. SSAR(stock-settled stock appreciation rights)
type CompensationType string

const (
	Csar      CompensationType = "CSAR"
	Option    CompensationType = "OPTION"
	OptionISO CompensationType = "OPTION_ISO"
	OptionNso CompensationType = "OPTION_NSO"
	Rsu       CompensationType = "RSU"
	Ssar      CompensationType = "SSAR"
)
