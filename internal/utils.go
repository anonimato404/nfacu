package internal

import "fmt"

func ParseLine(key, val string) string {
	return fmt.Sprintf(`    <add key="%s" value="%s" />`, key, val)
}
