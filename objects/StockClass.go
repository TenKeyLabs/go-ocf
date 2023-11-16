package objects

import (
	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/types"
	"github.com/tenkeylabs/go-ocf/types/conversionrights"
)

// Object describing a class of stock issued by the issuer
//
// Abstract object to be extended by all other objects
type StockClass struct {
	// Date on which the board approved the stock class
	BoardApprovalDate *string `json:"board_approval_date,omitempty"`
	// The type of this stock class (e.g. Preferred or Common)
	ClassType enums.StockClassType `json:"class_type"`
	// Unstructured text comments related to and stored for the object
	Comments []string `json:"comments,omitempty"`
	// List of stock class conversion rights possible for this stock class
	ConversionRights []conversionrights.StockClassConversionRights `json:"conversion_rights,omitempty"`
	// Default prefix for certificate numbers in certificated shares (e.g. CS- in CS-1). If
	// certificate IDs have a dash, the prefix should end in the dash like CS-
	DefaultIDPrefix string `json:"default_id_prefix"`
	// Identifier for the object
	ID string `json:"id"`
	// The initial number of shares authorized for this stock class
	InitialSharesAuthorized string `json:"initial_shares_authorized"`
	// The liquidation preference per share for this stock class
	LiquidationPreferenceMultiple *string `json:"liquidation_preference_multiple,omitempty"`
	// Name for the stock type (e.g. Series A Preferred or Class A Common)
	Name string `json:"name"`
	// Object type field
	ObjectType enums.ObjectType `json:"object_type"`
	// Per-share par value of this stock class
	ParValue *types.Monetary `json:"par_value,omitempty"`
	// The participation cap multiple per share for this stock class
	ParticipationCapMultiple *string `json:"participation_cap_multiple,omitempty"`
	// Per-share price this stock class was issued for
	PricePerShare *types.Monetary `json:"price_per_share,omitempty"`
	// Seniority of the stock - determines repayment priority. Seniority is ordered by
	// increasing number so that stock classes with a higher seniority have higher repayment
	// priority. The following properties hold for all stock classes for a given company:
	// a) transitivity: stock classes are absolutely stackable by seniority and in increasing
	// numerical order,
	// b) non-uniqueness: multiple stock classes can have the same Seniority number and
	// therefore have the same liquidation/repayment order.
	// In practice, stock classes with same seniority may be created at different points in time
	// and (for example, an extension of an existing preferred financing round), and also a new
	// stock class can be created with seniority between two existing stock classes, in which
	// case it is assigned some decimal number between the numbers representing seniority of the
	// respective classes.
	Seniority string `json:"seniority"`
	// Date on which the stockholders approved the stock class
	StockholderApprovalDate *string `json:"stockholder_approval_date,omitempty"`
	// The number of votes each share of this stock class gets
	VotesPerShare string `json:"votes_per_share"`
}
