package internal

import "github.com/busy-cloud/boat/config"

const MODULE = "influxdb"

func init() {
	config.SetDefault(MODULE, "url", "http://127.0.0.1:8086")
	config.SetDefault(MODULE, "org", "")
	config.SetDefault(MODULE, "bucket", "")
	config.SetDefault(MODULE, "token", "")
}
