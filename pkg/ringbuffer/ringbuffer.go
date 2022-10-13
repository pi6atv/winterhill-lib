package ringbuffer

import (
	"container/ring"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type Ring struct {
	ring *ring.Ring
}

func New(size int) *Ring {
	return &Ring{ring: ring.New(size)}
}

func (R *Ring) Add(a any) {
	R.ring.Value = a
	R.ring = R.ring.Next()
}

func (R *Ring) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(R.ring.Value)
	if err != nil {
		return nil, errors.Wrapf(err, "marshalling %+v", R.ring.Value)
	}
	out := fmt.Sprintf("[%s", buf)

	for p := R.ring.Next(); p != R.ring; p = p.Next() {
		buf, err = json.Marshal(p.Value)
		if err != nil {
			return nil, errors.Wrapf(err, "marshalling %+v", R.ring.Value)
		}
		out = fmt.Sprintf("%s,%s", out, buf)
	}
	out = fmt.Sprintf("%s]", out)
	return []byte(out), nil
}
