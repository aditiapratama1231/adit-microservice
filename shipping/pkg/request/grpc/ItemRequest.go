package request

import (
	"context"
	"log"

	pb "github.com/aditiapratama1231/adit-microservice/proto/item"
	"google.golang.org/grpc"
)

func ShowItemDetail(id int64) (*pb.ShowItemDetailResponse, error) {
	port := ":8082"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewItemsClient(conn)
	itemId := &pb.ShowItemDetailRequest{
		ItemId: id,
	}

	return client.ShowItemDetail(context.Background(), itemId)
}
