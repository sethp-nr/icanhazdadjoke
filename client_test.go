package icanhazdadjoke

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("The dad joke client", func() {
	var server *ghttp.Server
	var client *Client

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = NewClient(serverUrl(server))
	})

	JustBeforeEach(func() {
		Expect(client).To(Not(BeNil()))
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("fetching plain text joke", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/"),
					ghttp.VerifyHeader(http.Header{
						"User-Agent": []string{UserAgent},
					}),
					ghttp.VerifyHeader(http.Header{
						"Accept": []string{"text/plain"},
					}),
					// https://icanhazdadjoke.com/j/fiyPR7wPZDd
					ghttp.RespondWith(200, "When does a joke become a dad joke? When it becomes apparent."),
				),
			)
		})

		It("retrieves a random joke", func() {
			Expect(client.GetRandomJokeText()).To(ContainSubstring("apparent"))
		})
	})
})

func serverUrl(server *ghttp.Server) FuncOpt {
	return func(c *Client) {
		c.url = server.URL()
	}
}
