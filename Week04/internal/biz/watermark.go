package biz

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/lithammer/shortuuid/v3"
)

type watermarkService struct{}

func NewService() Service { return &watermarkService{} }

func (w *watermarkService) Get(_ context.Context, filters ...Filter) ([]Document, error) {
	// query the database using the filters and return the list of documents
	doc := Document{
		Content: "book",
		Title:   "Harry Potter and Half Blood Prince",
		Author:  "J.K. Rowling",
		Topic:   "Fiction and Magic",
	}
	return []Document{doc}, nil
}

func (w *watermarkService) Watermark(_ context.Context, ticketID, mark string) (int, error) {
	// update the database entry with watermark field as non empty
	// first check if the watermark status is not already in InProgress, Started or Finished state
	// If yes, then return invalid request
	// return error if no item found using the ticketID
	return http.StatusOK, nil
}

func (w *watermarkService) AddDocument(_ context.Context, doc *Document) (string, error) {
	// add the document entry in the database
	// return error if the doc is invalid and/or the database invalid entry error
	newTicketID := shortuuid.New()
	return newTicketID, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
