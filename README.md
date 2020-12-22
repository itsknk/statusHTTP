# statusHTTP
[![Go Reference](https://pkg.go.dev/badge/github.com/itsknk/statusHTTP.svg)](https://pkg.go.dev/github.com/itsknk/statusHTTP)

Generate HTTP Errors with ease.

statusHTTP is a library written in go that helps in generating HTTP errors in a structured manner.

## Install
If you don't have go installed, you can download it [here](https://golang.org/doc/install).
```
$ go get github.com/itsknk/statusHTTP
``` 
After installing you can import it just like everyother library.
```
import (
    "github.com/itsknk/statusHTTP"
)
```
## Usage
```
package main
import (
    "github.com/itsknk/statusHTTP"
	"net/http"
)
func myHandler(w http.ResponseWriter, r *http.Request) {
	statusHTTP.Forbidden(w, "Forbidden link")
}
func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
```

## Further Updates
1. Have to add functions for handling more errors.

## Contributing
- Fork it and then do the changes or else download the zip file, test to make sure nothing is going sideways.
- Make a pull request with a detailed explanation. 

## License
[MIT](https://github.com/itsknk/statusHTTP/blob/master/LICENSE)