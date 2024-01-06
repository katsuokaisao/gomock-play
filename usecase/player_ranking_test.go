package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/katsuokaisao/gomock-play/model"
	"github.com/katsuokaisao/gomock-play/repository"
	"github.com/katsuokaisao/gomock-play/repository/mock"
)

func Test_playerRanking_GetPlayerRanking(t *testing.T) {
	type fields struct {
		playerRepository func(ctrl *gomock.Controller) repository.PlayerRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []PlayerRankingItem
		wantErr bool
	}{
		{
			name: "success#normal",
			fields: fields{
				playerRepository: func(ctrl *gomock.Controller) repository.PlayerRepository {
					m := mock.NewMockPlayerRepository(ctrl)
					m.EXPECT().GetPlayerList(gomock.Any()).Return([]model.Player{
						{
							ID:      "1",
							Name:    "test1",
							Ranking: 1,
						},
						{
							ID:      "2",
							Name:    "test2",
							Ranking: 2,
						},
						{
							ID:      "3",
							Name:    "test3",
							Ranking: 3,
						},
					}, nil)
					return m
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: []PlayerRankingItem{
				{
					Ranking: []model.Player{
						{
							ID:      "1",
							Name:    "test1",
							Ranking: 1,
						},
						{
							ID:      "2",
							Name:    "test2",
							Ranking: 2,
						},
						{
							ID:      "3",
							Name:    "test3",
							Ranking: 3,
						},
					},
					Count: 3,
				},
			},
			wantErr: false,
		},
		{
			name: "success#normal2",
			fields: fields{
				playerRepository: func(ctrl *gomock.Controller) repository.PlayerRepository {
					m := mock.NewMockPlayerRepository(ctrl)
					m.EXPECT().GetPlayerList(gomock.Any()).Return([]model.Player{
						{
							ID:      "3",
							Name:    "test3",
							Ranking: 3,
						},
						{
							ID:      "1",
							Name:    "test1",
							Ranking: 1,
						},
						{
							ID:      "2",
							Name:    "test2",
							Ranking: 2,
						},
					}, nil)
					return m
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: []PlayerRankingItem{
				{
					Ranking: []model.Player{
						{
							ID:      "1",
							Name:    "test1",
							Ranking: 1,
						},
						{
							ID:      "2",
							Name:    "test2",
							Ranking: 2,
						},
						{
							ID:      "3",
							Name:    "test3",
							Ranking: 3,
						},
					},
					Count: 3,
				},
			},
			wantErr: false,
		},
		{
			name: "success#retry",
			fields: fields{
				playerRepository: func(ctrl *gomock.Controller) repository.PlayerRepository {
					m := mock.NewMockPlayerRepository(ctrl)
					gomock.InOrder(
						m.EXPECT().GetPlayerList(gomock.Any()).Return(nil, errors.New("timeout")),
						m.EXPECT().GetPlayerList(gomock.Any()).Return([]model.Player{
							{
								ID:      "3",
								Name:    "test3",
								Ranking: 3,
							},
							{
								ID:      "1",
								Name:    "test1",
								Ranking: 1,
							},
							{
								ID:      "2",
								Name:    "test2",
								Ranking: 2,
							},
						}, nil),
					)
					return m
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: []PlayerRankingItem{
				{
					Ranking: []model.Player{
						{
							ID:      "1",
							Name:    "test1",
							Ranking: 1,
						},
						{
							ID:      "2",
							Name:    "test2",
							Ranking: 2,
						},
						{
							ID:      "3",
							Name:    "test3",
							Ranking: 3,
						},
					},
					Count: 3,
				},
			},
			wantErr: false,
		},
		{
			name: "failed#retry",
			fields: fields{
				playerRepository: func(ctrl *gomock.Controller) repository.PlayerRepository {
					m := mock.NewMockPlayerRepository(ctrl)
					gomock.InOrder(
						m.EXPECT().GetPlayerList(gomock.Any()).Return(nil, errors.New("timeout")),
						m.EXPECT().GetPlayerList(gomock.Any()).Return(nil, errors.New("timeout")),
					)
					return m
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			// テストの内容に意味はない
			// Doは、mockの引数をカスタムロジックでテストすることができる
			name: "play#gomock.Do",
			fields: fields{
				playerRepository: func(ctrl *gomock.Controller) repository.PlayerRepository {
					m := mock.NewMockPlayerRepository(ctrl)
					gomock.InOrder(
						m.EXPECT().GetPlayerList(gomock.Any()).Do(
							func(ctx context.Context) {
								msg := ctx.Value("test").(string)
								if msg != "Hello World!" {
									t.Errorf("playerRanking.GetPlayerRanking() error occurred actual %v, want %v", msg, "Hello World!")
								}
							},
						).Return(
							[]model.Player{
								{
									ID:      "3",
									Name:    "test3",
									Ranking: 3,
								},
								{
									ID:      "1",
									Name:    "test1",
									Ranking: 1,
								},
								{
									ID:      "2",
									Name:    "test2",
									Ranking: 2,
								},
							},
							nil,
						),
					)
					return m
				},
			},
			args: args{
				ctx: context.WithValue(context.Background(), "test", "Hello World!"),
			},
			want: []PlayerRankingItem{
				{
					Ranking: []model.Player{
						{
							ID:      "1",
							Name:    "test1",
							Ranking: 1,
						},
						{
							ID:      "2",
							Name:    "test2",
							Ranking: 2,
						},
						{
							ID:      "3",
							Name:    "test3",
							Ranking: 3,
						},
					},
					Count: 3,
				},
			},
			wantErr: false,
		},
		{
			// テストの内容に意味はない
			// DoAndReturnは、mockの引数により返り値を変えることができる
			name: "play#gomock.DoAndReturn",
			fields: fields{
				playerRepository: func(ctrl *gomock.Controller) repository.PlayerRepository {
					m := mock.NewMockPlayerRepository(ctrl)
					gomock.InOrder(
						m.EXPECT().GetPlayerList(gomock.Any()).DoAndReturn(
							func(ctx context.Context) ([]model.Player, error) {
								msg := ctx.Value("test").(string)
								if msg != "Hello World!" {
									return nil, errors.New("error occurred")
								}

								return []model.Player{
									{
										ID:      "3",
										Name:    "test3",
										Ranking: 3,
									},
									{
										ID:      "1",
										Name:    "test1",
										Ranking: 1,
									},
									{
										ID:      "2",
										Name:    "test2",
										Ranking: 2,
									},
								}, nil
							},
						).AnyTimes(),
					)
					return m
				},
			},
			args: args{
				ctx: context.WithValue(context.Background(), "test", "Hello Japan!"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			p := &playerRanking{
				playerRepository: tt.fields.playerRepository(ctrl),
			}

			got, err := p.GetPlayerRanking(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("playerRanking.GetPlayerRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
