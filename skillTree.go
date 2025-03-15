package main

import "github.com/hajimehoshi/ebiten/v2"

type SkillTree struct {
	nodes []*Node
}

func (st *SkillTree) Update(offsetX, offsetY, windowOffsetX, windowOffsetY int, zoom float64) {
	for i := range st.nodes {
		st.nodes[i].offsetX = offsetX
		st.nodes[i].offsetY = offsetY
		st.nodes[i].Update(offsetX, offsetY, windowOffsetX, windowOffsetY, zoom)
	}
}

func (st *SkillTree) Draw(screen *ebiten.Image) {
	for i := range st.nodes {
		st.nodes[i].Draw(screen)
	}
}

func (st *SkillTree) AddNode(node *Node) {
	st.nodes = append(st.nodes, node)
}
