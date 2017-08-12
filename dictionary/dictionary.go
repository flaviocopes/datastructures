// Package dictionary creates a ValueDictionary data structure for the Item type
package dictionary

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Key the key of the dictionary
type Key generic.Type

// Value the content of the dictionary
type Value generic.Type

// ValueDictionary the set of Items
type ValueDictionary struct {
	items map[Key]Value
	lock  sync.RWMutex
}

// Set adds a new item to the dictionary
func (d *ValueDictionary) Set(k Key, v Value) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.items == nil {
		d.items = make(map[Key]Value)
	}
	d.items[k] = v
}

// Delete removes a value from the dictionary, given its key
func (d *ValueDictionary) Delete(k Key) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, ok := d.items[k]
	if ok {
		delete(d.items, k)
	}
	return ok
}

// Has returns true if the key exists in the dictionary
func (d *ValueDictionary) Has(k Key) bool {
	d.lock.RLock()
	defer d.lock.RUnlock()
	_, ok := d.items[k]
	return ok
}

// Get returns the value associated with the key
func (d *ValueDictionary) Get(k Key) Value {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.items[k]
}

// Clear removes all the items from the dictionary
func (d *ValueDictionary) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.items = make(map[Key]Value)
}

// Size returns the amount of elements in the dictionary
func (d *ValueDictionary) Size() int {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return len(d.items)
}

// Keys returns a slice of all the keys present
func (d *ValueDictionary) Keys() []Key {
	d.lock.RLock()
	defer d.lock.RUnlock()
	keys := []Key{}
	for i := range d.items {
		keys = append(keys, i)
	}
	return keys
}

// Values returns a slice of all the values present
func (d *ValueDictionary) Values() []Value {
	d.lock.RLock()
	defer d.lock.RUnlock()
	values := []Value{}
	for i := range d.items {
		values = append(values, d.items[i])
	}
	return values
}
