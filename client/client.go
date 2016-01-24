package client

import (
	"log"
	"reflect"

	"moetang.info/go/nekoq/use/rpc/client"

	"github.com/moetang/searchservice/api"
)

var (
	_ Client = &clientImpl{}

	insertMethod api.InsertDoc
	searchMethod api.SearchDoc
)

type Client interface {
	InsertDoc(*api.InsertDocumentReq) (*api.InsertDocumentResp, error)
	SearchDoc(*api.SearchDocumentReq) (*api.SearchDocumentResp, error)
}

type clientImpl struct {
	client.Client
}

func NewClient(connStr string) (Client, error) {
	methods := map[string]reflect.Type{
		"insert": reflect.Type(reflect.ValueOf(insertMethod).Type()),
		"search": reflect.Type(reflect.ValueOf(searchMethod).Type()),
	}
	serviceGlobalConfig := new(client.ServiceConfig)
	serviceGlobalConfig.ConnectionTimeout = 10
	serviceGlobalConfig.ServerAddr = connStr
	clientApi, err := client.NewServiceClient(make(map[string]string), methods, serviceGlobalConfig)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &clientImpl{
		Client: clientApi,
	}, nil
}

func (this *clientImpl) InsertDoc(req *api.InsertDocumentReq) (*api.InsertDocumentResp, error) {
	r, e := this.CallSync("insert", req, &client.AppInfo{})
	return r.(*api.InsertDocumentResp), e
}

func (this *clientImpl) SearchDoc(req *api.SearchDocumentReq) (*api.SearchDocumentResp, error) {
	r, e := this.CallSync("search", req, &client.AppInfo{})
	return r.(*api.SearchDocumentResp), e
}
