package main

import (
	"context"
	"fmt"
	"github.com/toasterson/graphql-go-pets-example/helloworld"
	"google.golang.org/grpc"
	//_ "google.golang.org/grpc/resolver/passthrough"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	graphql_go_pets_example "github.com/toasterson/graphql-go-pets-example"
)

const defaultPort = "8080"

type helloServer struct {
}

func (h *helloServer) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if req.Name != "" {
		return &helloworld.HelloReply{
			Message: fmt.Sprintf("Hello %s!", req.Name),
		}, nil
	}
	return &helloworld.HelloReply{
		Message: "Hello World!",
	}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := os.Remove("/tmp/example.sock"); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	lis, err := net.Listen("unix", "/tmp/example.sock")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, &helloServer{})
	go grpcServer.Serve(lis)

	go func() {
		conn, err := grpc.Dial("passthrough:///unix:///tmp/example.sock", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := helloworld.NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "Toasty"})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Message)
	}()

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql_go_pets_example.NewExecutableSchema(graphql_go_pets_example.Config{Resolvers: &graphql_go_pets_example.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
