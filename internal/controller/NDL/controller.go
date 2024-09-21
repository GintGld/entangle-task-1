package controller

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/big"
	"net/http"
	"task1/internal/lib/utils/sl"

	"github.com/gorilla/mux"
)

type controller struct {
	log    *slog.Logger
	nglSrv NGLService
}

type NGLService interface {
	NGLAmount(context.Context) (big.Int, error)
}

func New(
	log *slog.Logger,
	nglSrv NGLService,
) *mux.Router {
	ctr := controller{
		log:    log,
		nglSrv: nglSrv,
	}

	mux := mux.NewRouter()

	mux.HandleFunc("/getTotalSupply", ctr.getTotalSupply)

	return mux
}

// get TotalSupply returns total supply of NGL.
func (c *controller) getTotalSupply(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	num, err := c.nglSrv.NGLAmount(req.Context())
	if err != nil {
		c.log.Error("failed to get NGL amount", sl.Err(err))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	msg, _ := json.Marshal(struct {
		Amount string `json:"amount"`
	}{
		Amount: num.String(),
	})

	resp.WriteHeader(http.StatusOK)
	if _, err := resp.Write(msg); err != nil {
		c.log.Error("failed to send response", sl.Err(err))
	}
}
