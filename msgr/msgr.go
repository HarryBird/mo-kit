package msgr

import "fmt"

// W print msg like {harry} ...
func W(fname string) string {
	return fmt.Sprintf("{%s}", fname)
}
