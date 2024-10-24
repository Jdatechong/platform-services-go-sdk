/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package partnercentersellv1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PartnerCenterSellV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(partnerCenterSellService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(partnerCenterSellService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
				URL: "https://partnercentersellv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(partnerCenterSellService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_CENTER_SELL_URL": "https://partnercentersellv1/api",
				"PARTNER_CENTER_SELL_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
				})
				Expect(partnerCenterSellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := partnerCenterSellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerCenterSellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerCenterSellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerCenterSellService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
					URL: "https://testService/api",
				})
				Expect(partnerCenterSellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerCenterSellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerCenterSellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerCenterSellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerCenterSellService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
				})
				err := partnerCenterSellService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerCenterSellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerCenterSellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerCenterSellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerCenterSellService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_CENTER_SELL_URL": "https://partnercentersellv1/api",
				"PARTNER_CENTER_SELL_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(partnerCenterSellService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_CENTER_SELL_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(partnerCenterSellService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = partnercentersellv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateRegistration(createRegistrationOptions *CreateRegistrationOptions) - Operation response error`, func() {
		createRegistrationPath := "/registration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRegistrationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRegistration with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("Company Representative")
				primaryContactModel.Email = core.StringPtr("companyrep@email.com")

				// Construct an instance of the CreateRegistrationOptions model
				createRegistrationOptionsModel := new(partnercentersellv1.CreateRegistrationOptions)
				createRegistrationOptionsModel.AccountID = core.StringPtr("4a5c3c51b97a446fbb1d0e1ef089823b")
				createRegistrationOptionsModel.CompanyName = core.StringPtr("Beautiful Company")
				createRegistrationOptionsModel.PrimaryContact = primaryContactModel
				createRegistrationOptionsModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				createRegistrationOptionsModel.ProviderAccessGroup = core.StringPtr("testString")
				createRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreateRegistration(createRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreateRegistration(createRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRegistration(createRegistrationOptions *CreateRegistrationOptions)`, func() {
		createRegistrationPath := "/registration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRegistrationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "company_name": "CompanyName", "primary_contact": {"name": "Name", "email": "Email"}, "default_private_catalog_id": "DefaultPrivateCatalogID", "provider_access_group": "ProviderAccessGroup", "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
				}))
			})
			It(`Invoke CreateRegistration successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("Company Representative")
				primaryContactModel.Email = core.StringPtr("companyrep@email.com")

				// Construct an instance of the CreateRegistrationOptions model
				createRegistrationOptionsModel := new(partnercentersellv1.CreateRegistrationOptions)
				createRegistrationOptionsModel.AccountID = core.StringPtr("4a5c3c51b97a446fbb1d0e1ef089823b")
				createRegistrationOptionsModel.CompanyName = core.StringPtr("Beautiful Company")
				createRegistrationOptionsModel.PrimaryContact = primaryContactModel
				createRegistrationOptionsModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				createRegistrationOptionsModel.ProviderAccessGroup = core.StringPtr("testString")
				createRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreateRegistrationWithContext(ctx, createRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreateRegistration(createRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreateRegistrationWithContext(ctx, createRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRegistrationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "company_name": "CompanyName", "primary_contact": {"name": "Name", "email": "Email"}, "default_private_catalog_id": "DefaultPrivateCatalogID", "provider_access_group": "ProviderAccessGroup", "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
				}))
			})
			It(`Invoke CreateRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreateRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("Company Representative")
				primaryContactModel.Email = core.StringPtr("companyrep@email.com")

				// Construct an instance of the CreateRegistrationOptions model
				createRegistrationOptionsModel := new(partnercentersellv1.CreateRegistrationOptions)
				createRegistrationOptionsModel.AccountID = core.StringPtr("4a5c3c51b97a446fbb1d0e1ef089823b")
				createRegistrationOptionsModel.CompanyName = core.StringPtr("Beautiful Company")
				createRegistrationOptionsModel.PrimaryContact = primaryContactModel
				createRegistrationOptionsModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				createRegistrationOptionsModel.ProviderAccessGroup = core.StringPtr("testString")
				createRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreateRegistration(createRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRegistration with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("Company Representative")
				primaryContactModel.Email = core.StringPtr("companyrep@email.com")

				// Construct an instance of the CreateRegistrationOptions model
				createRegistrationOptionsModel := new(partnercentersellv1.CreateRegistrationOptions)
				createRegistrationOptionsModel.AccountID = core.StringPtr("4a5c3c51b97a446fbb1d0e1ef089823b")
				createRegistrationOptionsModel.CompanyName = core.StringPtr("Beautiful Company")
				createRegistrationOptionsModel.PrimaryContact = primaryContactModel
				createRegistrationOptionsModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				createRegistrationOptionsModel.ProviderAccessGroup = core.StringPtr("testString")
				createRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreateRegistration(createRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRegistrationOptions model with no property values
				createRegistrationOptionsModelNew := new(partnercentersellv1.CreateRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreateRegistration(createRegistrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("Company Representative")
				primaryContactModel.Email = core.StringPtr("companyrep@email.com")

				// Construct an instance of the CreateRegistrationOptions model
				createRegistrationOptionsModel := new(partnercentersellv1.CreateRegistrationOptions)
				createRegistrationOptionsModel.AccountID = core.StringPtr("4a5c3c51b97a446fbb1d0e1ef089823b")
				createRegistrationOptionsModel.CompanyName = core.StringPtr("Beautiful Company")
				createRegistrationOptionsModel.PrimaryContact = primaryContactModel
				createRegistrationOptionsModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				createRegistrationOptionsModel.ProviderAccessGroup = core.StringPtr("testString")
				createRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreateRegistration(createRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRegistration(getRegistrationOptions *GetRegistrationOptions) - Operation response error`, func() {
		getRegistrationPath := "/registration/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRegistrationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRegistration with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetRegistrationOptions model
				getRegistrationOptionsModel := new(partnercentersellv1.GetRegistrationOptions)
				getRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				getRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetRegistration(getRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetRegistration(getRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRegistration(getRegistrationOptions *GetRegistrationOptions)`, func() {
		getRegistrationPath := "/registration/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRegistrationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "company_name": "CompanyName", "primary_contact": {"name": "Name", "email": "Email"}, "default_private_catalog_id": "DefaultPrivateCatalogID", "provider_access_group": "ProviderAccessGroup", "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
				}))
			})
			It(`Invoke GetRegistration successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetRegistrationOptions model
				getRegistrationOptionsModel := new(partnercentersellv1.GetRegistrationOptions)
				getRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				getRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetRegistrationWithContext(ctx, getRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetRegistration(getRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetRegistrationWithContext(ctx, getRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRegistrationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "company_name": "CompanyName", "primary_contact": {"name": "Name", "email": "Email"}, "default_private_catalog_id": "DefaultPrivateCatalogID", "provider_access_group": "ProviderAccessGroup", "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
				}))
			})
			It(`Invoke GetRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRegistrationOptions model
				getRegistrationOptionsModel := new(partnercentersellv1.GetRegistrationOptions)
				getRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				getRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetRegistration(getRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRegistration with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetRegistrationOptions model
				getRegistrationOptionsModel := new(partnercentersellv1.GetRegistrationOptions)
				getRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				getRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetRegistration(getRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRegistrationOptions model with no property values
				getRegistrationOptionsModelNew := new(partnercentersellv1.GetRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetRegistration(getRegistrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetRegistrationOptions model
				getRegistrationOptionsModel := new(partnercentersellv1.GetRegistrationOptions)
				getRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				getRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetRegistration(getRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRegistration(updateRegistrationOptions *UpdateRegistrationOptions) - Operation response error`, func() {
		updateRegistrationPath := "/registration/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRegistrationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRegistration with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the RegistrationPatch model
				registrationPatchModel := new(partnercentersellv1.RegistrationPatch)
				registrationPatchModel.CompanyName = core.StringPtr("testString")
				registrationPatchModel.PrimaryContact = primaryContactModel
				registrationPatchModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				registrationPatchModel.ProviderAccessGroup = core.StringPtr("testString")
				registrationPatchModelAsPatch, asPatchErr := registrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateRegistrationOptions model
				updateRegistrationOptionsModel := new(partnercentersellv1.UpdateRegistrationOptions)
				updateRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				updateRegistrationOptionsModel.RegistrationPatch = registrationPatchModelAsPatch
				updateRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateRegistration(updateRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateRegistration(updateRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRegistration(updateRegistrationOptions *UpdateRegistrationOptions)`, func() {
		updateRegistrationPath := "/registration/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRegistrationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "company_name": "CompanyName", "primary_contact": {"name": "Name", "email": "Email"}, "default_private_catalog_id": "DefaultPrivateCatalogID", "provider_access_group": "ProviderAccessGroup", "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
				}))
			})
			It(`Invoke UpdateRegistration successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the RegistrationPatch model
				registrationPatchModel := new(partnercentersellv1.RegistrationPatch)
				registrationPatchModel.CompanyName = core.StringPtr("testString")
				registrationPatchModel.PrimaryContact = primaryContactModel
				registrationPatchModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				registrationPatchModel.ProviderAccessGroup = core.StringPtr("testString")
				registrationPatchModelAsPatch, asPatchErr := registrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateRegistrationOptions model
				updateRegistrationOptionsModel := new(partnercentersellv1.UpdateRegistrationOptions)
				updateRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				updateRegistrationOptionsModel.RegistrationPatch = registrationPatchModelAsPatch
				updateRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateRegistrationWithContext(ctx, updateRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateRegistration(updateRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateRegistrationWithContext(ctx, updateRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRegistrationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "company_name": "CompanyName", "primary_contact": {"name": "Name", "email": "Email"}, "default_private_catalog_id": "DefaultPrivateCatalogID", "provider_access_group": "ProviderAccessGroup", "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
				}))
			})
			It(`Invoke UpdateRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the RegistrationPatch model
				registrationPatchModel := new(partnercentersellv1.RegistrationPatch)
				registrationPatchModel.CompanyName = core.StringPtr("testString")
				registrationPatchModel.PrimaryContact = primaryContactModel
				registrationPatchModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				registrationPatchModel.ProviderAccessGroup = core.StringPtr("testString")
				registrationPatchModelAsPatch, asPatchErr := registrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateRegistrationOptions model
				updateRegistrationOptionsModel := new(partnercentersellv1.UpdateRegistrationOptions)
				updateRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				updateRegistrationOptionsModel.RegistrationPatch = registrationPatchModelAsPatch
				updateRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateRegistration(updateRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateRegistration with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the RegistrationPatch model
				registrationPatchModel := new(partnercentersellv1.RegistrationPatch)
				registrationPatchModel.CompanyName = core.StringPtr("testString")
				registrationPatchModel.PrimaryContact = primaryContactModel
				registrationPatchModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				registrationPatchModel.ProviderAccessGroup = core.StringPtr("testString")
				registrationPatchModelAsPatch, asPatchErr := registrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateRegistrationOptions model
				updateRegistrationOptionsModel := new(partnercentersellv1.UpdateRegistrationOptions)
				updateRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				updateRegistrationOptionsModel.RegistrationPatch = registrationPatchModelAsPatch
				updateRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateRegistration(updateRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRegistrationOptions model with no property values
				updateRegistrationOptionsModelNew := new(partnercentersellv1.UpdateRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateRegistration(updateRegistrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the RegistrationPatch model
				registrationPatchModel := new(partnercentersellv1.RegistrationPatch)
				registrationPatchModel.CompanyName = core.StringPtr("testString")
				registrationPatchModel.PrimaryContact = primaryContactModel
				registrationPatchModel.DefaultPrivateCatalogID = core.StringPtr("testString")
				registrationPatchModel.ProviderAccessGroup = core.StringPtr("testString")
				registrationPatchModelAsPatch, asPatchErr := registrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateRegistrationOptions model
				updateRegistrationOptionsModel := new(partnercentersellv1.UpdateRegistrationOptions)
				updateRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				updateRegistrationOptionsModel.RegistrationPatch = registrationPatchModelAsPatch
				updateRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateRegistration(updateRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteRegistration(deleteRegistrationOptions *DeleteRegistrationOptions)`, func() {
		deleteRegistrationPath := "/registration/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRegistrationPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := partnerCenterSellService.DeleteRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRegistrationOptions model
				deleteRegistrationOptionsModel := new(partnercentersellv1.DeleteRegistrationOptions)
				deleteRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				deleteRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = partnerCenterSellService.DeleteRegistration(deleteRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRegistration with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeleteRegistrationOptions model
				deleteRegistrationOptionsModel := new(partnercentersellv1.DeleteRegistrationOptions)
				deleteRegistrationOptionsModel.RegistrationID = core.StringPtr("testString")
				deleteRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := partnerCenterSellService.DeleteRegistration(deleteRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRegistrationOptions model with no property values
				deleteRegistrationOptionsModelNew := new(partnercentersellv1.DeleteRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = partnerCenterSellService.DeleteRegistration(deleteRegistrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateOnboardingProduct(createOnboardingProductOptions *CreateOnboardingProductOptions) - Operation response error`, func() {
		createOnboardingProductPath := "/products"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOnboardingProductPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateOnboardingProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("name")
				primaryContactModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the CreateOnboardingProductOptions model
				createOnboardingProductOptionsModel := new(partnercentersellv1.CreateOnboardingProductOptions)
				createOnboardingProductOptionsModel.Type = core.StringPtr("service")
				createOnboardingProductOptionsModel.PrimaryContact = primaryContactModel
				createOnboardingProductOptionsModel.EccnNumber = core.StringPtr("testString")
				createOnboardingProductOptionsModel.EroClass = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Unspsc = core.Float64Ptr(float64(72.5))
				createOnboardingProductOptionsModel.TaxAssessment = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Support = onboardingProductSupportModel
				createOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreateOnboardingProduct(createOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreateOnboardingProduct(createOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateOnboardingProduct(createOnboardingProductOptions *CreateOnboardingProductOptions)`, func() {
		createOnboardingProductPath := "/products"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOnboardingProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "type": "software", "primary_contact": {"name": "Name", "email": "Email"}, "private_catalog_id": "PrivateCatalogID", "private_catalog_offering_id": "PrivateCatalogOfferingID", "global_catalog_offering_id": "GlobalCatalogOfferingID", "staging_global_catalog_offering_id": "StagingGlobalCatalogOfferingID", "approver_resource_id": "4eb40ee6-d5d6-4328-a52e-06654eab8775", "iam_registration_id": "IamRegistrationID", "eccn_number": "EccnNumber", "ero_class": "EroClass", "unspsc": 6, "tax_assessment": "TaxAssessment", "support": {"escalation_contacts": [{"name": "Name", "email": "Email", "role": "Role"}]}}`)
				}))
			})
			It(`Invoke CreateOnboardingProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("name")
				primaryContactModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the CreateOnboardingProductOptions model
				createOnboardingProductOptionsModel := new(partnercentersellv1.CreateOnboardingProductOptions)
				createOnboardingProductOptionsModel.Type = core.StringPtr("service")
				createOnboardingProductOptionsModel.PrimaryContact = primaryContactModel
				createOnboardingProductOptionsModel.EccnNumber = core.StringPtr("testString")
				createOnboardingProductOptionsModel.EroClass = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Unspsc = core.Float64Ptr(float64(72.5))
				createOnboardingProductOptionsModel.TaxAssessment = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Support = onboardingProductSupportModel
				createOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreateOnboardingProductWithContext(ctx, createOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreateOnboardingProduct(createOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreateOnboardingProductWithContext(ctx, createOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOnboardingProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "type": "software", "primary_contact": {"name": "Name", "email": "Email"}, "private_catalog_id": "PrivateCatalogID", "private_catalog_offering_id": "PrivateCatalogOfferingID", "global_catalog_offering_id": "GlobalCatalogOfferingID", "staging_global_catalog_offering_id": "StagingGlobalCatalogOfferingID", "approver_resource_id": "4eb40ee6-d5d6-4328-a52e-06654eab8775", "iam_registration_id": "IamRegistrationID", "eccn_number": "EccnNumber", "ero_class": "EroClass", "unspsc": 6, "tax_assessment": "TaxAssessment", "support": {"escalation_contacts": [{"name": "Name", "email": "Email", "role": "Role"}]}}`)
				}))
			})
			It(`Invoke CreateOnboardingProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreateOnboardingProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("name")
				primaryContactModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the CreateOnboardingProductOptions model
				createOnboardingProductOptionsModel := new(partnercentersellv1.CreateOnboardingProductOptions)
				createOnboardingProductOptionsModel.Type = core.StringPtr("service")
				createOnboardingProductOptionsModel.PrimaryContact = primaryContactModel
				createOnboardingProductOptionsModel.EccnNumber = core.StringPtr("testString")
				createOnboardingProductOptionsModel.EroClass = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Unspsc = core.Float64Ptr(float64(72.5))
				createOnboardingProductOptionsModel.TaxAssessment = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Support = onboardingProductSupportModel
				createOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreateOnboardingProduct(createOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateOnboardingProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("name")
				primaryContactModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the CreateOnboardingProductOptions model
				createOnboardingProductOptionsModel := new(partnercentersellv1.CreateOnboardingProductOptions)
				createOnboardingProductOptionsModel.Type = core.StringPtr("service")
				createOnboardingProductOptionsModel.PrimaryContact = primaryContactModel
				createOnboardingProductOptionsModel.EccnNumber = core.StringPtr("testString")
				createOnboardingProductOptionsModel.EroClass = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Unspsc = core.Float64Ptr(float64(72.5))
				createOnboardingProductOptionsModel.TaxAssessment = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Support = onboardingProductSupportModel
				createOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreateOnboardingProduct(createOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateOnboardingProductOptions model with no property values
				createOnboardingProductOptionsModelNew := new(partnercentersellv1.CreateOnboardingProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreateOnboardingProduct(createOnboardingProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateOnboardingProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("name")
				primaryContactModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the CreateOnboardingProductOptions model
				createOnboardingProductOptionsModel := new(partnercentersellv1.CreateOnboardingProductOptions)
				createOnboardingProductOptionsModel.Type = core.StringPtr("service")
				createOnboardingProductOptionsModel.PrimaryContact = primaryContactModel
				createOnboardingProductOptionsModel.EccnNumber = core.StringPtr("testString")
				createOnboardingProductOptionsModel.EroClass = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Unspsc = core.Float64Ptr(float64(72.5))
				createOnboardingProductOptionsModel.TaxAssessment = core.StringPtr("testString")
				createOnboardingProductOptionsModel.Support = onboardingProductSupportModel
				createOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreateOnboardingProduct(createOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetOnboardingProduct(getOnboardingProductOptions *GetOnboardingProductOptions) - Operation response error`, func() {
		getOnboardingProductPath := "/products/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOnboardingProductPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetOnboardingProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetOnboardingProductOptions model
				getOnboardingProductOptionsModel := new(partnercentersellv1.GetOnboardingProductOptions)
				getOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				getOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetOnboardingProduct(getOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetOnboardingProduct(getOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetOnboardingProduct(getOnboardingProductOptions *GetOnboardingProductOptions)`, func() {
		getOnboardingProductPath := "/products/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOnboardingProductPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "type": "software", "primary_contact": {"name": "Name", "email": "Email"}, "private_catalog_id": "PrivateCatalogID", "private_catalog_offering_id": "PrivateCatalogOfferingID", "global_catalog_offering_id": "GlobalCatalogOfferingID", "staging_global_catalog_offering_id": "StagingGlobalCatalogOfferingID", "approver_resource_id": "4eb40ee6-d5d6-4328-a52e-06654eab8775", "iam_registration_id": "IamRegistrationID", "eccn_number": "EccnNumber", "ero_class": "EroClass", "unspsc": 6, "tax_assessment": "TaxAssessment", "support": {"escalation_contacts": [{"name": "Name", "email": "Email", "role": "Role"}]}}`)
				}))
			})
			It(`Invoke GetOnboardingProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetOnboardingProductOptions model
				getOnboardingProductOptionsModel := new(partnercentersellv1.GetOnboardingProductOptions)
				getOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				getOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetOnboardingProductWithContext(ctx, getOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetOnboardingProduct(getOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetOnboardingProductWithContext(ctx, getOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOnboardingProductPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "type": "software", "primary_contact": {"name": "Name", "email": "Email"}, "private_catalog_id": "PrivateCatalogID", "private_catalog_offering_id": "PrivateCatalogOfferingID", "global_catalog_offering_id": "GlobalCatalogOfferingID", "staging_global_catalog_offering_id": "StagingGlobalCatalogOfferingID", "approver_resource_id": "4eb40ee6-d5d6-4328-a52e-06654eab8775", "iam_registration_id": "IamRegistrationID", "eccn_number": "EccnNumber", "ero_class": "EroClass", "unspsc": 6, "tax_assessment": "TaxAssessment", "support": {"escalation_contacts": [{"name": "Name", "email": "Email", "role": "Role"}]}}`)
				}))
			})
			It(`Invoke GetOnboardingProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetOnboardingProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOnboardingProductOptions model
				getOnboardingProductOptionsModel := new(partnercentersellv1.GetOnboardingProductOptions)
				getOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				getOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetOnboardingProduct(getOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetOnboardingProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetOnboardingProductOptions model
				getOnboardingProductOptionsModel := new(partnercentersellv1.GetOnboardingProductOptions)
				getOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				getOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetOnboardingProduct(getOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetOnboardingProductOptions model with no property values
				getOnboardingProductOptionsModelNew := new(partnercentersellv1.GetOnboardingProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetOnboardingProduct(getOnboardingProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetOnboardingProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetOnboardingProductOptions model
				getOnboardingProductOptionsModel := new(partnercentersellv1.GetOnboardingProductOptions)
				getOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				getOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetOnboardingProduct(getOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateOnboardingProduct(updateOnboardingProductOptions *UpdateOnboardingProductOptions) - Operation response error`, func() {
		updateOnboardingProductPath := "/products/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOnboardingProductPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateOnboardingProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the OnboardingProductPatch model
				onboardingProductPatchModel := new(partnercentersellv1.OnboardingProductPatch)
				onboardingProductPatchModel.PrimaryContact = primaryContactModel
				onboardingProductPatchModel.EccnNumber = core.StringPtr("testString")
				onboardingProductPatchModel.EroClass = core.StringPtr("testString")
				onboardingProductPatchModel.Unspsc = core.Float64Ptr(float64(12345))
				onboardingProductPatchModel.TaxAssessment = core.StringPtr("testString")
				onboardingProductPatchModel.Support = onboardingProductSupportModel
				onboardingProductPatchModelAsPatch, asPatchErr := onboardingProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateOnboardingProductOptions model
				updateOnboardingProductOptionsModel := new(partnercentersellv1.UpdateOnboardingProductOptions)
				updateOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				updateOnboardingProductOptionsModel.OnboardingProductPatch = onboardingProductPatchModelAsPatch
				updateOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateOnboardingProduct(updateOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateOnboardingProduct(updateOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateOnboardingProduct(updateOnboardingProductOptions *UpdateOnboardingProductOptions)`, func() {
		updateOnboardingProductPath := "/products/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOnboardingProductPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "type": "software", "primary_contact": {"name": "Name", "email": "Email"}, "private_catalog_id": "PrivateCatalogID", "private_catalog_offering_id": "PrivateCatalogOfferingID", "global_catalog_offering_id": "GlobalCatalogOfferingID", "staging_global_catalog_offering_id": "StagingGlobalCatalogOfferingID", "approver_resource_id": "4eb40ee6-d5d6-4328-a52e-06654eab8775", "iam_registration_id": "IamRegistrationID", "eccn_number": "EccnNumber", "ero_class": "EroClass", "unspsc": 6, "tax_assessment": "TaxAssessment", "support": {"escalation_contacts": [{"name": "Name", "email": "Email", "role": "Role"}]}}`)
				}))
			})
			It(`Invoke UpdateOnboardingProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the OnboardingProductPatch model
				onboardingProductPatchModel := new(partnercentersellv1.OnboardingProductPatch)
				onboardingProductPatchModel.PrimaryContact = primaryContactModel
				onboardingProductPatchModel.EccnNumber = core.StringPtr("testString")
				onboardingProductPatchModel.EroClass = core.StringPtr("testString")
				onboardingProductPatchModel.Unspsc = core.Float64Ptr(float64(12345))
				onboardingProductPatchModel.TaxAssessment = core.StringPtr("testString")
				onboardingProductPatchModel.Support = onboardingProductSupportModel
				onboardingProductPatchModelAsPatch, asPatchErr := onboardingProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateOnboardingProductOptions model
				updateOnboardingProductOptionsModel := new(partnercentersellv1.UpdateOnboardingProductOptions)
				updateOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				updateOnboardingProductOptionsModel.OnboardingProductPatch = onboardingProductPatchModelAsPatch
				updateOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateOnboardingProductWithContext(ctx, updateOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateOnboardingProduct(updateOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateOnboardingProductWithContext(ctx, updateOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOnboardingProductPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "type": "software", "primary_contact": {"name": "Name", "email": "Email"}, "private_catalog_id": "PrivateCatalogID", "private_catalog_offering_id": "PrivateCatalogOfferingID", "global_catalog_offering_id": "GlobalCatalogOfferingID", "staging_global_catalog_offering_id": "StagingGlobalCatalogOfferingID", "approver_resource_id": "4eb40ee6-d5d6-4328-a52e-06654eab8775", "iam_registration_id": "IamRegistrationID", "eccn_number": "EccnNumber", "ero_class": "EroClass", "unspsc": 6, "tax_assessment": "TaxAssessment", "support": {"escalation_contacts": [{"name": "Name", "email": "Email", "role": "Role"}]}}`)
				}))
			})
			It(`Invoke UpdateOnboardingProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateOnboardingProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the OnboardingProductPatch model
				onboardingProductPatchModel := new(partnercentersellv1.OnboardingProductPatch)
				onboardingProductPatchModel.PrimaryContact = primaryContactModel
				onboardingProductPatchModel.EccnNumber = core.StringPtr("testString")
				onboardingProductPatchModel.EroClass = core.StringPtr("testString")
				onboardingProductPatchModel.Unspsc = core.Float64Ptr(float64(12345))
				onboardingProductPatchModel.TaxAssessment = core.StringPtr("testString")
				onboardingProductPatchModel.Support = onboardingProductSupportModel
				onboardingProductPatchModelAsPatch, asPatchErr := onboardingProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateOnboardingProductOptions model
				updateOnboardingProductOptionsModel := new(partnercentersellv1.UpdateOnboardingProductOptions)
				updateOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				updateOnboardingProductOptionsModel.OnboardingProductPatch = onboardingProductPatchModelAsPatch
				updateOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateOnboardingProduct(updateOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateOnboardingProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the OnboardingProductPatch model
				onboardingProductPatchModel := new(partnercentersellv1.OnboardingProductPatch)
				onboardingProductPatchModel.PrimaryContact = primaryContactModel
				onboardingProductPatchModel.EccnNumber = core.StringPtr("testString")
				onboardingProductPatchModel.EroClass = core.StringPtr("testString")
				onboardingProductPatchModel.Unspsc = core.Float64Ptr(float64(12345))
				onboardingProductPatchModel.TaxAssessment = core.StringPtr("testString")
				onboardingProductPatchModel.Support = onboardingProductSupportModel
				onboardingProductPatchModelAsPatch, asPatchErr := onboardingProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateOnboardingProductOptions model
				updateOnboardingProductOptionsModel := new(partnercentersellv1.UpdateOnboardingProductOptions)
				updateOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				updateOnboardingProductOptionsModel.OnboardingProductPatch = onboardingProductPatchModelAsPatch
				updateOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateOnboardingProduct(updateOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateOnboardingProductOptions model with no property values
				updateOnboardingProductOptionsModelNew := new(partnercentersellv1.UpdateOnboardingProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateOnboardingProduct(updateOnboardingProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateOnboardingProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				primaryContactModel.Name = core.StringPtr("testString")
				primaryContactModel.Email = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

				// Construct an instance of the OnboardingProductPatch model
				onboardingProductPatchModel := new(partnercentersellv1.OnboardingProductPatch)
				onboardingProductPatchModel.PrimaryContact = primaryContactModel
				onboardingProductPatchModel.EccnNumber = core.StringPtr("testString")
				onboardingProductPatchModel.EroClass = core.StringPtr("testString")
				onboardingProductPatchModel.Unspsc = core.Float64Ptr(float64(12345))
				onboardingProductPatchModel.TaxAssessment = core.StringPtr("testString")
				onboardingProductPatchModel.Support = onboardingProductSupportModel
				onboardingProductPatchModelAsPatch, asPatchErr := onboardingProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateOnboardingProductOptions model
				updateOnboardingProductOptionsModel := new(partnercentersellv1.UpdateOnboardingProductOptions)
				updateOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				updateOnboardingProductOptionsModel.OnboardingProductPatch = onboardingProductPatchModelAsPatch
				updateOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateOnboardingProduct(updateOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteOnboardingProduct(deleteOnboardingProductOptions *DeleteOnboardingProductOptions)`, func() {
		deleteOnboardingProductPath := "/products/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteOnboardingProductPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteOnboardingProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := partnerCenterSellService.DeleteOnboardingProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteOnboardingProductOptions model
				deleteOnboardingProductOptionsModel := new(partnercentersellv1.DeleteOnboardingProductOptions)
				deleteOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				deleteOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = partnerCenterSellService.DeleteOnboardingProduct(deleteOnboardingProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteOnboardingProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeleteOnboardingProductOptions model
				deleteOnboardingProductOptionsModel := new(partnercentersellv1.DeleteOnboardingProductOptions)
				deleteOnboardingProductOptionsModel.ProductID = core.StringPtr("testString")
				deleteOnboardingProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := partnerCenterSellService.DeleteOnboardingProduct(deleteOnboardingProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteOnboardingProductOptions model with no property values
				deleteOnboardingProductOptionsModelNew := new(partnercentersellv1.DeleteOnboardingProductOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = partnerCenterSellService.DeleteOnboardingProduct(deleteOnboardingProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogProduct(createCatalogProductOptions *CreateCatalogProductOptions) - Operation response error`, func() {
		createCatalogProductPath := "/products/testString/catalog_products"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogProductPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCatalogProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My product display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My product description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My product description long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the CreateCatalogProductOptions model
				createCatalogProductOptionsModel := new(partnercentersellv1.CreateCatalogProductOptions)
				createCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogProductOptionsModel.Name = core.StringPtr("1p-service-08-06")
				createCatalogProductOptionsModel.Active = core.BoolPtr(true)
				createCatalogProductOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogProductOptionsModel.Kind = core.StringPtr("service")
				createCatalogProductOptionsModel.Tags = []string{"keyword", "support_ibm"}
				createCatalogProductOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogProductOptionsModel.ID = core.StringPtr("testString")
				createCatalogProductOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogProductOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogProductOptionsModel.Images = globalCatalogProductImagesModel
				createCatalogProductOptionsModel.Metadata = globalCatalogProductMetadataModel
				createCatalogProductOptionsModel.Env = core.StringPtr("testString")
				createCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogProduct(createCatalogProductOptions *CreateCatalogProductOptions)`, func() {
		createCatalogProductPath := "/products/testString/catalog_products"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "service", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "images": {"image": "Image"}, "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "other": {"PC": {"support": {"url": "URL", "status_url": "StatusURL", "locations": ["Locations"], "languages": ["Languages"], "process": "Process", "process_i18n": {"mapKey": "Inner"}, "support_type": "community", "support_escalation": {"contact": "Contact", "escalation_wait_time": {"value": 5, "type": "Type"}, "response_wait_time": {"value": 5, "type": "Type"}}, "support_details": [{"type": "support_site", "contact": "Contact", "response_wait_time": {"value": 5, "type": "Type"}, "availability": {"times": [{"day": 3, "start_time": "StartTime", "end_time": "EndTime"}], "timezone": "Timezone", "always_available": false}}]}}, "composite": {"composite_kind": "service", "composite_tag": "CompositeTag", "children": [{"kind": "service", "name": "Name"}]}}}}`)
				}))
			})
			It(`Invoke CreateCatalogProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My product display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My product description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My product description long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the CreateCatalogProductOptions model
				createCatalogProductOptionsModel := new(partnercentersellv1.CreateCatalogProductOptions)
				createCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogProductOptionsModel.Name = core.StringPtr("1p-service-08-06")
				createCatalogProductOptionsModel.Active = core.BoolPtr(true)
				createCatalogProductOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogProductOptionsModel.Kind = core.StringPtr("service")
				createCatalogProductOptionsModel.Tags = []string{"keyword", "support_ibm"}
				createCatalogProductOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogProductOptionsModel.ID = core.StringPtr("testString")
				createCatalogProductOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogProductOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogProductOptionsModel.Images = globalCatalogProductImagesModel
				createCatalogProductOptionsModel.Metadata = globalCatalogProductMetadataModel
				createCatalogProductOptionsModel.Env = core.StringPtr("testString")
				createCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreateCatalogProductWithContext(ctx, createCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreateCatalogProductWithContext(ctx, createCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "service", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "images": {"image": "Image"}, "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "other": {"PC": {"support": {"url": "URL", "status_url": "StatusURL", "locations": ["Locations"], "languages": ["Languages"], "process": "Process", "process_i18n": {"mapKey": "Inner"}, "support_type": "community", "support_escalation": {"contact": "Contact", "escalation_wait_time": {"value": 5, "type": "Type"}, "response_wait_time": {"value": 5, "type": "Type"}}, "support_details": [{"type": "support_site", "contact": "Contact", "response_wait_time": {"value": 5, "type": "Type"}, "availability": {"times": [{"day": 3, "start_time": "StartTime", "end_time": "EndTime"}], "timezone": "Timezone", "always_available": false}}]}}, "composite": {"composite_kind": "service", "composite_tag": "CompositeTag", "children": [{"kind": "service", "name": "Name"}]}}}}`)
				}))
			})
			It(`Invoke CreateCatalogProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreateCatalogProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My product display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My product description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My product description long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the CreateCatalogProductOptions model
				createCatalogProductOptionsModel := new(partnercentersellv1.CreateCatalogProductOptions)
				createCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogProductOptionsModel.Name = core.StringPtr("1p-service-08-06")
				createCatalogProductOptionsModel.Active = core.BoolPtr(true)
				createCatalogProductOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogProductOptionsModel.Kind = core.StringPtr("service")
				createCatalogProductOptionsModel.Tags = []string{"keyword", "support_ibm"}
				createCatalogProductOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogProductOptionsModel.ID = core.StringPtr("testString")
				createCatalogProductOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogProductOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogProductOptionsModel.Images = globalCatalogProductImagesModel
				createCatalogProductOptionsModel.Metadata = globalCatalogProductMetadataModel
				createCatalogProductOptionsModel.Env = core.StringPtr("testString")
				createCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCatalogProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My product display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My product description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My product description long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the CreateCatalogProductOptions model
				createCatalogProductOptionsModel := new(partnercentersellv1.CreateCatalogProductOptions)
				createCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogProductOptionsModel.Name = core.StringPtr("1p-service-08-06")
				createCatalogProductOptionsModel.Active = core.BoolPtr(true)
				createCatalogProductOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogProductOptionsModel.Kind = core.StringPtr("service")
				createCatalogProductOptionsModel.Tags = []string{"keyword", "support_ibm"}
				createCatalogProductOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogProductOptionsModel.ID = core.StringPtr("testString")
				createCatalogProductOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogProductOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogProductOptionsModel.Images = globalCatalogProductImagesModel
				createCatalogProductOptionsModel.Metadata = globalCatalogProductMetadataModel
				createCatalogProductOptionsModel.Env = core.StringPtr("testString")
				createCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCatalogProductOptions model with no property values
				createCatalogProductOptionsModelNew := new(partnercentersellv1.CreateCatalogProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateCatalogProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My product display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My product description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My product description long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the CreateCatalogProductOptions model
				createCatalogProductOptionsModel := new(partnercentersellv1.CreateCatalogProductOptions)
				createCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogProductOptionsModel.Name = core.StringPtr("1p-service-08-06")
				createCatalogProductOptionsModel.Active = core.BoolPtr(true)
				createCatalogProductOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogProductOptionsModel.Kind = core.StringPtr("service")
				createCatalogProductOptionsModel.Tags = []string{"keyword", "support_ibm"}
				createCatalogProductOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogProductOptionsModel.ID = core.StringPtr("testString")
				createCatalogProductOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogProductOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogProductOptionsModel.Images = globalCatalogProductImagesModel
				createCatalogProductOptionsModel.Metadata = globalCatalogProductMetadataModel
				createCatalogProductOptionsModel.Env = core.StringPtr("testString")
				createCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogProduct(getCatalogProductOptions *GetCatalogProductOptions) - Operation response error`, func() {
		getCatalogProductPath := "/products/testString/catalog_products/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogProductPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalogProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogProductOptions model
				getCatalogProductOptionsModel := new(partnercentersellv1.GetCatalogProductOptions)
				getCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.Env = core.StringPtr("testString")
				getCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetCatalogProduct(getCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetCatalogProduct(getCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogProduct(getCatalogProductOptions *GetCatalogProductOptions)`, func() {
		getCatalogProductPath := "/products/testString/catalog_products/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogProductPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "service", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "images": {"image": "Image"}, "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "other": {"PC": {"support": {"url": "URL", "status_url": "StatusURL", "locations": ["Locations"], "languages": ["Languages"], "process": "Process", "process_i18n": {"mapKey": "Inner"}, "support_type": "community", "support_escalation": {"contact": "Contact", "escalation_wait_time": {"value": 5, "type": "Type"}, "response_wait_time": {"value": 5, "type": "Type"}}, "support_details": [{"type": "support_site", "contact": "Contact", "response_wait_time": {"value": 5, "type": "Type"}, "availability": {"times": [{"day": 3, "start_time": "StartTime", "end_time": "EndTime"}], "timezone": "Timezone", "always_available": false}}]}}, "composite": {"composite_kind": "service", "composite_tag": "CompositeTag", "children": [{"kind": "service", "name": "Name"}]}}}}`)
				}))
			})
			It(`Invoke GetCatalogProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetCatalogProductOptions model
				getCatalogProductOptionsModel := new(partnercentersellv1.GetCatalogProductOptions)
				getCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.Env = core.StringPtr("testString")
				getCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetCatalogProductWithContext(ctx, getCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetCatalogProduct(getCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetCatalogProductWithContext(ctx, getCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogProductPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "service", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "images": {"image": "Image"}, "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "other": {"PC": {"support": {"url": "URL", "status_url": "StatusURL", "locations": ["Locations"], "languages": ["Languages"], "process": "Process", "process_i18n": {"mapKey": "Inner"}, "support_type": "community", "support_escalation": {"contact": "Contact", "escalation_wait_time": {"value": 5, "type": "Type"}, "response_wait_time": {"value": 5, "type": "Type"}}, "support_details": [{"type": "support_site", "contact": "Contact", "response_wait_time": {"value": 5, "type": "Type"}, "availability": {"times": [{"day": 3, "start_time": "StartTime", "end_time": "EndTime"}], "timezone": "Timezone", "always_available": false}}]}}, "composite": {"composite_kind": "service", "composite_tag": "CompositeTag", "children": [{"kind": "service", "name": "Name"}]}}}}`)
				}))
			})
			It(`Invoke GetCatalogProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetCatalogProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogProductOptions model
				getCatalogProductOptionsModel := new(partnercentersellv1.GetCatalogProductOptions)
				getCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.Env = core.StringPtr("testString")
				getCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetCatalogProduct(getCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCatalogProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogProductOptions model
				getCatalogProductOptionsModel := new(partnercentersellv1.GetCatalogProductOptions)
				getCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.Env = core.StringPtr("testString")
				getCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetCatalogProduct(getCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCatalogProductOptions model with no property values
				getCatalogProductOptionsModelNew := new(partnercentersellv1.GetCatalogProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetCatalogProduct(getCatalogProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCatalogProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogProductOptions model
				getCatalogProductOptionsModel := new(partnercentersellv1.GetCatalogProductOptions)
				getCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogProductOptionsModel.Env = core.StringPtr("testString")
				getCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetCatalogProduct(getCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogProduct(updateCatalogProductOptions *UpdateCatalogProductOptions) - Operation response error`, func() {
		updateCatalogProductPath := "/products/testString/catalog_products/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogProductPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCatalogProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My updated display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the GlobalCatalogProductPatch model
				globalCatalogProductPatchModel := new(partnercentersellv1.GlobalCatalogProductPatch)
				globalCatalogProductPatchModel.Active = core.BoolPtr(true)
				globalCatalogProductPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogProductPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogProductPatchModel.Tags = []string{"testString"}
				globalCatalogProductPatchModel.Images = globalCatalogProductImagesModel
				globalCatalogProductPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogProductPatchModel.Metadata = globalCatalogProductMetadataModel
				globalCatalogProductPatchModelAsPatch, asPatchErr := globalCatalogProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogProductOptions model
				updateCatalogProductOptionsModel := new(partnercentersellv1.UpdateCatalogProductOptions)
				updateCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.GlobalCatalogProductPatch = globalCatalogProductPatchModelAsPatch
				updateCatalogProductOptionsModel.Env = core.StringPtr("testString")
				updateCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogProduct(updateCatalogProductOptions *UpdateCatalogProductOptions)`, func() {
		updateCatalogProductPath := "/products/testString/catalog_products/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogProductPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "service", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "images": {"image": "Image"}, "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "other": {"PC": {"support": {"url": "URL", "status_url": "StatusURL", "locations": ["Locations"], "languages": ["Languages"], "process": "Process", "process_i18n": {"mapKey": "Inner"}, "support_type": "community", "support_escalation": {"contact": "Contact", "escalation_wait_time": {"value": 5, "type": "Type"}, "response_wait_time": {"value": 5, "type": "Type"}}, "support_details": [{"type": "support_site", "contact": "Contact", "response_wait_time": {"value": 5, "type": "Type"}, "availability": {"times": [{"day": 3, "start_time": "StartTime", "end_time": "EndTime"}], "timezone": "Timezone", "always_available": false}}]}}, "composite": {"composite_kind": "service", "composite_tag": "CompositeTag", "children": [{"kind": "service", "name": "Name"}]}}}}`)
				}))
			})
			It(`Invoke UpdateCatalogProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My updated display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the GlobalCatalogProductPatch model
				globalCatalogProductPatchModel := new(partnercentersellv1.GlobalCatalogProductPatch)
				globalCatalogProductPatchModel.Active = core.BoolPtr(true)
				globalCatalogProductPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogProductPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogProductPatchModel.Tags = []string{"testString"}
				globalCatalogProductPatchModel.Images = globalCatalogProductImagesModel
				globalCatalogProductPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogProductPatchModel.Metadata = globalCatalogProductMetadataModel
				globalCatalogProductPatchModelAsPatch, asPatchErr := globalCatalogProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogProductOptions model
				updateCatalogProductOptionsModel := new(partnercentersellv1.UpdateCatalogProductOptions)
				updateCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.GlobalCatalogProductPatch = globalCatalogProductPatchModelAsPatch
				updateCatalogProductOptionsModel.Env = core.StringPtr("testString")
				updateCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateCatalogProductWithContext(ctx, updateCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateCatalogProductWithContext(ctx, updateCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogProductPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "service", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "images": {"image": "Image"}, "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "other": {"PC": {"support": {"url": "URL", "status_url": "StatusURL", "locations": ["Locations"], "languages": ["Languages"], "process": "Process", "process_i18n": {"mapKey": "Inner"}, "support_type": "community", "support_escalation": {"contact": "Contact", "escalation_wait_time": {"value": 5, "type": "Type"}, "response_wait_time": {"value": 5, "type": "Type"}}, "support_details": [{"type": "support_site", "contact": "Contact", "response_wait_time": {"value": 5, "type": "Type"}, "availability": {"times": [{"day": 3, "start_time": "StartTime", "end_time": "EndTime"}], "timezone": "Timezone", "always_available": false}}]}}, "composite": {"composite_kind": "service", "composite_tag": "CompositeTag", "children": [{"kind": "service", "name": "Name"}]}}}}`)
				}))
			})
			It(`Invoke UpdateCatalogProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateCatalogProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My updated display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the GlobalCatalogProductPatch model
				globalCatalogProductPatchModel := new(partnercentersellv1.GlobalCatalogProductPatch)
				globalCatalogProductPatchModel.Active = core.BoolPtr(true)
				globalCatalogProductPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogProductPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogProductPatchModel.Tags = []string{"testString"}
				globalCatalogProductPatchModel.Images = globalCatalogProductImagesModel
				globalCatalogProductPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogProductPatchModel.Metadata = globalCatalogProductMetadataModel
				globalCatalogProductPatchModelAsPatch, asPatchErr := globalCatalogProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogProductOptions model
				updateCatalogProductOptionsModel := new(partnercentersellv1.UpdateCatalogProductOptions)
				updateCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.GlobalCatalogProductPatch = globalCatalogProductPatchModelAsPatch
				updateCatalogProductOptionsModel.Env = core.StringPtr("testString")
				updateCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCatalogProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My updated display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the GlobalCatalogProductPatch model
				globalCatalogProductPatchModel := new(partnercentersellv1.GlobalCatalogProductPatch)
				globalCatalogProductPatchModel.Active = core.BoolPtr(true)
				globalCatalogProductPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogProductPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogProductPatchModel.Tags = []string{"testString"}
				globalCatalogProductPatchModel.Images = globalCatalogProductImagesModel
				globalCatalogProductPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogProductPatchModel.Metadata = globalCatalogProductMetadataModel
				globalCatalogProductPatchModelAsPatch, asPatchErr := globalCatalogProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogProductOptions model
				updateCatalogProductOptionsModel := new(partnercentersellv1.UpdateCatalogProductOptions)
				updateCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.GlobalCatalogProductPatch = globalCatalogProductPatchModelAsPatch
				updateCatalogProductOptionsModel.Env = core.StringPtr("testString")
				updateCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCatalogProductOptions model with no property values
				updateCatalogProductOptionsModelNew := new(partnercentersellv1.UpdateCatalogProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCatalogProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My updated display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel

				// Construct an instance of the GlobalCatalogProductPatch model
				globalCatalogProductPatchModel := new(partnercentersellv1.GlobalCatalogProductPatch)
				globalCatalogProductPatchModel.Active = core.BoolPtr(true)
				globalCatalogProductPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogProductPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogProductPatchModel.Tags = []string{"testString"}
				globalCatalogProductPatchModel.Images = globalCatalogProductImagesModel
				globalCatalogProductPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogProductPatchModel.Metadata = globalCatalogProductMetadataModel
				globalCatalogProductPatchModelAsPatch, asPatchErr := globalCatalogProductPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogProductOptions model
				updateCatalogProductOptionsModel := new(partnercentersellv1.UpdateCatalogProductOptions)
				updateCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogProductOptionsModel.GlobalCatalogProductPatch = globalCatalogProductPatchModelAsPatch
				updateCatalogProductOptionsModel.Env = core.StringPtr("testString")
				updateCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCatalogProduct(deleteCatalogProductOptions *DeleteCatalogProductOptions)`, func() {
		deleteCatalogProductPath := "/products/testString/catalog_products/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCatalogProductPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCatalogProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := partnerCenterSellService.DeleteCatalogProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCatalogProductOptions model
				deleteCatalogProductOptionsModel := new(partnercentersellv1.DeleteCatalogProductOptions)
				deleteCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				deleteCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				deleteCatalogProductOptionsModel.Env = core.StringPtr("testString")
				deleteCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = partnerCenterSellService.DeleteCatalogProduct(deleteCatalogProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCatalogProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeleteCatalogProductOptions model
				deleteCatalogProductOptionsModel := new(partnercentersellv1.DeleteCatalogProductOptions)
				deleteCatalogProductOptionsModel.ProductID = core.StringPtr("testString")
				deleteCatalogProductOptionsModel.CatalogProductID = core.StringPtr("testString")
				deleteCatalogProductOptionsModel.Env = core.StringPtr("testString")
				deleteCatalogProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := partnerCenterSellService.DeleteCatalogProduct(deleteCatalogProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCatalogProductOptions model with no property values
				deleteCatalogProductOptionsModelNew := new(partnercentersellv1.DeleteCatalogProductOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = partnerCenterSellService.DeleteCatalogProduct(deleteCatalogProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogPlan(createCatalogPlanOptions *CreateCatalogPlanOptions) - Operation response error`, func() {
		createCatalogPlanPath := "/products/testString/catalog_products/testString/catalog_plans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogPlanPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCatalogPlan with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My plan")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My plan description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My plan long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(false)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("paid")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the CreateCatalogPlanOptions model
				createCatalogPlanOptionsModel := new(partnercentersellv1.CreateCatalogPlanOptions)
				createCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Name = core.StringPtr("free-plan2")
				createCatalogPlanOptionsModel.Active = core.BoolPtr(true)
				createCatalogPlanOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogPlanOptionsModel.Kind = core.StringPtr("plan")
				createCatalogPlanOptionsModel.Tags = []string{"ibm_created"}
				createCatalogPlanOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogPlanOptionsModel.ID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogPlanOptionsModel.Metadata = globalCatalogPlanMetadataModel
				createCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogPlan(createCatalogPlanOptions *CreateCatalogPlanOptions)`, func() {
		createCatalogPlanPath := "/products/testString/catalog_products/testString/catalog_plans"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogPlanPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "plan", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "pricing": {"type": "free", "origin": "global_catalog"}, "plan": {"allow_internal_users": true, "bindable": true}}}`)
				}))
			})
			It(`Invoke CreateCatalogPlan successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My plan")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My plan description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My plan long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(false)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("paid")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the CreateCatalogPlanOptions model
				createCatalogPlanOptionsModel := new(partnercentersellv1.CreateCatalogPlanOptions)
				createCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Name = core.StringPtr("free-plan2")
				createCatalogPlanOptionsModel.Active = core.BoolPtr(true)
				createCatalogPlanOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogPlanOptionsModel.Kind = core.StringPtr("plan")
				createCatalogPlanOptionsModel.Tags = []string{"ibm_created"}
				createCatalogPlanOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogPlanOptionsModel.ID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogPlanOptionsModel.Metadata = globalCatalogPlanMetadataModel
				createCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreateCatalogPlanWithContext(ctx, createCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreateCatalogPlanWithContext(ctx, createCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogPlanPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "plan", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "pricing": {"type": "free", "origin": "global_catalog"}, "plan": {"allow_internal_users": true, "bindable": true}}}`)
				}))
			})
			It(`Invoke CreateCatalogPlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreateCatalogPlan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My plan")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My plan description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My plan long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(false)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("paid")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the CreateCatalogPlanOptions model
				createCatalogPlanOptionsModel := new(partnercentersellv1.CreateCatalogPlanOptions)
				createCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Name = core.StringPtr("free-plan2")
				createCatalogPlanOptionsModel.Active = core.BoolPtr(true)
				createCatalogPlanOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogPlanOptionsModel.Kind = core.StringPtr("plan")
				createCatalogPlanOptionsModel.Tags = []string{"ibm_created"}
				createCatalogPlanOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogPlanOptionsModel.ID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogPlanOptionsModel.Metadata = globalCatalogPlanMetadataModel
				createCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCatalogPlan with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My plan")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My plan description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My plan long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(false)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("paid")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the CreateCatalogPlanOptions model
				createCatalogPlanOptionsModel := new(partnercentersellv1.CreateCatalogPlanOptions)
				createCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Name = core.StringPtr("free-plan2")
				createCatalogPlanOptionsModel.Active = core.BoolPtr(true)
				createCatalogPlanOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogPlanOptionsModel.Kind = core.StringPtr("plan")
				createCatalogPlanOptionsModel.Tags = []string{"ibm_created"}
				createCatalogPlanOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogPlanOptionsModel.ID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogPlanOptionsModel.Metadata = globalCatalogPlanMetadataModel
				createCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCatalogPlanOptions model with no property values
				createCatalogPlanOptionsModelNew := new(partnercentersellv1.CreateCatalogPlanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateCatalogPlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My plan")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My plan description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My plan long description.")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(false)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("paid")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the CreateCatalogPlanOptions model
				createCatalogPlanOptionsModel := new(partnercentersellv1.CreateCatalogPlanOptions)
				createCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Name = core.StringPtr("free-plan2")
				createCatalogPlanOptionsModel.Active = core.BoolPtr(true)
				createCatalogPlanOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogPlanOptionsModel.Kind = core.StringPtr("plan")
				createCatalogPlanOptionsModel.Tags = []string{"ibm_created"}
				createCatalogPlanOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogPlanOptionsModel.ID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogPlanOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogPlanOptionsModel.Metadata = globalCatalogPlanMetadataModel
				createCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				createCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogPlan(getCatalogPlanOptions *GetCatalogPlanOptions) - Operation response error`, func() {
		getCatalogPlanPath := "/products/testString/catalog_products/testString/catalog_plans/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPlanPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalogPlan with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogPlanOptions model
				getCatalogPlanOptionsModel := new(partnercentersellv1.GetCatalogPlanOptions)
				getCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogPlan(getCatalogPlanOptions *GetCatalogPlanOptions)`, func() {
		getCatalogPlanPath := "/products/testString/catalog_products/testString/catalog_plans/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPlanPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "plan", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "pricing": {"type": "free", "origin": "global_catalog"}, "plan": {"allow_internal_users": true, "bindable": true}}}`)
				}))
			})
			It(`Invoke GetCatalogPlan successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetCatalogPlanOptions model
				getCatalogPlanOptionsModel := new(partnercentersellv1.GetCatalogPlanOptions)
				getCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetCatalogPlanWithContext(ctx, getCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetCatalogPlanWithContext(ctx, getCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPlanPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "plan", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "pricing": {"type": "free", "origin": "global_catalog"}, "plan": {"allow_internal_users": true, "bindable": true}}}`)
				}))
			})
			It(`Invoke GetCatalogPlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetCatalogPlan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogPlanOptions model
				getCatalogPlanOptionsModel := new(partnercentersellv1.GetCatalogPlanOptions)
				getCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCatalogPlan with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogPlanOptions model
				getCatalogPlanOptionsModel := new(partnercentersellv1.GetCatalogPlanOptions)
				getCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCatalogPlanOptions model with no property values
				getCatalogPlanOptionsModelNew := new(partnercentersellv1.GetCatalogPlanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCatalogPlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogPlanOptions model
				getCatalogPlanOptionsModel := new(partnercentersellv1.GetCatalogPlanOptions)
				getCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				getCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogPlan(updateCatalogPlanOptions *UpdateCatalogPlanOptions) - Operation response error`, func() {
		updateCatalogPlanPath := "/products/testString/catalog_products/testString/catalog_plans/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogPlanPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCatalogPlan with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("free")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the GlobalCatalogPlanPatch model
				globalCatalogPlanPatchModel := new(partnercentersellv1.GlobalCatalogPlanPatch)
				globalCatalogPlanPatchModel.Active = core.BoolPtr(true)
				globalCatalogPlanPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogPlanPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogPlanPatchModel.Tags = []string{"testString"}
				globalCatalogPlanPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogPlanPatchModel.Metadata = globalCatalogPlanMetadataModel
				globalCatalogPlanPatchModelAsPatch, asPatchErr := globalCatalogPlanPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogPlanOptions model
				updateCatalogPlanOptionsModel := new(partnercentersellv1.UpdateCatalogPlanOptions)
				updateCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.GlobalCatalogPlanPatch = globalCatalogPlanPatchModelAsPatch
				updateCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogPlan(updateCatalogPlanOptions *UpdateCatalogPlanOptions)`, func() {
		updateCatalogPlanPath := "/products/testString/catalog_products/testString/catalog_plans/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogPlanPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "plan", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "pricing": {"type": "free", "origin": "global_catalog"}, "plan": {"allow_internal_users": true, "bindable": true}}}`)
				}))
			})
			It(`Invoke UpdateCatalogPlan successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("free")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the GlobalCatalogPlanPatch model
				globalCatalogPlanPatchModel := new(partnercentersellv1.GlobalCatalogPlanPatch)
				globalCatalogPlanPatchModel.Active = core.BoolPtr(true)
				globalCatalogPlanPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogPlanPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogPlanPatchModel.Tags = []string{"testString"}
				globalCatalogPlanPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogPlanPatchModel.Metadata = globalCatalogPlanMetadataModel
				globalCatalogPlanPatchModelAsPatch, asPatchErr := globalCatalogPlanPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogPlanOptions model
				updateCatalogPlanOptionsModel := new(partnercentersellv1.UpdateCatalogPlanOptions)
				updateCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.GlobalCatalogPlanPatch = globalCatalogPlanPatchModelAsPatch
				updateCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateCatalogPlanWithContext(ctx, updateCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateCatalogPlanWithContext(ctx, updateCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogPlanPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "plan", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "pricing": {"type": "free", "origin": "global_catalog"}, "plan": {"allow_internal_users": true, "bindable": true}}}`)
				}))
			})
			It(`Invoke UpdateCatalogPlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateCatalogPlan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("free")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the GlobalCatalogPlanPatch model
				globalCatalogPlanPatchModel := new(partnercentersellv1.GlobalCatalogPlanPatch)
				globalCatalogPlanPatchModel.Active = core.BoolPtr(true)
				globalCatalogPlanPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogPlanPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogPlanPatchModel.Tags = []string{"testString"}
				globalCatalogPlanPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogPlanPatchModel.Metadata = globalCatalogPlanMetadataModel
				globalCatalogPlanPatchModelAsPatch, asPatchErr := globalCatalogPlanPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogPlanOptions model
				updateCatalogPlanOptionsModel := new(partnercentersellv1.UpdateCatalogPlanOptions)
				updateCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.GlobalCatalogPlanPatch = globalCatalogPlanPatchModelAsPatch
				updateCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCatalogPlan with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("free")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the GlobalCatalogPlanPatch model
				globalCatalogPlanPatchModel := new(partnercentersellv1.GlobalCatalogPlanPatch)
				globalCatalogPlanPatchModel.Active = core.BoolPtr(true)
				globalCatalogPlanPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogPlanPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogPlanPatchModel.Tags = []string{"testString"}
				globalCatalogPlanPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogPlanPatchModel.Metadata = globalCatalogPlanMetadataModel
				globalCatalogPlanPatchModelAsPatch, asPatchErr := globalCatalogPlanPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogPlanOptions model
				updateCatalogPlanOptionsModel := new(partnercentersellv1.UpdateCatalogPlanOptions)
				updateCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.GlobalCatalogPlanPatch = globalCatalogPlanPatchModelAsPatch
				updateCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCatalogPlanOptions model with no property values
				updateCatalogPlanOptionsModelNew := new(partnercentersellv1.UpdateCatalogPlanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCatalogPlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				globalCatalogMetadataPricingModel.Type = core.StringPtr("free")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel

				// Construct an instance of the GlobalCatalogPlanPatch model
				globalCatalogPlanPatchModel := new(partnercentersellv1.GlobalCatalogPlanPatch)
				globalCatalogPlanPatchModel.Active = core.BoolPtr(true)
				globalCatalogPlanPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogPlanPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogPlanPatchModel.Tags = []string{"testString"}
				globalCatalogPlanPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogPlanPatchModel.Metadata = globalCatalogPlanMetadataModel
				globalCatalogPlanPatchModelAsPatch, asPatchErr := globalCatalogPlanPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogPlanOptions model
				updateCatalogPlanOptionsModel := new(partnercentersellv1.UpdateCatalogPlanOptions)
				updateCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.GlobalCatalogPlanPatch = globalCatalogPlanPatchModelAsPatch
				updateCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				updateCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCatalogPlan(deleteCatalogPlanOptions *DeleteCatalogPlanOptions)`, func() {
		deleteCatalogPlanPath := "/products/testString/catalog_products/testString/catalog_plans/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCatalogPlanPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCatalogPlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := partnerCenterSellService.DeleteCatalogPlan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCatalogPlanOptions model
				deleteCatalogPlanOptionsModel := new(partnercentersellv1.DeleteCatalogPlanOptions)
				deleteCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				deleteCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				deleteCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				deleteCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				deleteCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = partnerCenterSellService.DeleteCatalogPlan(deleteCatalogPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCatalogPlan with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeleteCatalogPlanOptions model
				deleteCatalogPlanOptionsModel := new(partnercentersellv1.DeleteCatalogPlanOptions)
				deleteCatalogPlanOptionsModel.ProductID = core.StringPtr("testString")
				deleteCatalogPlanOptionsModel.CatalogProductID = core.StringPtr("testString")
				deleteCatalogPlanOptionsModel.CatalogPlanID = core.StringPtr("testString")
				deleteCatalogPlanOptionsModel.Env = core.StringPtr("testString")
				deleteCatalogPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := partnerCenterSellService.DeleteCatalogPlan(deleteCatalogPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCatalogPlanOptions model with no property values
				deleteCatalogPlanOptionsModelNew := new(partnercentersellv1.DeleteCatalogPlanOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = partnerCenterSellService.DeleteCatalogPlan(deleteCatalogPlanOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogDeployment(createCatalogDeploymentOptions *CreateCatalogDeploymentOptions) - Operation response error`, func() {
		createCatalogDeploymentPath := "/products/testString/catalog_products/testString/catalog_plans/testString/catalog_deployments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogDeploymentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCatalogDeployment with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("brokerunique1234")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3c2")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the CreateCatalogDeploymentOptions model
				createCatalogDeploymentOptionsModel := new(partnercentersellv1.CreateCatalogDeploymentOptions)
				createCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Name = core.StringPtr("deployment-eu-de")
				createCatalogDeploymentOptionsModel.Active = core.BoolPtr(true)
				createCatalogDeploymentOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogDeploymentOptionsModel.Kind = core.StringPtr("deployment")
				createCatalogDeploymentOptionsModel.Tags = []string{"eu-gb"}
				createCatalogDeploymentOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogDeploymentOptionsModel.ID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogDeploymentOptionsModel.Metadata = globalCatalogDeploymentMetadataModel
				createCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogDeployment(createCatalogDeploymentOptions *CreateCatalogDeploymentOptions)`, func() {
		createCatalogDeploymentPath := "/products/testString/catalog_products/testString/catalog_plans/testString/catalog_deployments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogDeploymentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "deployment", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "deployment": {"broker": {"name": "Name", "guid": "Guid"}, "location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn"}}}`)
				}))
			})
			It(`Invoke CreateCatalogDeployment successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("brokerunique1234")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3c2")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the CreateCatalogDeploymentOptions model
				createCatalogDeploymentOptionsModel := new(partnercentersellv1.CreateCatalogDeploymentOptions)
				createCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Name = core.StringPtr("deployment-eu-de")
				createCatalogDeploymentOptionsModel.Active = core.BoolPtr(true)
				createCatalogDeploymentOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogDeploymentOptionsModel.Kind = core.StringPtr("deployment")
				createCatalogDeploymentOptionsModel.Tags = []string{"eu-gb"}
				createCatalogDeploymentOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogDeploymentOptionsModel.ID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogDeploymentOptionsModel.Metadata = globalCatalogDeploymentMetadataModel
				createCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreateCatalogDeploymentWithContext(ctx, createCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreateCatalogDeploymentWithContext(ctx, createCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogDeploymentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "deployment", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "deployment": {"broker": {"name": "Name", "guid": "Guid"}, "location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn"}}}`)
				}))
			})
			It(`Invoke CreateCatalogDeployment successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreateCatalogDeployment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("brokerunique1234")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3c2")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the CreateCatalogDeploymentOptions model
				createCatalogDeploymentOptionsModel := new(partnercentersellv1.CreateCatalogDeploymentOptions)
				createCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Name = core.StringPtr("deployment-eu-de")
				createCatalogDeploymentOptionsModel.Active = core.BoolPtr(true)
				createCatalogDeploymentOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogDeploymentOptionsModel.Kind = core.StringPtr("deployment")
				createCatalogDeploymentOptionsModel.Tags = []string{"eu-gb"}
				createCatalogDeploymentOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogDeploymentOptionsModel.ID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogDeploymentOptionsModel.Metadata = globalCatalogDeploymentMetadataModel
				createCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCatalogDeployment with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("brokerunique1234")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3c2")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the CreateCatalogDeploymentOptions model
				createCatalogDeploymentOptionsModel := new(partnercentersellv1.CreateCatalogDeploymentOptions)
				createCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Name = core.StringPtr("deployment-eu-de")
				createCatalogDeploymentOptionsModel.Active = core.BoolPtr(true)
				createCatalogDeploymentOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogDeploymentOptionsModel.Kind = core.StringPtr("deployment")
				createCatalogDeploymentOptionsModel.Tags = []string{"eu-gb"}
				createCatalogDeploymentOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogDeploymentOptionsModel.ID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogDeploymentOptionsModel.Metadata = globalCatalogDeploymentMetadataModel
				createCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCatalogDeploymentOptions model with no property values
				createCatalogDeploymentOptionsModelNew := new(partnercentersellv1.CreateCatalogDeploymentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateCatalogDeployment successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("brokerunique1234")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3c2")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the CreateCatalogDeploymentOptions model
				createCatalogDeploymentOptionsModel := new(partnercentersellv1.CreateCatalogDeploymentOptions)
				createCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Name = core.StringPtr("deployment-eu-de")
				createCatalogDeploymentOptionsModel.Active = core.BoolPtr(true)
				createCatalogDeploymentOptionsModel.Disabled = core.BoolPtr(false)
				createCatalogDeploymentOptionsModel.Kind = core.StringPtr("deployment")
				createCatalogDeploymentOptionsModel.Tags = []string{"eu-gb"}
				createCatalogDeploymentOptionsModel.ObjectProvider = catalogProductProviderModel
				createCatalogDeploymentOptionsModel.ID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.ObjectID = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.OverviewUi = globalCatalogOverviewUiModel
				createCatalogDeploymentOptionsModel.Metadata = globalCatalogDeploymentMetadataModel
				createCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				createCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogDeployment(getCatalogDeploymentOptions *GetCatalogDeploymentOptions) - Operation response error`, func() {
		getCatalogDeploymentPath := "/products/testString/catalog_products/testString/catalog_plans/testString/catalog_deployments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogDeploymentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalogDeployment with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogDeploymentOptions model
				getCatalogDeploymentOptionsModel := new(partnercentersellv1.GetCatalogDeploymentOptions)
				getCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogDeployment(getCatalogDeploymentOptions *GetCatalogDeploymentOptions)`, func() {
		getCatalogDeploymentPath := "/products/testString/catalog_products/testString/catalog_plans/testString/catalog_deployments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogDeploymentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "deployment", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "deployment": {"broker": {"name": "Name", "guid": "Guid"}, "location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn"}}}`)
				}))
			})
			It(`Invoke GetCatalogDeployment successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetCatalogDeploymentOptions model
				getCatalogDeploymentOptionsModel := new(partnercentersellv1.GetCatalogDeploymentOptions)
				getCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetCatalogDeploymentWithContext(ctx, getCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetCatalogDeploymentWithContext(ctx, getCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogDeploymentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "deployment", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "deployment": {"broker": {"name": "Name", "guid": "Guid"}, "location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn"}}}`)
				}))
			})
			It(`Invoke GetCatalogDeployment successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetCatalogDeployment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogDeploymentOptions model
				getCatalogDeploymentOptionsModel := new(partnercentersellv1.GetCatalogDeploymentOptions)
				getCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCatalogDeployment with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogDeploymentOptions model
				getCatalogDeploymentOptionsModel := new(partnercentersellv1.GetCatalogDeploymentOptions)
				getCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCatalogDeploymentOptions model with no property values
				getCatalogDeploymentOptionsModelNew := new(partnercentersellv1.GetCatalogDeploymentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCatalogDeployment successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogDeploymentOptions model
				getCatalogDeploymentOptionsModel := new(partnercentersellv1.GetCatalogDeploymentOptions)
				getCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				getCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogDeployment(updateCatalogDeploymentOptions *UpdateCatalogDeploymentOptions) - Operation response error`, func() {
		updateCatalogDeploymentPath := "/products/testString/catalog_products/testString/catalog_plans/testString/catalog_deployments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogDeploymentPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCatalogDeployment with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("another-broker")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3cf")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the GlobalCatalogDeploymentPatch model
				globalCatalogDeploymentPatchModel := new(partnercentersellv1.GlobalCatalogDeploymentPatch)
				globalCatalogDeploymentPatchModel.Active = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogDeploymentPatchModel.Tags = []string{"testString"}
				globalCatalogDeploymentPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogDeploymentPatchModel.Metadata = globalCatalogDeploymentMetadataModel
				globalCatalogDeploymentPatchModelAsPatch, asPatchErr := globalCatalogDeploymentPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogDeploymentOptions model
				updateCatalogDeploymentOptionsModel := new(partnercentersellv1.UpdateCatalogDeploymentOptions)
				updateCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.GlobalCatalogDeploymentPatch = globalCatalogDeploymentPatchModelAsPatch
				updateCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogDeployment(updateCatalogDeploymentOptions *UpdateCatalogDeploymentOptions)`, func() {
		updateCatalogDeploymentPath := "/products/testString/catalog_products/testString/catalog_plans/testString/catalog_deployments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogDeploymentPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "deployment", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "deployment": {"broker": {"name": "Name", "guid": "Guid"}, "location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn"}}}`)
				}))
			})
			It(`Invoke UpdateCatalogDeployment successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("another-broker")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3cf")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the GlobalCatalogDeploymentPatch model
				globalCatalogDeploymentPatchModel := new(partnercentersellv1.GlobalCatalogDeploymentPatch)
				globalCatalogDeploymentPatchModel.Active = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogDeploymentPatchModel.Tags = []string{"testString"}
				globalCatalogDeploymentPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogDeploymentPatchModel.Metadata = globalCatalogDeploymentMetadataModel
				globalCatalogDeploymentPatchModelAsPatch, asPatchErr := globalCatalogDeploymentPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogDeploymentOptions model
				updateCatalogDeploymentOptionsModel := new(partnercentersellv1.UpdateCatalogDeploymentOptions)
				updateCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.GlobalCatalogDeploymentPatch = globalCatalogDeploymentPatchModelAsPatch
				updateCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateCatalogDeploymentWithContext(ctx, updateCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateCatalogDeploymentWithContext(ctx, updateCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogDeploymentPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "object_id": "ObjectID", "name": "Name", "active": true, "disabled": true, "kind": "deployment", "overview_ui": {"en": {"display_name": "DisplayName", "description": "Description", "long_description": "LongDescription"}}, "tags": ["Tags"], "url": "URL", "object_provider": {"name": "Name", "email": "Email"}, "metadata": {"rc_compatible": true, "ui": {"strings": {"en": {"bullets": [{"description": "Description", "description_i18n": {"mapKey": "Inner"}, "title": "Title", "title_i18n": {"mapKey": "Inner"}}], "media": [{"caption": "Caption", "caption_i18n": {"mapKey": "Inner"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "embeddable_dashboard": "EmbeddableDashboard"}}, "urls": {"doc_url": "DocURL", "apidocs_url": "ApidocsURL", "terms_url": "TermsURL", "instructions_url": "InstructionsURL", "catalog_details_url": "CatalogDetailsURL", "custom_create_page_url": "CustomCreatePageURL", "dashboard": "Dashboard"}, "hidden": true, "side_by_side_index": 15}, "service": {"rc_provisionable": false, "iam_compatible": false, "bindable": true, "plan_updateable": true, "service_key_supported": false}, "deployment": {"broker": {"name": "Name", "guid": "Guid"}, "location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn"}}}`)
				}))
			})
			It(`Invoke UpdateCatalogDeployment successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateCatalogDeployment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("another-broker")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3cf")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the GlobalCatalogDeploymentPatch model
				globalCatalogDeploymentPatchModel := new(partnercentersellv1.GlobalCatalogDeploymentPatch)
				globalCatalogDeploymentPatchModel.Active = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogDeploymentPatchModel.Tags = []string{"testString"}
				globalCatalogDeploymentPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogDeploymentPatchModel.Metadata = globalCatalogDeploymentMetadataModel
				globalCatalogDeploymentPatchModelAsPatch, asPatchErr := globalCatalogDeploymentPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogDeploymentOptions model
				updateCatalogDeploymentOptionsModel := new(partnercentersellv1.UpdateCatalogDeploymentOptions)
				updateCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.GlobalCatalogDeploymentPatch = globalCatalogDeploymentPatchModelAsPatch
				updateCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCatalogDeployment with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("another-broker")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3cf")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the GlobalCatalogDeploymentPatch model
				globalCatalogDeploymentPatchModel := new(partnercentersellv1.GlobalCatalogDeploymentPatch)
				globalCatalogDeploymentPatchModel.Active = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogDeploymentPatchModel.Tags = []string{"testString"}
				globalCatalogDeploymentPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogDeploymentPatchModel.Metadata = globalCatalogDeploymentMetadataModel
				globalCatalogDeploymentPatchModelAsPatch, asPatchErr := globalCatalogDeploymentPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogDeploymentOptions model
				updateCatalogDeploymentOptionsModel := new(partnercentersellv1.UpdateCatalogDeploymentOptions)
				updateCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.GlobalCatalogDeploymentPatch = globalCatalogDeploymentPatchModelAsPatch
				updateCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCatalogDeploymentOptions model with no property values
				updateCatalogDeploymentOptionsModelNew := new(partnercentersellv1.UpdateCatalogDeploymentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCatalogDeployment successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel

				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				catalogProductProviderModel.Name = core.StringPtr("testString")
				catalogProductProviderModel.Email = core.StringPtr("testString")

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("another-broker")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3cf")

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel

				// Construct an instance of the GlobalCatalogDeploymentPatch model
				globalCatalogDeploymentPatchModel := new(partnercentersellv1.GlobalCatalogDeploymentPatch)
				globalCatalogDeploymentPatchModel.Active = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.Disabled = core.BoolPtr(true)
				globalCatalogDeploymentPatchModel.OverviewUi = globalCatalogOverviewUiModel
				globalCatalogDeploymentPatchModel.Tags = []string{"testString"}
				globalCatalogDeploymentPatchModel.ObjectProvider = catalogProductProviderModel
				globalCatalogDeploymentPatchModel.Metadata = globalCatalogDeploymentMetadataModel
				globalCatalogDeploymentPatchModelAsPatch, asPatchErr := globalCatalogDeploymentPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateCatalogDeploymentOptions model
				updateCatalogDeploymentOptionsModel := new(partnercentersellv1.UpdateCatalogDeploymentOptions)
				updateCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.GlobalCatalogDeploymentPatch = globalCatalogDeploymentPatchModelAsPatch
				updateCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				updateCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCatalogDeployment(deleteCatalogDeploymentOptions *DeleteCatalogDeploymentOptions)`, func() {
		deleteCatalogDeploymentPath := "/products/testString/catalog_products/testString/catalog_plans/testString/catalog_deployments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCatalogDeploymentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCatalogDeployment successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := partnerCenterSellService.DeleteCatalogDeployment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCatalogDeploymentOptions model
				deleteCatalogDeploymentOptionsModel := new(partnercentersellv1.DeleteCatalogDeploymentOptions)
				deleteCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = partnerCenterSellService.DeleteCatalogDeployment(deleteCatalogDeploymentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCatalogDeployment with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeleteCatalogDeploymentOptions model
				deleteCatalogDeploymentOptionsModel := new(partnercentersellv1.DeleteCatalogDeploymentOptions)
				deleteCatalogDeploymentOptionsModel.ProductID = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.CatalogProductID = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.CatalogPlanID = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.CatalogDeploymentID = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.Env = core.StringPtr("testString")
				deleteCatalogDeploymentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := partnerCenterSellService.DeleteCatalogDeployment(deleteCatalogDeploymentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCatalogDeploymentOptions model with no property values
				deleteCatalogDeploymentOptionsModelNew := new(partnercentersellv1.DeleteCatalogDeploymentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = partnerCenterSellService.DeleteCatalogDeployment(deleteCatalogDeploymentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateIamRegistration(createIamRegistrationOptions *CreateIamRegistrationOptions) - Operation response error`, func() {
		createIamRegistrationPath := "/products/testString/iam_registration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createIamRegistrationPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateIamRegistration with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the CreateIamRegistrationOptions model
				createIamRegistrationOptionsModel := new(partnercentersellv1.CreateIamRegistrationOptions)
				createIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Name = core.StringPtr("pet-store")
				createIamRegistrationOptionsModel.Enabled = core.BoolPtr(true)
				createIamRegistrationOptionsModel.ServiceType = core.StringPtr("service")
				createIamRegistrationOptionsModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				createIamRegistrationOptionsModel.AdditionalPolicyScopes = []string{"pet-store"}
				createIamRegistrationOptionsModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				createIamRegistrationOptionsModel.ParentIds = []string{}
				createIamRegistrationOptionsModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				createIamRegistrationOptionsModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				createIamRegistrationOptionsModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				createIamRegistrationOptionsModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				createIamRegistrationOptionsModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				createIamRegistrationOptionsModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				createIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateIamRegistration(createIamRegistrationOptions *CreateIamRegistrationOptions)`, func() {
		createIamRegistrationPath := "/products/testString/iam_registration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createIamRegistrationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "enabled": false, "service_type": "service", "actions": [{"id": "ID", "roles": ["Roles"], "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"hidden": true}}], "additional_policy_scopes": ["AdditionalPolicyScopes"], "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "parent_ids": ["ParentIds"], "resource_hierarchy_attribute": {"key": "Key", "value": "Value"}, "supported_anonymous_accesses": [{"attributes": {"account_id": "AccountID", "service_name": "ServiceName", "additional_properties": {"mapKey": "Inner"}}, "roles": ["Roles"]}], "supported_attributes": [{"key": "Key", "options": {"operators": ["stringEquals"], "hidden": true, "supported_attributes": ["SupportedAttributes"], "policy_types": ["access"], "is_empty_value_supported": false, "is_string_exists_false_value_supported": false, "key": "Key", "resource_hierarchy": {"key": {"key": "Key", "value": "Value"}, "value": {"key": "Key"}}}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "ui": {"input_type": "InputType", "input_details": {"type": "Type", "values": [{"value": "Value", "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}}], "gst": {"query": "Query", "value_property_name": "ValuePropertyName", "label_property_name": "LabelPropertyName", "input_option_label": "InputOptionLabel"}, "url": {"url_endpoint": "UrlEndpoint", "input_option_label": "InputOptionLabel"}}}}], "supported_authorization_subjects": [{"attributes": {"service_name": "ServiceName", "resource_type": "ResourceType"}, "roles": ["Roles"]}], "supported_roles": [{"id": "ID", "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"access_policy": true, "additional_properties_for_access_policy": {"mapKey": "Inner"}, "policy_type": ["access"], "account_type": "enterprise"}}], "supported_network": {"environment_attributes": [{"key": "Key", "values": ["Values"], "options": {"hidden": true}}]}}`)
				}))
			})
			It(`Invoke CreateIamRegistration successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the CreateIamRegistrationOptions model
				createIamRegistrationOptionsModel := new(partnercentersellv1.CreateIamRegistrationOptions)
				createIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Name = core.StringPtr("pet-store")
				createIamRegistrationOptionsModel.Enabled = core.BoolPtr(true)
				createIamRegistrationOptionsModel.ServiceType = core.StringPtr("service")
				createIamRegistrationOptionsModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				createIamRegistrationOptionsModel.AdditionalPolicyScopes = []string{"pet-store"}
				createIamRegistrationOptionsModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				createIamRegistrationOptionsModel.ParentIds = []string{}
				createIamRegistrationOptionsModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				createIamRegistrationOptionsModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				createIamRegistrationOptionsModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				createIamRegistrationOptionsModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				createIamRegistrationOptionsModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				createIamRegistrationOptionsModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				createIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreateIamRegistrationWithContext(ctx, createIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreateIamRegistrationWithContext(ctx, createIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createIamRegistrationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "enabled": false, "service_type": "service", "actions": [{"id": "ID", "roles": ["Roles"], "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"hidden": true}}], "additional_policy_scopes": ["AdditionalPolicyScopes"], "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "parent_ids": ["ParentIds"], "resource_hierarchy_attribute": {"key": "Key", "value": "Value"}, "supported_anonymous_accesses": [{"attributes": {"account_id": "AccountID", "service_name": "ServiceName", "additional_properties": {"mapKey": "Inner"}}, "roles": ["Roles"]}], "supported_attributes": [{"key": "Key", "options": {"operators": ["stringEquals"], "hidden": true, "supported_attributes": ["SupportedAttributes"], "policy_types": ["access"], "is_empty_value_supported": false, "is_string_exists_false_value_supported": false, "key": "Key", "resource_hierarchy": {"key": {"key": "Key", "value": "Value"}, "value": {"key": "Key"}}}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "ui": {"input_type": "InputType", "input_details": {"type": "Type", "values": [{"value": "Value", "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}}], "gst": {"query": "Query", "value_property_name": "ValuePropertyName", "label_property_name": "LabelPropertyName", "input_option_label": "InputOptionLabel"}, "url": {"url_endpoint": "UrlEndpoint", "input_option_label": "InputOptionLabel"}}}}], "supported_authorization_subjects": [{"attributes": {"service_name": "ServiceName", "resource_type": "ResourceType"}, "roles": ["Roles"]}], "supported_roles": [{"id": "ID", "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"access_policy": true, "additional_properties_for_access_policy": {"mapKey": "Inner"}, "policy_type": ["access"], "account_type": "enterprise"}}], "supported_network": {"environment_attributes": [{"key": "Key", "values": ["Values"], "options": {"hidden": true}}]}}`)
				}))
			})
			It(`Invoke CreateIamRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreateIamRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the CreateIamRegistrationOptions model
				createIamRegistrationOptionsModel := new(partnercentersellv1.CreateIamRegistrationOptions)
				createIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Name = core.StringPtr("pet-store")
				createIamRegistrationOptionsModel.Enabled = core.BoolPtr(true)
				createIamRegistrationOptionsModel.ServiceType = core.StringPtr("service")
				createIamRegistrationOptionsModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				createIamRegistrationOptionsModel.AdditionalPolicyScopes = []string{"pet-store"}
				createIamRegistrationOptionsModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				createIamRegistrationOptionsModel.ParentIds = []string{}
				createIamRegistrationOptionsModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				createIamRegistrationOptionsModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				createIamRegistrationOptionsModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				createIamRegistrationOptionsModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				createIamRegistrationOptionsModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				createIamRegistrationOptionsModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				createIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateIamRegistration with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the CreateIamRegistrationOptions model
				createIamRegistrationOptionsModel := new(partnercentersellv1.CreateIamRegistrationOptions)
				createIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Name = core.StringPtr("pet-store")
				createIamRegistrationOptionsModel.Enabled = core.BoolPtr(true)
				createIamRegistrationOptionsModel.ServiceType = core.StringPtr("service")
				createIamRegistrationOptionsModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				createIamRegistrationOptionsModel.AdditionalPolicyScopes = []string{"pet-store"}
				createIamRegistrationOptionsModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				createIamRegistrationOptionsModel.ParentIds = []string{}
				createIamRegistrationOptionsModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				createIamRegistrationOptionsModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				createIamRegistrationOptionsModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				createIamRegistrationOptionsModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				createIamRegistrationOptionsModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				createIamRegistrationOptionsModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				createIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateIamRegistrationOptions model with no property values
				createIamRegistrationOptionsModelNew := new(partnercentersellv1.CreateIamRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateIamRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the CreateIamRegistrationOptions model
				createIamRegistrationOptionsModel := new(partnercentersellv1.CreateIamRegistrationOptions)
				createIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Name = core.StringPtr("pet-store")
				createIamRegistrationOptionsModel.Enabled = core.BoolPtr(true)
				createIamRegistrationOptionsModel.ServiceType = core.StringPtr("service")
				createIamRegistrationOptionsModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				createIamRegistrationOptionsModel.AdditionalPolicyScopes = []string{"pet-store"}
				createIamRegistrationOptionsModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				createIamRegistrationOptionsModel.ParentIds = []string{}
				createIamRegistrationOptionsModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				createIamRegistrationOptionsModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				createIamRegistrationOptionsModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				createIamRegistrationOptionsModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				createIamRegistrationOptionsModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				createIamRegistrationOptionsModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				createIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				createIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateIamRegistration(updateIamRegistrationOptions *UpdateIamRegistrationOptions) - Operation response error`, func() {
		updateIamRegistrationPath := "/products/testString/iam_registration/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIamRegistrationPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateIamRegistration with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the IamServiceRegistrationPatch model
				iamServiceRegistrationPatchModel := new(partnercentersellv1.IamServiceRegistrationPatch)
				iamServiceRegistrationPatchModel.Enabled = core.BoolPtr(true)
				iamServiceRegistrationPatchModel.ServiceType = core.StringPtr("service")
				iamServiceRegistrationPatchModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				iamServiceRegistrationPatchModel.AdditionalPolicyScopes = []string{"pet-store"}
				iamServiceRegistrationPatchModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationPatchModel.ParentIds = []string{}
				iamServiceRegistrationPatchModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				iamServiceRegistrationPatchModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				iamServiceRegistrationPatchModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				iamServiceRegistrationPatchModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				iamServiceRegistrationPatchModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				iamServiceRegistrationPatchModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				iamServiceRegistrationPatchModelAsPatch, asPatchErr := iamServiceRegistrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateIamRegistrationOptions model
				updateIamRegistrationOptionsModel := new(partnercentersellv1.UpdateIamRegistrationOptions)
				updateIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.IamRegistrationPatch = iamServiceRegistrationPatchModelAsPatch
				updateIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateIamRegistration(updateIamRegistrationOptions *UpdateIamRegistrationOptions)`, func() {
		updateIamRegistrationPath := "/products/testString/iam_registration/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIamRegistrationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "enabled": false, "service_type": "service", "actions": [{"id": "ID", "roles": ["Roles"], "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"hidden": true}}], "additional_policy_scopes": ["AdditionalPolicyScopes"], "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "parent_ids": ["ParentIds"], "resource_hierarchy_attribute": {"key": "Key", "value": "Value"}, "supported_anonymous_accesses": [{"attributes": {"account_id": "AccountID", "service_name": "ServiceName", "additional_properties": {"mapKey": "Inner"}}, "roles": ["Roles"]}], "supported_attributes": [{"key": "Key", "options": {"operators": ["stringEquals"], "hidden": true, "supported_attributes": ["SupportedAttributes"], "policy_types": ["access"], "is_empty_value_supported": false, "is_string_exists_false_value_supported": false, "key": "Key", "resource_hierarchy": {"key": {"key": "Key", "value": "Value"}, "value": {"key": "Key"}}}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "ui": {"input_type": "InputType", "input_details": {"type": "Type", "values": [{"value": "Value", "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}}], "gst": {"query": "Query", "value_property_name": "ValuePropertyName", "label_property_name": "LabelPropertyName", "input_option_label": "InputOptionLabel"}, "url": {"url_endpoint": "UrlEndpoint", "input_option_label": "InputOptionLabel"}}}}], "supported_authorization_subjects": [{"attributes": {"service_name": "ServiceName", "resource_type": "ResourceType"}, "roles": ["Roles"]}], "supported_roles": [{"id": "ID", "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"access_policy": true, "additional_properties_for_access_policy": {"mapKey": "Inner"}, "policy_type": ["access"], "account_type": "enterprise"}}], "supported_network": {"environment_attributes": [{"key": "Key", "values": ["Values"], "options": {"hidden": true}}]}}`)
				}))
			})
			It(`Invoke UpdateIamRegistration successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the IamServiceRegistrationPatch model
				iamServiceRegistrationPatchModel := new(partnercentersellv1.IamServiceRegistrationPatch)
				iamServiceRegistrationPatchModel.Enabled = core.BoolPtr(true)
				iamServiceRegistrationPatchModel.ServiceType = core.StringPtr("service")
				iamServiceRegistrationPatchModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				iamServiceRegistrationPatchModel.AdditionalPolicyScopes = []string{"pet-store"}
				iamServiceRegistrationPatchModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationPatchModel.ParentIds = []string{}
				iamServiceRegistrationPatchModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				iamServiceRegistrationPatchModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				iamServiceRegistrationPatchModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				iamServiceRegistrationPatchModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				iamServiceRegistrationPatchModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				iamServiceRegistrationPatchModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				iamServiceRegistrationPatchModelAsPatch, asPatchErr := iamServiceRegistrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateIamRegistrationOptions model
				updateIamRegistrationOptionsModel := new(partnercentersellv1.UpdateIamRegistrationOptions)
				updateIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.IamRegistrationPatch = iamServiceRegistrationPatchModelAsPatch
				updateIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateIamRegistrationWithContext(ctx, updateIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateIamRegistrationWithContext(ctx, updateIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIamRegistrationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "enabled": false, "service_type": "service", "actions": [{"id": "ID", "roles": ["Roles"], "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"hidden": true}}], "additional_policy_scopes": ["AdditionalPolicyScopes"], "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "parent_ids": ["ParentIds"], "resource_hierarchy_attribute": {"key": "Key", "value": "Value"}, "supported_anonymous_accesses": [{"attributes": {"account_id": "AccountID", "service_name": "ServiceName", "additional_properties": {"mapKey": "Inner"}}, "roles": ["Roles"]}], "supported_attributes": [{"key": "Key", "options": {"operators": ["stringEquals"], "hidden": true, "supported_attributes": ["SupportedAttributes"], "policy_types": ["access"], "is_empty_value_supported": false, "is_string_exists_false_value_supported": false, "key": "Key", "resource_hierarchy": {"key": {"key": "Key", "value": "Value"}, "value": {"key": "Key"}}}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "ui": {"input_type": "InputType", "input_details": {"type": "Type", "values": [{"value": "Value", "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}}], "gst": {"query": "Query", "value_property_name": "ValuePropertyName", "label_property_name": "LabelPropertyName", "input_option_label": "InputOptionLabel"}, "url": {"url_endpoint": "UrlEndpoint", "input_option_label": "InputOptionLabel"}}}}], "supported_authorization_subjects": [{"attributes": {"service_name": "ServiceName", "resource_type": "ResourceType"}, "roles": ["Roles"]}], "supported_roles": [{"id": "ID", "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"access_policy": true, "additional_properties_for_access_policy": {"mapKey": "Inner"}, "policy_type": ["access"], "account_type": "enterprise"}}], "supported_network": {"environment_attributes": [{"key": "Key", "values": ["Values"], "options": {"hidden": true}}]}}`)
				}))
			})
			It(`Invoke UpdateIamRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateIamRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the IamServiceRegistrationPatch model
				iamServiceRegistrationPatchModel := new(partnercentersellv1.IamServiceRegistrationPatch)
				iamServiceRegistrationPatchModel.Enabled = core.BoolPtr(true)
				iamServiceRegistrationPatchModel.ServiceType = core.StringPtr("service")
				iamServiceRegistrationPatchModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				iamServiceRegistrationPatchModel.AdditionalPolicyScopes = []string{"pet-store"}
				iamServiceRegistrationPatchModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationPatchModel.ParentIds = []string{}
				iamServiceRegistrationPatchModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				iamServiceRegistrationPatchModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				iamServiceRegistrationPatchModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				iamServiceRegistrationPatchModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				iamServiceRegistrationPatchModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				iamServiceRegistrationPatchModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				iamServiceRegistrationPatchModelAsPatch, asPatchErr := iamServiceRegistrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateIamRegistrationOptions model
				updateIamRegistrationOptionsModel := new(partnercentersellv1.UpdateIamRegistrationOptions)
				updateIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.IamRegistrationPatch = iamServiceRegistrationPatchModelAsPatch
				updateIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateIamRegistration with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the IamServiceRegistrationPatch model
				iamServiceRegistrationPatchModel := new(partnercentersellv1.IamServiceRegistrationPatch)
				iamServiceRegistrationPatchModel.Enabled = core.BoolPtr(true)
				iamServiceRegistrationPatchModel.ServiceType = core.StringPtr("service")
				iamServiceRegistrationPatchModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				iamServiceRegistrationPatchModel.AdditionalPolicyScopes = []string{"pet-store"}
				iamServiceRegistrationPatchModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationPatchModel.ParentIds = []string{}
				iamServiceRegistrationPatchModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				iamServiceRegistrationPatchModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				iamServiceRegistrationPatchModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				iamServiceRegistrationPatchModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				iamServiceRegistrationPatchModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				iamServiceRegistrationPatchModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				iamServiceRegistrationPatchModelAsPatch, asPatchErr := iamServiceRegistrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateIamRegistrationOptions model
				updateIamRegistrationOptionsModel := new(partnercentersellv1.UpdateIamRegistrationOptions)
				updateIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.IamRegistrationPatch = iamServiceRegistrationPatchModelAsPatch
				updateIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateIamRegistrationOptions model with no property values
				updateIamRegistrationOptionsModelNew := new(partnercentersellv1.UpdateIamRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateIamRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

				// Construct an instance of the IamServiceRegistrationPatch model
				iamServiceRegistrationPatchModel := new(partnercentersellv1.IamServiceRegistrationPatch)
				iamServiceRegistrationPatchModel.Enabled = core.BoolPtr(true)
				iamServiceRegistrationPatchModel.ServiceType = core.StringPtr("service")
				iamServiceRegistrationPatchModel.Actions = []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}
				iamServiceRegistrationPatchModel.AdditionalPolicyScopes = []string{"pet-store"}
				iamServiceRegistrationPatchModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationPatchModel.ParentIds = []string{}
				iamServiceRegistrationPatchModel.ResourceHierarchyAttribute = iamServiceRegistrationResourceHierarchyAttributeModel
				iamServiceRegistrationPatchModel.SupportedAnonymousAccesses = []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}
				iamServiceRegistrationPatchModel.SupportedAttributes = []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}
				iamServiceRegistrationPatchModel.SupportedAuthorizationSubjects = []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}
				iamServiceRegistrationPatchModel.SupportedRoles = []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}
				iamServiceRegistrationPatchModel.SupportedNetwork = iamServiceRegistrationSupportedNetworkModel
				iamServiceRegistrationPatchModelAsPatch, asPatchErr := iamServiceRegistrationPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateIamRegistrationOptions model
				updateIamRegistrationOptionsModel := new(partnercentersellv1.UpdateIamRegistrationOptions)
				updateIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.IamRegistrationPatch = iamServiceRegistrationPatchModelAsPatch
				updateIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				updateIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteIamRegistration(deleteIamRegistrationOptions *DeleteIamRegistrationOptions)`, func() {
		deleteIamRegistrationPath := "/products/testString/iam_registration/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteIamRegistrationPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteIamRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := partnerCenterSellService.DeleteIamRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteIamRegistrationOptions model
				deleteIamRegistrationOptionsModel := new(partnercentersellv1.DeleteIamRegistrationOptions)
				deleteIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				deleteIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				deleteIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				deleteIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = partnerCenterSellService.DeleteIamRegistration(deleteIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteIamRegistration with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeleteIamRegistrationOptions model
				deleteIamRegistrationOptionsModel := new(partnercentersellv1.DeleteIamRegistrationOptions)
				deleteIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				deleteIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				deleteIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				deleteIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := partnerCenterSellService.DeleteIamRegistration(deleteIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteIamRegistrationOptions model with no property values
				deleteIamRegistrationOptionsModelNew := new(partnercentersellv1.DeleteIamRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = partnerCenterSellService.DeleteIamRegistration(deleteIamRegistrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIamRegistration(getIamRegistrationOptions *GetIamRegistrationOptions) - Operation response error`, func() {
		getIamRegistrationPath := "/products/testString/iam_registration/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIamRegistrationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetIamRegistration with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetIamRegistrationOptions model
				getIamRegistrationOptionsModel := new(partnercentersellv1.GetIamRegistrationOptions)
				getIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				getIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetIamRegistration(getIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetIamRegistration(getIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIamRegistration(getIamRegistrationOptions *GetIamRegistrationOptions)`, func() {
		getIamRegistrationPath := "/products/testString/iam_registration/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIamRegistrationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "enabled": false, "service_type": "service", "actions": [{"id": "ID", "roles": ["Roles"], "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"hidden": true}}], "additional_policy_scopes": ["AdditionalPolicyScopes"], "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "parent_ids": ["ParentIds"], "resource_hierarchy_attribute": {"key": "Key", "value": "Value"}, "supported_anonymous_accesses": [{"attributes": {"account_id": "AccountID", "service_name": "ServiceName", "additional_properties": {"mapKey": "Inner"}}, "roles": ["Roles"]}], "supported_attributes": [{"key": "Key", "options": {"operators": ["stringEquals"], "hidden": true, "supported_attributes": ["SupportedAttributes"], "policy_types": ["access"], "is_empty_value_supported": false, "is_string_exists_false_value_supported": false, "key": "Key", "resource_hierarchy": {"key": {"key": "Key", "value": "Value"}, "value": {"key": "Key"}}}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "ui": {"input_type": "InputType", "input_details": {"type": "Type", "values": [{"value": "Value", "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}}], "gst": {"query": "Query", "value_property_name": "ValuePropertyName", "label_property_name": "LabelPropertyName", "input_option_label": "InputOptionLabel"}, "url": {"url_endpoint": "UrlEndpoint", "input_option_label": "InputOptionLabel"}}}}], "supported_authorization_subjects": [{"attributes": {"service_name": "ServiceName", "resource_type": "ResourceType"}, "roles": ["Roles"]}], "supported_roles": [{"id": "ID", "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"access_policy": true, "additional_properties_for_access_policy": {"mapKey": "Inner"}, "policy_type": ["access"], "account_type": "enterprise"}}], "supported_network": {"environment_attributes": [{"key": "Key", "values": ["Values"], "options": {"hidden": true}}]}}`)
				}))
			})
			It(`Invoke GetIamRegistration successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetIamRegistrationOptions model
				getIamRegistrationOptionsModel := new(partnercentersellv1.GetIamRegistrationOptions)
				getIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				getIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetIamRegistrationWithContext(ctx, getIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetIamRegistration(getIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetIamRegistrationWithContext(ctx, getIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIamRegistrationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "enabled": false, "service_type": "service", "actions": [{"id": "ID", "roles": ["Roles"], "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"hidden": true}}], "additional_policy_scopes": ["AdditionalPolicyScopes"], "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "parent_ids": ["ParentIds"], "resource_hierarchy_attribute": {"key": "Key", "value": "Value"}, "supported_anonymous_accesses": [{"attributes": {"account_id": "AccountID", "service_name": "ServiceName", "additional_properties": {"mapKey": "Inner"}}, "roles": ["Roles"]}], "supported_attributes": [{"key": "Key", "options": {"operators": ["stringEquals"], "hidden": true, "supported_attributes": ["SupportedAttributes"], "policy_types": ["access"], "is_empty_value_supported": false, "is_string_exists_false_value_supported": false, "key": "Key", "resource_hierarchy": {"key": {"key": "Key", "value": "Value"}, "value": {"key": "Key"}}}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "ui": {"input_type": "InputType", "input_details": {"type": "Type", "values": [{"value": "Value", "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}}], "gst": {"query": "Query", "value_property_name": "ValuePropertyName", "label_property_name": "LabelPropertyName", "input_option_label": "InputOptionLabel"}, "url": {"url_endpoint": "UrlEndpoint", "input_option_label": "InputOptionLabel"}}}}], "supported_authorization_subjects": [{"attributes": {"service_name": "ServiceName", "resource_type": "ResourceType"}, "roles": ["Roles"]}], "supported_roles": [{"id": "ID", "description": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "display_name": {"default": "Default", "en": "En", "de": "De", "es": "Es", "fr": "Fr", "it": "It", "ja": "Ja", "ko": "Ko", "pt_br": "PtBr", "zh_tw": "ZhTw", "zh_cn": "ZhCn"}, "options": {"access_policy": true, "additional_properties_for_access_policy": {"mapKey": "Inner"}, "policy_type": ["access"], "account_type": "enterprise"}}], "supported_network": {"environment_attributes": [{"key": "Key", "values": ["Values"], "options": {"hidden": true}}]}}`)
				}))
			})
			It(`Invoke GetIamRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetIamRegistration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIamRegistrationOptions model
				getIamRegistrationOptionsModel := new(partnercentersellv1.GetIamRegistrationOptions)
				getIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				getIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetIamRegistration(getIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetIamRegistration with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetIamRegistrationOptions model
				getIamRegistrationOptionsModel := new(partnercentersellv1.GetIamRegistrationOptions)
				getIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				getIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetIamRegistration(getIamRegistrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetIamRegistrationOptions model with no property values
				getIamRegistrationOptionsModelNew := new(partnercentersellv1.GetIamRegistrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetIamRegistration(getIamRegistrationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetIamRegistration successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetIamRegistrationOptions model
				getIamRegistrationOptionsModel := new(partnercentersellv1.GetIamRegistrationOptions)
				getIamRegistrationOptionsModel.ProductID = core.StringPtr("testString")
				getIamRegistrationOptionsModel.ProgrammaticName = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Env = core.StringPtr("testString")
				getIamRegistrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetIamRegistration(getIamRegistrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResourceBroker(createResourceBrokerOptions *CreateResourceBrokerOptions) - Operation response error`, func() {
		createResourceBrokerPath := "/brokers"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceBrokerPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceBroker with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CreateResourceBrokerOptions model
				createResourceBrokerOptionsModel := new(partnercentersellv1.CreateResourceBrokerOptions)
				createResourceBrokerOptionsModel.AuthScheme = core.StringPtr("bearer")
				createResourceBrokerOptionsModel.Name = core.StringPtr("brokername")
				createResourceBrokerOptionsModel.BrokerURL = core.StringPtr("https://broker-url-for-my-service.com")
				createResourceBrokerOptionsModel.Type = core.StringPtr("provision_through")
				createResourceBrokerOptionsModel.AuthUsername = core.StringPtr("apikey")
				createResourceBrokerOptionsModel.AuthPassword = core.StringPtr("testString")
				createResourceBrokerOptionsModel.ResourceGroupCrn = core.StringPtr("crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f")
				createResourceBrokerOptionsModel.State = core.StringPtr("active")
				createResourceBrokerOptionsModel.AllowContextUpdates = core.BoolPtr(false)
				createResourceBrokerOptionsModel.CatalogType = core.StringPtr("service")
				createResourceBrokerOptionsModel.Region = core.StringPtr("global")
				createResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				createResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResourceBroker(createResourceBrokerOptions *CreateResourceBrokerOptions)`, func() {
		createResourceBrokerPath := "/brokers"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceBrokerPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"auth_username": "apikey", "auth_password": "AuthPassword", "auth_scheme": "bearer", "resource_group_crn": "ResourceGroupCrn", "state": "active", "broker_url": "BrokerURL", "allow_context_updates": false, "catalog_type": "CatalogType", "type": "provision_through", "name": "Name", "region": "Region", "account_id": "AccountID", "crn": "Crn", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": {"user_id": "UserID", "user_name": "UserName"}, "updated_by": {"user_id": "UserID", "user_name": "UserName"}, "deleted_by": {"user_id": "UserID", "user_name": "UserName"}, "guid": "Guid", "id": "ID", "url": "URL"}`)
				}))
			})
			It(`Invoke CreateResourceBroker successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the CreateResourceBrokerOptions model
				createResourceBrokerOptionsModel := new(partnercentersellv1.CreateResourceBrokerOptions)
				createResourceBrokerOptionsModel.AuthScheme = core.StringPtr("bearer")
				createResourceBrokerOptionsModel.Name = core.StringPtr("brokername")
				createResourceBrokerOptionsModel.BrokerURL = core.StringPtr("https://broker-url-for-my-service.com")
				createResourceBrokerOptionsModel.Type = core.StringPtr("provision_through")
				createResourceBrokerOptionsModel.AuthUsername = core.StringPtr("apikey")
				createResourceBrokerOptionsModel.AuthPassword = core.StringPtr("testString")
				createResourceBrokerOptionsModel.ResourceGroupCrn = core.StringPtr("crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f")
				createResourceBrokerOptionsModel.State = core.StringPtr("active")
				createResourceBrokerOptionsModel.AllowContextUpdates = core.BoolPtr(false)
				createResourceBrokerOptionsModel.CatalogType = core.StringPtr("service")
				createResourceBrokerOptionsModel.Region = core.StringPtr("global")
				createResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				createResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreateResourceBrokerWithContext(ctx, createResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreateResourceBrokerWithContext(ctx, createResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceBrokerPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"auth_username": "apikey", "auth_password": "AuthPassword", "auth_scheme": "bearer", "resource_group_crn": "ResourceGroupCrn", "state": "active", "broker_url": "BrokerURL", "allow_context_updates": false, "catalog_type": "CatalogType", "type": "provision_through", "name": "Name", "region": "Region", "account_id": "AccountID", "crn": "Crn", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": {"user_id": "UserID", "user_name": "UserName"}, "updated_by": {"user_id": "UserID", "user_name": "UserName"}, "deleted_by": {"user_id": "UserID", "user_name": "UserName"}, "guid": "Guid", "id": "ID", "url": "URL"}`)
				}))
			})
			It(`Invoke CreateResourceBroker successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreateResourceBroker(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateResourceBrokerOptions model
				createResourceBrokerOptionsModel := new(partnercentersellv1.CreateResourceBrokerOptions)
				createResourceBrokerOptionsModel.AuthScheme = core.StringPtr("bearer")
				createResourceBrokerOptionsModel.Name = core.StringPtr("brokername")
				createResourceBrokerOptionsModel.BrokerURL = core.StringPtr("https://broker-url-for-my-service.com")
				createResourceBrokerOptionsModel.Type = core.StringPtr("provision_through")
				createResourceBrokerOptionsModel.AuthUsername = core.StringPtr("apikey")
				createResourceBrokerOptionsModel.AuthPassword = core.StringPtr("testString")
				createResourceBrokerOptionsModel.ResourceGroupCrn = core.StringPtr("crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f")
				createResourceBrokerOptionsModel.State = core.StringPtr("active")
				createResourceBrokerOptionsModel.AllowContextUpdates = core.BoolPtr(false)
				createResourceBrokerOptionsModel.CatalogType = core.StringPtr("service")
				createResourceBrokerOptionsModel.Region = core.StringPtr("global")
				createResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				createResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateResourceBroker with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CreateResourceBrokerOptions model
				createResourceBrokerOptionsModel := new(partnercentersellv1.CreateResourceBrokerOptions)
				createResourceBrokerOptionsModel.AuthScheme = core.StringPtr("bearer")
				createResourceBrokerOptionsModel.Name = core.StringPtr("brokername")
				createResourceBrokerOptionsModel.BrokerURL = core.StringPtr("https://broker-url-for-my-service.com")
				createResourceBrokerOptionsModel.Type = core.StringPtr("provision_through")
				createResourceBrokerOptionsModel.AuthUsername = core.StringPtr("apikey")
				createResourceBrokerOptionsModel.AuthPassword = core.StringPtr("testString")
				createResourceBrokerOptionsModel.ResourceGroupCrn = core.StringPtr("crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f")
				createResourceBrokerOptionsModel.State = core.StringPtr("active")
				createResourceBrokerOptionsModel.AllowContextUpdates = core.BoolPtr(false)
				createResourceBrokerOptionsModel.CatalogType = core.StringPtr("service")
				createResourceBrokerOptionsModel.Region = core.StringPtr("global")
				createResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				createResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceBrokerOptions model with no property values
				createResourceBrokerOptionsModelNew := new(partnercentersellv1.CreateResourceBrokerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateResourceBroker successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CreateResourceBrokerOptions model
				createResourceBrokerOptionsModel := new(partnercentersellv1.CreateResourceBrokerOptions)
				createResourceBrokerOptionsModel.AuthScheme = core.StringPtr("bearer")
				createResourceBrokerOptionsModel.Name = core.StringPtr("brokername")
				createResourceBrokerOptionsModel.BrokerURL = core.StringPtr("https://broker-url-for-my-service.com")
				createResourceBrokerOptionsModel.Type = core.StringPtr("provision_through")
				createResourceBrokerOptionsModel.AuthUsername = core.StringPtr("apikey")
				createResourceBrokerOptionsModel.AuthPassword = core.StringPtr("testString")
				createResourceBrokerOptionsModel.ResourceGroupCrn = core.StringPtr("crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f")
				createResourceBrokerOptionsModel.State = core.StringPtr("active")
				createResourceBrokerOptionsModel.AllowContextUpdates = core.BoolPtr(false)
				createResourceBrokerOptionsModel.CatalogType = core.StringPtr("service")
				createResourceBrokerOptionsModel.Region = core.StringPtr("global")
				createResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				createResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceBroker(updateResourceBrokerOptions *UpdateResourceBrokerOptions) - Operation response error`, func() {
		updateResourceBrokerPath := "/brokers/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceBrokerPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceBroker with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the BrokerPatch model
				brokerPatchModel := new(partnercentersellv1.BrokerPatch)
				brokerPatchModel.AuthUsername = core.StringPtr("apikey")
				brokerPatchModel.AuthPassword = core.StringPtr("testString")
				brokerPatchModel.AuthScheme = core.StringPtr("bearer")
				brokerPatchModel.ResourceGroupCrn = core.StringPtr("testString")
				brokerPatchModel.State = core.StringPtr("active")
				brokerPatchModel.BrokerURL = core.StringPtr("https://my-updated-broker-url.com")
				brokerPatchModel.AllowContextUpdates = core.BoolPtr(true)
				brokerPatchModel.CatalogType = core.StringPtr("testString")
				brokerPatchModel.Type = core.StringPtr("provision_through")
				brokerPatchModel.Region = core.StringPtr("testString")
				brokerPatchModelAsPatch, asPatchErr := brokerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateResourceBrokerOptions model
				updateResourceBrokerOptionsModel := new(partnercentersellv1.UpdateResourceBrokerOptions)
				updateResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.BrokerPatch = brokerPatchModelAsPatch
				updateResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceBroker(updateResourceBrokerOptions *UpdateResourceBrokerOptions)`, func() {
		updateResourceBrokerPath := "/brokers/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceBrokerPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"auth_username": "apikey", "auth_password": "AuthPassword", "auth_scheme": "bearer", "resource_group_crn": "ResourceGroupCrn", "state": "active", "broker_url": "BrokerURL", "allow_context_updates": false, "catalog_type": "CatalogType", "type": "provision_through", "name": "Name", "region": "Region", "account_id": "AccountID", "crn": "Crn", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": {"user_id": "UserID", "user_name": "UserName"}, "updated_by": {"user_id": "UserID", "user_name": "UserName"}, "deleted_by": {"user_id": "UserID", "user_name": "UserName"}, "guid": "Guid", "id": "ID", "url": "URL"}`)
				}))
			})
			It(`Invoke UpdateResourceBroker successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the BrokerPatch model
				brokerPatchModel := new(partnercentersellv1.BrokerPatch)
				brokerPatchModel.AuthUsername = core.StringPtr("apikey")
				brokerPatchModel.AuthPassword = core.StringPtr("testString")
				brokerPatchModel.AuthScheme = core.StringPtr("bearer")
				brokerPatchModel.ResourceGroupCrn = core.StringPtr("testString")
				brokerPatchModel.State = core.StringPtr("active")
				brokerPatchModel.BrokerURL = core.StringPtr("https://my-updated-broker-url.com")
				brokerPatchModel.AllowContextUpdates = core.BoolPtr(true)
				brokerPatchModel.CatalogType = core.StringPtr("testString")
				brokerPatchModel.Type = core.StringPtr("provision_through")
				brokerPatchModel.Region = core.StringPtr("testString")
				brokerPatchModelAsPatch, asPatchErr := brokerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateResourceBrokerOptions model
				updateResourceBrokerOptionsModel := new(partnercentersellv1.UpdateResourceBrokerOptions)
				updateResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.BrokerPatch = brokerPatchModelAsPatch
				updateResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateResourceBrokerWithContext(ctx, updateResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateResourceBrokerWithContext(ctx, updateResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceBrokerPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"auth_username": "apikey", "auth_password": "AuthPassword", "auth_scheme": "bearer", "resource_group_crn": "ResourceGroupCrn", "state": "active", "broker_url": "BrokerURL", "allow_context_updates": false, "catalog_type": "CatalogType", "type": "provision_through", "name": "Name", "region": "Region", "account_id": "AccountID", "crn": "Crn", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": {"user_id": "UserID", "user_name": "UserName"}, "updated_by": {"user_id": "UserID", "user_name": "UserName"}, "deleted_by": {"user_id": "UserID", "user_name": "UserName"}, "guid": "Guid", "id": "ID", "url": "URL"}`)
				}))
			})
			It(`Invoke UpdateResourceBroker successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateResourceBroker(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BrokerPatch model
				brokerPatchModel := new(partnercentersellv1.BrokerPatch)
				brokerPatchModel.AuthUsername = core.StringPtr("apikey")
				brokerPatchModel.AuthPassword = core.StringPtr("testString")
				brokerPatchModel.AuthScheme = core.StringPtr("bearer")
				brokerPatchModel.ResourceGroupCrn = core.StringPtr("testString")
				brokerPatchModel.State = core.StringPtr("active")
				brokerPatchModel.BrokerURL = core.StringPtr("https://my-updated-broker-url.com")
				brokerPatchModel.AllowContextUpdates = core.BoolPtr(true)
				brokerPatchModel.CatalogType = core.StringPtr("testString")
				brokerPatchModel.Type = core.StringPtr("provision_through")
				brokerPatchModel.Region = core.StringPtr("testString")
				brokerPatchModelAsPatch, asPatchErr := brokerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateResourceBrokerOptions model
				updateResourceBrokerOptionsModel := new(partnercentersellv1.UpdateResourceBrokerOptions)
				updateResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.BrokerPatch = brokerPatchModelAsPatch
				updateResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateResourceBroker with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the BrokerPatch model
				brokerPatchModel := new(partnercentersellv1.BrokerPatch)
				brokerPatchModel.AuthUsername = core.StringPtr("apikey")
				brokerPatchModel.AuthPassword = core.StringPtr("testString")
				brokerPatchModel.AuthScheme = core.StringPtr("bearer")
				brokerPatchModel.ResourceGroupCrn = core.StringPtr("testString")
				brokerPatchModel.State = core.StringPtr("active")
				brokerPatchModel.BrokerURL = core.StringPtr("https://my-updated-broker-url.com")
				brokerPatchModel.AllowContextUpdates = core.BoolPtr(true)
				brokerPatchModel.CatalogType = core.StringPtr("testString")
				brokerPatchModel.Type = core.StringPtr("provision_through")
				brokerPatchModel.Region = core.StringPtr("testString")
				brokerPatchModelAsPatch, asPatchErr := brokerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateResourceBrokerOptions model
				updateResourceBrokerOptionsModel := new(partnercentersellv1.UpdateResourceBrokerOptions)
				updateResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.BrokerPatch = brokerPatchModelAsPatch
				updateResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceBrokerOptions model with no property values
				updateResourceBrokerOptionsModelNew := new(partnercentersellv1.UpdateResourceBrokerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateResourceBroker successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the BrokerPatch model
				brokerPatchModel := new(partnercentersellv1.BrokerPatch)
				brokerPatchModel.AuthUsername = core.StringPtr("apikey")
				brokerPatchModel.AuthPassword = core.StringPtr("testString")
				brokerPatchModel.AuthScheme = core.StringPtr("bearer")
				brokerPatchModel.ResourceGroupCrn = core.StringPtr("testString")
				brokerPatchModel.State = core.StringPtr("active")
				brokerPatchModel.BrokerURL = core.StringPtr("https://my-updated-broker-url.com")
				brokerPatchModel.AllowContextUpdates = core.BoolPtr(true)
				brokerPatchModel.CatalogType = core.StringPtr("testString")
				brokerPatchModel.Type = core.StringPtr("provision_through")
				brokerPatchModel.Region = core.StringPtr("testString")
				brokerPatchModelAsPatch, asPatchErr := brokerPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateResourceBrokerOptions model
				updateResourceBrokerOptionsModel := new(partnercentersellv1.UpdateResourceBrokerOptions)
				updateResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.BrokerPatch = brokerPatchModelAsPatch
				updateResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				updateResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceBroker(getResourceBrokerOptions *GetResourceBrokerOptions) - Operation response error`, func() {
		getResourceBrokerPath := "/brokers/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceBrokerPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceBroker with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetResourceBrokerOptions model
				getResourceBrokerOptionsModel := new(partnercentersellv1.GetResourceBrokerOptions)
				getResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetResourceBroker(getResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetResourceBroker(getResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceBroker(getResourceBrokerOptions *GetResourceBrokerOptions)`, func() {
		getResourceBrokerPath := "/brokers/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceBrokerPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"auth_username": "apikey", "auth_password": "AuthPassword", "auth_scheme": "bearer", "resource_group_crn": "ResourceGroupCrn", "state": "active", "broker_url": "BrokerURL", "allow_context_updates": false, "catalog_type": "CatalogType", "type": "provision_through", "name": "Name", "region": "Region", "account_id": "AccountID", "crn": "Crn", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": {"user_id": "UserID", "user_name": "UserName"}, "updated_by": {"user_id": "UserID", "user_name": "UserName"}, "deleted_by": {"user_id": "UserID", "user_name": "UserName"}, "guid": "Guid", "id": "ID", "url": "URL"}`)
				}))
			})
			It(`Invoke GetResourceBroker successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceBrokerOptions model
				getResourceBrokerOptionsModel := new(partnercentersellv1.GetResourceBrokerOptions)
				getResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetResourceBrokerWithContext(ctx, getResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetResourceBroker(getResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetResourceBrokerWithContext(ctx, getResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceBrokerPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"auth_username": "apikey", "auth_password": "AuthPassword", "auth_scheme": "bearer", "resource_group_crn": "ResourceGroupCrn", "state": "active", "broker_url": "BrokerURL", "allow_context_updates": false, "catalog_type": "CatalogType", "type": "provision_through", "name": "Name", "region": "Region", "account_id": "AccountID", "crn": "Crn", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": {"user_id": "UserID", "user_name": "UserName"}, "updated_by": {"user_id": "UserID", "user_name": "UserName"}, "deleted_by": {"user_id": "UserID", "user_name": "UserName"}, "guid": "Guid", "id": "ID", "url": "URL"}`)
				}))
			})
			It(`Invoke GetResourceBroker successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetResourceBroker(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceBrokerOptions model
				getResourceBrokerOptionsModel := new(partnercentersellv1.GetResourceBrokerOptions)
				getResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetResourceBroker(getResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceBroker with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetResourceBrokerOptions model
				getResourceBrokerOptionsModel := new(partnercentersellv1.GetResourceBrokerOptions)
				getResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetResourceBroker(getResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceBrokerOptions model with no property values
				getResourceBrokerOptionsModelNew := new(partnercentersellv1.GetResourceBrokerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetResourceBroker(getResourceBrokerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetResourceBroker successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetResourceBrokerOptions model
				getResourceBrokerOptionsModel := new(partnercentersellv1.GetResourceBrokerOptions)
				getResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				getResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetResourceBroker(getResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteResourceBroker(deleteResourceBrokerOptions *DeleteResourceBrokerOptions)`, func() {
		deleteResourceBrokerPath := "/brokers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteResourceBrokerPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["env"]).To(Equal([]string{"testString"}))
					// TODO: Add check for remove_from_account query parameter
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceBroker successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := partnerCenterSellService.DeleteResourceBroker(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceBrokerOptions model
				deleteResourceBrokerOptionsModel := new(partnercentersellv1.DeleteResourceBrokerOptions)
				deleteResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				deleteResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				deleteResourceBrokerOptionsModel.RemoveFromAccount = core.BoolPtr(true)
				deleteResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = partnerCenterSellService.DeleteResourceBroker(deleteResourceBrokerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceBroker with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceBrokerOptions model
				deleteResourceBrokerOptionsModel := new(partnercentersellv1.DeleteResourceBrokerOptions)
				deleteResourceBrokerOptionsModel.BrokerID = core.StringPtr("testString")
				deleteResourceBrokerOptionsModel.Env = core.StringPtr("testString")
				deleteResourceBrokerOptionsModel.RemoveFromAccount = core.BoolPtr(true)
				deleteResourceBrokerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := partnerCenterSellService.DeleteResourceBroker(deleteResourceBrokerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceBrokerOptions model with no property values
				deleteResourceBrokerOptionsModelNew := new(partnercentersellv1.DeleteResourceBrokerOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = partnerCenterSellService.DeleteResourceBroker(deleteResourceBrokerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProductBadges(listProductBadgesOptions *ListProductBadgesOptions) - Operation response error`, func() {
		listProductBadgesPath := "/product_badges"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProductBadgesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// TODO: Add check for start query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProductBadges with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListProductBadgesOptions model
				listProductBadgesOptionsModel := new(partnercentersellv1.ListProductBadgesOptions)
				listProductBadgesOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProductBadgesOptionsModel.Start = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.ListProductBadges(listProductBadgesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.ListProductBadges(listProductBadgesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProductBadges(listProductBadgesOptions *ListProductBadgesOptions)`, func() {
		listProductBadgesPath := "/product_badges"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProductBadgesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// TODO: Add check for start query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 1, "offset": 0, "total_count": 0, "first": {"href": "Href", "start": "9fab83da-98cb-4f18-a7ba-b6f0435c9673"}, "next": {"href": "Href", "start": "9fab83da-98cb-4f18-a7ba-b6f0435c9673"}, "previous": {"href": "Href", "start": "9fab83da-98cb-4f18-a7ba-b6f0435c9673"}, "last": {"href": "Href", "start": "9fab83da-98cb-4f18-a7ba-b6f0435c9673"}, "product_badges": [{"id": "ID", "label": "Label", "description": "Description", "internal_description": "InternalDescription", "learn_more_links": {"first_party": "FirstParty", "third_party": "ThirdParty"}, "get_started_link": "GetStartedLink", "tag": "Tag"}]}`)
				}))
			})
			It(`Invoke ListProductBadges successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the ListProductBadgesOptions model
				listProductBadgesOptionsModel := new(partnercentersellv1.ListProductBadgesOptions)
				listProductBadgesOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProductBadgesOptionsModel.Start = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.ListProductBadgesWithContext(ctx, listProductBadgesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.ListProductBadges(listProductBadgesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.ListProductBadgesWithContext(ctx, listProductBadgesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProductBadgesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					// TODO: Add check for start query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 1, "offset": 0, "total_count": 0, "first": {"href": "Href", "start": "9fab83da-98cb-4f18-a7ba-b6f0435c9673"}, "next": {"href": "Href", "start": "9fab83da-98cb-4f18-a7ba-b6f0435c9673"}, "previous": {"href": "Href", "start": "9fab83da-98cb-4f18-a7ba-b6f0435c9673"}, "last": {"href": "Href", "start": "9fab83da-98cb-4f18-a7ba-b6f0435c9673"}, "product_badges": [{"id": "ID", "label": "Label", "description": "Description", "internal_description": "InternalDescription", "learn_more_links": {"first_party": "FirstParty", "third_party": "ThirdParty"}, "get_started_link": "GetStartedLink", "tag": "Tag"}]}`)
				}))
			})
			It(`Invoke ListProductBadges successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.ListProductBadges(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProductBadgesOptions model
				listProductBadgesOptionsModel := new(partnercentersellv1.ListProductBadgesOptions)
				listProductBadgesOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProductBadgesOptionsModel.Start = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.ListProductBadges(listProductBadgesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProductBadges with error: Operation request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListProductBadgesOptions model
				listProductBadgesOptionsModel := new(partnercentersellv1.ListProductBadgesOptions)
				listProductBadgesOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProductBadgesOptionsModel.Start = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.ListProductBadges(listProductBadgesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListProductBadges successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListProductBadgesOptions model
				listProductBadgesOptionsModel := new(partnercentersellv1.ListProductBadgesOptions)
				listProductBadgesOptionsModel.Limit = core.Int64Ptr(int64(100))
				listProductBadgesOptionsModel.Start = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.ListProductBadges(listProductBadgesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProductBadge(getProductBadgeOptions *GetProductBadgeOptions) - Operation response error`, func() {
		getProductBadgePath := "/product_badges/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProductBadgePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProductBadge with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetProductBadgeOptions model
				getProductBadgeOptionsModel := new(partnercentersellv1.GetProductBadgeOptions)
				getProductBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetProductBadge(getProductBadgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetProductBadge(getProductBadgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProductBadge(getProductBadgeOptions *GetProductBadgeOptions)`, func() {
		getProductBadgePath := "/product_badges/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProductBadgePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "label": "Label", "description": "Description", "internal_description": "InternalDescription", "learn_more_links": {"first_party": "FirstParty", "third_party": "ThirdParty"}, "get_started_link": "GetStartedLink", "tag": "Tag"}`)
				}))
			})
			It(`Invoke GetProductBadge successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetProductBadgeOptions model
				getProductBadgeOptionsModel := new(partnercentersellv1.GetProductBadgeOptions)
				getProductBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetProductBadgeWithContext(ctx, getProductBadgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetProductBadge(getProductBadgeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetProductBadgeWithContext(ctx, getProductBadgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProductBadgePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "label": "Label", "description": "Description", "internal_description": "InternalDescription", "learn_more_links": {"first_party": "FirstParty", "third_party": "ThirdParty"}, "get_started_link": "GetStartedLink", "tag": "Tag"}`)
				}))
			})
			It(`Invoke GetProductBadge successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetProductBadge(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProductBadgeOptions model
				getProductBadgeOptionsModel := new(partnercentersellv1.GetProductBadgeOptions)
				getProductBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetProductBadge(getProductBadgeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProductBadge with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetProductBadgeOptions model
				getProductBadgeOptionsModel := new(partnercentersellv1.GetProductBadgeOptions)
				getProductBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetProductBadge(getProductBadgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProductBadgeOptions model with no property values
				getProductBadgeOptionsModelNew := new(partnercentersellv1.GetProductBadgeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetProductBadge(getProductBadgeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetProductBadge successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetProductBadgeOptions model
				getProductBadgeOptionsModel := new(partnercentersellv1.GetProductBadgeOptions)
				getProductBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetProductBadge(getProductBadgeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			partnerCenterSellService, _ := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
				URL:           "http://partnercentersellv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCatalogProductMediaItem successfully`, func() {
				caption := "testString"
				typeVar := "image"
				url := "testString"
				_model, err := partnerCenterSellService.NewCatalogProductMediaItem(caption, typeVar, url)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateCatalogDeploymentOptions successfully`, func() {
				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				Expect(catalogProductProviderModel).ToNot(BeNil())
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")
				Expect(catalogProductProviderModel.Name).To(Equal(core.StringPtr("IBM")))
				Expect(catalogProductProviderModel.Email).To(Equal(core.StringPtr("name.name@ibm.com")))

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				Expect(globalCatalogOverviewUiTranslatedContentModel).ToNot(BeNil())
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")
				Expect(globalCatalogOverviewUiTranslatedContentModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogOverviewUiTranslatedContentModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogOverviewUiTranslatedContentModel.LongDescription).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				Expect(globalCatalogOverviewUiModel).ToNot(BeNil())
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel
				Expect(globalCatalogOverviewUiModel.En).To(Equal(globalCatalogOverviewUiTranslatedContentModel))

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				Expect(catalogHighlightItemModel).ToNot(BeNil())
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}
				Expect(catalogHighlightItemModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(catalogHighlightItemModel.DescriptionI18n).To(Equal(map[string]string{"key1": "testString"}))
				Expect(catalogHighlightItemModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(catalogHighlightItemModel.TitleI18n).To(Equal(map[string]string{"key1": "testString"}))

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				Expect(catalogProductMediaItemModel).ToNot(BeNil())
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")
				Expect(catalogProductMediaItemModel.Caption).To(Equal(core.StringPtr("testString")))
				Expect(catalogProductMediaItemModel.CaptionI18n).To(Equal(map[string]string{"key1": "testString"}))
				Expect(catalogProductMediaItemModel.Thumbnail).To(Equal(core.StringPtr("testString")))
				Expect(catalogProductMediaItemModel.Type).To(Equal(core.StringPtr("image")))
				Expect(catalogProductMediaItemModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				Expect(globalCatalogMetadataUiStringsContentModel).ToNot(BeNil())
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")
				Expect(globalCatalogMetadataUiStringsContentModel.Bullets).To(Equal([]partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}))
				Expect(globalCatalogMetadataUiStringsContentModel.Media).To(Equal([]partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}))
				Expect(globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				Expect(globalCatalogMetadataUiStringsModel).ToNot(BeNil())
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel
				Expect(globalCatalogMetadataUiStringsModel.En).To(Equal(globalCatalogMetadataUiStringsContentModel))

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				Expect(globalCatalogMetadataUiUrlsModel).ToNot(BeNil())
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")
				Expect(globalCatalogMetadataUiUrlsModel.DocURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.ApidocsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.TermsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.InstructionsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.CatalogDetailsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.CustomCreatePageURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.Dashboard).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				Expect(globalCatalogMetadataUiModel).ToNot(BeNil())
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))
				Expect(globalCatalogMetadataUiModel.Strings).To(Equal(globalCatalogMetadataUiStringsModel))
				Expect(globalCatalogMetadataUiModel.Urls).To(Equal(globalCatalogMetadataUiUrlsModel))
				Expect(globalCatalogMetadataUiModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataUiModel.SideBySideIndex).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				Expect(globalCatalogMetadataServiceModel).ToNot(BeNil())
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)
				Expect(globalCatalogMetadataServiceModel.RcProvisionable).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.IamCompatible).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.PlanUpdateable).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.ServiceKeySupported).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the GlobalCatalogMetadataDeploymentBroker model
				globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
				Expect(globalCatalogMetadataDeploymentBrokerModel).ToNot(BeNil())
				globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("brokerunique1234")
				globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3c2")
				Expect(globalCatalogMetadataDeploymentBrokerModel.Name).To(Equal(core.StringPtr("brokerunique1234")))
				Expect(globalCatalogMetadataDeploymentBrokerModel.Guid).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3c2")))

				// Construct an instance of the GlobalCatalogMetadataDeployment model
				globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
				Expect(globalCatalogMetadataDeploymentModel).ToNot(BeNil())
				globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
				globalCatalogMetadataDeploymentModel.Location = core.StringPtr("eu-gb")
				globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")
				globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")
				Expect(globalCatalogMetadataDeploymentModel.Broker).To(Equal(globalCatalogMetadataDeploymentBrokerModel))
				Expect(globalCatalogMetadataDeploymentModel.Location).To(Equal(core.StringPtr("eu-gb")))
				Expect(globalCatalogMetadataDeploymentModel.LocationURL).To(Equal(core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb")))
				Expect(globalCatalogMetadataDeploymentModel.TargetCrn).To(Equal(core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb")))

				// Construct an instance of the GlobalCatalogDeploymentMetadata model
				globalCatalogDeploymentMetadataModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
				Expect(globalCatalogDeploymentMetadataModel).ToNot(BeNil())
				globalCatalogDeploymentMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogDeploymentMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogDeploymentMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogDeploymentMetadataModel.Deployment = globalCatalogMetadataDeploymentModel
				Expect(globalCatalogDeploymentMetadataModel.RcCompatible).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogDeploymentMetadataModel.Ui).To(Equal(globalCatalogMetadataUiModel))
				Expect(globalCatalogDeploymentMetadataModel.Service).To(Equal(globalCatalogMetadataServiceModel))
				Expect(globalCatalogDeploymentMetadataModel.Deployment).To(Equal(globalCatalogMetadataDeploymentModel))

				// Construct an instance of the CreateCatalogDeploymentOptions model
				productID := "testString"
				catalogProductID := "testString"
				catalogPlanID := "testString"
				createCatalogDeploymentOptionsName := "deployment-eu-de"
				createCatalogDeploymentOptionsActive := true
				createCatalogDeploymentOptionsDisabled := false
				createCatalogDeploymentOptionsKind := "deployment"
				createCatalogDeploymentOptionsTags := []string{"eu-gb"}
				var createCatalogDeploymentOptionsObjectProvider *partnercentersellv1.CatalogProductProvider = nil
				createCatalogDeploymentOptionsModel := partnerCenterSellService.NewCreateCatalogDeploymentOptions(productID, catalogProductID, catalogPlanID, createCatalogDeploymentOptionsName, createCatalogDeploymentOptionsActive, createCatalogDeploymentOptionsDisabled, createCatalogDeploymentOptionsKind, createCatalogDeploymentOptionsTags, createCatalogDeploymentOptionsObjectProvider)
				createCatalogDeploymentOptionsModel.SetProductID("testString")
				createCatalogDeploymentOptionsModel.SetCatalogProductID("testString")
				createCatalogDeploymentOptionsModel.SetCatalogPlanID("testString")
				createCatalogDeploymentOptionsModel.SetName("deployment-eu-de")
				createCatalogDeploymentOptionsModel.SetActive(true)
				createCatalogDeploymentOptionsModel.SetDisabled(false)
				createCatalogDeploymentOptionsModel.SetKind("deployment")
				createCatalogDeploymentOptionsModel.SetTags([]string{"eu-gb"})
				createCatalogDeploymentOptionsModel.SetObjectProvider(catalogProductProviderModel)
				createCatalogDeploymentOptionsModel.SetID("testString")
				createCatalogDeploymentOptionsModel.SetObjectID("testString")
				createCatalogDeploymentOptionsModel.SetOverviewUi(globalCatalogOverviewUiModel)
				createCatalogDeploymentOptionsModel.SetMetadata(globalCatalogDeploymentMetadataModel)
				createCatalogDeploymentOptionsModel.SetEnv("testString")
				createCatalogDeploymentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCatalogDeploymentOptionsModel).ToNot(BeNil())
				Expect(createCatalogDeploymentOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogDeploymentOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogDeploymentOptionsModel.CatalogPlanID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogDeploymentOptionsModel.Name).To(Equal(core.StringPtr("deployment-eu-de")))
				Expect(createCatalogDeploymentOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogDeploymentOptionsModel.Disabled).To(Equal(core.BoolPtr(false)))
				Expect(createCatalogDeploymentOptionsModel.Kind).To(Equal(core.StringPtr("deployment")))
				Expect(createCatalogDeploymentOptionsModel.Tags).To(Equal([]string{"eu-gb"}))
				Expect(createCatalogDeploymentOptionsModel.ObjectProvider).To(Equal(catalogProductProviderModel))
				Expect(createCatalogDeploymentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogDeploymentOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogDeploymentOptionsModel.OverviewUi).To(Equal(globalCatalogOverviewUiModel))
				Expect(createCatalogDeploymentOptionsModel.Metadata).To(Equal(globalCatalogDeploymentMetadataModel))
				Expect(createCatalogDeploymentOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogDeploymentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCatalogPlanOptions successfully`, func() {
				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				Expect(catalogProductProviderModel).ToNot(BeNil())
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")
				Expect(catalogProductProviderModel.Name).To(Equal(core.StringPtr("IBM")))
				Expect(catalogProductProviderModel.Email).To(Equal(core.StringPtr("name.name@ibm.com")))

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				Expect(globalCatalogOverviewUiTranslatedContentModel).ToNot(BeNil())
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My plan")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My plan description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My plan long description.")
				Expect(globalCatalogOverviewUiTranslatedContentModel.DisplayName).To(Equal(core.StringPtr("My plan")))
				Expect(globalCatalogOverviewUiTranslatedContentModel.Description).To(Equal(core.StringPtr("My plan description.")))
				Expect(globalCatalogOverviewUiTranslatedContentModel.LongDescription).To(Equal(core.StringPtr("My plan long description.")))

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				Expect(globalCatalogOverviewUiModel).ToNot(BeNil())
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel
				Expect(globalCatalogOverviewUiModel.En).To(Equal(globalCatalogOverviewUiTranslatedContentModel))

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				Expect(catalogHighlightItemModel).ToNot(BeNil())
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}
				Expect(catalogHighlightItemModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(catalogHighlightItemModel.DescriptionI18n).To(Equal(map[string]string{"key1": "testString"}))
				Expect(catalogHighlightItemModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(catalogHighlightItemModel.TitleI18n).To(Equal(map[string]string{"key1": "testString"}))

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				Expect(catalogProductMediaItemModel).ToNot(BeNil())
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")
				Expect(catalogProductMediaItemModel.Caption).To(Equal(core.StringPtr("testString")))
				Expect(catalogProductMediaItemModel.CaptionI18n).To(Equal(map[string]string{"key1": "testString"}))
				Expect(catalogProductMediaItemModel.Thumbnail).To(Equal(core.StringPtr("testString")))
				Expect(catalogProductMediaItemModel.Type).To(Equal(core.StringPtr("image")))
				Expect(catalogProductMediaItemModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				Expect(globalCatalogMetadataUiStringsContentModel).ToNot(BeNil())
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")
				Expect(globalCatalogMetadataUiStringsContentModel.Bullets).To(Equal([]partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}))
				Expect(globalCatalogMetadataUiStringsContentModel.Media).To(Equal([]partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}))
				Expect(globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				Expect(globalCatalogMetadataUiStringsModel).ToNot(BeNil())
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel
				Expect(globalCatalogMetadataUiStringsModel.En).To(Equal(globalCatalogMetadataUiStringsContentModel))

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				Expect(globalCatalogMetadataUiUrlsModel).ToNot(BeNil())
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")
				Expect(globalCatalogMetadataUiUrlsModel.DocURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.ApidocsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.TermsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.InstructionsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.CatalogDetailsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.CustomCreatePageURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.Dashboard).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				Expect(globalCatalogMetadataUiModel).ToNot(BeNil())
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))
				Expect(globalCatalogMetadataUiModel.Strings).To(Equal(globalCatalogMetadataUiStringsModel))
				Expect(globalCatalogMetadataUiModel.Urls).To(Equal(globalCatalogMetadataUiUrlsModel))
				Expect(globalCatalogMetadataUiModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataUiModel.SideBySideIndex).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				Expect(globalCatalogMetadataServiceModel).ToNot(BeNil())
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(false)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)
				Expect(globalCatalogMetadataServiceModel.RcProvisionable).To(Equal(core.BoolPtr(false)))
				Expect(globalCatalogMetadataServiceModel.IamCompatible).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.PlanUpdateable).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.ServiceKeySupported).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the GlobalCatalogMetadataPricing model
				globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
				Expect(globalCatalogMetadataPricingModel).ToNot(BeNil())
				globalCatalogMetadataPricingModel.Type = core.StringPtr("paid")
				globalCatalogMetadataPricingModel.Origin = core.StringPtr("pricing_catalog")
				Expect(globalCatalogMetadataPricingModel.Type).To(Equal(core.StringPtr("paid")))
				Expect(globalCatalogMetadataPricingModel.Origin).To(Equal(core.StringPtr("pricing_catalog")))

				// Construct an instance of the GlobalCatalogPlanMetadataPlan model
				globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
				Expect(globalCatalogPlanMetadataPlanModel).ToNot(BeNil())
				globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				globalCatalogPlanMetadataPlanModel.Bindable = core.BoolPtr(true)
				Expect(globalCatalogPlanMetadataPlanModel.AllowInternalUsers).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogPlanMetadataPlanModel.Bindable).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the GlobalCatalogPlanMetadata model
				globalCatalogPlanMetadataModel := new(partnercentersellv1.GlobalCatalogPlanMetadata)
				Expect(globalCatalogPlanMetadataModel).ToNot(BeNil())
				globalCatalogPlanMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogPlanMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogPlanMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogPlanMetadataModel.Pricing = globalCatalogMetadataPricingModel
				globalCatalogPlanMetadataModel.Plan = globalCatalogPlanMetadataPlanModel
				Expect(globalCatalogPlanMetadataModel.RcCompatible).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogPlanMetadataModel.Ui).To(Equal(globalCatalogMetadataUiModel))
				Expect(globalCatalogPlanMetadataModel.Service).To(Equal(globalCatalogMetadataServiceModel))
				Expect(globalCatalogPlanMetadataModel.Pricing).To(Equal(globalCatalogMetadataPricingModel))
				Expect(globalCatalogPlanMetadataModel.Plan).To(Equal(globalCatalogPlanMetadataPlanModel))

				// Construct an instance of the CreateCatalogPlanOptions model
				productID := "testString"
				catalogProductID := "testString"
				createCatalogPlanOptionsName := "free-plan2"
				createCatalogPlanOptionsActive := true
				createCatalogPlanOptionsDisabled := false
				createCatalogPlanOptionsKind := "plan"
				createCatalogPlanOptionsTags := []string{"ibm_created"}
				var createCatalogPlanOptionsObjectProvider *partnercentersellv1.CatalogProductProvider = nil
				createCatalogPlanOptionsModel := partnerCenterSellService.NewCreateCatalogPlanOptions(productID, catalogProductID, createCatalogPlanOptionsName, createCatalogPlanOptionsActive, createCatalogPlanOptionsDisabled, createCatalogPlanOptionsKind, createCatalogPlanOptionsTags, createCatalogPlanOptionsObjectProvider)
				createCatalogPlanOptionsModel.SetProductID("testString")
				createCatalogPlanOptionsModel.SetCatalogProductID("testString")
				createCatalogPlanOptionsModel.SetName("free-plan2")
				createCatalogPlanOptionsModel.SetActive(true)
				createCatalogPlanOptionsModel.SetDisabled(false)
				createCatalogPlanOptionsModel.SetKind("plan")
				createCatalogPlanOptionsModel.SetTags([]string{"ibm_created"})
				createCatalogPlanOptionsModel.SetObjectProvider(catalogProductProviderModel)
				createCatalogPlanOptionsModel.SetID("testString")
				createCatalogPlanOptionsModel.SetObjectID("testString")
				createCatalogPlanOptionsModel.SetOverviewUi(globalCatalogOverviewUiModel)
				createCatalogPlanOptionsModel.SetMetadata(globalCatalogPlanMetadataModel)
				createCatalogPlanOptionsModel.SetEnv("testString")
				createCatalogPlanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCatalogPlanOptionsModel).ToNot(BeNil())
				Expect(createCatalogPlanOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogPlanOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogPlanOptionsModel.Name).To(Equal(core.StringPtr("free-plan2")))
				Expect(createCatalogPlanOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogPlanOptionsModel.Disabled).To(Equal(core.BoolPtr(false)))
				Expect(createCatalogPlanOptionsModel.Kind).To(Equal(core.StringPtr("plan")))
				Expect(createCatalogPlanOptionsModel.Tags).To(Equal([]string{"ibm_created"}))
				Expect(createCatalogPlanOptionsModel.ObjectProvider).To(Equal(catalogProductProviderModel))
				Expect(createCatalogPlanOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogPlanOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogPlanOptionsModel.OverviewUi).To(Equal(globalCatalogOverviewUiModel))
				Expect(createCatalogPlanOptionsModel.Metadata).To(Equal(globalCatalogPlanMetadataModel))
				Expect(createCatalogPlanOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogPlanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCatalogProductOptions successfully`, func() {
				// Construct an instance of the CatalogProductProvider model
				catalogProductProviderModel := new(partnercentersellv1.CatalogProductProvider)
				Expect(catalogProductProviderModel).ToNot(BeNil())
				catalogProductProviderModel.Name = core.StringPtr("IBM")
				catalogProductProviderModel.Email = core.StringPtr("name.name@ibm.com")
				Expect(catalogProductProviderModel.Name).To(Equal(core.StringPtr("IBM")))
				Expect(catalogProductProviderModel.Email).To(Equal(core.StringPtr("name.name@ibm.com")))

				// Construct an instance of the GlobalCatalogOverviewUITranslatedContent model
				globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
				Expect(globalCatalogOverviewUiTranslatedContentModel).ToNot(BeNil())
				globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("My product display name.")
				globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("My product description.")
				globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("My product description long description.")
				Expect(globalCatalogOverviewUiTranslatedContentModel.DisplayName).To(Equal(core.StringPtr("My product display name.")))
				Expect(globalCatalogOverviewUiTranslatedContentModel.Description).To(Equal(core.StringPtr("My product description.")))
				Expect(globalCatalogOverviewUiTranslatedContentModel.LongDescription).To(Equal(core.StringPtr("My product description long description.")))

				// Construct an instance of the GlobalCatalogOverviewUI model
				globalCatalogOverviewUiModel := new(partnercentersellv1.GlobalCatalogOverviewUI)
				Expect(globalCatalogOverviewUiModel).ToNot(BeNil())
				globalCatalogOverviewUiModel.En = globalCatalogOverviewUiTranslatedContentModel
				Expect(globalCatalogOverviewUiModel.En).To(Equal(globalCatalogOverviewUiTranslatedContentModel))

				// Construct an instance of the GlobalCatalogProductImages model
				globalCatalogProductImagesModel := new(partnercentersellv1.GlobalCatalogProductImages)
				Expect(globalCatalogProductImagesModel).ToNot(BeNil())
				globalCatalogProductImagesModel.Image = core.StringPtr("testString")
				Expect(globalCatalogProductImagesModel.Image).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CatalogHighlightItem model
				catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
				Expect(catalogHighlightItemModel).ToNot(BeNil())
				catalogHighlightItemModel.Description = core.StringPtr("testString")
				catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
				catalogHighlightItemModel.Title = core.StringPtr("testString")
				catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}
				Expect(catalogHighlightItemModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(catalogHighlightItemModel.DescriptionI18n).To(Equal(map[string]string{"key1": "testString"}))
				Expect(catalogHighlightItemModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(catalogHighlightItemModel.TitleI18n).To(Equal(map[string]string{"key1": "testString"}))

				// Construct an instance of the CatalogProductMediaItem model
				catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
				Expect(catalogProductMediaItemModel).ToNot(BeNil())
				catalogProductMediaItemModel.Caption = core.StringPtr("testString")
				catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
				catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
				catalogProductMediaItemModel.Type = core.StringPtr("image")
				catalogProductMediaItemModel.URL = core.StringPtr("testString")
				Expect(catalogProductMediaItemModel.Caption).To(Equal(core.StringPtr("testString")))
				Expect(catalogProductMediaItemModel.CaptionI18n).To(Equal(map[string]string{"key1": "testString"}))
				Expect(catalogProductMediaItemModel.Thumbnail).To(Equal(core.StringPtr("testString")))
				Expect(catalogProductMediaItemModel.Type).To(Equal(core.StringPtr("image")))
				Expect(catalogProductMediaItemModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogMetadataUIStringsContent model
				globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
				Expect(globalCatalogMetadataUiStringsContentModel).ToNot(BeNil())
				globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
				globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
				globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")
				Expect(globalCatalogMetadataUiStringsContentModel.Bullets).To(Equal([]partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}))
				Expect(globalCatalogMetadataUiStringsContentModel.Media).To(Equal([]partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}))
				Expect(globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogMetadataUIStrings model
				globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
				Expect(globalCatalogMetadataUiStringsModel).ToNot(BeNil())
				globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel
				Expect(globalCatalogMetadataUiStringsModel.En).To(Equal(globalCatalogMetadataUiStringsContentModel))

				// Construct an instance of the GlobalCatalogMetadataUIUrls model
				globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
				Expect(globalCatalogMetadataUiUrlsModel).ToNot(BeNil())
				globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
				globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")
				Expect(globalCatalogMetadataUiUrlsModel.DocURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.ApidocsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.TermsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.InstructionsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.CatalogDetailsURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.CustomCreatePageURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogMetadataUiUrlsModel.Dashboard).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogMetadataUI model
				globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
				Expect(globalCatalogMetadataUiModel).ToNot(BeNil())
				globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
				globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
				globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
				globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))
				Expect(globalCatalogMetadataUiModel.Strings).To(Equal(globalCatalogMetadataUiStringsModel))
				Expect(globalCatalogMetadataUiModel.Urls).To(Equal(globalCatalogMetadataUiUrlsModel))
				Expect(globalCatalogMetadataUiModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataUiModel.SideBySideIndex).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the GlobalCatalogMetadataService model
				globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
				Expect(globalCatalogMetadataServiceModel).ToNot(BeNil())
				globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.Bindable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
				globalCatalogMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)
				Expect(globalCatalogMetadataServiceModel.RcProvisionable).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.IamCompatible).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.PlanUpdateable).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogMetadataServiceModel.ServiceKeySupported).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the SupportTimeInterval model
				supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
				Expect(supportTimeIntervalModel).ToNot(BeNil())
				supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
				supportTimeIntervalModel.Type = core.StringPtr("testString")
				Expect(supportTimeIntervalModel.Value).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(supportTimeIntervalModel.Type).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				Expect(supportEscalationModel).ToNot(BeNil())
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
				supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel
				Expect(supportEscalationModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(supportEscalationModel.EscalationWaitTime).To(Equal(supportTimeIntervalModel))
				Expect(supportEscalationModel.ResponseWaitTime).To(Equal(supportTimeIntervalModel))

				// Construct an instance of the SupportDetailsItemAvailabilityTime model
				supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
				Expect(supportDetailsItemAvailabilityTimeModel).ToNot(BeNil())
				supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
				supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
				supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")
				Expect(supportDetailsItemAvailabilityTimeModel.Day).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(supportDetailsItemAvailabilityTimeModel.StartTime).To(Equal(core.StringPtr("testString")))
				Expect(supportDetailsItemAvailabilityTimeModel.EndTime).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SupportDetailsItemAvailability model
				supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
				Expect(supportDetailsItemAvailabilityModel).ToNot(BeNil())
				supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
				supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
				supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				Expect(supportDetailsItemAvailabilityModel.Times).To(Equal([]partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}))
				Expect(supportDetailsItemAvailabilityModel.Timezone).To(Equal(core.StringPtr("testString")))
				Expect(supportDetailsItemAvailabilityModel.AlwaysAvailable).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the SupportDetailsItem model
				supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
				Expect(supportDetailsItemModel).ToNot(BeNil())
				supportDetailsItemModel.Type = core.StringPtr("support_site")
				supportDetailsItemModel.Contact = core.StringPtr("testString")
				supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
				supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel
				Expect(supportDetailsItemModel.Type).To(Equal(core.StringPtr("support_site")))
				Expect(supportDetailsItemModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(supportDetailsItemModel.ResponseWaitTime).To(Equal(supportTimeIntervalModel))
				Expect(supportDetailsItemModel.Availability).To(Equal(supportDetailsItemAvailabilityModel))

				// Construct an instance of the GlobalCatalogProductMetadataOtherPCSupport model
				globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
				Expect(globalCatalogProductMetadataOtherPcSupportModel).ToNot(BeNil())
				globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
				globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
				globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
				globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
				globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
				globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}
				Expect(globalCatalogProductMetadataOtherPcSupportModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogProductMetadataOtherPcSupportModel.StatusURL).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogProductMetadataOtherPcSupportModel.Locations).To(Equal([]string{"testString"}))
				Expect(globalCatalogProductMetadataOtherPcSupportModel.Languages).To(Equal([]string{"testString"}))
				Expect(globalCatalogProductMetadataOtherPcSupportModel.Process).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n).To(Equal(map[string]string{"key1": "testString"}))
				Expect(globalCatalogProductMetadataOtherPcSupportModel.SupportType).To(Equal(core.StringPtr("community")))
				Expect(globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation).To(Equal(supportEscalationModel))
				Expect(globalCatalogProductMetadataOtherPcSupportModel.SupportDetails).To(Equal([]partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}))

				// Construct an instance of the GlobalCatalogProductMetadataOtherPC model
				globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
				Expect(globalCatalogProductMetadataOtherPcModel).ToNot(BeNil())
				globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel
				Expect(globalCatalogProductMetadataOtherPcModel.Support).To(Equal(globalCatalogProductMetadataOtherPcSupportModel))

				// Construct an instance of the GlobalCatalogProductMetadataOtherCompositeChild model
				globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
				Expect(globalCatalogProductMetadataOtherCompositeChildModel).ToNot(BeNil())
				globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")
				Expect(globalCatalogProductMetadataOtherCompositeChildModel.Kind).To(Equal(core.StringPtr("service")))
				Expect(globalCatalogProductMetadataOtherCompositeChildModel.Name).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GlobalCatalogProductMetadataOtherComposite model
				globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
				Expect(globalCatalogProductMetadataOtherCompositeModel).ToNot(BeNil())
				globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
				globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
				globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}
				Expect(globalCatalogProductMetadataOtherCompositeModel.CompositeKind).To(Equal(core.StringPtr("service")))
				Expect(globalCatalogProductMetadataOtherCompositeModel.CompositeTag).To(Equal(core.StringPtr("testString")))
				Expect(globalCatalogProductMetadataOtherCompositeModel.Children).To(Equal([]partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}))

				// Construct an instance of the GlobalCatalogProductMetadataOther model
				globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
				Expect(globalCatalogProductMetadataOtherModel).ToNot(BeNil())
				globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
				globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel
				Expect(globalCatalogProductMetadataOtherModel.PC).To(Equal(globalCatalogProductMetadataOtherPcModel))
				Expect(globalCatalogProductMetadataOtherModel.Composite).To(Equal(globalCatalogProductMetadataOtherCompositeModel))

				// Construct an instance of the GlobalCatalogProductMetadata model
				globalCatalogProductMetadataModel := new(partnercentersellv1.GlobalCatalogProductMetadata)
				Expect(globalCatalogProductMetadataModel).ToNot(BeNil())
				globalCatalogProductMetadataModel.RcCompatible = core.BoolPtr(true)
				globalCatalogProductMetadataModel.Ui = globalCatalogMetadataUiModel
				globalCatalogProductMetadataModel.Service = globalCatalogMetadataServiceModel
				globalCatalogProductMetadataModel.Other = globalCatalogProductMetadataOtherModel
				Expect(globalCatalogProductMetadataModel.RcCompatible).To(Equal(core.BoolPtr(true)))
				Expect(globalCatalogProductMetadataModel.Ui).To(Equal(globalCatalogMetadataUiModel))
				Expect(globalCatalogProductMetadataModel.Service).To(Equal(globalCatalogMetadataServiceModel))
				Expect(globalCatalogProductMetadataModel.Other).To(Equal(globalCatalogProductMetadataOtherModel))

				// Construct an instance of the CreateCatalogProductOptions model
				productID := "testString"
				createCatalogProductOptionsName := "1p-service-08-06"
				createCatalogProductOptionsActive := true
				createCatalogProductOptionsDisabled := false
				createCatalogProductOptionsKind := "service"
				createCatalogProductOptionsTags := []string{"keyword", "support_ibm"}
				var createCatalogProductOptionsObjectProvider *partnercentersellv1.CatalogProductProvider = nil
				createCatalogProductOptionsModel := partnerCenterSellService.NewCreateCatalogProductOptions(productID, createCatalogProductOptionsName, createCatalogProductOptionsActive, createCatalogProductOptionsDisabled, createCatalogProductOptionsKind, createCatalogProductOptionsTags, createCatalogProductOptionsObjectProvider)
				createCatalogProductOptionsModel.SetProductID("testString")
				createCatalogProductOptionsModel.SetName("1p-service-08-06")
				createCatalogProductOptionsModel.SetActive(true)
				createCatalogProductOptionsModel.SetDisabled(false)
				createCatalogProductOptionsModel.SetKind("service")
				createCatalogProductOptionsModel.SetTags([]string{"keyword", "support_ibm"})
				createCatalogProductOptionsModel.SetObjectProvider(catalogProductProviderModel)
				createCatalogProductOptionsModel.SetID("testString")
				createCatalogProductOptionsModel.SetObjectID("testString")
				createCatalogProductOptionsModel.SetOverviewUi(globalCatalogOverviewUiModel)
				createCatalogProductOptionsModel.SetImages(globalCatalogProductImagesModel)
				createCatalogProductOptionsModel.SetMetadata(globalCatalogProductMetadataModel)
				createCatalogProductOptionsModel.SetEnv("testString")
				createCatalogProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCatalogProductOptionsModel).ToNot(BeNil())
				Expect(createCatalogProductOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogProductOptionsModel.Name).To(Equal(core.StringPtr("1p-service-08-06")))
				Expect(createCatalogProductOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogProductOptionsModel.Disabled).To(Equal(core.BoolPtr(false)))
				Expect(createCatalogProductOptionsModel.Kind).To(Equal(core.StringPtr("service")))
				Expect(createCatalogProductOptionsModel.Tags).To(Equal([]string{"keyword", "support_ibm"}))
				Expect(createCatalogProductOptionsModel.ObjectProvider).To(Equal(catalogProductProviderModel))
				Expect(createCatalogProductOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogProductOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogProductOptionsModel.OverviewUi).To(Equal(globalCatalogOverviewUiModel))
				Expect(createCatalogProductOptionsModel.Images).To(Equal(globalCatalogProductImagesModel))
				Expect(createCatalogProductOptionsModel.Metadata).To(Equal(globalCatalogProductMetadataModel))
				Expect(createCatalogProductOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateIamRegistrationOptions successfully`, func() {
				// Construct an instance of the IamServiceRegistrationDescriptionObject model
				iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
				Expect(iamServiceRegistrationDescriptionObjectModel).ToNot(BeNil())
				iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("View dashboard")
				Expect(iamServiceRegistrationDescriptionObjectModel.Default).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.En).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.De).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.Es).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.Fr).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.It).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.Ja).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.Ko).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.PtBr).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.ZhTw).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDescriptionObjectModel.ZhCn).To(Equal(core.StringPtr("View dashboard")))

				// Construct an instance of the IamServiceRegistrationDisplayNameObject model
				iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
				Expect(iamServiceRegistrationDisplayNameObjectModel).ToNot(BeNil())
				iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("View dashboard")
				iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("View dashboard")
				Expect(iamServiceRegistrationDisplayNameObjectModel.Default).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.En).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.De).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.Es).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.Fr).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.It).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.Ja).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.Ko).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.PtBr).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.ZhTw).To(Equal(core.StringPtr("View dashboard")))
				Expect(iamServiceRegistrationDisplayNameObjectModel.ZhCn).To(Equal(core.StringPtr("View dashboard")))

				// Construct an instance of the IamServiceRegistrationActionOptions model
				iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
				Expect(iamServiceRegistrationActionOptionsModel).ToNot(BeNil())
				iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)
				Expect(iamServiceRegistrationActionOptionsModel.Hidden).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the IamServiceRegistrationAction model
				iamServiceRegistrationActionModel := new(partnercentersellv1.IamServiceRegistrationAction)
				Expect(iamServiceRegistrationActionModel).ToNot(BeNil())
				iamServiceRegistrationActionModel.ID = core.StringPtr("pet-store.dashboard.view")
				iamServiceRegistrationActionModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}
				iamServiceRegistrationActionModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationActionModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationActionModel.Options = iamServiceRegistrationActionOptionsModel
				Expect(iamServiceRegistrationActionModel.ID).To(Equal(core.StringPtr("pet-store.dashboard.view")))
				Expect(iamServiceRegistrationActionModel.Roles).To(Equal([]string{"crn:v1:bluemix:public:iam::::serviceRole:Reader", "crn:v1:bluemix:public:iam::::serviceRole:Manager", "crn:v1:bluemix:public:iam::::serviceRole:Writer", "crn:v1:bluemix:public:iam::::role:Operator"}))
				Expect(iamServiceRegistrationActionModel.Description).To(Equal(iamServiceRegistrationDescriptionObjectModel))
				Expect(iamServiceRegistrationActionModel.DisplayName).To(Equal(iamServiceRegistrationDisplayNameObjectModel))
				Expect(iamServiceRegistrationActionModel.Options).To(Equal(iamServiceRegistrationActionOptionsModel))

				// Construct an instance of the IamServiceRegistrationResourceHierarchyAttribute model
				iamServiceRegistrationResourceHierarchyAttributeModel := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
				Expect(iamServiceRegistrationResourceHierarchyAttributeModel).ToNot(BeNil())
				iamServiceRegistrationResourceHierarchyAttributeModel.Key = core.StringPtr("testString")
				iamServiceRegistrationResourceHierarchyAttributeModel.Value = core.StringPtr("testString")
				Expect(iamServiceRegistrationResourceHierarchyAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(iamServiceRegistrationResourceHierarchyAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccessAttributes model
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModel).ToNot(BeNil())
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("serviceName", core.StringPtr("pet-store"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("accountId", core.StringPtr("25543245345"))
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperty("testAttribute", core.StringPtr("dsgdsfgsd576456"))
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties).To(Equal(map[string]string{"key1": "testString"}))
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModel.GetProperties()).ToNot(BeEmpty())
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModel.GetProperty("serviceName")).To(Equal(core.StringPtr("pet-store")))
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModel.GetProperty("accountId")).To(Equal(core.StringPtr("25543245345")))
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModel.GetProperty("testAttribute")).To(Equal(core.StringPtr("dsgdsfgsd576456")))

				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperties(nil)
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModel.GetProperties()).To(BeEmpty())

				iamServiceRegistrationSupportedAnonymousAccessAttributesModelExpectedMap := make(map[string]*string)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModelExpectedMap["serviceName"] = core.StringPtr("pet-store")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModelExpectedMap["accountId"] = core.StringPtr("25543245345")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModelExpectedMap["testAttribute"] = core.StringPtr("dsgdsfgsd576456")
				iamServiceRegistrationSupportedAnonymousAccessAttributesModel.SetProperties(iamServiceRegistrationSupportedAnonymousAccessAttributesModelExpectedMap)
				iamServiceRegistrationSupportedAnonymousAccessAttributesModelActualMap := iamServiceRegistrationSupportedAnonymousAccessAttributesModel.GetProperties()
				Expect(iamServiceRegistrationSupportedAnonymousAccessAttributesModelActualMap).To(Equal(iamServiceRegistrationSupportedAnonymousAccessAttributesModelExpectedMap))

				// Construct an instance of the IamServiceRegistrationSupportedAnonymousAccess model
				iamServiceRegistrationSupportedAnonymousAccessModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
				Expect(iamServiceRegistrationSupportedAnonymousAccessModel).ToNot(BeNil())
				iamServiceRegistrationSupportedAnonymousAccessModel.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
				iamServiceRegistrationSupportedAnonymousAccessModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}
				Expect(iamServiceRegistrationSupportedAnonymousAccessModel.Attributes).To(Equal(iamServiceRegistrationSupportedAnonymousAccessAttributesModel))
				Expect(iamServiceRegistrationSupportedAnonymousAccessModel.Roles).To(Equal([]string{"crn:v1:bluemix:public:iam::::serviceRole:Reader"}))

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyKey model
				supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
				Expect(supportedAttributesOptionsResourceHierarchyKeyModel).ToNot(BeNil())
				supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")
				Expect(supportedAttributesOptionsResourceHierarchyKeyModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(supportedAttributesOptionsResourceHierarchyKeyModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchyValue model
				supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
				Expect(supportedAttributesOptionsResourceHierarchyValueModel).ToNot(BeNil())
				supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")
				Expect(supportedAttributesOptionsResourceHierarchyValueModel.Key).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SupportedAttributesOptionsResourceHierarchy model
				supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
				Expect(supportedAttributesOptionsResourceHierarchyModel).ToNot(BeNil())
				supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
				supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel
				Expect(supportedAttributesOptionsResourceHierarchyModel.Key).To(Equal(supportedAttributesOptionsResourceHierarchyKeyModel))
				Expect(supportedAttributesOptionsResourceHierarchyModel.Value).To(Equal(supportedAttributesOptionsResourceHierarchyValueModel))

				// Construct an instance of the SupportedAttributesOptions model
				supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
				Expect(supportedAttributesOptionsModel).ToNot(BeNil())
				supportedAttributesOptionsModel.Operators = []string{"stringMatch", "stringEquals"}
				supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
				supportedAttributesOptionsModel.SupportedAttributes = []string{"testString"}
				supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
				supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
				supportedAttributesOptionsModel.Key = core.StringPtr("testString")
				supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel
				Expect(supportedAttributesOptionsModel.Operators).To(Equal([]string{"stringMatch", "stringEquals"}))
				Expect(supportedAttributesOptionsModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(supportedAttributesOptionsModel.SupportedAttributes).To(Equal([]string{"testString"}))
				Expect(supportedAttributesOptionsModel.PolicyTypes).To(Equal([]string{"access"}))
				Expect(supportedAttributesOptionsModel.IsEmptyValueSupported).To(Equal(core.BoolPtr(true)))
				Expect(supportedAttributesOptionsModel.IsStringExistsFalseValueSupported).To(Equal(core.BoolPtr(true)))
				Expect(supportedAttributesOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(supportedAttributesOptionsModel.ResourceHierarchy).To(Equal(supportedAttributesOptionsResourceHierarchyModel))

				// Construct an instance of the SupportedAttributeUiInputValue model
				supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
				Expect(supportedAttributeUiInputValueModel).ToNot(BeNil())
				supportedAttributeUiInputValueModel.Value = core.StringPtr("staticValue")
				supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				Expect(supportedAttributeUiInputValueModel.Value).To(Equal(core.StringPtr("staticValue")))
				Expect(supportedAttributeUiInputValueModel.DisplayName).To(Equal(iamServiceRegistrationDisplayNameObjectModel))

				// Construct an instance of the SupportedAttributeUiInputGst model
				supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
				Expect(supportedAttributeUiInputGstModel).ToNot(BeNil())
				supportedAttributeUiInputGstModel.Query = core.StringPtr("ghost query")
				supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("instance")
				supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
				supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("{name} - {instance_id}")
				Expect(supportedAttributeUiInputGstModel.Query).To(Equal(core.StringPtr("ghost query")))
				Expect(supportedAttributeUiInputGstModel.ValuePropertyName).To(Equal(core.StringPtr("instance")))
				Expect(supportedAttributeUiInputGstModel.LabelPropertyName).To(Equal(core.StringPtr("testString")))
				Expect(supportedAttributeUiInputGstModel.InputOptionLabel).To(Equal(core.StringPtr("{name} - {instance_id}")))

				// Construct an instance of the SupportedAttributeUiInputURL model
				supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
				Expect(supportedAttributeUiInputUrlModel).ToNot(BeNil())
				supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
				supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")
				Expect(supportedAttributeUiInputUrlModel.UrlEndpoint).To(Equal(core.StringPtr("testString")))
				Expect(supportedAttributeUiInputUrlModel.InputOptionLabel).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SupportedAttributeUiInputDetails model
				supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
				Expect(supportedAttributeUiInputDetailsModel).ToNot(BeNil())
				supportedAttributeUiInputDetailsModel.Type = core.StringPtr("gst")
				supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
				supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
				supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel
				Expect(supportedAttributeUiInputDetailsModel.Type).To(Equal(core.StringPtr("gst")))
				Expect(supportedAttributeUiInputDetailsModel.Values).To(Equal([]partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}))
				Expect(supportedAttributeUiInputDetailsModel.Gst).To(Equal(supportedAttributeUiInputGstModel))
				Expect(supportedAttributeUiInputDetailsModel.URL).To(Equal(supportedAttributeUiInputUrlModel))

				// Construct an instance of the SupportedAttributeUi model
				supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
				Expect(supportedAttributeUiModel).ToNot(BeNil())
				supportedAttributeUiModel.InputType = core.StringPtr("selector")
				supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel
				Expect(supportedAttributeUiModel.InputType).To(Equal(core.StringPtr("selector")))
				Expect(supportedAttributeUiModel.InputDetails).To(Equal(supportedAttributeUiInputDetailsModel))

				// Construct an instance of the IamServiceRegistrationSupportedAttribute model
				iamServiceRegistrationSupportedAttributeModel := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
				Expect(iamServiceRegistrationSupportedAttributeModel).ToNot(BeNil())
				iamServiceRegistrationSupportedAttributeModel.Key = core.StringPtr("testAttribute")
				iamServiceRegistrationSupportedAttributeModel.Options = supportedAttributesOptionsModel
				iamServiceRegistrationSupportedAttributeModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedAttributeModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedAttributeModel.Ui = supportedAttributeUiModel
				Expect(iamServiceRegistrationSupportedAttributeModel.Key).To(Equal(core.StringPtr("testAttribute")))
				Expect(iamServiceRegistrationSupportedAttributeModel.Options).To(Equal(supportedAttributesOptionsModel))
				Expect(iamServiceRegistrationSupportedAttributeModel.DisplayName).To(Equal(iamServiceRegistrationDisplayNameObjectModel))
				Expect(iamServiceRegistrationSupportedAttributeModel.Description).To(Equal(iamServiceRegistrationDescriptionObjectModel))
				Expect(iamServiceRegistrationSupportedAttributeModel.Ui).To(Equal(supportedAttributeUiModel))

				// Construct an instance of the SupportAuthorizationSubjectAttribute model
				supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
				Expect(supportAuthorizationSubjectAttributeModel).ToNot(BeNil())
				supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
				supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")
				Expect(supportAuthorizationSubjectAttributeModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(supportAuthorizationSubjectAttributeModel.ResourceType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the IamServiceRegistrationSupportedAuthorizationSubject model
				iamServiceRegistrationSupportedAuthorizationSubjectModel := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
				Expect(iamServiceRegistrationSupportedAuthorizationSubjectModel).ToNot(BeNil())
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes = supportAuthorizationSubjectAttributeModel
				iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles = []string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}
				Expect(iamServiceRegistrationSupportedAuthorizationSubjectModel.Attributes).To(Equal(supportAuthorizationSubjectAttributeModel))
				Expect(iamServiceRegistrationSupportedAuthorizationSubjectModel.Roles).To(Equal([]string{"crn:v1:bluemix:public:iam::::serviceRole:Writer"}))

				// Construct an instance of the SupportedRoleOptions model
				supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
				Expect(supportedRoleOptionsModel).ToNot(BeNil())
				supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
				supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
				supportedRoleOptionsModel.PolicyType = []string{"access"}
				supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")
				Expect(supportedRoleOptionsModel.AccessPolicy).To(Equal(core.BoolPtr(true)))
				Expect(supportedRoleOptionsModel.AdditionalPropertiesForAccessPolicy).To(Equal(map[string]string{"key1": "testString"}))
				Expect(supportedRoleOptionsModel.PolicyType).To(Equal([]string{"access"}))
				Expect(supportedRoleOptionsModel.AccountType).To(Equal(core.StringPtr("enterprise")))

				// Construct an instance of the IamServiceRegistrationSupportedRole model
				iamServiceRegistrationSupportedRoleModel := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
				Expect(iamServiceRegistrationSupportedRoleModel).ToNot(BeNil())
				iamServiceRegistrationSupportedRoleModel.ID = core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")
				iamServiceRegistrationSupportedRoleModel.Description = iamServiceRegistrationDescriptionObjectModel
				iamServiceRegistrationSupportedRoleModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel
				iamServiceRegistrationSupportedRoleModel.Options = supportedRoleOptionsModel
				Expect(iamServiceRegistrationSupportedRoleModel.ID).To(Equal(core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")))
				Expect(iamServiceRegistrationSupportedRoleModel.Description).To(Equal(iamServiceRegistrationDescriptionObjectModel))
				Expect(iamServiceRegistrationSupportedRoleModel.DisplayName).To(Equal(iamServiceRegistrationDisplayNameObjectModel))
				Expect(iamServiceRegistrationSupportedRoleModel.Options).To(Equal(supportedRoleOptionsModel))

				// Construct an instance of the EnvironmentAttributeOptions model
				environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
				Expect(environmentAttributeOptionsModel).ToNot(BeNil())
				environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)
				Expect(environmentAttributeOptionsModel.Hidden).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the EnvironmentAttribute model
				environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
				Expect(environmentAttributeModel).ToNot(BeNil())
				environmentAttributeModel.Key = core.StringPtr("networkType")
				environmentAttributeModel.Values = []string{"public"}
				environmentAttributeModel.Options = environmentAttributeOptionsModel
				Expect(environmentAttributeModel.Key).To(Equal(core.StringPtr("networkType")))
				Expect(environmentAttributeModel.Values).To(Equal([]string{"public"}))
				Expect(environmentAttributeModel.Options).To(Equal(environmentAttributeOptionsModel))

				// Construct an instance of the IamServiceRegistrationSupportedNetwork model
				iamServiceRegistrationSupportedNetworkModel := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
				Expect(iamServiceRegistrationSupportedNetworkModel).ToNot(BeNil())
				iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}
				Expect(iamServiceRegistrationSupportedNetworkModel.EnvironmentAttributes).To(Equal([]partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}))

				// Construct an instance of the CreateIamRegistrationOptions model
				productID := "testString"
				createIamRegistrationOptionsName := "pet-store"
				createIamRegistrationOptionsModel := partnerCenterSellService.NewCreateIamRegistrationOptions(productID, createIamRegistrationOptionsName)
				createIamRegistrationOptionsModel.SetProductID("testString")
				createIamRegistrationOptionsModel.SetName("pet-store")
				createIamRegistrationOptionsModel.SetEnabled(true)
				createIamRegistrationOptionsModel.SetServiceType("service")
				createIamRegistrationOptionsModel.SetActions([]partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel})
				createIamRegistrationOptionsModel.SetAdditionalPolicyScopes([]string{"pet-store"})
				createIamRegistrationOptionsModel.SetDisplayName(iamServiceRegistrationDisplayNameObjectModel)
				createIamRegistrationOptionsModel.SetParentIds([]string{})
				createIamRegistrationOptionsModel.SetResourceHierarchyAttribute(iamServiceRegistrationResourceHierarchyAttributeModel)
				createIamRegistrationOptionsModel.SetSupportedAnonymousAccesses([]partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel})
				createIamRegistrationOptionsModel.SetSupportedAttributes([]partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel})
				createIamRegistrationOptionsModel.SetSupportedAuthorizationSubjects([]partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel})
				createIamRegistrationOptionsModel.SetSupportedRoles([]partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel})
				createIamRegistrationOptionsModel.SetSupportedNetwork(iamServiceRegistrationSupportedNetworkModel)
				createIamRegistrationOptionsModel.SetEnv("testString")
				createIamRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createIamRegistrationOptionsModel).ToNot(BeNil())
				Expect(createIamRegistrationOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(createIamRegistrationOptionsModel.Name).To(Equal(core.StringPtr("pet-store")))
				Expect(createIamRegistrationOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createIamRegistrationOptionsModel.ServiceType).To(Equal(core.StringPtr("service")))
				Expect(createIamRegistrationOptionsModel.Actions).To(Equal([]partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel}))
				Expect(createIamRegistrationOptionsModel.AdditionalPolicyScopes).To(Equal([]string{"pet-store"}))
				Expect(createIamRegistrationOptionsModel.DisplayName).To(Equal(iamServiceRegistrationDisplayNameObjectModel))
				Expect(createIamRegistrationOptionsModel.ParentIds).To(Equal([]string{}))
				Expect(createIamRegistrationOptionsModel.ResourceHierarchyAttribute).To(Equal(iamServiceRegistrationResourceHierarchyAttributeModel))
				Expect(createIamRegistrationOptionsModel.SupportedAnonymousAccesses).To(Equal([]partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel}))
				Expect(createIamRegistrationOptionsModel.SupportedAttributes).To(Equal([]partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel}))
				Expect(createIamRegistrationOptionsModel.SupportedAuthorizationSubjects).To(Equal([]partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel}))
				Expect(createIamRegistrationOptionsModel.SupportedRoles).To(Equal([]partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel}))
				Expect(createIamRegistrationOptionsModel.SupportedNetwork).To(Equal(iamServiceRegistrationSupportedNetworkModel))
				Expect(createIamRegistrationOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(createIamRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateOnboardingProductOptions successfully`, func() {
				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				Expect(primaryContactModel).ToNot(BeNil())
				primaryContactModel.Name = core.StringPtr("name")
				primaryContactModel.Email = core.StringPtr("name.name@ibm.com")
				Expect(primaryContactModel.Name).To(Equal(core.StringPtr("name")))
				Expect(primaryContactModel.Email).To(Equal(core.StringPtr("name.name@ibm.com")))

				// Construct an instance of the OnboardingProductSupportEscalationContactItems model
				onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
				Expect(onboardingProductSupportEscalationContactItemsModel).ToNot(BeNil())
				onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
				onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")
				Expect(onboardingProductSupportEscalationContactItemsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(onboardingProductSupportEscalationContactItemsModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(onboardingProductSupportEscalationContactItemsModel.Role).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the OnboardingProductSupport model
				onboardingProductSupportModel := new(partnercentersellv1.OnboardingProductSupport)
				Expect(onboardingProductSupportModel).ToNot(BeNil())
				onboardingProductSupportModel.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}
				Expect(onboardingProductSupportModel.EscalationContacts).To(Equal([]partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}))

				// Construct an instance of the CreateOnboardingProductOptions model
				createOnboardingProductOptionsType := "service"
				var createOnboardingProductOptionsPrimaryContact *partnercentersellv1.PrimaryContact = nil
				createOnboardingProductOptionsModel := partnerCenterSellService.NewCreateOnboardingProductOptions(createOnboardingProductOptionsType, createOnboardingProductOptionsPrimaryContact)
				createOnboardingProductOptionsModel.SetType("service")
				createOnboardingProductOptionsModel.SetPrimaryContact(primaryContactModel)
				createOnboardingProductOptionsModel.SetEccnNumber("testString")
				createOnboardingProductOptionsModel.SetEroClass("testString")
				createOnboardingProductOptionsModel.SetUnspsc(float64(72.5))
				createOnboardingProductOptionsModel.SetTaxAssessment("testString")
				createOnboardingProductOptionsModel.SetSupport(onboardingProductSupportModel)
				createOnboardingProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createOnboardingProductOptionsModel).ToNot(BeNil())
				Expect(createOnboardingProductOptionsModel.Type).To(Equal(core.StringPtr("service")))
				Expect(createOnboardingProductOptionsModel.PrimaryContact).To(Equal(primaryContactModel))
				Expect(createOnboardingProductOptionsModel.EccnNumber).To(Equal(core.StringPtr("testString")))
				Expect(createOnboardingProductOptionsModel.EroClass).To(Equal(core.StringPtr("testString")))
				Expect(createOnboardingProductOptionsModel.Unspsc).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(createOnboardingProductOptionsModel.TaxAssessment).To(Equal(core.StringPtr("testString")))
				Expect(createOnboardingProductOptionsModel.Support).To(Equal(onboardingProductSupportModel))
				Expect(createOnboardingProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateRegistrationOptions successfully`, func() {
				// Construct an instance of the PrimaryContact model
				primaryContactModel := new(partnercentersellv1.PrimaryContact)
				Expect(primaryContactModel).ToNot(BeNil())
				primaryContactModel.Name = core.StringPtr("Company Representative")
				primaryContactModel.Email = core.StringPtr("companyrep@email.com")
				Expect(primaryContactModel.Name).To(Equal(core.StringPtr("Company Representative")))
				Expect(primaryContactModel.Email).To(Equal(core.StringPtr("companyrep@email.com")))

				// Construct an instance of the CreateRegistrationOptions model
				createRegistrationOptionsAccountID := "4a5c3c51b97a446fbb1d0e1ef089823b"
				createRegistrationOptionsCompanyName := "Beautiful Company"
				var createRegistrationOptionsPrimaryContact *partnercentersellv1.PrimaryContact = nil
				createRegistrationOptionsModel := partnerCenterSellService.NewCreateRegistrationOptions(createRegistrationOptionsAccountID, createRegistrationOptionsCompanyName, createRegistrationOptionsPrimaryContact)
				createRegistrationOptionsModel.SetAccountID("4a5c3c51b97a446fbb1d0e1ef089823b")
				createRegistrationOptionsModel.SetCompanyName("Beautiful Company")
				createRegistrationOptionsModel.SetPrimaryContact(primaryContactModel)
				createRegistrationOptionsModel.SetDefaultPrivateCatalogID("testString")
				createRegistrationOptionsModel.SetProviderAccessGroup("testString")
				createRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRegistrationOptionsModel).ToNot(BeNil())
				Expect(createRegistrationOptionsModel.AccountID).To(Equal(core.StringPtr("4a5c3c51b97a446fbb1d0e1ef089823b")))
				Expect(createRegistrationOptionsModel.CompanyName).To(Equal(core.StringPtr("Beautiful Company")))
				Expect(createRegistrationOptionsModel.PrimaryContact).To(Equal(primaryContactModel))
				Expect(createRegistrationOptionsModel.DefaultPrivateCatalogID).To(Equal(core.StringPtr("testString")))
				Expect(createRegistrationOptionsModel.ProviderAccessGroup).To(Equal(core.StringPtr("testString")))
				Expect(createRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateResourceBrokerOptions successfully`, func() {
				// Construct an instance of the CreateResourceBrokerOptions model
				createResourceBrokerOptionsAuthScheme := "bearer"
				createResourceBrokerOptionsName := "brokername"
				createResourceBrokerOptionsBrokerURL := "https://broker-url-for-my-service.com"
				createResourceBrokerOptionsType := "provision_through"
				createResourceBrokerOptionsModel := partnerCenterSellService.NewCreateResourceBrokerOptions(createResourceBrokerOptionsAuthScheme, createResourceBrokerOptionsName, createResourceBrokerOptionsBrokerURL, createResourceBrokerOptionsType)
				createResourceBrokerOptionsModel.SetAuthScheme("bearer")
				createResourceBrokerOptionsModel.SetName("brokername")
				createResourceBrokerOptionsModel.SetBrokerURL("https://broker-url-for-my-service.com")
				createResourceBrokerOptionsModel.SetType("provision_through")
				createResourceBrokerOptionsModel.SetAuthUsername("apikey")
				createResourceBrokerOptionsModel.SetAuthPassword("testString")
				createResourceBrokerOptionsModel.SetResourceGroupCrn("crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f")
				createResourceBrokerOptionsModel.SetState("active")
				createResourceBrokerOptionsModel.SetAllowContextUpdates(false)
				createResourceBrokerOptionsModel.SetCatalogType("service")
				createResourceBrokerOptionsModel.SetRegion("global")
				createResourceBrokerOptionsModel.SetEnv("testString")
				createResourceBrokerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceBrokerOptionsModel).ToNot(BeNil())
				Expect(createResourceBrokerOptionsModel.AuthScheme).To(Equal(core.StringPtr("bearer")))
				Expect(createResourceBrokerOptionsModel.Name).To(Equal(core.StringPtr("brokername")))
				Expect(createResourceBrokerOptionsModel.BrokerURL).To(Equal(core.StringPtr("https://broker-url-for-my-service.com")))
				Expect(createResourceBrokerOptionsModel.Type).To(Equal(core.StringPtr("provision_through")))
				Expect(createResourceBrokerOptionsModel.AuthUsername).To(Equal(core.StringPtr("apikey")))
				Expect(createResourceBrokerOptionsModel.AuthPassword).To(Equal(core.StringPtr("testString")))
				Expect(createResourceBrokerOptionsModel.ResourceGroupCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f")))
				Expect(createResourceBrokerOptionsModel.State).To(Equal(core.StringPtr("active")))
				Expect(createResourceBrokerOptionsModel.AllowContextUpdates).To(Equal(core.BoolPtr(false)))
				Expect(createResourceBrokerOptionsModel.CatalogType).To(Equal(core.StringPtr("service")))
				Expect(createResourceBrokerOptionsModel.Region).To(Equal(core.StringPtr("global")))
				Expect(createResourceBrokerOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(createResourceBrokerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCatalogDeploymentOptions successfully`, func() {
				// Construct an instance of the DeleteCatalogDeploymentOptions model
				productID := "testString"
				catalogProductID := "testString"
				catalogPlanID := "testString"
				catalogDeploymentID := "testString"
				deleteCatalogDeploymentOptionsModel := partnerCenterSellService.NewDeleteCatalogDeploymentOptions(productID, catalogProductID, catalogPlanID, catalogDeploymentID)
				deleteCatalogDeploymentOptionsModel.SetProductID("testString")
				deleteCatalogDeploymentOptionsModel.SetCatalogProductID("testString")
				deleteCatalogDeploymentOptionsModel.SetCatalogPlanID("testString")
				deleteCatalogDeploymentOptionsModel.SetCatalogDeploymentID("testString")
				deleteCatalogDeploymentOptionsModel.SetEnv("testString")
				deleteCatalogDeploymentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCatalogDeploymentOptionsModel).ToNot(BeNil())
				Expect(deleteCatalogDeploymentOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogDeploymentOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogDeploymentOptionsModel.CatalogPlanID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogDeploymentOptionsModel.CatalogDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogDeploymentOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogDeploymentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCatalogPlanOptions successfully`, func() {
				// Construct an instance of the DeleteCatalogPlanOptions model
				productID := "testString"
				catalogProductID := "testString"
				catalogPlanID := "testString"
				deleteCatalogPlanOptionsModel := partnerCenterSellService.NewDeleteCatalogPlanOptions(productID, catalogProductID, catalogPlanID)
				deleteCatalogPlanOptionsModel.SetProductID("testString")
				deleteCatalogPlanOptionsModel.SetCatalogProductID("testString")
				deleteCatalogPlanOptionsModel.SetCatalogPlanID("testString")
				deleteCatalogPlanOptionsModel.SetEnv("testString")
				deleteCatalogPlanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCatalogPlanOptionsModel).ToNot(BeNil())
				Expect(deleteCatalogPlanOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogPlanOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogPlanOptionsModel.CatalogPlanID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogPlanOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogPlanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCatalogProductOptions successfully`, func() {
				// Construct an instance of the DeleteCatalogProductOptions model
				productID := "testString"
				catalogProductID := "testString"
				deleteCatalogProductOptionsModel := partnerCenterSellService.NewDeleteCatalogProductOptions(productID, catalogProductID)
				deleteCatalogProductOptionsModel.SetProductID("testString")
				deleteCatalogProductOptionsModel.SetCatalogProductID("testString")
				deleteCatalogProductOptionsModel.SetEnv("testString")
				deleteCatalogProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCatalogProductOptionsModel).ToNot(BeNil())
				Expect(deleteCatalogProductOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogProductOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogProductOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteIamRegistrationOptions successfully`, func() {
				// Construct an instance of the DeleteIamRegistrationOptions model
				productID := "testString"
				programmaticName := "testString"
				deleteIamRegistrationOptionsModel := partnerCenterSellService.NewDeleteIamRegistrationOptions(productID, programmaticName)
				deleteIamRegistrationOptionsModel.SetProductID("testString")
				deleteIamRegistrationOptionsModel.SetProgrammaticName("testString")
				deleteIamRegistrationOptionsModel.SetEnv("testString")
				deleteIamRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteIamRegistrationOptionsModel).ToNot(BeNil())
				Expect(deleteIamRegistrationOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteIamRegistrationOptionsModel.ProgrammaticName).To(Equal(core.StringPtr("testString")))
				Expect(deleteIamRegistrationOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(deleteIamRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteOnboardingProductOptions successfully`, func() {
				// Construct an instance of the DeleteOnboardingProductOptions model
				productID := "testString"
				deleteOnboardingProductOptionsModel := partnerCenterSellService.NewDeleteOnboardingProductOptions(productID)
				deleteOnboardingProductOptionsModel.SetProductID("testString")
				deleteOnboardingProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteOnboardingProductOptionsModel).ToNot(BeNil())
				Expect(deleteOnboardingProductOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOnboardingProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRegistrationOptions successfully`, func() {
				// Construct an instance of the DeleteRegistrationOptions model
				registrationID := "testString"
				deleteRegistrationOptionsModel := partnerCenterSellService.NewDeleteRegistrationOptions(registrationID)
				deleteRegistrationOptionsModel.SetRegistrationID("testString")
				deleteRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRegistrationOptionsModel).ToNot(BeNil())
				Expect(deleteRegistrationOptionsModel.RegistrationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceBrokerOptions successfully`, func() {
				// Construct an instance of the DeleteResourceBrokerOptions model
				brokerID := "testString"
				deleteResourceBrokerOptionsModel := partnerCenterSellService.NewDeleteResourceBrokerOptions(brokerID)
				deleteResourceBrokerOptionsModel.SetBrokerID("testString")
				deleteResourceBrokerOptionsModel.SetEnv("testString")
				deleteResourceBrokerOptionsModel.SetRemoveFromAccount(true)
				deleteResourceBrokerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceBrokerOptionsModel).ToNot(BeNil())
				Expect(deleteResourceBrokerOptionsModel.BrokerID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceBrokerOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceBrokerOptionsModel.RemoveFromAccount).To(Equal(core.BoolPtr(true)))
				Expect(deleteResourceBrokerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogDeploymentOptions successfully`, func() {
				// Construct an instance of the GetCatalogDeploymentOptions model
				productID := "testString"
				catalogProductID := "testString"
				catalogPlanID := "testString"
				catalogDeploymentID := "testString"
				getCatalogDeploymentOptionsModel := partnerCenterSellService.NewGetCatalogDeploymentOptions(productID, catalogProductID, catalogPlanID, catalogDeploymentID)
				getCatalogDeploymentOptionsModel.SetProductID("testString")
				getCatalogDeploymentOptionsModel.SetCatalogProductID("testString")
				getCatalogDeploymentOptionsModel.SetCatalogPlanID("testString")
				getCatalogDeploymentOptionsModel.SetCatalogDeploymentID("testString")
				getCatalogDeploymentOptionsModel.SetEnv("testString")
				getCatalogDeploymentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogDeploymentOptionsModel).ToNot(BeNil())
				Expect(getCatalogDeploymentOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogDeploymentOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogDeploymentOptionsModel.CatalogPlanID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogDeploymentOptionsModel.CatalogDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogDeploymentOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogDeploymentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogPlanOptions successfully`, func() {
				// Construct an instance of the GetCatalogPlanOptions model
				productID := "testString"
				catalogProductID := "testString"
				catalogPlanID := "testString"
				getCatalogPlanOptionsModel := partnerCenterSellService.NewGetCatalogPlanOptions(productID, catalogProductID, catalogPlanID)
				getCatalogPlanOptionsModel.SetProductID("testString")
				getCatalogPlanOptionsModel.SetCatalogProductID("testString")
				getCatalogPlanOptionsModel.SetCatalogPlanID("testString")
				getCatalogPlanOptionsModel.SetEnv("testString")
				getCatalogPlanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogPlanOptionsModel).ToNot(BeNil())
				Expect(getCatalogPlanOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogPlanOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogPlanOptionsModel.CatalogPlanID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogPlanOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogPlanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogProductOptions successfully`, func() {
				// Construct an instance of the GetCatalogProductOptions model
				productID := "testString"
				catalogProductID := "testString"
				getCatalogProductOptionsModel := partnerCenterSellService.NewGetCatalogProductOptions(productID, catalogProductID)
				getCatalogProductOptionsModel.SetProductID("testString")
				getCatalogProductOptionsModel.SetCatalogProductID("testString")
				getCatalogProductOptionsModel.SetEnv("testString")
				getCatalogProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogProductOptionsModel).ToNot(BeNil())
				Expect(getCatalogProductOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogProductOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogProductOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetIamRegistrationOptions successfully`, func() {
				// Construct an instance of the GetIamRegistrationOptions model
				productID := "testString"
				programmaticName := "testString"
				getIamRegistrationOptionsModel := partnerCenterSellService.NewGetIamRegistrationOptions(productID, programmaticName)
				getIamRegistrationOptionsModel.SetProductID("testString")
				getIamRegistrationOptionsModel.SetProgrammaticName("testString")
				getIamRegistrationOptionsModel.SetEnv("testString")
				getIamRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getIamRegistrationOptionsModel).ToNot(BeNil())
				Expect(getIamRegistrationOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(getIamRegistrationOptionsModel.ProgrammaticName).To(Equal(core.StringPtr("testString")))
				Expect(getIamRegistrationOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(getIamRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetOnboardingProductOptions successfully`, func() {
				// Construct an instance of the GetOnboardingProductOptions model
				productID := "testString"
				getOnboardingProductOptionsModel := partnerCenterSellService.NewGetOnboardingProductOptions(productID)
				getOnboardingProductOptionsModel.SetProductID("testString")
				getOnboardingProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getOnboardingProductOptionsModel).ToNot(BeNil())
				Expect(getOnboardingProductOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(getOnboardingProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProductBadgeOptions successfully`, func() {
				// Construct an instance of the GetProductBadgeOptions model
				badgeID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductBadgeOptionsModel := partnerCenterSellService.NewGetProductBadgeOptions(badgeID)
				getProductBadgeOptionsModel.SetBadgeID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getProductBadgeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProductBadgeOptionsModel).ToNot(BeNil())
				Expect(getProductBadgeOptionsModel.BadgeID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getProductBadgeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRegistrationOptions successfully`, func() {
				// Construct an instance of the GetRegistrationOptions model
				registrationID := "testString"
				getRegistrationOptionsModel := partnerCenterSellService.NewGetRegistrationOptions(registrationID)
				getRegistrationOptionsModel.SetRegistrationID("testString")
				getRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRegistrationOptionsModel).ToNot(BeNil())
				Expect(getRegistrationOptionsModel.RegistrationID).To(Equal(core.StringPtr("testString")))
				Expect(getRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceBrokerOptions successfully`, func() {
				// Construct an instance of the GetResourceBrokerOptions model
				brokerID := "testString"
				getResourceBrokerOptionsModel := partnerCenterSellService.NewGetResourceBrokerOptions(brokerID)
				getResourceBrokerOptionsModel.SetBrokerID("testString")
				getResourceBrokerOptionsModel.SetEnv("testString")
				getResourceBrokerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceBrokerOptionsModel).ToNot(BeNil())
				Expect(getResourceBrokerOptionsModel.BrokerID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceBrokerOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(getResourceBrokerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewIamServiceRegistrationSupportedAnonymousAccessAttributes successfully`, func() {
				accountID := "testString"
				serviceName := "testString"
				additionalProperties := map[string]string{"key1": "testString"}
				_model, err := partnerCenterSellService.NewIamServiceRegistrationSupportedAnonymousAccessAttributes(accountID, serviceName, additionalProperties)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListProductBadgesOptions successfully`, func() {
				// Construct an instance of the ListProductBadgesOptions model
				listProductBadgesOptionsModel := partnerCenterSellService.NewListProductBadgesOptions()
				listProductBadgesOptionsModel.SetLimit(int64(100))
				listProductBadgesOptionsModel.SetStart(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				listProductBadgesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProductBadgesOptionsModel).ToNot(BeNil())
				Expect(listProductBadgesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listProductBadgesOptionsModel.Start).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(listProductBadgesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPrimaryContact successfully`, func() {
				name := "testString"
				email := "testString"
				_model, err := partnerCenterSellService.NewPrimaryContact(name, email)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSupportedRoleOptions successfully`, func() {
				accessPolicy := true
				_model, err := partnerCenterSellService.NewSupportedRoleOptions(accessPolicy)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateCatalogDeploymentOptions successfully`, func() {
				// Construct an instance of the UpdateCatalogDeploymentOptions model
				productID := "testString"
				catalogProductID := "testString"
				catalogPlanID := "testString"
				catalogDeploymentID := "testString"
				globalCatalogDeploymentPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateCatalogDeploymentOptionsModel := partnerCenterSellService.NewUpdateCatalogDeploymentOptions(productID, catalogProductID, catalogPlanID, catalogDeploymentID, globalCatalogDeploymentPatch)
				updateCatalogDeploymentOptionsModel.SetProductID("testString")
				updateCatalogDeploymentOptionsModel.SetCatalogProductID("testString")
				updateCatalogDeploymentOptionsModel.SetCatalogPlanID("testString")
				updateCatalogDeploymentOptionsModel.SetCatalogDeploymentID("testString")
				updateCatalogDeploymentOptionsModel.SetGlobalCatalogDeploymentPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateCatalogDeploymentOptionsModel.SetEnv("testString")
				updateCatalogDeploymentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCatalogDeploymentOptionsModel).ToNot(BeNil())
				Expect(updateCatalogDeploymentOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogDeploymentOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogDeploymentOptionsModel.CatalogPlanID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogDeploymentOptionsModel.CatalogDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogDeploymentOptionsModel.GlobalCatalogDeploymentPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateCatalogDeploymentOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogDeploymentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCatalogPlanOptions successfully`, func() {
				// Construct an instance of the UpdateCatalogPlanOptions model
				productID := "testString"
				catalogProductID := "testString"
				catalogPlanID := "testString"
				globalCatalogPlanPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateCatalogPlanOptionsModel := partnerCenterSellService.NewUpdateCatalogPlanOptions(productID, catalogProductID, catalogPlanID, globalCatalogPlanPatch)
				updateCatalogPlanOptionsModel.SetProductID("testString")
				updateCatalogPlanOptionsModel.SetCatalogProductID("testString")
				updateCatalogPlanOptionsModel.SetCatalogPlanID("testString")
				updateCatalogPlanOptionsModel.SetGlobalCatalogPlanPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateCatalogPlanOptionsModel.SetEnv("testString")
				updateCatalogPlanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCatalogPlanOptionsModel).ToNot(BeNil())
				Expect(updateCatalogPlanOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogPlanOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogPlanOptionsModel.CatalogPlanID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogPlanOptionsModel.GlobalCatalogPlanPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateCatalogPlanOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogPlanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCatalogProductOptions successfully`, func() {
				// Construct an instance of the UpdateCatalogProductOptions model
				productID := "testString"
				catalogProductID := "testString"
				globalCatalogProductPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateCatalogProductOptionsModel := partnerCenterSellService.NewUpdateCatalogProductOptions(productID, catalogProductID, globalCatalogProductPatch)
				updateCatalogProductOptionsModel.SetProductID("testString")
				updateCatalogProductOptionsModel.SetCatalogProductID("testString")
				updateCatalogProductOptionsModel.SetGlobalCatalogProductPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateCatalogProductOptionsModel.SetEnv("testString")
				updateCatalogProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCatalogProductOptionsModel).ToNot(BeNil())
				Expect(updateCatalogProductOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogProductOptionsModel.CatalogProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogProductOptionsModel.GlobalCatalogProductPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateCatalogProductOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateIamRegistrationOptions successfully`, func() {
				// Construct an instance of the UpdateIamRegistrationOptions model
				productID := "testString"
				programmaticName := "testString"
				iamRegistrationPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateIamRegistrationOptionsModel := partnerCenterSellService.NewUpdateIamRegistrationOptions(productID, programmaticName, iamRegistrationPatch)
				updateIamRegistrationOptionsModel.SetProductID("testString")
				updateIamRegistrationOptionsModel.SetProgrammaticName("testString")
				updateIamRegistrationOptionsModel.SetIamRegistrationPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateIamRegistrationOptionsModel.SetEnv("testString")
				updateIamRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateIamRegistrationOptionsModel).ToNot(BeNil())
				Expect(updateIamRegistrationOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateIamRegistrationOptionsModel.ProgrammaticName).To(Equal(core.StringPtr("testString")))
				Expect(updateIamRegistrationOptionsModel.IamRegistrationPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateIamRegistrationOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(updateIamRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateOnboardingProductOptions successfully`, func() {
				// Construct an instance of the UpdateOnboardingProductOptions model
				productID := "testString"
				onboardingProductPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateOnboardingProductOptionsModel := partnerCenterSellService.NewUpdateOnboardingProductOptions(productID, onboardingProductPatch)
				updateOnboardingProductOptionsModel.SetProductID("testString")
				updateOnboardingProductOptionsModel.SetOnboardingProductPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateOnboardingProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateOnboardingProductOptionsModel).ToNot(BeNil())
				Expect(updateOnboardingProductOptionsModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateOnboardingProductOptionsModel.OnboardingProductPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateOnboardingProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateRegistrationOptions successfully`, func() {
				// Construct an instance of the UpdateRegistrationOptions model
				registrationID := "testString"
				registrationPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateRegistrationOptionsModel := partnerCenterSellService.NewUpdateRegistrationOptions(registrationID, registrationPatch)
				updateRegistrationOptionsModel.SetRegistrationID("testString")
				updateRegistrationOptionsModel.SetRegistrationPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateRegistrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRegistrationOptionsModel).ToNot(BeNil())
				Expect(updateRegistrationOptionsModel.RegistrationID).To(Equal(core.StringPtr("testString")))
				Expect(updateRegistrationOptionsModel.RegistrationPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateRegistrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceBrokerOptions successfully`, func() {
				// Construct an instance of the UpdateResourceBrokerOptions model
				brokerID := "testString"
				brokerPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateResourceBrokerOptionsModel := partnerCenterSellService.NewUpdateResourceBrokerOptions(brokerID, brokerPatch)
				updateResourceBrokerOptionsModel.SetBrokerID("testString")
				updateResourceBrokerOptionsModel.SetBrokerPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateResourceBrokerOptionsModel.SetEnv("testString")
				updateResourceBrokerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceBrokerOptionsModel).ToNot(BeNil())
				Expect(updateResourceBrokerOptionsModel.BrokerID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceBrokerOptionsModel.BrokerPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateResourceBrokerOptionsModel.Env).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceBrokerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalBrokerPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.BrokerPatch)
			model.AuthUsername = core.StringPtr("apikey")
			model.AuthPassword = core.StringPtr("testString")
			model.AuthScheme = core.StringPtr("bearer")
			model.ResourceGroupCrn = core.StringPtr("testString")
			model.State = core.StringPtr("active")
			model.BrokerURL = core.StringPtr("testString")
			model.AllowContextUpdates = core.BoolPtr(true)
			model.CatalogType = core.StringPtr("testString")
			model.Type = core.StringPtr("provision_through")
			model.Region = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.BrokerPatch
			err = partnercentersellv1.UnmarshalBrokerPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCatalogHighlightItem successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.CatalogHighlightItem)
			model.Description = core.StringPtr("testString")
			model.DescriptionI18n = map[string]string{"key1": "testString"}
			model.Title = core.StringPtr("testString")
			model.TitleI18n = map[string]string{"key1": "testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.CatalogHighlightItem
			err = partnercentersellv1.UnmarshalCatalogHighlightItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCatalogProductMediaItem successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.CatalogProductMediaItem)
			model.Caption = core.StringPtr("testString")
			model.CaptionI18n = map[string]string{"key1": "testString"}
			model.Thumbnail = core.StringPtr("testString")
			model.Type = core.StringPtr("image")
			model.URL = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.CatalogProductMediaItem
			err = partnercentersellv1.UnmarshalCatalogProductMediaItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCatalogProductProvider successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.CatalogProductProvider)
			model.Name = core.StringPtr("testString")
			model.Email = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.CatalogProductProvider
			err = partnercentersellv1.UnmarshalCatalogProductProvider(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalEnvironmentAttribute successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.EnvironmentAttribute)
			model.Key = core.StringPtr("testString")
			model.Values = []string{"testString"}
			model.Options = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.EnvironmentAttribute
			err = partnercentersellv1.UnmarshalEnvironmentAttribute(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalEnvironmentAttributeOptions successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.EnvironmentAttributeOptions)
			model.Hidden = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.EnvironmentAttributeOptions
			err = partnercentersellv1.UnmarshalEnvironmentAttributeOptions(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogDeploymentMetadata successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
			model.RcCompatible = core.BoolPtr(true)
			model.Ui = nil
			model.Service = nil
			model.Deployment = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogDeploymentMetadata
			err = partnercentersellv1.UnmarshalGlobalCatalogDeploymentMetadata(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogDeploymentPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogDeploymentPatch)
			model.Active = core.BoolPtr(true)
			model.Disabled = core.BoolPtr(true)
			model.OverviewUi = nil
			model.Tags = []string{"testString"}
			model.ObjectProvider = nil
			model.Metadata = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogDeploymentPatch
			err = partnercentersellv1.UnmarshalGlobalCatalogDeploymentPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogMetadataDeployment successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
			model.Broker = nil
			model.Location = core.StringPtr("testString")
			model.LocationURL = core.StringPtr("testString")
			model.TargetCrn = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogMetadataDeployment
			err = partnercentersellv1.UnmarshalGlobalCatalogMetadataDeployment(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogMetadataDeploymentBroker successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
			model.Name = core.StringPtr("testString")
			model.Guid = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogMetadataDeploymentBroker
			err = partnercentersellv1.UnmarshalGlobalCatalogMetadataDeploymentBroker(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogMetadataPricing successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogMetadataPricing)
			model.Type = core.StringPtr("free")
			model.Origin = core.StringPtr("global_catalog")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogMetadataPricing
			err = partnercentersellv1.UnmarshalGlobalCatalogMetadataPricing(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogMetadataService successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogMetadataService)
			model.RcProvisionable = core.BoolPtr(true)
			model.IamCompatible = core.BoolPtr(true)
			model.Bindable = core.BoolPtr(true)
			model.PlanUpdateable = core.BoolPtr(true)
			model.ServiceKeySupported = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogMetadataService
			err = partnercentersellv1.UnmarshalGlobalCatalogMetadataService(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogMetadataUI successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogMetadataUI)
			model.Strings = nil
			model.Urls = nil
			model.Hidden = core.BoolPtr(true)
			model.SideBySideIndex = core.Float64Ptr(float64(72.5))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogMetadataUI
			err = partnercentersellv1.UnmarshalGlobalCatalogMetadataUI(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogMetadataUIStrings successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
			model.En = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogMetadataUIStrings
			err = partnercentersellv1.UnmarshalGlobalCatalogMetadataUIStrings(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogMetadataUIStringsContent successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
			model.Bullets = nil
			model.Media = nil
			model.EmbeddableDashboard = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogMetadataUIStringsContent
			err = partnercentersellv1.UnmarshalGlobalCatalogMetadataUIStringsContent(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogMetadataUIUrls successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
			model.DocURL = core.StringPtr("testString")
			model.ApidocsURL = core.StringPtr("testString")
			model.TermsURL = core.StringPtr("testString")
			model.InstructionsURL = core.StringPtr("testString")
			model.CatalogDetailsURL = core.StringPtr("testString")
			model.CustomCreatePageURL = core.StringPtr("testString")
			model.Dashboard = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogMetadataUIUrls
			err = partnercentersellv1.UnmarshalGlobalCatalogMetadataUIUrls(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogOverviewUI successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogOverviewUI)
			model.En = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogOverviewUI
			err = partnercentersellv1.UnmarshalGlobalCatalogOverviewUI(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogOverviewUITranslatedContent successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
			model.DisplayName = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.LongDescription = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogOverviewUITranslatedContent
			err = partnercentersellv1.UnmarshalGlobalCatalogOverviewUITranslatedContent(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogPlanMetadata successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogPlanMetadata)
			model.RcCompatible = core.BoolPtr(true)
			model.Ui = nil
			model.Service = nil
			model.Pricing = nil
			model.Plan = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogPlanMetadata
			err = partnercentersellv1.UnmarshalGlobalCatalogPlanMetadata(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogPlanMetadataPlan successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
			model.AllowInternalUsers = core.BoolPtr(true)
			model.Bindable = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogPlanMetadataPlan
			err = partnercentersellv1.UnmarshalGlobalCatalogPlanMetadataPlan(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogPlanPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogPlanPatch)
			model.Active = core.BoolPtr(true)
			model.Disabled = core.BoolPtr(true)
			model.OverviewUi = nil
			model.Tags = []string{"testString"}
			model.ObjectProvider = nil
			model.Metadata = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogPlanPatch
			err = partnercentersellv1.UnmarshalGlobalCatalogPlanPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogProductImages successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogProductImages)
			model.Image = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogProductImages
			err = partnercentersellv1.UnmarshalGlobalCatalogProductImages(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogProductMetadata successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogProductMetadata)
			model.RcCompatible = core.BoolPtr(true)
			model.Ui = nil
			model.Service = nil
			model.Other = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogProductMetadata
			err = partnercentersellv1.UnmarshalGlobalCatalogProductMetadata(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogProductMetadataOther successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
			model.PC = nil
			model.Composite = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogProductMetadataOther
			err = partnercentersellv1.UnmarshalGlobalCatalogProductMetadataOther(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogProductMetadataOtherComposite successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
			model.CompositeKind = core.StringPtr("service")
			model.CompositeTag = core.StringPtr("testString")
			model.Children = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogProductMetadataOtherComposite
			err = partnercentersellv1.UnmarshalGlobalCatalogProductMetadataOtherComposite(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogProductMetadataOtherCompositeChild successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
			model.Kind = core.StringPtr("service")
			model.Name = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild
			err = partnercentersellv1.UnmarshalGlobalCatalogProductMetadataOtherCompositeChild(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogProductMetadataOtherPC successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
			model.Support = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogProductMetadataOtherPC
			err = partnercentersellv1.UnmarshalGlobalCatalogProductMetadataOtherPC(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogProductMetadataOtherPCSupport successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
			model.URL = core.StringPtr("testString")
			model.StatusURL = core.StringPtr("testString")
			model.Locations = []string{"testString"}
			model.Languages = []string{"testString"}
			model.Process = core.StringPtr("testString")
			model.ProcessI18n = map[string]string{"key1": "testString"}
			model.SupportType = core.StringPtr("community")
			model.SupportEscalation = nil
			model.SupportDetails = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport
			err = partnercentersellv1.UnmarshalGlobalCatalogProductMetadataOtherPCSupport(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGlobalCatalogProductPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.GlobalCatalogProductPatch)
			model.Active = core.BoolPtr(true)
			model.Disabled = core.BoolPtr(true)
			model.OverviewUi = nil
			model.Tags = []string{"testString"}
			model.Images = nil
			model.ObjectProvider = nil
			model.Metadata = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.GlobalCatalogProductPatch
			err = partnercentersellv1.UnmarshalGlobalCatalogProductPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationAction successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationAction)
			model.ID = core.StringPtr("testString")
			model.Roles = []string{"testString"}
			model.Description = nil
			model.DisplayName = nil
			model.Options = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationAction
			err = partnercentersellv1.UnmarshalIamServiceRegistrationAction(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationActionOptions successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationActionOptions)
			model.Hidden = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationActionOptions
			err = partnercentersellv1.UnmarshalIamServiceRegistrationActionOptions(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationDescriptionObject successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
			model.Default = core.StringPtr("testString")
			model.En = core.StringPtr("testString")
			model.De = core.StringPtr("testString")
			model.Es = core.StringPtr("testString")
			model.Fr = core.StringPtr("testString")
			model.It = core.StringPtr("testString")
			model.Ja = core.StringPtr("testString")
			model.Ko = core.StringPtr("testString")
			model.PtBr = core.StringPtr("testString")
			model.ZhTw = core.StringPtr("testString")
			model.ZhCn = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationDescriptionObject
			err = partnercentersellv1.UnmarshalIamServiceRegistrationDescriptionObject(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationDisplayNameObject successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
			model.Default = core.StringPtr("testString")
			model.En = core.StringPtr("testString")
			model.De = core.StringPtr("testString")
			model.Es = core.StringPtr("testString")
			model.Fr = core.StringPtr("testString")
			model.It = core.StringPtr("testString")
			model.Ja = core.StringPtr("testString")
			model.Ko = core.StringPtr("testString")
			model.PtBr = core.StringPtr("testString")
			model.ZhTw = core.StringPtr("testString")
			model.ZhCn = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationDisplayNameObject
			err = partnercentersellv1.UnmarshalIamServiceRegistrationDisplayNameObject(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationPatch)
			model.Enabled = core.BoolPtr(true)
			model.ServiceType = core.StringPtr("service")
			model.Actions = nil
			model.AdditionalPolicyScopes = []string{"testString"}
			model.DisplayName = nil
			model.ParentIds = []string{"testString"}
			model.ResourceHierarchyAttribute = nil
			model.SupportedAnonymousAccesses = nil
			model.SupportedAttributes = nil
			model.SupportedAuthorizationSubjects = nil
			model.SupportedRoles = nil
			model.SupportedNetwork = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationPatch
			err = partnercentersellv1.UnmarshalIamServiceRegistrationPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationResourceHierarchyAttribute successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
			model.Key = core.StringPtr("testString")
			model.Value = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute
			err = partnercentersellv1.UnmarshalIamServiceRegistrationResourceHierarchyAttribute(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationSupportedAnonymousAccess successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
			model.Attributes = nil
			model.Roles = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess
			err = partnercentersellv1.UnmarshalIamServiceRegistrationSupportedAnonymousAccess(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationSupportedAnonymousAccessAttributes successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
			model.AccountID = core.StringPtr("testString")
			model.ServiceName = core.StringPtr("testString")
			model.AdditionalProperties = map[string]string{"key1": "testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes
			err = partnercentersellv1.UnmarshalIamServiceRegistrationSupportedAnonymousAccessAttributes(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationSupportedAttribute successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
			model.Key = core.StringPtr("testString")
			model.Options = nil
			model.DisplayName = nil
			model.Description = nil
			model.Ui = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationSupportedAttribute
			err = partnercentersellv1.UnmarshalIamServiceRegistrationSupportedAttribute(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationSupportedAuthorizationSubject successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
			model.Attributes = nil
			model.Roles = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject
			err = partnercentersellv1.UnmarshalIamServiceRegistrationSupportedAuthorizationSubject(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationSupportedNetwork successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
			model.EnvironmentAttributes = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationSupportedNetwork
			err = partnercentersellv1.UnmarshalIamServiceRegistrationSupportedNetwork(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIamServiceRegistrationSupportedRole successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
			model.ID = core.StringPtr("testString")
			model.Description = nil
			model.DisplayName = nil
			model.Options = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.IamServiceRegistrationSupportedRole
			err = partnercentersellv1.UnmarshalIamServiceRegistrationSupportedRole(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalOnboardingProductPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.OnboardingProductPatch)
			model.PrimaryContact = nil
			model.EccnNumber = core.StringPtr("testString")
			model.EroClass = core.StringPtr("testString")
			model.Unspsc = core.Float64Ptr(float64(72.5))
			model.TaxAssessment = core.StringPtr("testString")
			model.Support = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.OnboardingProductPatch
			err = partnercentersellv1.UnmarshalOnboardingProductPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalOnboardingProductSupport successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.OnboardingProductSupport)
			model.EscalationContacts = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.OnboardingProductSupport
			err = partnercentersellv1.UnmarshalOnboardingProductSupport(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalOnboardingProductSupportEscalationContactItems successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
			model.Name = core.StringPtr("testString")
			model.Email = core.StringPtr("testString")
			model.Role = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.OnboardingProductSupportEscalationContactItems
			err = partnercentersellv1.UnmarshalOnboardingProductSupportEscalationContactItems(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPrimaryContact successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.PrimaryContact)
			model.Name = core.StringPtr("testString")
			model.Email = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.PrimaryContact
			err = partnercentersellv1.UnmarshalPrimaryContact(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRegistrationPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.RegistrationPatch)
			model.CompanyName = core.StringPtr("testString")
			model.PrimaryContact = nil
			model.DefaultPrivateCatalogID = core.StringPtr("testString")
			model.ProviderAccessGroup = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.RegistrationPatch
			err = partnercentersellv1.UnmarshalRegistrationPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportAuthorizationSubjectAttribute successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
			model.ServiceName = core.StringPtr("testString")
			model.ResourceType = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportAuthorizationSubjectAttribute
			err = partnercentersellv1.UnmarshalSupportAuthorizationSubjectAttribute(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportDetailsItem successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportDetailsItem)
			model.Type = core.StringPtr("support_site")
			model.Contact = core.StringPtr("testString")
			model.ResponseWaitTime = nil
			model.Availability = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportDetailsItem
			err = partnercentersellv1.UnmarshalSupportDetailsItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportDetailsItemAvailability successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportDetailsItemAvailability)
			model.Times = nil
			model.Timezone = core.StringPtr("testString")
			model.AlwaysAvailable = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportDetailsItemAvailability
			err = partnercentersellv1.UnmarshalSupportDetailsItemAvailability(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportDetailsItemAvailabilityTime successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
			model.Day = core.Float64Ptr(float64(72.5))
			model.StartTime = core.StringPtr("testString")
			model.EndTime = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportDetailsItemAvailabilityTime
			err = partnercentersellv1.UnmarshalSupportDetailsItemAvailabilityTime(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportEscalation successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportEscalation)
			model.Contact = core.StringPtr("testString")
			model.EscalationWaitTime = nil
			model.ResponseWaitTime = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportEscalation
			err = partnercentersellv1.UnmarshalSupportEscalation(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportTimeInterval successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportTimeInterval)
			model.Value = core.Float64Ptr(float64(72.5))
			model.Type = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportTimeInterval
			err = partnercentersellv1.UnmarshalSupportTimeInterval(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedAttributeUi successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedAttributeUi)
			model.InputType = core.StringPtr("testString")
			model.InputDetails = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedAttributeUi
			err = partnercentersellv1.UnmarshalSupportedAttributeUi(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedAttributeUiInputDetails successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedAttributeUiInputDetails)
			model.Type = core.StringPtr("testString")
			model.Values = nil
			model.Gst = nil
			model.URL = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedAttributeUiInputDetails
			err = partnercentersellv1.UnmarshalSupportedAttributeUiInputDetails(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedAttributeUiInputGst successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedAttributeUiInputGst)
			model.Query = core.StringPtr("testString")
			model.ValuePropertyName = core.StringPtr("testString")
			model.LabelPropertyName = core.StringPtr("testString")
			model.InputOptionLabel = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedAttributeUiInputGst
			err = partnercentersellv1.UnmarshalSupportedAttributeUiInputGst(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedAttributeUiInputURL successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedAttributeUiInputURL)
			model.UrlEndpoint = core.StringPtr("testString")
			model.InputOptionLabel = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedAttributeUiInputURL
			err = partnercentersellv1.UnmarshalSupportedAttributeUiInputURL(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedAttributeUiInputValue successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedAttributeUiInputValue)
			model.Value = core.StringPtr("testString")
			model.DisplayName = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedAttributeUiInputValue
			err = partnercentersellv1.UnmarshalSupportedAttributeUiInputValue(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedAttributesOptions successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedAttributesOptions)
			model.Operators = []string{"stringEquals"}
			model.Hidden = core.BoolPtr(true)
			model.SupportedAttributes = []string{"testString"}
			model.PolicyTypes = []string{"access"}
			model.IsEmptyValueSupported = core.BoolPtr(true)
			model.IsStringExistsFalseValueSupported = core.BoolPtr(true)
			model.Key = core.StringPtr("testString")
			model.ResourceHierarchy = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedAttributesOptions
			err = partnercentersellv1.UnmarshalSupportedAttributesOptions(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedAttributesOptionsResourceHierarchy successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
			model.Key = nil
			model.Value = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedAttributesOptionsResourceHierarchy
			err = partnercentersellv1.UnmarshalSupportedAttributesOptionsResourceHierarchy(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedAttributesOptionsResourceHierarchyKey successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
			model.Key = core.StringPtr("testString")
			model.Value = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey
			err = partnercentersellv1.UnmarshalSupportedAttributesOptionsResourceHierarchyKey(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedAttributesOptionsResourceHierarchyValue successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
			model.Key = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue
			err = partnercentersellv1.UnmarshalSupportedAttributesOptionsResourceHierarchyValue(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSupportedRoleOptions successfully`, func() {
			// Construct an instance of the model.
			model := new(partnercentersellv1.SupportedRoleOptions)
			model.AccessPolicy = core.BoolPtr(true)
			model.AdditionalPropertiesForAccessPolicy = map[string]string{"key1": "testString"}
			model.PolicyType = []string{"access"}
			model.AccountType = core.StringPtr("enterprise")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *partnercentersellv1.SupportedRoleOptions
			err = partnercentersellv1.UnmarshalSupportedRoleOptions(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}