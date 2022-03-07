package readdata

import "math"

type nextItem struct {
	val   float64
	index int
}

// CalculateDistance calculates the distance between two nodes:
func CalculateDistance(node1 Node, node2 Node) float64 {
	return math.Sqrt(math.Pow(node1.x-node2.x, 2) + math.Pow(node1.y-node2.y, 2))
}

// Fitness calculates objective function (fitness function):
func Fitness(data Nodes) float64 {
	res := 0.0
	for i := 0; i < len(data)-1; i++ {
		res += CalculateDistance(data[i], data[i+1])
	}
	return res
}

// Path takes a Nodes struct and return the id numbers of node to tsp path:
func Path(data Nodes) []int {
	var p []int
	for _, node := range data {
		p = append(p, node.id)
	}
	return p
}

// AppendSolution takes Nodes named 'data' and a number named 'r', it returns a []Node appending r'th object:
func AppendSolution(data Nodes, r int) []Node {

	solution := make([]Node, 0)

	for i := 0; i < len(data); i++ {
		if r == i {
			var n Node
			n.id = data[i].id
			n.x = data[i].x
			n.y = data[i].y
			n.chosen = true
			solution = append(solution, n)
		}
	}

	return solution
}

// VisitNode takes an index and nodes as data and make the given index of node visited (chosen: true)
func VisitNode(data Nodes, index int) Nodes {
	solution := make([]Node, 0)
	for i := 0; i < len(data); i++ {
		var n Node
		n.id = data[i].id
		n.x = data[i].x
		n.y = data[i].y

		if index == data[i].id {
			n.chosen = true
		} else {
			n.chosen = data[i].chosen
		}
		solution = append(solution, n)
	}
	return solution
}

// IsThere returns true if there exists any node chosen:
func IsThere(data Nodes) (res bool) {
	for i := 0; i < len(data); i++ {
		if !data[i].chosen {
			res = true
		}
	}
	return
}

// MinSlice returns the index number of the min. value of a slice:
func MinSlice(v []nextItem) (ind int) {

	m := v[0].val
	ind = v[0].index

	for i := 1; i < len(v); i++ {
		if v[i].val < m {
			m = v[i].val
			ind = v[i].index
		}
	}

	return
}

// NextItemFinder returns the next node whose distance is minimum compared to others:
func NextItemFinder(rem Nodes, curr Node) int {

	nList := make([]nextItem, 0)

	for i := 0; i < len(rem); i++ {
		if !(curr.id == rem[i].id) && !rem[i].chosen {
			v := CalculateDistance(curr, rem[i])

			var it nextItem
			it.val = v
			it.index = rem[i].id

			nList = append(nList, it)
		}

	}

	return MinSlice(nList)
}

// Changelist takes a list and returns that list:
func Changelist(cur Nodes) Nodes {
	next := make(Nodes, 0)
	for i := 0; i < len(cur); i++ {
		var n Node
		n.id = cur[i].id
		n.x = cur[i].x
		n.y = cur[i].y
		n.chosen = cur[i].chosen

		next = append(next, n)
	}
	return next
}
