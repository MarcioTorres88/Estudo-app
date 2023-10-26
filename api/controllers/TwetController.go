package controllers

import ("net/http"
	"github.com/gin-gonic/gin"
	entities "api/api/entities"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}
func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet() 
	
	if err := ctx.ShouldBindJSON(&tweet); err!= nil {
        return
    }
	
	t.tweets = append(t.tweets, *tweet)
	ctx.JSON(http.StatusOK, t.tweets)
}
