package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	codeTmpl = `package main

import(
	
	_ "embed"
)

var (
	//go:embed inputs/day_{{.Day}}.txt
	File string	
)

func tmp() int {
	return 0
}`

	testTmpl = `
func Test{{ .FuncName1 | Title }}(t *testing.T) {
	want := {{ .Answer1 }}
	if got := {{ .FuncName1 }}(); got != want {
		t.Errorf("{{ .FuncName1 }}() = %d, want = %d", got, want)
	}
}

func Test{{ .FuncName2 | Title }}(t *testing.T) {
	want := {{ .Answer2 }}
	if got := {{ .FuncName2 }}(); got != want {
		t.Errorf("{{ .FuncName2 }}() = %d, want = %d", got, want)
	}
}
`

	benchmarkTmpl = `
func Benchmark{{ .FuncName1 | Title }}(b *testing.B) {
	for i := 0; i < b.N; i++ {
		{{ .FuncName1 }}()
	}
}

func Benchmark{{ .FuncName2 | Title }}(b *testing.B) {
	for i := 0; i < b.N; i++ {
		{{ .FuncName2 }}()
	}
}
`
)

type FuncsSpecs struct {
	FuncName1 string
	FuncName2 string
	Answer1   int
	Answer2   int
}

// Populate the day's file with templated code.
func MakeCodeTmpl(d int) {
	inF, err := os.Create(fmt.Sprintf("inputs/day_%d.txt", d))
	if err != nil {
		log.Fatal(err)
	}
	inF.Close()

	x := struct{ Day int }{d}
	tmpl, err := template.New("tmp").Parse(codeTmpl)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(fmt.Sprintf("day_%d.go", d))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	if err := tmpl.Execute(w, x); err != nil {
		log.Fatal(err)
	}
	w.Flush()
}

func AddTestsAndBenchmarks(funcName1, funcName2 string, answer1, answer2 int) {
	x := &FuncsSpecs{
		funcName1,
		funcName2,
		answer1,
		answer2,
	}
	funcMap := template.FuncMap{
		"Title": strings.Title,
	}
	tests, err := template.New("tests").Funcs(funcMap).Parse(testTmpl)
	if err != nil {
		log.Fatal(err)
	}
	WriteTmpl(x, tests, "correctness_test.go")

	bm, err := template.New("benchmarks").Funcs(funcMap).Parse(benchmarkTmpl)
	if err != nil {
		log.Fatal(err)
	}
	WriteTmpl(x, bm, "benchmark_test.go")
}

func WriteTmpl(x *FuncsSpecs, tmpl *template.Template, fname string) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	if err := tmpl.Execute(w, x); err != nil {
		log.Fatal(err)
	}
	w.Flush()
}
