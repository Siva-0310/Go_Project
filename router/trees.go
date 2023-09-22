package router

import "fmt"

type Node struct {
	PathLetter map[rune]*Node
	route      *Route
	CurrPath   []rune
}

func Insert(path []rune, root *Node, route *Route) {
	len_1 := len(root.CurrPath)
	len_2 := len(path)
	fmt.Println(root.route)
	if len_1 == 0 {
		root.route = route
		root.CurrPath = path
		return
	}
	i := commonPrefix(root.CurrPath, path, len_1, len_2)
	temp1 := &Node{
		CurrPath:   make([]rune, 0),
		PathLetter: map[rune]*Node{},
	}
	if i == 0 {
		root.PathLetter[path[0]] = temp1
		Insert(path, temp1, route)
		return
	}
	if len_2-i == 0 {
		temp_path := root.CurrPath[i:]
		temp_route := root.route
		root.CurrPath = root.CurrPath[:i]
		root.PathLetter[temp_path[0]] = temp1
		root.route = route
		Insert(temp_path, temp1, temp_route)
		return
	}

	path1 := path[i:]
	val := root.PathLetter[path1[0]]
	if val == nil {
		root.PathLetter[path1[0]] = temp1
		Insert(path1, temp1, route)
	} else {
		Insert(path1, val, route)
	}
	if len_1-i != 0 {
		temp2 := &Node{
			CurrPath:   make([]rune, 0),
			PathLetter: map[rune]*Node{},
		}
		path2 := root.CurrPath[i:]
		root.CurrPath = root.CurrPath[:i]
		temp_route := root.route
		root.route = nil
		root.PathLetter[path2[0]] = temp2
		Insert(path2, temp2, temp_route)
	}
}

func Search(path []rune, root *Node) *Route {
	len_1 := len(root.CurrPath)
	len_2 := len(path)
	i := commonPrefix(root.CurrPath, path, len_1, len_2)
	fmt.Println(root, path)
	if len_2-i == 0 {
		return root.route
	}
	if i == 0 {
		return nil
	}
	path1 := path[i:]
	val := root.PathLetter[path1[0]]
	if val == nil {
		return nil
	}
	return Search(path1, val)
}
func commonPrefix(path1 []rune, path2 []rune, len_1 int, len_2 int) int {
	i := 0
	for i < len_1 && i < len_2 && path1[i] == path2[i] {
		i += 1
	}
	return i
}
