package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/fatih/color"
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

// Print a bingo board for Day 04
func (b *board) String() string {
	rows := make([][]string, 5)

	for val, cell := range b.values {
		if len(rows[cell.row]) == 0 {
			rows[cell.row] = make([]string, 5)
		}
		rows[cell.row][cell.col] = val
	}

	buf := new(strings.Builder)
	w := tabwriter.NewWriter(buf, 0, 0, 1, ' ', tabwriter.AlignRight)
	for _, row := range rows {
		for _, cellVal := range row {
			thisCell := b.values[cellVal]
			if thisCell.marked {
				fmt.Fprintf(w, "%s\t", color.BlueString(cellVal))
			} else {
				fmt.Fprintf(w, "%s\t", color.HiWhiteString(cellVal))
			}
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	return buf.String()
}

// Populate the day's file with templated code.
func makeCodeTmpl(d int) {
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

type funcsSpecs struct {
	FuncName1 string
	FuncName2 string
	Answer1   int
	Answer2   int
}

func addTestsAndBenchmarks(funcName1, funcName2 string, answer1, answer2 int) {
	x := &funcsSpecs{
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
	writeTmpl(x, tests, "correctness_test.go")

	bm, err := template.New("benchmarks").Funcs(funcMap).Parse(benchmarkTmpl)
	if err != nil {
		log.Fatal(err)
	}
	writeTmpl(x, bm, "benchmark_test.go")
}

func writeTmpl(x *funcsSpecs, tmpl *template.Template, fname string) {
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
