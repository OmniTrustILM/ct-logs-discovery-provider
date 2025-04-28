package db

import (
	"CZERTAINLY-CT-Logs-Discovery-Provider/internal/config"
	"github.com/lib/pq"
)

func tbl(name string) string {
	return pq.QuoteIdentifier(config.Get().Database.Schema) +
		"." + pq.QuoteIdentifier(name)
}
