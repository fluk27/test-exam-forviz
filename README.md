# test-exam-forviz
api create update delete read borrow and return book.
### Setup and Run:

### Setup

#### install
    
 - [install Go](https://go.dev/doc/install)

- [install sqlite](https://www.sqlite.org/download.html)
#### config at path "config/confg.yaml"
1. config number port 
    ```bash
        port: { { app-port } }
    ```
2. config sqlite  database name 
    ```bash
        dbname: { { sqlite-dbname } }
    ```
3. config sqlite database path 
    ```bash
        dbpath: { { sqlite-dbpath } }
    ```
### Run Go
1. run install all package.

    ```bash
    go mod tidy
    ```

    2. start go server.

    ```bash
    go run cmd/main.go
    ```
### test Go
1. run unit test all files and display coverage.

    ```bash
    go test ./... -cover
    ```