package readdata

import (
	"math/rand"
	"time"
)

func Initial_solution(data Nodes) Nodes {
	rand.Seed(time.Now().UnixMilli())

	r := rand.Intn(len(data))
	solution := Append_solution(data, r)

	data = Visit_node(data, data[r].id)
	cur_node := data[r]

	for IsThere(data) {
		nxt := Next_item_finder(data, cur_node)
		data = Visit_node(data, nxt)
		for i := 0; i < len(data); i++ {
			if nxt == data[i].id {
				solution = append(solution, data[i])
				cur_node = data[i]
			}
		}
	}

	return solution
}
