package health

type Service interface {
	CheckHealth() Status
} 