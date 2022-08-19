package disjoint_set

type DSUNode struct {
	val    int
	parent *DSUNode
	rank   int
}

func MakeSet(x int) *DSUNode {
	d := &DSUNode{
		val:  x,
		rank: 0,
	}
	d.parent = d
	return d
}

func Union(x, y *DSUNode) {
	if x.rank > y.rank {
		y.parent = x
	} else {
		x.parent = y
		if x.rank == y.rank {
			y.rank++
		}
	}
}

func FindSet(x *DSUNode) *DSUNode {
	if x.parent != x {
		x.parent = FindSet(x.parent)
	}
	return x.parent
}
