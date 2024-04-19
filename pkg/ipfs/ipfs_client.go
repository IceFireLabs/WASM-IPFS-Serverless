package ipfs

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type IPFSClient struct {
	scheme          string
	host            string
	port            int
	LassieClientURL string
}

func NewIPFSClient(scheme string, host string, port int) *IPFSClient {

	return &IPFSClient{
		scheme:          scheme,
		host:            host,
		port:            port,
		LassieClientURL: fmt.Sprintf("%s://%s:%d/ipfs", scheme, host, port),
	}

}

func (ipfsC *IPFSClient) GetURLFromCID(cid string) (cidUrl string, err error) {
	cidUrl = fmt.Sprintf("%s/%s", ipfsC.LassieClientURL, cid)

	return cidUrl, err
}

func (ipfsC *IPFSClient) GetDataFromCID(cid string) (data []byte, err error) {
	cidUrl := fmt.Sprintf("%s/%s", ipfsC.LassieClientURL, cid)

	req, err := http.NewRequest("GET", cidUrl, nil)
	req.Header.Set("Accept", "*/*")
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 20}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return
}
