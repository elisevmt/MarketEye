package graphZero

import (
	"errors"
	"fmt"
	"sort"
)

type Node struct { // Валютный узел, привязанный к бирже. Например, BTC Binance - это валютный узел, который привязан к бирже, имеет своё название и список направлений обмена [] Branch
	Name       string    `json:"name"`       // Name - название валюты (BTC, ETH, USDT, ...)
	Market     *Market   `json:"market"`     // Market - ссылка на биржу, к которой привязан узел
	Directions []*Branch `json:"directions"` // Directions - cписок направлений обмена, например, для BTC: BTCUSDT, BTCETH, BTCBUSD и т.д..
}

func NewNode(
	market *Market,
	name string) *Node {
	return &Node{
		Market:     market,
		Name:       name,
		Directions: nil,
	}
}

func (node *Node) Wire(
	partner *Node,
	weight float64,
	weightSize int64,
	count int64, symbol string,
	leftPrecision int64,
	rightPrecision int64,
	tradeFee float64,
	asks []*Lot,
	bids []*Lot,
) *Branch {
	branch := NewBranch(node.Market, node, partner, weight, weightSize, count, symbol, leftPrecision, rightPrecision, tradeFee, asks, bids)
	node.Directions = append(node.Directions, branch)
	return branch
}

func (node *Node) GetNode(
	market *Market,
	nodeName string,
) (*Node, error) {
	for _, el := range node.Directions {
		if el.Market == market && el.Right.Name == nodeName {
			return el.Right, nil
		}
	}
	return nil, errors.New("NodeNotFound")
}

func (node *Node) GetBranch(
	market *Market,
	nodeName string,
) (*Branch, error) {
	for _, el := range node.Directions {
		if el.Right.Name == nodeName && el.Right.Market == market {
			return el, nil
		}
	}
	return nil, errors.New("BranchNotFound")
}

func (node *Node) ToString() string {
	str := "Root: \n"
	for i, _ := range node.Directions {
		if i == len(node.Directions)/2 {
			str += fmt.Sprintf("%s---|", node.Name)
			continue
		}
		str += fmt.Sprintf("       |----%s: %f\n", node.Directions[i].Right.Name, node.Directions[i].LRWeight)
	}
	return str
}

func (node *Node) BranchReverseSort() {
	sort.Slice(node.Directions, func(i, j int) bool {
		return node.Directions[i].ExCount > node.Directions[j].ExCount
	})
}

func (node *Node) BranchRawSort() {
	sort.Slice(node.Directions, func(i, j int) bool {
		return node.Directions[i].ExCount < node.Directions[j].ExCount
	})
}
