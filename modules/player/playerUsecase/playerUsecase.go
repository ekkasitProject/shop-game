package playerUsecase

import "github.com/ekkasitProject/shop-game/modules/player/playerRepository"

type (
	PlayerUsecaseService interface {
	}

	playerUsecase struct {
		playerRepository playerRepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(playerRepository playerRepository.PlayerRepositoryService) PlayerUsecaseService {
	return &playerUsecase{
		playerRepository: playerRepository,
	}
}
