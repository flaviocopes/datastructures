package dictionary

import (
	"fmt"
	"testing"
)

func populateDictionary(count int, start int) *ValueDictionary {
	dict := ValueDictionary{}
	for i := start; i < (start + count); i++ {
		dict.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	return &dict
}

func TestSet(t *testing.T) {
	dict := populateDictionary(3, 0)
	if size := dict.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	dict.Set("key1", "value1") //should not add a new one, just change the existing one
	if size := dict.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	dict.Set("key4", "value4") //should add it
	if size := dict.Size(); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}
}

func TestDelete(t *testing.T) {
	dict := populateDictionary(3, 0)
	dict.Delete("key2")
	if size := dict.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}

func TestClear(t *testing.T) {
	dict := populateDictionary(3, 0)
	dict.Clear()
	if size := dict.Size(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}
}

func TestHas(t *testing.T) {
	dict := populateDictionary(3, 0)
	has := dict.Has("key2")
	if !has {
		t.Errorf("expected key2 to be there")
	}
	dict.Delete("key2")
	has = dict.Has("key2")
	if has {
		t.Errorf("expected key2 to be removed")
	}
	dict.Delete("key1")
	has = dict.Has("key1")
	if has {
		t.Errorf("expected key1 to be removed")
	}
}

func TestKeys(t *testing.T) {
	dict := populateDictionary(3, 0)
	items := dict.Keys()
	if len(items) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", len(items))
	}
	dict = populateDictionary(520, 0)
	items = dict.Keys()
	if len(items) != 520 {
		t.Errorf("wrong count, expected 520 and got %d", len(items))
	}
}

func TestValues(t *testing.T) {
	dict := populateDictionary(3, 0)
	items := dict.Values()
	if len(items) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", len(items))
	}
	dict = populateDictionary(520, 0)
	items = dict.Values()
	if len(items) != 520 {
		t.Errorf("wrong count, expected 520 and got %d", len(items))
	}
}

func TestSize(t *testing.T) {
	dict := populateDictionary(3, 0)
	items := dict.Values()
	if len(items) != dict.Size() {
		t.Errorf("wrong count, expected %d and got %d", dict.Size(), len(items))
	}
	dict = populateDictionary(0, 0)
	items = dict.Values()
	if len(items) != dict.Size() {
		t.Errorf("wrong count, expected %d and got %d", dict.Size(), len(items))
	}
	dict = populateDictionary(10000, 0)
	items = dict.Values()
	if len(items) != dict.Size() {
		t.Errorf("wrong count, expected %d and got %d", dict.Size(), len(items))
	}
}
