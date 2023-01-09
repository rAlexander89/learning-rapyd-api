package rapyd

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	rapydUtils "github.com/rAlexander89/rapyd-banking/rapyd_utils"
)

type SignatureData struct {
	Method    string
	Path      string
	Salt      string
	Timestamp string
	AccessKey string
	SecretKey string
	Body      string
}

type signer struct {
	accessKey []byte
	secretKey []byte
}

type Signer interface {
	secretKeyString() string
	accessKeyString() string
	signRequest(r *http.Request, body []byte) error
}

const (
	SaltHeader      = "salt"
	TimestampHeader = "timestamp"
	AccessKeyHeader = "access_key"
	SignatureHeader = "signature"

	ContentTypeHeader  = "Content-Type"
	DefaultContentType = "application/json"
)

func NewSigner(accessKey []byte, secretKey []byte) Signer {
	return &signer{
		accessKey: accessKey,
		secretKey: secretKey,
	}
}

func (s *signer) signRequest(r *http.Request, body []byte) error {
	salt, err := rapydUtils.GenSalt()
	if err != nil {
		return errors.New("error getting salt")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	r.Header.Add(AccessKeyHeader, s.accessKeyString())
	r.Header.Add(SaltHeader, salt)
	r.Header.Add(TimestampHeader, timestamp)
	r.Header.Add(ContentTypeHeader, DefaultContentType)

	data := SignatureData{
		Method:    strings.ToLower(r.Method),
		Path:      r.URL.RequestURI(),
		Salt:      salt,
		Timestamp: timestamp,
	}

	if len(body) != 0 {
		data.Body = string(body)
	}

	signature := s.signData(data)
	r.Header.Add(SignatureHeader, base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(signature))))
	return nil
}

func (s *signer) secretKeyString() string {
	return string(s.secretKey)
}

func (s *signer) accessKeyString() string {
	return string(s.accessKey)
}

func (s *signer) signData(data SignatureData) []byte {
	if data.AccessKey == "" {
		data.AccessKey = s.accessKeyString()
	}

	if data.SecretKey == "" {
		data.SecretKey = s.secretKeyString()
	}

	toSign := data.Method + data.Path + data.Salt + data.Timestamp + data.AccessKey + data.SecretKey + data.Body

	h := hmac.New(sha256.New, s.secretKey)
	h.Write([]byte(toSign))
	return h.Sum(nil)
}

func (c *RapydClient) Sign(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.Resolve(path), nil)
	if err != nil {
		return nil, errors.New("unable to form a new request")
	}

	err = c.signRequest(req, nil)
	if err != nil {
		return nil, errors.New("error signing request")
	}
	res, err := c.Do(req)

	if err != nil {
		return nil, errors.New("error sending request")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errorResponse, _ := io.ReadAll(res.Body)
		// string formatter: %d == decimal integer, %s == string
		return nil, fmt.Errorf("error: status code %d, response %s", res.StatusCode, string(errorResponse))
	}
	print("res body: ", res.Body)

	return io.ReadAll(res.Body)
}
