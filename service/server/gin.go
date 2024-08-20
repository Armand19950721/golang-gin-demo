package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"service/controller"
	"service/protos/Common"
	"service/utils"
	"service/utils/ginResult"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	staticToken = utils.GetEnv("STATIC_TOKEN")
)

func InitGinServer() {
	r := setRouter()
	err := r.Run(utils.GetEnv("GIN_SERVER_IP"))
	if err != nil {
		log.Fatalf("upload server run fail", err.Error())
		return
	}
}

func setRouter() *gin.Engine {
	route := gin.Default()

	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowHeaders: []string{"authorization", "Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method",
			"Access-Control-Request-Headers"},
	}
	// 設定靜態檔案路徑
	prefix := "/api/"

	route.Use(cors.New(corsConf))
	route.Use(printRequest())
	route.GET(prefix+"Echo", controller.Echo)
	route.Static(prefix+"staticDir", "./staticDir")
	route.Use(middleware)

	return route
}

func checkWhiteList(route string) (res bool, routeName string) {
	// 分割路由路径
	parts := strings.Split(route, "/")

	// 定义白名单路由部分的 map
	whiteList := map[string]bool{
		"GetExcel": true,
	}

	// 遍历路由部分，检查是否在白名单中
	for _, part := range parts[2:] {
		routeName = part // 會抓到最後一個route 之後拿去給permission辨識
		if whiteList[part] {
			res = true
			return res, routeName
		}
	}

	res = false
	return res, routeName
}

func printRequest() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 打印请求方法和URL路径
		fmt.Printf("Request Method: %s, URL: %s\n", c.Request.Method, c.Request.URL.Path)

		if strings.Contains(c.Request.URL.Path, "static") {
			c.Next()
			return
		}

		// 获取查询参数
		queryParams, _ := url.ParseQuery(c.Request.URL.RawQuery)
		fmt.Printf("Query Parameters: %v\n", queryParams)

		// 读取请求体数据
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			// 处理错误
			c.Abort()
			return
		}

		// 打印请求体数据
		fmt.Printf("Request Body: %s\n", body)

		// 将读取的请求体重新设置回请求体中
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		c.Next()
	}
}

func middleware(ctx *gin.Context) {
	reply := ginResult.CommonResult{}

	// check route need token
	isInWhiteList, routeName := checkWhiteList(ctx.Request.URL.Path)
	utils.PrintObj([]string{utils.ToJson(isInWhiteList), routeName}, "isInWhiteList, routeName")

	if isInWhiteList {
		ctx.Next()
		return
	}

	// get token
	token := ctx.Request.Header.Get("Authorization")
	utils.PrintObj(token, "token")

	// check if is picbot token
	pass := token == staticToken
	utils.PrintObj(pass, "pass")

	if pass {
		ctx.Next()
		return
	}

	// can not pass
	utils.PrintObj("Abort")
	reply.Code = Common.ErrorCodes_INVAILD_TOKEN
	ctx.JSON(http.StatusOK, ginResult.GetGinResult(reply))
	ctx.Abort()
}
