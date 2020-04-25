// Migrations is a struct that consists of all informations helpful to migrate database   
type Migrations struct {
	operations []Operation
}

// Operation is a common interface for all database functionalites such as create, delete tables or 
// create, delete, modify fields
type Operation interface {
	// Add defines a function when an operation is to be added on present databases or 
	// simply for migrate add
	Add( database string )

	// Remove defines a function when an operation is to be removed from present database or 
	// simple for migrate subtract
	Remove(database string)
}

// Field consist of all information of their respective fields
type Field struct {
	// Name of the field
	Name string
	// Type of the ORM
	OrmType ORMType
	// Properties consist of meta information for fields 
	Properties map[string]string
}

// CreateModel is an opertion to create table with certain fields provided  
type CreateModel struct {
	Name string
	Fields []Field
}

// TODO: Add `Add` and `Remove` functions for CreateModel struct to make elgible
// for Operation interface
func (*CreateModel) Add( database string ) {}
func (*CreateModel) Remove( database string ) {}

// AddField is an operation to modify table and add a new field 
type AddField struct {
	ModelName string
	Field Field
}

// TODO: Add `Add` and `Remove` functions for AddField struct to make elgible
// for Operation interface
func (*AddField) Add( database string ) {}
func (*AddField) Remove( database string ) {}

// RemoveModel is an operation to remove a Model from database this command
// cannot be reversed So Remove function should through error
type RemoveModel struct {
	ModelName string
}

// TODO: Add `Add` and `Remove` functions for RemoveModel struct to make elgible
// for Operation interface
func (*RemoveModel) Add( database string ) {}
func (*RemoveModel) Remove( database string ) {}

// RemoveField is an operation to remove a field from corresponding Model this command
// cannot be reversed So Remove function should through error
type RemoveField struct {
	ModelName string
	FieldName string
}

// TODO: Add `Add` and `Remove` functions for RemoveField struct to make elgible
// for Operation interface
func (*RemoveField) Add( database string ) {}
func (*RemoveField) Remove( database string ) {}
