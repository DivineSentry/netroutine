package netroutine

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
)

const (
	idHashSHA512 = "HashSHA512"
)

func init() {
	blocks[idHashSHA512] = &HashSHA512{}
}

type HashSHA512 struct {
	FromKey string
	ToKey   string
}

func (b *HashSHA512) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *HashSHA512) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *HashSHA512) kind() string {
	return idHashSHA512
}

func (b *HashSHA512) Run(ctx context.Context, wce *Environment) (string, Status) {
	hasher := sha512.New()

	sv, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	s, err := toString(sv)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	hasher.Write([]byte(s))

	sha512s := hex.EncodeToString(hasher.Sum(nil))

	wce.setData(b.ToKey, sha512s)

	return log(b, setWorkingData(b.ToKey, sha512s), Success)
}
