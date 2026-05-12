package main

import "math/rand"

type RandomizedSet struct {
	sets     []int
	indexMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		sets:     []int{},
		indexMap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, v := this.indexMap[val]; v {
		return false
	}

	this.indexMap[val] = len(this.sets)
	this.sets = append(this.sets, val)
	return true

}

func (this *RandomizedSet) Remove(val int) bool {
	idx, v := this.indexMap[val]

	if !v {
		return false
	}

	// 마지막 요소를 가져온다
	lastIdx := len(this.sets) - 1
	lastVal := this.sets[lastIdx]

	// 삭제할 위치에 마지막 요소를 가져온다
	this.sets[idx] = lastVal
	this.indexMap[lastVal] = idx

	this.sets = this.sets[:lastIdx]
	delete(this.indexMap, val)

	return true

}

func (this *RandomizedSet) GetRandom() int {
	return this.sets[rand.Intn(len(this.sets))]
}
