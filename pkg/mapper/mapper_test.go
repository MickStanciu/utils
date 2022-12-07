package mapper_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/MickStanciu/utils/pkg/mapper"
	"github.com/stretchr/testify/assert"
)

func TestGenericMapper(t *testing.T) {
	type from struct {
		FIELD1 string
		Meta   struct {
			FIELD1 string
			FIELD2 int
		}
	}

	type to struct {
		FIELD1 string
		FIELD2 string
	}

	req := []from{
		{
			FIELD1: "F1",
			Meta: struct {
				FIELD1 string
				FIELD2 int
			}{FIELD1: "F1", FIELD2: 1},
		},
		{
			FIELD1: "F2",
			Meta: struct {
				FIELD1 string
				FIELD2 int
			}{FIELD1: "F2", FIELD2: 2},
		},
	}

	transformer := func(data from) *to {
		return &to{
			FIELD1: data.FIELD1,
			FIELD2: fmt.Sprintf("%s - %d", data.Meta.FIELD1, data.Meta.FIELD2),
		}
	}
	res := mapper.GenericMapper(req, transformer)
	assert.EqualValues(t, 2, len(res))
	assert.IsType(t, []*to{}, res)

	shouldBe := []*to{
		{
			FIELD1: "F1",
			FIELD2: "F1 - 1",
		},
		{
			FIELD1: "F2",
			FIELD2: "F2 - 2",
		},
	}

	eq := reflect.DeepEqual(shouldBe, res)
	assert.True(t, eq)

}
