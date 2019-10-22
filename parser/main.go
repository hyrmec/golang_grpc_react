package main

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
	"google.golang.org/grpc"
	pb "grpc_ms/protobuf"
	"log"
	"net"
	"net/http"
	"strings"
)

const (
	port = ":50052"
)

type ParserService struct{}

func (s *ParserService) AddParse(ctx context.Context, in *pb.Ticker) (*pb.Stock, error) {
	log.Printf("Добавляем акцию на парсинг: %v", in.Name)
	ticker := getTicker(in.Name)
	log.Printf("Tiker is found %s", ticker)
	return ticker, nil
}

func getTicker(name string) *pb.Stock {
	name = strings.ToUpper(name)
	res, err := http.Get("https://finviz.com/quote.ashx?t=" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var titles []string
	var values []string

	doc.Find("td .snapshot-td2-cp").Each(func(i int, s *goquery.Selection) {
		titles = append(titles, s.Text())
	})
	doc.Find("td .snapshot-td2").Each(func(i int, s *goquery.Selection) {
		values = append(values, s.Text())
	})
	stoc := pb.Stock{Name: name}
	for i, value := range titles {
		if value == "P/E" {
			stoc.Eps = values[i]
		} else if value == "Price" {
			stoc.Price = values[i]
		}
	}
	return &stoc
}

func parseTikers() {
	// Парсим тикеры всех доступных акций
}

func parseSelectableTikers() {
	// Парсим интересные тикеры
}

func main() {
	// cron для запуска парсинга раз в неделю
	c := cron.New()
	err := c.AddFunc("@every 5s", parseSelectableTikers)
	// err := c.AddFunc("@weekly", gogog)
	if err != nil {
		log.Printf("Error %v", err)
	}
	c.Start()
	log.Print("Start")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterParserServer(s, &ParserService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
