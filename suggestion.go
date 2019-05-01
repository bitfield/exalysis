package exalysis

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"regexp"
	"time"

	"github.com/exercism/exalysis/exam"
	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/exercism/exalysis/track/diffsquares"
	"github.com/exercism/exalysis/track/isogram"
	"github.com/exercism/exalysis/track/luhn"
	"github.com/exercism/exalysis/track/paraletterfreq"
	"github.com/exercism/exalysis/track/raindrops"
	"github.com/exercism/exalysis/track/scrabble"
	"github.com/exercism/go-analyzer/analyzer"
	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/astrav"
)

var exercisePkgs = map[string]extypes.SuggestionFunc{
	"twofer":      nil,
	"hamming":     nil,
	"raindrops":   raindrops.Suggest,
	"scrabble":    scrabble.Suggest,
	"isogram":     isogram.Suggest,
	"diffsquares": diffsquares.Suggest,
	"luhn":        luhn.Suggest,
	"letter":      paraletterfreq.Suggest,
}

// GetSuggestions selects the package suggestion routine and returns the suggestions
func GetSuggestions(codePath string) (string, string) {
	var r = extypes.NewResponse()
	folder := astrav.NewFolder(http.Dir(codePath), "")
	_, err := folder.ParseFolder()
	if err != nil {
		addGreeting(r, "", "there")
		r.AppendTodoTpl(gtpl.Compile)
		return r.GetAnswerString(), rating(r, nil, "")
	}

	// TODO: clean this up after all exercises are switched to the analyzer

	var pkgName = folder.Pkg.Name()
	analyze := pkgName == "twofer" || pkgName == "hamming"

	pkg, suggFunc := getExercisePkg(folder)
	var msg string
	if pkg == nil && !analyze {
		msg = "I don't have any specific knowledge on the %q exercise, but I'll still run my general code checks on it:"
	} else {
		msg = "I know the %q exercise, so I'll try to make some specific suggestions about it, as well as the general code checks:"
	}
	fmt.Println(aurora.Sprintf(aurora.Gray(msg), aurora.Green(pkgName)))

	addGreeting(r, pkgName, getStudentName(codePath))
	examRes, err := exam.All(folder, r, pkgName, analyze)
	if err != nil {
		log.Fatal(err)
	}

	if analyze {
		getSuggsAnalyzer(r, codePath)
	} else {
		getSuggestions(pkg, r, suggFunc)
	}
	addTip(r, pkgName)
	return r.GetAnswerString(), rating(r, examRes, pkgName)
}

func getSuggsAnalyzer(r *extypes.Response, codePath string) {
	res := analyzer.Analyze("", codePath)
	for _, err := range res.Errors {
		log.Printf("ERROR on %s:\n", codePath)
		log.Println(err)
	}

	for _, cmt := range res.Comments {
		switch cmt.Category() {
		case sugg.CtgTodo:
			r.AppendTodo(cmt)
		case sugg.CtgImprovement:
			r.AppendImprovement(cmt)
		case sugg.CtgThought:
			r.AppendComment(cmt)
		case sugg.CtgBlock:
			r.AppendBlock(cmt)
		default:
			log.Println("unknown comment category encountered")
			r.AppendImprovement(cmt)
		}
	}
}

func getSuggestions(pkg *astrav.Package, r *extypes.Response, suggFunc extypes.SuggestionFunc) {
	if suggFunc != nil {
		suggFunc(pkg, r)
	}
}

var student = regexp.MustCompile("users/([^/]*)/go/")

func getStudentName(codePath string) string {
	absPath, err := filepath.Abs(codePath)
	if err != nil {
		return ""
	}

	submatch := student.FindStringSubmatch(absPath)
	if 1 < len(submatch) {
		return submatch[1]
	}
	return ""
}

func getExercisePkg(folder *astrav.Folder) (*astrav.Package, extypes.SuggestionFunc) {
	for name, pkg := range folder.Pkgs {
		if sg, ok := exercisePkgs[name]; ok {
			return pkg, sg
		}
	}
	return nil, nil
}

func addGreeting(r *extypes.Response, pkg, student string) {
	r.SetGreeting(gtpl.Greeting.Format(student))
	if pkg == "twofer" {
		r.AppendGreeting(gtpl.NewcomerGreeting)
	}
}

func addTip(r *extypes.Response, pkgName string) {
	if pkgName == "twofer" {
		// For the first exercise, give some useful hints about Exercism and the Go track.
		r.AppendOutro(gtpl.Hints)
		return
	}
	if pkgName == "account" {
		// For the final core exercise, congratulate the student on graduating.
		r.AppendOutro(gtpl.Congrats)
		return
	}
	// For other exercises, give a randomly-selected tip.
	if r.LenSuggestions() < 3 {
		rand.Seed(time.Now().UnixNano())
		t := rand.Intn(len(gtpl.Tips))
		r.AppendTip(gtpl.Tips[t])
	}
}

func rating(r *extypes.Response, examRes *exam.Result, pkgName string) string {
	rating := aurora.Gray("Rating Suggestion\n").String()
	rating += fmt.Sprintf("Todos:\t\t%d\n", aurora.Red(r.LenTodos()))
	rating += fmt.Sprintf("Suggestions:\t%d\n", aurora.Brown(r.LenImprovements()))
	rating += fmt.Sprintf("Comments:\t%d\n", aurora.Green(r.LenComments()))

	approve := approval(r, examRes, pkgName)
	rating += fmt.Sprintf("Suggestion:\t%s\n", approve)
	return rating
}

func approval(r *extypes.Response, examRes *exam.Result, pkgName string) aurora.Value {
	var gofmt, golint bool
	if examRes != nil {
		gofmt = examRes.GoFmt
		golint = examRes.GoLint
	}

	// don't be so strict on the first exercises
	switch pkgName {
	case "twofer":
		gofmt = true
		golint = true
	case "hamming":
		golint = true
	}

	if !gofmt {
		return aurora.Red("NO APPROVAL")
	}
	if !golint {
		return aurora.Red("NO APPROVAL")
	}
	if r.LenTodos() != 0 {
		return aurora.Red("NO APPROVAL")
	}

	l := r.LenImprovements()
	switch {
	case 5 < l:
		return aurora.Red("NO APPROVAL")
	case 2 < l:
		return aurora.Magenta("MAYBE APPROVE")
	case 1 < l:
		return aurora.Brown("LIKELY APPROVE")
	}

	return aurora.Green("APPROVE")
}
