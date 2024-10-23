package words

import (
	context "context"
	"dtclients/config"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client WordsClient

func init() {
	conn, err := grpc.NewClient(
		config.Cfg.Words.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("connect words client failed: %v", err)
	}
	defer conn.Close()
	client = NewWordsClient(conn)
}

func CheckSensitive(text string) (*CheckSensitiveResponse, error) {
	resp, err := client.CheckSensitive(context.Background(), &CheckSensitiveRequest{Text: text})
	if err != nil {
		log.Fatalf("could not check sensitive: %v", err)
	}
	return resp, nil
}