package yrbtree

import (
    "error"
)

type Color int

const (
    Unknown Color = 0
    Black Color = 1
    Red   Color = 2
)

type Value interface {
    compare(val value) int // -1/0/1 less/equal/bigger then param val
}

type Node struct {
    val *Value
    left *Node
    right *Node
    color Color

    parent *Node
}

func (n Node) compare(n2 Node) int {
    return n.val.compare(n2.val)
}

func (n *Node) uncle() *Node {
    if n.parent == nil {
        return nil
    }
    return n.parent.right
}

func (n *Node) pparent() *Node {
    if n.parent == nil {
        return nil
    }
    return n.parent.parent
}

func (n *Node) lrotate() error {
    if n.parent == nil {
        return errors.New("can not lrotate in root")
    }

    if n.right == nil {
        return errors.New("can not lrotate in node without right node")
    }

    r := n.right
    p := n.parent

    // n is ok
    n.parent = r
    n.right = r.left

    // r.left is ok
    if r.left != nil {
        r.left.parent = n
    }

    // r is ok
    r.parent = p
    r.left = n
    return nil
}

func (n *Node) rrotate() error {
    if n.parent == nil {
        return errors.New("can not lrotate in root")
    }

    if n.left == nil {
        return errors.New("can not lrotate in node without left node")
    }

    l := n.left
    p := n.parent

    // n is ok
    n.parent = l
    n.left = l.right

    // l.right is ok
    if l.right != nil {
        l.right.parent = n
    }

    // l is ok
    l.parent = p
    l.right = n
    return nil
}

type RBTree struct {
    root *Node
}

func InitRBTree(val Value) *RBTree {
    r := new(Node)
    r.left = r.right = nil
    r.color = Black
    r.val = &val
    return &RBTree{root: r}
}

func (r *RBTree) EnQueue(val Value) error {
    // find node position
    p := r.root
    inst := &Node{val: &val, color: Red}
    for {
        if val.compare(p.val) == -1 {
            if p.left == nil {
                p.left = inst
                inst.parent = p
                break
            } else {
                p = p.left
                continue
            }
        } else {
            if p.right == nil {
                p.right = inst
                inst.parent = p
                break
            } else {
                p = p.right
                continue
            }
        }
    }

    if inst.parent.parent == nil {
        return nil
    }

    if inst.uncle() == nil && inst.parent.left == inst{
        inst.parent.parent.rrotate()
        inst.parent.color = Black
        inst.parent.right.color = Red
        return nil
    }

    if inst.uncle() == nil && inst.parent.right == inst {
        inst.parent.lrotate()
    }

    if inst.parent.color == Black {
        return nil
    }

    if inst.uncle().color == Red {

    }
}

func (r *RBTree) DeQueue(val Value) error {

}
