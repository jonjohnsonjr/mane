# mane

A better `main` function.

Note that this is mostly for my personal debuging workflows, and I might make breaking changes.

It is expected that you have [`otel-desktop-viewer`](https://github.com/CtrlSpice/otel-desktop-viewer) and `go` installed.

## pprof

This will pop open the go pprof web view:

```go
package main

import (
	"context"
	"log"

	"github.com/jonjohnsonjr/mane/pprof"
	"exampel.com/my/cobra/based/cli"
)

func main() {
	log.Fatal(pprof.Main(context.Background(), cli.New().ExecuteContext))
}
```

## trace

This will pop open the go pprof web view:

```go
package main

import (
	"context"
	"log"

	"github.com/jonjohnsonjr/mane/trace"
	"exampel.com/my/cobra/based/cli"
)

func main() {
	log.Fatal(trace.Main(context.Background(), cli.New().ExecuteContext))
}
```

## both

```go
package main

import (
	"context"
	"log"

	"github.com/jonjohnsonjr/mane/pprof"
	"github.com/jonjohnsonjr/mane/trace"
	"exampel.com/my/cobra/based/cli"
)

func main() {
	log.Fatal(pprof.Main(context.Background(), mainE))
}

func mainE(ctx context.Context) error {
	return trace.Main(ctx, cli.New().ExecuteContext)
}
```
