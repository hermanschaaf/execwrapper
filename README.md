Execwrapper
-----------

A simple wrapper for Go `exec` that allows you to execute arbitrary strings and pipe commands. For example:

```go
package main

import (
	"fmt"
	"github.com/hermanschaaf/execwrapper"
)

func main() {
	output, err := execwrapper.Command("echo \"why\nhello\nthere\" | grep \"hello\"").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output)) // prints "hello"
}

```
