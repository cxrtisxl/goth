package goth_test

import (
	"reflect"
	"testing"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/azureadv2"
	"github.com/markbates/goth/providers/faux"
	"github.com/stretchr/testify/assert"
)

func Test_ConfigProvider(t *testing.T) {
	a := assert.New(t)
	providerCfg := azureadv2.Config{
		ClientKey: "6731de76-14a6-49ae-97bc-6eba6914391e",
		Secret:    "foo",
	}
	callbackURL := "http://localhost:3000"
	providerCfg.CallbackURL = callbackURL
	provider := providerCfg.Build()
	goth.UseProviders(provider)

	a.Equal(len(goth.GetProviders()), 1)
	a.Equal(goth.GetProviders()[provider.Name()], provider)

	// Gets provider, reflect the underlying azureadv2.Provider and
	// compares it's callbackURL with the one that was set in azureadv2.Config
	v := reflect.ValueOf(goth.GetProviders()[provider.Name()]).Elem()
	a.Equal(v.FieldByName("CallbackURL").String(), callbackURL)

	goth.ClearProviders()
}

func Test_UseProviders(t *testing.T) {
	a := assert.New(t)

	provider := &faux.Provider{}
	goth.UseProviders(provider)
	a.Equal(len(goth.GetProviders()), 1)
	a.Equal(goth.GetProviders()[provider.Name()], provider)
	goth.ClearProviders()
}

func Test_GetProvider(t *testing.T) {
	a := assert.New(t)

	provider := &faux.Provider{}
	goth.UseProviders(provider)

	p, err := goth.GetProvider(provider.Name())
	a.NoError(err)
	a.Equal(p, provider)

	_, err = goth.GetProvider("unknown")
	a.Error(err)
	a.Equal(err.Error(), "no provider for unknown exists")
	goth.ClearProviders()
}
