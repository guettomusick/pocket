package types

type Request struct {
	nonce uint32
	ch    chan Work
}

func (r *Request) Nonce() uint32 {
	return r.nonce
}

func (r *Request) Response() <-chan Work {
	return r.ch
}
