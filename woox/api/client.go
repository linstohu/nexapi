package api

import (
	"log"

	validator "github.com/go-playground/validator/v10"
)

type WooXClient struct {
	basePath    string
	key, secret string
	// debug mode
	debug bool
	// logger
	logger *log.Logger
	// validate struct fields
	validate *validator.Validate
}

type WooXCfg struct {
	BasePath string `validate:"required"`
	Key      string
	Secret   string
	Debug    bool
	// Logger
	Logger *log.Logger
}

func NewWooXClient(cfg *WooXCfg) (*WooXClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := WooXClient{
		basePath: cfg.BasePath,
		key:      cfg.Key,
		secret:   cfg.Secret,
		debug:    cfg.Debug,
		logger:   cfg.Logger,

		validate: validator,
	}

	if cli.logger == nil {
		cli.logger = log.Default()
	}

	return &cli, nil
}
