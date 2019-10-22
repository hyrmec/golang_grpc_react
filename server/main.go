package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	pb "grpc_ms/protobuf"
)

const (
	port          = 9090
	originAddress = "http://localhost:3000"
	addressParser = "localhost:50052"
)

type StockService struct{}

var stocks = []*pb.Stock{
	{
		Name:  "MTSS",
		Eps:   "25",
		Price: "250",
	},
	{
		Name:  "GPR",
		Eps:   "20",
		Price: "180",
	},
	{
		Name:  "ART",
		Eps:   "4",
		Price: "10",
	},
}

func (s *StockService) GetStock(ctx context.Context, in *pb.Ticker) (*pb.Stock, error) {
	log.Printf("Хотим получить информацию по акции: %v", in.Name)
	for _, stock := range stocks {
		if stock.Name == in.Name {
			return stock, nil
		}
	}
	return nil, grpc.Errorf(codes.NotFound, "")
}

func (s *StockService) GetAllStocks(ctx context.Context, in *empty.Empty) (*pb.AllStocks, error) {
	log.Print("Хотим получить информацию по всем акциям!")
	return &pb.AllStocks{Data: []*pb.Stock{{Name: "MTSS", Eps: "23.5", Price: "125"}}}, nil
}

func (s *StockService) Add(ctx context.Context, in *pb.Ticker) (*empty.Empty, error) {
	out := new(empty.Empty)
	conn, err := grpc.Dial(addressParser, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewParserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddParse(ctx, &pb.Ticker{Name: in.Name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	} else {
		log.Printf("Добавляем акцию: %v", in.Name)
		log.Print(r)
		stocks = append(stocks, r)
		log.Print(stocks)
	}
	return out, nil
}

func originFunc(origin string) bool {
	if origin == originAddress {
		return true
	} else {
		return false
	}
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterStocksServer(grpcServer, &StockService{})

	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(originFunc))
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}
