package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/Seiji-Ikeda32/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transferReqest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,gt=0"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferReqest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}

	if !server.validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx.Request.Context(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (sever *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := sever.store.GetAccount(ctx.Request.Context(), accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponce(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return false
	}
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return false
	}

	return true
}
