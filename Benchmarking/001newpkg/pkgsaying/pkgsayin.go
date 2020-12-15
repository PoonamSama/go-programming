//Package saying welcomes you to golang
package saying

import "fmt"

//Welcome function greets you
func Welcome(s string) string {

	return fmt.Sprint("hello dear", s)
}
