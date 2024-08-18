package recommendation

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/proto/recommendation"
	"google.golang.org/grpc"
)

type Client struct {
    conn   *grpc.ClientConn
    client pb.RecommendationServiceClient
}

func NewClient(address string) (*Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, fmt.Errorf("failed to connect to recommendation service: %v", err)
    }

    client := pb.NewRecommendationServiceClient(conn)
    return &Client{conn: conn, client: client}, nil
}

func (c *Client) GetRecommendations(ctx context.Context, userID string, articleID uint64) ([]uint64, error) {
    resp, err := c.client.GetRecommendations(ctx, &pb.RecommendationRequest{
        UserId:    userID,
        ArticleId: articleID,
    })
    if err != nil {
        return nil, fmt.Errorf("error getting recommendations: %v", err)
    }
    fmt.Printf("Received recommendations: %v\n", resp.ProductIds)
    return resp.ProductIds, nil
}

func (c *Client) Close() {
    if err := c.conn.Close(); err != nil {
        log.Printf("Error closing gRPC connection: %v", err)
    }
}