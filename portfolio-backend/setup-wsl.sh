#!/bin/bash
# WSL2 Ubuntu Environment Setup Script for kp-backend
# Usage: bash setup-wsl.sh

set -e

echo "==================================="
echo "WSL2 Ubuntu Setup for kp-backend"
echo "==================================="
echo ""

# -----------------------------------------------
# System packages update
# -----------------------------------------------
echo "[1/7] Updating system packages..."
sudo apt-get update -y
sudo apt-get upgrade -y

# -----------------------------------------------
# Install Docker
# -----------------------------------------------
echo "[2/7] Installing Docker..."
if ! command -v docker &> /dev/null; then
  sudo apt-get install -y ca-certificates curl gnupg lsb-release
  sudo install -m 0755 -d /etc/apt/keyrings
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg \
    | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
  sudo chmod a+r /etc/apt/keyrings/docker.gpg
  echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] \
    https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) stable" \
    | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
  sudo apt-get update -y
  sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
  echo "Docker installed."
else
  echo "Docker already installed. Skipping."
fi

# -----------------------------------------------
# Install Docker Compose (standalone v2)
# -----------------------------------------------
echo "[3/7] Installing Docker Compose..."
if ! command -v docker-compose &> /dev/null; then
  COMPOSE_VERSION=$(curl -fsSL https://api.github.com/repos/docker/compose/releases/latest \
    | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')
  sudo curl -fsSL \
    "https://github.com/docker/compose/releases/download/${COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" \
    -o /usr/local/bin/docker-compose
  sudo chmod +x /usr/local/bin/docker-compose
  echo "Docker Compose ${COMPOSE_VERSION} installed."
else
  echo "Docker Compose already installed. Skipping."
fi

# -----------------------------------------------
# Install Go
# -----------------------------------------------
echo "[4/7] Installing Go..."
if ! command -v go &> /dev/null; then
  GO_VERSION="1.25.0"
  curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" -o /tmp/go.tar.gz
  sudo rm -rf /usr/local/go
  sudo tar -C /usr/local -xzf /tmp/go.tar.gz
  rm /tmp/go.tar.gz
  # Add Go to PATH for this session
  export PATH="$PATH:/usr/local/go/bin"
  # Persist to .bashrc / .profile if not already set
  if ! grep -q '/usr/local/go/bin' ~/.bashrc 2>/dev/null; then
    echo 'export PATH="$PATH:/usr/local/go/bin"' >> ~/.bashrc
  fi
  echo "Go ${GO_VERSION} installed."
else
  echo "Go $(go version) already installed. Skipping."
fi

# -----------------------------------------------
# Install Make and MySQL Client
# -----------------------------------------------
echo "[5/7] Installing Make and MySQL Client..."
sudo apt-get install -y make mysql-client

# -----------------------------------------------
# Configure Docker group
# -----------------------------------------------
echo "[6/7] Configuring Docker group..."
if ! groups "$USER" | grep -q docker; then
  sudo usermod -aG docker "$USER"
  echo "Added $USER to docker group. You may need to run 'newgrp docker' or re-login."
else
  echo "User $USER is already in the docker group."
fi

# Start Docker service
if ! sudo service docker status > /dev/null 2>&1; then
  sudo service docker start
  echo "Docker service started."
else
  echo "Docker service already running."
fi

# -----------------------------------------------
# Download Go dependencies
# -----------------------------------------------
echo "[7/7] Downloading Go dependencies..."
go mod download

# -----------------------------------------------
# Setup .env
# -----------------------------------------------
if [ ! -f .env ]; then
  cp .env.example .env
  echo ".env file created from .env.example"
fi

echo ""
echo "==================================="
echo "Setup complete!"
echo ""
echo "Next steps:"
echo "  make up    - Build and start all containers (DB + app + phpMyAdmin)"
echo "  make logs  - View container logs"
echo "  make ps    - Show container status"
echo ""
echo "Access URLs after 'make up':"
echo "  GraphQL Playground: http://localhost:8080/"
echo "  GraphQL API:        http://localhost:8080/query"
echo "  phpMyAdmin:         http://localhost:8081/"
echo "==================================="
