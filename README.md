
![](https://raw.githubusercontent.com/williamauhk/golang-backend-challenge/master/screenshot_for_result_console.png)

## Background

 The open source community has mostly settled on using Github and its releases feature to publish releases and are also mostly using Semantic Versioning as their versioning structure.

## Intro

This project  want to be able to track when new versions of all of these applications are released. It's a simple application that gives us the highest patch version of every release between a minimum version and the highest released version. reads the Github Releases list, uses SemVer for comparison and takes a path to a file as its first argument when executed.

## Command and files

 the format of file.txt:
```
repository,min_version
kubernetes/kubernetes,1.8.0
prometheus/prometheus,2.2.0
```
takes a path to a file as its first argument when executed.
```
go run main.go file.txt
```
and it should produce output to stdout in the format of:
```
latest versions of kubernetes/kubernetes: [1.10.1 1.9.6 1.8.11]
latest versions of prometheus/prometheus: [2.2.1]
```

## Testing

Run
```
go test
```

Which contains some test cases that should pass when the application is ready.
```
{[1.8.11 1.9.6 1.10.1 1.9.5 1.8.10 1.10.0 1.7.14 1.8.9 1.9.5] [1.10.1 1.9.6 1.8.11] 0xc420060840}
[1.10.1 1.9.6 1.8.11]
{[1.8.11 1.9.6 1.10.1 1.9.5 1.8.10 1.10.0 1.7.14 1.8.9 1.9.5] [1.10.1 1.9.6] 0xc420060880}
[1.10.1 1.9.6]
{[1.10.1 1.9.5 1.8.10 1.10.0 1.7.14 1.8.9 1.9.5] [1.10.1] 0xc4200608c0}
[1.10.1]
{[2.2.1 2.2.0] [2.2.1] 0xc420060900}
[2.2.1]
PASS
```

