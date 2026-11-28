package main

import (
	"flag"
	"sort"

	"github.com/feodorpol2006-crypto/task-3/internal/config"
	"github.com/feodorpol2006-crypto/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	xmlData, err := currency.ParseXML(cfg.InputPath)
	if err != nil {
		panic(err)
	}

	convertedData, err := currency.ConvertData(xmlData.Items)
	if err != nil {
		panic(err)
	}

	sort.Slice(convertedData, func(i, j int) bool {
		return convertedData[i].Value > convertedData[j].Value
	})

	if err := currency.SaveToJSON(cfg.OutputPath, convertedData); err != nil {
		panic(err)
	}
}
