package adjustment

type Adjustment struct {
	*StockClassAuthorizedSharesAdjustment
	*StockClassConversionRatioAdjustment
	*StockPlanPoolAdjustment
}
