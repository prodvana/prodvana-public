package cmdutil

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/gosuri/uitable"
	"github.com/mattn/go-isatty"
)

var (
	OutputFormat string = "table"
)

// Output format => supports color output
var outputFormats = map[string]bool{
	"table": true,
	"json":  false,
}

func ValidOutputFormats() []string {
	r := make([]string, len(outputFormats))
	idx := 0
	for k := range outputFormats {
		r[idx] = k
		idx++
	}
	return r
}

func SetOutputFormat(format string) error {
	if _, ok := outputFormats[strings.ToLower(format)]; !ok {
		return fmt.Errorf("Unknown output format '%s'. Must be one of [%s]", format, strings.Join(ValidOutputFormats(), ", "))
	}
	OutputFormat = format
	return nil
}

func ShouldPrintColor() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) && outputFormats[OutputFormat]
}

func WriteStdout(listing *OutputListing) error {
	return WriteOutput(os.Stdout, listing)
}

func WriteOutput(writer io.Writer, listing *OutputListing) error {
	if listing == nil {
		return nil
	}
	var err error
	switch OutputFormat {
	case "json":
		err = listing.WriteJson(writer)
	default:
		err = listing.WriteTable(writer)
	}
	return err

}

type OutputListing struct {
	headers []string
	rows    [][]any
}

func NewOutputListing(headers ...string) *OutputListing {
	return &OutputListing{
		headers: headers,
		rows:    [][]any{},
	}
}

func (l *OutputListing) AddRow(values ...any) {
	if len(values) != len(l.headers) {
		panic(fmt.Sprintf("row length (%d) does not match header count (%d)", len(values), len(l.headers)))
	}
	l.rows = append(l.rows, values)
}

func toInterface(strs []string) []interface{} {
	result := make([]interface{}, len(strs))
	for idx, str := range strs {
		result[idx] = str
	}
	return result
}

func (l *OutputListing) WriteTable(output io.Writer) error {
	table := uitable.New()
	table.AddRow(toInterface(l.headers)...)
	for _, row := range l.rows {
		table.AddRow(row...)
	}
	tableBytes := table.Bytes()
	tableBytes = append(tableBytes, '\n')
	_, err := output.Write(tableBytes)
	return err
}

func (l *OutputListing) WriteJson(output io.Writer) error {
	result := make([]map[string]any, len(l.rows))
	fields := camelCase(l.headers)
	for idx := range l.rows {
		row := map[string]any{}
		for valIdx, value := range l.rows[idx] {
			row[fields[valIdx]] = value
		}
		result[idx] = row
	}
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return err
	}
	jsonBytes = append(jsonBytes, '\n')
	_, err = output.Write(jsonBytes)
	return err
}

func camelCase(headers []string) []string {
	result := make([]string, len(headers))
	for idx, header := range headers {
		// split on whitespace, removing runs between words
		fields := strings.Fields(header)
		if len(fields) == 0 {
			result[idx] = strings.ToLower(header)
			continue
		}
		builder := strings.Builder{}
		for idx, field := range fields {
			for rdx, rn := range field {
				// first character of each word should be upper case except the
				// first one
				if rdx == 0 && idx > 0 {
					builder.WriteRune(unicode.ToUpper(rn))
				} else {
					builder.WriteRune(unicode.ToLower(rn))
				}
			}
		}
		result[idx] = builder.String()
	}
	return result
}
