package boot

import (
	"github.com/busy-cloud/boat/boot"
	_ "github.com/busy-cloud/influxdb/apis"
	"github.com/busy-cloud/influxdb/influx2"
)

func init() {
	boot.Register("influxdb", &boot.Task{
		Startup:  influx2.Startup,
		Shutdown: nil,
		Depends:  []string{"log", "mqtt", "database"},
	})
}
