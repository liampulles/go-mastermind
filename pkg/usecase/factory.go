package usecase

import (
	"fmt"

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
	var result domain.Combination
	for i, elem := range request {
		col, err := asColour(elem)
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

func asColour(in string) (domain.Colour, error) {
	r, err := domain.StringToRune(in)
	if err != nil {
		return domain.Colour(-1), fmt.Errorf("cannot make %s a rune: %w", in, err)
	}

	col, ok := domain.ShortCodeToColour[r]
	if !ok {
		return domain.Colour(-1), fmt.Errorf("no colour corresponds to %s", in)
	}
	return col, nil
}
