# AWS Cost Monitor

This project sets up cost monitoring for a single AWS account.

## Architecture

- Go program using AWS Go SDK queries AWS API for resource usage.
- Resources not tagged with `team` are flagged and/or automatically deleted.
- Properly tagged resources are:
    1. Compared against set of rules specified for given `team`
    2. If it breaks a rule, the resource is flagged and/or automatically deleted
- Data persisted to Elasticsearch
- Kibana dashboards to show rule violations per `team`
- Input should be yaml. Example:
```
devops:
    env: prod
    ec2:
        - maxsize: t2.micro
        - maxcount:
    asg:
        ...
```

## Future Work

- Add to a billing account which has access to all accounts.

## Local Development

### Set up Go

- For Mac, download Go binary: https://golang.org/doc/install?download=go1.16.3.darwin-amd64.pkg
- Verify version with `go version`
- `cd /path/to/aws-cost-monitor`
- `go run cmd/aws-cost-monitor/main.go`

### Set up AWS

- Create an AWS account (free tier is fine)
- Install AWS SDK for Go `go get -u github.com/aws/aws-sdk-go/...`
- Retrieve access key ID and secret access key of the IAM user
- Import the AWS SDK for Go packages in relevant `.go` files
```
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)
```