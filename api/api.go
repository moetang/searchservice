package api

import "moetang.info/go/nekoq/use/rpc"

type InsertDoc func(*InsertDocumentReq) (*InsertDocumentResp, rpc.ErrStr)

type InsertDocumentReq struct {
	Id      uint64
	Content string
}

type InsertDocumentResp struct {
	Success bool
}

type SearchDoc func(*SearchDocumentReq) (*SearchDocumentResp, rpc.ErrStr)

type SearchDocumentReq struct {
	Keyword string
}

type SearchDocumentResp struct {
	DocIds []uint64
}
