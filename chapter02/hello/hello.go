package hello

import "fmt"
import "strings"

func Say(names []string) string {
	if len(names) == 0 {
		return fmt.Sprint("Hello World!")
	}
	return fmt.Sprintf("Hello %s!", strings.Join(names, ", "))
}