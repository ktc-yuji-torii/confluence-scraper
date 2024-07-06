# Golang DevContainers

Golang DevContainersは、Visual Studio Code Dev Containersを使用してGoの開発環境を提供するプロジェクトです。このセットアップは、必要なツールや依存関係がすべて事前にインストールされた一貫性のある環境を提供することで、開発プロセスを簡素化することを目的としています。

## 目次

- [紹介](#紹介)
- [特徴](#特徴)
- [前提条件](#前提条件)
- [インストール](#インストール)
- [使い方](#使い方)
- [コントリビューション](#コントリビューション)
- [ライセンス](#ライセンス)

## 紹介

このリポジトリには、Visual Studio CodeのDev Containersを使用して完全な機能を持つGoの開発環境を作成するための設定ファイルとセットアップ手順が含まれています。Dev Containersを使用すると、異なる開発環境間で一貫性を保ちながらプロジェクトを開くことができます。

## 特徴

- Goのための事前に設定された開発環境。
- Visual Studio Codeとの統合。
- 必要なツールや依存関係の自動インストール。
- x86_64およびARM64アーキテクチャをサポート。

## 前提条件

このプロジェクトを使用する前に、システムに以下がインストールされていることを確認してください：

- [Docker](https://www.docker.com/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [Remote - Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## インストール

開発環境をセットアップするには、以下の手順に従ってください：

1. リポジトリをクローンします：

    ```sh
    git clone https://github.com/ktc-yuji-torii/golang-devcontainers.git
    cd golang-devcontainers
    ```

2. リポジトリをVisual Studio Codeで開きます：

    ```sh
    code .
    ```

3. フォルダーをコンテナで再オープンします：
    - `F1`キーを押してコマンドパレットを開きます。
    - `Remote-Containers: Reopen in Container`と入力し、選択します。

VS Codeはコンテナイメージをビルドし、事前に設定された環境で開発コンテナを開始します。

## 使い方

開発コンテナが実行されると、Goアプリケーションの開発を開始できます。コンテナには、以下を含む必要なツールや依存関係がすべて事前にインストールされています：

- Go
- 一般的なGoツール（例：`gofmt`、`goimports`）
- VS Code Go拡張機能

### プロジェクトのビルド

コンテナ内でGoプロジェクトをビルドするには、以下のコマンドを使用します：

```sh
go build -o your-output-binary
```

### テストの実行

テストを実行するには、以下のコマンドを使用します：

```sh
go test ./...
```

## コントリビューション

Golang DevContainersの改善にご協力ください。コントリビューションの手順は以下の通りです：

1. リポジトリをフォークします。
2. 新しいブランチを作成します (`git checkout -b feature/your-feature-name`)。
3. 変更を加えます。
4. 変更をコミットします (`git commit -m 'Add some feature'`)。
5. ブランチにプッシュします (`git push origin feature/your-feature-name`)。
6. プルリクエストを開きます。

コードがプロジェクトのコーディング標準に従っていること、および適切なテストが含まれていることを確認してください。
