# Имя бинарного файла
BINARY_NAME=delivery

# Исходные файлы
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

# Цель по умолчанию
all: build

# Сборка приложения
build:
	@echo "Building the application..."
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/app/main.go

# Запуск приложения
run: build
	@echo "Running the application..."
	./$(BINARY_NAME)

# Очистка
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Запуск тестов
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

.PHONY: all build run clean test
