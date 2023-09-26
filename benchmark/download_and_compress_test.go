package cloudbucket_test

import (
	"testing"

	"github.com/kkrajkumar1198/Zocket/cloudbucket"
)

// BenchmarkDownloadAndCompressImages benchmarks the DownloadAndCompressImages function.
func BenchmarkDownloadAndCompressImages(b *testing.B) {

	imageNames := []string{"nike-1.jpg", "nike-2.jpg", "nike-3.jpg"}

	// Bench mark test
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cloudbucket.DownloadAndCompressImages(imageNames)
		if err != nil {
			b.Fatalf("Benchmark failed: %v", err)
		}
	}
}
