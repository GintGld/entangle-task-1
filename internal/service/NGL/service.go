package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"

	"task1/internal/lib/utils/sl"
	"task1/internal/models"
)

type NGL struct {
	log *slog.Logger
	c   NGLClient
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name NGLClient
type NGLClient interface {
	Amount(ctx context.Context) (models.NGLAmountResp, error)
}

func New(
	log *slog.Logger,
	client NGLClient,
) *NGL {
	return &NGL{
		log: log,
		c:   client,
	}
}

// NGLAmount returns currect NGL token amount in big.Int.
func (n *NGL) NGLAmount(ctx context.Context) (big.Int, error) {
	const op = "NGL.NGLAmount"

	log := n.log.With(
		slog.String("op", op),
	)

	res, err := n.c.Amount(ctx)
	if err != nil {
		log.Error("failed to process Amount request", sl.Err(err))
		return big.Int{}, fmt.Errorf("%s: %w", op, err)
	}

	if len(res.Supply) == 0 {
		log.Error("supply slice is empty")
		return big.Int{}, errors.New("no supply")
	}

	if len(res.Supply) > 1 {
		log.Warn("supply slice has more tha one element, take only first")
	}

	return res.Supply[0].Amount.BigInt(), nil
}
