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

### Running Go app
To run in development mode:
1. Go to terminal
2. Go to your project folder and run `cd Project-Golang-Graohql-00/olympus/pkg/main`
3. Run `go run main.go`

### Running Graphql on localhost
- Use Postman and create new Graphql request type
- Graphql URL `http://localhost:8080/graphql`
### Queries
`query GodList {
godList {
description
id
name
otherNames
}
}`

### Mutations
`mutation AddGod {
    addGod(name: "Hera", description: "Wife of Zeus") {
        description
        id
        name
        otherNames
    }
}
`
