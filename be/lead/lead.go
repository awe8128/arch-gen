package lead

import "context"

type Lead struct {
	Id *uint
}

func NewLead(ctx context.Context, id *uint) *Lead {

	ent := &Lead{
		Id: id,
	}

	return ent
}
