package service

import (
	"context"
	"errors"
	"os"

	watermark "github.com/elliottlin2020/Week04/internal/biz"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	AddDocumentEndpoint   endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	WatermarkEndpoint     endpoint.Endpoint
}

func NewEndpointSet(svc watermark.Service) Set {
	return Set{
		GetEndpoint:         MakeGetEndpoint(svc),
		AddDocumentEndpoint: MakeAddDocumentEndpoint(svc),
		WatermarkEndpoint:   MakeWatermarkEndpoint(svc),
	}
}

func MakeGetEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		docs, err := svc.Get(ctx, req.Filters...)
		if err != nil {
			return GetResponse{docs, err.Error()}, nil
		}
		return GetResponse{docs, ""}, nil
	}
}

func MakeAddDocumentEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddDocumentRequest)
		ticketID, err := svc.AddDocument(ctx, req.Document)
		if err != nil {
			return AddDocumentResponse{TicketID: ticketID, Err: err.Error()}, nil
		}
		return AddDocumentResponse{TicketID: ticketID, Err: ""}, nil
	}
}

func MakeWatermarkEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WatermarkRequest)
		code, err := svc.Watermark(ctx, req.TicketID, req.Mark)
		if err != nil {
			return WatermarkResponse{Code: code, Err: err.Error()}, nil
		}
		return WatermarkResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context, filters ...watermark.Filter) ([]watermark.Document, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})
	if err != nil {
		return []watermark.Document{}, err
	}
	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return []watermark.Document{}, errors.New(getResp.Err)
	}
	return getResp.Documents, nil
}

func (s *Set) AddDocument(ctx context.Context, doc *watermark.Document) (string, error) {
	resp, err := s.AddDocumentEndpoint(ctx, AddDocumentRequest{Document: doc})
	if err != nil {
		return "", err
	}
	adResp := resp.(AddDocumentResponse)
	if adResp.Err != "" {
		return "", errors.New(adResp.Err)
	}
	return adResp.TicketID, nil
}

func (s *Set) Watermark(ctx context.Context, ticketID, mark string) (int, error) {
	resp, err := s.WatermarkEndpoint(ctx, WatermarkRequest{TicketID: ticketID, Mark: mark})
	wmResp := resp.(WatermarkResponse)
	if err != nil {
		return wmResp.Code, err
	}
	if wmResp.Err != "" {
		return wmResp.Code, errors.New(wmResp.Err)
	}
	return wmResp.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
