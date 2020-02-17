package jobs

import (
    "fmt"
    "testing"
)

var unsorted Queue
var sorted Queue

func TestSortQueue(t *testing.T) {
    node1 := Node{Schedule:80}
    node2 := Node{Schedule:1204}
    node3 := Node{Schedule:365}
    node4 := Node{Schedule:201}
    sorted.Items = append(sorted.Items, node1, node4, node3, node2)
    unsorted.Items = append(unsorted.Items, node1, node2, node3, node4)
    cases := []struct {Name string; A []Node; Expected []Node} {
        {"sort an unsorted array of nodes", unsorted.Items, sorted.Items},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
                actual := sortQueue(testCase.A)
                for inc, _ := range unsorted.Items {
                    if (actual[inc].Schedule) != testCase.Expected[inc].Schedule {t.Errorf("TestGetYaml %d failed - expected: '%v' got: '%v'", i+1, actual[inc].Schedule, testCase.Expected[inc].Schedule)}
                }
        })
    }
}
