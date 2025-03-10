package health

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) CheckHealth() Status {
	return Status{
		Status:  "success",
		Message: "Hello World from Differ Template API!",
	}
} 