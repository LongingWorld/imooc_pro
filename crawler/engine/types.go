package engine

type Request struct {
	URL string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	URL string
	Id string
	Type string
	Payload interface{}
}


func NilParserFunc([]byte) ParserResult{
	return ParserResult{}
}