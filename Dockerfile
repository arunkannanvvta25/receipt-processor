FROM golang:1.22.5
WORKDIR /receipt-processor
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["go", "run", "main.go"]
EXPOSE 8080

# docker build -f Dockerfile -t my-receipt-processor .
# docker run --name receipt-processor-container -d -p 5001:8080 my-receipt-processor
# endpoints: http://localhost:5001/receipts/process  