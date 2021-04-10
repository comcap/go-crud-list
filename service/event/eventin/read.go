package eventin

type ReadInput struct {
	EventID string `json:"-" validate:"required"`
}

//func MakeTestReadInput() (input *ReadInput) {
//	return &ReadInput{
//		EventID: "test",
//	}
//}
