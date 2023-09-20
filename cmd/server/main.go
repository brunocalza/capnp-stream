package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/brunocalza/stream"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:"+os.Getenv("PORT"))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	ctx := context.Background()
	slog.Info("Listening", "port", os.Getenv("PORT"))

	server, err := stream.NewFileUploaderServer(ctx)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		defer conn.Close()

		client := stream.FileUploader_ServerToClient(server)
		rpcConn := rpc.NewConn(rpc.NewStreamTransport(conn), &rpc.Options{
			BootstrapClient: capnp.Client(client),
		})
		defer rpcConn.Close()

		// Block until the connection terminates.
		select {
		case <-rpcConn.Done():
			slog.Info("connection closed")
		case <-ctx.Done():
			conn.Close()
		}

	}
}
