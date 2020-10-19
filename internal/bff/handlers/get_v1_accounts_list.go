package handlers

import (
	"context"
	"fmt"
	"github.com/anz-bank/pkg/log"
	"github.com/ashwinsajiv/sysl-up/gen/pkg/servers/bff"
	"github.com/ashwinsajiv/sysl-up/gen/pkg/servers/bff/up"
)

type GetV1AccountsList struct {
}

func NewGetV1AccountsList() GetV1AccountsList {
	return GetV1AccountsList{}
}

func (h GetV1AccountsList) GetV1AccountsList(ctx context.Context,
	_ *bff.GetV1AccountsListRequest,
	client bff.GetV1AccountsListClient) (*bff.GetAccountResponse, error) {
	setJSONResponseContentType(ctx)
	request := up.GetAccountsListRequest{PageSize: Int64(1)}
	response, err := client.UpGetAccountsList(ctx, &request)
	if err != nil {
		log.Error(ctx, fmt.Errorf("unsuccessful response: %w", err))
		return nil, bff.InternalServerError
	}
	if response == nil {
		log.Error(ctx, fmt.Errorf("response contains no result"))
		return nil, bff.InternalServerError
	}
	balance := response.Data[0].Attributes.Balance.Value
	return &bff.GetAccountResponse{Amount: balance}, nil
}
