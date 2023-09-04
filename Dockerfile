# Use a image with go
FROM golang:latest

# Define workspace in the container
WORKDIR /app

# Copy Go script and question files in the container (if file change, do not forget to user --no-cache tag on build)
COPY main.go /app/main.go
COPY knowledge_questions.json /app/knowledge_questions.json
COPY practice_questions.json /app/practice_questions.json


# Execute Go script 'go run main.go'
CMD ["go", "run", "main.go"]
