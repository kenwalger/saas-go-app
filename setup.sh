#!/bin/bash

# Setup script for SaaS Go App

echo "Setting up SaaS Go App..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "Error: Node.js is not installed. Please install Node.js 18 or later."
    exit 1
fi

echo "Installing Go dependencies..."
go mod download
go mod tidy

echo "Setting up frontend..."
cd web/frontend
if [ ! -d "node_modules" ]; then
    echo "Installing npm dependencies..."
    npm install
else
    echo "npm dependencies already installed"
fi
cd ../..

echo ""
echo "Setup complete!"
echo ""
echo "Next steps:"
echo "1. Create a .env file (see .env.example)"
echo "2. Start the backend: make run"
echo "3. Start the frontend: cd web/frontend && npm run dev"
echo ""

