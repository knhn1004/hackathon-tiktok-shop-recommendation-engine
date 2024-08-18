package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	protoDir := filepath.Join("cmd", "generate_proto", "protos")
	outputDir := "internal/proto/recommendation"

	// Check if protoDir exists
	if _, err := os.Stat(protoDir); os.IsNotExist(err) {
		fmt.Printf("Proto directory does not exist: %s\n", protoDir)
		os.Exit(1)
	}

	// Create the output directory if it doesn't exist
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Printf("Failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	protoFile := filepath.Join(protoDir, "recommendation.proto")

	// Check if protoFile exists
	if _, err := os.Stat(protoFile); os.IsNotExist(err) {
		fmt.Printf("Proto file does not exist: %s\n", protoFile)
		os.Exit(1)
	}

	// Generate Go code and gRPC code
	cmd := exec.Command("protoc",
		fmt.Sprintf("--proto_path=%s", protoDir),
		fmt.Sprintf("--go_out=%s", outputDir),
		fmt.Sprintf("--go-grpc_out=%s", outputDir),
		"--go_opt=paths=source_relative",
		"--go-grpc_opt=paths=source_relative",
		protoFile,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to generate Protocol Buffer files: %v\n", err)
		os.Exit(1)
	}

	// Check if files were actually generated
	expectedFiles := []string{
		filepath.Join(outputDir, "recommendation.pb.go"),
		filepath.Join(outputDir, "recommendation_grpc.pb.go"),
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Printf("Expected generated file not found: %s\n", file)
			os.Exit(1)
		} else {
			fmt.Printf("Generated file found: %s\n", file)
		}
	}

	fmt.Println("Protocol Buffer files generated successfully.")
}