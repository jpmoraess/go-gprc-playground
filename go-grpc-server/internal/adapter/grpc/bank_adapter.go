package grpc

import (
	"context"
	"github.com/jpmoraess/grpc-proto/protogen/go/bank"
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

func (a *GrpcAdapter) GetCurrentBalance(ctx context.Context, req *bank.CurrentBalanceRequest) (*bank.CurrentBalanceResponse, error) {
	now := time.Now()
	balance := a.bankService.GetCurrentBalance(req.AccountNumber)
	return &bank.CurrentBalanceResponse{
		Amount: balance,
		CurrentDate: &date.Date{
			Year:  int32(now.Year()),
			Month: int32(now.Month()),
			Day:   int32(now.Day()),
		},
	}, nil
}
