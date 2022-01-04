package teemap_test

import (
	"testing"
	"math/rand"
	"time"
	"github.com/BenjaminB64/go-teemap"
)

func TestTeeMap(t *testing.T) {
	teemap := teemap.New[string, string]()
	teemap.Set("key1", "value1")

	if teemap.Get("key1") != "value1" {
		t.Fatal()
	}
}

func FuzzStringTeeMap(f *testing.F) {
	teemap := teemap.New[string, string]()
	for	i := 0; i < 500; i++ {
		f.Add(RandStringBytesMaskImprSrc(20), RandStringBytesMaskImprSrc(20))
	}

	f.Fuzz(func(t *testing.T, fuzzKey string, fuzzValue string) {
		t.Parallel()
		teemap.Set(fuzzKey, fuzzValue)

		if v := teemap.Get(fuzzKey); v != fuzzValue {
			t.Fatalf("Value doesn't match %s != %s", v, fuzzValue)
		}
	})
}


// Generate random string (https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go)
var src = rand.NewSource(time.Now().UnixNano())
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}
