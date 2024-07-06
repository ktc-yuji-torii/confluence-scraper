# Confluence Scraper

Confluence Scraper は、Confluence のページデータを取得し、構造化された形式で保存するツールです。このプロジェクトは、Confluence の REST API を使用してデータを取得し、JSON 形式で保存します。

## 目次

-   [インストール](#インストール)
-   [設定](#設定)
-   [使い方](#使い方)
-   [コントリビューション](#コントリビューション)
-   [ライセンス](#ライセンス)

## インストール

Confluence Scraper を使用するには、システムに Go がインストールされている必要があります。Go のインストール方法については、公式サイト [golang.org](https://golang.org/) を参照してください。

1. リポジトリをクローンします:

    ```sh
    git clone https://github.com/your-username/confluence-scraper.git
    cd confluence-scraper
    ```

2. 依存関係をインストールします:

    ```sh
    go mod tidy
    ```

3. プロジェクトをビルドします:

    ```sh
    go build -o confluence-scraper main.go
    ```

## 設定

コマンドライン引数を通じて設定パラメータを提供する必要があります。必要なパラメータは以下の通りです：

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

### VS Code でのデバッグ

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
