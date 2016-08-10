package main_test

import (
	. "github.com/dcarley/swt-wifi-login"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"testing"
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
})
