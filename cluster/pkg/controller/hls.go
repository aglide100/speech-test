package controller

import (
	"github.com/aglide100/speech-test/cluster/pkg/db"
	"github.com/patrickmn/go-cache"
)

type HlsController struct {
	db *db.Database
	c *cache.Cache
}

func NewHlsController(db *db.Database, c *cache.Cache) *HlsController {
	return &HlsController{
		db : db,
		c : c,
	}
}
