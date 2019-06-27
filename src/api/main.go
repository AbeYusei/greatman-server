package main

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type arg struct {
	path    string
	handler func(c *gin.Context)
}

func main() {
	router := gin.Default()
	router.Use(cors())

	m := msgAPIArg()
	router.GET(m.path, m.handler)

	router.Run()

	log.Info("server started.")
}

func msgAPIArg() *arg {
	path := "/msg"
	handler := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": buildMsg(),
		})
	}
	return &arg{path: path, handler: handler}
}

func buildMsg() string {
	targets := [5]string{"東京", "日本", "アメリカ", "世界", "資本主義経済"}
	adjectives := [5]string{"二度と見れないほどに", "根絶やしにするように", "抵抗する気も起きないほどに", "産まれてきたことを後悔させるほどに", "人が住めないような"}
	whatJohnUnDoesArr := [5]string{"火の海にしてやる", "すべてを破壊してやる", "むちゃくちゃにしてやる", "やってやる", "バキバキにしてやる"}

	i := random(1, 5)

	return targets[i] + "を" + adjectives[i] + whatJohnUnDoesArr[i]
}

func random(min int, max int) int {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	return rand.Intn((max+1)-min) + min
}

// for dev
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
