package models

const (
	MaximumCashRegisterOneThousandBahtBankNotes  = 10
	MaximumCashRegisterFiveHundredBahtBankNotes  = 20
	MaximumCashRegisterOneHundredBahtBankNotes   = 15
	MaximumCashRegisterFiftyBahtBankNotes        = 20
	MaximumCashRegisterTwentyBahtBankNotes       = 30
	MaximumCashRegisterTenBahtBankCoins          = 20
	MaximumCashRegisterFiveBahtBankCoins         = 20
	MaximumCashRegisterOneBahtBankCoins          = 20
	MaximumCashRegisterTwentyFiveSatangBankCoins = 50
)

var CashRegisterIndexMap = map[float64]BankNoteAndCoin{
	1000: {
		Index:       0,
		Value:       1000,
		MaxQuantity: MaximumCashRegisterOneThousandBahtBankNotes,
	},
	500: {
		Index:       1,
		Value:       500,
		MaxQuantity: MaximumCashRegisterFiveHundredBahtBankNotes,
	},
	100: {
		Index:       2,
		Value:       100,
		MaxQuantity: MaximumCashRegisterOneHundredBahtBankNotes,
	},
	50: {
		Index:       3,
		Value:       50,
		MaxQuantity: MaximumCashRegisterFiftyBahtBankNotes,
	},
	20: {
		Index:       4,
		Value:       20,
		MaxQuantity: MaximumCashRegisterTwentyBahtBankNotes,
	},
	10: {
		Index:       5,
		Value:       10,
		MaxQuantity: MaximumCashRegisterTenBahtBankCoins,
	},
	5: {
		Index:       6,
		Value:       5,
		MaxQuantity: MaximumCashRegisterFiveBahtBankCoins,
	},
	1: {
		Index:       7,
		Value:       1,
		MaxQuantity: MaximumCashRegisterOneBahtBankCoins,
	},
	0.25: {
		Index:       8,
		Value:       0.25,
		MaxQuantity: MaximumCashRegisterTwentyFiveSatangBankCoins,
	},
}

var CashRegisterBankNotes = []float64{1000, 500, 100, 50, 20, 10, 5, 1, 0.25}

type CashRegister struct {
	BankNoteAndCoins []BankNoteAndCoin `json:"cash_register"`
}

type PaidInfo struct {
	TotalPaid         float64 `json:"total_paid"`
	TotalProductPrice float64 `json:"total_product_price"`
}

type ChangeInfo struct {
	TotalChange           float64           `json:"total_change"`
	Changes               []BankNoteAndCoin `json:"changes"`
	CashRegisterAfterPaid []BankNoteAndCoin `json:"cash_register_after_paid"`
}

type BankNoteAndCoin struct {
	MaxQuantity int     `json:"max_quantity"`
	Quantity    int     `json:"quantity"`
	Value       float64 `json:"value"`
	Index       int     `json:"index"`
}
