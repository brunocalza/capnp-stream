package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"os"

	"capnproto.org/go/capnp/v3/flowcontrol"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/brunocalza/stream"
	"github.com/schollz/progressbar/v3"
)

func main() {
	ctx := context.Background()

	conn, err := tls.Dial("tcp", "localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rpcConn := rpc.NewConn(rpc.NewStreamTransport(conn), nil)
	defer rpcConn.Close()

	client := stream.FileUploader(rpcConn.Bootstrap(ctx))
	err = upload(context.Background(), client, "sample.txt")
	fmt.Println(err)
}

func upload(ctx context.Context, client stream.FileUploader, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return fmt.Errorf("fstat: %s", err)
	}

	bar := progressbar.DefaultBytes(
		fi.Size(),
		"Uploading file...",
	)

	f, r := client.Upload(ctx, func(params stream.FileUploader_upload_Params) error {
		_ = params.SetFilename(filename)

		return nil
	})
	defer r()

	callback := f.Callback()
	callback.SetFlowLimiter(flowcontrol.NewFixedLimiter(1 << 24))

	buf := make([]byte, 1<<20) // 1MiB
	for {
		n, err := file.Read(buf[:cap(buf)])
		buf = buf[:n]
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("read file: %s", err)
			}
			break
		}

		f, r := callback.Write(context.Background(), func(p stream.FileUploader_Callback_write_Params) error {
			return p.SetChunk(buf)
		})
		r()

		if _, err := f.Struct(); err != nil {
			return fmt.Errorf("waiting for write: %s", err)
		}

		// if err := callback.Write(ctx, func(p stream.FileUploader_Callback_write_Params) error {
		// 	return p.SetChunk(buf)
		// }); err != nil {
		// 	return err
		// }

		_, _ = bar.Write(buf)
	}

	doneFuture, doneRelease := callback.Done(ctx, func(p stream.FileUploader_Callback_done_Params) error {
		return nil
	})
	defer doneRelease()
	if _, err := doneFuture.Struct(); err != nil {
		return fmt.Errorf("done: %s", err)
	}

	if err := client.WaitStreaming(); err != nil {
		return fmt.Errorf("wait streaming: %s", err)
	}

	return nil
}
