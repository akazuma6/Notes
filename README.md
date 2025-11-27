# Notes

go run main.go

curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"mypassword123"}'