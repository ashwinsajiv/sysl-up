// Code generated by sysl DO NOT EDIT.
package up

import (
	"context"
	"net/http"

	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/handlerinitialiser"

	"github.com/go-chi/chi"
)

// Router interface for Up
type Router interface {
	Route(router *chi.Mux)
}

// ServiceRouter for Up API
type ServiceRouter struct {
	gc               core.RestGenCallback
	svcHandler       *ServiceHandler
	basePathFromSpec string
}

// swaggerFile is a struct to store the swagger file system
type swaggerFile struct {
	file http.FileSystem
}

// swagger will receive the embedded swagger file if it is generated by the resource application
var swagger = swaggerFile{}

// NewServiceRouter creates a new service router for Up
func NewServiceRouter(gc core.RestGenCallback, svcHandler *ServiceHandler) handlerinitialiser.HandlerInitialiser {
	return &ServiceRouter{gc, svcHandler, ""}
}

// WireRoutes ...
//nolint:funlen
func (s *ServiceRouter) WireRoutes(ctx context.Context, r chi.Router) {
	r.Route(core.SelectBasePath(s.basePathFromSpec, s.gc.BasePath()), func(r chi.Router) {
		s.gc.AddMiddleware(ctx, r)
		r.Delete("/transactions/{transactionId}/relationships/tags", s.svcHandler.DeleteTransactionsRelationshipsTagsHandler)
		r.Delete("/webhooks/{id}", s.svcHandler.DeleteWebhooksHandler)
		r.Get("/accounts", s.svcHandler.GetAccountsListHandler)
		r.Get("/accounts/{accountId}/transactions", s.svcHandler.GetAccountsTransactionsListHandler)
		r.Get("/accounts/{id}", s.svcHandler.GetAccountsHandler)
		r.Get("/categories", s.svcHandler.GetCategoriesListHandler)
		r.Get("/categories/{id}", s.svcHandler.GetCategoriesHandler)
		r.Get("/tags", s.svcHandler.GetTagsListHandler)
		r.Get("/transactions", s.svcHandler.GetTransactionsListHandler)
		r.Get("/transactions/{id}", s.svcHandler.GetTransactionsHandler)
		r.Get("/util/ping", s.svcHandler.GetUtilPingListHandler)
		r.Get("/v1/accounts", s.svcHandler.GetV1AccountsListHandler)
		r.Get("/webhooks", s.svcHandler.GetWebhooksListHandler)
		r.Get("/webhooks/{id}", s.svcHandler.GetWebhooksHandler)
		r.Get("/webhooks/{webhookId}/logs", s.svcHandler.GetWebhooksLogsListHandler)
		r.Post("/transactions/{transactionId}/relationships/tags", s.svcHandler.PostTransactionsRelationshipsTagsHandler)
		r.Post("/webhooks", s.svcHandler.PostWebhooksHandler)
		r.Post("/webhooks/{webhookId}/ping", s.svcHandler.PostWebhooksPingHandler)
	})
}

// Config ...
func (s *ServiceRouter) Config() interface{} {
	return s.gc.Config()
}

// Name ...
func (s *ServiceRouter) Name() string {
	return "Up"
}