package commands_test

import (
	"fmt"

	cli_fakes "github.com/SAP/cf-mta-plugin/cli/fakes"
	"github.com/SAP/cf-mta-plugin/clients/models"
	mtafake "github.com/SAP/cf-mta-plugin/clients/mtaclient/fakes"
	"github.com/SAP/cf-mta-plugin/commands"
	"github.com/SAP/cf-mta-plugin/testutil"
	"github.com/SAP/cf-mta-plugin/ui"
	util_fakes "github.com/SAP/cf-mta-plugin/util/fakes"
	plugin_fakes "github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MtasCommand", func() {
	Describe("Execute", func() {
		const org = "test-org"
		const space = "test-space"
		const user = "test-user"

		var name string
		var cliConnection *plugin_fakes.FakeCliConnection
		var clientFactory *commands.TestClientFactory
		var command *commands.MtasCommand
		var oc = testutil.NewUIOutputCapturer()
		var ex = testutil.NewUIExpector()
		var testTokenFactory *commands.TestTokenFactory

		var getOutputLines = func(mtas [][]string) []string {
			lines := []string{}
			lines = append(lines,
				fmt.Sprintf("Getting multi-target apps in org %s / space %s as %s...\n", org, space, user))
			lines = append(lines, "OK\n")
			if mtas != nil {
				lines = append(lines, testutil.GetTableOutputLines([]string{"mta id", "version"}, mtas)...)
			} else {
				lines = append(lines, "No multi-target apps found\n")
			}
			return lines
		}

		BeforeEach(func() {
			ui.DisableTerminalOutput(true)
			name = command.GetPluginCommand().Name
			cliConnection = cli_fakes.NewFakeCliConnectionBuilder().
				CurrentOrg("test-org-guid", org, nil).
				CurrentSpace("test-space-guid", space, nil).
				Username(user, nil).
				AccessToken("bearer test-token", nil).
				APIEndpoint("https://api.test.ondemand.com", nil).Build()
			mtaClient := mtafake.NewFakeMtaClientBuilder().
				GetMtas(nil, nil).Build()
			clientFactory = commands.NewTestClientFactory(mtaClient, nil)
			command = &commands.MtasCommand{}
			testTokenFactory = commands.NewTestTokenFactory(cliConnection)
			deployServiceURLCalculator := util_fakes.NewDeployServiceURLFakeCalculator("deploy-service.test.ondemand.com")
			command.InitializeAll(name, cliConnection, testutil.NewCustomTransport(200, nil), nil, clientFactory, testTokenFactory, deployServiceURLCalculator)
		})

		// unknown flag - error
		Context("with an unknown flag", func() {
			It("should print incorrect usage, call cf help, and exit with a non-zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-a"}).ToInt()
				})
				ex.ExpectFailure(status, output, "Incorrect usage. Unknown or wrong flag.")
				Expect(cliConnection.CliCommandArgsForCall(0)).To(Equal([]string{"help", name}))
			})
		})

		// wrong arguments - error
		Context("with wrong arguments", func() {
			It("should print incorrect usage, call cf help, and exit with a non-zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"x", "y", "z"}).ToInt()
				})
				ex.ExpectFailure(status, output, "Incorrect usage. Wrong arguments.")
				Expect(cliConnection.CliCommandArgsForCall(0)).To(Equal([]string{"help", name}))
			})
		})

		// can't connect to backend - error
		Context("when can't connect to backend", func() {
			const host = "x"
			It("should print an error and exit with a non-zero status", func() {
				clientFactory.MtaClient = mtafake.NewFakeMtaClientBuilder().
					GetMtas(nil, fmt.Errorf("Get https://%s/rest/test/test/components: dial tcp: lookup %s: no such host", host, host)).Build()
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-u", host}).ToInt()
				})
				ex.ExpectFailureOnLine(status, output, "Could not get deployed components:", 2)
			})
		})

		// backend returns an an error response - error
		Context("with an error response returned by the backend", func() {
			It("should print an error and exit with a non-zero status", func() {
				clientFactory.MtaClient = mtafake.NewFakeMtaClientBuilder().
					GetMtas(nil, fmt.Errorf("unknown error (status 404)")).Build()
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{}).ToInt()
				})
				ex.ExpectFailureOnLine(status, output, "Could not get deployed components:", 1)
			})
		})

		// backend returns an empty response - success
		Context("with an empty response returned by the backend", func() {
			It("should print a message and exit with zero status", func() {
				clientFactory.MtaClient = mtafake.NewFakeMtaClientBuilder().
					GetMtas([]*models.Mta{}, nil).Build()
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{}).ToInt()
				})
				ex.ExpectSuccessWithOutput(status, output, getOutputLines(nil))
			})
		})

		// backend returns a non-empty response - success
		Context("with a non-empty response returned by the backend containing an unknown MTA version", func() {
			It("should print a table with all deployed MTAs and exit with zero status", func() {
				clientFactory.MtaClient = mtafake.NewFakeMtaClientBuilder().
					GetMtas([]*models.Mta{testutil.GetMta("org.cloudfoundry.samples.music", "0.0.0-unknown",
						[]*models.Module{testutil.GetMtaModule("spring-music", []string{"postgresql"}, []string{})},
						[]string{"postgresql"})}, nil).Build()
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{}).ToInt()
				})
				ex.ExpectSuccessWithOutput(status, output,
					getOutputLines([][]string{[]string{"org.cloudfoundry.samples.music", "?"}}))
			})
		})

		// backend returns a non-empty response - success
		Context("with a non-empty response returned by the backend", func() {
			It("should print a table with all deployed MTAs and exit with zero status", func() {
				clientFactory.MtaClient = mtafake.NewFakeMtaClientBuilder().
					GetMtas([]*models.Mta{testutil.GetMta("org.cloudfoundry.samples.music", "1.0",
						[]*models.Module{testutil.GetMtaModule("spring-music", []string{"postgresql"}, []string{})},
						[]string{"postgresql"})}, nil).Build()
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{}).ToInt()
				})
				ex.ExpectSuccessWithOutput(status, output,
					getOutputLines([][]string{[]string{"org.cloudfoundry.samples.music", "1.0"}}))
			})
		})
	})
})
