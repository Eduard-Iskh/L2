package match

import (
	"regexp"
	"strings"

	"example.com/internal/cfg"
)

// IsMatch проверяет, совпадает ли строка с шаблоном согласно заданной конфигурации.
// Возвращает true при совпадении и false в противном случае.
func IsMatch(line string, config cfg.Config, re *regexp.Regexp) bool {
	pattern := config.Pattern
	if config.Fixed {
		if config.IgnoreCase {
			return strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		}
		return strings.Contains(line, pattern)
	}

	return re.MatchString(line)
}
