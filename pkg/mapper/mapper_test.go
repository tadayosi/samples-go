package mapper

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

type Object struct {
	PropOne   bool   `property:"prop-one" json:"propOne"`
	PropTwo   int    `property:"prop-two" json:"propTwo"`
	PropThree string `property:"prop-three" json:"propThree"`
}

type Object2 struct {
	Raw []byte `json:",inline"`
}

func TestMap2Object(t *testing.T) {
	data := map[string]interface{}{
		"prop-one":   "false",
		"prop-two":   "1",
		"prop-three": "test",
	}
	obj := &Object{
		PropOne:   true,
		PropTwo:   0,
		PropThree: "",
	}

	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			Metadata:         &mapstructure.Metadata{},
			WeaklyTypedInput: true,
			TagName:          "property",
			Result:           obj,
		},
	)
	assert.Nil(t, err)

	err = decoder.Decode(data)
	assert.Nil(t, err)
	assert.False(t, obj.PropOne)
	assert.Equal(t, 1, obj.PropTwo)
	assert.Equal(t, "test", obj.PropThree)
}

func TestObject2JSON(t *testing.T) {
	obj := &Object{
		PropOne:   false,
		PropTwo:   1,
		PropThree: "test",
	}

	data, err := json.Marshal(obj)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"propOne":false,"propTwo":1,"propThree":"test"}`, string(data))
}

func TestMap2Object2JSON(t *testing.T) {
	data := map[string]interface{}{
		"prop-one":   "false",
		"prop-two":   "1",
		"prop-three": "test",
	}
	o := &Object{
		PropOne:   true,
		PropTwo:   0,
		PropThree: "",
	}
	obj := reflect.New(reflect.TypeOf(o)).Interface()

	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			Metadata:         &mapstructure.Metadata{},
			WeaklyTypedInput: true,
			TagName:          "property",
			Result:           obj,
		},
	)
	assert.Nil(t, err)

	err = decoder.Decode(data)

	result, err := json.Marshal(obj)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"propOne":false,"propTwo":1,"propThree":"test"}`, string(result))
}

func TestJSON2Object(t *testing.T) {
	data := []byte(`{"propOne":false,"propTwo":1,"propThree":"test"}`)
	obj := &Object{}

	err := json.Unmarshal(data, obj)

	assert.Nil(t, err)
	assert.False(t, obj.PropOne)
	assert.Equal(t, 1, obj.PropTwo)
	assert.Equal(t, "test", obj.PropThree)
}
