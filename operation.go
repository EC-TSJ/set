/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package set

import (
	"bytes"
	"ec-tsj/core"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type (
	// Dato
	T = core.T
	//
	uSet map[T]struct{}

	// Element tipo de dato
	Element struct {
		D T
	}
	// Set dato de []Element
	Set []Element

	// An OrderedPair represents a 2-tuple of values.
	OrderedPair struct {
		First  T
		Second T
	}
)

// makeSet
func makeSet() uSet {
	return make(uSet)
}

// Set
func (ls Set) ToSet() ISet {
	set := NewSet()
	for _, v := range ls {
		set.Add(v)
	}
	return set
}

// interface Stringer
func (ls Set) String() string {
	ts := make([]string, 0, len(ls))
	for _, elem := range ls {
		if fmt.Sprintf("%T", elem) == "set.Element" {
			ts = append(ts, fmt.Sprintf("%v", elem.D))
		} else {
			ts = append(ts, fmt.Sprintf("%v", elem))
		}
	}
	return fmt.Sprintf("{{%s}}", strings.Join(ts, ", "))
}

// Equal says whether two 2-tuples contain the same values in the same order.
func (pair *OrderedPair) Equal(other OrderedPair) bool {
	if pair.First == other.First &&
		pair.Second == other.Second {
		return true
	}
	return false
}

func (set *uSet) Add(Ts ...T) bool {
	if v, ok := Ts[0].([]T); ok {
		for _, f := range v {
			_, found := (*set)[f]
			if found {
				continue //False if it existed already
			}
			(*set)[f] = struct{}{} //Element{f}

		}
	} else if v, ok := Ts[0].([]Element); ok {
		for _, f := range v {
			_, found := (*set)[f]
			if found {
				continue //False if it existed already
			}
			(*set)[f] = struct{}{} //f)
		}
	} else {
		for _, v := range Ts {
			_, found := (*set)[v]
			if found {
				continue //False if it existed already
			}
			(*set)[v] = struct{}{}
		}
	}
	return true
}

// Contiene un elemento?
func (set *uSet) Contains(i ...T) bool {
	for _, val := range i {
		if _, ok := (*set)[val]; !ok {
			return false
		}
	}
	return true
}

// Remove a single element from the set.
func (set *uSet) Remove(key T) error { // TODO: Cambiar el return. Poner clave de retorno
	_, exists := (*set)[key]
	if !exists {
		return fmt.Errorf("Remove Error: T no existe en el Element")
	}
	delete(*set, key)
	return nil
}

func (set *uSet) IsSubset(other ISet) bool {
	_ = other.(*uSet)
	if set.Cardinality() > other.Cardinality() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (set *uSet) IsProperSubset(other ISet) bool {
	return set.IsSubset(other) && !set.Equal(other)
}

func (set *uSet) IsSuperset(other ISet) bool {
	return other.IsSubset(set)
}

func (set *uSet) IsProperSuperset(other ISet) bool {
	return set.IsSuperset(other) && !set.Equal(other)
}

func (set *uSet) Union(other ISet) ISet {
	o := other.(*uSet)

	unionedSet := makeSet()

	for elem := range *set {
		unionedSet.Add(elem)
	}
	for elem := range *o {
		unionedSet.Add(elem)
	}
	return &unionedSet
}

func (set *uSet) Intersect(other ISet) ISet {
	o := other.(*uSet)

	intersection := makeSet()
	// loop over smaller set
	if set.Cardinality() < other.Cardinality() {
		for elem := range *set {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range *o {
			if set.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return &intersection
}

// of the method. Otherwise, Difference will
// panic.
func (set *uSet) Difference(other ISet) ISet {
	_ = other.(*uSet)

	difference := makeSet()
	for elem := range *set {
		if !other.Contains(elem) {
			difference.Add(elem)
		}
	}
	return &difference
}

// Returns a new set with all elements which are
// in either this set or the other set but not in both.
//
// Note that the argument to SymmetricDifference
// must be of the same type as the receiver
// of the method. Otherwise, SymmetricDifference
// will panic.
func (set *uSet) SymmetricDifference(other ISet) ISet {
	_ = other.(*uSet)

	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Removes all elements from the set, leaving
// the empty set.
func (set *uSet) Clear() {
	*set = makeSet()
}

// Returns the number of elements in the set.
func (set *uSet) Cardinality() int {
	return len(*set)
}

// Returns the number of elements in the set.
//var Size = (*uSet).Cardinality

// Iterates over elements and executes the passed func against each element.
// If passed func returns true, stop iteration at the time.
func (set *uSet) Each(cb func(T) bool) {
	for elem := range *set {
		if cb(elem) {
			break
		}
	}
}

// Returns a channel of elements that you can
// range over.
func (set *uSet) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for elem := range *set {
			ch <- elem
		}
		close(ch)
	}()
	return ch
}

// Returns an Iterator object that you can
// use to range over the set.
func (set *uSet) Iterator() *Iterator {
	iterator, ch, stopCh := newIterator()
	go func() {
	L:
		for elem := range *set {
			select {
			case <-stopCh:
				break L
			case ch <- elem:
			}
		}
		close(ch)
	}()
	return iterator
}

// Determines if two sets are equal to each
// other. If they have the same cardinality
// and contain the same elements, they are
// considered equal. The order in which
// the elements were added is irrelevant.
//
// Note that the argument to Equal must be
// of the same type as the receiver of the
// method. Otherwise, Equal will panic.
func (set *uSet) Equal(other ISet) bool {
	_ = other.(*uSet)
	if set.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Returns a clone of the set using the same
// implementation, duplicating all keys.
func (set *uSet) Clone() ISet {
	clonedSet := makeSet()
	for elem := range *set {
		clonedSet.Add(elem)
	}
	return &clonedSet
}

// Provides a convenient string representation
// of the current state of the set.
func (set *uSet) String() string {
	ts := make([]string, 0, len(*set))
	for elem := range *set {
		if v, ok := elem.(Element); ok {
			ts = append(ts, fmt.Sprintf("%v", v.D))
		} else {
			ts = append(ts, fmt.Sprintf("%v", elem))
		}
	}
	return fmt.Sprintf("{{%s}}", strings.Join(ts, ", "))
}

// Pop removes and returns an arbitrary T from the set.
func (set *uSet) Pop() T {
	for T := range *set {
		delete(*set, T)
		return T
	}
	return nil
}

// Returns all subsets of a given set (Power Element).
func (set *uSet) PowerSet() ISet {
	op := makeSet()
	powSet := &op
	nullset := makeSet()
	powSet.Add(&nullset)
	for es := range *set {
		u := makeSet()
		j := powSet.Iter()
		for er := range j {
			p := makeSet()
			if reflect.TypeOf(er).Name() == "" {
				k := er.(*uSet)
				for ek := range *(k) {
					p.Add(ek)
				}
			} else {
				p.Add(er)
			}
			p.Add(es)
			u.Add(&p)
		}

		powSet = powSet.Union(&u).(*uSet)
	}
	return powSet
}

// Returns the Cartesian Product of two sets.
func (set *uSet) CartesianProduct(other ISet) ISet {
	o := other.(*uSet)
	op := makeSet()
	cartProduct := &op
	for i := range *set {
		for j := range *o {
			elem := OrderedPair{First: i, Second: j}
			cartProduct.Add(elem)
		}
	}
	return cartProduct
}

// Returns the members of the set as a slice.
func (set *uSet) ToSlice() []Element {
	keys := make([]Element, 0, set.Cardinality())
	for elem := range *set {
		keys = append(keys, Element{elem})
	}
	return keys
}

// MarshalJSON creates a JSON array from the set, it marshals all elements
func (set *uSet) MarshalJSON() ([]byte, error) {
	ts := make([]string, 0, set.Cardinality())
	for elem := range *set {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}
		ts = append(ts, string(b))
	}
	return []byte(fmt.Sprintf("[%s]", strings.Join(ts, ","))), nil
}

// UnmarshalJSON recreates a set from a JSON array, it only decodes
// primitive types. Numbers are decoded as json.Number.
func (set *uSet) UnmarshalJSON(b []byte) error {
	var i []interface{}
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&i)
	if err != nil {
		return err
	}
	for _, v := range i {
		switch t := v.(type) {
		case []interface{}, map[string]interface{}:
			continue
		default:
			set.Add(t)
		}
	}
	return nil
}
