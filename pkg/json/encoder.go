package json

import (
		"crypto/md5"
		"fmt"
		"github.com/go-kratos/kratos/v2/log"
		json "github.com/json-iterator/go"
		"io"
)

func Encode(v interface{}) string {
		var data, err = json.Marshal(v)
		if err != nil {
				log.Errorw("function", "JsonEncode","Error", err)
				return ``
		}
		return string(data)
}

func Beautify(v interface{}) string {
		var data, err = json.MarshalIndent(v, "", `  `)
		if err != nil {
				log.Errorw("function", "Beautify","Error", err)
				return ``
		}
		return string(data)
}

func Bytes(v interface{}) ([]byte, error) {
		return json.Marshal(v)
}

func Decode(data []byte, v interface{}) error {
		return json.Unmarshal(data, v)
}

func DecodeReader(data io.Reader, v interface{}) error {
		return json.NewDecoder(data).Decode(v)
}

func HashCode(v interface{}) string {
		var data = Encode(v)
		if data == "" {
				return ``
		}
		var hashCodec = md5.New()
		hashCodec.Write([]byte(data))
		return fmt.Sprintf("%x", hashCodec.Sum(nil))
}

func CopyBind(src interface{}, dst interface{}) error {
		bytes, err := Bytes(src)
		if err != nil {
				return err
		}
		return Decode(bytes, dst)
}
