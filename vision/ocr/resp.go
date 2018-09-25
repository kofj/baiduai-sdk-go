package ocr

// Resp API 响应
type Resp struct {
	bytes []byte
}

// Bytes 返回 bytes
func (r *Resp) Bytes() []byte {
	return r.bytes
}
