package slack

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Signature headers
const (
	hSignature = "X-Slack-Signature"
	hTimestamp = "X-Slack-Request-Timestamp"
)

// SecretsVerifier contains the information needed to verify that the request comes from Slack
type SecretsVerifier struct {
	signature []byte
	hmac      hash.Hash
}

func unsafeSignatureVerifier(header http.Header, secret string) (_ SecretsVerifier, err error) {
	var (
		bsignature []byte
	)

	signature := header.Get(hSignature)
	stimestamp := header.Get(hTimestamp)

	if signature == "" || stimestamp == "" {
		return SecretsVerifier{}, errors.New("missing headers")
	}

	if bsignature, err = hex.DecodeString(strings.TrimPrefix(signature, "v0=")); err != nil {
		return SecretsVerifier{}, err
	}

	hash := hmac.New(sha256.New, []byte(secret))
	if _, err = hash.Write([]byte(fmt.Sprintf("v0:%s:", stimestamp))); err != nil {
		return SecretsVerifier{}, err
	}

	return SecretsVerifier{
		signature: bsignature,
		hmac:      hash,
	}, nil
}

// NewSecretsVerifier returns a SecretsVerifier object in exchange for an http.Header object and signing secret
func NewSecretsVerifier(header http.Header, secret string) (sv SecretsVerifier, err error) {
	var (
		timestamp int64
	)

	stimestamp := header.Get(hTimestamp)

	if sv, err = unsafeSignatureVerifier(header, secret); err != nil {
		return SecretsVerifier{}, err
	}

	if timestamp, err = strconv.ParseInt(stimestamp, 10, 64); err != nil {
		return SecretsVerifier{}, err
	}

	diff := absDuration(time.Since(time.Unix(timestamp, 0)))
	if diff > 5*time.Minute {
		return SecretsVerifier{}, fmt.Errorf("timestamp is too old")
	}

	return sv, err
}

func (v *SecretsVerifier) Write(body []byte) (n int, err error) {
	return v.hmac.Write(body)
}

// Ensure compares the signature sent from Slack with the actual computed hash to judge validity
func (v SecretsVerifier) Ensure() error {
	computed := v.hmac.Sum(nil)
	// use hmac.Equal prevent leaking timing information.
	if hmac.Equal(computed, v.signature) {
		return nil
	}

	return fmt.Errorf("Expected signing signature: %s, but computed: %s", v.signature, computed)
}

func abs64(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func absDuration(n time.Duration) time.Duration {
	return time.Duration(abs64(int64(n)))
}
