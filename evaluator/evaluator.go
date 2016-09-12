package evaluator

import (
	"errors"
	"strings"
)

// Evaluator is the main struct of the program
// It receives the poker hands and then decides which one is the winner
type Evaluator struct {
	Hands   []Hand
	Winners []Hand
}

// ErrEmptyHands is used when no hands were given to the evaluator.
var ErrEmptyHands = errors.New("No hands given")

// ErrHandNotComplete is used when the hand has not 5 cards.
var ErrHandNotComplete = errors.New("You must have 5 cards to a hand be complete")

// Decide evaluates a group of poker hands and decides which one is the winner
func (e *Evaluator) Decide() error {
	if len(e.Hands) == 0 {
		return ErrEmptyHands
	}

	winningHand := e.Hands[0]
	winners := []Hand{e.Hands[0]}
	for i, hand := range e.Hands {
		if i == 0 {
			continue
		}

		switch hand.Compare(winningHand) {
		case 1:
			winningHand = hand
			winners = []Hand{hand}
		case 0:
			winners = append(winners, hand)
		}

	}

	e.Winners = winners

	return nil
}

// ParseHand parses a string in the format "Playername, AS, 1H, 2H, JD, KC" and add it to
// his set of hands for evaluation
func (e *Evaluator) ParseHand(name string, cards []string) error {
	if len(cards) != 5 {
		return ErrHandNotComplete
	}

	hand := Hand{
		Player: name,
	}

	for _, c := range cards {
		card, err := ParseCard(c)
		if err != nil {
			return err
		}

		hand.Cards = append(hand.Cards, *card)
	}

	hand.Classify()
	e.Hands = append(e.Hands, hand)

	return nil
}

// Resolution gives a string with the player names of the winner hands
func (e *Evaluator) Resolution() string {
	var winners []string
	for _, winner := range e.Winners {
		winners = append(winners, winner.Player)
	}
	return strings.Join(winners, ", ")
}
