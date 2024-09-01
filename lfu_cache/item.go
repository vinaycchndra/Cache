package lfu_cache

type Item struct {
	key        string
	index      int32
	time_stamp int32
	frequency  int32
	value      any
}

func newItem(key string, value any, freq, time_stamp int32) *Item {
	return &Item{key: key, value: value, frequency: freq, time_stamp: time_stamp}
}

func (i Item) lessThan(j *Item) bool {
	if i.frequency < j.frequency || (i.frequency == j.frequency && i.time_stamp < j.time_stamp) {
		return true
	}
	return false
}

func (i *Item) setValue(value any) {
	i.value = value
}

func (i *Item) getValue() any {
	return i.value
}

func (i *Item) getKey() string {
	return i.key
}

func (i *Item) getIndex() int32 {
	return i.index
}

func (i *Item) setIndex(ind int32) {
	i.index = ind
}

func (i *Item) setTime(ind int32) {
	i.time_stamp = ind
}

func (i *Item) increaseCount() {
	i.frequency++
}
