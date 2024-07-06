# Confluence Scraper

Confluence Scraper is a powerful tool designed to retrieve and save structured Confluence page data. This project leverages Confluence's REST API to comprehensively extract data and store it in JSON format.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Build Instructions](#build-instructions)
- [Debugging with VS Code](#debugging-with-vs-code)
- [Contributing](#contributing)
- [License](#license)

## Installation

To effectively utilize Confluence Scraper, ensure that Go is installed on your system. For Go installation instructions, refer to the official site [golang.org](https://golang.org/).

1. Clone the repository:

    ```sh
    git clone https://github.com/your-username/confluence-scraper.git
    cd confluence-scraper
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Configuration

Configuration parameters are provided via command-line arguments. The required parameters are as follows:

- `baseURL`: The base URL of your Confluence instance.
- `username`: Your Confluence username.
- `apiToken`: Your Confluence API token.
- `parentPageID`: The parent page ID in Confluence.
- `debug`: Enable debug mode (optional).

## Usage

To run Confluence Scraper, use the following command:

```sh
./confluence-scraper --baseURL=https://your-confluence-instance.atlassian.net/wiki --username=your-username --apiToken=your-api-token --parentPageID=your-parent-page-id --debug=true
```

## Build Instructions

This project includes a `Makefile` to facilitate the build process. The `Makefile` automatically detects the system architecture and sets the appropriate `GOARCH` value.

1. Ensure you are in the project directory.
2. Run the build command:

    ```sh
    make build
    ```

## Debugging with VS Code

If you are using VS Code, you can use the following `launch.json` configuration to debug the application:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Go Program",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go",
            "args": [
                "--baseURL=https://your-confluence-instance.atlassian.net/wiki",
                "--username=your-username",
                "--apiToken=your-api-token",
                "--parentPageID=your-parent-page-id",
                "--debug=true"
            ],
            "env": {},
            "cwd": "${workspaceFolder}"
        }
    ]
}
```

## Contributing

We welcome contributions to improve Confluence Scraper. To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature-name`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature/your-feature-name`).
6. Open a pull request.

Please ensure your code adheres to the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. For more details, see the [LICENSE](LICENSE) file.
