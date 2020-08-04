module dashboard

go 1.14

require (
	addons/news v0.0.0-00010101000000-000000000000
	addons/newsreader v0.0.1
	addons/stocks v0.0.1
	addons/weather v0.0.1
	github.com/mum4k/termdash v0.12.1
	github.com/spf13/viper v1.7.0
	gopkg.in/yaml.v2 v2.2.4
)

replace addons/stocks => ./addons/stocks

replace addons/weather => ./addons/weather

replace addons/newsreader => ./addons/newsreader

replace addons/news => ./addons/news
