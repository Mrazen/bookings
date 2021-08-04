# go run .\cmd\web\main.go .\cmd\web\middleware.go .\cmd\web\routes.go
go build -o bookings.exe .\cmd\web\.
Start-Process bookings.exe