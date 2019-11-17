package meander

import (
	"errors"
	"strings"
)

// Cost ...
type Cost int8

const (
	_ Cost = iota
	// Cost1 ...
	Cost1
	// Cost2 ...
	Cost2
	// Cost3 ...
	Cost3
	// Cost4 ...
	Cost4
	// Cost5 ...
	Cost5
)

var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (l Cost) String() string {
	for s, v := range costStrings {
		if l == v {
			return s
		}
	}
	return "invalid"
}

// ParseCost ...
func ParseCost(s string) Cost {
	return costStrings[s]
}

// CostRange ...
type CostRange struct {
	From Cost
	To   Cost
}

// String ...
func (r CostRange) String() string {
	return r.From.String() + "..." + r.To.String()
}

// ParseCostRange ...
func ParseCostRange(s string) (CostRange, error) {
	var r CostRange
	segs := strings.Split(s, "...")
	if len(segs) != 2 {
		return r, errors.New("invalid cost range")
	}
	r.From = ParseCost(segs[0])
	r.To = ParseCost(segs[1])
	return r, nil
}
