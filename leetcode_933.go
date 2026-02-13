package main

type RecentCounter struct {
	f []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		f: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.f = append(this.f, t)
	limit := t - 3000

	for len(this.f) > 0 && this.f[0] < limit {
		this.f = this.f[1:]
	}

	return len(this.f)
}
