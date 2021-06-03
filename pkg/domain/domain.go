package domain

import "math/rand"

type Colour int

const (
	Red Colour = iota
	Blue
	Green
	Yellow
	Purple
	White
)

var ShortCodeToColour map[rune]Colour = map[rune]Colour{
	'R': Red,
	'B': Blue,
	'G': Green,
	'Y': Yellow,
	'P': Purple,
	'W': White,
}

var ColourToShortCode map[Colour]rune = map[Colour]rune{
	Red:    'R',
	Blue:   'B',
	Green:  'G',
	Yellow: 'Y',
	Purple: 'P',
	White:  'W',
}

type Rating int

const (
	Nothing Rating = iota
	GoodColourBadPosition
	Perfect
)

type Combination [4]Colour

type Evaluation [4]Rating

func (e Evaluation) IsPerfect() bool {
	for _, r := range e {
		if r != Perfect {
			return false
		}
	}
	return true
}

type Engine interface {
	Evaluate(secret *Combination, guess *Combination) *Evaluation
	CreateSecret() *Combination
}

type EngineImpl struct{}

var _ Engine = &EngineImpl{}

func NewEngineImpl() *EngineImpl {
	return &EngineImpl{}
}

func (e *EngineImpl) Evaluate(secret *Combination, guess *Combination) *Evaluation {
	var result Evaluation

	guessColours := make(map[Colour]bool)
	for _, guessColour := range guess {
		guessColours[guessColour] = true
	}

	for i := 0; i < len(secret); i++ {
		guessColour := guess[i]
		secretColour := secret[i]
		if guessColour == secretColour {
			result[i] = Perfect
		} else if _, ok := guessColours[guessColour]; ok {
			result[i] = GoodColourBadPosition
		} else {
			result[i] = Nothing
		}
	}
	return &result
}

func (e *EngineImpl) CreateSecret() *Combination {
	var result Combination
	for i := 0; i < len(result); i++ {
		result[i] = Colour(rand.Intn(len(ShortCodeToColour)))
	}
	return &result
}
