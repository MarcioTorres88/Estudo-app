package main

import( 
"github.com/gin-gonic/gin"
"net/http"
)
func main()  {
    app := gin.Default()
    app.GET("/v1",func(ctx *gin.Context){
    ctx.JSON(http.StatusOK,gin.H{
        "ok":true,
    }) 
    })
    app.Run("localhost:3001")
}