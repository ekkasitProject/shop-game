package paymentHandler

import (
	"github.com/ekkasitProject/shop-game/config"
	"github.com/ekkasitProject/shop-game/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface {
	}

	PaymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &PaymentHttpHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
