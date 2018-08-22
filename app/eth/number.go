package eth

import (
	"encoding/json"
	"strconv"
	"time"
)

type (
	ethInt struct {
		Value int64
	}

	ethTimestamp struct {
		Value time.Time
	}
)

func parseInt(data []byte) (int64, error) {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return -1, err
	}

	if len(str) == 0 {
		return -1, nil
	}

	return strconv.ParseInt(str, 0, 64)
}

func (e *ethInt) UnmarshalJSON(data []byte) error {
	i, err := parseInt(data)
	if err != nil {
		return err
	}
	e.Value = i
	return nil
}

func (e *ethTimestamp) UnmarshalJSON(data []byte) error {
	i, err := parseInt(data)
	if err != nil {
		return err
	}
	e.Value = time.Unix(i, 0)
	return nil
}

func ethNumFromInt(num int64) string {
	return "0x" + strconv.FormatInt(num, 16)
}

// func ethNumFromBigInt(num *big.Int) string {
// 	return "0x" + strconv.FormatInt(num, 16)
// }
