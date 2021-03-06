package format

import (
	"bytes"
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/pflag"
)

func render(rows [][]string) []byte {
	header := []string{"FlagSet", "Name", "Value", "Type"}

	b := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(b)
	table.SetHeader(header)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(true)
	table.SetBorder(true)
	table.SetTablePadding("\t") // pad with tabs

	for _, v := range rows {
		table.Append(v)
	}
	table.Render()

	return b.Bytes()
}

// RenderCode renders a snippet with lines numbers
func RenderCode(code string) string {
	// Go editors use tabs instead of spaces, fix them to proper printing
	code = strings.ReplaceAll(code, "\t", "    ")
	text := fmt.Sprintf("\n %02d |", 1)
	lines := strings.Split(code, "\n")
	for n, l := range lines {
		text = text + fmt.Sprintf("\n %02d |  %s", n+2, l)
	}
	return text + fmt.Sprintf("\n %02d |", len(lines)+2)
}

// RenderFlagSet renders a flag.FlagSet as a table
func RenderFlagSet(fss ...*flag.FlagSet) {
	rows := [][]string{}
	for _, fs := range fss {
		fs.VisitAll(func(f *flag.Flag) {
			rows = append(rows, []string{fs.Name(), f.Name, fmt.Sprintf("%v", f.Value), reflect.TypeOf(f.Value).String()})
		})
	}
	fmt.Printf("%s\n", render(rows))
}

// RenderPFlagSet renders a pflag.FlagSet as a table
func RenderPFlagSet(fss ...*pflag.FlagSet) {
	rows := [][]string{}
	for _, fs := range fss {
		fs.VisitAll(func(f *pflag.Flag) {
			rows = append(rows, []string{"", f.Name, fmt.Sprintf("%v", f.Value), reflect.TypeOf(f.Value).String()})
		})
	}
	fmt.Printf("%s\n", render(rows))
}
