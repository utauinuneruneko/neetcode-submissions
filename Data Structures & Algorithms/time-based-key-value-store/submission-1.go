type State struct {
	nums []int
	vals map[int]string
}

type TimeMap struct {
	dict map[string]State
}

func Constructor() TimeMap {
	return TimeMap{
		dict: make(map[string]State),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.dict[key]; !ok {
		this.dict[key] = State{
			nums: []int{timestamp},
			vals: map[int]string{
				timestamp: value,
			},
		}
	} else {
		state := this.dict[key]
		state.nums = append(state.nums, timestamp)
		state.vals[timestamp] = value
		this.dict[key] = state
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.dict[key]; !ok {
		return ""
	}
	state := this.dict[key]
	if _, ok := state.vals[timestamp]; ok {
		return state.vals[timestamp]
	}

	nums := state.nums
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2+left
		if nums[mid] < timestamp {
			left = mid+1
		} else {
			right = mid-1
		}
	}
	if right < 0 {
		return ""
	}
	return state.vals[nums[right]]
}
