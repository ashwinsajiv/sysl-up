// Code generated by sysl DO NOT EDIT.
package up

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Service interface for Up
type Service interface {
	DeleteTransactionsRelationshipsTags(ctx context.Context, req *DeleteTransactionsRelationshipsTagsRequest) (*http.Header, error)
	DeleteWebhooks(ctx context.Context, req *DeleteWebhooksRequest) (*http.Header, error)
	GetAccountsList(ctx context.Context, req *GetAccountsListRequest) (*ListAccountsResponse, error)
	GetAccountsTransactionsList(ctx context.Context, req *GetAccountsTransactionsListRequest) (*ListTransactionsResponse, error)
	GetAccounts(ctx context.Context, req *GetAccountsRequest) (*GetAccountResponse, error)
	GetCategoriesList(ctx context.Context, req *GetCategoriesListRequest) (*ListCategoriesResponse, error)
	GetCategories(ctx context.Context, req *GetCategoriesRequest) (*GetCategoryResponse, error)
	GetTagsList(ctx context.Context, req *GetTagsListRequest) (*ListTagsResponse, error)
	GetTransactionsList(ctx context.Context, req *GetTransactionsListRequest) (*ListTransactionsResponse, error)
	GetTransactions(ctx context.Context, req *GetTransactionsRequest) (*GetTransactionResponse, error)
	GetUtilPingList(ctx context.Context, req *GetUtilPingListRequest) (*PingResponse, error)
	GetWebhooksList(ctx context.Context, req *GetWebhooksListRequest) (*ListWebhooksResponse, error)
	GetWebhooks(ctx context.Context, req *GetWebhooksRequest) (*GetWebhookResponse, error)
	GetWebhooksLogsList(ctx context.Context, req *GetWebhooksLogsListRequest) (*ListWebhookDeliveryLogsResponse, error)
	PostTransactionsRelationshipsTags(ctx context.Context, req *PostTransactionsRelationshipsTagsRequest) (*http.Header, error)
	PostWebhooks(ctx context.Context, req *PostWebhooksRequest) (*CreateWebhookResponse, error)
	PostWebhooksPing(ctx context.Context, req *PostWebhooksPingRequest) (*WebhookEventCallback, error)
}

// Client for Up API
type Client struct {
	Client  *http.Client
	URL     string
	Headers map[string][]string
}

// NewClient for Up
func NewClient(client *http.Client, serviceURL string) *Client {
	return &Client{client, serviceURL, nil}
}

// DeleteTransactionsRelationshipsTags ...
func (s *Client) DeleteTransactionsRelationshipsTags(ctx context.Context, req *DeleteTransactionsRelationshipsTagsRequest) (*http.Header, error) {
	required := []string{}
	u, err := url.Parse(fmt.Sprintf("%s/transactions/%v/relationships/tags", s.URL, req.TransactionId))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "DELETE",
		URLString:     u.String(),
		Body:          req.Request,
		Required:      required,
		OKResponse:    nil,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- DELETE "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	return &result.HTTPResponse.Header, nil
}

// DeleteWebhooks ...
func (s *Client) DeleteWebhooks(ctx context.Context, req *DeleteWebhooksRequest) (*http.Header, error) {
	required := []string{}
	u, err := url.Parse(fmt.Sprintf("%s/webhooks/%v", s.URL, req.ID))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "DELETE",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    nil,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- DELETE "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	return &result.HTTPResponse.Header, nil
}

