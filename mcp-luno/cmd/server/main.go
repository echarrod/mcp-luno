package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/echarrod/mcp-luno/internal/config"
	"github.com/echarrod/mcp-luno/internal/logging"
	"github.com/echarrod/mcp-luno/internal/server"
	"github.com/joho/godotenv"
)

const (
	appName    = "mcp-luno"
	appVersion = "0.1.0"
)

func main() {
	// Try different possible locations for the .env file
	envPaths := []string{
		".env",    // Current directory
		"../.env", // Parent directory
	}

	envLoaded := false
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			log.Printf("Successfully loaded environment from %s", path)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		log.Println("Warning: No .env file found or unable to load it. Make sure environment variables are set.")
		// Print current directory for debugging
		if cwd, err := os.Getwd(); err == nil {
			log.Printf("Current working directory: %s", cwd)
		}
	}

	// Parse command line flags
	transportType := flag.String("transport", "stdio", "Transport type (stdio or sse)")
	sseAddr := flag.String("sse-address", "localhost:8080", "Address for SSE transport")
	lunoDomain := flag.String("domain", "", "Luno API domain (default: api.luno.com)")
	logLevel := flag.String("log-level", "info", "Log level (debug, info, warn, error)")
	flag.Parse()

	// Set up basic logger first
	level := parseLogLevel(*logLevel)
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	logger := slog.New(consoleHandler)
	slog.SetDefault(logger)

	// Load configuration
	cfg, err := config.Load(*lunoDomain)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create MCP server with logging hooks
	mcpServer := server.NewMCPServer(appName, appVersion, cfg, logging.MCPHooks())

	// Now enhance the logger with MCP notification capability
	mcpHandler := logging.NewMCPNotificationHandler(mcpServer, level)
	multiHandler := logging.NewMultiHandler(consoleHandler, mcpHandler)
	enhancedLogger := slog.New(multiHandler)
	slog.SetDefault(enhancedLogger)

	// Setup signal handling for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		slog.Info("Received shutdown signal")
		cancel()
	}()

	// Start the server with the selected transport
	switch *transportType {
	case "stdio":
		slog.Info("Starting Luno MCP server using stdio transport")
		if err := server.ServeStdio(ctx, mcpServer); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	case "sse":
		slog.Info("Starting Luno MCP server using SSE transport", slog.String("address", *sseAddr))
		if err := server.ServeSSE(ctx, mcpServer, *sseAddr); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	default:
		log.Fatalf("Invalid transport type: %s. Must be 'stdio' or 'sse'", *transportType)
	}
}

func parseLogLevel(level string) slog.Level {
	var l slog.Level
	if err := l.UnmarshalText([]byte(level)); err != nil {
		return slog.LevelInfo
	}
	return l
}
