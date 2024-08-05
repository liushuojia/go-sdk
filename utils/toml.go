package utils

import "github.com/BurntSushi/toml"

// type Config struct {
// 	Name  string `toml:"name"`
// 	Port  int    `toml:"port"`
// 	Debug bool   `toml:"debug"`
// 	Token struct {
// 		Access  Token
// 		Refresh Token
// 	} `toml:"token"`
// 	MysqlMap map[string]Mysql `toml:"mysql"`
// }
// type Token struct {
// 	Key    string `toml:"key"`
// 	Expire string `toml:"expire"`
// 	Issuer string `toml:"issuer"` // textproto or json
// }
// type Mysql struct {
// 	Address  string `toml:"address"`
// 	Username string `toml:"username"`
// 	Password string `toml:"password"`
// 	Database string `toml:"database"`
// }

func Read(filePath string, config interface{}) error {
	// 配置文件路径路径、解析的结构体
	if _, err := toml.DecodeFile(filePath, config); err != nil {
		return err
	}
	return nil
}
