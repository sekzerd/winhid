$env:GOOS = "windows"
$env:GOARCH = "386" 
$env:CGO_ENABLED = 1
$env:GOGCCFLAGS = "-m32 -mthreads -fmessage-length=0"
go build .