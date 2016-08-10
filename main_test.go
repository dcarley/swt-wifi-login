package main_test

import (
	. "github.com/dcarley/swt-wifi-login"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"testing"

	"github.com/onsi/gomega/ghttp"
)

func TestSwtWifiLogin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SwtWifiLogin Suite")
}

var _ = Describe("SwtWifiLogin", func() {
	const (
		Username     = "myuser"
		Password     = "mypass"
		PasswordHash = "e727d1464ae12436e899a726da5b2f11d8381b26" // `echo -n "<Password>" | shasum -a1`
		BaseURL      = "http://localhost/cws"
	)

	Describe("LoginRequest", func() {
		var req *http.Request

		BeforeEach(func() {
			var err error
			req, err = LoginRequest(Username, Password, BaseURL)
			Expect(err).ToNot(HaveOccurred())
		})

		It("has correct base URL", func() {
			Expect(req.URL.String()).To(HavePrefix(BaseURL))
		})

		It("has a GET method", func() {
			Expect(req.Method).To(Equal("GET"))
		})

		It("sets rq query param", func() {
			Expect(req.URL.Query().Get("rq")).To(Equal("login"))
		})

		It("sets username query param", func() {
			Expect(req.URL.Query().Get("username")).To(Equal(Username))
		})

		It("sets password query param as hash", func() {
			Expect(req.URL.Query().Get("password")).To(Equal(PasswordHash))
		})
	})

	Describe("Login", func() {
		var (
			client       *http.Client
			server       *ghttp.Server
			request      *http.Request
			responseCode int
			responseBody string
		)

		BeforeEach(func() {
			client = &http.Client{}
			server = ghttp.NewServer()

			var err error
			request, err = LoginRequest(Username, Password, server.URL()+"/")
			Expect(err).ToNot(HaveOccurred())

			headers := http.Header{}
			headers.Set("Content-type", "application/json")

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", request.URL.Path),
					ghttp.RespondWithPtr(&responseCode, &responseBody, headers),
				),
			)
		})

		AfterEach(func() {
			server.Close()
		})

		Context("200 status code and 0 errorcode", func() {
			BeforeEach(func() {
				responseCode = http.StatusOK
				responseBody = `{"errorcode": 0}`
			})

			It("returns no errors", func() {
				Expect(Login(request, client)).To(Succeed())
			})
		})

		Context("200 status code and non-0 errorcode", func() {
			BeforeEach(func() {
				responseCode = http.StatusOK
				responseBody = `{"errorcode": 101}`
			})

			It("returns an error with status code and response body", func() {
				Expect(
					Login(request, client),
				).To(
					MatchError(LoginError{responseCode, responseBody}),
				)
			})
		})

		Context("non-200 status code and unparseable response body", func() {
			BeforeEach(func() {
				responseCode = http.StatusServiceUnavailable
				responseBody = `there was a problem`
			})

			It("returns an error with status code and response body", func() {
				Expect(
					Login(request, client),
				).To(
					MatchError(LoginError{responseCode, responseBody}),
				)
			})
		})

		Context("200 status code and unparseable response body", func() {
			BeforeEach(func() {
				responseCode = http.StatusOK
				responseBody = `there was a problem`
			})

			It("returns an error with JSON decoder details", func() {
				Expect(
					Login(request, client),
				).To(
					MatchError(ContainSubstring("invalid character")),
				)
			})
		})
	})
})
