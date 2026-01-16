package utils

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}

func JsonStatus(message string) []byte {
	m := map[string]string{"message": message}
	b, _ := json.Marshal(m)
	return b
}
