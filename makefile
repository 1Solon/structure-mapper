BINARY_NAME=./build/structure-mapper.exe
SRC=./src/main.go

build: $(BINARY_NAME)

$(BINARY_NAME): $(SRC)
	go build -o $(BINARY_NAME) $(SRC)

run: build
	$(BINARY_NAME)

clean:
	go clean
ifeq ($(OS),Windows_NT)
	del build\\structure-mapper.exe
else
	rm -f $(BINARY_NAME)
endif

.PHONY: build run clean