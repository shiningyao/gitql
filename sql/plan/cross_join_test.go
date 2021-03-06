package plan

import (
	"io"
	"testing"

	"github.com/gitql/gitql/mem"
	"github.com/gitql/gitql/sql"

	"github.com/stretchr/testify/assert"
)

var lSchema = sql.Schema{
	sql.Field{"lcol1", sql.String},
	sql.Field{"lcol2", sql.String},
	sql.Field{"lcol3", sql.Integer},
	sql.Field{"lcol4", sql.BigInteger},
}

var rSchema = sql.Schema{
	sql.Field{"rcol1", sql.String},
	sql.Field{"rcol2", sql.String},
	sql.Field{"rcol3", sql.Integer},
	sql.Field{"rcol4", sql.BigInteger},
}

func TestCrossJoin(t *testing.T) {
	assert := assert.New(t)

	resultSchema := sql.Schema{
		sql.Field{"lcol1", sql.String},
		sql.Field{"lcol2", sql.String},
		sql.Field{"lcol3", sql.Integer},
		sql.Field{"lcol4", sql.BigInteger},
		sql.Field{"rcol1", sql.String},
		sql.Field{"rcol2", sql.String},
		sql.Field{"rcol3", sql.Integer},
		sql.Field{"rcol4", sql.BigInteger},
	}

	ltable := mem.NewTable("left", lSchema)
	rtable := mem.NewTable("right", rSchema)
	insertData(assert, ltable)
	insertData(assert, rtable)

	j := NewCrossJoin(ltable, rtable)

	assert.Equal(resultSchema, j.Schema())

	iter, err := j.RowIter()
	assert.Nil(err)
	assert.NotNil(iter)

	row, err := iter.Next()
	assert.Nil(err)
	assert.NotNil(row)

	assert.Equal(8, len(row.Fields()))

	assert.Equal("col1_1", row.Fields()[0])
	assert.Equal("col2_1", row.Fields()[1])
	assert.Equal(int32(1111), row.Fields()[2])
	assert.Equal(int64(2222), row.Fields()[3])
	assert.Equal("col1_1", row.Fields()[4])
	assert.Equal("col2_1", row.Fields()[5])
	assert.Equal(int32(1111), row.Fields()[6])
	assert.Equal(int64(2222), row.Fields()[7])

	row, err = iter.Next()
	assert.Nil(err)
	assert.NotNil(row)

	assert.Equal("col1_1", row.Fields()[0])
	assert.Equal("col2_1", row.Fields()[1])
	assert.Equal(int32(1111), row.Fields()[2])
	assert.Equal(int64(2222), row.Fields()[3])
	assert.Equal("col1_2", row.Fields()[4])
	assert.Equal("col2_2", row.Fields()[5])
	assert.Equal(int32(3333), row.Fields()[6])
	assert.Equal(int64(4444), row.Fields()[7])

	for i := 0; i < 2; i++ {
		row, err = iter.Next()
		assert.Nil(err)
		assert.NotNil(row)
	}

	// total: 4 rows
	row, err = iter.Next()
	assert.NotNil(err)
	assert.Equal(err, io.EOF)
	assert.Nil(row)
}

func TestCrossJoin_Empty(t *testing.T) {
	assert := assert.New(t)

	ltable := mem.NewTable("left", lSchema)
	rtable := mem.NewTable("right", rSchema)
	insertData(assert, ltable)

	j := NewCrossJoin(ltable, rtable)

	iter, err := j.RowIter()
	assert.Nil(err)
	assert.NotNil(iter)

	row, err := iter.Next()
	assert.Equal(io.EOF, err)
	assert.Nil(row)

	ltable = mem.NewTable("left", lSchema)
	rtable = mem.NewTable("right", rSchema)
	insertData(assert, rtable)

	j = NewCrossJoin(ltable, rtable)

	iter, err = j.RowIter()
	assert.Nil(err)
	assert.NotNil(iter)

	row, err = iter.Next()
	assert.Equal(io.EOF, err)
	assert.Nil(row)
}

func insertData(assert *assert.Assertions, table *mem.Table) {
	err := table.Insert("col1_1", "col2_1", int32(1111), int64(2222))
	assert.Nil(err)
	err = table.Insert("col1_2", "col2_2", int32(3333), int64(4444))
	assert.Nil(err)
}
