#!/bin/bash

set -e

echo "==================================="
echo "WSL2 Ubuntu Setup Script for kp-backend"
echo "==================================="
echo ""

# Color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    echo -e "${GREEN}✓${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_info() {
    echo -e "${YELLOW}ℹ${NC} $1"
}

# Check if running in WSL
if ! grep -qi microsoft /proc/version && ! grep -qi wsl /proc/version; then
    print_error "This script is designed to run in WSL2 Ubuntu"
    exit 1
fi

print_info "Starting WSL2 Ubuntu setup..."
echo ""

# Update system packages
print_info "Updating system packages..."
sudo apt update && sudo apt upgrade -y > /dev/null 2>&1
print_status "System packages updated"

# Install Docker
if ! command -v docker &> /dev/null; then
    print_info "Installing Docker..."
    sudo apt install -y docker.io > /dev/null 2>&1
    print_status "Docker installed"
else
    print_status "Docker already installed"
fi

# Install Docker Compose
if ! command -v docker-compose &> /dev/null; then
    print_info "Installing Docker Compose..."
    sudo apt install -y docker-compose > /dev/null 2>&1
    print_status "Docker Compose installed"
else
    print_status "Docker Compose already installed"
fi

# Install Make
if ! command -v make &> /dev/null; then
    print_info "Installing Make..."
    sudo apt install -y make > /dev/null 2>&1
    print_status "Make installed"
else
    print_status "Make already installed"
fi

# Install Go
if ! command -v go &> /dev/null; then
    print_info "Installing Go..."
    sudo apt install -y golang-go > /dev/null 2>&1
    print_status "Go installed"
else
    print_status "Go already installed"
fi

# Install Git
if ! command -v git &> /dev/null; then
    print_info "Installing Git..."
    sudo apt install -y git > /dev/null 2>&1
    print_status "Git installed"
else
    print_status "Git already installed"
fi

# Install MySQL Client
if ! command -v mysql &> /dev/null; then
    print_info "Installing MySQL Client..."
    sudo apt install -y mysql-client > /dev/null 2>&1
    print_status "MySQL Client installed"
else
    print_status "MySQL Client already installed"
fi

# Configure Docker to run without sudo
print_info "Configuring Docker to run without sudo..."
sudo usermod -aG docker $USER 2>/dev/null || true
print_status "Docker group configured"

# Start Docker service
print_info "Starting Docker service..."
sudo service docker start > /dev/null 2>&1
print_status "Docker service started"

# Set Git autocrlf to prevent line ending issues
print_info "Configuring Git line ending settings..."
git config --global core.autocrlf input
print_status "Git line ending settings configured"

echo ""
print_info "Installing Go dependencies..."
go mod download > /dev/null 2>&1
print_status "Go dependencies installed"

echo ""
print_info "Creating .env file from .env-sample..."
[ -f .env ] || cp .env-sample .env
print_status ".env file created"

echo ""
print_info "Starting Docker containers with database migration..."
make up > /dev/null 2>&1

# Wait for containers to be healthy
echo ""
print_info "Waiting for containers to be fully initialized..."
sleep 5

# Check container status
if docker-compose ps | grep -q "portfolio-app.*Up"; then
    print_status "Application container is running"
else
    print_error "Application container failed to start"
    echo ""
    print_info "Checking Docker logs:"
    docker-compose logs app
    exit 1
fi

if docker-compose ps | grep -q "db.*Up"; then
    print_status "Database container is running"
else
    print_error "Database container failed to start"
    echo ""
    print_info "Checking Docker logs:"
    docker-compose logs db
    exit 1
fi

if docker-compose ps | grep -q "phpmyadmin.*Up"; then
    print_status "phpMyAdmin container is running"
else
    print_error "phpMyAdmin container failed to start"
    echo ""
    print_info "Checking Docker logs:"
    docker-compose logs phpmyadmin
    exit 1
fi

echo ""
echo "==================================="
print_status "Setup completed successfully!"
echo "==================================="
echo ""
echo "Available commands:"
echo "  make env           - Setup environment (copy .env-sample to .env)"
echo "  make up            - Start containers and run migrations"
echo "  make down          - Stop and remove containers"
echo "  make logs          - View container logs"
echo "  make ps            - Show container status"
echo ""
echo "Access the application:"
echo "  GraphQL Playground: http://localhost:8080"
echo "  phpMyAdmin:         http://localhost:8081"
echo ""
print_info "To apply group changes immediately, run:"
echo "  newgrp docker"
echo ""
