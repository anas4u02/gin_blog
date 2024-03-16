package routing

import (
	"blog/pkg/config"
	"fmt"
)

func Serve() {
	configs := config.Get()
	router := GetRouter()

	err := router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
	if err != nil {
		return
	}
}
