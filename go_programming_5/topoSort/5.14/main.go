package main

//Node 有向图节点
type Node struct {
	name        string
	childNodes  []*Node //学习完当前课程后才可以学习的"进阶"课程
	parentNodes []*Node //学习当前课程需要首先学习的"基础"课程
}

//!+table
//exercise5.10 prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

type graph struct {
	name   string
	parent int
	child  []*graph
}

//!-table

//!+main
//exercise5.14 广度优先遍历拓扑结构
func main() {
	record := make(map[string]*graph)
	var queue []string
	for key, strings := range prereqs {
		if record[key] == nil {
			record[key] = &graph{name: key,
				parent: len(strings), child: nil}
		}

		for _, str := range strings {
			if record[str] == nil {
				record[str] = &graph{name: str, parent: len(prereqs[str])}
			}
			record[str].child = append(record[str].child, record[key])
		}
	}

	queue = searchadd(record, queue)
	index := 0
	for len(queue) > 0 {
		str := queue[index]
		println(str)
		for _, node := range record[str].child {
			n := node
			n.parent -= 1
			if n.parent == 0 {
				queue = append(queue, node.name)
			}
		}
		delete(record, str)
		queue = queue[1:]
		if len(queue) == 0 {
			queue = searchadd(record, queue)
		}
	}

}

func searchadd(record map[string]*graph, queue []string) []string {
	for _, g := range record {
		if g.parent == 0 {
			queue = append(queue, g.name)
			return queue
		}
	}
	return nil
}
