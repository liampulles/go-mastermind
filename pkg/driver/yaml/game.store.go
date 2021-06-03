package yaml

import (
	"fmt"
	"os"
	"path"

	"github.com/liampulles/go-mastermind/pkg/domain"
	"github.com/liampulles/go-mastermind/pkg/usecase"
	"gopkg.in/yaml.v2"
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

func (gs *GameStore) New(secret *domain.Combination) error {
	game := GameYAML{
		Secret: asStringCombination(secret),
		Won:    false,
	}
	if err := writeGame(&game); err != nil {
		return err
	}
	return nil
}

func (gs *GameStore) IsWon() (bool, error) {
	game, err := readGame()
	if err != nil {
		return false, err
	}

	return game.Won, nil
}

func (gs *GameStore) GetSecret() (*domain.Combination, error) {
	game, err := readGame()
	if err != nil {
		return nil, err
	}

	comb, err := asDomainCombination(&game.Secret)
	if err != nil {
		return nil, fmt.Errorf("could not create combination: %w", err)
	}
	return comb, nil
}

func (gs *GameStore) EndGame() error {
	game, err := readGame()
	if err != nil {
		return err
	}
	game.Won = true
	if err := writeGame(game); err != nil {
		return err
	}
	return nil
}

func readGame() (*GameYAML, error) {
	gameFile, err := fileLocation()
	if err != nil {
		return nil, err
	}
	bytes, err := os.ReadFile(gameFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("no game exists for id %s", gameFile)
		}
		return nil, fmt.Errorf("could not read game file: %w", err)
	}

	game := GameYAML{}
	if err := yaml.Unmarshal(bytes, &game); err != nil {
		return nil, fmt.Errorf("could not parse game file YAML: %w", err)
	}
	return &game, nil
}

func writeGame(game *GameYAML) error {
	gameFile, err := fileLocation()
	if err != nil {
		return err
	}
	bytes, err := yaml.Marshal(&game)
	if err != nil {
		return fmt.Errorf("could not marshal game to YAML: %w", err)
	}
	if err := os.WriteFile(gameFile, bytes, 0777); err != nil {
		return fmt.Errorf("could not write YAML to file: %w", err)
	}
	return nil
}

func fileLocation() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("could not identify config dir: %w", err)
	}
	return path.Join(cfgDir, "gomastermind.game"), nil
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
