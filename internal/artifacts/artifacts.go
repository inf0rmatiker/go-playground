package artifacts

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/walle/targz"
)

// Gameplan:
// 0. Ensure cfg.ArtifactsDir exists, create if not.
// 1. Unpack/extract tarball to cfg.ArtifactsDir location.
// 2. Process the unpacked contents.
// 3. Update repofiles baseurls for relevant repos.
// 4. Update symlinks

// ExtractArtifacts extracts the provided artifacts archive (tgz file) to
// the configured artifacts directory, creating it if it does not already exist.
func ExtractArtifacts(ctx context.Context, archivePath, artifactsDir string, logger *log.Logger) error {

	fqdn := GetFqdn()
	logger.Infof("FQDN: %s", fqdn)

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
		logger.Infof("%s -> %s", header.Name, targetPath)
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

			// Catch *.repo files and fix their baseurls to point to Admin node's artifact webserver
			// instead of DAOS public packages repo.
			if filepath.Ext(header.Name) == ".repo" {
				log.Infof("Found repofile: %s", header.Name)
				directory := filepath.Dir(strings.TrimPrefix(header.Name, "./"))
				if err = ReplaceBaseurl(ctx, outFile, tarReader, fqdn, directory, logger); err != nil {
					return fmt.Errorf("failed to update baseurl in repofile %s: %w", targetPath, err)
				}
			} else {
				// Just copy original contents (as an IO stream) from tarReader to outFile
				if _, err := io.Copy(outFile, tarReader); err != nil {
					return fmt.Errorf("failed to copy content to file %s: %w", targetPath, err)
				}
			}
		default:
			log.Printf("Skipping unsupported tar entry type: %c for %s\n", header.Typeflag, header.Name)
		}
	}
}

func ExtractArtifactsWithLib(ctx context.Context, archivePath, artifactsDir string, logger *log.Logger) error {
	// Create the artifacts directory if it does not already exist. If it does
	// already exist, no error will be returned.
	if err := os.MkdirAll(artifactsDir, 0644); err != nil {
		logger.Errorf("Failed to create artifacts directory %s: %v", artifactsDir, err)
		return err
	}

	return targz.Extract(archivePath, artifactsDir)
}

func GetFqdn() string {
	cmd := exec.Command("/bin/hostname", "-f")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
	fqdn := out.String()
	fqdn = fqdn[:len(fqdn)-1] // removing EOL
	return fqdn
}

func ReplaceBaseurl(ctx context.Context, outFile *os.File, tarReader *tar.Reader, fqdn, targetDir string, logger *log.Logger) error {
	// Read file line by line, looking for baseurl lines to update
	scanner := bufio.NewScanner(tarReader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "baseurl=") {
			log.Infof("Found baseurl line: %s", line)
			line = fmt.Sprintf("baseurl=https://%s/artifacts/%s", fqdn, targetDir)
		}
		if _, err := outFile.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write to file %s: %w", outFile.Name(), err)
		}
	}
	return nil
}
