package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kimbu-chat/web-socket-manager-go/internal/apiproto"
)

var client pb.CentrifugoApiClient
var conenction *grpc.ClientConn

func initGRPCCleint() {
	addr := os.Getenv("CENTRIFUGO_GRPC_ADDRESS")

	var err error
	conenction, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
