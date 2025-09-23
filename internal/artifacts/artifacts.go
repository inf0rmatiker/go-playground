package artifacts

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// Gameplan:
// 0. Ensure cfg.ArtifactsDir exists, create if not.
// 1. Unpack/extract tarball to cfg.ArtifactsDir location.
// 2. Process the unpacked contents.
// 3. Update repofiles baseurls for relevant repos.
// 4. Update symlinks

// ExtractArtifacts extracts the provided artifacts archive (tgz file) to
// the configured artifacts directory, creating it if it does not already exist.
func ExtractArtifacts(ctx context.Context, logger *log.Logger, archivePath, artifactsDir string) error {

	// Create the artifacts directory if it does not already exist. If it does
	// already exist, no error will be returned.
	if err := os.MkdirAll(artifactsDir, 0644); err != nil {
		logger.Errorf("Failed to create artifacts directory %s: %v", artifactsDir, err)
		return err
	}

	// Open the file
	file, err := os.Open(archivePath)
	if err != nil {
		logger.Errorf("Failed to open artifacts archive %s: %v", archivePath, err)
		return err
	}
	defer file.Close()

	// Create a gunzip reader on the file
	gunzipReader, err := gzip.NewReader(file)
	if err != nil {
		logger.Errorf("Failed to create gunzip reader: %v", err)
		return err
	}
	defer gunzipReader.Close()

	// Create tar reader on the gunzip reader
	tarReader := tar.NewReader(gunzipReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			// End of the tarball archive, return
			return nil
		}
		if err != nil {
			logger.Errorf("Failed to read tarball header: %v", err)
			return err
		}

		targetPath := filepath.Join(artifactsDir, header.Name)
		logger.Infof("Extracting to %s", targetPath)
		switch header.Typeflag {
		case tar.TypeDir:
			// Create directory with same permissions as in the tarball
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetPath, err)
			}
		case tar.TypeReg:
			// Create file and copy content
			outFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", targetPath, err)
			}
			defer outFile.Close() // Ensure each file is closed

			// Stream from tarReader to outFile
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("failed to copy content to file %s: %w", targetPath, err)
			}
		default:
			log.Printf("Skipping unsupported tar entry type: %c for %s\n", header.Typeflag, header.Name)
		}
	}
}
