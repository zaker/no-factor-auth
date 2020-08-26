package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"testing"

	"github.com/equinor/no-factor-auth/tests"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnParseError(t *testing.T) {
	jsonString := []byte("malformed json")

	_, err := ParseExtraClaims(jsonString)
	assert.Error(t, err, "Malformed json should return error")

}

func TestParseJson(t *testing.T) {
	oidValue := "b213-61024b63a7ea"
	arrValue := "some-value"

	jsonString := []byte(`{"oid":"` + oidValue + `", "arr_elem_key":["` + arrValue + `","foo"]}`)

	result, err := ParseExtraClaims(jsonString)

	if assert.NoError(t, err) {
		assert.Equal(t, oidValue, result["oid"], "Simple json elements should be parsed")

		arrElements := extractArrayElement(result)
		assert.True(t, contains(arrElements, arrValue), "Json array should be parsed")
	}

}

func contains(s []string, searchterm string) bool {
	sort.Strings(s)
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func extractArrayElement(result map[string]interface{}) []string {
	var items []string

	object := reflect.ValueOf(result["arr_elem_key"])
	for i := 0; i < object.Len(); i++ {
		s := fmt.Sprint(object.Index(i).Interface())
		items = append(items, s)
	}

	return items
}

func TestToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := tests.NewMockContext(ctrl)
	ctx.EXPECT().
		QueryParam("redirect_uri").
		Return("http://example.com")
	ctx.EXPECT().
		QueryParam("client_id").
		Return("1234")
	ctx.EXPECT().
		QueryParam("grant_type").
		Return("client_credential")

	ctx.EXPECT().
		QueryParam("code").
		Return("code")

	ctx.EXPECT().
		QueryParam("client_secret").
		Return("secret")

	ctx.EXPECT().
		QueryParam("extra_claims").
		Return("")

	ctx.EXPECT().
		Request().
		Return(&http.Request{Host: "http://localhost"})

	ctx.EXPECT().
		JSON(http.StatusOK, gomock.Any()).
		Return(nil)

	t.Run("On correct query params the token should be received in a json struct", func(t *testing.T) {
		if err := Token(ctx); err != nil {
			t.Errorf("Token() error = %v", err)
		}
	})

}
