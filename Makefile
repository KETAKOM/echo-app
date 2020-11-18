GO_CMD = go
VARSION = -u

all: tools fmt

tools:
	$(GO_CMD) get $(VARSION)
	$(GO_CMD) mod tidy

fmt:
	$(GO_CMD) fmt ./...
