package api

import (
	"encoding/gob"
	"github.com/cespare/xxhash"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genset=true
type NamespacedName struct {
	Namespace string
	Name      string
}

// +genset=true
type Matcher struct {
	Name      string
	Namespace string
	Selector  *metav1.LabelSelector
}

func (m *Matcher) MapIndex() uint64 {
	if m == nil {
		return 0
	}
	h := xxhash.New()
	_ = gob.NewEncoder(h).Encode(m)
	return h.Sum64()
}

func (m Matcher) Compare(other Matcher) int {
	if m.Name < other.Name {
		return -1
	}
	if m.Name > other.Name {
		return +1
	}
	if m.Namespace < other.Namespace {
		return -1
	}
	if m.Namespace > other.Namespace {
		return +1
	}
	lhsIdx := m.MapIndex()
	rhsIdx := other.MapIndex()
	if lhsIdx < rhsIdx {
		return -1
	} else if lhsIdx > rhsIdx {
		return +1
	}
	return 0
}

func (m Matcher) Equal(other Matcher) bool {
	return m.MapIndex() == other.MapIndex()
}
