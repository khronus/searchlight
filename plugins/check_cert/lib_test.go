package check_cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	. "github.com/appscode/searchlight/data"
	"github.com/appscode/searchlight/pkg/icinga"
	"github.com/appscode/searchlight/plugins"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/util/cert"
)

func generateCertificate(expirity time.Duration) ([]byte, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"AppsCode Inc."},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(expirity),
	}

	certByte, err := x509.CreateCertificate(rand.Reader, &template, &template, priv.Public(), priv)
	if err != nil {
		return nil, err
	}

	certificate, err := x509.ParseCertificate(certByte)
	if err != nil {
		return nil, err
	}
	return cert.EncodeCertPEM(certificate), nil
}

var _ = Describe("check_cert", func() {
	var secret, secret2 *core.Secret
	var client corev1.SecretInterface
	var plugin plugins.PluginInterface
	var opts options

	BeforeEach(func() {
		secret = &core.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "secret-1",
				Namespace: "demo",
				Labels:    make(map[string]string),
			},
			Data: make(map[string][]byte),
		}
	})

	AfterEach(func() {
		if client != nil {
			client.Delete(secret.Name, &metav1.DeleteOptions{})
		}
	})

	Describe("when there is a valid certificate", func() {
		JustBeforeEach(func() {
			client = client.CoreV1().Secrets(secret.Namespace)
			opts = options{
				Namespace:  secret.Namespace,
				SecretName: secret.Name,
				Warning:    time.Hour * 7 * 24,
				Critical:   time.Hour * 5 * 24,
			}
			plugin = NewPlugin(client, opts)
		})
		Context("with nearly expire", func() {
			It("with in 4 days", func() {
				cert, err := generateCertificate(time.Hour * 4 * 24)
				Expect(err).ShouldNot(HaveOccurred())
				secret.Data["client.cert"] = cert
				_, err = client.Create(secret)
				Expect(err).ShouldNot(HaveOccurred())
				opts.SecretKey = []string{"client.cert"}

				state, _ := plugin.Check()
				Expect(state).Should(BeIdenticalTo(icinga.Critical))
			})
			It("with in 6 days", func() {
				cert, err := generateCertificate(time.Hour * 6 * 24)
				Expect(err).ShouldNot(HaveOccurred())
				secret.Data["client.cert"] = cert
				_, err = client.Create(secret)
				Expect(err).ShouldNot(HaveOccurred())
				opts.SecretKey = []string{"client.cert"}

				state, _ := plugin.Check()
				Expect(state).Should(BeIdenticalTo(icinga.Warning))
			})
			It("with in 10 days", func() {
				cert, err := generateCertificate(time.Hour * 10 * 24)
				Expect(err).ShouldNot(HaveOccurred())
				secret.Data["client.cert"] = cert
				_, err = client.Create(secret)
				Expect(err).ShouldNot(HaveOccurred())
				opts.SecretKey = []string{"client.cert"}

				state, _ := plugin.Check()
				Expect(state).Should(BeIdenticalTo(icinga.OK))
			})
		})

		Context("with label", func() {
			JustBeforeEach(func() {
				secret.Labels["app/searchlight"] = "cert"
				opts.SecretName = ""
				opts.Selector = labels.SelectorFromSet(secret.Labels).String()
			})
			It("with in 10 days", func() {
				cert, err := generateCertificate(time.Hour * 10 * 24)
				Expect(err).ShouldNot(HaveOccurred())
				secret.Data["client.cert"] = cert
				_, err = client.Create(secret)
				Expect(err).ShouldNot(HaveOccurred())
				opts.SecretKey = []string{"client.cert"}

				state, _ := plugin.Check()
				Expect(state).Should(BeIdenticalTo(icinga.OK))
			})

			Context("for multiple secret", func() {
				JustBeforeEach(func() {
					secret.Labels["app/searchlight"] = "cert"
					opts.SecretName = ""
					opts.Selector = labels.SelectorFromSet(secret.Labels).String()

					// Another Secret
					secret2 = &core.Secret{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "secret-2",
							Namespace: "demo",
							Labels: map[string]string{
								"app/searchlight": "cert",
							},
						},
						Data: make(map[string][]byte),
					}
				})
				AfterEach(func() {
					client.Delete(secret2.Name, &metav1.DeleteOptions{})
				})
				It("both contain valid certificate", func() {
					cert, err := generateCertificate(time.Hour * 10 * 24)
					Expect(err).ShouldNot(HaveOccurred())
					secret.Data["client.cert"] = cert
					_, err = client.Create(secret)
					Expect(err).ShouldNot(HaveOccurred())

					cert, err = generateCertificate(time.Hour * 10 * 24)
					Expect(err).ShouldNot(HaveOccurred())
					secret2.Data["client.cert"] = cert
					_, err = client.Create(secret2)
					Expect(err).ShouldNot(HaveOccurred())

					opts.SecretKey = []string{"client.cert"}

					state, _ := plugin.Check()
					Expect(state).Should(BeIdenticalTo(icinga.OK))
				})

				It("one contains nearly expire certificate", func() {
					cert, err := generateCertificate(time.Hour * 10 * 24)
					Expect(err).ShouldNot(HaveOccurred())
					secret.Data["client.cert"] = cert
					_, err = client.Create(secret)
					Expect(err).ShouldNot(HaveOccurred())

					cert, err = generateCertificate(time.Hour * 2 * 24)
					Expect(err).ShouldNot(HaveOccurred())
					secret2.Data["client.cert"] = cert
					_, err = client.Create(secret2)
					Expect(err).ShouldNot(HaveOccurred())

					opts.SecretKey = []string{"client.cert"}

					state, _ := plugin.Check()
					Expect(state).Should(BeIdenticalTo(icinga.Critical))
				})
			})
		})
	})

	Describe("when there is an invalid secret", func() {
		JustBeforeEach(func() {
			client = client.CoreV1().Secrets(secret.Namespace)
			opts = options{
				Namespace:  secret.Namespace,
				SecretName: secret.Name,
			}
			plugin = NewPlugin(client, opts)
		})
		Context("with invalid certificate", func() {
			It("should be Unknown", func() {
				cert, err := generateCertificate(time.Hour * 10 * 24)
				cert[0] = cert[0] + 1
				Expect(err).ShouldNot(HaveOccurred())
				secret.Data["client.cert"] = cert
				_, err = client.Create(secret)
				Expect(err).ShouldNot(HaveOccurred())
				opts.SecretKey = []string{"client.cert"}

				state, _ := plugin.Check()
				Expect(state).Should(BeIdenticalTo(icinga.Unknown))
			})
		})
		Context("with invalid secret key", func() {
			It("should be Warning", func() {
				_, err := client.Create(secret)
				Expect(err).ShouldNot(HaveOccurred())
				opts.SecretKey = []string{"unknown.cert"}

				state, _ := plugin.Check()
				Expect(state).Should(BeIdenticalTo(icinga.Warning))
			})
		})
	})

	Describe("when there is a tls secret", func() {
		JustBeforeEach(func() {
			client = client.CoreV1().Secrets(secret.Namespace)
			secret.Type = core.SecretTypeTLS
			opts = options{
				Namespace:  secret.Namespace,
				SecretName: secret.Name,
			}
			plugin = NewPlugin(client, opts)
		})
		Context("with valid certificate", func() {
			It("with key", func() {
				cert, err := generateCertificate(time.Hour * 10 * 24)
				Expect(err).ShouldNot(HaveOccurred())
				secret.Data["tls.crt"] = cert
				_, err = client.Create(secret)
				Expect(err).ShouldNot(HaveOccurred())
				opts.SecretKey = []string{"tls.crt"}

				state, _ := plugin.Check()
				Expect(state).Should(BeIdenticalTo(icinga.OK))
			})
			It("without key", func() {
				cert, err := generateCertificate(time.Hour * 10 * 24)
				Expect(err).ShouldNot(HaveOccurred())
				secret.Data["tls.crt"] = cert
				_, err = client.Create(secret)
				Expect(err).ShouldNot(HaveOccurred())

				state, _ := plugin.Check()
				Expect(state).Should(BeIdenticalTo(icinga.OK))
			})
			It("with invalid key", func() {
				cert, err := generateCertificate(time.Hour * 10 * 24)
				Expect(err).ShouldNot(HaveOccurred())
				secret.Data["tls-invalid.crt"] = cert
				_, err = client.Create(secret)
				Expect(err).ShouldNot(HaveOccurred())

				state, _ := plugin.Check()
				Expect(state).Should(BeIdenticalTo(icinga.Warning))
			})
		})
	})

	Describe("Check bindata support", func() {
		Context("bindata contain plugin info", func() {
			It("should be succeeded", func() {
				ic, err := LoadClusterChecks()
				Expect(err).ShouldNot(HaveOccurred())
				found := false
				for _, c := range ic.Command {
					if c.Name == "cert" {
						found = true
						break
					}
				}
				Expect(found).Should(BeTrue())
			})
		})
	})
})
