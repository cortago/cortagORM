package model

import "time"

// ORMType is a generic interface for all ORM Types
type ORMType interface {
	// FieldDataType returns data type for any corresponding database.
	FieldDataType(database string) string

	// ormType ensures that no other type accidentally implements the interface
	ormType()
}

// AutoField is an IntegerField that automatically increments according to available IDs
type AutoField uint32

// BigAutoField is much like AutoField but for 64 bit integer
type BigAutoField uint64

// BigIntegerField is a 64 bit integer
type BigIntegerField int64

// BooleanField is a true/false field
type BooleanField bool

// CharField is a string field. Default limit is 255 characters
type CharField string

// TimeField is time.Time instance of golang time package
type TimeDateField time.Time

// DecimalField is a decimal number
type DecimalField float64

// DurationField is time.Duration instance of golang time package
type DurationField time.Duration

// EmailField is a CharField with email validations(TODO: Add Email Validations)
type EmailField string

// IntegerField is a 32 instance bit integer
type IntegerField int32

// PositiveIntegerField is a 32 bit unsigned integer
type PositiveIntegerField uint32

// PositiveSmallIntegerField is a 8 bit unsigned integer
type PositiveSmallIntegerField uint8

// SmallIntegerField is a 8 bit integer
type SmallIntegerField int8

// TextField is for long text
type TextField string

func (*AutoField) ormType()                 {}
func (*BigAutoField) ormType()              {}
func (*BigIntegerField) ormType()           {}
func (*BooleanField) ormType()              {}
func (*CharField) ormType()                 {}
func (*DecimalField) ormType()              {}
func (*DurationField) ormType()             {}
func (*EmailField) ormType()                {}
func (*IntegerField) ormType()              {}
func (*PositiveIntegerField) ormType()      {}
func (*PositiveSmallIntegerField) ormType() {}
func (*SmallIntegerField) ormType()         {}
func (*TextField) ormType()                 {}
