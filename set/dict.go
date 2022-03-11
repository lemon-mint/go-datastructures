/*
Copyright 2014 Workiva, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package set is a simple unordered set implemented with a map.  This set
is threadsafe which decreases performance.

TODO: Actually write custom hashmap using the hash/fnv hasher.

TODO: Our Set implementation Could be further optimized by getting the uintptr
of the generic interface{} used and using that as the key; Golang maps handle
uintptr much better than the generic interace{} key.
*/

package set

import (
	"sync"
)

// Set is an implementation of ISet using the builtin map type. Set is threadsafe.
type Set[T comparable] struct {
	items     map[T]struct{}
	lock      sync.RWMutex
	flattened []T
}

// Add will add the provided items to the set.
func (set *Set[T]) Add(items ...T) {
	set.lock.Lock()
	defer set.lock.Unlock()

	set.flattened = nil
	for _, item := range items {
		set.items[item] = struct{}{}
	}
}

// Remove will remove the given items from the set.
func (set *Set[T]) Remove(items ...T) {
	set.lock.Lock()
	defer set.lock.Unlock()

	set.flattened = nil
	for _, item := range items {
		delete(set.items, item)
	}
}

// Exists returns a bool indicating if the given item exists in the set.
func (set *Set[T]) Exists(item T) bool {
	set.lock.RLock()

	_, ok := set.items[item]

	set.lock.RUnlock()

	return ok
}

// Flatten will return a list of the items in the set.
func (set *Set[T]) Flatten() []T {
	set.lock.Lock()
	defer set.lock.Unlock()

	if set.flattened != nil {
		return set.flattened
	}

	set.flattened = make([]T, 0, len(set.items))
	for item := range set.items {
		set.flattened = append(set.flattened, item)
	}
	return set.flattened
}

// Len returns the number of items in the set.
func (set *Set[T]) Len() int64 {
	set.lock.RLock()

	size := int64(len(set.items))

	set.lock.RUnlock()

	return size
}

// Clear will remove all items from the set.
func (set *Set[T]) Clear() {
	set.lock.Lock()

	set.items = map[T]struct{}{}

	set.lock.Unlock()
}

// All returns a bool indicating if all of the supplied items exist in the set.
func (set *Set[T]) All(items ...T) bool {
	set.lock.RLock()
	defer set.lock.RUnlock()

	for _, item := range items {
		if _, ok := set.items[item]; !ok {
			return false
		}
	}

	return true
}

// Dispose will add this set back into the pool.
func (set *Set[T]) Dispose() {
	set.lock.Lock()
	defer set.lock.Unlock()

	for k := range set.items {
		delete(set.items, k)
	}

	//this is so we don't hang onto any references
	for i := 0; i < len(set.flattened); i++ {
		var zero T
		set.flattened[i] = zero
	}

	set.flattened = set.flattened[:0]
}

// New is the constructor for sets. It will pull from a reuseable memory pool if it can.
// Takes a list of items to initialize the set with.
func New[T comparable](items ...T) *Set[T] {
	set := &Set[T]{
		items: make(map[T]struct{}, 10),
	}
	for _, item := range items {
		set.items[item] = struct{}{}
	}

	if len(items) > 0 {
		set.flattened = nil
	}

	return set
}
