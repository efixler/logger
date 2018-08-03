# logger
Go logger that can be used with or without Appengine. Most useful for porting code out of Appengine Standard.
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[godocs]: https://godoc.org/github.com/efixler/logger

## Installation

`go get github.com/efixler/logger`

## Usage

````
import (
	"github.com/efixler/logger"
)

logger.Context.Errorf(ctx, "There was an error: %s", err)

 ````

See the [Godoc](https://godoc.org/github.com/efixler/logger) for details and more examples. 
