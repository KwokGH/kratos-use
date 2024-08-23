package unique

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewID() string {
	return primitive.NewObjectID().Hex()
}

func GetMd5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetFileSha1(file io.Reader) (string, error) {
	_sha1 := sha1.New()
	if _, err := io.Copy(_sha1, file); err != nil {
		return "", err
	}
	sha1 := hex.EncodeToString(_sha1.Sum(nil))

	return sha1, nil
}
