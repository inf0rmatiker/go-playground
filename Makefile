.DEFAULT_GOAL := bin


# fmt - Run go fmt on all go files
.PHONY: fmt
fmt:
	go fmt ./...

# vet - Run vet fmt on all go files
.PHONY: vet
vet:
	go vet ./...

# bin - Builds the go binaries
.PHONY: bin
bin: fmt vet
	go build -v -o main .

.PHONY: clean
clean:
	rm -f main

