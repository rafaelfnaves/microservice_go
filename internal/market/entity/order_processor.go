package entity

type OrderProcessor struct {
	Transaction *Transaction
}

func NewOrderProcessor(transaction *Transaction) *OrderProcessor {
	return &OrderProcessor{
		Transaction: transaction,
	}
}

func (op *OrderProcessor) Process() {
	shares := op.CalculateShares()
	op.UpdatePositions(shares)
	op.UpdateOrders(shares)
	op.Transaction.Total = float64(shares) * op.Transaction.Price
}

func (op *OrderProcessor) CalculateShares() int {
	avaiableShares := op.Transaction.Shares
	if op.Transaction.BuyingOrder.PendingShares < avaiableShares {
		avaiableShares = op.Transaction.BuyingOrder.PendingShares
	}

	if op.Transaction.SellingOrder.PendingShares < avaiableShares {
		avaiableShares = op.Transaction.SellingOrder.PendingShares
	}

	return avaiableShares
}

func (op *OrderProcessor) UpdatePositions(shares int) {
	op.Transaction.SellingOrder.Investor.UpdateAssetPosition(op.Transaction.BuyingOrder.Asset.ID, -shares)
	op.Transaction.BuyingOrder.Investor.UpdateAssetPosition(op.Transaction.BuyingOrder.Asset.ID, shares)
}

func (op *OrderProcessor) UpdateOrders(shares int) {
	op.Transaction.BuyingOrder.ApplyTrade(shares)
	op.Transaction.SellingOrder.ApplyTrade(shares)
}
