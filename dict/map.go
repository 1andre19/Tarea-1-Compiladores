package dict

const MAP_SIZE = 100

type MapItem struct {
	key   string
	value int
}

type HashMap struct {
	entries [][]MapItem
}

func NewDict() *HashMap {
	return &HashMap{
		entries: make([][]MapItem, MAP_SIZE),
	}
}

// djb2 hash function
// http://www.cse.yorku.ca/~oz/hash.html
func hash(key string) uint32 {
	var hash uint32 = 5381

	for _, v := range key {
		hash = hash*33 + uint32(v)
	}

	return hash % MAP_SIZE
}

func (h *HashMap) Set(key string, value int) {
	index := hash(key)

	// chaining to avoid collisions
	for i := range h.entries[index] {
		if h.entries[index][i].key == key {
			h.entries[index][i].value = value
			return
		}
	}

	h.entries[index] = append(h.entries[index], MapItem{key, value})
}

func (h *HashMap) Delete(key string) {
	index := hash(key)

	// chaining to avoid collisions
	for i := range h.entries[index] {
		if h.entries[index][i].key == key {
			h.entries[index] = append(h.entries[index][:i], h.entries[index][i+1:]...)
			return
		}
	}
}

func (h *HashMap) Get(key string) (int, bool) {
	index := hash(key)

	for i := range h.entries[index] {
		if h.entries[index][i].key == key {
			return h.entries[index][i].value, true
		}
	}
	return -1, false
}
