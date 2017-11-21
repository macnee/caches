package expirations

//structs
type ExpirationHeap struct {
	tree []Expiree
}

//interface implementations

func (e *ExpirationHeap) Push(x Expiree) {
	e.tree = append(e.tree, x)
	e.heapUp(e.getTail())
}

func (e *ExpirationHeap) Pop() []uint {
	retval := e.tree[0].getKeys()
	e.swap(0, e.getTail())
	e.delete(e.getTail())
	if !e.isEmpty() {
		e.heapDown(0)
	}
	return retval
}

//struct specific implementations

func (e *ExpirationHeap) heapUp (current uint) {
	if len(e.tree) == 0 {
		panic("we are trying to heap up on an empty tree")
	}

	if current == 0 {
		//we are now the root node, nowhere to go, we're done
		return
	}

	parent := e.getParentIndex(current)

	if e.tree[current].ComesBefore(e.tree[parent]) {
		e.swap(current, parent)
		e.heapUp(parent)
	} else if e.tree[current].IsEqual(e.tree[parent]) {
		//we need to merge the two nodes in question, take the tail and put it at current,
		//and then heap down
		e.merge(current, parent)
		tail := e.getTail()
		e.swap(current, tail)
		e.delete(tail)
		if !e.isEmpty() {
			e.heapDown(current)
		}
	}

	//else, we are already sorted!
}

func (e *ExpirationHeap) heapDown (current uint) {
	if (len(e.tree) == 0) {
		panic("we are trying to heap down on an empty tree")
	}

	if current == e.getTail() {
		//we are the last node, we're done!
		return
	}

	child := e.getMinChildIndex(current)
	if child < current {
		panic("somehow got a child closer to root than current")
	} else if current == child {
		//we didn't get a child sooner than us, we're done!
		return
	} else if e.tree[child].ComesBefore(e.tree[current]){
		e.swap(child, current)
		e.heapDown(child)
	}
}

func (e *ExpirationHeap) getParentIndex (idx uint) (uint) {
	if idx == 0 {
		panic ("We are looking for the parent of a root node, this should never happen.")
	} else if idx % 2 == 0 {
		return (idx / 2) - 1
	} else {
		return idx/ 2
	}
}

func (e *ExpirationHeap) getMinChildIndex (current uint) (uint) {
	left := (current * 2) + 1
	if left > e.getTail() {
		//if left is out of bounds, no children
		return current
	} else if left  == e.getTail() {
		//our left node is the last node in the array.
		if e.tree[left].ComesBefore(e.tree[current]) {
			return left
		}
		return current
	} else {
		//our right node is either at the end, or at a lower index. we must check both l&r
		right := left + 1
		if e.tree[left].ComesBefore(e.tree[current]) && e.tree[left].ComesBefore(e.tree[right]) {
			//left is bigger than both current and right
			return left
		} else if e.tree[right].ComesBefore(e.tree[current]) {
			return right
		} else {
			return current
		}
	}
}

func (e *ExpirationHeap) isEmpty() bool {
	return len(e.tree) == 0
}

func (e *ExpirationHeap) getTail() uint {
	return uint(len(e.tree) - 1)
}

func (e *ExpirationHeap) swap(a uint, b uint) {
	tmp := e.tree[a]
	e.tree[a] = e.tree[b]
	e.tree[b] = tmp
}

func (e *ExpirationHeap) merge(from uint, to uint) {
	e.tree[to].Append(e.tree[from])
}

func (e *ExpirationHeap) delete(i uint) {
	if e.getTail() == i {
		e.tree = e.tree[:i]
	} else if i == 0 {
		e.tree = e.tree[1:]
	} else {
		e.tree = append(e.tree[:i], e.tree[i+1:]...)
	}
}

//constructors
func NewExpirationHeap() ExpirationHeap {
	return ExpirationHeap{}
}