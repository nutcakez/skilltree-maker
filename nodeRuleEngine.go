package main

type NodeRuleEngine struct {
	rules []Rule
}

type Rule struct {
	Name  string
	Check func(n *Node, tree *SkillTree) bool
}

func NoMoreRank1(n *Node, tree *SkillTree) bool {
	if n.active {
		return false
	}

	for _, node := range tree.Nodes {
		if node.active && node.tags["rank"] == "1" {
			return false
		}
	}
	return true
}

func NoMoreRank2(n *Node, tree *SkillTree) bool {
	if n.active {
		return false
	}

	for _, node := range tree.Nodes {
		if node.active && node.tags["rank"] == "2" {
			return false
		}
	}
	return true
}
