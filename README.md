# goweblinks

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/goweblinks)
[![Build Status](https://travis-ci.org/shamsher31/goweblinks.svg)](https://travis-ci.org/shamsher31/goweblinks)

This will give you list of web links from given webpage.

### How to install
```go
go get github.com/shamsher31/goweblinks
```

### How to use

Create a file called links.go and save the following content.

```go
package main

import (
  "flag"
  "fmt"
  "github.com/shamsher31/goweblinks"
  "net/http"
  "os"
)

func main() {

  flag.Parse()
  args := flag.Args()

  if len(args) < 1 {
    fmt.Println("Please specify start page")
    os.Exit(1)
  }

  getWebPage(args[0])

}

func getWebPage(uri string) {
  
  resp, err := http.Get(uri)

  if err != nil {
    return
  }

  defer resp.Body.Close()

  links := weblinks.Get(resp.Body)

  for _, link := range links {
    fmt.Println(link)
  }

}

```

Run this through command line, notice providing website url

```go
go run links.go http://google.com

```
You will get the following output

```go
http://www.goole.com/
http://www.goole.com/about/
http://www.goole.com/schools-and-education/
http://www.goole.com/blogs-facebook/
http://www.goole.com/business-directory/
http://www.goole.com/charities-support-groups/
http://www.goole.com/gallery/
http://www.goole.com/history/
http://www.goole.com/news/
http://www.goole.com/sport/
http://www.jdoqocy.com/click-1948850-10857968-1353517711000
http://www.goole.com/acceptable-use/
http://www.goole.com/terms-conditions/
http://www.goole.com/privacy/

```

### Why
This package will be useful when you are writing web crawler application and you want to fetch
all the links from the provided website url.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
