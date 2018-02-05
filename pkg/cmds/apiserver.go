package cmds

import (
	"fmt"
	"io"
	"net"

	"github.com/appscode/searchlight/pkg/admission/plugin"
	"github.com/appscode/searchlight/pkg/apiserver"
	"github.com/spf13/cobra"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

const defaultEtcdPathPrefix = "/registry/monitoring.appscode.com"

type APIServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions

	AdmissionHooks []apiserver.AdmissionHook

	StdOut io.Writer
	StdErr io.Writer
}

func NewAPIServerOptions(out, errOut io.Writer, admissionHooks ...apiserver.AdmissionHook) *APIServerOptions {
	o := &APIServerOptions{
		// TODO we will nil out the etcd storage options.  This requires a later level of k8s.io/apiserver
		RecommendedOptions: genericoptions.NewRecommendedOptions(defaultEtcdPathPrefix, apiserver.Codecs.LegacyCodec(admissionv1beta1.SchemeGroupVersion)),

		AdmissionHooks: admissionHooks,

		StdOut: out,
		StdErr: errOut,
	}
	o.RecommendedOptions.Etcd = nil

	return o
}

// NewCommandStartMaster provides a CLI handler for 'start master' command
func NewCommandStartAPIServer(out, errOut io.Writer, stopCh <-chan struct{}) *cobra.Command {
	o := NewAPIServerOptions(out, errOut)

	cmd := &cobra.Command{
		Use:   "apiserver",
		Short: "Launch Searchlight API server",
		Long:  "Launch Searchlight API server",
		RunE: func(c *cobra.Command, args []string) error {
			o.AdmissionHooks = []apiserver.AdmissionHook{
				&plugin.AdmissionHook{},
			}

			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.RecommendedOptions.AddFlags(flags)

	return cmd
}

func (o APIServerOptions) Validate(args []string) error {
	return nil
}

func (o *APIServerOptions) Complete() error {
	return nil
}

func (o APIServerOptions) Config() (*apiserver.Config, error) {
	// TODO have a "real" external address
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewRecommendedConfig(apiserver.Codecs)
	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig: apiserver.ExtraConfig{
			AdmissionHooks: o.AdmissionHooks,
		},
	}
	return config, nil
}

func (o APIServerOptions) RunServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}
	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}
