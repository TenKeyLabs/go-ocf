package objects

import (
	"encoding/json"
	"fmt"
	"maps"

	"github.com/tenkeylabs/go-ocf/enums"
	"github.com/tenkeylabs/go-ocf/objects/transactions/acceptance"
	"github.com/tenkeylabs/go-ocf/objects/transactions/adjustment"
	"github.com/tenkeylabs/go-ocf/objects/transactions/cancellation"
	"github.com/tenkeylabs/go-ocf/objects/transactions/conversion"
	"github.com/tenkeylabs/go-ocf/objects/transactions/exercise"
	"github.com/tenkeylabs/go-ocf/objects/transactions/issuance"
	"github.com/tenkeylabs/go-ocf/objects/transactions/reissuance"
	"github.com/tenkeylabs/go-ocf/objects/transactions/release"
	"github.com/tenkeylabs/go-ocf/objects/transactions/repurchase"
	"github.com/tenkeylabs/go-ocf/objects/transactions/retraction"
	"github.com/tenkeylabs/go-ocf/objects/transactions/returntopool"
	"github.com/tenkeylabs/go-ocf/objects/transactions/split"
	"github.com/tenkeylabs/go-ocf/objects/transactions/transfer"
	"github.com/tenkeylabs/go-ocf/objects/transactions/vesting"
	"github.com/tenkeylabs/go-ocf/primitives/objects"
	"github.com/tenkeylabs/go-ocf/primitives/objects/transactions"
)

type Transaction struct {
	objects.Object
	transactions.Transaction

	//Acceptance
	ConvertibleAcceptance        *acceptance.ConvertibleAcceptance
	EquityCompensationAcceptance *acceptance.EquityCompensationAcceptance
	PlanSecurityAcceptance       *acceptance.PlanSecurityAcceptance
	StockAcceptance              *acceptance.StockAcceptance
	WarrantAcceptance            *acceptance.WarrantAcceptance

	//Adjustment
	IssuerAuthorizedSharesAdjustment     *adjustment.IssuerAuthorizedSharesAdjustment
	StockClassAuthorizedSharesAdjustment *adjustment.StockClassAuthorizedSharesAdjustment
	StockClassConversionRatioAdjustment  *adjustment.StockClassConversionRatioAdjustment
	StockPlanPoolAdjustment              *adjustment.StockPlanPoolAdjustment

	//Cancellation
	ConvertibleCancellation        *cancellation.ConvertibleCancellation
	EquityCompensationCancellation *cancellation.EquityCompensationCancellation
	PlanSecurityCancellation       *cancellation.PlanSecurityCancellation
	StockCancellation              *cancellation.StockCancellation
	WarrantCancellation            *cancellation.WarrantCancellation

	//Conversion
	ConvertibleConversion *conversion.ConvertibleConversion
	StockConversion       *conversion.StockConversion

	//Exercise
	EquityCompensationExercise *exercise.EquityCompensationExercise
	PlanSecurityExercise       *exercise.PlanSecurityExercise
	WarrantExercise            *exercise.WarrantExercise

	//Issuance
	ConvertibleIssuance        *issuance.ConvertibleIssuance
	EquityCompensationIssuance *issuance.EquityCompensationIssuance
	PlanSecurityIssuance       *issuance.PlanSecurityIssuance
	StockIssuance              *issuance.StockIssuance
	WarrantIssuance            *issuance.WarrantIssuance

	//Reissuance
	StockReissuance *reissuance.StockReissuance

	//Release
	EquityCompensationRelease *release.EquityCompensationRelease
	PlanSecurityRelease       *release.PlanSecurityRelease

	//Repurchase
	StockRepurchase *repurchase.StockRepurchase

	//Retraction
	ConvertibleRetraction        *retraction.ConvertibleRetraction
	EquityCompensationRetraction *retraction.EquityCompensationRetraction
	PlanSecurityRetraction       *retraction.PlanSecurityRetraction
	StockRetraction              *retraction.StockRetraction
	WarrantRetraction            *retraction.WarrantRetraction

	//ReturnToPool
	StockPlanReturnToPool *returntopool.StockPlanReturnToPool

	//Split
	StockClassSplit *split.StockClassSplit

	//Transfer
	ConvertibleTransfer        *transfer.ConvertibleTransfer
	EquityCompensationTransfer *transfer.EquityCompensationTransfer
	PlanSecurityTransfer       *transfer.PlanSecurityTransfer
	StockTransfer              *transfer.StockTransfer
	WarrantTransfer            *transfer.WarrantTransfer

	//Vesting
	VestingAcceleration *vesting.VestingAcceleration
	VestingEvent        *vesting.VestingEvent
	VestingStart        *vesting.VestingStart
}

