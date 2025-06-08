package grpcserver

import (
	"net"
	"time"

	"github.com/chiyonn/swarmyard/api/proto/pricefeed"
	"github.com/chiyonn/swarmyard/internal/logger"
	"google.golang.org/grpc"
)

type PriceFeedServerImpl struct {
	pricefeed.UnimplementedPriceFeedServer
}

func (s *PriceFeedServerImpl) SubscribePrices(req *pricefeed.PriceRequest, stream pricefeed.PriceFeed_SubscribePricesServer) error {
	snapshot := &pricefeed.PriceSnapshotList{
		Base: req.Base,
		Rates: []*pricefeed.PriceSnapshot{
			{
				Pair:      req.Base + "/USDT",
				Price:     123.45,
				Timestamp: 999999,
			},
		},
		Timestamp: 999999,
	}

	for {
		time.Sleep(5 * time.Second)
		if err := stream.Send(snapshot); err != nil {
			return err
		}
	}
}

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func New() (*Server, error) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	pricefeed.RegisterPriceFeedServer(s, &PriceFeedServerImpl{})

	return &Server{
		grpcServer: s,
		listener:   lis,
	}, nil
}

func (s *Server) Start() error {
	logger.Info("gRPC server listening on :50052")
	return s.grpcServer.Serve(s.listener)
}
