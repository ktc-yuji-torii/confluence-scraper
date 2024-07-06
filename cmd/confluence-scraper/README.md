# Confluence Scraper

Confluence Scraper は、Confluence のページデータを取得し、構造化された形式で保存する強力なツールです。このプロジェクトは、Confluence の REST API を活用してデータを包括的に抽出し、その情報を JSON 形式で保存します。

## 目次

-   [インストール](#インストール)
-   [設定](#設定)
-   [使い方](#使い方)
-   [ビルド手順](#ビルド手順)
-   [VS Code でのデバッグ](#vs-codeでのデバッグ)
-   [コントリビューション](#コントリビューション)
-   [ライセンス](#ライセンス)

## インストール

Confluence Scraper を効果的に利用するには、システムに Go がインストールされている必要があります。Go のインストール方法については、公式サイト [golang.org](https://golang.org/) を参照してください。

1. リポジトリをクローンします:

    ```sh
    git clone https://github.com/your-username/confluence-scraper.git
    cd confluence-scraper
    ```

2. 依存関係をインストールします:

    ```sh
    go mod tidy
    ```

## 設定

設定パラメータはコマンドライン引数を通じて提供します。必要なパラメータは以下の通りです：

-   `baseURL`: Confluence インスタンスのベース URL
-   `username`: Confluence のユーザー名
-   `apiToken`: Confluence の API トークン
-   `parentPageID`: Confluence の親ページ ID
-   `debug`: デバッグモードを有効にする（オプション）

## 使い方

Confluence Scraper を実行するには、以下のコマンドを使用します：

```sh
./confluence-scraper --baseURL=https://your-confluence-instance.atlassian.net/wiki --username=your-username --apiToken=your-api-token --parentPageID=your-parent-page-id --debug=true
```

## ビルド手順

このプロジェクトには、ビルドプロセスを容易にするための`Makefile`が含まれています。`Makefile`はシステムのアーキテクチャを自動的に検出し、適切な`GOARCH`値を設定します。

1. プロジェクトディレクトリにいることを確認してください。
2. ビルドコマンドを実行します：

    ```sh
    make build
    ```

### `Makefile`の例

```makefile
.PHONY: build
build:
	@ARCH=`uname -m`; \
	echo "アーキテクチャ: $$ARCH"; \
	if [ "$$ARCH" = "x86_64" ]; then \
		GOARCH="amd64"; \
	elif [ "$$ARCH" = "aarch64" ]; then \
		GOARCH="arm64"; \
	else \
		echo "Unsupported architecture: $$ARCH"; \
		exit 1; \
	fi; \
	echo "GOARCH: $$GOARCH"; \
	CGO_ENABLED=0 GOOS=linux GOARCH=$$GOARCH go build \
		-ldflags="-w -s" \
		-trimpath \
		-o confluence-scraper \
		main.go
	@echo "ビルド完了"
```

## VS Code でのデバッグ

VS Code を使用している場合、以下の`launch.json`設定を使用してアプリケーションをデバッグできます：

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

## コントリビューション

Confluence Scraper の改善にご協力ください。コントリビューションの手順は以下の通りです：

1. リポジトリをフォークします。
2. 新しいブランチを作成します (`git checkout -b feature/your-feature-name`)。
3. 変更を加えます。
4. 変更をコミットします (`git commit -m 'Add some feature'`)。
5. ブランチにプッシュします (`git push origin feature/your-feature-name`)。
6. プルリクエストを開きます。

コードがプロジェクトのコーディング標準に従っていること、および適切なテストが含まれていることを確認してください。

## ライセンス

このプロジェクトは MIT ライセンスの下でライセンスされています。詳細については、[LICENSE](LICENSE)ファイルを参照してください。
