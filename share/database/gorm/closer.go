package gorm

import "gorm.io/gorm"

type Closer struct {
	DB *gorm.DB
}

func (c *Closer) Close() error {
	raw, _ := c.DB.DB()
	return raw.Close()
}
