package rules

import (
	"github.com/letsssgooo/custom-linter/internal/config"
	"github.com/letsssgooo/custom-linter/internal/rules/english"
	"github.com/letsssgooo/custom-linter/internal/rules/lowercase"
	"github.com/letsssgooo/custom-linter/internal/rules/sensitive"
	"github.com/letsssgooo/custom-linter/internal/rules/symbols"
)

type Result struct {
	Reports []string
	Fixed   string
}

func CheckAll(cfg config.Config, msg string) Result {
	var checks []func(string) (bool, string, string)
	if cfg.Lowercase {
		checks = append(checks, lowercase.Check)
	}
	if cfg.English {
		checks = append(checks, english.Check)
	}
	if cfg.Symbols {
		checks = append(checks, symbols.Check)
	}
	if cfg.Sensitive {
		checks = append(checks, sensitive.Check)
	}

	out := make([]string, 0, len(checks))
	fixed := msg
	for _, check := range checks {
		ok, report, next := check(fixed)
		if !ok {
			out = append(out, report)
		}
		fixed = next
	}

	return Result{Reports: out, Fixed: fixed}
}
