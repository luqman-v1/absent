package repo

type GC struct {
	StartDate string
	EndDate   string
	Services  IServices
}

func (r *GC) ListEvent() ([]byte, error) {
	body, err := r.Services.GetListEvents(r.StartDate, r.EndDate)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func NewGC() *GC {
	return &GC{
		Services: NewServices(),
	}
}
