package retraction

type Retraction struct {
	*ConvertibleRetraction
	*EquityCompensationRetraction
	*PlanSecurityRetraction
	*StockRetraction
	*WarrantRetraction
}
