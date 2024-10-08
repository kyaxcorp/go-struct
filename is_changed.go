package _struct

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
)

func getBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Checksum(data interface{}) (string, error) {
	bytes, err := getBytes(data)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:]), nil
}

func IsChanged(data interface{}, prevHash string) (changed bool, newHash string, err error) {
	newHash, err = Checksum(data)
	if err != nil {
		return
	}
	if newHash != prevHash {
		changed = true
		return
	}
	return
}
