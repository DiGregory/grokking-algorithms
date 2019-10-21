package main

type node struct {
	name      string
	haveMango bool
}

type graph struct {
	nodes map[node][]node
}
type queue struct {
	nodes []node
}

func (q *queue) popleft() (n node) {

	n = q.nodes[0]
	q.nodes = q.nodes[1:]
	return
}

func createGraph() (graph) {

	return graph{nodes: map[node][]node{}}
}

func createNode(n string, m bool) (node) {
	return node{n, m}
}

func (g *graph) addNodeNeigbors(n node, neighbors ...node) () {

	g.nodes[n] = append(g.nodes[n], neighbors...)

}
func (g *graph)FindMangoSeller(startNode node) (node) {
	if startNode.haveMango{
		return startNode
	}
	searchQueue := queue{}
	searchQueue.nodes = append(searchQueue.nodes, startNode)
	searched := []node{}
LOOP:
	for len(searchQueue.nodes) > 0 {
		person := searchQueue.popleft()
		for _, v := range searched {
			if v.name == person.name {
				continue LOOP
			}
		}

		if person.haveMango == true {
			return person
		} else {
			searchQueue.nodes=append(searchQueue.nodes,g.nodes[person]...)
			searched=append(searched,person)

		}

	}
	return node{}
}

