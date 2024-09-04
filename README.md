## A sample project using Golang, Graphql and REST based API endpoints for learning and understanding.

### GoLang Naming Rules and Conventions

#### File Name: snake_case

- Source files are all lower case with underscore separating multiple words.
Compound file names are separated with _. eg: main_schema.go, user_detail.go
- File names that begin with “.” or “_” are ignored by the go tool
- Files with the suffix _test.go are only compiled and run by the go test tool.

#### Functions and Methods: CamelCase
- Exported functions should start with uppercase:

`writeToDB // unexported, only visible within the package`

`WriteToDB // exported, visible within the package`

#### Constants: CONSTANT_VARIABLE
- All capital letters and use underscore _ to separate words.

#### Variables
- Generally, use relatively simple (short) name.
`user to u`

`userID to uid`

- **Boolean** variable, should start with **Has**, **Can** or **Is**, etc.
- **Index** variable should be single letter: `i, j, k`



### Go Toolchain

1. **`go build`**: Compiles the Go code into an executable binary.
2. **`go run`**: Compiles and runs a Go program.
3. **`go test`**: Runs tests for a Go package.
4. **`go fmt`**: Formats Go source code according to the Go style guidelines.
5. **`go vet`**: Analyzes Go code for potential issues and bugs.
6. **`go get`**: Downloads and installs Go packages and dependencies.
7. **`go install`**: Compiles and installs a Go binary.
8. **`go mod`**: Manages Go modules, which handle dependencies and module versions.
9. **`go list`**: Lists information about Go packages.

### Running Go app on localhost
- Pre-requisite: Your system should have Golang and Postgres DB installed and running. You can follow any tutorial available online to install. 


- Database setup:
  1. Create Database in Postgres DB with <db-name> of your choice.
  2. Create a <db-role> with Password and privileges on the <db-name>.
  3. Create a table to store data:

  `CREATE TABLE gods ( id SERIAL PRIMARY KEY,  name TEXT NOT NULL, description TEXT, other_names TEXT );`


- To run in development mode:
  1. Go to terminal
  2. Go to your project folder and run `cd Project-Golang-Graohql-00/olympus/pkg/main`
  3. Run `go run main.go`

### Making REST API calls
- Use Postman and create new REST request type
- Sample Rest URL `http://localhost:8080/gods/add`

### Making Graphql API calls
- Use Postman and create new Graphql request type
- Graphql URL `http://localhost:8080/graphql`
#### Sample Graphql Query
`query GodList {
godList {
description
id
name
otherNames
}
}`

#### Sample Graphql Mutation
`mutation AddGod {
    addGod(name: "Hera", description: "Wife of Zeus") {
        description
        id
        name
        otherNames
    }
}
`
