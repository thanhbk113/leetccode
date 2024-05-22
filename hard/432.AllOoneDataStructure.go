package main

import (
	"fmt"
)

type DLinkList[T comparable] struct {
	next  *DLinkList[T]
	prev  *DLinkList[T]
	count int
	val   T
}

type AllOne struct {
	head   *DLinkList[string] // save min value
	tail   *DLinkList[string] // save max value
	counts map[string]*DLinkList[string]
}

const NoneFound = ""

func Constructor() AllOne {
	o := AllOne{
		head:   &DLinkList[string]{val: NoneFound, count: -1},
		tail:   &DLinkList[string]{val: NoneFound, count: 9e2},
		counts: make(map[string]*DLinkList[string]),
	}

	o.head.next = o.tail
	o.tail.prev = o.head
	return o
}

func (this *AllOne) Inc(key string) {
	it := this.counts[key]
	if it == nil {
		it = &DLinkList[string]{this.head.next, this.head, 0, key}
		it.prev.next = it
		it.next.prev = it
		this.counts[key] = it
	}
	it.count++
	for it.count > it.next.count {
		right(it)
	}
}

func right[T comparable](it *DLinkList[T]) {
	//l-it-r-r2
	l, r, r2 := it.prev, it.next, it.next.next
	it.prev, it.next = r, r.next
	r2.prev = it
	r.prev, r.next = l, it
	l.next = r
	// l-r-it-r2
}

func (this *AllOne) Dec(key string) {
	it := this.counts[key]
	it.count--
	for it.count < it.prev.count {
		right(it.prev)
	}

	if it.count == 0 {
		it.prev.next = it.next
		it.next.prev = it.prev
		it.prev = nil
		it.next = nil
		delete(this.counts, key)
	}
}

func (this *AllOne) GetMaxKey() string {
	return this.tail.prev.val
}

func (this *AllOne) GetMinKey() string {
	return this.head.next.val
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

func main() {
	a := Constructor()
	a.Inc("a")
	a.Dec("a")
	a.Inc("b")
	a.Inc("b")
	a.Inc("b")
	a.Inc("b")
	a.Dec("b")
	a.Dec("b")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())
}
