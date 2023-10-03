package models

import "MarketEye/pkg/graphZero"

type OrderPool struct {
	Pool       []*MarginOrder
	Executed   bool
	ExitAmount *graphZero.Amount
}

type MarginOrder struct {
	Amount   *graphZero.Amount
	Branch   *graphZero.Branch
	Executed bool
}

type CalculatedArbitrateResponse struct {
	Way      []*graphZero.Branch
	InAmount *graphZero.Amount
	Margin   float64
}
