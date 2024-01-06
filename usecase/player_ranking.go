package usecase

import (
	"context"
	"errors"
	"sort"

	"github.com/katsuokaisao/gomock-play/model"
	"github.com/katsuokaisao/gomock-play/repository"
)

type PlayerRanking interface {
	GetPlayerRanking(ctx context.Context) ([]PlayerRankingItem, error)
}

type playerRanking struct {
	playerRepository repository.PlayerRepository
}

func NewPlayerRanking(playerRepository repository.PlayerRepository) PlayerRanking {
	return &playerRanking{
		playerRepository: playerRepository,
	}
}

type PlayerRankingItem struct {
	Ranking []model.Player
	Count   int
}

func (p *playerRanking) GetPlayerRanking(ctx context.Context) ([]PlayerRankingItem, error) {
	players, err := p.playerRepository.GetPlayerList(ctx)
	if err != nil {
		// retry
		players, err = p.playerRepository.GetPlayerList(ctx)
		if err != nil {
			return nil, err
		}
	}

	if len(players) == 0 {
		return nil, errors.New("player not found")
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Ranking < players[j].Ranking
	})

	return []PlayerRankingItem{
		{
			Ranking: players,
			Count:   len(players),
		},
	}, nil
}
