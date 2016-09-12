package evaluator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var comparetests = []struct {
	in  Hand
	arg Hand
	out int
}{
	{Hand{Category: Flush}, Hand{Category: Flush}, 0},
	{Hand{Category: Flush}, Hand{Category: ThreeOfAKind}, 1},
	{Hand{Category: Flush}, Hand{Category: OnePair}, 1},
	{Hand{Category: Flush}, Hand{Category: HighCard}, 1},
	{Hand{Category: ThreeOfAKind}, Hand{Category: Flush}, -1},
	{Hand{Category: ThreeOfAKind}, Hand{Category: ThreeOfAKind}, 0},
	{Hand{Category: ThreeOfAKind}, Hand{Category: OnePair}, 1},
	{Hand{Category: ThreeOfAKind}, Hand{Category: HighCard}, 1},
	{Hand{Category: OnePair}, Hand{Category: Flush}, -1},
	{Hand{Category: OnePair}, Hand{Category: ThreeOfAKind}, -1},
	{Hand{Category: OnePair}, Hand{Category: OnePair}, 0},
	{Hand{Category: OnePair}, Hand{Category: HighCard}, 1},
	{Hand{Category: HighCard}, Hand{Category: Flush}, -1},
	{Hand{Category: HighCard}, Hand{Category: ThreeOfAKind}, -1},
	{Hand{Category: HighCard}, Hand{Category: OnePair}, -1},
	{Hand{Category: HighCard}, Hand{Category: HighCard}, 0},
}

func TestHandCompare(t *testing.T) {
	for _, lt := range comparetests {
		assert.Equal(t, lt.out, lt.in.Compare(lt.arg), fmt.Sprintf("Expected %+v .Compare %+v to be %d", lt.in, lt.arg, lt.out))
	}
}

// Joe, 3H, 4H, 5H, 6H, 8H
// Bob, 3C, 3D, 3S, 8C, 10D
// Sally, AC, 10C, 5C, 2S, 2C
var classifytests = []struct {
	in          Hand
	category    HandCategory
	higherValue CardValue
}{
	{Hand{Cards: []Card{
		Card{Suit: Hearts, Value: Trey},
		Card{Suit: Hearts, Value: Four},
		Card{Suit: Hearts, Value: Five},
		Card{Suit: Hearts, Value: Six},
		Card{Suit: Hearts, Value: Eight},
	}}, Flush, Eight},
	{Hand{Cards: []Card{
		Card{Suit: Clubs, Value: Trey},
		Card{Suit: Diamonds, Value: Trey},
		Card{Suit: Spades, Value: Trey},
		Card{Suit: Clubs, Value: Eight},
		Card{Suit: Diamonds, Value: Ten},
	}}, ThreeOfAKind, Trey},
	{Hand{Cards: []Card{
		Card{Suit: Clubs, Value: Ace},
		Card{Suit: Clubs, Value: Ten},
		Card{Suit: Clubs, Value: Five},
		Card{Suit: Spades, Value: Deuce},
		Card{Suit: Clubs, Value: Deuce},
	}}, OnePair, Deuce},
	{Hand{Cards: []Card{
		Card{Suit: Clubs, Value: Ace},
		Card{Suit: Spades, Value: Ace},
		Card{Suit: Clubs, Value: Five},
		Card{Suit: Spades, Value: Deuce},
		Card{Suit: Clubs, Value: Deuce},
	}}, HighCard, Ace}, // TwoPair (unsupported)
	{Hand{Cards: []Card{
		Card{Suit: Clubs, Value: Ace},
		Card{Suit: Spades, Value: Ace},
		Card{Suit: Diamonds, Value: Ace},
		Card{Suit: Spades, Value: Deuce},
		Card{Suit: Clubs, Value: Deuce},
	}}, HighCard, Ace}, // FullHouse (unsupported)
}

func TestHandClassify(t *testing.T) {
	for _, lt := range classifytests {
		lt.in.Classify()
		assert.Equal(t, lt.category, lt.in.Category, fmt.Sprintf("Expected %+v Category to be %d", lt.in, lt.category))
		assert.Equal(t, lt.higherValue, lt.in.HigherCardValue, fmt.Sprintf("Expected %+v HigherCardValue to be %s", lt.in, lt.higherValue))
	}
}
