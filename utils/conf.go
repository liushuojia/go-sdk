package utils

import "github.com/jinzhu/configor"

func Read(c interface{}, path ...string) error {
	return configor.Load(c, path...)
}
