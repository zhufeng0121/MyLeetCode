package SwordToOffer

type RecentCounter struct {
	counter []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		counter: make([]int, 0),
	}

}

func (this *RecentCounter) Ping(t int) int {
	if len(this.counter) == 0 {
		this.counter = append(this.counter, t)
		return len(this.counter)
	}

	if t-this.counter[0] <= 3000 {
		this.counter = append(this.counter, t)
	} else {
		this.counter = append(this.counter, t)
		for t-this.counter[0] > 3000 {
			this.counter = this.counter[1:]
		}

	}
	return len(this.counter)

}
