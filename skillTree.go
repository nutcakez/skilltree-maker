package main

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
