module dashboard

go 1.14

require addons/stocks v0.0.1

require (
	addons/weather v0.0.1
	github.com/spf13/viper v1.7.0
	gopkg.in/yaml.v2 v2.2.4
)

replace addons/stocks => ./addons/stocks

replace addons/weather => ./addons/weather
