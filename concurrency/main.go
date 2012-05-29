package main

import (
	"fmt"
	"sync"
)

const (
	NORMAL_REQUEST = iota
	EXIT
	NORMAL_RESPONSE
)

type UserConn struct {
	Name string
	nr_req int
}

type Request struct {
	Type int
}

type Response struct {
	Type int
}

func (c *UserConn) ReadRequest() *Request {
	if c.nr_req <= 0 {
		return &Request{EXIT}
	}
	c.nr_req --
	return &Request{NORMAL_REQUEST}
}

func (c *UserConn) WriteResponse(res *Response) {
	fmt.Printf("send respond %v to user %v\n",
			   res.Type, c.Name)
}

func calculate(req *Request) *Response {
	return &Response{NORMAL_RESPONSE}
}

func writeRes(conn *UserConn,
			  ch chan *Response) {
	for r := range ch {
		conn.WriteResponse(r)
	}
}

func process(req *Request,
			 ch chan *Response,
			 wg *sync.WaitGroup) {
	res := calculate(req)
	ch <-res
	wg.Done()
}

func ServeClient(conn *UserConn,
				 quit chan bool) {
	var wg sync.WaitGroup
	resCh := make(chan *Response)

	go writeRes(conn, resCh)
	defer func () {
		wg.Wait()
		quit <- true
	}()

	for {
		// 读取一个请求，
		//  判断类型
		// 如果用户请求关闭，
		//  则函数返回
		req := conn.ReadRequest()
		switch req.Type {
		case NORMAL_REQUEST:
			wg.Add(1)
			go process(req, resCh, &wg)
		case EXIT:
			return
		}
	}
}

func main() {
	quit := make(chan bool)
	conn := &UserConn{"alice", 10}
	go ServeClient(conn, quit)
	<-quit
}

