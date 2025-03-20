package main

import (
	"flag"
	"fmt"
	"log"
	"mcp-1panel/operations/system"
	"mcp-1panel/operations/website"
	"mcp-1panel/utils"
	"os"
	"path/filepath"

	"github.com/mark3labs/mcp-go/server"
)

var (
	Version = utils.Version
)

func setupLogger() (*os.File, error) {
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("create log dir error: %v\n", err)
		return nil, err
	}

	logFilePath := filepath.Join(logDir, "mcp-1panel.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open log file error: %v\n", err)
		return nil, err
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return logFile, nil
}

func newMCPServer() *server.MCPServer {
	return server.NewMCPServer(
		"mcp-1panel",
		Version,
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)
}

func addTools(s *server.MCPServer) {
	s.AddTool(system.GetSystemInfoTool, system.GetSystemInfoHandle)
	s.AddTool(system.GetDashboardInfoTool, system.GetDashboardInfoHandle)
	s.AddTool(website.ListWebsitesTool, website.ListWebsiteHandle)
	s.AddTool(website.CreateWebsiteTool, website.CreateWebsiteHandle)
}

func runServer(transport string, port int64) error {
	mcpServer := newMCPServer()
	addTools(mcpServer)

	if transport == "sse" {
		log.Printf("SSE server listening on :%d", port)
		sseServer := server.NewSSEServer(mcpServer, server.WithBaseURL(fmt.Sprintf("http://localhost:%d", port)))
		if err := sseServer.Start(fmt.Sprintf(":%d", port)); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	} else {
		log.Printf("Run Stdio server")
		if err := server.ServeStdio(mcpServer); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}
	return nil
}

func main() {
	var (
		transport   string
		accessToken string
		host        string
		seePort     int64
	)
	flag.StringVar(&transport, "transport", "sse", "Transport type (stdio or sse)")
	flag.Int64Var(&seePort, "sse-port", 8000, "The port to start the sse server on")
	flag.StringVar(&accessToken, "token", "", "1Panel api key")
	flag.StringVar(&host, "host", "", "1Panel host (example:http://127.0.0.1:9999)")
	flag.Parse()

	logFile, err := setupLogger()
	if err != nil {
		fmt.Printf("setup logger error: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	if accessToken != "" {
		utils.SetAccessToken(accessToken)
	}
	if host != "" {
		utils.SetHost(host)
	}

	if err := runServer(transport, seePort); err != nil {
		fmt.Printf("server run error: %v\n", err)
		panic(err)
	}
}
