package teemap_test

import (
	"testing"

	"github.com/BenjaminB64/go-teemap"
)

func TestTeeMap(t *testing.T) {
	teemap := teemap.New()
	t.Log(teemap)
}
