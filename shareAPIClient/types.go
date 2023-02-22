package shareAPIClient

type QueryStringParameters interface{}
type MultiValueQueryStringParameters interface{}

type Req struct {
	Path                            string                          `json:"path"`
	HttpMethod                      string                          `json:"httpMethod"`
	QueryStringParameters           QueryStringParameters           `json:"queryStringParameters,omitempty"`
	MultiValueQueryStringParameters MultiValueQueryStringParameters `json:"multiValueQueryStringParameters,omitempty"`
}

type GetShareParam struct {
	PublicKey string `json:"publicKey"`
	Msg       string `json:"msg"`
	SignedMsg string `json:"signedMsg"`
}

type SetupMultiValParam struct {
	PkList []string `json:"pkList"`
}

type SetupParam struct {
	N         string `json:"n"`
	T         string `json:"t"`
	Msg       string `json:"msg"`
	SignedMsg string `json:"signedMsg"`
}

type GetShareRespBody struct {
	Pk       string `json:"pk"`
	EncShare string `json:"encShare"`
	Index    string `json:"index"`
}

type GetShareResp struct {
	Body GetShareRespBody `json:"body"`
}

type Response struct {
	Body string `json:"body"`
}
