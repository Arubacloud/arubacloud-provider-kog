# Makefile for Arubacloud Provider KOG Plugins

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

# Build directory
BUILD_DIR=plugins/bin

# Plugin directories
PLUGINS=compute-plugin container-plugin database-plugin network-plugin project-plugin schedule-plugin security-plugin storage-plugin

# Default target
.DEFAULT_GOAL := help

.PHONY: all build clean test fmt fmt-check vet check deps help $(PLUGINS)

# Build all plugins
all: build

# Build all plugins
build: $(PLUGINS)
	@echo "✅ All plugins built successfully!"

# Build individual plugins
$(PLUGINS):
	@echo "🔨 Building $@..."
	@mkdir -p $(BUILD_DIR)
	@cd plugins/cmd/$@ && $(GOBUILD) -o ../../../$(BUILD_DIR)/$@ .
	@echo "✅ $@ built successfully!"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete!"

# Run tests for all plugins
test:
	@echo "🧪 Running tests..."
	@for plugin in $(PLUGINS); do \
		echo "Testing $$plugin..."; \
		cd plugins/cmd/$$plugin && $(GOTEST) -v ./... || exit 1; \
		cd ../../..; \
	done
	@echo "✅ All tests passed!"

# Format code
fmt:
	@echo "📝 Formatting code..."
	@for plugin in $(PLUGINS); do \
		echo "Formatting $$plugin..."; \
		cd plugins/cmd/$$plugin && $(GOFMT) -w ./...; \
		cd ../../..; \
	done
	@echo "✅ Formatting complete!"

# Check if code is formatted (without modifying files)
fmt-check:
	@echo "🔍 Checking code formatting..."
	@for plugin in $(PLUGINS); do \
		echo "Checking formatting for $$plugin..."; \
		cd plugins/cmd/$$plugin && \
		UNFORMATTED=$$($(GOFMT) -l . 2>/dev/null); \
		if [ -n "$$UNFORMATTED" ]; then \
			echo "❌ $$plugin needs formatting:"; \
			echo "$$UNFORMATTED"; \
			echo "Run 'make fmt' to fix."; \
			cd ../../..; \
			exit 1; \
		fi; \
		cd ../../..; \
	done
	@echo "✅ All code is properly formatted!"

# Run go vet
vet:
	@echo "🔍 Running go vet..."
	@for plugin in $(PLUGINS); do \
		echo "Vetting $$plugin..."; \
		cd plugins/cmd/$$plugin && $(GOVET) ./... || exit 1; \
		cd ../../..; \
	done
	@echo "✅ Vet complete!"

# Run all checks (format check, vet, test)
check: fmt-check vet test
	@echo "✅ All checks passed!"

# Download dependencies for all plugins
deps:
	@echo "📦 Downloading dependencies..."
	@for plugin in $(PLUGINS); do \
		echo "Downloading dependencies for $$plugin..."; \
		cd plugins/cmd/$$plugin && $(GOMOD) download && $(GOMOD) tidy; \
		cd ../../..; \
	done
	@echo "✅ Dependencies downloaded!"

# Tidy all go.mod files
tidy:
	@echo "🧹 Tidying go.mod files..."
	@for plugin in $(PLUGINS); do \
		echo "Tidying $$plugin..."; \
		cd plugins/cmd/$$plugin && $(GOMOD) tidy; \
		cd ../../..; \
	done
	@echo "✅ Tidy complete!"

# Build and run a specific plugin (usage: make run PLUGIN=compute-plugin)
run:
	@if [ -z "$(PLUGIN)" ]; then \
		echo "❌ Please specify PLUGIN=plugin-name (e.g., make run PLUGIN=compute-plugin)"; \
		exit 1; \
	fi
	@echo "🚀 Running $(PLUGIN)..."
	@cd plugins/cmd/$(PLUGIN) && $(GOBUILD) -o ../../../$(BUILD_DIR)/$(PLUGIN) . && ../../../$(BUILD_DIR)/$(PLUGIN)

# Help target
help:
	@echo "Arubacloud Provider KOG Plugins - Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  make build          - Build all plugins"
	@echo "  make <plugin-name>  - Build a specific plugin (e.g., make compute-plugin)"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make check          - Run all checks (fmt-check, vet, test)"
	@echo "  make test           - Run tests for all plugins"
	@echo "  make fmt            - Format code for all plugins"
	@echo "  make fmt-check      - Check if code is formatted (without modifying)"
	@echo "  make vet            - Run go vet for all plugins"
	@echo "  make deps           - Download dependencies for all plugins"
	@echo "  make tidy           - Tidy go.mod files for all plugins"
	@echo "  make run PLUGIN=... - Build and run a specific plugin"
	@echo ""
	@echo "Available plugins:"
	@for plugin in $(PLUGINS); do \
		echo "  - $$plugin"; \
	done
	@echo ""
	@echo "Examples:"
	@echo "  make build                    # Build all plugins"
	@echo "  make compute-plugin           # Build only compute-plugin"
	@echo "  make check                    # Run all checks (fmt-check, vet, test)"
	@echo "  make fmt                      # Format all code"
	@echo "  make fmt-check                # Check formatting without modifying"
	@echo "  make vet                      # Run go vet on all plugins"
	@echo "  make run PLUGIN=compute-plugin # Build and run compute-plugin"

