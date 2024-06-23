package controller

import (
	"CoFiler/service"
	"CoFiler/utils/logging"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"sync"
)

type FileHandler struct {
	service *service.FileService
}

func NewFileHandler(service *service.FileService) *FileHandler {
	return &FileHandler{
		service: service,
	}
}

func FileRouter(e *gin.Engine, h *FileHandler) {
	v1 := e.Group("/api/v1")
	fileV1 := v1.Group("/file")

	fileV1.Use()
	{
		fileV1.POST("upload", h.uploadFile)
	}
}

func (h *FileHandler) uploadFile(c *gin.Context) {
	logger := logging.FromContext(c)
	logger.Info("UploadFile")

	var err error

	// 요청 사이즈 확인
	if err = c.Request.ParseMultipartForm(10 << 20); err != nil {
		GeneralResponse(c, http.StatusBadRequest, "Exceed file size", 0, "")
		return
	}

	if files, ok := c.Request.MultipartForm.File["files"]; !ok || len(files) == 0 {
		GeneralResponse(c, http.StatusBadRequest, "No files uploaded", 0, "")
		return
	} else {
		var wg sync.WaitGroup
		errChan := make(chan error, len(files))

		for _, fileHeader := range files {
			wg.Add(1)
			go func(fileHeader *multipart.FileHeader) {
				defer wg.Done()
				if file, gErr := fileHeader.Open(); gErr != nil {
					errChan <- gErr
					return
				} else {
					defer file.Close()
					if gErr = h.service.Save(fileHeader.Filename, file); gErr != nil {
						errChan <- gErr
					}
				}
			}(fileHeader)
		}
		wg.Wait()
		close(errChan)

		for err := range errChan {
			if err != nil {
				GeneralResponse(c, http.StatusInternalServerError, "Error saving the file", 0, err.Error())
				return
			}
		}
		GeneralResponse(c, http.StatusOK, "Success", 0, "")
	}
}
