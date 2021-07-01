package main

import (
	"flag"

	"bitbucket.org/proflead/cdn/configs"
	cdn_log_parser "bitbucket.org/proflead/cdn/internal"

	"github.com/burntSushi/toml"
)

var (
	log_initial_file, log_one_file, log_two_file string
	configPath                                   string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := configs.NewConfig()
	if _, err := toml.DecodeFile(configPath, cfg); err != nil {
		panic("Config fiel doesn't read")
	}

	cdn_log_parser.StartParse(cfg.LogPath, cfg.LogLimit)
}
