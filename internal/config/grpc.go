package config

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kimbu-chat/web-socket-manager-go/internal/apiproto"
)

var client pb.CentrifugoApiClient
var conenction *grpc.ClientConn

type keyAuth struct {
	key string
}

func (t keyAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "apikey " + t.key,
	}, nil
}

func (t keyAuth) RequireTransportSecurity() bool {
	return false
}

func initGRPCCleint() {
	var err error

	conenction, err = grpc.Dial(CentrifugoGRPCAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(keyAuth{CentrifugoAPIKey()}))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
		return
	}

	client = pb.NewCentrifugoApiClient(conenction)
}

func GetGRPCClient() pb.CentrifugoApiClient {
	return client
}

func closeGRPCConnection() error {
	return conenction.Close()
}
