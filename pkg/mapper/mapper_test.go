package mapstruct

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

type Object struct {
	Prop1 bool   `property:"prop1" json:"prop1"`
	Prop2 int    `property:"prop2" json:"prop2"`
	Prop3 string `property:"prop3" json:"prop3"`
}

type Object2 struct {
	Raw []byte `json:",inline"`
}

func TestMap2Object(t *testing.T) {
	data := map[string]interface{}{
		"prop1": "false",
		"prop2": "1",
		"prop3": "test",
	}
	obj := &Object{}

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
	assert.False(t, obj.Prop1)
	assert.Equal(t, 1, obj.Prop2)
	assert.Equal(t, "test", obj.Prop3)
}

func TestObject2JSON(t *testing.T) {
	obj := &Object{
		Prop1: false,
		Prop2: 1,
		Prop3: "test",
	}

	data, err := json.Marshal(obj)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"prop1":false,"prop2":1,"prop3":"test"}`, string(data))
}

func TestJSON2Object(t *testing.T) {
	data := []byte(`{"prop1":false,"prop2":1,"prop3":"test"}`)
	obj := &Object{}

	err := json.Unmarshal(data, obj)

	assert.Nil(t, err)
	assert.False(t, obj.Prop1)
	assert.Equal(t, 1, obj.Prop2)
	assert.Equal(t, "test", obj.Prop3)
}
