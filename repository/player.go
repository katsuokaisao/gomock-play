package repository

import (
	"context"

	"github.com/katsuokaisao/gomock-play/model"
)

//go:generate mockgen -source=$GOFILE -destination=mock/${GOFILE} -package=mock
type PlayerRepository interface {
	GetPlayerList(context.Context) ([]model.Player, error)
}
