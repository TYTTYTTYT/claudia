package tools

import (
	"sync"
)

type node[K comparable, T any] struct {
	data T
	key K
	next *node[K, T]
	pre *node[K, T]
}

type LRUCache[K comparable, V any] struct {
	lock sync.Mutex
	capacity, count int
	k2n map[K]*node[K, V]
	head *node[K, V]
	tail *node[K, V]
}

func NewLRU[K comparable, V any](capacity int) *LRUCache[K, V] {
	cache := LRUCache[K, V]{
		lock: sync.Mutex{},
		capacity: capacity,
		count: 0,
		k2n: make(map[K]*node[K, V]),
		head: nil,
		tail: nil,
	}
	return &cache
}

func (cache *LRUCache[K, V]) Get(key K) (V, bool) {
	var result V
	if n, ok := cache.k2n[key]; ok {
		cache.lock.Lock()
		defer cache.lock.Unlock()

		result = n.data
		nxt, pre := n.next, n.pre

		if n == cache.tail && pre != nil {
			cache.tail = pre
		}


		if pre != nil {
			pre.next = nxt
		}
		if nxt != nil {
			nxt.pre = pre
		}
		if cache.head != nil {
			cache.head.pre = n
		}
		
		n.pre = nil
		n.next = cache.head
		cache.head = n


		return n.data, true
	} else {
		return result, false
	}
}

func (cache *LRUCache[K, V]) Put(key K, value V) {
	n := node[K, V]{value, key, nil, nil}

	cache.lock.Lock()
	defer cache.lock.Unlock()

	if cache.head != nil {
		cache.head.pre = &n
	}
	
	n.next = cache.head
	cache.head = &n
	cache.k2n[key] = &n
	if cache.tail == nil {
		cache.tail = &n
	}

	cache.count += 1
	if cache.count > cache.capacity {
		key = cache.tail.key
		delete(cache.k2n, key)
		cache.tail = cache.tail.pre
		if cache.tail != nil {
			cache.tail.next = nil
		}
		cache.count -= 1
	}
}
