package eventin

type DeleteInput struct {
	ID string `json:"-" validate:"required"`
}

//func MakeTestDeleteInput() (input *DeleteInput) {
//	return &DeleteInput{
//		ID: "test",
//	}
//}