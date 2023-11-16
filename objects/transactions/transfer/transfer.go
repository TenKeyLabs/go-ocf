package transfer

type Transfer struct {
	*ConvertibleTransfer
	*EquityCompensationTransfer
	*PlanSecurityTransfer
	*StockTransfer
	*WarrantTransfer
}
