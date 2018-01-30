package main

import (
	"errors"

	plugin "github.com/ipfs/go-ipfs/plugin"
	ipfsConfig "github.com/ipfs/go-ipfs/repo/config"
	serialize "github.com/ipfs/go-ipfs/repo/fsrepo/serialize"

	jaegerConfig "github.com/uber/jaeger-client-go/config"
	opentracing "gx/ipfs/QmWLWmRVSiagqP15jczsGME1qpob6HDbtbHAY2he9W5iUo/opentracing-go"
)

// Plugins is exported list of plugins that will be loaded
var Plugins = []plugin.Plugin{
	&jaegerPlugin{},
}

type jaegerPlugin struct{}

var _ plugin.PluginTracer = (*jaegerPlugin)(nil)

var tracerName = "#ERRNAME"

func (*jaegerPlugin) Name() string {
	return "jaeger"
}

func (*jaegerPlugin) Version() string {
	return "0.0.1"
}

// Init will attempt to read the ipfs PeerID and store its value
// for use in the InitTracer method
func (*jaegerPlugin) Init(config string) error {
	if len(config) == 0 {
		return errors.New("No config directory given for Jaeger tracer")
	}
	configFilename, err := ipfsConfig.Filename(config)
	if err != nil {
		return err
	}
	cfg, err := serialize.Load(configFilename)
	if err != nil {
		return err
	}
	tracerName = cfg.Identity.PeerID
	return nil
}

//Initalize a Jaeger tracer and set it as the global tracer in opentracing api
func (*jaegerPlugin) InitTracer() (opentracing.Tracer, error) {
	tracerCfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans: true,
		},
	}
	//we are ignoring the closer for now
	tracer, _, err := tracerCfg.New(tracerName)
	if err != nil {
		return nil, err
	}
	return tracer, nil
}
