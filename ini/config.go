package ini

import (
	"fmt"
	"log"

	"app/global"

	"github.com/spf13/viper"
)

const ConfigFile = "./config/dev.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(ConfigFile)

	if err := v.ReadInConfig(); err != nil {
		log.Printf("err:%s\n", err)
	}
	// 注意：&global后不能直接加结构体,cause 结构体是个type
	if err := v.Unmarshal(&global.Config); err != nil {
		log.Printf("err:%s\n", err)
	}

	fmt.Printf("%T   %v ---", global.Config, global.Config)
}
