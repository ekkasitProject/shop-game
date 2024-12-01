package playerHandler

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/player"
	"github.com/ekkasitProject/shop-game/modules/player/playerUsecase"
	"github.com/ekkasitProject/shop-game/pkg/request"
	"github.com/ekkasitProject/shop-game/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	PlayerHttpHandlerService interface {
		CreatePlayer(c echo.Context) error
		FindOnePlayerProfile(c echo.Context) error
		AddPlayerMoney(c echo.Context) error
		GetPlayerSavingAccount(c echo.Context) error
	}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}

func (h *playerHttpHandler) CreatePlayer(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(player.CreatePlayerReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	res, err := h.playerUsecase.CreatePlayer(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	return response.SuccessResponse(c, http.StatusAccepted, res)
}

func (h *playerHttpHandler) FindOnePlayerProfile(c echo.Context) error {
	ctx := context.Background()
	playerId := strings.TrimPrefix(c.Param("player_id"), "player:")
	res, err := h.playerUsecase.FindOnePlayerProfile(ctx, playerId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	return response.SuccessResponse(c, http.StatusAccepted, res)

}
func (h *playerHttpHandler) AddPlayerMoney(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	playerId := strings.TrimPrefix(c.Param("player_id"), "player:")
	req := new(player.CreatePlayerTransactionReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	req.PlayerId = playerId
	res, err := h.playerUsecase.AddPlayerMoney(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusAccepted, res)
}

func (h *playerHttpHandler) GetPlayerSavingAccount(c echo.Context) error {
	ctx := context.Background()

	playerId := strings.TrimPrefix(c.Param("player_id"), "player:")
	log.Printf("Debug: playerId = %s", playerId)

	res, err := h.playerUsecase.GetPlayerSavingAccount(ctx, playerId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusAccepted, res)
}
