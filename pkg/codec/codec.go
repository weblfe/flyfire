package codec

var (
		codecs =make(map[string]Codec)
)

type Codec interface {
		Encode(interface{})([]byte,error)
		Decode([]byte,interface{}) error
}

func Register(name string,codec Codec)  {
		codecs[name] = codec
}

func GetCodec(name string) (Codec,bool) {
		c,ok:=codecs[name]
		return c,ok
}

func init()  {
		Register("json",jsonCodec{})
}