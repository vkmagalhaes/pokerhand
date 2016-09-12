# Pokerhand

Library which evaluates who are the winner(s) among several 5 card poker hands.

In this moment this project only implement a subset of the regular poker hands:

- Flush
- Three of a Kind
- One pair
- High Card

### Input

Collection of players in the showdown.

- Player Name
- 5 Cards (each specifying the card rank and suit of the card)

### Output

Collection of winning players (more than one in case of a tie)

### Example

Input:
- 3 (number of hands to be analyzed)
- Joe, 3H, 4H, 5H, 6H, 8H
- Bob, 3C, 3D, 3S, 8C, 10D
- Sally, AC, 10C, 5C, 2S, 2C

Output:
- Joe

## Dependencies

You must have `golang 1.7+` or `docker` installed.

## Running

Using go:
`make run`

Using docker:
`make run_docker`

## Testing

`make test`

## Building

`make build`

## Code Architecture

I use DDD to organize the code and keep it clean and enjoyable during its evolution.

The next links will help you get insights of how DDD should looks like in a go application:

1. http://www.citerus.se/go-ddd
2. http://www.citerus.se/part-2-domain-driven-design-in-go/
3. http://www.citerus.se/part-3-domain-driven-design-in-go/
