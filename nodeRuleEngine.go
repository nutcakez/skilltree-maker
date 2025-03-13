package main

type NodeRuleEngine struct {
	rules []Rule
}

type Rule struct {
	Name  string
	Check func(n *Node, tree *SkillTree) bool
}
