# Simple Gin Gonic REST
## A simple CRUD of Book
## Creating project
```bash
mkdir simple-web-service-gin
cd simple-web-service-gin
go mod init simple-web-service-gin
```

## In memory database
```go
var books = []book{
	{ID: "1", Title: "The Lord of the Rings - The fellowship of the ring", Author: "J. R. R. Tolkien", Price: 1.00},
	{ID: "2", Title: "The Lord of the Rings  - The two towers", Author: "J. R. R. Tolkien", Price: 2.00},
	{ID: "3", Title: "The Lord of the Rings  - The return of the king", Author: "J. R. R. Tolkien", Price: 3.00},
}
```

## Test

### GET all books
```bash
curl localhost:8080/books
```

### GET specific book
```bash
curl localhost:8080/books/1
```

### POST create a book
```bash
curl http://localhost:8080/books --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4","title": "The Hobbit","author": "Tolkien","price": 4.00}'
```