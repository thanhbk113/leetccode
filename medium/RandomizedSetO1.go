package main

import "math/rand"

type RandomizedSet struct {
	values []int
	indexs map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values: make([]int, 0),
		indexs: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexs[val]; ok {
		return false
	}

	this.values = append(this.values, val)
	this.indexs[val] = len(this.values) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.indexs[val]; ok {
		lastValue := this.values[len(this.values)-1]
		this.values[idx] = lastValue
		this.indexs[lastValue] = idx
		this.values = this.values[:len(this.values)-1]
		delete(this.indexs, val)
		return true
	}

	return false

}

func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	a := Constructor()

	a.Insert(1)
	a.Insert(2)
	a.Insert(3)
	a.Remove(1)
}
