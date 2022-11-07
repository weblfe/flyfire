package codec

import "githu.com/weblfe/flyfire/pkg/json"

type jsonCodec struct {}

func (jsonCodec)Encode(v interface{})([]byte,error)  {
		return json.Bytes(v)
}

func (jsonCodec)Decode(data []byte,returnValue interface{}) error  {
		return json.Decode(data,returnValue)
}
