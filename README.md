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

#### Installation

Like many good things in life, simply install with `go get`:

    go get github.com/hermanschaaf/execwrapper


#### Explanation

This is really simple, and using this package is totally overkill. The trick is to wrap normal exec commands with `bash -c`:

    output = exec.Command("bash", "-c", cmd)
    return output

It's possible that this will be extended later to deal with more platforms, in which case it might actually warrant using this package :)