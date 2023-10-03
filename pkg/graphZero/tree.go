package graphZero

import (
	"errors"
)

type Tree struct { // Tree - это дерево всех валют, доступных арбитражному алгоритму. У основания может быть любая валюта, которая для нас является ключевой. В случае нашего арбитража - USDT
	Root     *Node     // Root - основание дерева (Tether)
	Nodes    []*Node   // Nodes - список всех валютных узлов, доступных QCA
	Branches []*Branch // Branches - список всех валютных ветвей
}

func NewTree(
	root *Node,
) *Tree {
	return &Tree{
		Root:  root,
		Nodes: []*Node{root},
	}
}

func (tree *Tree) Append(
	node *Node,
) {
	_, err := tree.Find(node.Name, node.Market)
	if err != nil {
		tree.Nodes = append(tree.Nodes, node)
	}
}

func (tree *Tree) GetTreeKernel(
	market []*Market,
	node *string,
	left *string,
	right *string,
	symbol *string,
) (answer KernelMarketPrice) {
	targetBranches := BranchFilterByLeft(left, tree.Branches)
	targetBranches = BranchFilterByRight(right, targetBranches)
	targetBranches = BranchFilterBySymbol(symbol, targetBranches)
	targetBranches = BranchFilterByNode(node, targetBranches)
	answer = BranchFilterByMarket(market, targetBranches)
	return
}

func (tree *Tree) Wire(
	nodeA,
	nodeB *Node,
	weight float64,
	weightSize int64,
	exCount int64,
	exSymbol string,
	leftPrecision int64,
	rightPrecision int64,
	tradeFee float64,
	asks []*Lot,
	bids []*Lot,
) *Branch {
	treeMemberA, err := tree.Find(nodeA.Name, nodeA.Market)
	if err != nil {
		treeMemberA = nodeA
	}
	treeMemberB, err := tree.Find(nodeB.Name, nodeB.Market)
	if err != nil {
		treeMemberB = nodeB
	}
	branch := treeMemberA.Wire(treeMemberB, weight, weightSize, exCount, exSymbol, leftPrecision, rightPrecision, tradeFee, asks, bids)
	tree.Branches = append(tree.Branches, branch)
	return branch
}

func (tree *Tree) Find(
	nodeName string,
	market *Market,
) (*Node, error) {
	for _, x := range tree.Nodes {
		if x.Name == nodeName && x.Market == market {
			return x, nil
		}
	}
	return nil, errors.New("NodeNotFound")
}

func (tree *Tree) GetBranch(
	Symbol string,
	market *Market,
) (*Branch, error) {
	for _, x := range tree.Branches {
		if x.ExSymbol == Symbol && x.Market == market && x.Left.Name+x.Right.Name == x.ExSymbol {
			return x, nil
		}
	}
	return nil, errors.New("BranchNotFound")
}

func (tree *Tree) WayExec(
	amount *Amount,
	way []*Branch,
) float64 {
	copyAmount := NewAmount(amount.Market, amount.Value, amount.Node)
	startValue := copyAmount.Value
	for _, x := range way {
		copyAmount.Move(x)
	}
	endValue := copyAmount.Value
	return endValue - startValue
}
