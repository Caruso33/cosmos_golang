package parts

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func GetConnection() (*grpc.ClientConn, error) {
	var grpc_addr = "grpc-cosmoshub-ia.cosmosia.notional.ventures:443"

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		panic(errors.Wrap(err, "cannot load root CA certs"))
	}
	creds := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	// conn, err := grpc.DialContext(ctx, addr,
	// 	grpc.WithTransportCredentials(creds),
	// )
	grpcConn, err := grpc.Dial(
		grpc_addr,
		grpc.WithTransportCredentials(creds),
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to grpc %v", err.Error())
	}

	return grpcConn, nil
}
