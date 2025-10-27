package main

import (
	"fmt"

	"github.com/inf0rmatiker/go-playground/internal/generics"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Printf("Hello world\n")
	log.Info("Hello world from logger\n")

	log.Infof("'a' lt 'b' = %t", generics.Compare("a", "b", "lt"))
	log.Infof("'b' lt 'a' = %t", generics.Compare("b", "a", "lt"))
	log.Infof("1 lt 3 = %t", generics.Compare(1, 3, "lt"))

}
