package server

import (
	"github.com/ekkasitProject/shop-game/modules/payment/paymentHandler"
	"github.com/ekkasitProject/shop-game/modules/payment/paymentRepository"
	"github.com/ekkasitProject/shop-game/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	paymentRepository := paymentRepository.NewPaymentRepository(s.db)
	paymentUsecase := paymentUsecase.NewPaymentUsecase(paymentRepository)
	paymentHttpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, paymentUsecase)
	paymentQueueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, paymentUsecase)

	_ = paymentHttpHandler
	_ = paymentQueueHandler

	payment := s.app.Group("/payment_v1")

	// Health check
	payment.GET("/", s.healthCheckService)
}
