package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/ktc-yuji-torii/confluence-scraper/config"
	"github.com/ktc-yuji-torii/confluence-scraper/internal/client"
	"github.com/ktc-yuji-torii/confluence-scraper/models"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logger *slog.Logger
	race   bool
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "confluence-scraper",
	Short: "Confluenceのページ情報を再帰的に取得するCLIツール",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(cmd)
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		// ログの初期化を行う
		var level slog.Level
		if cfg.Debug {
			level = slog.LevelDebug
		} else {
			level = slog.LevelInfo
		}
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
		logger.Info("Logger initialized", "config", struct {
			BaseURL      string `json:"baseURL"`
			Username     string `json:"username"`
			ParentPageID string `json:"pageID"`
			Debug        bool   `json:"debug"`
		}{
			BaseURL:      cfg.BaseURL,
			Username:     cfg.Username,
			ParentPageID: cfg.ParentPageID,
			Debug:        cfg.Debug,
		})

		// アプリケーションのメイン処理
		confluenceClient := client.NewConfluenceClient(*cfg, logger)
		logger.Info("Starting to fetch pages recursively", "cfg.ParentPageID", cfg.ParentPageID)

		pages, err := confluenceClient.GetChildPagesRecursively(cfg.ParentPageID, *cfg)
		if err != nil {
			logger.Error("Error fetching child pages recursively", "error", err)
			return
		}

		logger.Info("Total Pages fetched", "count", len(pages))

		// JSONファイルに保存
		fileName := generateFileName(cfg.ParentPageID)
		outputFilePath := filepath.Join("output", fileName)
		err = savePagesToFile(outputFilePath, pages)
		if err != nil {
			logger.Error("Error saving pages to file", "error", err)
			return
		}

		logger.Info("Successfully saved all pages", "fileName", fileName)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// コマンドライン引数のフラグを定義
	rootCmd.PersistentFlags().String("baseURL", "", "Base URL of the Confluence instance")
	rootCmd.PersistentFlags().String("username", "", "Username for Confluence")
	rootCmd.PersistentFlags().String("apiToken", "", "API token for Confluence")
	rootCmd.PersistentFlags().String("parentPageID", "", "Parent Page ID in Confluence")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug mode")
	rootCmd.PersistentFlags().BoolVar(&race, "race", false, "Enable race detection")

	// 必須のフラグを設定
	rootCmd.MarkPersistentFlagRequired("baseURL")
	rootCmd.MarkPersistentFlagRequired("username")
	rootCmd.MarkPersistentFlagRequired("apiToken")
	rootCmd.MarkPersistentFlagRequired("parentPageID")
}

func initConfig() {
	// 環境変数を読み込む設定
	viper.AutomaticEnv()
}

// JSONファイル名を生成する関数
func generateFileName(spaceKey string) string {
	return fmt.Sprintf("%s.json", spaceKey)
}

// ページ情報をJSONファイルに保存する関数
func savePagesToFile(filePath string, pages []models.OutputPage) error {
	// outputディレクトリを作成
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	// JSONファイルを作成
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating JSON file: %w", err)
	}
	defer file.Close()

	// JSONエンコード
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	data := map[string]interface{}{
		"pages": pages,
	}
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("error encoding data to JSON: %w", err)
	}

	return nil
}

func main() {
	// コマンドを実行
	if err := rootCmd.Execute(); err != nil {
		if logger != nil {
			logger.Error("Command execution failed", "error", err)
		} else {
			log.Fatalf("Command execution failed: %v", err)
		}
		os.Exit(1)
	}
}
