package block

type block struct {
	Data         string
	Hash         string
	PreviousHash string
}

func New(data, hash, previousHash string) *block {
	return &block{
		Data:         data,
		Hash:         hash,
		PreviousHash: previousHash,
	}
}
