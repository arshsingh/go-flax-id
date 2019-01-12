Go implementation of Flax ID: https://github.com/ergeon/python-flax-id

#### Usage

```go
package main

import (
    "github.com/arshsingh/go-flax-id"
)

func main() {
    id := flaxid.New()  // 4e3tRv_8hyBIuXj5
    ...
}
```

Get an ID for a specified timestamp instead of the current moment:

```go
timestamp := time.Unix(1424300000, 0)
id := flaxid.ForTimestamp(timestamp)  // -EkPZk2dAFoHm_t4
```
