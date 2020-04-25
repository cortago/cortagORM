package model

// ModelBasic is basic structure which can be embedded
// For eg:-
// type Person struct {
//		cortagORM.ModelBasic
// }
type ModelBasic struct {
	ID        AutoField
	CreatedAt TimeDateField
	UpdatedAt TimeDateField
	DeletedAt TimeDateField
}

// Model is a common interface for all newly defined models in the framework
// for any newly defined model it should follow this interface.
type Model interface {
	// ModelToStrng defines a function which converts any model object into string,
	ModelToString() string
}
