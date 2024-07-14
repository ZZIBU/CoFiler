package rpc

import (
	"CoFiler/config"
	cofiler "CoFiler/rpc/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CofilerClient struct {
	client           *grpc.ClientConn
	tagServiceClient cofiler.TagServiceClient
}

func NewCofilerClient(config *config.Config) (*CofilerClient, error) {
	c := &CofilerClient{}

	// TODO : 서비스 도메인, 네임스페이스, 포트 magic string 및 number 수정하기
	address := fmt.Sprintf("%s.%s:%d", "tagify", "tagify", 8080)

	if client, err := grpc.NewClient(
		address,
		// TODO : Client 옵션 알아보고 추가하기 ex) timeout
		grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.client = client
		c.tagServiceClient = cofiler.NewTagServiceClient(c.client)
	}

	return c, nil
}
