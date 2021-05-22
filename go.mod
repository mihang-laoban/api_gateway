module api_gateway

go 1.16

require (
	dev-go-base v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
)

replace dev-go-base => ../dev-go-base

replace dev-go-base/server => ../dev-go-base
