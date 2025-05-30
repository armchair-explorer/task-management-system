Steps to run:

After cloning this repo run:

go mod tidy

then inside the directory run -> go run cmd/task-service/main.go

Then run docker-compose up --build to start the db

sample curls:

Request:
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Write integration tests",
    "description": "Write tests for the reader service",
    "status": "Pending"
  }'

Response:
{"id":5,"title":"Write integration tests","description":"Write tests for the reader service","status":"Pending","created_at":"2025-05-30T19:38:18.290293Z","updated_at":"2025-05-30T19:38:18.290293Z"}

Request:
curl -X DELETE http://localhost:8080/tasks/8                     
 
Response: 
{"id":8,"title":"","description":"","status":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}
