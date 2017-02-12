// Package twelve outputs the lyrics to 'The Twelve Days of Christmas'
package twelve

import (
	"fmt"
)

const testVersion = 1

type lyric struct {
	gift    string
	ordinal string
}

var lyrics = map[int]lyric{
	1:  lyric{"a Partridge in a Pear Tree", "first"},
	2:  lyric{"two Turtle Doves", "second"},
	3:  lyric{"three French Hens", "third"},
	4:  lyric{"four Calling Birds", "fourth"},
	5:  lyric{"five Gold Rings", "fifth"},
	6:  lyric{"six Geese-a-Laying", "sixth"},
	7:  lyric{"seven Swans-a-Swimming", "seventh"},
	8:  lyric{"eight Maids-a-Milking", "eighth"},
	9:  lyric{"nine Ladies Dancing", "ninth"},
	10: lyric{"ten Lords-a-Leaping", "tenth"},
	11: lyric{"eleven Pipers Piping", "eleventh"},
	12: lyric{"twelve Drummers Drumming", "twelfth"},
}

// Song returns a a string of the full song the Twelve Days of Christmas.
func Song() string {
	var song string

	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}

	return song
}

// Verse returns one lyric of the song the Twelve Days of Christmas.
func Verse(s int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s.", lyrics[s].ordinal, gifts(s))
}

func gifts(max int) string {
	var gifts string

	for i := 1; i <= max; i++ {
		if i == 2 {
			gifts = fmt.Sprintf("%s %s", "and", gifts)
		}
		gifts = fmt.Sprintf("%s, %s", lyrics[i].gift, gifts)
		if i == 1 {
			gifts = lyrics[i].gift
		}
	}

	return gifts
}
