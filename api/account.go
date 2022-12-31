package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/Seiji-Ikeda32/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountReqest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR JPY"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountReqest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx.Request.Context(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id", binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	account, err := server.store.GetAccount(ctx.Request.Context(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponce(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id", binding:"required,min=1"`
	PageSize int32 `form:"page_size", binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	fmt.Println(arg.Limit)
	fmt.Println(arg.Offset)

	accounts, err := server.store.ListAccounts(ctx.Request.Context(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
