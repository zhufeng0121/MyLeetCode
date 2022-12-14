package SwordToOffer

type MovingAverage struct {
	window []int
	size   int
	sum    int
}

/** Initialize your data structure here. */
func ConstructorI(size int) MovingAverage {
	return MovingAverage{
		window: make([]int, 0, size),
		size:   size,
		sum:    0,
	}

}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.window) == this.size {
		this.sum -= this.window[0]
		this.window = this.window[1:]
	}
	this.window = append(this.window, val)
	this.sum += val

	return float64(this.sum) / float64(len(this.window))

}