type MarshalStruct struct {
	objects.Object
	transactions.Transaction
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	aliasTransaction := MarshalStruct{}
	if err := json.Unmarshal(data, &aliasTransaction); err != nil {
		return err
	}

	t.Object = aliasTransaction.Object
	t.Transaction = aliasTransaction.Transaction

	var err error
	switch t.ObjectType {
	//Acceptance
	case enums.ObjectTxConvertibleAcceptance:
		var convertibleAcceptance acceptance.ConvertibleAcceptance
		err = json.Unmarshal(data, &convertibleAcceptance)
		t.ConvertibleAcceptance = &convertibleAcceptance
	case enums.ObjectTxEquityCompensationAcceptance:
		var equityCompensationAcceptance acceptance.EquityCompensationAcceptance
		err = json.Unmarshal(data, &equityCompensationAcceptance)
		t.EquityCompensationAcceptance = &equityCompensationAcceptance
	case enums.ObjectTxPlanSecurityAcceptance:
		var planSecurityAcceptance acceptance.PlanSecurityAcceptance
		err = json.Unmarshal(data, &planSecurityAcceptance)
		t.PlanSecurityAcceptance = &planSecurityAcceptance
	case enums.ObjectTxStockAcceptance:
		var stockAcceptance acceptance.StockAcceptance
		err = json.Unmarshal(data, &stockAcceptance)
		t.StockAcceptance = &stockAcceptance
	case enums.ObjectTxWarrantAcceptance:
		var warrantAcceptance acceptance.WarrantAcceptance
		err = json.Unmarshal(data, &warrantAcceptance)
		t.WarrantAcceptance = &warrantAcceptance

	//Adjustment
	case enums.ObjectTxIssuerAuthorizedSharesAdjustment:
		var issuerAuthorizedSharesAdjustment adjustment.IssuerAuthorizedSharesAdjustment
		err = json.Unmarshal(data, &issuerAuthorizedSharesAdjustment)
		t.IssuerAuthorizedSharesAdjustment = &issuerAuthorizedSharesAdjustment
	case enums.ObjectTxStockClassAuthorizedSharesAdjustment:
		var stockClassAuthorizedSharesAdjustment adjustment.StockClassAuthorizedSharesAdjustment
		err = json.Unmarshal(data, &stockClassAuthorizedSharesAdjustment)
		t.StockClassAuthorizedSharesAdjustment = &stockClassAuthorizedSharesAdjustment
	case enums.ObjectTxStockClassConversionRatioAdjustment:
		var stockClassConversionRatioAdjustment adjustment.StockClassConversionRatioAdjustment
		err = json.Unmarshal(data, &stockClassConversionRatioAdjustment)
		t.StockClassConversionRatioAdjustment = &stockClassConversionRatioAdjustment
	case enums.ObjectTxStockPlanPoolAdjustment:
		var stockPlanPoolAdjustment adjustment.StockPlanPoolAdjustment
		err = json.Unmarshal(data, &stockPlanPoolAdjustment)
		t.StockPlanPoolAdjustment = &stockPlanPoolAdjustment

	//Cancellation
	case enums.ObjectTxConvertibleCancellation:
		var convertibleCancellation cancellation.ConvertibleCancellation
		err = json.Unmarshal(data, &convertibleCancellation)
		t.ConvertibleCancellation = &convertibleCancellation
	case enums.ObjectTxEquityCompensationCancellation:
		var equityCompensationCancellation cancellation.EquityCompensationCancellation
		err = json.Unmarshal(data, &equityCompensationCancellation)
		t.EquityCompensationCancellation = &equityCompensationCancellation
	case enums.ObjectTxPlanSecurityCancellation:
		var planSecurityCancellation cancellation.PlanSecurityCancellation
		err = json.Unmarshal(data, &planSecurityCancellation)
		t.PlanSecurityCancellation = &planSecurityCancellation
	case enums.ObjectTxStockCancellation:
		var stockCancellation cancellation.StockCancellation
		err = json.Unmarshal(data, &stockCancellation)
		t.StockCancellation = &stockCancellation
	case enums.ObjectTxWarrantCancellation:
		var warrantCancellation cancellation.WarrantCancellation
		err = json.Unmarshal(data, &warrantCancellation)
		t.WarrantCancellation = &warrantCancellation

	//Conversion
	case enums.ObjectTxConvertibleConversion:
		var convertibleConversion conversion.ConvertibleConversion
		err = json.Unmarshal(data, &convertibleConversion)
		t.ConvertibleConversion = &convertibleConversion
	case enums.ObjectTxStockConversion:
		var stockConversion conversion.StockConversion
		err = json.Unmarshal(data, &stockConversion)
		t.StockConversion = &stockConversion

	//Exercise
	case enums.ObjectTxEquityCompensationExercise:
		var equityCompensationExercise exercise.EquityCompensationExercise
		err = json.Unmarshal(data, &equityCompensationExercise)
		t.EquityCompensationExercise = &equityCompensationExercise
	case enums.ObjectTxPlanSecurityExercise:
		var planSecurityExercise exercise.PlanSecurityExercise
		err = json.Unmarshal(data, &planSecurityExercise)
		t.PlanSecurityExercise = &planSecurityExercise
	case enums.ObjectTxWarrantExercise:
		var warrantExercise exercise.WarrantExercise
		err = json.Unmarshal(data, &warrantExercise)
		t.WarrantExercise = &warrantExercise

	//Issuance
	case enums.ObjectTxStockIssuance:
		var stockIssuance issuance.StockIssuance
		err = json.Unmarshal(data, &stockIssuance)
		t.StockIssuance = &stockIssuance
	case enums.ObjectTxConvertibleIssuance:
		var convertibleIssuance issuance.ConvertibleIssuance
		err = json.Unmarshal(data, &convertibleIssuance)
		t.ConvertibleIssuance = &convertibleIssuance
	case enums.ObjectTxPlanSecurityIssuance:
		var planSecurityIssuance issuance.PlanSecurityIssuance
		err = json.Unmarshal(data, &planSecurityIssuance)
		t.PlanSecurityIssuance = &planSecurityIssuance
	case enums.ObjectTxEquityCompensationIssuance:
		var equityCompensationIssuance issuance.EquityCompensationIssuance
		err = json.Unmarshal(data, &equityCompensationIssuance)
		t.EquityCompensationIssuance = &equityCompensationIssuance
	case enums.ObjectTxWarrantIssuance:
		var warrantIssuance issuance.WarrantIssuance
		err = json.Unmarshal(data, &warrantIssuance)
		t.WarrantIssuance = &warrantIssuance

	//Reissuance
	case enums.ObjectTxStockReissuance:
		var stockReissuance reissuance.StockReissuance
		err = json.Unmarshal(data, &stockReissuance)
		t.StockReissuance = &stockReissuance

	//Release
	case enums.ObjectTxEquityCompensationRelease:
		var equityCompensationRelease release.EquityCompensationRelease
		err = json.Unmarshal(data, &equityCompensationRelease)
		t.EquityCompensationRelease = &equityCompensationRelease
	case enums.ObjectTxPlanSecurityRelease:
		var planSecurityRelease release.PlanSecurityRelease
		err = json.Unmarshal(data, &planSecurityRelease)
		t.PlanSecurityRelease = &planSecurityRelease

	//Repurchase
	case enums.ObjectTxStockRepurchase:
		var stockRepurchase repurchase.StockRepurchase
		err = json.Unmarshal(data, &stockRepurchase)
		t.StockRepurchase = &stockRepurchase

	//Retraction
	case enums.ObjectTxConvertibleRetraction:
		var convertibleRetraction retraction.ConvertibleRetraction
		err = json.Unmarshal(data, &convertibleRetraction)
		t.ConvertibleRetraction = &convertibleRetraction
	case enums.ObjectTxEquityCompensationRetraction:
		var equityCompensationRetraction retraction.EquityCompensationRetraction
		err = json.Unmarshal(data, &equityCompensationRetraction)
		t.EquityCompensationRetraction = &equityCompensationRetraction
	case enums.ObjectTxPlanSecurityRetraction:
		var planSecurityRetraction retraction.PlanSecurityRetraction
		err = json.Unmarshal(data, &planSecurityRetraction)
		t.PlanSecurityRetraction = &planSecurityRetraction
	case enums.ObjectTxStockRetraction:
		var stockRetraction retraction.StockRetraction
		err = json.Unmarshal(data, &stockRetraction)
		t.StockRetraction = &stockRetraction
	case enums.ObjectTxWarrantRetraction:
		var warrantRetraction retraction.WarrantRetraction
		err = json.Unmarshal(data, &warrantRetraction)
		t.WarrantRetraction = &warrantRetraction

	//ReturnToPool
	case enums.ObjectTxStockPlanReturnToPool:
		var stockPlanReturnToPool returntopool.StockPlanReturnToPool
		err = json.Unmarshal(data, &stockPlanReturnToPool)
		t.StockPlanReturnToPool = &stockPlanReturnToPool

	//Split
	case enums.ObjectTxStockClassSplit:
		var stockClassSplit split.StockClassSplit
		err = json.Unmarshal(data, &stockClassSplit)
		t.StockClassSplit = &stockClassSplit

	//Transfer
	case enums.ObjectTxConvertibleTransfer:
		var convertibleTransfer transfer.ConvertibleTransfer
		err = json.Unmarshal(data, &convertibleTransfer)
		t.ConvertibleTransfer = &convertibleTransfer
	case enums.ObjectTxEquityCompensationTransfer:
		var equityCompensationTransfer transfer.EquityCompensationTransfer
		err = json.Unmarshal(data, &equityCompensationTransfer)
		t.EquityCompensationTransfer = &equityCompensationTransfer
	case enums.ObjectTxPlanSecurityTransfer:
		var planSecurityTransfer transfer.PlanSecurityTransfer
		err = json.Unmarshal(data, &planSecurityTransfer)
		t.PlanSecurityTransfer = &planSecurityTransfer
	case enums.ObjectTxStockTransfer:
		var stockTransfer transfer.StockTransfer
		err = json.Unmarshal(data, &stockTransfer)
		t.StockTransfer = &stockTransfer
	case enums.ObjectTxWarrantTransfer:
		var warrantTransfer transfer.WarrantTransfer
		err = json.Unmarshal(data, &warrantTransfer)
		t.WarrantTransfer = &warrantTransfer

	//Vesting
	case enums.ObjectTxVestingAcceleration:
		var vestingAcceleration vesting.VestingAcceleration
		err = json.Unmarshal(data, &vestingAcceleration)
		t.VestingAcceleration = &vestingAcceleration
	case enums.ObjectTxVestingEvent:
		var vestingEvent vesting.VestingEvent
		err = json.Unmarshal(data, &vestingEvent)
		t.VestingEvent = &vestingEvent
	case enums.ObjectTxVestingStart:
		var vestingStart vesting.VestingStart
		err = json.Unmarshal(data, &vestingStart)
		t.VestingStart = &vestingStart
	}

	return err
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	transactionData, err := t.getTransactionData()
	if err != nil {
		return nil, err
	}

	dataMap := map[string]any{}
	err = json.Unmarshal(transactionData, &dataMap)
	if err != nil {
		return nil, err
	}

	txWrapper := MarshalStruct{
		Object:      t.Object,
		Transaction: t.Transaction,
	}

	txData, err := json.Marshal(txWrapper)
	if err != nil {
		return nil, err
	}

	txMap := map[string]any{}
	err = json.Unmarshal(txData, &txMap)
	if err != nil {
		return nil, err
	}

	maps.Copy(txMap, dataMap)

	return json.Marshal(txMap)
}

