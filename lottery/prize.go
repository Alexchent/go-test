package lottery

import (
	_ "embed"
	"math/rand"
)

type GiftType = string

const (
	THANKS GiftType = "thanks"
	VIP    GiftType = "vip"
	COIN   GiftType = "coin"
	COUPON GiftType = "coupon"
)

// Gift 奖品
type Gift struct {
	Title  string
	Type   GiftType
	Pic    string
	Weight float32 `json:"weight"` // 权重
	Index  int     // 位置，用于计算中奖区间
}

type GiftList []Gift

//func main() {
//	gifts := &GiftList{}
//	json.Unmarshal([]byte(giftConf), &gifts)
//	ret := gifts.sort()
//
//	// 开始抽奖
//	res := closeAnGame(10, 3, ret)
//	fmt.Println(res)
//}

func closeAnGame(num, max int, gifts GiftList) (ret GiftList) {
	winn := 0
	for i := 0; i < num; i++ {
		if winn == max {
			return
		}
		item := gifts.PrizeDraw()
		if item.Type == THANKS {
			continue
		}
		winn++
		ret = append(ret, *item)
	}
	return
}

func (l GiftList) PrizeDraw() *Gift {
	randNum := rand.Intn(l[len(l)-1].Index) + 1
	//fmt.Println(randNum)
	for _, gift := range l {
		if randNum <= gift.Index {
			return &gift
		}
	}
	return nil
}

// 计算奖品的中奖区间，同时对奖品排序
func (l GiftList) sort() GiftList {
	ret := make(GiftList, len(l)) // 声明切片是就确定好长度，避免append扩容的性能损耗
	index := 0
	for i, gift := range l {
		index += int(gift.Weight * 10)
		gift.Index = index
		ret[i] = gift
	}
	return ret
}
