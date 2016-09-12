package evaluator

import (
	"errors"
	"regexp"
)

// CardSuit types
type CardSuit string

// CardValue types
type CardValue string

// Possible CardSuit and CardValue values
const (
	Spades   CardSuit = "S"
	Hearts   CardSuit = "H"
	Diamonds CardSuit = "D"
	Clubs    CardSuit = "C"

	Ace   CardValue = "A"
	Deuce CardValue = "2"
	Trey  CardValue = "3"
	Four  CardValue = "4"
	Five  CardValue = "5"
	Six   CardValue = "6"
	Seven CardValue = "7"
	Eight CardValue = "8"
	Nine  CardValue = "9"
	Ten   CardValue = "10"
	Jack  CardValue = "J"
	Queen CardValue = "Q"
	King  CardValue = "K"
)

var (
	// CardRegex used to validate parsing to card
	CardRegex = regexp.MustCompile(`\b(A|2|3|4|5|6|7|8|9|10|J|Q|K)(S|H|D|C)\b`)
	// ErrCardParse is used when was not possible to parse a string to a card
	ErrCardParse = errors.New("Could not parse string to card")
)

// Card of poker deck
type Card struct {
	Suit  CardSuit
	Value CardValue
}

// ParseCard receives a string in the format "(CardValue)(CardSuit)" and returns the
// Card which corresponds to it
func ParseCard(str string) (*Card, error) {
	res := CardRegex.FindAllStringSubmatch(str, 1) // ex: "3H" -> [[3H 3 H]]
	if len(res) == 0 {
		return nil, ErrCardParse
	}

	if len(res[0]) != 3 {
		return nil, ErrCardParse
	}

	card := &Card{
		Suit:  CardSuit(res[0][2]),
		Value: CardValue(res[0][1]),
	}

	return card, nil
}

// Less return true if c has value smaller than c2
func (c *Card) Less(c2 Card) bool {
	return c.intValue() < c2.intValue()
}

func (c *Card) intValue() int {
	switch c.Value {
	case Ace:
		return 14
	case Deuce:
		return 2
	case Trey:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Ten:
		return 10
	case Jack:
		return 11
	case Queen:
		return 12
	case King:
		return 13
	}
	return 0
}
