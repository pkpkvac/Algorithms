package timebasedkeyvaluestore

import (
	"testing"
)

func TestTimeMap(t *testing.T) {

	t.Run("leetcode example 1", func(t *testing.T) {
		// TimeMap timeMap = new TimeMap();
		timeMap := Constructor()

		// timeMap.set("alice", "happy", 1);
		timeMap.Set("alice", "happy", 1)

		// timeMap.get("alice", 1); // return "happy"
		result1 := timeMap.Get("alice", 1)
		if result1 != "happy" {
			t.Errorf("Get(\"alice\", 1) = %v, want %v", result1, "happy")
		}

		// timeMap.get("alice", 2); // return "happy", there is no value stored for timestamp 2, thus we return the value at timestamp 1.
		result2 := timeMap.Get("alice", 2)
		if result2 != "happy" {
			t.Errorf("Get(\"alice\", 2) = %v, want %v", result2, "happy")
		}

		// timeMap.set("alice", "sad", 3);
		timeMap.Set("alice", "sad", 3)

		// timeMap.get("alice", 3); // return "sad"
		result3 := timeMap.Get("alice", 3)
		if result3 != "sad" {
			t.Errorf("Get(\"alice\", 3) = %v, want %v", result3, "sad")
		}
	})
}
