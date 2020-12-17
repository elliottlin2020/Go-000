package service

import (
	watermark "github.com/elliottlin2020/Week04/internal/biz"
)

type GetRequest struct {
	Filters []watermark.Filter `json:"filters,omitempty"`
}

type GetResponse struct {
	Documents []watermark.Document `json:"documents"`
	Err       string               `json:"err,omitempty"`
}

type WatermarkRequest struct {
	TicketID string `json:"ticketID"`
	Mark     string `json:"mark"`
}

type WatermarkResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

type AddDocumentRequest struct {
	Document *watermark.Document `json:"document"`
}

type AddDocumentResponse struct {
	TicketID string `json:"ticketID"`
	Err      string `json:"err,omitempty"`
}
