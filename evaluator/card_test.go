package evaluator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cardtests = []struct {
	in  string
	out interface{}
}{
	{"AS", Card{Suit: Spades, Value: Ace}},
	{"2S", Card{Suit: Spades, Value: Deuce}},
	{"3S", Card{Suit: Spades, Value: Trey}},
	{"4S", Card{Suit: Spades, Value: Four}},
	{"5S", Card{Suit: Spades, Value: Five}},
	{"6S", Card{Suit: Spades, Value: Six}},
	{"7S", Card{Suit: Spades, Value: Seven}},
	{"8S", Card{Suit: Spades, Value: Eight}},
	{"9S", Card{Suit: Spades, Value: Nine}},
	{"10S", Card{Suit: Spades, Value: Ten}},
	{"JS", Card{Suit: Spades, Value: Jack}},
	{"QS", Card{Suit: Spades, Value: Queen}},
	{"KS", Card{Suit: Spades, Value: King}},
	{"AH", Card{Suit: Hearts, Value: Ace}},
	{"2H", Card{Suit: Hearts, Value: Deuce}},
	{"3H", Card{Suit: Hearts, Value: Trey}},
	{"4H", Card{Suit: Hearts, Value: Four}},
	{"5H", Card{Suit: Hearts, Value: Five}},
	{"6H", Card{Suit: Hearts, Value: Six}},
	{"7H", Card{Suit: Hearts, Value: Seven}},
	{"8H", Card{Suit: Hearts, Value: Eight}},
	{"9H", Card{Suit: Hearts, Value: Nine}},
	{"10H", Card{Suit: Hearts, Value: Ten}},
	{"JH", Card{Suit: Hearts, Value: Jack}},
	{"QH", Card{Suit: Hearts, Value: Queen}},
	{"KH", Card{Suit: Hearts, Value: King}},
	{"AC", Card{Suit: Clubs, Value: Ace}},
	{"2C", Card{Suit: Clubs, Value: Deuce}},
	{"3C", Card{Suit: Clubs, Value: Trey}},
	{"4C", Card{Suit: Clubs, Value: Four}},
	{"5C", Card{Suit: Clubs, Value: Five}},
	{"6C", Card{Suit: Clubs, Value: Six}},
	{"7C", Card{Suit: Clubs, Value: Seven}},
	{"8C", Card{Suit: Clubs, Value: Eight}},
	{"9C", Card{Suit: Clubs, Value: Nine}},
	{"10C", Card{Suit: Clubs, Value: Ten}},
	{"JC", Card{Suit: Clubs, Value: Jack}},
	{"QC", Card{Suit: Clubs, Value: Queen}},
	{"KC", Card{Suit: Clubs, Value: King}},
	{"AD", Card{Suit: Diamonds, Value: Ace}},
	{"2D", Card{Suit: Diamonds, Value: Deuce}},
	{"3D", Card{Suit: Diamonds, Value: Trey}},
	{"4D", Card{Suit: Diamonds, Value: Four}},
	{"5D", Card{Suit: Diamonds, Value: Five}},
	{"6D", Card{Suit: Diamonds, Value: Six}},
	{"7D", Card{Suit: Diamonds, Value: Seven}},
	{"8D", Card{Suit: Diamonds, Value: Eight}},
	{"9D", Card{Suit: Diamonds, Value: Nine}},
	{"10D", Card{Suit: Diamonds, Value: Ten}},
	{"JD", Card{Suit: Diamonds, Value: Jack}},
	{"QD", Card{Suit: Diamonds, Value: Queen}},
	{"KD", Card{Suit: Diamonds, Value: King}},
	{"11S", nil},
	{"AA", nil},
	{"13H", nil},
	{"1D", nil},
}

func TestParseCard(t *testing.T) {
	for _, ct := range cardtests {
		card, err := ParseCard(ct.in)
		if ct.out != nil {
			if assert.NoError(t, err, "it should have parsed succesfully") {
				assert.Equal(t, ct.out, *card)
			}
		} else {
			assert.Error(t, err, "it should have failed")
		}
	}
}

var lesstests = []struct {
	in  Card
	arg Card
	out bool
}{
	{Card{Suit: Spades, Value: Deuce}, Card{Suit: Spades, Value: Trey}, true},
	{Card{Suit: Hearts, Value: Trey}, Card{Suit: Clubs, Value: Four}, true},
	{Card{Suit: Clubs, Value: Four}, Card{Suit: Hearts, Value: Five}, true},
	{Card{Suit: Spades, Value: Five}, Card{Suit: Spades, Value: Six}, true},
	{Card{Suit: Hearts, Value: Six}, Card{Suit: Diamonds, Value: Seven}, true},
	{Card{Suit: Clubs, Value: Seven}, Card{Suit: Hearts, Value: Eight}, true},
	{Card{Suit: Spades, Value: Eight}, Card{Suit: Spades, Value: Nine}, true},
	{Card{Suit: Diamonds, Value: Nine}, Card{Suit: Diamonds, Value: Ten}, true},
	{Card{Suit: Hearts, Value: Ten}, Card{Suit: Spades, Value: Jack}, true},
	{Card{Suit: Clubs, Value: Jack}, Card{Suit: Hearts, Value: Queen}, true},
	{Card{Suit: Spades, Value: Queen}, Card{Suit: Clubs, Value: King}, true},
	{Card{Suit: Diamonds, Value: King}, Card{Suit: Diamonds, Value: Ace}, true},
	{Card{Suit: Hearts, Value: Deuce}, Card{Suit: Spades, Value: Ace}, true},
	{Card{Suit: Spades, Value: Trey}, Card{Suit: Spades, Value: Deuce}, false},
	{Card{Suit: Clubs, Value: Four}, Card{Suit: Hearts, Value: Trey}, false},
	{Card{Suit: Hearts, Value: Five}, Card{Suit: Clubs, Value: Four}, false},
	{Card{Suit: Spades, Value: Six}, Card{Suit: Spades, Value: Five}, false},
	{Card{Suit: Diamonds, Value: Seven}, Card{Suit: Hearts, Value: Six}, false},
	{Card{Suit: Hearts, Value: Eight}, Card{Suit: Clubs, Value: Seven}, false},
	{Card{Suit: Spades, Value: Nine}, Card{Suit: Spades, Value: Eight}, false},
	{Card{Suit: Diamonds, Value: Ten}, Card{Suit: Diamonds, Value: Nine}, false},
	{Card{Suit: Spades, Value: Jack}, Card{Suit: Hearts, Value: Ten}, false},
	{Card{Suit: Hearts, Value: Queen}, Card{Suit: Clubs, Value: Jack}, false},
	{Card{Suit: Clubs, Value: King}, Card{Suit: Spades, Value: Queen}, false},
	{Card{Suit: Diamonds, Value: Ace}, Card{Suit: Diamonds, Value: King}, false},
	{Card{Suit: Spades, Value: Ace}, Card{Suit: Hearts, Value: Deuce}, false},
}

func TestCardLess(t *testing.T) {
	for _, lt := range lesstests {
		assert.Equal(t, lt.out, lt.in.Less(lt.arg), fmt.Sprintf("Expected %+v .Less %+v to be %t", lt.in, lt.arg, lt.out))
	}
}
