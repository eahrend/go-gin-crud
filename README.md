# go-gin-crud
Basic Golang CRUD app with a MySQL backend.
api:
    environment:
      - PORT=8080      
    build:
      context: .
    ports:
      - "8080:8080"
    depends_on:
      - sql