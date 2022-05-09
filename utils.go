package csidb

func NewRow(columns []string, values []string) Row {
	var entries []Entry
	for i, c := range columns {
		entries = append(entries, Entry{Column: c, Value: values[i]})
	}
	return Row{Entries: entries}
}