// GetAccountsList ...
func (s *Client) GetAccountsList(ctx context.Context, req *GetAccountsListRequest) (*ListAccountsResponse, error) {
	required := []string{}
	var okResponse ListAccountsResponse
	u, err := url.Parse(fmt.Sprintf("%s/accounts", s.URL))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	q := u.Query()
	if req.PageSize != nil {
		q.Add("page[size]", fmt.Sprintf("%v", *req.PageSize))
	}
	u.RawQuery = q.Encode()
	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkListAccountsResponseResponse, ok := result.Response.(*ListAccountsResponse)
	if ok {
		valErr := validator.Validate(OkListAccountsResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkListAccountsResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetAccountsTransactionsList ...
func (s *Client) GetAccountsTransactionsList(ctx context.Context, req *GetAccountsTransactionsListRequest) (*ListTransactionsResponse, error) {
	required := []string{}
	var okResponse ListTransactionsResponse
	u, err := url.Parse(fmt.Sprintf("%s/accounts/%v/transactions", s.URL, req.AccountId))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	q := u.Query()
	if req.PageSize != nil {
		q.Add("page[size]", fmt.Sprintf("%v", *req.PageSize))
	}
	if req.FilterStatus != nil {
		q.Add("filter[status]", fmt.Sprintf("%v", *req.FilterStatus))
	}
	if req.FilterSince != nil {
		q.Add("filter[since]", *req.FilterSince)
	}
	if req.FilterUntil != nil {
		q.Add("filter[until]", *req.FilterUntil)
	}
	if req.FilterCategory != nil {
		q.Add("filter[category]", *req.FilterCategory)
	}
	if req.FilterTag != nil {
		q.Add("filter[tag]", *req.FilterTag)
	}
	u.RawQuery = q.Encode()
	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkListTransactionsResponseResponse, ok := result.Response.(*ListTransactionsResponse)
	if ok {
		valErr := validator.Validate(OkListTransactionsResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkListTransactionsResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetAccounts ...
func (s *Client) GetAccounts(ctx context.Context, req *GetAccountsRequest) (*GetAccountResponse, error) {
	required := []string{}
	var okResponse GetAccountResponse
	u, err := url.Parse(fmt.Sprintf("%s/accounts/%v", s.URL, req.ID))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkGetAccountResponseResponse, ok := result.Response.(*GetAccountResponse)
	if ok {
		valErr := validator.Validate(OkGetAccountResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkGetAccountResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetCategoriesList ...
func (s *Client) GetCategoriesList(ctx context.Context, req *GetCategoriesListRequest) (*ListCategoriesResponse, error) {
	required := []string{}
	var okResponse ListCategoriesResponse
	u, err := url.Parse(fmt.Sprintf("%s/categories", s.URL))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	q := u.Query()
	if req.FilterParent != nil {
		q.Add("filter[parent]", *req.FilterParent)
	}
	u.RawQuery = q.Encode()
	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkListCategoriesResponseResponse, ok := result.Response.(*ListCategoriesResponse)
	if ok {
		valErr := validator.Validate(OkListCategoriesResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkListCategoriesResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetCategories ...
func (s *Client) GetCategories(ctx context.Context, req *GetCategoriesRequest) (*GetCategoryResponse, error) {
	required := []string{}
	var okResponse GetCategoryResponse
	u, err := url.Parse(fmt.Sprintf("%s/categories/%v", s.URL, req.ID))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkGetCategoryResponseResponse, ok := result.Response.(*GetCategoryResponse)
	if ok {
		valErr := validator.Validate(OkGetCategoryResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkGetCategoryResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetTagsList ...
func (s *Client) GetTagsList(ctx context.Context, req *GetTagsListRequest) (*ListTagsResponse, error) {
	required := []string{}
	var okResponse ListTagsResponse
	u, err := url.Parse(fmt.Sprintf("%s/tags", s.URL))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	q := u.Query()
	if req.PageSize != nil {
		q.Add("page[size]", fmt.Sprintf("%v", *req.PageSize))
	}
	u.RawQuery = q.Encode()
	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkListTagsResponseResponse, ok := result.Response.(*ListTagsResponse)
	if ok {
		valErr := validator.Validate(OkListTagsResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkListTagsResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetTransactionsList ...
func (s *Client) GetTransactionsList(ctx context.Context, req *GetTransactionsListRequest) (*ListTransactionsResponse, error) {
	required := []string{}
	var okResponse ListTransactionsResponse
	u, err := url.Parse(fmt.Sprintf("%s/transactions", s.URL))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	q := u.Query()
	if req.PageSize != nil {
		q.Add("page[size]", fmt.Sprintf("%v", *req.PageSize))
	}
	if req.FilterStatus != nil {
		q.Add("filter[status]", fmt.Sprintf("%v", *req.FilterStatus))
	}
	if req.FilterSince != nil {
		q.Add("filter[since]", *req.FilterSince)
	}
	if req.FilterUntil != nil {
		q.Add("filter[until]", *req.FilterUntil)
	}
	if req.FilterCategory != nil {
		q.Add("filter[category]", *req.FilterCategory)
	}
	if req.FilterTag != nil {
		q.Add("filter[tag]", *req.FilterTag)
	}
	u.RawQuery = q.Encode()
	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkListTransactionsResponseResponse, ok := result.Response.(*ListTransactionsResponse)
	if ok {
		valErr := validator.Validate(OkListTransactionsResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkListTransactionsResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetTransactions ...
func (s *Client) GetTransactions(ctx context.Context, req *GetTransactionsRequest) (*GetTransactionResponse, error) {
	required := []string{}
	var okResponse GetTransactionResponse
	u, err := url.Parse(fmt.Sprintf("%s/transactions/%v", s.URL, req.ID))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkGetTransactionResponseResponse, ok := result.Response.(*GetTransactionResponse)
	if ok {
		valErr := validator.Validate(OkGetTransactionResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkGetTransactionResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetUtilPingList ...
func (s *Client) GetUtilPingList(ctx context.Context, req *GetUtilPingListRequest) (*PingResponse, error) {
	required := []string{}
	var okResponse PingResponse
	var errorResponse ErrorResponse
	u, err := url.Parse(fmt.Sprintf("%s/util/ping", s.URL))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: &errorResponse,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		response, ok := err.(*restlib.HTTPResult)
		if !ok {
			return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
		}
		return nil, common.CreateDownstreamError(ctx, common.DownstreamResponseError, response.HTTPResponse, response.Body, &errorResponse)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkPingResponseResponse, ok := result.Response.(*PingResponse)
	if ok {
		valErr := validator.Validate(OkPingResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkPingResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetWebhooksList ...
func (s *Client) GetWebhooksList(ctx context.Context, req *GetWebhooksListRequest) (*ListWebhooksResponse, error) {
	required := []string{}
	var okResponse ListWebhooksResponse
	u, err := url.Parse(fmt.Sprintf("%s/webhooks", s.URL))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	q := u.Query()
	if req.PageSize != nil {
		q.Add("page[size]", fmt.Sprintf("%v", *req.PageSize))
	}
	u.RawQuery = q.Encode()
	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkListWebhooksResponseResponse, ok := result.Response.(*ListWebhooksResponse)
	if ok {
		valErr := validator.Validate(OkListWebhooksResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkListWebhooksResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetWebhooks ...
func (s *Client) GetWebhooks(ctx context.Context, req *GetWebhooksRequest) (*GetWebhookResponse, error) {
	required := []string{}
	var okResponse GetWebhookResponse
	u, err := url.Parse(fmt.Sprintf("%s/webhooks/%v", s.URL, req.ID))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkGetWebhookResponseResponse, ok := result.Response.(*GetWebhookResponse)
	if ok {
		valErr := validator.Validate(OkGetWebhookResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkGetWebhookResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// GetWebhooksLogsList ...
func (s *Client) GetWebhooksLogsList(ctx context.Context, req *GetWebhooksLogsListRequest) (*ListWebhookDeliveryLogsResponse, error) {
	required := []string{}
	var okResponse ListWebhookDeliveryLogsResponse
	u, err := url.Parse(fmt.Sprintf("%s/webhooks/%v/logs", s.URL, req.WebhookId))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	q := u.Query()
	if req.PageSize != nil {
		q.Add("page[size]", fmt.Sprintf("%v", *req.PageSize))
	}
	u.RawQuery = q.Encode()
	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkListWebhookDeliveryLogsResponseResponse, ok := result.Response.(*ListWebhookDeliveryLogsResponse)
	if ok {
		valErr := validator.Validate(OkListWebhookDeliveryLogsResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkListWebhookDeliveryLogsResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// PostTransactionsRelationshipsTags ...
func (s *Client) PostTransactionsRelationshipsTags(ctx context.Context, req *PostTransactionsRelationshipsTagsRequest) (*http.Header, error) {
	required := []string{}
	u, err := url.Parse(fmt.Sprintf("%s/transactions/%v/relationships/tags", s.URL, req.TransactionId))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "POST",
		URLString:     u.String(),
		Body:          req.Request,
		Required:      required,
		OKResponse:    nil,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- POST "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	return &result.HTTPResponse.Header, nil
}

// PostWebhooks ...
func (s *Client) PostWebhooks(ctx context.Context, req *PostWebhooksRequest) (*CreateWebhookResponse, error) {
	required := []string{}
	var okResponse CreateWebhookResponse
	u, err := url.Parse(fmt.Sprintf("%s/webhooks", s.URL))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "POST",
		URLString:     u.String(),
		Body:          req.Request,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- POST "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkCreateWebhookResponseResponse, ok := result.Response.(*CreateWebhookResponse)
	if ok {
		valErr := validator.Validate(OkCreateWebhookResponseResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkCreateWebhookResponseResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}

// PostWebhooksPing ...
func (s *Client) PostWebhooksPing(ctx context.Context, req *PostWebhooksPingRequest) (*WebhookEventCallback, error) {
	required := []string{}
	var okResponse WebhookEventCallback
	u, err := url.Parse(fmt.Sprintf("%s/webhooks/%v/ping", s.URL, req.WebhookId))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "POST",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Up <- POST "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkWebhookEventCallbackResponse, ok := result.Response.(*WebhookEventCallback)
	if ok {
		valErr := validator.Validate(OkWebhookEventCallbackResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkWebhookEventCallbackResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}
