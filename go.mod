module github.com/endgate-systems/lm-cli

go 1.24.1

replace github.com/endgate-systems/lm-automation => /opt/lm-cli

require (
	github.com/endgate-systems/lm-automation v0.0.0-00010101000000-000000000000
	github.com/jmoiron/sqlx v1.4.0
	github.com/knadh/listmonk v1.1.0
	github.com/spf13/cobra v1.9.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/yuin/goldmark v1.3.4 // indirect
	gopkg.in/volatiletech/null.v6 v6.0.0-20170828023728-0bef4e07ae1b // indirect
)
