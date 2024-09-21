package service

import (
	"context"
	"errors"
	"log/slog"
	"math/big"
	"os"
	"task1/internal/models"
	"task1/internal/service/NGL/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](t T) *T {
	return &t
}

func TestAmount(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type clientRes struct {
		resp models.NGLAmountResp
		err  error
	}
	type want struct {
		amount big.Int
		err    error
	}
	tests := []struct {
		name      string
		args      args
		clientRes clientRes
		want      want
	}{
		{
			name: "main line",
			args: args{ctx: context.Background()},
			clientRes: clientRes{
				resp: models.NGLAmountResp{Supply: []models.NGLSupply{
					{
						Amount: ptr(models.AmountType(*big.NewInt(10493))),
					},
				}},
				err: nil,
			},
			want: want{
				amount: *big.NewInt(10493),
				err:    nil,
			},
		},
		{
			name:      "error from client",
			args:      args{ctx: context.Background()},
			clientRes: clientRes{models.NGLAmountResp{}, errors.New("some http error")},
			want:      want{big.Int{}, errors.New("NGL.NGLAmount: some http error")},
		},
		{
			name:      "empty supply",
			args:      args{ctx: context.Background()},
			clientRes: clientRes{models.NGLAmountResp{Supply: []models.NGLSupply{}}, nil},
			want:      want{big.Int{}, errors.New("no supply")},
		},
		{
			name: "supply with length > 1",
			args: args{ctx: context.Background()},
			clientRes: clientRes{models.NGLAmountResp{Supply: []models.NGLSupply{
				{Amount: ptr(models.AmountType(*big.NewInt(516843526)))},
				{Amount: ptr(models.AmountType(*big.NewInt(238239)))},
			}}, nil},
			want: want{*big.NewInt(516843526), nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.NewNGLClient(t)

			client.
				On("Amount", tt.args.ctx).
				Return(tt.clientRes.resp, tt.clientRes.err)

			ngl := NGL{
				log: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
				c:   client,
			}

			res, err := ngl.NGLAmount(tt.args.ctx)
			assert.Equal(t, tt.want.amount, res)
			if tt.want.err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.want.err.Error())
			}
		})
	}
}
