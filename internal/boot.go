package internal

import (
	"github.com/busy-cloud/boat/boot"
)

func init() {
	boot.Register("influxdb", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"log", "mqtt", "database"},
	})
}
