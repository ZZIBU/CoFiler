package router

import (
	"Core/service"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"sync"
)

type FileRouter struct {
	router  *Router
	service *service.FileService
}

func InitFileRouter(router *Router, fileService *service.FileService) {
	fr := &FileRouter{
		router:  router,
		service: fileService,
	}

	baseUri := "/file"
	fr.router.POST(baseUri+"/upload", fr.uploadFile)
}

func (fr *FileRouter) uploadFile(c *gin.Context) {
	var err error

	// 요청 사이즈 확인
	if err = c.Request.ParseMultipartForm(10 << 20); err != nil {
		fr.router.GeneralResponse(c, http.StatusBadRequest, "Exceed file size", 0, "")
		return
	}

	if files, ok := c.Request.MultipartForm.File["files"]; !ok || len(files) == 0 {
		fr.router.GeneralResponse(c, http.StatusBadRequest, "No files uploaded", 0, "")
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
					if gErr = fr.service.Save(fileHeader.Filename, file); gErr != nil {
						errChan <- gErr
					}
				}
			}(fileHeader)
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			if err != nil {
				fr.router.GeneralResponse(c, http.StatusInternalServerError, "Error saving the file", 0, err.Error())
				return
			}
		}

		fr.router.GeneralResponse(c, http.StatusOK, "Success", 0, "")
	}
}
