package stream

import (
	"bytes"
	context "context"
	"fmt"
	"io"

	capnp "capnproto.org/go/capnp/v3"
	"cloud.google.com/go/storage"
)

type FileUploaderServer struct {
	client *storage.Client
}

func NewFileUploaderServer(ctx context.Context) (*FileUploaderServer, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %w", err)
	}

	return &FileUploaderServer{
		client,
	}, nil
}

var _ FileUploader_Server = (*FileUploaderServer)(nil)

func (s *FileUploaderServer) Upload(ctx context.Context, call FileUploader_upload) error {
	filename, err := call.Args().Filename()
	if err != nil {
		return nil
	}

	results, err := call.AllocResults()
	if err != nil {
		return err
	}

	wc := s.client.Bucket("tableland-staging").Object(filename).NewWriter(context.Background())
	callback := &Callback{
		wc: wc,
	}

	err = results.SetCallback(FileUploader_Callback(capnp.NewClient(FileUploader_Callback_NewServer(callback))))
	if err != nil {
		return err
	}

	return nil
}

type Callback struct {
	wc *storage.Writer
}

func (c *Callback) Write(_ context.Context, params FileUploader_Callback_write) error {
	chunk, err := params.Args().Chunk()
	if err != nil {
		return nil
	}

	if _, err := io.Copy(c.wc, bytes.NewBuffer(chunk)); err != nil {
		return err
	}

	return nil
}

func (c *Callback) Done(_ context.Context, p FileUploader_Callback_done) error {
	if err := c.wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	return nil
}
