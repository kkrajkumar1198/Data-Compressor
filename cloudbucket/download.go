package cloudbucket

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kkrajkumar1198/Zocket/initializers"
)

// DownloadAndSaveImageToLocal downloads an image file from Google Cloud Storage (GCS)
// and saves it to the local machine.
func DownloadAndSaveImageToLocal(ctx *gin.Context) {

	client, err := initializers.GetGCSClient(ctx)

	if err != nil {
		log.Fatal(ctx, "Failed to create GCS client: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create GCS client"})
		return
	}

	filename := ctx.PostForm("filename") // Get the filename from the route parameter

	// Construct the ob	ject path
	objectPath := filename // Adjust the path as needed
	log.Println("asdfasfdasfasdfasfasdfasfasd", objectPath, filename)
	rc, err := client.Bucket(GCSBucket).Object(objectPath).NewReader(ctx)
	if err != nil {
		log.Println(ctx, "Failed to create GCS reader: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create GCS reader: %v", err)})
		return
	}
	defer rc.Close()

	// Set the content type based on the image file format (e.g., "image/jpeg", "image/png", etc.)
	// You may need to determine the content type based on the file extension or other criteria.
	contentType := "image/jpeg" // Adjust as needed based on your image file format
	ctx.Header("Content-Type", contentType)

	// Create a file on the local machine to save the downloaded image
	localFilePath := fmt.Sprintf("/Users/rk/Documents/assignment/Zocket/savehere/%s", filename) // Adjust the local path as needed
	localFile, err := os.Create(localFilePath)
	if err != nil {
		log.Fatal(ctx, "Failed to create local file: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create local file"})
		return
	}
	defer localFile.Close()

	// Copy the image file content to both the response body and the local file
	if _, err := io.Copy(io.MultiWriter(ctx.Writer, localFile), rc); err != nil {
		log.Fatal(ctx, "Failed to copy file content: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file content"})
		return
	}

	// Optionally, you can send a success response to the client
	ctx.JSON(http.StatusOK, gin.H{"message": "Image downloaded and saved successfully", "localFilePath": localFilePath})
}
