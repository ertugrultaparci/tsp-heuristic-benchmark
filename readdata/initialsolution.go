package readdata

import (
	"math/rand"
	"time"
)

func InitialSolution(data Nodes) Nodes {
	rand.Seed(time.Now().UnixMilli())

	r := rand.Intn(len(data))
	solution := AppendSolution(data, r)

	data = VisitNode(data, data[r].id)
	curNode := data[r]

	for IsThere(data) {
		nxt := NextItemFinder(data, curNode)
		data = VisitNode(data, nxt)
		for i := 0; i < len(data); i++ {
			if nxt == data[i].id {
				solution = append(solution, data[i])
				curNode = data[i]
			}
		}
	}

	return solution
}
