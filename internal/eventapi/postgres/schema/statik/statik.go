// Code generated by statik. DO NOT EDIT.

package statik

import (
	"github.com/rakyll/statik/fs"
)

const EventapiSql = "eventapi/sql" // static asset namespace

func init() {
	data := "PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x16\x00	\x00001_initial_schema.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE jobset\n(\n    id        bigserial PRIMARY KEY,\n    queue     varchar(512) NOT NULL,\n    jobset    varchar(512) NOT NULL,\n);\n\nCREATE TABLE event\n(\n    jobset    bigint,\n    sequence  bigint,\n    msg bytea,\n    PRIMARY KEY (jobset, sequence)\n);PK\x07\x08\x1f\xc8r\x17\xfe\x00\x00\x00\xfe\x00\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x1f\xc8r\x17\xfe\x00\x00\x00\xfe\x00\x00\x00\x16\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00001_initial_schema.sqlUT\x05\x00\x01\x80Cm8PK\x05\x06\x00\x00\x00\x00\x01\x00\x01\x00M\x00\x00\x00K\x01\x00\x00\x00\x00"
	fs.RegisterWithNamespace("eventapi/sql", data)
}