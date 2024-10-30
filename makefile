BINARY_NAME=./build/structure-mapper.exe
INSTALL_DIR=/usr/local/bin

build: $(BINARY_NAME)

$(BINARY_NAME):
	go build -o $(BINARY_NAME)

run: build
	$(BINARY_NAME)

install: build
	go install github.com/1Solon/structure-mapper

clean:
	go clean
ifeq ($(OS),Windows_NT)
	del build\\structure-mapper.exe
else
	rm -f $(BINARY_NAME)
endif

.PHONY: build run clean install