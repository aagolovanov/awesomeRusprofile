package main

import (
	"fmt"
	"github.com/aagolovanov/awesomeRusprofile/pkg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"path"
	"strings"
)

//func main() {
//	comp, err := pkg.GetMainInfo("5258081758")
//	if err != nil {
//		log.Panicln(err)
//	}
//
//	kpp, err := pkg.GetCompanyKPP(comp)
//	if err != nil {
//		log.Panicln(err)
//	}
//
//	fmt.Printf(
//		"ИНН %s\n"+
//			"КПП: %s\n"+
//			"Название: %s\n"+
//			"ФИО Рук.:%s\n", comp.INN, kpp, comp.Name, comp.FIO)
//}

const (
	port     = 8888
	httpPort = 8080
)

func main() {
	go startHTTP()

	startGRPC()
}

func registerSwaggerUI(mux *http.ServeMux) {

	mux.HandleFunc("/swaggerui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/openapiv2/api.swagger.json")
	})

	mux.HandleFunc("/swaggerui/", func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, "/swaggerui/")
		if p == "" {
			http.ServeFile(w, r, path.Join("./static", "index.html"))
		} else {
			p = path.Join("./static", p)
			http.ServeFile(w, r, p)
		}
	})

	//mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(2)
	//	p := strings.TrimPrefix(r.URL.Path, "/static")
	//	p = path.Join("./static", p)
	//	http.ServeFile(w, r, p)
	//})
}

func startHTTP() {
	mux := http.NewServeMux()

	gw, err := registerGatewayEndpoints()
	if err != nil {
		log.Fatalf("error while registering gateway ep: %v", err)
	}

	registerSwaggerUI(mux)

	mux.Handle("/", gw)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux); err != nil {
		log.Fatalf("HTTP server stopped with error: %v", err)
	}
}

func registerGatewayEndpoints() (http.Handler, error) {
	h := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pkg.RegisterScraperHandlerFromEndpoint(context.Background(), h, fmt.Sprintf(":%d", port), opts); err != nil {
		return nil, fmt.Errorf("error while registering handler: %w", err)
	}

	return h, nil
}

func startGRPC() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("couldn't start gRPC server: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pkg.RegisterScraperServer(
		grpcServer,
		pkg.MyScraper{},
	)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server stopped with error: %v", err)
	}
}
