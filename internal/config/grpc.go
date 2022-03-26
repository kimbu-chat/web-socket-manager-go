package config

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kimbu-chat/web-socket-manager-go/internal/apiproto"
)

var client pb.CentrifugoApiClient
var conenction *grpc.ClientConn

func InitGRPCCleint() {
	addr := os.Getenv("CENTRIFUGO_GRPC_CLIENT_ADDRESS")

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	client = pb.NewCentrifugoApiClient(conn)
}

func GetGRPCClient() pb.CentrifugoApiClient {
	return client
}

func CloseGRPCConnection() {
	conenction.Close()
}
