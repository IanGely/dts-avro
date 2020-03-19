// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     record.avsc
 */
package avro

import (
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/vm"
)

  
type SourceType int32

const (

	SourceTypeMySQL SourceType = 0

	SourceTypeOracle SourceType = 1

	SourceTypeSQLServer SourceType = 2

	SourceTypePostgreSQL SourceType = 3

	SourceTypeMongoDB SourceType = 4

	SourceTypeRedis SourceType = 5

	SourceTypeDB2 SourceType = 6

	SourceTypePPAS SourceType = 7

	SourceTypeDRDS SourceType = 8

	SourceTypeHBASE SourceType = 9

	SourceTypeHDFS SourceType = 10

	SourceTypeFILE SourceType = 11

	SourceTypeOTHER SourceType = 12

)

func (e SourceType) String() string {
	switch e {

	case SourceTypeMySQL:
		return "MySQL"

	case SourceTypeOracle:
		return "Oracle"

	case SourceTypeSQLServer:
		return "SQLServer"

	case SourceTypePostgreSQL:
		return "PostgreSQL"

	case SourceTypeMongoDB:
		return "MongoDB"

	case SourceTypeRedis:
		return "Redis"

	case SourceTypeDB2:
		return "DB2"

	case SourceTypePPAS:
		return "PPAS"

	case SourceTypeDRDS:
		return "DRDS"

	case SourceTypeHBASE:
		return "HBASE"

	case SourceTypeHDFS:
		return "HDFS"

	case SourceTypeFILE:
		return "FILE"

	case SourceTypeOTHER:
		return "OTHER"

	}
	return "unknown"
}

func writeSourceType(r SourceType, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewSourceTypeValue(raw string) (r SourceType, err error) {
	switch raw {

	case "MySQL":
		return SourceTypeMySQL, nil

	case "Oracle":
		return SourceTypeOracle, nil

	case "SQLServer":
		return SourceTypeSQLServer, nil

	case "PostgreSQL":
		return SourceTypePostgreSQL, nil

	case "MongoDB":
		return SourceTypeMongoDB, nil

	case "Redis":
		return SourceTypeRedis, nil

	case "DB2":
		return SourceTypeDB2, nil

	case "PPAS":
		return SourceTypePPAS, nil

	case "DRDS":
		return SourceTypeDRDS, nil

	case "HBASE":
		return SourceTypeHBASE, nil

	case "HDFS":
		return SourceTypeHDFS, nil

	case "FILE":
		return SourceTypeFILE, nil

	case "OTHER":
		return SourceTypeOTHER, nil

	}

	return -1, fmt.Errorf("invalid value for SourceType: '%s'", raw)
}