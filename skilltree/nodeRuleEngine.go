package skilltree

type NodeRuleEngine struct {
	rules []Rule
}

type Rule struct {
	Name  string
	Check func(n *Node, tree *SkillTree) bool
}

// Example for filtering for rank1
func NoMoreRank1(n *Node, tree *SkillTree) bool {
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

func NoMoreRank2(n *Node, tree *SkillTree) bool {
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

func (nre *NodeRuleEngine) addRule(name string, rule func(n *Node, st *SkillTree) bool) {
	nre.rules = append(nre.rules, Rule{
		Name:  name,
		Check: rule,
	})
}

func (nre *NodeRuleEngine) Check(n *Node, st *SkillTree) bool {
	for _, rule := range nre.rules {
		if !rule.Check(n, st) {
			return false
		}
	}

	return true
}
