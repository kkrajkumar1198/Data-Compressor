package cloudbucket

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kkrajkumar1198/Zocket/initializers"
)

// Test Function - this is not used anywhere in the project

func Upload(c *gin.Context) {
	ctx := c.Request.Context()
	err := c.Request.ParseMultipartForm(100 << 20) // Max Size Limit is 100 MB
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// fileKey is the name of key passed with form request
	fhs := c.Request.MultipartForm.File["fileKey"]
	// Multiple files can be passed as part of the form request
	var fileLinks []string
	for _, fh := range fhs {
		link, err := UploadFileToGCSBucket(ctx, fh)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fileLinks = append(fileLinks, link)
	}
	c.JSON(http.StatusOK, gin.H{"fileLinks": fileLinks})
}

func UploadFileToGCSBucket(ctx context.Context, fh *multipart.FileHeader) (string, error) {

	date := time.Now()
	dateStr := date.Format("01-02-2006")

	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	filename := generateFileNameForGCS(ctx, fh.Filename)
	filepath := fmt.Sprintf("%s/%s%s%s", dateStr, filename, initializers.Delimitor, fh.Filename)
	client, err := initializers.GetGCSClient()
	if err != nil {
		return "", err
	}
	wc := client.Bucket(initializers.GCSBucket).UserProject(initializers.ProjectID).Object(filepath).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	return filepath, nil
}

// generateFileNameForGCS will generate the resource path for a file.
// Combination of current time and filename to generate a unique entry.
func generateFileNameForGCS(ctx context.Context, name string) string {
	time := time.Now().UnixNano()
	var strArr []string
	strArr = append(strArr, name)
	strArr = append(strArr, strconv.Itoa(int(time)))
	var filename string
	for _, str := range strArr {
		filename = filename + str
	}
	return filename
}
