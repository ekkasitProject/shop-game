package paymentHandler

import (
	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface {
	}

	PaymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &PaymentQueueHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
