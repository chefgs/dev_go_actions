## Table of Contents
- [Build Status](#build-status)
- [Introduction](#introduction)
- [My Workflow](#my-workflow)
- [Submission Category](#submission-category)
- [Yaml File and Link to Code](#yaml-file-and-link-to-code)
- [GitHub Action Workflow Run Output](#github-action-workflow-run-output)
- [Additional Info](#additional-info)
  - [How to Add GitHub Actions workflow](#how-to-add-github-actions-workflow)
  - [Go Source Code Details](#go-source-code-details)
  - [Go REST API Unit testing](#go-rest-api-unit-testing)

## Build Status
![GO CI workflow](https://github.com/chefgs/dev_go_actions/actions/workflows/go_apiauth.yml/badge.svg)

## Introduction
- Created a GitHub Actions workflow for the submission of Dev.to  `actionshackathon21` 
- This action performs the Go continuous integration process on the Go source code repository
- Interested *Go developers and DevOps Engineers* can use this workflow to create the `Continuous Integration` for their GitHub repo

## My Workflow
### Go Continuous Integration workflow
- The workflow we created will be doing the Continuous Integration process on this repository.
- Whenever there is a code check-in happens on `main branch`, then CI workflow will be triggered
- Below are the details of this workflow
  - Code checkout into the workspace
  - Install Go and runs the Go linting process for doing code review
  - Get the build dependencies
  - Builds the code
  - Runs the Unit test, if it is success it moves to next stage
  - Docker image is created for the code and pushed into `docker hub` registry 
  - Finally workflow posts the status of each step to a `slack channel`
- [Here is the link](github https://github.com/chefgs/dev_go_actions/blob/main/.github/workflows/go_apiauth.yml ) to the GitHub actions workflow `Yaml`

### Marketplace GitHub Actions used in this Workflow 
I've leveraged existing actions available from [GitHub Actions Marketplace](https://github.com/marketplace?type=actions)
- `actions/checkout@v2` - [Action](https://github.com/marketplace/actions/checkout) that is used to checkout code. 
- `reviewdog/action-golangci-lint@v2` - Go installation and linting [action](https://github.com/marketplace/actions/run-golangci-lint-with-reviewdog).
- `mr-smithers-excellent/docker-build-push@v5` - Docker build and push [action](https://github.com/marketplace/actions/docker-build-push-action).

## Submission Category 
- **DIY Deployments** 

## Yaml File and Link to Code
- GitHub Actions Workflow `Yaml`
```yaml
name: Go_CI_Workflow

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:

  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
    - name: Check out code to Build
      uses: actions/checkout@v2
      
    - name: Install Go and Run Code Linting
      id: Install-Go
      uses: reviewdog/action-golangci-lint@v2
      
    - name: Get dependencies to Build
      id: Get-dependencies-to-Build
      run: |
        go get -v -t -d ./...
        if [ -f Gopkg.toml ]; then
            curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
            dep ensure
        fi
      
    - name: Build Code
      id: Build-Code
      run: |
        go build -v .
        
    - name: Unit Test
      id: Unit-Test-Run
      run: |
        go test

    - name: Build & push Docker image
      id: Build-and-Push-Docker-Image
      uses: mr-smithers-excellent/docker-build-push@v5
      with:
        image: gsdockit/goapiauth
        tags: latest
        registry: docker.io
        dockerfile: Dockerfile
        username: ${{ secrets.DOCKER_USERNAME }}
        password: ${{ secrets.DOCKER_PASSWORD }}
        
    - name: Posting Action Workflow updates to Slack
      uses: act10ns/slack@v1
      with: 
        status: ${{ job.status }}
        steps: ${{ toJson(steps) }}
        channel: '#github-updates'
      if: always()
    env:
      SLACK_WEBHOOK_URL: ${{ secrets.SLACK_WEBHOOK_URL }}
```

- The full source code can be accessible in this [GitHub repo](https://github.com/chefgs/dev_go_actions)

{% github https://github.com/chefgs/dev_go_actions %}

## GitHub Action Workflow Run Output
- Here is how it runs the workflow and it can be viewable from the `Actions` tab
![Go CI Run](https://github.com/chefgs/repo_images/blob/master/GO_CI_workflowRun.png?raw=true)

___
## Additional Info

**Other details that might be helpful to users starting to write GitHub Actions**

## How to Add GitHub Actions workflow
- Our workflow gets created in `.github/workflows` as an `.yaml` file

### Method 1:
- Step 1: We can find `Actions` tab on the repo
- Step 2: Click on the tab and choose `New Workflow`
- Step 3: Choose workflow from the pre-defined template
- Step 4: Edit the workflow yaml according to the required stages 

or 

### Method 2:
- Create workflow for yourself from scratch using the [documentation guide](https://docs.github.com/en/actions/quickstart#introduction)

> *Enjoy creating the GitHub Workflow*  

---
## How to Integrate Slack with GitHub Actions Workflow
Here is the list of steps to integrate slack with github actions,

- Add `github` app to the slack channel
- Add `incoming webhook` app and pick `webhook url` and keep it
- Run the below command in slack channel to `subscribe` to repo updates. (while adding, it asks for auth with github and we choose repo) 
```
/github subscribe user/repo
```
- Add `incoming webhook URL` as `secret` in the specific repo we want to get updates
- Add the below code at the end of Action workflow to get the job updates in slack
```
  - name: Posting Action Workflow updates to Slack
      uses: act10ns/slack@v1
      with: 
        status: ${{ job.status }}
        steps: ${{ toJson(steps) }}
        channel: '#system-tester'
      if: always()
    env:
      SLACK_WEBHOOK_URL: ${{ secrets.SLACK_WEBHOOK_URL }}
```
- Each steps in workflow should have `id: step-name` to get the status of the step in Slack update

---
**Now let's see the details about the code and it's functionality...**
## Go Source Code Details

## Code Objective
- We will be creating a REST API that listens on `localhost` port `1357` and has the API versioning with a query string parameter.
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

