package controllers_test

import (
	"testing"

	"github.com/katsu105/go-intermediate/controllers"
	"github.com/katsu105/go-intermediate/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
