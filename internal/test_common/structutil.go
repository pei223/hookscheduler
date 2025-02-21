package test_common

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func DeepCopy[T any](src T) (*T, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}

	var dst T
	err = json.Unmarshal(b, &dst)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	return &dst, nil
}
