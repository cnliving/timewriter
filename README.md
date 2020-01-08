# TimeWriter

TimeWriter implements io.Writer to roll daily and compress log file time

# Version

2020-01-12

​    1. use independent coroutines to process

​    1.1 the names of  daily  log file

​    1.2 compress log file

​    1.3  delete expired  log file

   2.delay compressed history file for easy reading by other programs

## Overview

* implements io.Writer. You can easily use in golang log, GORM, grpclog etc.
* daily roll log, the file name's prefix is process name
* compress to gz for old file

## Getting Started

**Example**

To use TimeWriter, you can git clone [https://github.com/longbozhan/timewriter](https://github.com/longbozhan/timewriter), and import like this:

```go
package main

import (
	"log"
)

func main() {
	timeWriter := &TimeWriter{
		Dir:        "./log",
		Compress:   true,
		ReserveDay: 30,
	}
	logDebug := log.New(timeWriter, " [Debug] ", log.LstdFlags)
	logDebug.Println("this is debug")
}

```

## Reference

* [lumberject](https://github.com/natefinch/lumberjack)
