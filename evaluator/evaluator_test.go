package evaluator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHand struct {
	name  string
	cards []string
}

// Joe, 3H, 4H, 5H, 6H, 8H
// Bob, 3C, 3D, 3S, 8C, 10D
// Sally, AC, 10C, 5C, 2S, 2C
var evaluatortests = []struct {
	hands []testHand
	out   interface{}
}{
	{[]testHand{
		testHand{name: "Joe", cards: []string{"3H", "4H", "5H", "6H", "8H"}},
		testHand{name: "Bob", cards: []string{"3C", "3D", "3S", "8C", "10D"}},
		testHand{name: "Sally", cards: []string{"AC", "10C", "5C", "2S", "2C"}},
	}, "Joe"},
	{[]testHand{
		testHand{name: "Joe", cards: []string{"3H", "4H", "5H", "6H", "8H"}},
		testHand{name: "Bob", cards: []string{"3C", "4C", "5C", "6C", "8C"}},
		testHand{name: "Sally", cards: []string{"AC", "10C", "5C", "2S", "2C"}},
	}, "Joe, Bob"},
	{[]testHand{
		testHand{name: "Joe", cards: []string{"3H", "4C", "5H", "6H", "8H"}},
		testHand{name: "Bob", cards: []string{"3C", "4H", "5C", "6C", "8C"}},
		testHand{name: "Sally", cards: []string{"AC", "10C", "5C", "4S", "2C"}},
	}, "Joe, Bob, Sally"},
	{[]testHand{
		testHand{name: "Joe", cards: []string{"3H", "4H", "5H", "6H", "8D"}},
		testHand{name: "Bob", cards: []string{"3C", "3D", "3S", "8C", "10D"}},
		testHand{name: "Sally", cards: []string{"AC", "10C", "5C", "2S", "2C"}},
	}, "Bob"},
	{[]testHand{
		testHand{name: "Joe", cards: []string{"3H", "4H", "5H", "6H", "8D"}},
		testHand{name: "Bob", cards: []string{"3C", "3D", "3S", "8C", "10D"}},
		testHand{name: "Sally", cards: []string{"AC", "10C", "2D", "2S", "2C"}},
	}, "Bob, Sally"},
	{[]testHand{
		testHand{name: "Joe", cards: []string{"3H", "4H", "5H", "6H", "8D"}},
		testHand{name: "Bob", cards: []string{"3C", "5D", "3S", "8C", "10D"}},
		testHand{name: "Sally", cards: []string{"AC", "10C", "2D", "2S", "2C"}},
	}, "Sally"},
	{[]testHand{
		testHand{name: "Joe", cards: []string{"3H", "4H", "5H", "6H", "8H"}},
		testHand{name: "Bob", cards: []string{"3C", "3D", "3S", "8C", "10D"}},
		testHand{name: "Sally", cards: []string{"AC", "10C", "5C", "2S", "2C"}},
		testHand{name: "Luke", cards: []string{"AC", "10C", "5C", "4S", "2C"}},
	}, "Joe"},
	{[]testHand{
		testHand{name: "Joe", cards: []string{"3H", "4H", "5H", "6H", "8H"}},
		testHand{name: "Bob", cards: []string{"3C", "3D", "3S", "8C", "10D"}},
		testHand{name: "Sally", cards: []string{"AC", "10C", "5C", "2S", "2C"}},
		testHand{name: "Luke", cards: []string{"AC", "10C", "5C", "4S", "2C"}},
		testHand{name: "Han", cards: []string{"10S", "JS", "QS", "KS", "AS"}},
	}, "Joe, Han"},
	{[]testHand{
		testHand{name: "Luke", cards: []string{"AC", "10C", "5C", "4S", "2C"}},
		testHand{name: "Han", cards: []string{"10S", "JS", "QS", "KS", "AS"}},
	}, "Han"},
	{[]testHand{}, nil},
}

func TestEvaluator(t *testing.T) {
	for _, et := range evaluatortests {
		e := Evaluator{}
		for _, hand := range et.hands {
			e.ParseHand(hand.name, hand.cards)
		}
		err := e.Decide()
		if et.out != nil {
			if assert.NoError(t, err, "not expecting to return error") {
				assert.Equal(t, et.out, e.Resolution(), fmt.Sprintf("Expected %+v to decide for %s", e, et.out))
			}
		} else {
			assert.Error(t, err, "expecting return of error")
		}
	}
}
