package main

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/6/10 上午11:55
 */

import (
	"test/router"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//
//	if r.URL.Path[1:] == "EXIT" {
//		os.Exit(0)
//	}else if r.URL.Path[1:] == "genP10" {
//		fmt.Fprintf(w, "Hi there, I love %s!\n", "SCONE")
//		p10, _ := GenP10(34997, "12", RandStringInt(), "SM2")
//		fmt.Fprintf(w, "p10 = %s\n",p10)
//	}
//}

//func main() {
//	http.HandleFunc("/", handler)
//	http.ListenAndServe(":9443", nil)
//}
func main() {
	//r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "Hello, Geektutu")
	//})
	//r.Run(":9443") // listen and serve on 0.0.0.0:8080

	route := router.CreateRouter()
	route.Run(":9443")

}
