type Values struct{
	value string
	timestamp int
}
type TimeMap struct {
	store map[string][]Values
}

func Constructor() TimeMap {
	return TimeMap{
		store : make(map[string][]Values),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key],Values{
		value : value,
		timestamp : timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries,exists := this.store[key]
	if !exists || len(entries) == 0{
		return ""
	}
	left := 0
	right := len(entries) - 1
	res := ""

	for left <= right{
		mid := left + (right - left)/2
		if entries[mid].timestamp <= timestamp{
			left = mid + 1
			res = entries[mid].value
		}else{
			right = mid - 1
		}
	}
	return res
}
