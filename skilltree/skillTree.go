package skilltree

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type eventListener interface {
	ClickedNode(nodeIndex int)
}

type SkillTree struct {
	Nodes          []*Node
	ruleEngine     *NodeRuleEngine
	eventListeners []eventListener
}

func (st *SkillTree) Update(offsetX, offsetY, windowOffsetX, windowOffsetY int, zoom float64) (*int, string) {
	for i := range st.Nodes {
		st.Nodes[i].offsetX = offsetX
		st.Nodes[i].offsetY = offsetY
		clicked, hovered := st.Nodes[i].Update(offsetX, offsetY, windowOffsetX, windowOffsetY, zoom)
		if clicked {
			for x := range st.eventListeners {
				st.eventListeners[x].ClickedNode(i)
			}
			return &i, st.Nodes[i].HoverText
		}
		if hovered {
			return nil, st.Nodes[i].HoverText
		}
	}

	return nil, ""
}

func (st *SkillTree) Draw(screen *ebiten.Image) {
	for i := range st.Nodes {
		st.Nodes[i].DrawLines(screen)
	}
	for i := range st.Nodes {
		st.Nodes[i].Draw(screen)
	}
}

func (st *SkillTree) AddNode(node *Node) {
	node.RuleEngine = st.ruleEngine
	st.Nodes = append(st.Nodes, node)
}

func (st *SkillTree) AddRule(name string, rule func(n *Node, st *SkillTree) bool) {
	st.ruleEngine.addRule(name, rule)
}

func (st *SkillTree) RuleCheck(node *Node) bool {
	return st.ruleEngine.Check(node, st)
}

func (st *SkillTree) AddEventListener(listener eventListener) {
	st.eventListeners = append(st.eventListeners, listener)
}
