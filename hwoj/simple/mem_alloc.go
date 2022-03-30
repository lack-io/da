package simple

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Memory struct {
	Val int64
	Num int64
}

func (m *Memory) String() string {
	return fmt.Sprintf("{%d:%d}", m.Val, m.Num)
}

func MemoryAlloc() ([]bool, error) {
	row1 := ""
	_, err := fmt.Scanln(&row1)
	if err != nil {
		return nil, err
	}

	row2 := ""
	_, err = fmt.Scanln(&row2)
	if err != nil {
		return nil, err
	}

	return apply(getMemPool(row1), getApplyMem(row2)), nil
}

func getMemPool(s string) []*Memory {
	mems := make([]*Memory, 0)
	parts := strings.Split(s, ",")
	for _, part := range parts {
		item := strings.TrimSpace(part)
		memSp := strings.Split(item, ":")
		if len(memSp) != 2 {
			continue
		}
		val, _ := strconv.ParseInt(memSp[0], 10, 64)
		if val == 0 {
			continue
		}
		num, _ := strconv.ParseInt(memSp[1], 10, 64)
		if num == 0 {
			continue
		}
		mems = append(mems, &Memory{val, num})
	}

	sort.Slice(mems, func(i, j int) bool {
		return mems[i].Val < mems[j].Val
	})

	return mems
}

func getApplyMem(s string) []int64 {
	out := make([]int64, 0)
	parts := strings.Split(s, ",")
	for _, part := range parts {
		item := strings.TrimSpace(part)
		n, _ := strconv.ParseInt(item, 10, 64)
		if n == 0 {
			continue
		}
		out = append(out, n)
	}

	return out
}

func apply(pool []*Memory, mems []int64) []bool {
	//fmt.Println(pool, mems)
	result := make([]bool, len(mems))
	for i, mem := range mems {
		ok := false
		for _, p := range pool {
			if p.Val >= mem && p.Num > 0 {
				ok = true
				p.Num -= 1
				break
			}
		}
		result[i] = ok
	}

	return result
}
