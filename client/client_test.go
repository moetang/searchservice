package client_test

import (
	"github.com/moetang/searchservice/api"
	"github.com/moetang/searchservice/client"
	"log"
	"testing"
	"time"
)

func TestSample(t *testing.T) {
	c, e := client.NewClient("tcp://127.0.0.1:14357")
	if e != nil {
		t.Fatal(e)
	}

	insertReq := &api.InsertDocumentReq{
		Id:      1,
		Content: "中国最大的搜索引擎是百度搜索引擎",
	}
	resp, err := c.InsertDoc(insertReq)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(resp.Success)

	time.Sleep(1 * time.Second)

	searchReq := &api.SearchDocumentReq{
		Keyword:"百科",
	}
	searchResp, err := c.SearchDoc(searchReq)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(searchResp.DocIds)

}
