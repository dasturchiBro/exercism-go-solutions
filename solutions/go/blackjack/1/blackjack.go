package blackjack

func ParseCard(card string) int {
    cards := map[string]int{
        "ace": 11,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
        "ten": 10,
        "jack": 10,
        "queen": 10,
        "king": 10,
    }
	c, ok := cards[card]
    if !ok {
        return 0
    }
    return c
}

func FirstTurn(c1, c2, d string) string {
    cSum := ParseCard(c1) + ParseCard(c2)
    if c1 == "ace" && c2 == "ace" {
        return "P"
    } else if cSum == 21 && (ParseCard(d) != 10 && ParseCard(d) != 11) {
        return "W"
    } else if (cSum >= 17 && cSum <=20) || (ParseCard(d) < 7 && cSum >= 12 && cSum <= 16) || cSum == 21 && (ParseCard(d) == 10 || ParseCard(d) == 11) {
        return "S"
    } else if cSum <= 11 || ParseCard(d) >= 7  {
        return "H"
    }
    return "H"
}