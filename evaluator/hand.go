package evaluator

// HandCategory defines the category of a poker hand
type HandCategory int

// Possible CardSuit and CardValue values
const (
	Flush        HandCategory = 0
	ThreeOfAKind HandCategory = 1
	OnePair      HandCategory = 2
	HighCard     HandCategory = 3
)

// Hand of poker of a player with five cards
type Hand struct {
	Player          string
	Cards           []Card
	Category        HandCategory
	HigherCardValue CardValue

	cardSuitTable  map[CardSuit]int
	cardValueTable map[CardValue]int
}

// Compare returns 1 if the hand1 is more valuable than hand2, 0 if they match and -1 otherwise
func (h Hand) Compare(h2 Hand) int {
	switch {
	case h.Category < h2.Category:
		return 1
	case h.Category == h2.Category:
		return 0
	}
	return -1
}

// Classify scans the hand cards and set the pokerhand classification for it
func (h *Hand) Classify() {
	h.buildCardTables()
	switch {
	case h.isFlush():
		h.Category = Flush
		return
	case h.isThreeOfAKind():
		h.Category = ThreeOfAKind
		return
	case h.isOnePair():
		h.Category = OnePair
		return
	}
	h.setHigherCardValue()
	h.Category = HighCard
}

func (h *Hand) isFlush() bool {
	if len(h.cardSuitTable) != 1 {
		return false
	}

	h.setHigherCardValue()
	return true
}

func (h *Hand) isThreeOfAKind() bool {
	if len(h.cardValueTable) != 3 {
		return false
	}

	for key, value := range h.cardValueTable {
		if value == 3 {
			h.HigherCardValue = key
			return true
		}
	}

	return false
}

func (h *Hand) isOnePair() bool {
	if len(h.cardValueTable) != 4 {
		return false
	}

	for key, value := range h.cardValueTable {
		if value == 2 {
			h.HigherCardValue = key
			return true
		}
	}

	return false
}

func (h *Hand) buildCardTables() {
	h.cardValueTable = make(map[CardValue]int)
	for _, card := range h.Cards {
		h.cardValueTable[card.Value] = h.cardValueTable[card.Value] + 1
	}

	h.cardSuitTable = make(map[CardSuit]int)
	for _, card := range h.Cards {
		h.cardSuitTable[card.Suit] = h.cardSuitTable[card.Suit] + 1
	}
}

func (h *Hand) setHigherCardValue() {
	higherCard := h.Cards[0]
	for _, card := range h.Cards {
		if higherCard.Less(card) {
			higherCard = card
		}
	}

	h.HigherCardValue = higherCard.Value
}
