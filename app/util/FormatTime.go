package util

import (
	"time"
	"github.com/janaonline/icmyc/config"
)

func FormatMongoDate(t time.Time) string  {
	return t.Format(config.Get().MagicalDate)
}