/*
	Hans Nicolaus
	29 Aug 2023
*/

package main

import (
	"os"
	"os/signal"
	"path/filepath"

	"bibit.id/challenge/handler"
	"bibit.id/challenge/model"
	"bibit.id/challenge/repo"
	"bibit.id/challenge/server"
	"bibit.id/challenge/usecase"
	"gopkg.in/yaml.v2"
)

const (
	fileName = "bibit.local.yaml"
)

func main() {
	stockRepo := repo.New()
	stockUsecase := usecase.New(stockRepo)
	stockHandler := handler.New(stockUsecase)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	cfg := getConfig()

	go server.ServeGRPC(cfg, stockHandler)
	go server.ServeKafka(cfg, stockHandler)

	<-signals
}

func getConfig() model.Config {
	var file *os.File

	cfg, err := func() (model.Config, error) {
		currentDir, err := os.Getwd()
		if err != nil {
			return model.Config{}, err
		}

		path := filepath.Join(currentDir, fileName)

		file, err = os.Open(filepath.Clean(path))
		if err != nil {
			return model.Config{}, err
		}

		var cfg model.Config
		if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
			return model.Config{}, err
		}

		return cfg, nil
	}()

	if err != nil || cfg == (model.Config{}) {
		return model.DefaultConfigLocal
	}

	_ = file.Close()
	return cfg
}
