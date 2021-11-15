package aliclient

import (
	"errors"
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

const RegionId = "RegionId"

// Client def .
type Client struct {
	*alidns.Client
}

// NewAliClient 创建阿里客户端
func NewAliClient() (*Client, error) {
	accessKeyID := os.Getenv("ALI_ACCESSKEYID")
	accessKeySecret := os.Getenv("ALI_ACCESSKEY_SECRET")
	if accessKeyID == "" {
		return nil, errors.New("plase exec: export ALI_ACCESSKEYID = 阿里云AccessKeyID")
	}
	if accessKeySecret == "" {
		return nil, errors.New("plase exec: export ALI_ACCESSKEY_SECRET = 阿里云AccessKeySecret")
	}
	client, err := alidns.NewClientWithAccessKey(RegionId, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("new aliclient error:%s", err.Error())
	}
	return &Client{client}, err
}
