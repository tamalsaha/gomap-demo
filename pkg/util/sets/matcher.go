/*
Copyright The Kubernetes Authors.

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

// Code generated by set-gen. DO NOT EDIT.

package sets

import (
	"reflect"
	"sort"

	"github.com/tamalsaha/gomap-demo/pkg/api"
)

// sets.Matcher is a set of api.Matchers, implemented via map[sets.Matcher]struct{} for minimal memory consumption.
type Matcher struct {
	keys   map[uint64]*api.Matcher
	values map[uint64]Empty
}

// NewMatcher creates a *Matcher from a list of values.
func NewMatcher(items ...*api.Matcher) *Matcher {
	ss := &Matcher{
		keys:   map[uint64]*api.Matcher{},
		values: map[uint64]Empty{},
	}
	ss.Insert(items...)
	return ss
}

// *MatcherKeySet creates a *Matcher from a keys of a map[sets.Matcher](? extends interface{}).
// If the value passed in is not actually a map, this will panic.
func MatcherKeySet(theMap interface{}) *Matcher {
	v := reflect.ValueOf(theMap)
	ret := &Matcher{
		keys:   map[uint64]*api.Matcher{},
		values: map[uint64]Empty{},
	}

	for _, keyValue := range v.MapKeys() {
		ret.Insert(keyValue.Interface().(*api.Matcher))
	}
	return ret
}

// Insert adds items to the set.
func (s *Matcher) Insert(items ...*api.Matcher) *Matcher {
	for idx, item := range items {
		kidx := item.MapIndex()
		s.keys[kidx] = items[idx]
		s.values[kidx] = Empty{}
	}
	return s
}

// Delete removes all items from the set.
func (s *Matcher) Delete(items ...*api.Matcher) *Matcher {
	for _, item := range items {
		kidx := item.MapIndex()
		delete(s.keys, kidx)
		delete(s.values, kidx)
	}
	return s
}

// Has returns true if and only if item is contained in the set.
func (s *Matcher) Has(item *api.Matcher) bool {
	_, contained := s.values[item.MapIndex()]
	return contained
}

// HasAll returns true if and only if all items are contained in the set.
func (s *Matcher) HasAll(items ...*api.Matcher) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s *Matcher) HasAny(items ...*api.Matcher) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

// Difference returns a set of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}
func (s *Matcher) Difference(s2 *Matcher) *Matcher {
	result := NewMatcher()
	for kidx, key := range s.keys {
		if _, exists := s2.keys[kidx]; !exists {
			result.Insert(key)
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}
func (s1 *Matcher) Union(s2 *Matcher) *Matcher {
	result := NewMatcher()
	for _, key := range s1.keys {
		result.Insert(key)
	}
	for _, key := range s2.keys {
		result.Insert(key)
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 *Matcher) Intersection(s2 *Matcher) *Matcher {
	var walk, other *Matcher
	result := NewMatcher()
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for kidx, key := range walk.keys {
		if _, exists := other.keys[kidx]; exists {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 *Matcher) IsSuperset(s2 *Matcher) bool {
	for kidx := range s2.keys {
		if _, exists := s1.keys[kidx]; !exists {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two api are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 *Matcher) Equal(s2 *Matcher) bool {
	return len(s1.keys) == len(s2.keys) && s1.IsSuperset(s2)
}

type sortableSliceOfMatcher []*api.Matcher

func (s sortableSliceOfMatcher) Len() int           { return len(s) }
func (s sortableSliceOfMatcher) Less(i, j int) bool { return lessMatcher(s[i], s[j]) }
func (s sortableSliceOfMatcher) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// List returns the contents as a sorted sets.Matcher slice.
func (s *Matcher) List() []*api.Matcher {
	res := make(sortableSliceOfMatcher, 0, len(s.keys))
	for kidx := range s.keys {
		res = append(res, s.keys[kidx])
	}
	sort.Sort(res)
	return []*api.Matcher(res)
}

// UnsortedList returns the slice with contents in random order.
func (s *Matcher) UnsortedList() []*api.Matcher {
	res := make([]*api.Matcher, 0, len(s.keys))
	for kidx := range s.keys {
		res = append(res, s.keys[kidx])
	}
	return res
}

// Returns a single element from the set.
func (s *Matcher) PopAny() (*api.Matcher, bool) {
	for kidx, key := range s.keys {
		delete(s.keys, kidx)
		delete(s.values, kidx)
		return key, true
	}
	return nil, false
}

// Len returns the size of the set.
func (s *Matcher) Len() int {
	return len(s.keys)
}

func lessMatcher(lhs, rhs *api.Matcher) bool {
	return lhs.Compare(*rhs) < 0
}