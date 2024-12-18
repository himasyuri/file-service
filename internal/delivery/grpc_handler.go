package delivery

import (
	"context"
	"io"

	"file-storage/internal/domain"
	pb "file-storage/pkg/fileservicepb"
)

type GRPCHandler struct {
	pb.UnimplementedFileServiceServer
	fileService domain.FileService
}

func NewGRPCHandler(service domain.FileService) *GRPCHandler {
	return &GRPCHandler{fileService: service}
}

func (h *GRPCHandler) UploadFile(stream pb.FileService_UploadFileServer) error {
	var fileName string
	var data []byte

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if chunk.GetMetadata() != nil {
			fileName = chunk.GetMetadata().GetName()
		} else {
			data = append(data, chunk.GetContent()...)
		}
	}

	return h.fileService.UploadFile(context.Background(), fileName, data)
}

func (h *GRPCHandler) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	files, err := h.fileService.ListFiles(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListFilesResponse{}
	for _, f := range files {
		resp.Files = append(resp.Files, &pb.FileMetadata{
			Name:      f.FileName,
			CreatedAt: f.CreatedAt.Unix(),
			UpdatedAt: f.UpdatedAt.Unix(),
		})
	}
	return resp, nil
}

func (h *GRPCHandler) DownloadFile(req *pb.DownloadFileRequest, stream pb.FileService_DownloadFileServer) error {
	data, err := h.fileService.DownloadFile(context.Background(), req.GetName())
	if err != nil {
		return err
	}

	return stream.Send(&pb.FileChunk{Content: data})
}

