package readdata

import "math"

type next_item struct {
	val   float64
	index int
}

// it calculates the distance between two nodes:
func Calculate_distance(node1 Node, node2 Node) float64 {
	return math.Sqrt(math.Pow(node1.x_coord-node2.x_coord, 2) + math.Pow(node1.y_coord-node2.y_coord, 2))
}

// it calculates objective function (fitness function):
func Fitness(data Nodes) float64 {
	res := 0.0
	for i := 0; i < len(data)-1; i++ {
		res += Calculate_distance(data[i], data[i+1])
	}
	return res
}

// it takes a Nodes struct and return the id numbers of node to tsp path:
func Path(data Nodes) []int {
	p := []int{}
	for _, node := range data {
		p = append(p, node.id)
	}
	return p
}

// it takes Nodes named 'data' and a number named 'r', it returns a []Node appending r'th object:

func Append_solution(data Nodes, r int) []Node {

	solution := make([]Node, 0)

	for i := 0; i < len(data); i++ {
		if r == i {
			var n Node
			n.id = data[i].id
			n.x_coord = data[i].x_coord
			n.y_coord = data[i].y_coord
			n.choosen = true
			solution = append(solution, n)
		}
	}

	return solution
}

// it takes an index and nodes as data and make the given index of node visited (choosen: true)
func Visit_node(data Nodes, index int) Nodes {
	solution := make([]Node, 0)
	for i := 0; i < len(data); i++ {
		var n Node
		n.id = data[i].id
		n.x_coord = data[i].x_coord
		n.y_coord = data[i].y_coord

		if index == data[i].id {
			n.choosen = true
		} else {
			n.choosen = data[i].choosen
		}
		solution = append(solution, n)
	}
	return solution
}

// it returns true if there exists any node chosen:
func IsThere(data Nodes) bool {
	res := false
	for i := 0; i < len(data); i++ {
		if !data[i].choosen {
			res = true
		}
	}
	return res
}

// it returns the index number of the min. value of a slice:
func MinSlice(v []next_item) int {

	m := v[0].val
	ind := v[0].index

	for i := 1; i < len(v); i++ {
		if v[i].val < m {
			m = v[i].val
			ind = v[i].index
		}
	}

	return ind
}

// it returns the next node whose distance is minimum compared to others:
func Next_item_finder(rem Nodes, curr Node) int {

	n_list := make([]next_item, 0)

	for i := 0; i < len(rem); i++ {
		if !(curr.id == rem[i].id) && !rem[i].choosen {
			v := Calculate_distance(curr, rem[i])

			var it next_item
			it.val = v
			it.index = rem[i].id

			n_list = append(n_list, it)
		}

	}

	return MinSlice(n_list)
}

// it takes a list and returns that list:
func Changelist(cur Nodes) Nodes {
	next := make(Nodes, 0)
	for i := 0; i < len(cur); i++ {
		var n Node
		n.id = cur[i].id
		n.x_coord = cur[i].x_coord
		n.y_coord = cur[i].y_coord
		n.choosen = cur[i].choosen

		next = append(next, n)
	}
	return next
}
