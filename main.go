package main

import (
	"context"

	"github.com/katsuokaisao/gomock-play/model"
	"github.com/katsuokaisao/gomock-play/repository"
	"github.com/katsuokaisao/gomock-play/usecase"
)

func main() {
	u := usecase.NewPlayerRanking(NewEmpty())
	u.GetPlayerRanking(context.Background())
}

type empty struct{}

func NewEmpty() repository.PlayerRepository {
	return &empty{}
}

func (e *empty) GetPlayerList(ctx context.Context) ([]model.Player, error) {
	return nil, nil
}
