package matchers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/onsi/gomega"
	"reflect"
)

type jsonMatcher struct {
	expectedJson string
}

func MatchJson(expectedJson string) gomega.OmegaMatcher {
	return jsonMatcher{expectedJson: expectedJson}
}

func (matcher jsonMatcher) Match(actual interface{}) (success bool, message string, err error) {
	var actualJson string

	switch actual := actual.(type) {
	case string:
		actualJson = actual
	default:
		err = errors.New(fmt.Sprintf("expected string, not %T", actual))
		return
	}

	result1, err := ToJson(actualJson)
	if err != nil {
		return
	}
	result2, err := ToJson(matcher.expectedJson)
	if err != nil {
		return
	}

	success = reflect.DeepEqual(result1, result2)
	if success {
		message = fmt.Sprintf("expected \n\t%#v\nnot to equal:\n\t%#v", result1, result2)
	} else {
		message = fmt.Sprintf("expected \n\t%#v\nto equal:\n\t%#v", result1, result2)
	}
	return
}

func ToJson(data string) (map[string]interface{}, error) {
	result := new(map[string]interface{})
	err := json.Unmarshal([]byte(data), result)
	return *result, err
}
