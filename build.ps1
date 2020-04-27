$env:GOOS="linux";$env:GOARCH="amd64"; go build .\clientA.go
$env:GOOS="linux";$env:GOARCH="amd64"; go build .\clientB.go
$env:GOOS="windows";$env:GOARCH="amd64"; go build .\clientA.go
$env:GOOS="windows";$env:GOARCH="amd64"; go build .\clientB.go