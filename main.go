package main

import (
	"encoding/json"
	"fmt"
	"github.com/tamalsaha/gomap-demo/pkg/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cespare/xxhash"
	"github.com/google/go-cmp/cmp"
	"github.com/kr/pretty"
	"k8s.io/apimachinery/pkg/types"
)

func main() {
	m := map[api.Matcher][]types.NamespacedName{}
	fmt.Println("a", m)

	m2 := api.Matcher{
		Name:      "",
		Namespace: "",
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"a": "b",
			},
			MatchExpressions: nil,
		},
	}
	data, err := m2.Selector.Marshal()
	if err != nil {
		panic(err)
	}
	d2, err := json.Marshal(m2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println(string(d2))
	fmt.Printf("%# v", pretty.Formatter(m2))
	fmt.Println()
	fmt.Printf("%T", m2)

	var m3 *api.Matcher
	fmt.Println(m3.MapIndex())
	var m4 api.Matcher
	fmt.Println(m4.MapIndex())

	fmt.Println(m2.MapIndex())

	xxhash.Sum64String("")

	fmt.Println("------------")
	fmt.Println(cmp.Equal(m2, m4))
	fmt.Println(cmp.Equal(m2, m2))
	fmt.Println(cmp.Equal(m4, m4))
}
