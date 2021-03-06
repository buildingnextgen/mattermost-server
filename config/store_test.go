package config_test

import (
	"fmt"
	"testing"

	"github.com/mattermost/mattermost-server/config"
	"github.com/stretchr/testify/require"
)

func TestNewStore(t *testing.T) {
	sqlSettings := mainHelper.GetSqlSettings()

	t.Run("database dsn", func(t *testing.T) {
		ds, err := config.NewStore(fmt.Sprintf("%s://%s", *sqlSettings.DriverName, *sqlSettings.DataSource), false)
		require.NoError(t, err)
		ds.Close()
	})

	t.Run("database dsn, watch ignored", func(t *testing.T) {
		ds, err := config.NewStore(fmt.Sprintf("%s://%s", *sqlSettings.DriverName, *sqlSettings.DataSource), true)
		require.NoError(t, err)
		ds.Close()
	})

	t.Run("file dsn", func(t *testing.T) {
		fs, err := config.NewStore("config.json", false)
		require.NoError(t, err)
		fs.Close()
	})

	t.Run("file dsn, watch", func(t *testing.T) {
		fs, err := config.NewStore("config.json", true)
		require.NoError(t, err)
		fs.Close()
	})
}
