#!/bin/bash

# Nome do CLI
CLI_NAME="expense-tracker"

# Compilar para Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o ./bin/${CLI_NAME}

# Compilar para Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o ./bin/${CLI_NAME}.exe 

echo "Build concluído!"