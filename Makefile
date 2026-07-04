ARTIFACT = todo_server
SRC = cmd/server/main.go
OBJ = bin/$(ARTIFACT)


build: $(SRC)
	@echo "Building $(ARTIFACT)....."
	@go build -o $(OBJ) $(SRC)
	@echo "Build successful....."

run: build
	@echo "Running $(ARTIFACT)....."
	@./$(OBJ)
