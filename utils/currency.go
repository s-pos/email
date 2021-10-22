package utils

import (
	"fmt"
	"strings"

	"github.com/dustin/go-humanize"
)

func FormatIndonesiaCurrency(m int64) string {
	return fmt.Sprintf("Rp %s", strings.ReplaceAll(humanize.Comma(m), ",", "."))
}
