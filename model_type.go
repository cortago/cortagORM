package model

import (
	"time"
)

// AutoField is an IntegerField that automatically increments according to available IDs
type AutoField uint32

// BigAutoField is much like AutoField but for 64 bit integer  
type BigAutoField uint64

// BigIntegerField is a 64 bit integer
type BigIntegerField int64

// BooleanField is a true/false field 
type BooleanField bool

// CharField is a string field
type CharField string

// TimeField is time.Time instance of golang time package
type TimeField time.Time

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
