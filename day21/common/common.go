package common

import (
	"adventofcode2024/utils"
	"math"
)

var numericKeyPad = map[rune]utils.Position{
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	'0': {3, 1},
	'A': {3, 2},
}

var directionalKeyPad = map[rune]utils.Position{
	'^': {0, 1},
	'A': {0, 2},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}

var directionChar = map[utils.Direction]rune{
	{0, 1}:  '>',
	{1, 0}:  'v',
	{0, -1}: '<',
	{-1, 0}: '^',
}

type MoveKey struct{ start, end utils.Position }
type CacheKey struct {
	code  string
	depth int
}

func SequenceLength(code []rune, indir int, maxindir int, cache map[CacheKey]int, movesMap map[MoveKey][][]rune) (_len int) {
	var key = CacheKey{string(code), indir}
	if l, found := cache[key]; found {
		return l
	}
	var from = numericKeyPad['A']
	var pad = numericKeyPad
	if indir > 0 {
		from = directionalKeyPad['A']
		pad = directionalKeyPad
	}
	for _, char := range code {
		var to = pad[char]
		if indir > 0 {
			var found bool
			var key = MoveKey{from, to}
			if sequences, found = movesMap[key]; !found {
				sequences = make([][]rune, 0)
				sequence = make([]rune, 0)
				gen(from, to, pad)
				movesMap[key] = sequences
			}
		} else {
			sequences = make([][]rune, 0)
			sequence = make([]rune, 0)
			gen(from, to, pad)
		}
		var maxLength = math.MaxInt
		if indir == maxindir {
			for _, seq := range sequences {
				if len(seq) < maxLength {
					maxLength = len(seq)
				}
			}
		} else {
			for _, sequence := range sequences {
				var length = SequenceLength(sequence, indir+1, maxindir, cache, movesMap)
				if length < maxLength {
					maxLength = length
				}
			}
		}
		_len += maxLength
		from = to
	}
	cache[key] = _len
	return
}

var sequence []rune
var sequences [][]rune

func gen(from utils.Position, to utils.Position, pad map[rune]utils.Position) {
	var seq = make([]rune, len(sequence))
	copy(seq, sequence)
	if to == from {
		seq = append(seq, 'A')
		sequences = append(sequences, seq)
		return
	}
	var dist = utils.Abs(from.Row-to.Row) + utils.Abs(from.Col-to.Col)
	for _, direction := range utils.UpRightDownLeft {
		var next = utils.Position{Row: from.Row + direction[0], Col: from.Col + direction[1]}
		if utils.Abs(next.Row-to.Row)+utils.Abs(next.Col-to.Col) < dist {
			var found = false
			for _, v := range pad {
				if v == next {
					found = true
					break
				}
			}
			if !found {
				continue
			}
			sequence = append(sequence, directionChar[direction])
			gen(next, to, pad)
			sequence = sequence[0 : len(sequence)-1]
		}
	}
}
