main_package_path = ./cmd/myweather-api
datafiles = ./configs
binary_name = weather-api
BIN_DIR = bin
CONFIGS_DIR = bin/configs

# Detect OS properly
UNAME_S := $(shell uname -s)
ifeq ($(OS), Windows_NT)
    OS := Windows_NT
		TARGET = $(binary_name).exe
    GOOS = windows
    GOARCH = amd64
		datafiles = .\configs
		CONFIGS_DIR = .\bin\configs
		COPY = xcopy /E /I /Y
    MKDIR = mkdir
		BUILD_CMD = powershell -Command "$$env:GOOS='linux'; $$env:GOARCH='arm64'; go build -o bin/weather-api $(main_package_path)/main.go"
else
    OS := $(UNAME_S)
		ifeq ($(OS), Linux)
			TARGET = $(binary_name)
			GOOS = linux
			GOARCH = arm64
		else ifeq ($(OS), Darwin)
			GOOS = darwin
			GOARCH = amd64
		endif
		TARGET = $(binary_name)
		COPY = cp -r
		MKDIR = mkdir -p
		BUILD_CMD = GOOS=linux GOARCH=arm64 go build -o bin/weather-api $(main_package_path)/main.go
endif

# Build API for detected system
build:
	@echo "Building for $(OS) ($(GOOS)-$(GOARCH))..."
	go build -o $(BIN_DIR)/$(TARGET) $(main_package_path)/main.go
	$(MKDIR) $(CONFIGS_DIR)
	$(COPY) $(datafiles) $(CONFIGS_DIR)

# Linux ARM64 build (Raspberry Pi) - i deployed this on my own raspberry pi
build-raspdist:
	$(BUILD_CMD)
	$(MKDIR) $(CONFIGS_DIR)
	$(COPY) $(datafiles) $(CONFIGS_DIR)

ifeq ($(OS), Windows_NT)
    clean:
			if exist $(BIN_DIR) rmdir /S /Q $(BIN_DIR)
else
    clean:
			@if [ -d "$(BIN_DIR)" ]; then rm -rf $(BIN_DIR); fi
endif

# Build for detected OS
all: clean build
