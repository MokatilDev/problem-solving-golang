package time_based_key_value_store

type Value struct {
	val       string
	timestamp int
}

type TimeMap struct {
	timeMap map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{
		timeMap: make(map[string][]Value),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if values, ok := this.timeMap[key]; ok {
		this.timeMap[key] = append(values, Value{val: value, timestamp: timestamp})
	} else {
		this.timeMap[key] = []Value{{val: value, timestamp: timestamp}}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if values, ok := this.timeMap[key]; ok {
		top := values[len(values)-1]
		if top.timestamp <= timestamp {
			return top.val
		}

		if values[0].timestamp > timestamp {
			return ""
		}

		l, r := 0, len(values)-1

		for l+1 < r {
			mid := (l + r) / 2
			if values[mid].timestamp > timestamp {
				r = mid
			} else {
				l = mid
			}
		}

		return values[l].val
	}

	return ""
}
