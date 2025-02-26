"# deepits" 
## Installation and Running

### Prerequisites
- Go (version 1.16 or higher)
- Git

### Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/gecko-team-2024/deepits.git
    ```
2. Navigate to the project directory:
    ```sh
    cd deepits
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Project
1. Build the project:
    ```sh
    go run cmd/master.go
    go run slave/slave.go
    ```
2. Run the executable:
    ```sh
    ./deepits
    ```

### Testing
To run tests:
```sh
go test ./...
```