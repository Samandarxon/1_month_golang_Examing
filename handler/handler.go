package handler

import (
	"playground/newProject/config"
	"playground/newProject/storage"
)

type handler struct {
	strg storage.StorageI
	cfg  config.Config
}

func NewHandler(strg storage.StorageI, conf config.Config) *handler {
	return &handler{strg: strg, cfg: conf}
}
