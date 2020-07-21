# Golang mod Serverless template

## Requirements

- Serverless CLI
- AWS SAM local CLI
- Docker
- NodeJS
- Make

## Instructions

Don't forget to initialize the go mod package and install the `aws-lambda-go`
package:

```console
go mod init <package-name>
go get github.com/aws/aws-lambda-go
```

Install serverless dependencies and configurations:

```console
npm install
```
