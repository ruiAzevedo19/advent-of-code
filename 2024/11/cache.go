package solution

type Cache struct {
	values map[int]map[int]int
}

func NewCache() (cache *Cache) {
	return &Cache{
		values: map[int]map[int]int{},
	}
}

func (c *Cache) isCached(stone int, blink int) (numberOfStones int, ok bool) {
	blinks, ok := c.values[stone]
	if !ok {
		return 0, false
	}

	numberOfStones, ok = blinks[blink]
	if !ok {
		return 0, false
	}

	return numberOfStones, true
}

func (c *Cache) Add(stone int, blink int, numberOfStones int) {
	_, ok := c.values[stone]
	if !ok {
		c.values[stone] = map[int]int{}
	}

	_, ok = c.values[stone][blink]
	if !ok {
		c.values[stone][blink] = numberOfStones
	}
}
