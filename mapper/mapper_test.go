package mapper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"

	. "github.com/tadayosi/samples-go/util"
)

type Object struct {
	PropOne   *bool  `property:"prop-one" json:"propOne,omitempty"`
	PropTwo   int    `property:"prop-two" json:"propTwo,omitempty"`
	PropThree string `property:"prop-three" json:"propThree,omitempty"`
}

type Child struct {
	Object   `property:",squash"`
	PropFour string `property:"prop-four" json:"propFour,omitempty"`
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
		PropOne:   BoolP(true),
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

func TestMap2Child(t *testing.T) {
	data := map[string]interface{}{
		"prop-one":   "false",
		"prop-two":   "1",
		"prop-three": "test",
		"prop-four":  "child-test",
	}
	child := &Child{
		Object: Object{
			PropOne:   BoolP(true),
			PropTwo:   0,
			PropThree: "",
		},
		PropFour: "",
	}

	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			Metadata:         &mapstructure.Metadata{},
			WeaklyTypedInput: true,
			TagName:          "property",
			Result:           child,
			ErrorUnused:      true,
		},
	)
	assert.Nil(t, err)

	err = decoder.Decode(data)
	assert.Nil(t, err)
	assert.False(t, *child.PropOne)
	assert.Equal(t, 1, child.PropTwo)
	assert.Equal(t, "test", child.PropThree)
	assert.Equal(t, "child-test", child.PropFour)
}

func TestObject2JSON(t *testing.T) {
	obj := &Object{
		PropOne:   BoolP(false),
		PropTwo:   1,
		PropThree: "test",
	}

	data, err := json.Marshal(obj)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"propOne":false,"propTwo":1,"propThree":"test"}`, string(data))
}

func TestChild2JSON(t *testing.T) {
	child := &Child{
		Object: Object{
			PropOne:   BoolP(false),
			PropTwo:   1,
			PropThree: "test",
		},
		PropFour: "child-test",
	}

	data, err := json.Marshal(child)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"propOne":false,"propTwo":1,"propThree":"test","propFour":"child-test"}`, string(data))
}

func TestObject2JSON_PropOneEmpty(t *testing.T) {
	obj := &Object{
		//PropOne:   BoolP(false),
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
		PropOne:   BoolP(true),
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
	assert.Nil(t, err)

	result, err := json.Marshal(obj)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"propOne":false,"propTwo":1,"propThree":"test"}`, string(result))
}

func TestMap2Child2JSON(t *testing.T) {
	data := map[string]interface{}{
		"prop-one":   "false",
		"prop-two":   "1",
		"prop-three": "test",
		"prop-four":  "child-test",
	}
	c := &Child{
		Object: Object{
			PropOne:   BoolP(true),
			PropTwo:   0,
			PropThree: "",
		},
		PropFour: "",
	}
	child := reflect.New(reflect.TypeOf(c)).Interface()

	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			Metadata:         &mapstructure.Metadata{},
			WeaklyTypedInput: true,
			TagName:          "property",
			Result:           child,
		},
	)
	assert.Nil(t, err)

	err = decoder.Decode(data)
	assert.Nil(t, err)

	result, err := json.Marshal(child)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"propOne":false,"propTwo":1,"propThree":"test","propFour":"child-test"}`, string(result))
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
