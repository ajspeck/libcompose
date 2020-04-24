package ctx

import (
	"github.com/ajspeck/libcompose/docker/auth"
	"github.com/ajspeck/libcompose/docker/client"
	"github.com/ajspeck/libcompose/project"
	cliconfig "github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/configfile"
)

// Context holds context meta information about a libcompose project and docker
// client information (like configuration file, builder to use, …)
type Context struct {
	project.Context
	ClientFactory client.Factory
	ConfigDir     string
	ConfigFile    *configfile.ConfigFile
	AuthLookup    auth.Lookup
}

// LookupConfig tries to load the docker configuration files, if any.
func (c *Context) LookupConfig() error {
	if c.ConfigFile != nil {
		return nil
	}

	config, err := cliconfig.Load(c.ConfigDir)
	if err != nil {
		return err
	}

	c.ConfigFile = config

	return nil
}
