# GitHub Actions Hackathon by Dev Community

## Go Continuous Integration workflow
- The workflow is doing the Continuous Integration process on this Go code repository.
- Below are the details of this workflow
  - Code checkout into the workspace
  - Install Go executable and runs the Go linting process for doing code review
  - Get the build dependencies
  - Builds the code
  - Runs the Unit test, if it is success it moves to next stage
  - Finally docker image is created for the code and pushed into `docker hub` registry 

## Marketplace GitHub Actions used in this Workflow 
We levergaed existing actions in this workflow from [GitHub Actions Marketplace](https://github.com/marketplace?type=actions)
1. `actions/checkout@v2` - [Action](https://github.com/marketplace/actions/checkout) that is used to checkout code. 
2. `reviewdog/action-golangci-lint@v2` - Go installation and linting [action](https://github.com/marketplace/actions/run-golangci-lint-with-reviewdog).
3. `mr-smithers-excellent/docker-build-push@v5` - Docker build and push [action[(https://github.com/marketplace/actions/docker-build-push-action).

## How to Add GitHub Actions workflow
- Our workflow gets created in `.githuh/workflows` as an `.yaml` file
### Method 1:
- We can find `Actions` tab on the repo
- Click on the tab and choose `New Worflow`
- Choose workflow from the pre-defined template

or 

### Method 2:
- Create workflow for yourself from scratch using the [documentation guide](https://docs.github.com/en/actions/quickstart#introduction)

> *Enjoy coding the GitHub Workflow for the repo*  

---

**Now let's see the details about the code and it's functionality...**
## Go Source Code Details

## Objective
- We will be creating REST API that listens on `localhost` port `1357` and has the API versioning with query string parameter.
- Sample URL format we are planning to create, `http://localhost:1357/api/v1/PersonId/Id456`

## Function main
Add the below code function `main()`

## Adding API Versioning and Basic authentication
```
  // Define gin router
  router := gin.Default()

  // Create Sub Router for customized API version and basic auth
  subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
    "basic_auth_user": "userpass",
  }))
```

## Passing Query String Parameters
```
  subRouterAuthenticated.GET("/:IdValue", GetMethod)
```
## REST API Listening on Port
```
  listenPort := "1357"
  // Listen and Server on the LocalHost:Port
  router.Run(":"+listenPort)
```

## Get Method
- Define a `GetMethod` function and add the following code 
- It fetches and prints the `Person IdValue` from the query string parameter passed in the API URL
```
func GetMethod(c *gin.Context) {
  fmt.Println("\n'GetMethod' called")
  IdValue := c.Params.ByName("IdValue")
  message := "GetMethod Called With Param: " + IdValue
  c.JSON(http.StatusOK, message)

  // Print the Request Payload in console
  ReqPayload := make([]byte, 1024)
  ReqPayload, err := c.GetRawData()
  if err != nil {
        fmt.Println(err)
        return
  }
  fmt.Println("Request Payload Data: ", string(ReqPayload))
}
```

## Go REST API Unit testing
- Go `testing` module can be used for creating unit testing code for Go source
- Go testing module code has been found in `api_authtest.go`
- Run the command `go test` to run the tests
