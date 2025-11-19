package timebasedkeyvaluestore

type Entry struct {
	value     string
	timestamp int
}

type Key string

type TimeMap struct {
	m map[Key][]Entry
}

func Constructor() TimeMap {

	tm := TimeMap{m: make(map[Key][]Entry)}

	return tm
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {

	e := Entry{
		value:     value,
		timestamp: timestamp,
	}

	tm.m[Key(key)] = append(tm.m[Key(key)], e)
}

func (tm *TimeMap) Get(key string, timestamp int) string {

	// String get(String key, int timestamp) Returns the most
	// recent value of key if set was previously called
	// on it and the most recent timestamp for that key
	// prev_timestamp is less than or equal to the given
	// timestamp (prev_timestamp <= timestamp). If there are
	// no values, it returns "".
	// Note: For all calls to set, the timestamps are in strictly
	// increasing order.

	entries := tm.m[Key(key)]

	if entries == nil {
		return ""
	}

	// perform a binary search on the entries, specifically
	// seeking for ts <= timestamp
	l := 0
	r := len(entries) - 1
	for l <= r {
		mid := l + (r-l)/2
		// target = timestamp
		if entries[mid].timestamp == timestamp {
			return entries[mid].value
		}
		if entries[mid].timestamp > timestamp {
			// this means our solution is to the left
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if r >= 0 && entries[r].timestamp <= timestamp {
		return entries[r].value
	} else {
		return ""
	}
}
