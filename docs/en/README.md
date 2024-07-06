# Golang DevContainers

Golang DevContainers is a project that provides a development environment for Go using Visual Studio Code Dev Containers. This setup aims to simplify the development process by offering a consistent environment with all necessary tools and dependencies pre-installed.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This repository contains the configuration files and setup instructions to create a fully-featured development environment for Go using Visual Studio Code's Dev Containers. Dev Containers allow you to open a project in a containerized environment, ensuring consistency across different development setups.

## Features

- Pre-configured development environment for Go.
- Integration with Visual Studio Code.
- Automatic installation of necessary tools and dependencies.
- Supports both x86_64 and ARM64 architectures.

## Prerequisites

Before using this project, ensure you have the following installed on your system:

- [Docker](https://www.docker.com/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [Remote - Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Installation

Follow these steps to set up the development environment:

1. Clone the repository:

    ```sh
    git clone https://github.com/ktc-yuji-torii/golang-devcontainers.git
    cd golang-devcontainers
    ```

2. Open the repository in Visual Studio Code:

    ```sh
    code .
    ```

3. Reopen the folder in the container:
    - Press `F1` to open the command palette.
    - Type `Remote-Containers: Reopen in Container` and select it.

VS Code will build the container image and start a development container with the pre-configured environment.

## Usage

Once the development container is running, you can start developing your Go applications. The container comes with all necessary tools and dependencies pre-installed, including:

- Go
- Common Go tools (e.g., `gofmt`, `goimports`)
- VS Code Go extension

### Building Your Project

To build your Go project inside the container, use the following command:

```sh
go build -o your-output-binary
```

### Running Tests

To run tests, use the following command:

```sh
go test ./...
```

## Contributing

We welcome contributions to improve Golang Devcontainers. To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature-name`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature/your-feature-name`).
6. Open a pull request.

Please ensure your code adheres to the project's coding standards and includes appropriate tests.
