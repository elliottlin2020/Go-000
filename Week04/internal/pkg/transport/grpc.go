package transport

import (
	"context"

	pb "github.com/elliottlin2020/Week04/api"
	watermark "github.com/elliottlin2020/Week04/internal/biz"
	svc "github.com/elliottlin2020/Week04/internal/service"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	get           grpctransport.Handler
	status        grpctransport.Handler
	addDocument   grpctransport.Handler
	watermark     grpctransport.Handler
	serviceStatus grpctransport.Handler
}

func NewGRPCServer(ep svc.Set) pb.WatermarkServer {
	return &grpcServer{
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGRPCGetRequest,
			decodeGRPCGetResponse,
		),
		addDocument: grpctransport.NewServer(
			ep.AddDocumentEndpoint,
			decodeGRPCAddDocumentRequest,
			decodeGRPCAddDocumentResponse,
		),
		watermark: grpctransport.NewServer(
			ep.WatermarkEndpoint,
			decodeGRPCWatermarkRequest,
			decodeGRPCWatermarkResponse,
		),
	}
}

func (g *grpcServer) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetReply), nil
}

func (g *grpcServer) AddDocument(ctx context.Context, r *pb.AddDocumentRequest) (*pb.AddDocumentReply, error) {
	_, rep, err := g.addDocument.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddDocumentReply), nil
}

func (g *grpcServer) Watermark(ctx context.Context, r *pb.WatermarkRequest) (*pb.WatermarkReply, error) {
	_, rep, err := g.watermark.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WatermarkReply), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetRequest)
	var filters []watermark.Filter
	for _, f := range req.Filters {
		filters = append(filters, watermark.Filter{Key: f.Key, Value: f.Value})
	}
	return svc.GetRequest{Filters: filters}, nil
}

func decodeGRPCWatermarkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.WatermarkRequest)
	return svc.WatermarkRequest{TicketID: req.TicketID, Mark: req.Mark}, nil
}

func decodeGRPCAddDocumentRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddDocumentRequest)
	doc := &watermark.Document{
		Content:   req.Document.Content,
		Title:     req.Document.Title,
		Author:    req.Document.Author,
		Topic:     req.Document.Topic,
		Watermark: req.Document.Watermark,
	}
	return svc.AddDocumentRequest{Document: doc}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetReply)
	var docs []watermark.Document
	for _, d := range reply.Documents {
		doc := watermark.Document{
			Content:   d.Content,
			Title:     d.Title,
			Author:    d.Author,
			Topic:     d.Topic,
			Watermark: d.Watermark,
		}
		docs = append(docs, doc)
	}
	return svc.GetResponse{Documents: docs, Err: reply.Err}, nil
}

func decodeGRPCWatermarkResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.WatermarkReply)
	return svc.WatermarkResponse{Code: int(reply.Code), Err: reply.Err}, nil
}

func decodeGRPCAddDocumentResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.AddDocumentReply)
	return svc.AddDocumentResponse{TicketID: reply.TicketID, Err: reply.Err}, nil
}
