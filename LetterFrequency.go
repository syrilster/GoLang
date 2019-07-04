package main

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	euro = `Freude schöner Götterfunken
Tochter aus Elysium,
Wir betreten feuertrunken,
Himmlische, dein Heiligtum!
Deine Zauber binden wieder
Was die Mode streng geteilt;
Alle Menschen werden Brüder,
Wo dein sanfter Flügel weilt.`

	dutch = `Wilhelmus van Nassouwe
ben ik, van Duitsen bloed,
den vaderland getrouwe
blijf ik tot in den dood.
Een Prinse van Oranje
ben ik, vrij, onverveerd,
den Koning van Hispanje
heb ik altijd geëerd.`

	us = `O say can you see by the dawn's early light,
What so proudly we hailed at the twilight's last gleaming,
Whose broad stripes and bright stars through the perilous fight,
O'er the ramparts we watched, were so gallantly streaming?
And the rockets' red glare, the bombs bursting in air,
Gave proof through the night that our flag was still there;
O say does that star-spangled banner yet wave,
O'er the land of the free and the home of the brave?`
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}

	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	channel := make(chan FreqMap)
	result := FreqMap{}

	for _, inputString := range texts {
		go func(inputString string) {
			channel <- Frequency(inputString)
		}(inputString)
	}

	for range texts {
		for letter, count := range <-channel {
			result[letter] += count
		}
	}

	return result
}

func main() {
	seq := Frequency(euro + dutch + us)
	con := ConcurrentFrequency([]string{euro, dutch, us})
	if !reflect.DeepEqual(con, seq) {
		e := errors.New("ConcurrentFrequency wrong result")
		fmt.Print(e)
	} else {
		fmt.Println("All cases PASSED !!")
	}
}
