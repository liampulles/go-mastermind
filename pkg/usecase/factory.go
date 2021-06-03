package usecase

import (
	"fmt"
	"strings"

	"github.com/liampulles/go-mastermind/pkg/domain"
)

type Factory interface {
	CreateCombination(*Request) (*domain.Combination, error)
	CreateResponse(eval *domain.Evaluation, gameWon bool) *Response
}

type FactoryImpl struct {
	engine domain.Engine
}

var _ Factory = &FactoryImpl{}

func NewFactoryImpl(engine domain.Engine) *FactoryImpl {
	return &FactoryImpl{
		engine: engine,
	}
}

func (f *FactoryImpl) CreateCombination(request *Request) (*domain.Combination, error) {
	str := strings.ToUpper(string(*request))
	if len(str) != 4 {
		return nil, fmt.Errorf("guess must be exactly 4 letters")
	}

	var result domain.Combination
	for i, r := range str {
		col, err := asColour(r)
		if err != nil {
			return nil, fmt.Errorf("could not map element %d of request: %w", i, err)
		}
		result[i] = col
	}
	return &result, nil
}

func (f *FactoryImpl) CreateResponse(eval *domain.Evaluation, gameWon bool) *Response {
	return &Response{
		Evaluation: eval,
		GameWon:    gameWon,
	}
}

func asColour(in rune) (domain.Colour, error) {
	col, ok := domain.ShortCodeToColour[in]
	if !ok {
		return domain.Colour(-1), fmt.Errorf("no colour corresponds to %s", string(in))
	}
	return col, nil
}
