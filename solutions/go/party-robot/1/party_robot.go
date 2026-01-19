package partyrobot

import "fmt"

func Welcome(msg string) string {
    return fmt.Sprintf("Welcome to my party, %s!", msg)
}

func HappyBirthday(name string, age int) string {
    return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, number int, name2 string, direct string, distance float64) string {
    msg := fmt.Sprintf(
        "Welcome to my party, %s!\nYou have been assigned to table %.3d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.",
        name,
        number,
        direct,
        distance,
        name2,
    )
    return msg
}