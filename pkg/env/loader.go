package env

import (
		"github.com/go-kratos/kratos/v2/log"
		"github.com/joho/godotenv"
		"os"
		"path/filepath"
)

func LoadFile(file ...string) bool {
		if err := godotenv.Load(abs(file)...); err != nil {
				log.Errorw("load env Error: %v", err)
				return false
		}
		return true
}

func abs(files []string) []string {
		if len(files) <= 0 {
				return files
		}
		for i, fs := range files {
				if filepath.IsAbs(fs) {
						continue
				}
				if absFs, err := filepath.Abs(fs); err == nil {
						files[i] = absFs
				}
		}
		return files
}

func AutoLoad() bool {
		pwd, _ := os.Getwd()
		envFile := filepath.Join(pwd, "/.env")
		if err := godotenv.Load(envFile); err != nil {
				return false
		}
		return true
}

