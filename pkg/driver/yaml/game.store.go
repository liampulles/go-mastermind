package yaml

import (
	"fmt"
	"os"

	"github.com/liampulles/go-mastermind/pkg/domain"
	"github.com/liampulles/go-mastermind/pkg/usecase"
	"gopkg.in/yaml.v2"
)

const (
	gameFile = "./gomastermind.game"
)

type GameYAML struct {
	Secret [4]string
	Won    bool
}

type GameStore struct{}

var _ usecase.GameStore = &GameStore{}

func NewGameStore() *GameStore {
	return &GameStore{}
}

func (gs *GameStore) New(secret *domain.Combination) (usecase.GameIdentifier, error) {
	game := GameYAML{
		Secret: asStringCombination(secret),
		Won:    false,
	}
	if err := writeGame(usecase.GameIdentifier(gameFile), &game); err != nil {
		return usecase.GameIdentifier(""), err
	}
	return usecase.GameIdentifier(gameFile), nil
}

func (gs *GameStore) GetCurrentGameIdentifier() (usecase.GameIdentifier, error) {
	if _, err := os.Stat(gameFile); err != nil {
		if os.IsNotExist(err) {
			return usecase.GameIdentifier(""), fmt.Errorf("no current game exists")
		}
		return usecase.GameIdentifier(""), fmt.Errorf("could not identify state of game file: %w", err)
	}
	return gameFile, nil
}

func (gs *GameStore) IsWon(id usecase.GameIdentifier) (bool, error) {
	game, err := readGame(id)
	if err != nil {
		return false, err
	}

	return game.Won, nil
}

func (gs *GameStore) GetSecret(id usecase.GameIdentifier) (*domain.Combination, error) {
	game, err := readGame(id)
	if err != nil {
		return nil, err
	}

	comb, err := asDomainCombination(&game.Secret)
	if err != nil {
		return nil, fmt.Errorf("could not create combination: %w", err)
	}
	return comb, nil
}

func (gs *GameStore) EndGame(id usecase.GameIdentifier) error {
	game, err := readGame(id)
	if err != nil {
		return err
	}
	game.Won = true
	if err := writeGame(id, game); err != nil {
		return err
	}
	return nil
}

func readGame(id usecase.GameIdentifier) (*GameYAML, error) {
	bytes, err := os.ReadFile(string(id))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("no game exists for id %s", string(id))
		}
		return nil, fmt.Errorf("could not read game file: %w", err)
	}

	game := GameYAML{}
	if err := yaml.Unmarshal(bytes, &game); err != nil {
		return nil, fmt.Errorf("could not parse game file YAML: %w", err)
	}
	return &game, nil
}

func writeGame(id usecase.GameIdentifier, game *GameYAML) error {
	bytes, err := yaml.Marshal(&game)
	if err != nil {
		return fmt.Errorf("could not marshal game to YAML: %w", err)
	}
	if err := os.WriteFile(string(id), bytes, os.ModeAppend); err != nil {
		return fmt.Errorf("could not write YAML to file: %w", err)
	}
	return nil
}

func asStringCombination(comb *domain.Combination) [4]string {
	var result [4]string
	for i, col := range comb {
		result[i] = string(domain.ColourToShortCode[col])
	}
	return result
}

func asDomainCombination(comb *[4]string) (*domain.Combination, error) {
	var result domain.Combination
	for i, str := range comb {
		r, err := domain.StringToRune(str)
		if err != nil {
			return nil, fmt.Errorf("could not convert %s to rune: %w", str, err)
		}
		col := domain.ShortCodeToColour[r]
		result[i] = col
	}
	return &result, nil
}
