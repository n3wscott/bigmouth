# bigmouth

Web-based CloudEvents sender.

## Usage

TBD: usage.

See also, the CloudEvents [Spec](https://github.com/cloudevents/spec) or
[golang SDK](https://github.com/cloudevents/sdk-go) to get started sending
CloudEvents formatted events.

## Running Locally

```shell
FILE_PATH=./cmd/bigmouth/kodata go run cmd/bigmouth/main.go
```

## Running on Kubernetes

### From Release v0.1.0 (pending)

To install into your default namespace
```shell
kubectl apply -f https://github.com/n3wscott/bigmouth/releases/download/v0.1.0/release.yaml
```

### From Source

```shell
ko apply -f config/bigmouth.yaml
```
