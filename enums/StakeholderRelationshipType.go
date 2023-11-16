package enums

// What is the current relationship of the stakeholder to the issuer?
//
// Enumeration of types of relationships between stakeholder and issuer
type StakeholderRelationshipType string

const (
	Advisor                              StakeholderRelationshipType = "ADVISOR"
	BoardMember                          StakeholderRelationshipType = "BOARD_MEMBER"
	Consultant                           StakeholderRelationshipType = "CONSULTANT"
	Employee                             StakeholderRelationshipType = "EMPLOYEE"
	EnumStakeholderRelationshipTypeOTHER StakeholderRelationshipType = "OTHER"
	ExAdvisor                            StakeholderRelationshipType = "EX_ADVISOR"
	ExConsultant                         StakeholderRelationshipType = "EX_CONSULTANT"
	ExEmployee                           StakeholderRelationshipType = "EX_EMPLOYEE"
	Executive                            StakeholderRelationshipType = "EXECUTIVE"
	Founder                              StakeholderRelationshipType = "FOUNDER"
	Investor                             StakeholderRelationshipType = "INVESTOR"
	NonUsEmployee                        StakeholderRelationshipType = "NON_US_EMPLOYEE"
	Officer                              StakeholderRelationshipType = "OFFICER"
)
