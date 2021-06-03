package domain

import (
	"math/rand"
	"time"
)

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

	secretColours := make(map[Colour]int)
	for _, secretColour := range secret {
		secretColours[secretColour] += 1
	}

	// First sweep for perfect matches
	for i := 0; i < len(secret); i++ {
		guessColour := guess[i]
		secretColour := secret[i]
		if guessColour == secretColour {
			secretColours[guessColour] -= 1
		}
	}
	// Second sweep to assign result
	for i := 0; i < len(secret); i++ {
		guessColour := guess[i]
		secretColour := secret[i]
		if guessColour == secretColour {
			result[i] = Perfect
		} else if remaining, ok := secretColours[guessColour]; ok && remaining > 0 {
			result[i] = GoodColourBadPosition
		} else {
			result[i] = Nothing
		}
	}

	// Mix the evaluation around ;)
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })

	return &result
}

func (e *EngineImpl) CreateSecret() *Combination {
	var result Combination
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(result); i++ {
		result[i] = Colour(rand.Intn(len(ShortCodeToColour)))
	}
	return &result
}
