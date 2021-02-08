package main

import (
	"TeaPicker/store"
	"TeaPicker/tree"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

/*
1、实现GRPC通信接口	done
2、loading配置文件   简化至flag处理  done
3、单独线程进行磁盘持久化，永久保存数据？AOF（记录读写记录）：RDB（生成时间快照） done
4、主从集群？  就暂时压力而言没必要  pending
*/
import (
	"flag"
	"github.com/json-iterator/go"
)

var Rootmap *tree.Mapping
var record *store.Record
var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	dataPath    string
	port        string
	cacheSize   int
	cacheExpire int
	tuning      bool
)

func init() {
	//初始化参数
	flag.StringVar(&dataPath, "P", "data.tmd", "数据文件路径：/home/abel/data.tmd。有则读取，无则创建。")
	flag.StringVar(&port, "p", "9988", "端口号")
	flag.IntVar(&cacheSize, "cache-size", 256, "缓存内容数量")
	flag.IntVar(&cacheExpire, "cache-expire", 3600, "缓存时间(s)")
	flag.BoolVar(&tuning, "m", false, "优化数据文件")
}

func main() {
	flag.Parse()

	fmt.Printf("缓存内容数量：%d\n", cacheSize)
	fmt.Printf("缓存内容时间：%d\n", cacheExpire)
	//初始化数据
	record = store.CreateRecord(dataPath)
	Rootmap = tree.CreateMapping(cacheSize, cacheExpire)

	record.RegisterStoreStruct(store.TAG_Tree, Rootmap)
	if tuning {
		//数据文件调优
		record.Optimize()
		os.Exit(1)
	}

	fmt.Println("加载数据...")
	//启动加载数据项到内存
	record.Recovery()
	fmt.Println("数据加载完毕")

	//启动服务
	listener, err := net.Listen("tcp", ":"+port)
	if nil != err {
		panic(err)
	} else {
		fmt.Println("Listen at:" + port + ".....")
	}
	defer listener.Close()

	rpcServer := grpc.NewServer()
	tree.RegisterAddServiceServer(rpcServer, &Service{})
	tree.RegisterDelServiceServer(rpcServer, &Service{})
	tree.RegisterSearchServiceServer(rpcServer, &Service{})
	tree.RegisterSearchLenovosServiceServer(rpcServer, &Service{})
	rpcServer.Serve(listener)
}

//服务包装
type Service struct {
	tree.AddServiceServer
	tree.DelServiceServer
	tree.SearchServiceServer
	tree.SearchLenovosServiceServer
}

//增加定时器功能够需要做到协程安全
func (s *Service) Add(ctx context.Context, request *tree.AddRequest) (*tree.BooleanResponse, error) {
	b := Rootmap.AddWord(request.Index, request.Key, request.JsonData, request.Chinese)
	if b {
		addWord := tree.AddWordCommand{request.Index, request.Key, request.JsonData, request.Chinese}
		argsString, err := json.MarshalToString(&addWord)
		if err == nil {
			record.AddLog(store.TAG_Tree, store.OperateAdd, argsString)
		}
	}
	return &tree.BooleanResponse{b}, nil
}

func (s *Service) Del(ctx context.Context, request *tree.DelRequest) (*tree.BooleanResponse, error) {
	b := Rootmap.DelWord(request.Index, request.Key, request.Chinese)
	if b {
		delWord := tree.DelWordCommand{request.Index, request.Key, request.Chinese}
		argsString, err := json.MarshalToString(&delWord)
		if err == nil {
			record.AddLog(store.TAG_Tree, store.OperateDel, argsString)
		}
	}
	return &tree.BooleanResponse{b}, nil
}

func (s *Service) Search(ctx context.Context, request *tree.SearchRequest) (*tree.SearchResponse, error) {
	res := Rootmap.Search(request.Index, request.Context)
	ret := []*tree.SearchResult{}
	for _, v := range res {
		result := tree.SearchResult{Str: v.Str, Data: v.Data}
		ret = append(ret, &result)
	}
	return &tree.SearchResponse{ret}, nil
}

func (s *Service) SearchLenovos(ctx context.Context, request *tree.SearchLenovosRequest) (*tree.SearchResponse, error)  {
	res := Rootmap.SearchLenovos(request.Index, request.Context)
	ret := []*tree.SearchResult{}
	for _, v := range res {
		result := tree.SearchResult{Str: v.Str, Data: v.Data}
		ret = append(ret, &result)
	}
	return &tree.SearchResponse{ret}, nil
}
