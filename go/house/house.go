// Package house outputs the nursery rhyme 'This is the House that Jack Built'.
package house

const testVersion = 1

type nounPhrase struct {
	determiner, complementation, noun, eol string
}

var nounPhrases = []nounPhrase{
	{complementation: "that lay in the ", noun: "house that Jack built", eol: "."},
	{complementation: "that ate the ", noun: "malt", eol: "\n"},
	{complementation: "that killed the ", noun: "rat", eol: "\n"},
	{complementation: "that worried the ", noun: "cat", eol: "\n"},
	{complementation: "that tossed the ", noun: "dog", eol: "\n"},
	{complementation: "that milked the ", noun: "cow with the crumpled horn", eol: "\n"},
	{complementation: "that kissed the ", noun: "maiden all forlorn", eol: "\n"},
	{complementation: "that married the ", noun: "man all tattered and torn", eol: "\n"},
	{complementation: "that woke the ", noun: "priest all shaven and shorn", eol: "\n"},
	{complementation: "that kept the ", noun: "rooster that crowed in the morn", eol: "\n"},
	{complementation: "that belonged to the ", noun: "farmer sowing his corn", eol: "\n"},
	{complementation: "killed the ", noun: "horse and the hound and the horn", eol: "\n"},
}

// Verse returns a string of the verse number specified
func Verse(n int) string {
	var verse string
	first := true

	for i := n - 1; i > -1; i-- {
		if first {
			first = false
			verse += "This is the " + nounPhrases[i].noun + nounPhrases[i].eol
		} else {
			verse += nounPhrases[i].complementation + nounPhrases[i].noun + nounPhrases[i].eol
		}

	}

	return verse

}

// Song returns all verses of the nursery rhyme
func Song() string {
	var song string
	for v := 1; v < 13; v++ {
		song += Verse(v)
		if v != 12 {
			song += "\n\n"
		}
	}
	return song
}
