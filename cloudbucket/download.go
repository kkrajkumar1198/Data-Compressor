package cloudbucket

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kkrajkumar1198/Zocket/initializers"
	"golang.org/x/sync/errgroup"
)

// DownloadAndCompressImages downloads and compresses the specified images from GCS
// and returns the list of compressed file locations.
func DownloadAndCompressImages(ImagesNameList []string) ([]string, error) {
	ctx := context.Background()
	client, err := initializers.GetGCSClient()
	if err != nil {
		log.Printf("Failed to Connect to GCS Client: %s", err)
		return nil, err
	}

	// Create a temporary directory to store downloaded images
	tempDir := "/Users/rk/Documents/assignment/Zocket/savehere/"
	erro := os.MkdirAll(tempDir, os.ModePerm)
	if erro != nil {
		log.Printf("Failed to create a temp folder: %s", erro)
		return nil, erro
	}
	defer os.RemoveAll(tempDir)

	// Use an errgroup to synchronize goroutines
	var eg errgroup.Group

	// Slice to store the compressed filenames
	var compressedFilenames []string

	// Download and compress each file concurrently
	for _, filename := range ImagesNameList {
		filename := filename

		eg.Go(func() error {
			// Remove the file extension
			baseFilename := strings.TrimSuffix(filename, filepath.Ext(filename))

			// Download the image from GCS
			object := client.Bucket(initializers.GCSBucket).Object(filename)
			if object == nil {
				log.Printf("Failed to create GCS object for file %s", filename)
				return fmt.Errorf("failed to create gcs object for file %s", filename)
			}
			rc, err := object.NewReader(ctx)
			if err != nil {
				log.Printf("Failed to create GCS object reader for file %s: %s", filename, err)
				return err
			}
			defer rc.Close()

			// Create a new zip file for this image
			zipFileName := fmt.Sprintf("/Users/rk/Documents/assignment/Zocket/compressed/%s.zip", baseFilename)
			zipFile, err := os.Create(zipFileName)
			if err != nil {
				log.Printf("Failed to create a Zip folder for file %s: %s", baseFilename, err)
				return err
			}
			defer zipFile.Close()

			// Create a zip writer to write compressed image
			zipWriter := zip.NewWriter(zipFile)
			defer zipWriter.Close()

			// Add the image to the zip archive
			zipImageFile, err := zipWriter.Create(filename)
			if err != nil {
				log.Printf("Failed to create zip file for file %s: %s", filename, err)
				return err
			}

			// Copy the image content to the zip archive
			_, err = io.Copy(zipImageFile, rc)
			if err != nil {
				log.Printf("Failed to copy image content to zip archive for file %s: %s", filename, err)
				return err
			}

			compressedFilenames = append(compressedFilenames, zipFileName)

			return nil
		})
	}

	// Wait for all goroutines to complete and handle errors
	if err := eg.Wait(); err != nil {
		log.Printf("Failed to download and compress images: %s", err)
		return nil, err
	}

	return compressedFilenames, nil
}
