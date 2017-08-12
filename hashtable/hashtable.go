// Package hashtable creates a ValueHashtable data structure for the Item type
package hashtable

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Key the key of the dictionary
type Key generic.Type

// Value the content of the dictionary
type Value generic.Type

// ValueHashtable the set of Items
type ValueHashtable struct {
	items map[int]Value
	lock  sync.RWMutex
}

// the hash() private function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
func hash(k Key) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

// Put item with value v and key k into the hashtable
func (ht *ValueHashtable) Put(k Key, v Value) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := hash(k)
	if ht.items == nil {
		ht.items = make(map[int]Value)
	}
	ht.items[i] = v
}

// Remove item with key k from hashtable
func (ht *ValueHashtable) Remove(k Key) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := hash(k)
	delete(ht.items, i)
}

// Get item with key k from the hashtable
func (ht *ValueHashtable) Get(k Key) Value {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	i := hash(k)
	return ht.items[i]
}

// Size returns the number of the hashtable elements
func (ht *ValueHashtable) Size() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return len(ht.items)
}
