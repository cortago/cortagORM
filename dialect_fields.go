package model

import "log"

const (
	mysql      = "mysql"
	postgresql = "postgresql"
)

func (*AutoField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "INT"
	case postgresql:
		return "integer"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*BigAutoField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "BIGINT"
	case postgresql:
		return "bigint"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*BigIntegerField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "BIGINT"
	case postgresql:
		return "bigint"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*BooleanField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "BOOLEAN"
	case postgresql:
		return "boolean"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*CharField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "VARCHAR(255)"
	case postgresql:
		return "varchar(255)"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*TimeDateField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "DATETIME"
	case postgresql:
		return "timestamp"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*DecimalField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "DECIMAL"
	case postgresql:
		return "decimal"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*DurationField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "BIGINT"
	case postgresql:
		return "bigint"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*EmailField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "VARCHAR(255)"
	case postgresql:
		return "varchar(255)"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*IntegerField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "INT"
	case postgresql:
		return "integer"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*PositiveIntegerField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "INT"
	case postgresql:
		return "integer"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*SmallIntegerField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "SMALLINT"
	case postgresql:
		return "smallint"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}

func (*TextField) FieldDataType(database string) string {
	switch database {
	case mysql:
		return "TEXT"
	case postgresql:
		return "text"
	default:
		log.Fatalf("%s not supported as per now.", database)
		return ""
	}
}
