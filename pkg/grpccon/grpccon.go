package grpccon

import (
	"errors"
	"log"
	"net"

	authPb "github.com/ekkasitProject/shop-game/modules/auth/authPb"
	itemPb "github.com/ekkasitProject/shop-game/modules/item/itemPb"
	playerPb "github.com/ekkasitProject/shop-game/modules/player/playerPb"

	"github.com/ekkasitProject/shop-game/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		Player() playerPb.PlayerGrpcServiceClient
		Item() itemPb.ItemGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}

	grpcAuth struct {
		secretKey string
	}
)

// func (g *grpcAuth) unaryAuthorization(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		log.Printf("Error: Metadata not found")
// 		return nil, errors.New("error: metadata not found")
// 	}

// 	authHeader, ok := md["auth"]
// 	if !ok {
// 		log.Printf("Error: Metadata not found")
// 		return nil, errors.New("error: metadata not found")
// 	}

// 	if len(authHeader) == 0 {
// 		log.Printf("Error: Metadata not found")
// 		return nil, errors.New("error: metadata not found")
// 	}

// 	claims, err := jwtauth.ParseToken(g.secretKey, string(authHeader[0]))
// 	if err != nil {
// 		log.Printf("Error: Parse token failed: %s", err.Error())
// 		return nil, errors.New("error: token is invalid")
// 	}
// 	log.Printf("claims: %v", claims)

// 	return handler(ctx, req)
// }

func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Player() playerPb.PlayerGrpcServiceClient {
	return playerPb.NewPlayerGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Item() itemPb.ItemGrpcServiceClient {
	return itemPb.NewItemGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("Error: Grpc client connection failed: %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}

	return &grpcClientFactory{
		client: clientConn,
	}, nil
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	// grpcAuth := &grpcAuth{
	// 	secretKey: cfg.ApiSceretKey,
	// }

	// opts = append(opts, grpc.UnaryInterceptor(grpcAuth.unaryAuthorization))

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}

	return grpcServer, lis
}
