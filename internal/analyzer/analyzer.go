package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"log/slog"
	"os"
	"strconv"

	"github.com/letsssgooo/custom-linter/internal/config"
	"github.com/letsssgooo/custom-linter/internal/rules"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "customlinter",
	Doc:  "Linter for checking log notes",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	cfg, err := config.Load(os.Getenv("LINTER_CONFIG"))
	slog.Warn("error while loading config using default values", "error", err)

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			_, msg, ok := checkLogCall(pass, call)
			if !ok {
				return true
			}

			results := rules.CheckAll(cfg, msg)
			for _, report := range results.Reports {
				pass.Report(analysis.Diagnostic{
					Pos:     call.Pos(),
					End:     call.End(),
					Message: report,
					SuggestedFixes: []analysis.SuggestedFix{
						{
							TextEdits: []analysis.TextEdit{
								{
									Pos:     call.Args[0].Pos(),
									End:     call.Args[0].End(),
									NewText: []byte(strconv.Quote(results.Fixed)),
								},
							},
						},
					},
				})
			}

			return true
		})
	}

	return nil, nil
}

func checkLogCall(pass *analysis.Pass, call *ast.CallExpr) (kind, msg string, ok bool) {
	selectorExpr, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	methodName := selectorExpr.Sel.Name
	if !isLogMethod(methodName) {
		return "", "", false
	}

	if ident, ok := selectorExpr.X.(*ast.Ident); ok {
		if pkgName, ok := pass.TypesInfo.Uses[ident].(*types.PkgName); ok {
			switch pkgName.Imported().Path() {
			case "log":
				msg, ok = extractLogMessage(call)
				if ok {
					return "log", msg, true
				}
			case "log/slog":
				msg, ok = extractLogMessage(call)
				if ok {
					return "slog", msg, true
				}
			}
		}
	}

	if isZapSelector(pass, selectorExpr) {
		msg, ok = extractLogMessage(call)
		if ok {
			return "zap", msg, true
		}
	}

	return "", "", false
}

func isZapSelector(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	selection := pass.TypesInfo.Selections[selector]
	if selection == nil {
		return false
	}

	recv := selection.Recv()
	named, ok := getNamed(recv)
	if !ok || named.Obj() == nil || named.Obj().Pkg() == nil {
		return false
	}
	pkgPath := named.Obj().Pkg().Path()
	typeName := named.Obj().Name()

	if pkgPath != "go.uber.org/zap" {
		return false
	}
	return typeName == "Logger" || typeName == "SugaredLogger"
}

func getNamed(t types.Type) (*types.Named, bool) {
	if p, ok := t.(*types.Pointer); ok {
		t = p.Elem()
	}
	n, ok := t.(*types.Named)
	return n, ok
}

func extractLogMessage(call *ast.CallExpr) (string, bool) {
	if len(call.Args) == 0 {
		return "", false
	}

	bl, ok := call.Args[0].(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return "", false
	}
	unquote, err := strconv.Unquote(bl.Value)
	if err != nil {
		return "", false
	}

	return unquote, true
}

func isLogMethod(name string) bool {
	levels := []string{
		"Trace", "Debug", "Info", "Warn", "Warning", "Error", "Fatal", "Panic",
		"Tracef", "Debugf", "Infof", "Warnf", "Warningf", "Errorf", "Fatalf", "Panicf",
		"Debugw", "Infow", "Warnw", "Errorw", "Fatalw", "Panicw",
		"Log", "Printf", "Logf", "Logw", "Printw", "Print", "Println",
		"Fatalln", "Panicln", "Infoln", "Warnln", "Errorln", "Traceln", "Debugln",
	}

	for _, level := range levels {
		if name == level {
			return true
		}
	}
	return false
}
