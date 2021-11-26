// Golang REST API program
package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)
/*
func PostMethod(c *gin.Context) {
  fmt.Println("\napi.go 'PostMethod' called")
  message := "PostMethod called"
  c.JSON(http.StatusOK, message)
}
*/
func GetMethod(c *gin.Context) {
  fmt.Println("\n'GetMethod' called")
  IdValue := c.Params.ByName("IdValue")
  message := "GetMethod Called With Param: " + IdValue
  c.JSON(http.StatusOK, message)

  ReqPayload := make([]byte, 1024)
  ReqPayload, err := c.GetRawData()
  if err != nil {
        fmt.Println(err)
        return
  }
  fmt.Println("Request Payload Data: ", string(ReqPayload))
}

func PostMethod(c *gin.Context) {
  fmt.Println("\n'PostMethod' called")
  IdValue := c.Params.ByName("IdValue")
  message := "PostMethod Called With Param: " + IdValue
  c.JSON(http.StatusOK, message)

  ReqPayload := make([]byte, 1024)
  ReqPayload, err := c.GetRawData()
  if err != nil {
        fmt.Println(err)
        return
  }
  fmt.Println("Request Payload Data: ", string(ReqPayload))
}

func main() {
  router := gin.Default()

  // Create Sub Router for  customised API version
  subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
    "admin": "adminpass",
  }))

  subRouterAuthenticated.POST("/:IdValue", PostMethod)
  subRouterAuthenticated.GET("/:IdValue", GetMethod)

  listenPort := "4000"
  // Listen and Server on the LocalHost:Port
  router.Run(":"+listenPort)
}
