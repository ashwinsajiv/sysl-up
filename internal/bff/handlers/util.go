package handlers

import (
	"context"

	"github.com/anz-bank/sysl-go/common"
)

// setJSONResponseContentType sets the content type of the response to the appropriate application/json value.
func setJSONResponseContentType(ctx context.Context) {
	headers, _ := common.RespHeaderAndStatusFromContext(ctx)
	headers["Content-Type"] = []string{"application/json; charset=utf-8"}
}

func Int64(i int64) *int64 {
	return &i
}
