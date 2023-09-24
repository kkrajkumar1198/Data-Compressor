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

const (
	// GCSBucket name
	GCSBucket = "zocket-source-imgs"
	// ProjectID Google Project ID name
	ProjectID = "zocket-400012"
	delimitor = "_"
)

// Upload API will take Multi Form data as an input and store the object to Google storage
func Upload(c *gin.Context) {
	ctx := c.Request.Context()
	err := c.Request.ParseMultipartForm(100 << 20) // Max Size Limit is 100 MB
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// fileKey is the name of key passed with multi-form request
	fhs := c.Request.MultipartForm.File["fileKey"]
	// Multiple files can be passed as part of the multi-form request
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

// UploadFileToGCSBucket will create a date-wise directory bucket
func UploadFileToGCSBucket(ctx context.Context, fh *multipart.FileHeader) (string, error) {

	date := time.Now()
	dateStr := date.Format("01-02-2006") // MM-DD-YYYY Format

	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	filename := generateFileNameForGCS(ctx, fh.Filename)
	filepath := fmt.Sprintf("%s/%s%s%s", dateStr, filename, delimitor, fh.Filename)
	client, err := initializers.GetGCSClient(ctx)
	if err != nil {
		return "", err
	}
	wc := client.Bucket(GCSBucket).UserProject(ProjectID).Object(filepath).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	return filepath, nil
}

// generateFileNameForGCS will generate the resource path for a file.
// It will use a combination of current time and filename to generate a unique entry.
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
