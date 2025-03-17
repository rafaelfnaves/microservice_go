package entity

type OrderProcessor struct {
	Transaction *Transaction
}

func NewOrderProcessor(transaction *Transaction) *OrderProcessor {
	return &OrderProcessor{
		Transaction: transaction,
	}
}

func (op *OrderProcessor) Process() {}
