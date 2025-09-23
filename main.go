package main

import (
	"context"
	"fmt"
	"os"

	"github.com/inf0rmatiker/go-playground/internal/artifacts"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Printf("Hello world\n")

	args := os.Args
	var src, dest string
	if len(args) != 3 {
		log.Errorf("Usage: %s <source> <destination>", args[0])
	}
	src = args[1]
	dest = args[2]

	err := artifacts.ExtractArtifacts(context.TODO(), log.StandardLogger(), src, dest)
	if err != nil {
		log.Errorf("Error extracting artifacts: %v", err)
	}
}
