package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {

	server := &http.Server{

		Addr: "127.0.0.1:8000",

		ReadTimeout: 3 * time.Second,

		WriteTimeout: 5 * time.Second,

		// IdleTimeout: 0 * time.Second,
	}

	http.HandleFunc("/healthcheck", healthCheck)
	http.HandleFunc("/echo", echoHandler)
	log.Println("go http srever")
	log.Fatal(server.ListenAndServe())
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

	// log.Printf("loginUuid : %s \n", loginUuid)
	log.Printf("%+v", map[string]interface{}{
		"name": "echoHandler",
	})

	upgrader := &websocket.Upgrader{
		//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	connect, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("%+v", map[string]interface{}{
			"name": "echoHandler connect err",
		}, err)
		if connect != nil {
			connect.Close()
		}
		return
	}

	for {
		// isConnect := socket.SocketHandler(w, r, connCore)
		// if !isConnect {
		// 	break
		// }
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL)
	if r.URL.Path != "/healthcheck" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// http.ServeFile(w, r, "./client/healthCheck.html")
	fmt.Fprintf(w, "healthCheck ver : "+os.Getenv("version"))

}

func f(n int) []int {
	sum := e(n, 0)
	factor := []int{sum}
	if n-sum >= 1 {
		factor = append(factor, f(n-sum)...)
	}
	return factor
}

func e(n int, t int) int {
	if n < 2 {
		return int(math.Pow(2, float64(t)))
	}
	sum := e(n/2, t+1)
	return sum
}

type node struct {
	Val       int
	LeftNode  *node
	RightNode *node
}

func sumNode() int {

	root := node{
		Val: 1,
		LeftNode: &node{
			Val: 2,
			LeftNode: &node{
				Val: 4,
			},
			RightNode: &node{
				Val: 5,
				LeftNode: &node{
					Val: 8,
				},
				RightNode: &node{
					Val: 9,
				},
			},
		},
		RightNode: &node{
			Val: 3,
			LeftNode: &node{
				Val: 6,
				LeftNode: &node{
					Val: 10,
				},
				RightNode: &node{
					Val: 11,
				},
			},
			RightNode: &node{
				Val: 7,
			},
		},
	}

	sum := add(&root)

	return sum
}

func add(root *node) int {
	sum := root.Val
	if root.LeftNode != nil {
		sum = sum + add(root.LeftNode)
	}
	if root.RightNode != nil {
		sum = sum + add(root.RightNode)
	}
	return sum
}
