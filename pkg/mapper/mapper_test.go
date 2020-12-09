package mapper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

type Object struct {
	PropOne   *bool  `property:"prop-one" json:"propOne,omitempty"`
	PropTwo   int    `property:"prop-two" json:"propTwo,omitempty"`
	PropThree string `property:"prop-three" json:"propThree,omitempty"`
}

func (o *Object) String() string {
	return fmt.Sprintf("Object[PropOne=%v, PropTwo=%v, PropThree=%v]",
		o.PropOne, o.PropTwo, o.PropThree)
}

func TestMap2Object(t *testing.T) {
	data := map[string]interface{}{
		"prop-one":   "false",
		"prop-two":   "1",
		"prop-three": "test",
	}
	obj := &Object{
		PropOne:   &[]bool{true}[0],
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
	assert.False(t, *obj.PropOne)
	assert.Equal(t, 1, obj.PropTwo)
	assert.Equal(t, "test", obj.PropThree)
}

func TestObject2JSON(t *testing.T) {
	obj := &Object{
		PropOne:   &[]bool{false}[0],
		PropTwo:   1,
		PropThree: "test",
	}

	data, err := json.Marshal(obj)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"propOne":false,"propTwo":1,"propThree":"test"}`, string(data))
}

func TestObject2JSON_PropOneEmpty(t *testing.T) {
	obj := &Object{
		//PropOne:   &[]bool{false}[0],
		PropTwo:   1,
		PropThree: "test",
	}
	assert.Nil(t, obj.PropOne)

	data, err := json.Marshal(obj)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"propTwo":1,"propThree":"test"}`, string(data))
}

func TestMap2Object2JSON(t *testing.T) {
	data := map[string]interface{}{
		"prop-one":   "false",
		"prop-two":   "1",
		"prop-three": "test",
	}
	o := &Object{
		PropOne:   &[]bool{true}[0],
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
	assert.False(t, *obj.PropOne)
	assert.Equal(t, 1, obj.PropTwo)
	assert.Equal(t, "test", obj.PropThree)
}
