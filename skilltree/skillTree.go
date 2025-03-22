package skilltree

import "github.com/hajimehoshi/ebiten/v2"

type SkillTree struct {
	Nodes []*Node
}

func (st *SkillTree) Update(offsetX, offsetY, windowOffsetX, windowOffsetY int, zoom float64) {
	for i := range st.Nodes {
		st.Nodes[i].offsetX = offsetX
		st.Nodes[i].offsetY = offsetY
		st.Nodes[i].Update(offsetX, offsetY, windowOffsetX, windowOffsetY, zoom)
	}
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
	st.Nodes = append(st.Nodes, node)
}
