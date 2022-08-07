package protocol

import (
	"bytes"
	handler "github.com/captainjonas/godisson/pkg/godisson/client/handler"
)

type Decoder interface {
	decode(buf bytes.Buffer, state handler.State) error
}
