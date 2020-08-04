package feato_test

import (
	"testing"

	"github.com/imantung/feato"
	"github.com/stretchr/testify/require"
)

func TestRegisterGlobal(t *testing.T) {
	feato.RegisterGlobal([]*feato.Feature{
		{Name: "name01", Flag: feato.Enabled},
		{Name: "name02", Flag: feato.Disabled},
		{Name: "group01", Childs: []*feato.Feature{
			{Name: "name03", Flag: feato.Enabled},
			{Name: "name04", Flag: feato.Disabled},
			{Name: "sub01", Flag: feato.Enabled, Childs: []*feato.Feature{
				{Name: "name05", Flag: feato.Enabled},
				{Name: "name06", Flag: feato.Disabled},
			}},
		}},
	})
	defer feato.Instance.ClearAll()

	require.Equal(t, feato.FlagMap{
		"name01":               feato.Enabled,
		"name02":               feato.Disabled,
		"group01.name03":       feato.Enabled,
		"group01.name04":       feato.Disabled,
		"group01.sub01":        feato.Enabled,
		"group01.sub01.name05": feato.Enabled,
		"group01.sub01.name06": feato.Disabled,
	}, feato.Instance)

	require.True(t, feato.IsEnabled("name01"))
	require.False(t, feato.IsEnabled("group01.sub01.name06"))
	require.False(t, feato.IsEnabled("not-found"))

}
