package rapyd

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/pkg/errors"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	SaltLen         = 12
	SaltHeader      = "salt"
	TimestampHeader = "timestamp"
	AccessKeyHeader = "access_key"
	SignatureHeader = "signature"

	ContentTypeHeader  = "Content-Type"
	DefaultContentType = "application/json"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type SignatureData struct {
	Method    string
	Path      string
	Salt      string
	Timestamp string
	AccessKey string
	SecretKey string
	Body      string
}

func (c *client) getRandomRune() (*rune, error) {
	lettersLen := len(letters)

	index, err := rand.Int(rand.Reader, new(big.Int).SetInt64(int64(lettersLen)))
	if err != nil {
		return nil, errors.Wrap(err, "error picking random number")
	}

	return &letters[index.Int64()], nil
}

// GetSalt for generating salt for certain length
func (c *client) getSalt(len int) (string, error) {
	salt := make([]rune, len)

	for i := range salt {
		r, err := c.getRandomRune()
		if err != nil {
			return "", errors.New("error getting next rune")
		}

		salt[i] = *r
	}
	return string(salt), nil
}

func (c *client) signData(data SignatureData) []byte {
	if data.AccessKey == "" {
		data.AccessKey = string(c.accessKey)
	}

	if data.SecretKey == "" {
		data.SecretKey = string(c.secretKey)
	}

	toSign := data.Method + data.Path + data.Salt + data.Timestamp + data.AccessKey + data.SecretKey + data.Body

	h := hmac.New(sha256.New, c.secretKey)
	h.Write([]byte(toSign))
	return h.Sum(nil)
}

func (c *client) signRequest(r *http.Request, body []byte) error {
	salt, err := c.getSalt(SaltLen)
	if err != nil {
		return errors.Wrap(err, "error getting salt")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	r.Header.Add(AccessKeyHeader, string(c.accessKey))
	r.Header.Add(SaltHeader, salt)
	r.Header.Add(TimestampHeader, timestamp)
	r.Header.Add(ContentTypeHeader, DefaultContentType)

	data := SignatureData{
		Method:    strings.ToLower(r.Method),
		Path:      r.URL.Path,
		Salt:      salt,
		Timestamp: timestamp,
	}

	if len(body) != 0 {
		data.Body = string(body)
	}

	signature := c.signData(data)
	r.Header.Add(SignatureHeader, base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(signature))))
	return nil
}
