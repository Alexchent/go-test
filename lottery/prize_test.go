package lottery

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"
)

//go:embed red_packet_rain.json
var giftConf string

// go test -v prize_test.go lottery.go
func TestCloseAnGame(t *testing.T) {
	gifts := &GiftList{}
	json.Unmarshal([]byte(giftConf), &gifts)
	ret := gifts.sort()

	// 开始抽奖
	res := closeAnGame(10, 3, ret)
	if len(res) > 3 {
		t.Error("no cool")
	}
	fmt.Println(res)
}
