package enums

// Object type field
//
// Enumeration of object types
type ObjectType string

const (
	//File Types
	ObjectIssuer              ObjectType = "ISSUER"
	ObjectStakeholder         ObjectType = "STAKEHOLDER"
	ObjectStockClass          ObjectType = "STOCK_CLASS"
	ObjectStockLegendTemplate ObjectType = "STOCK_LEGEND_TEMPLATE"
	ObjectStockPlan           ObjectType = "STOCK_PLAN"
	ObjectValuation           ObjectType = "VALUATION"
	ObjectVestingTerms        ObjectType = "VESTING_TERMS"

	//Transaction Types

	//Acceptance
	ObjectTxConvertibleAcceptance        ObjectType = "TX_CONVERTIBLE_ACCEPTANCE"
	ObjectTxEquityCompensationAcceptance ObjectType = "TX_EQUITY_COMPENSATION_ACCEPTANCE"
	ObjectTxPlanSecurityAcceptance       ObjectType = "TX_PLAN_SECURITY_ACCEPTANCE"
	ObjectTxStockAcceptance              ObjectType = "TX_STOCK_ACCEPTANCE"
	ObjectTxWarrantAcceptance            ObjectType = "TX_WARRANT_ACCEPTANCE"

	//Adjustment
	ObjectTxIssuerAuthorizedSharesAdjustment     ObjectType = "TX_ISSUER_AUTHORIZED_SHARES_ADJUSTMENT"
	ObjectTxStockClassAuthorizedSharesAdjustment ObjectType = "TX_STOCK_CLASS_AUTHORIZED_SHARES_ADJUSTMENT"
	ObjectTxStockClassConversionRatioAdjustment  ObjectType = "TX_STOCK_CLASS_CONVERSION_RATIO_ADJUSTMENT"
	ObjectTxStockPlanPoolAdjustment              ObjectType = "TX_STOCK_PLAN_POOL_ADJUSTMENT"

	//Cancellation
	ObjectTxConvertibleCancellation        ObjectType = "TX_CONVERTIBLE_CANCELLATION"
	ObjectTxEquityCompensationCancellation ObjectType = "TX_EQUITY_COMPENSATION_CANCELLATION"
	ObjectTxPlanSecurityCancellation       ObjectType = "TX_PLAN_SECURITY_CANCELLATION"
	ObjectTxStockCancellation              ObjectType = "TX_STOCK_CANCELLATION"
	ObjectTxWarrantCancellation            ObjectType = "TX_WARRANT_CANCELLATION"

	//Conversion
	ObjectTxConvertibleConversion ObjectType = "TX_CONVERTIBLE_CONVERSION"
	ObjectTxStockConversion       ObjectType = "TX_STOCK_CONVERSION"

	//Exercise
	ObjectTxEquityCompensationExercise ObjectType = "TX_EQUITY_COMPENSATION_EXERCISE"
	ObjectTxPlanSecurityExercise       ObjectType = "TX_PLAN_SECURITY_EXERCISE"
	ObjectTxWarrantExercise            ObjectType = "TX_WARRANT_EXERCISE"

	//Issuance
	ObjectTxConvertibleIssuance        ObjectType = "TX_CONVERTIBLE_ISSUANCE"
	ObjectTxEquityCompensationIssuance ObjectType = "TX_EQUITY_COMPENSATION_ISSUANCE"
	ObjectTxPlanSecurityIssuance       ObjectType = "TX_PLAN_SECURITY_ISSUANCE"
	ObjectTxStockIssuance              ObjectType = "TX_STOCK_ISSUANCE"
	ObjectTxWarrantIssuance            ObjectType = "TX_WARRANT_ISSUANCE"

	//Reissuance
	ObjectTxStockReissuance ObjectType = "TX_STOCK_REISSUANCE"

	//Release
	ObjectTxEquityCompensationRelease ObjectType = "TX_EQUITY_COMPENSATION_RELEASE"
	ObjectTxPlanSecurityRelease       ObjectType = "TX_PLAN_SECURITY_RELEASE"

	//Repurchase
	ObjectTxStockRepurchase ObjectType = "TX_STOCK_REPURCHASE"

	//Retraction
	ObjectTxConvertibleRetraction        ObjectType = "TX_CONVERTIBLE_RETRACTION"
	ObjectTxEquityCompensationRetraction ObjectType = "TX_EQUITY_COMPENSATION_RETRACTION"
	ObjectTxPlanSecurityRetraction       ObjectType = "TX_PLAN_SECURITY_RETRACTION"
	ObjectTxStockRetraction              ObjectType = "TX_STOCK_RETRACTION"
	ObjectTxWarrantRetraction            ObjectType = "TX_WARRANT_RETRACTION"

	//ReturnToPool
	ObjectTxStockPlanReturnToPool ObjectType = "TX_STOCK_PLAN_RETURN_TO_POOL"

	//Split
	ObjectTxStockClassSplit ObjectType = "TX_STOCK_CLASS_SPLIT"

	//Transfer
	ObjectTxConvertibleTransfer        ObjectType = "TX_CONVERTIBLE_TRANSFER"
	ObjectTxEquityCompensationTransfer ObjectType = "TX_EQUITY_COMPENSATION_TRANSFER"
	ObjectTxPlanSecurityTransfer       ObjectType = "TX_PLAN_SECURITY_TRANSFER"
	ObjectTxStockTransfer              ObjectType = "TX_STOCK_TRANSFER"
	ObjectTxWarrantTransfer            ObjectType = "TX_WARRANT_TRANSFER"

	//Vesting
	ObjectTxVestingAcceleration ObjectType = "TX_VESTING_ACCELERATION"
	ObjectTxVestingEvent        ObjectType = "TX_VESTING_EVENT"
	ObjectTxVestingStart        ObjectType = "TX_VESTING_START"
)
