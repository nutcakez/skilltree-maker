package main

import "github.com/nutcakez/skilltree-maker/skilltree"

type NodeRuleEngine struct {
	rules []Rule
}

type Rule struct {
	Name  string
	Check func(n *skilltree.Node, tree *skilltree.SkillTree) bool
}

func NoMoreRank1(n *skilltree.Node, tree *skilltree.SkillTree) bool {
	if n.Active {
		return false
	}

	for _, node := range tree.Nodes {
		if node.Active && node.Tags["rank"] == "1" {
			return false
		}
	}
	return true
}

func NoMoreRank2(n *skilltree.Node, tree *skilltree.SkillTree) bool {
	if n.Active {
		return false
	}

	for _, node := range tree.Nodes {
		if node.Active && node.Tags["rank"] == "2" {
			return false
		}
	}
	return true
}