func (t *Transaction) getTransactionData() ([]byte, error) {
	switch t.ObjectType {
	//Acceptance
	case enums.ObjectTxConvertibleAcceptance:
		return json.Marshal(t.ConvertibleAcceptance)
	case enums.ObjectTxEquityCompensationAcceptance:
		return json.Marshal(t.EquityCompensationAcceptance)
	case enums.ObjectTxPlanSecurityAcceptance:
		return json.Marshal(t.PlanSecurityAcceptance)
	case enums.ObjectTxStockAcceptance:
		return json.Marshal(t.StockAcceptance)
	case enums.ObjectTxWarrantAcceptance:
		return json.Marshal(t.WarrantAcceptance)

	//Adjustment
	case enums.ObjectTxIssuerAuthorizedSharesAdjustment:
		return json.Marshal(t.IssuerAuthorizedSharesAdjustment)
	case enums.ObjectTxStockClassAuthorizedSharesAdjustment:
		return json.Marshal(t.StockClassAuthorizedSharesAdjustment)
	case enums.ObjectTxStockClassConversionRatioAdjustment:
		return json.Marshal(t.StockClassConversionRatioAdjustment)
	case enums.ObjectTxStockPlanPoolAdjustment:
		return json.Marshal(t.StockPlanPoolAdjustment)

	//Cancellation
	case enums.ObjectTxConvertibleCancellation:
		return json.Marshal(t.ConvertibleCancellation)
	case enums.ObjectTxEquityCompensationCancellation:
		return json.Marshal(t.EquityCompensationCancellation)
	case enums.ObjectTxPlanSecurityCancellation:
		return json.Marshal(t.PlanSecurityCancellation)
	case enums.ObjectTxStockCancellation:
		return json.Marshal(t.StockCancellation)
	case enums.ObjectTxWarrantCancellation:
		return json.Marshal(t.WarrantCancellation)

	//Conversion
	case enums.ObjectTxConvertibleConversion:
		return json.Marshal(t.ConvertibleConversion)
	case enums.ObjectTxStockConversion:
		return json.Marshal(t.StockConversion)

	//Exercise
	case enums.ObjectTxEquityCompensationExercise:
		return json.Marshal(t.EquityCompensationExercise)
	case enums.ObjectTxPlanSecurityExercise:
		return json.Marshal(t.PlanSecurityExercise)
	case enums.ObjectTxWarrantExercise:
		return json.Marshal(t.WarrantExercise)

	//Issuance
	case enums.ObjectTxStockIssuance:
		return json.Marshal(t.StockIssuance)
	case enums.ObjectTxConvertibleIssuance:
		return json.Marshal(t.ConvertibleIssuance)
	case enums.ObjectTxPlanSecurityIssuance:
		return json.Marshal(t.PlanSecurityIssuance)
	case enums.ObjectTxEquityCompensationIssuance:
		return json.Marshal(t.EquityCompensationIssuance)
	case enums.ObjectTxWarrantIssuance:
		return json.Marshal(t.WarrantIssuance)

	//Reissuance
	case enums.ObjectTxStockReissuance:
		return json.Marshal(t.StockReissuance)

	//Release
	case enums.ObjectTxEquityCompensationRelease:
		return json.Marshal(t.EquityCompensationRelease)
	case enums.ObjectTxPlanSecurityRelease:
		return json.Marshal(t.PlanSecurityRelease)

	//Repurchase
	case enums.ObjectTxStockRepurchase:
		return json.Marshal(t.StockRepurchase)

	//Retraction
	case enums.ObjectTxConvertibleRetraction:
		return json.Marshal(t.ConvertibleRetraction)
	case enums.ObjectTxEquityCompensationRetraction:
		return json.Marshal(t.EquityCompensationRetraction)
	case enums.ObjectTxPlanSecurityRetraction:
		return json.Marshal(t.PlanSecurityRetraction)
	case enums.ObjectTxStockRetraction:
		return json.Marshal(t.StockRetraction)
	case enums.ObjectTxWarrantRetraction:
		return json.Marshal(t.WarrantRetraction)

	//ReturnToPool
	case enums.ObjectTxStockPlanReturnToPool:
		return json.Marshal(t.StockPlanReturnToPool)

	//Split
	case enums.ObjectTxStockClassSplit:
		return json.Marshal(t.StockClassSplit)

	//Transfer
	case enums.ObjectTxConvertibleTransfer:
		return json.Marshal(t.ConvertibleTransfer)
	case enums.ObjectTxEquityCompensationTransfer:
		return json.Marshal(t.EquityCompensationTransfer)
	case enums.ObjectTxPlanSecurityTransfer:
		return json.Marshal(t.PlanSecurityTransfer)
	case enums.ObjectTxStockTransfer:
		return json.Marshal(t.StockTransfer)
	case enums.ObjectTxWarrantTransfer:
		return json.Marshal(t.WarrantTransfer)

	//Vesting
	case enums.ObjectTxVestingAcceleration:
		return json.Marshal(t.VestingAcceleration)
	case enums.ObjectTxVestingEvent:
		return json.Marshal(t.VestingEvent)
	case enums.ObjectTxVestingStart:
		return json.Marshal(t.VestingStart)
	}

	return nil, fmt.Errorf(fmt.Sprintf("unhandled transaction type %v", t.ObjectType))
}
