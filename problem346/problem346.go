package problem346

type MovingAverage struct {
	data []int
	size int
}

func Constructor(size int) MovingAverage {

	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {

	if len(this.data) == 0 {
		this.data = append(this.data, val)
		return float64(val)
	}

	sum := this.data[len(this.data)-1] + val
	this.data = append(this.data, sum)
	if len(this.data) <= this.size {
		return float64(sum) / float64(len(this.data))
	}

	return float64(sum-this.data[len(this.data)-this.size-1]) / float64(this.size)

}
