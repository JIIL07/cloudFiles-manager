package app

import (
	"github.com/JIIL07/jcloud/internal/client/models"
	"github.com/JIIL07/jcloud/internal/client/storage"
	"github.com/JIIL07/jcloud/internal/config"
	"github.com/JIIL07/jcloud/internal/logger"
	"github.com/JIIL07/jcloud/pkg/home"
	"log/slog"
)

type ClientContext struct {
	Cfg    *config.ClientConfig
	common service

	// Services
	FileService     *FileService
	StorageService  *StorageService
	PathsService    *PathsService
	LoggerService   *LoggerService
	AnchorService   *AnchorService
	DeltaService    *DeltaService
	SnapshotService *SnapshotService
}

type service struct {
	Context *ClientContext
}

type FileService struct {
	*service
	F *models.File
}

func NewFileService() *FileService {
	fs := &FileService{}
	return fs
}

type StorageService struct {
	*service
	S *storage.SQLite
}

type PathsService struct {
	*service
	P *home.Paths
}

type LoggerService struct {
	*service
	L *slog.Logger
}

type AnchorService struct {
	*service
}

type DeltaService struct {
	*service
}

type SnapshotService struct {
	*service
}

func NewAppContext(cfg *config.ClientConfig) (*ClientContext, error) {
	p := home.SetPaths()

	s := storage.MustInit()

	l := logger.NewClientLogger(p.JlogFile)

	context := &ClientContext{
		Cfg:    cfg,
		common: service{},
	}

	context.common.Context = context

	context.FileService = &FileService{
		service: &context.common,
		F:       &models.File{},
	}
	context.StorageService = &StorageService{
		service: &context.common,
		S:       s,
	}
	context.PathsService = &PathsService{
		service: &context.common,
		P:       p,
	}
	context.LoggerService = &LoggerService{
		service: &context.common,
		L:       l,
	}
	context.AnchorService = &AnchorService{
		service: &context.common,
	}
	context.DeltaService = &DeltaService{
		service: &context.common,
	}
	context.SnapshotService = &SnapshotService{
		service: &context.common,
	}

	return context, nil
}
