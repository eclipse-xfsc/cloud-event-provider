package cloudeventprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptionalConfig(t *testing.T) {
	// optional config timeoutInSec is missing
	t.Setenv("CLOUDEVENTPROVIDER_MESSAGING_PROTOCOL", "nats")
	t.Setenv("CLOUDEVENTPROVIDER_MESSAGING_NATS_URL", "http://localhost:4222")
	t.Setenv("CLOUDEVENTPROVIDER_MESSAGING_NATS_SUBJECT", "events")
	t.Setenv("CLOUDEVENTPROVIDER_MESSAGING_NATS_QUEUEGROUP", "test")

	conf, err := loadConfig()
	require.NoError(t, err)

	assert.Equal(t, conf.Protocol, ProtocolTypeNats)
	assert.Implements(t, (*protocolConfig)(nil), conf.Settings)

}

func TestMissingConfig(t *testing.T) {
	// subject config is missing
	t.Setenv("CLOUDEVENTPROVIDER_MESSAGING_PROTOCOL", "nats")
	t.Setenv("CLOUDEVENTPROVIDER_MESSAGING_NATS_URL", "http://localhost:4222")

	conf, err := loadConfig()

	require.Error(t, err)
	assert.Nil(t, conf)
	assert.ErrorIs(t, err, ErrConfigKeyMissing)
}
