package influx2

import "github.com/busy-cloud/boat/config"

const MODULE = "influx"

func init() {
	config.Register(MODULE, "url", "http://127.0.0.1:8086")
	config.Register(MODULE, "org", "")
	config.Register(MODULE, "bucket", "")
	config.Register(MODULE, "token", "")
}
