package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/banks/proto-play/playpb"
	"github.com/golang/protobuf/proto"
	"github.com/mr-tron/base58/base58"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative playpb/play.proto

func main() {
	t := playpb.Token{
		ClusterId:      "project/df4f087f-a357-4da7-b0f8-43e754a7fb9a/hashicorp.waypoint.entrypoint/web",
		SpClientId:     []byte("J5Cy3RQOwuUd98ogRUDBQA5nGrLn3Pcp"),
		SpClientSecret: []byte("tznTdErfJpj7F2W3g4RAPly68AkZdhOBPEcZXXu9pBLSmFMsbYvajyxDFj3HkKds"),
	}

	rawBs, err := proto.Marshal(&t)
	if err != nil {
		log.Fatal(err)
	}

	encoded := base58.Encode(rawBs)

	fmt.Println("=> Basic base58-proto encoding")
	fmt.Printf("%s (%d bytes)\n\n", encoded, len(encoded))

	rawID, err := base64.StdEncoding.DecodeString(string(t.SpClientId))
	if err != nil {
		log.Fatal(err)
	}
	rawS, err := base64.StdEncoding.DecodeString(string(t.SpClientSecret))
	if err != nil {
		log.Fatal(err)
	}

	t2 := playpb.Token{
		ClusterId:      "project/df4f087f-a357-4da7-b0f8-43e754a7fb9a/hashicorp.waypoint.entrypoint/web",
		SpClientId:     rawID,
		SpClientSecret: rawS,
	}

	rawBs, err = proto.Marshal(&t2)
	if err != nil {
		log.Fatal(err)
	}

	encoded = base58.Encode(rawBs)

	fmt.Println("=> base58-proto encoding with Oauth decoded to raw bytes")
	fmt.Printf("%s (%d bytes)\n\n", encoded, len(encoded))

	t3 := playpb.Token2{
		//ClusterId:      "project/df4f087f-a357-4da7-b0f8-43e754a7fb9a/hashicorp.waypoint.entrypoint/web",
		ProjectId:      uuid2bytes("df4f087f-a357-4da7-b0f8-43e754a7fb9a"),
		ResourceType:   playpb.ResourceType_WP_ENTRYPOINT,
		ResourceId:     "web",
		SpClientId:     rawID,
		SpClientSecret: rawS,
	}

	rawBs, err = proto.Marshal(&t3)
	if err != nil {
		log.Fatal(err)
	}

	encoded = base58.Encode(rawBs)

	fmt.Println("=> base58-proto encoding with raw oauth and raw location")
	fmt.Printf("%s (%d bytes)\n\n", encoded, len(encoded))
}

func uuid2bytes(uuid string) []byte {
	hexStr := strings.ReplaceAll(uuid, "-", "")
	dec, err := hex.DecodeString(hexStr)
	if err != nil {
		log.Fatal(err)
	}
	return dec
}

// output:
// 27ot46j87D5pqfDhj7pQFjovV65EYEtR68Dk2n9TxdrYCqT1Yh6tKGUDE1Z4c1oTcA9fx6CHdnfu\
// oanCMivPLpYatX5EHvE3oQiwCfur8qwWFwFgKihkcK9J5mK3J4VGve3udwTP39dvcHhdSAAq7Q6q\
// ZwKRxCnN3fE83DrMyjaR3ivtF6Rfgkq3PXvL4911CLxGv9yghXYwVqwjVz24k7AoEPehBPPcPYz5\
// 3SZqw9NqU5hVunKc2J (246 bytes)
