package scrabble

import (
	"go/ast"
	"reflect"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/exercism/exalysis/track/scrabble/tpl"
	"github.com/tehsphinx/astrav"
)

// Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	testGoRoutine,
	testRegex,
	testMapInFunc,
	testFlattenMap,
	testMapRuneInt,
	testSliceRuneConv,
	testTrySwitch,
	testIfElseToSwitch,
	testRuneLoop,
}

func testSliceRuneConv(pkg *astrav.Package, r *extypes.Response) {
	calls := pkg.FindByNodeType(astrav.NodeTypeCallExpr)
	for _, call := range calls {
		if reflect.TypeOf(call.(*astrav.CallExpr).Fun).String() != "*ast.ArrayType" {
			continue
		}

		if call.(*astrav.CallExpr).NodeName() == "[]rune" {
			r.AppendImprovementTpl(tpl.SliceRuneConv)
		}
	}
}

func testRegex(pkg *astrav.Package, r *extypes.Response) {
	rgx := pkg.FindFirstByName("Score").FindFirstByName("regexp.MustCompile")
	if rgx == nil {
		rgx = pkg.FindFirstByName("Score").FindFirstByName("regexp.Compile")
	}

	if rgx != nil &&
		rgx.(*astrav.SelectorExpr).X != nil &&
		rgx.(*astrav.SelectorExpr).X.(*ast.Ident).Name == "regexp" {

		lit := rgx.Parent().ChildByNodeType(astrav.NodeTypeBasicLit)
		if lit != nil && lit.IsValueType("string") {
			// is a static regex
			r.AppendTodoTpl(tpl.Regex)
			r.AppendCommentTpl(tpl.Challenge)
			r.AppendOutro(gtpl.Regex)
		}
	}
}

func testFlattenMap(pkg *astrav.Package, r *extypes.Response) {
	entryFn := pkg.FindFirstByName("Score")
	loopCount := len(entryFn.FindNodeTypeInCallTree(astrav.NodeTypeForStmt))
	loopCount += len(entryFn.FindNodeTypeInCallTree(astrav.NodeTypeRangeStmt))
	if 1 < loopCount {
		r.AppendImprovementTpl(tpl.FlattenMap)
	}
}

func testMapRuneInt(pkg *astrav.Package, r *extypes.Response) {
	if r.HasSuggestion(tpl.FlattenMap) {
		return
	}
	if len(pkg.FindByValueType("map[string]int")) != 0 {
		r.AppendImprovementTpl(tpl.MapRune.Format("map[string]int"))
		return
	}
	if len(pkg.FindByValueType("map[int]string")) != 0 {
		r.AppendImprovementTpl(tpl.MapRune.Format("map[int]string"))
		return
	}
}

func testTrySwitch(pkg *astrav.Package, r *extypes.Response) {
	if r.HasSuggestion(tpl.FlattenMap) {
		return
	}
	if len(pkg.FindByValueType("map[rune]int")) != 0 {
		r.AppendCommentTpl(tpl.TrySwitch)
		return
	}
	if len(pkg.FindByValueType("map[string]int")) != 0 {
		r.AppendCommentTpl(tpl.TrySwitch)
		return
	}
	if len(pkg.FindByValueType("map[int]string")) != 0 {
		r.AppendCommentTpl(tpl.TrySwitch)
		return
	}
}

func testRuneLoop(pkg *astrav.Package, r *extypes.Response) {
	ranges := pkg.FindFirstByName("Score").FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, rng := range ranges {
		l := rng.(*astrav.RangeStmt)
		if l.Value() == nil {
			if l.Key() != nil {
				r.AppendImprovementTpl(tpl.LoopRuneNotByte)
			}
		} else {
			var isByte bool
			for _, ident := range rng.FindIdentByName(l.Value().(*astrav.Ident).Name) {
				if ident.IsValueType("byte") {
					isByte = true
				}
			}
			if isByte {
				r.AppendImprovementTpl(tpl.LoopRuneNotByte)
			}
		}

		if rng.FindByName("string") != nil &&
			!r.HasSuggestion(tpl.MapRune) &&
			!r.HasSuggestion(tpl.FlattenMap) {

			r.AppendImprovementTpl(tpl.TypeConversion)
			return
		}
	}
}

func testGoRoutine(pkg *astrav.Package, r *extypes.Response) {
	goStmts := pkg.FindByNodeType(astrav.NodeTypeGoStmt)
	if len(goStmts) != 0 {
		r.AppendTodoTpl(tpl.GoRoutines)
	}
}

func testIfElseToSwitch(pkg *astrav.Package, r *extypes.Response) {
	ranges := pkg.FindFirstByName("Score").FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, rng := range ranges {
		ifs := rng.FindByNodeType(astrav.NodeTypeIfStmt)
		if 5 < len(ifs) {
			r.AppendTodoTpl(tpl.IfsToSwitch)
			return
		}
	}
}

func testMapInFunc(pkg *astrav.Package, r *extypes.Response) {
	fn := pkg.FindFirstByName("Score")
	maps := fn.FindMaps()

	var hasMapDef bool
	for _, m := range maps {
		if !m.IsNodeType(astrav.NodeTypeIdent) {
			hasMapDef = true
		}
	}
	if hasMapDef {
		r.AppendTodoTpl(tpl.MoveMap)
	}
}
