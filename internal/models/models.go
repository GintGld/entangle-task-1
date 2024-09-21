package models

import (
	"encoding/json"
	"math/big"
)

type NGLAmountResp struct {
	Supply     []NGLSupply   `json:"supply"`
	Pagination NGLPagination `json:"pagination"`
}

type NGLSupply struct {
	Denom  string      `json:"denom"`
	Amount *AmountType `json:"amount"`
}

type NGLPagination struct {
	NextKey interface{} `json:"next_key"`
	Total   string      `json:"total"`
}

type AmountType big.Int

// Custom Unmarshal to convert json string
// to big.Int
func (a *AmountType) UnmarshalJSON(data []byte) error {
	var tmp string

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var i big.Int
	i.SetString(tmp, 10)

	*a = AmountType(i)

	return nil
}

func (a *AmountType) BigInt() big.Int {
	return big.Int(*a)
}
