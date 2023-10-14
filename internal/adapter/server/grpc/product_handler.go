package grpcserver

import (
	"context"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/payment"
	"github.com/rs/zerolog/log"
)

// ListProduct is sending stream of product
func (a *GrpcServerAdapter) ListProduct(req *payment.ListProductRequest, stream payment.PaymentService_ListProductServer) error {
	products, err := a.paymentService.GetAllProducts(context.Background())
	if err != nil {
		return err
	}

	context := stream.Context()

	for _, product := range products {
		select {
		case <-context.Done():
			log.Info().Msg("client cancelled request")
			return nil
		default:
			stream.Send(&payment.Product{
				Code: product.Code,
				Name: product.Name,
			})
		}
	}

	return nil
}
