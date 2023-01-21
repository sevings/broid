# broid
A browser fingerprinting library that only runs on the server. 

# Installing

```shell
go get github.com/sevings/broid
```

# Usage
Import the package into your project. Then create a new BrowserIDBuilder.
The default BrowserIDBuilder contains functions for computing a BrowserID based
on the "User-Agent", "Accept", "Accept-Encoding" and "Accept-Language" headers.
This allows you to distinguish between hundreds of browser versions. You can 
improve accuracy by adding more browser characteristics for computing using
AddField(). Call Build() to obtain the BrowserID of the browser that made
the request.

```go
package main

import (
	"fmt"
	"github.com/sevings/broid"
	"io"
	"log"
	"net/http"
)

func main() {
	builder := broid.NewDefaultBrowserIDBuilder()

	idHandler := func(w http.ResponseWriter, req *http.Request) {
		id := builder.Build(req).String()
		io.WriteString(w, fmt.Sprintf("The ID of your browser is %s.", id))
	}

	http.HandleFunc("/id", idHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```
