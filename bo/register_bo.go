package bo

type RegisterBO struct {
	userName string
	password string
	name     string
	age      int
}

func (bo *RegisterBO) New() *RegisterBO {
	return &RegisterBO{}
}
