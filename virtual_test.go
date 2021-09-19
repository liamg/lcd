package lcd

import (
	"testing"

	"github.com/liamg/lcd/virtual"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_4BitInitialisation(t *testing.T) {

	device := virtual.New(16, 2)

	dataBus := device.DataPins()

	_, err := New1602(
		device.RegisterSelectPin(),
		device.EnablePin(),
		dataBus[4],
		dataBus[5],
		dataBus[6],
		dataBus[7],
	)
	require.NoError(t, err)

	assert.True(t, device.IsInitialised())

}
