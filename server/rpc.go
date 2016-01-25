package server

import (
	"log"
	"reflect"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"

	"moetang.info/go/nekoq/use/rpc"
	"moetang.info/go/nekoq/use/rpc/server"

	"github.com/moetang/searchservice/api"
)

var (
	searchengine = &engine.Engine{}

	insertMethod api.InsertDoc = insertDocMethod
	searchMethod api.SearchDoc = searchDocMethod
)

type ServerConfig struct {
	ListenString   string
	DictionaryPath string
}

func StartAndListen(config *ServerConfig) {
	searchengine.Init(types.EngineInitOptions{
		SegmenterDictionaries: config.DictionaryPath,
	})

	methods := map[string]reflect.Value{
		"insert": reflect.ValueOf(insertMethod),
		"search": reflect.ValueOf(searchMethod),
	}
	serverGlobalConfig := new(server.ServiceConfig)
	serverGlobalConfig.Listen = config.ListenString
	server, err := server.NewServer(make(map[string]string), methods, serverGlobalConfig)
	if err != nil {
		log.Fatalln(err)
		return
	}
	var _ = server
}

func insertDocMethod(req *api.InsertDocumentReq) (*api.InsertDocumentResp, rpc.ErrStr) {
	// 将文档加入索引
	searchengine.IndexDocument(req.Id, types.DocumentIndexData{Content: req.Content})

	resp := &api.InsertDocumentResp{}
	resp.Success = true
	return resp, ""
}

func searchDocMethod(req *api.SearchDocumentReq) (*api.SearchDocumentResp, rpc.ErrStr) {
	resp := searchengine.Search(types.SearchRequest{Text: req.Keyword})

	response := &api.SearchDocumentResp{}
	response.DocIds = make([]uint64, len(resp.Docs))

	for i := 0; i < len(resp.Docs); i++ {
		response.DocIds[i] = resp.Docs[i].DocId
	}

	return response, ""
}
