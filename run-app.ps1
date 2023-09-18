go build main.go

Start-Process "chrome" -ArgumentList "--new-window", "http://localhost:8000"

./main.exe
