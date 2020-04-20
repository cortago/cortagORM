package model

// Model is basic structure which can be embedded
// For eg:-
// type Person struct {
//		cortagORM.Model	
// } 
type Model struct {
	ID        AutoField
	CreatedAt TimeField
	UpdatedAt TimeField
	DeletedAt TimeField
}
