package internal

import (
	"sync"

	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/naveego/plugin-nifi/internal/pub"
)

type Server struct {
	mu  *sync.Mutex
	log hclog.Logger
}

// NewServer creates a new publisher Server.
func NewServer(log hclog.Logger) pub.PublisherServer {
	return &Server{
		mu:  &sync.Mutex{},
		log: log,
	}
}

func (s *Server) ConnectSession(*pub.ConnectRequest, pub.Publisher_ConnectSessionServer) error {
	panic("implement me")
}

func (s *Server) ConfigureConnection(context.Context, *pub.ConfigureConnectionRequest) (*pub.ConfigureConnectionResponse, error) {
	panic("implement me")
}

func (s *Server) ConfigureQuery(context.Context, *pub.ConfigureQueryRequest) (*pub.ConfigureQueryResponse, error) {
	panic("implement me")
}

func (s *Server) ConfigureRealTime(ctx context.Context, req *pub.ConfigureRealTimeRequest) (*pub.ConfigureRealTimeResponse, error) {
	if req.Form == nil {
		req.Form = &pub.ConfigurationFormRequest{}
	}

	resp := &pub.ConfigureRealTimeResponse{
		Form: &pub.ConfigurationFormResponse{
			DataJson:   req.Form.DataJson,
			SchemaJson: settingsSchema,
		},
	}

	return resp, nil
}

func (s *Server) BeginOAuthFlow(context.Context, *pub.BeginOAuthFlowRequest) (*pub.BeginOAuthFlowResponse, error) {
	panic("implement me")
}

func (s *Server) CompleteOAuthFlow(context.Context, *pub.CompleteOAuthFlowRequest) (*pub.CompleteOAuthFlowResponse, error) {
	panic("implement me")
}

func (s *Server) Connect(ctx context.Context, req *pub.ConnectRequest) (*pub.ConnectResponse, error) {
	return new(pub.ConnectResponse), nil
}

func (s *Server) DiscoverShapes(ctx context.Context, req *pub.DiscoverShapesRequest) (*pub.DiscoverShapesResponse, error) {
	return &pub.DiscoverShapesResponse{}, nil
}

func (s *Server) PublishStream(req *pub.PublishRequest, stream pub.Publisher_PublishStreamServer) error {
	return nil
}

func (s *Server) Disconnect(context.Context, *pub.DisconnectRequest) (*pub.DisconnectResponse, error) {
	return nil, nil
}
