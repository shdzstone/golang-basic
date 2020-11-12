package filelisting

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"golang-basic/lang/errorhandling/filelistingserver/errwrapper"
	"strings"
	"testing"
)

/*
1.http测试
* 通过使用假的Request/Response：只是假的测试，测试粒度比较高，类似于单元测试
* 通过起服务器：集成度比较高，代码覆盖度比较高
*/

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}
func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

var tests = []struct {
	h       errwrapper.AppHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
}

//此处的http请求只是假数据
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errwrapper.ErrWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)

		//recoder转换为response
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

//此处真正强大的是专门起了一个server来实现HTTP请求的发送及返回
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errwrapper.ErrWrapper(tt.h)
		//函数式编程
		//注：http.HandleFunc(f)是把f函数转换为HandlerFunc类型，然后传进NewServer函数。
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("Expect(%d,%s);"+"got(%d,%s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}
