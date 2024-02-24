# mane

A better `main` function.

Note that this is mostly for my personal debuging workflows, and I might make breaking changes.

## Main

Turn this boilerplate:

```go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()

	if err := foo(ctx); err != nil {
		log.Fatal(err)
	}
}
```

Into this:

```go
package main

import (
	"github.com/jonjohnsonjr/mane"
)

func main() {
    mane.Main(foo)
}
```

## pprof

Dependency: `go`

This will pop open the go pprof web view:

```go
package main

import (
	"context"
	"log"

	"github.com/jonjohnsonjr/mane"
	"github.com/jonjohnsonjr/mane/pprof"
	"example.com/my/cobra/based/cli"
)

func main() {
	log.Fatal(pprof.Main(mane.Context(), cli.New().ExecuteContext))
}
```

## trace

```go
package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"

	"github.com/jonjohnsonjr/mane/trace"
	"github.com/jonjohnsonjr/mane"
)

func main() {
    f, err := os.CreateTemp("", "my-trace-file")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    trace.Main(mane.Context(), f, func(ctx context.Context) error {
	    if err := foo(ctx); err != nil {
	            return err
	    }

	    return bar(ctx)
    })

    log.Printf("Writing tracefile to %s", f.Name())
}

func foo(ctx context.Context) error {
	ctx, span := otel.Tracer("example.com/something").Start(ctx, "foo")
	defer span.End()

	time.Sleep(1*time.Second)

	return nil
}

func bar(ctx context.Context) error {
	ctx, span := otel.Tracer("example.com/something").Start(ctx, "bar")
	defer span.End()

	time.Sleep(2*time.Second)

	return nil
}
```

## both

This does both!

```go
package main

import (
	"context"
	"log"

	"github.com/jonjohnsonjr/mane"
	"github.com/jonjohnsonjr/mane/pprof"
	"github.com/jonjohnsonjr/mane/trace"
	"example.com/my/cobra/based/cli"
)

func main() {
	log.Fatal(pprof.Main(mane.Context(), mainE))
}

func mainE(ctx context.Context) error {
	return trace.Main(ctx, cli.New().ExecuteContext)
}
```
