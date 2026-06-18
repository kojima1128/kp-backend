#!/bin/bash

set -e

echo "=========================================="
echo "Portfolio Backend WSL Setup Script"
echo "=========================================="
echo ""

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_step() {
    echo -e "${BLUE}→ $1${NC}"
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

# Check if running on WSL
print_step "Checking if running on WSL..."
if ! grep -qi microsoft /proc/version; then
    print_warning "Not running on WSL. This script is designed for WSL2 Ubuntu."
    print_warning "Please run this script inside WSL Ubuntu."
    exit 1
fi
print_success "Running on WSL"
echo ""

# Update package list
print_step "Updating package list..."
sudo apt update > /dev/null 2>&1
print_success "Package list updated"
echo ""

# Install Docker
print_step "Installing Docker..."
if ! command -v docker &> /dev/null; then
    sudo apt install -y docker.io > /dev/null 2>&1
    print_success "Docker installed"
else
    print_success "Docker already installed"
fi

# Configure Docker to run without sudo
print_step "Configuring Docker to run without sudo..."
if ! groups | grep -q docker; then
    sudo usermod -aG docker $USER
    print_success "User added to docker group"
    print_warning "Please run: 'newgrp docker' to apply changes, or restart your terminal"
else
    print_success "Docker group already configured"
fi
echo ""

# Install Docker Compose
print_step "Installing Docker Compose..."
if ! command -v docker-compose &> /dev/null; then
    sudo apt install -y docker-compose > /dev/null 2>&1
    print_success "Docker Compose installed"
else
    print_success "Docker Compose already installed"
fi
echo ""

# Install Make
print_step "Installing Make..."
if ! command -v make &> /dev/null; then
    sudo apt install -y make > /dev/null 2>&1
    print_success "Make installed"
else
    print_success "Make already installed"
fi
echo ""

# Install Go
print_step "Installing Go..."
if ! command -v go &> /dev/null; then
    sudo apt install -y golang-go > /dev/null 2>&1
    print_success "Go installed"
else
    print_success "Go already installed"
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}')
echo "  Go version: $GO_VERSION"
echo ""

# Install MySQL Client (optional but useful)
print_step "Installing MySQL Client..."
if ! command -v mysql &> /dev/null; then
    sudo apt install -y mysql-client > /dev/null 2>&1
    print_success "MySQL Client installed"
else
    print_success "MySQL Client already installed"
fi
echo ""

# Navigate to portfolio-backend directory
print_step "Navigating to portfolio-backend directory..."
cd "$(dirname "$0")/portfolio-backend"
print_success "Current directory: $(pwd)"
echo ""

# Download Go dependencies
print_step "Downloading Go dependencies..."
go mod download > /dev/null 2>&1
print_success "Go dependencies downloaded"
echo ""

# Start Docker daemon if not running
print_step "Starting Docker daemon..."
if ! sudo service docker status > /dev/null 2>&1; then
    sudo service docker start > /dev/null 2>&1
    print_success "Docker daemon started"
    sleep 2
else
    print_success "Docker daemon already running"
fi
echo ""

# Build and start containers
print_step "Building and starting containers with database migration..."
make up
echo ""

print_success "=========================================="
print_success "Setup Complete!"
print_success "=========================================="
echo ""
echo -e "${GREEN}Your development environment is ready!${NC}"
echo ""
echo "Next steps:"
echo "1. Open VSCode in the portfolio-backend directory:"
echo "   code portfolio-backend"
echo ""
echo "2. Access GraphQL Playground:"
echo "   http://localhost:8080"
echo ""
echo "3. Useful commands:"
echo "   make up       - Start containers with DB migration"
echo "   make down     - Stop and remove containers (delete volumes)"
echo "   make stop     - Stop containers without removing volumes"
echo "   make logs     - View container logs"
echo "   make ps       - Show container status"
echo ""
echo "4. Database access (MySQL):"
echo "   mysql -h localhost -u user -p -D common_db"
echo "   Password: password"
echo ""
print_warning "If you added docker group, run 'newgrp docker' to apply it without restarting"
echo ""
