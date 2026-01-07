package p286

func wallsAndGates(rooms [][]int) {
	rows, cols := len(rooms), len(rooms[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	q := newQueue()

	var addRoom func(r, c int)
	addRoom = func(r, c int) {
		if r < 0 || r == rows || c < 0 || c == cols || visited[r][c] || rooms[r][c] == -1 {
			return
		}
		visited[r][c] = true
		q.Enqueue([]int{r, c})
	}

	for r := range rows {
		for c := range cols {
			if rooms[r][c] == 0 {
				q.Enqueue([]int{r, c})
				visited[r][c] = true
			}
		}
	}

	dist := 0
	for q.Len() > 0 {
		for range q.Len() {
			item := q.Dequeue()
			r, c := item[0], item[1]
			rooms[r][c] = dist
			addRoom(r+1, c)
			addRoom(r-1, c)
			addRoom(r, c+1)
			addRoom(r, c-1)
		}
		dist += 1
	}
}

type queue struct {
	items [][]int // list of x,y coords
}

func newQueue() *queue {
	return &queue{
		items: make([][]int, 0),
	}
}

func (r *queue) Enqueue(coords []int) {
	r.items = append(r.items, coords)
}

func (r *queue) Dequeue() []int {
	item := r.items[0]
	r.items = r.items[1:]
	return item
}

func (r *queue) Len() int {
	return len(r.items)
}
