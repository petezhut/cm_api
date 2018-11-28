package services

type Service interface {
	Start()
	Stop()
	Restart()
}

