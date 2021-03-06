# Go SDK for [CloudEvents](https://github.com/cloudevents/spec)

The Golong SDK v1 can be found on the `release-1.x.y` branch, or one of the
releases for v0, or v1 published releases.

**With v1.1.0:**

[Master](https://github.com/cloudevents/sdk-go/tree/master) will now be the base
of the effort for v2.0.0 of this SDK and will contain breaking changes or
missing libraries.

Future work on v1.X.Y releases will branch off of
[release-1.y.z](https://github.com/cloudevents/sdk-go/tree/release-1.y.z). To
add a bugfix to a v1.X.Y release, please make a PR to that branch and we can do
releases as needed on the v1 SDK. No date for EOL on v1 support yet, that will
be determined by the progress on v2.

The CloudEvents golang team is working hard to bring you v2.0.0 of the SDK.

**With v1.0.0:**

The API that exists under
[`pkg/cloudevents`](https://github.com/cloudevents/sdk-go/tree/release-1.y.z/pkg/cloudevents)
will follow semver rules. This applies to the root
[`./alias.go`](https://github.com/cloudevents/sdk-go/tree/release-1.y.z/alias.go)
file as well.

Even though `pkg/cloudevents` is v1.0.0, there could still be minor bugs and
performance issues. We will continue to track and fix these issues as they come
up. Please file a pull request or issue if you experience problems.

The API that exists under
[`pkg/bindings`](https://github.com/cloudevents/sdk-go/tree/release-1.y.z/pkg/bindings)
is a new API that will become SDK v2.x, and will replace `pkg/cloudevents`. This
area is still under heavy development and will not be following the same semver
rules as `pkg/cloudevents`. If a release is required to ship changes to
`pkg/bindings`, a bug fix release will be issued (x.y.z+1).

We will target ~2 months of development to release v2 of this SDK with an end
date of March 27, 2020. You can read more about the plan for SDK v2 in the
[SDK v2 planning doc](https://github.com/cloudevents/sdk-go/tree/release-1.y.z/docs/SDK_v2.md).

This SDK current supports the following versions of CloudEvents:

- v1.0
- v0.3
- v0.2
- v0.1

## Working with CloudEvents

Package
[cloudevents](https://github.com/cloudevents/sdk-go/tree/release-1.y.z/pkg/cloudevents)
provides primitives to work with CloudEvents specification:
https://github.com/cloudevents/spec.

Import this repo to get the `cloudevents` package:

```go
import "github.com/cloudevents/sdk-go"
```

Receiving a cloudevents.Event via the HTTP Transport:

```go
func Receive(event cloudevents.Event) {
	// do something with event.Context and event.Data (via event.DataAs(foo)
}

func main() {
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), Receive));
}
```

Creating a minimal CloudEvent in version 1.0:

```go
event := cloudevents.NewEvent()
event.SetID("ABC-123")
event.SetType("com.cloudevents.readme.sent")
event.SetSource("http://localhost:8080/")
event.SetData(data)
```

Sending a cloudevents.Event via the HTTP Transport with Binary v1.0 encoding:

```go
t, err := cloudevents.NewHTTPTransport(
	cloudevents.WithTarget("http://localhost:8080/"),
	cloudevents.WithEncoding(cloudevents.HTTPBinaryV1),
)
if err != nil {
	panic("failed to create transport, " + err.Error())
}

c, err := cloudevents.NewClient(t)
if err != nil {
	panic("unable to create cloudevent client: " + err.Error())
}
if err := c.Send(ctx, event); err != nil {
	panic("failed to send cloudevent: " + err.Error())
}
```

Or, the transport can be set to produce CloudEvents using the selected encoding
but not change the provided event version, here the client is set to output
structured encoding:

```go
t, err := cloudevents.NewHTTPTransport(
	cloudevents.WithTarget("http://localhost:8080/"),
	cloudevents.WithStructuredEncoding(),
)
```

If you are using advanced transport features or have implemented your own
transport integration, provide it to a client so your integration does not
change:

```go
t, err := cloudevents.NewHTTPTransport(
	cloudevents.WithPort(8181),
	cloudevents.WithPath("/events/")
)
// or a custom transport: t := &custom.MyTransport{Cool:opts}

c, err := cloudevents.NewClient(t, opts...)
```

Checkout the sample
[sender](https://github.com/cloudevents/sdk-go/tree/release-1.y.z/cmd/samples/http/sender)
and
[receiver](https://github.com/cloudevents/sdk-go/tree/release-1.y.z/cmd/samples/http/receiver)
applications for working demo.
