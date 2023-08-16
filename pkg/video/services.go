package video

type Service interface{}

type videoSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &videoSvc{repo: r}
}
