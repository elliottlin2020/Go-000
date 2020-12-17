package biz

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, filters ...Filter) ([]Document, error)
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	AddDocument(ctx context.Context, doc *Document) (string, error)
}
