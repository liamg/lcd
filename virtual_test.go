package lcd

import (
	"fmt"
	"testing"

	"github.com/liamg/lcd/virtual"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Initialisation(t *testing.T) {

	tests := []struct {
		cols     uint8
		rows     uint8
		fontSize FontSize
	}{
		{
			cols:     16,
			rows:     2,
			fontSize: FontSize5x8,
		},
		{
			cols:     16,
			rows:     2,
			fontSize: FontSize5x11,
		},
		{
			cols:     16,
			rows:     1,
			fontSize: FontSize5x8,
		},
		{
			cols:     8,
			rows:     1,
			fontSize: FontSize5x11,
		},
		{
			cols:     8,
			rows:     2,
			fontSize: FontSize5x8,
		},
		{
			cols:     16,
			rows:     4,
			fontSize: FontSize5x8,
		},
		{
			cols:     20,
			rows:     2,
			fontSize: FontSize5x8,
		},
		{
			cols:     20,
			rows:     4,
			fontSize: FontSize5x11,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("4-bit %dx%d %s", test.cols, test.rows, test.fontSize), func(t *testing.T) {

			device := virtual.New(test.fontSize.Width(), test.fontSize.Height(), test.cols, test.rows)

			dataBus := device.DataPins()

			_, err := New(
				test.cols,
				test.rows,
				test.fontSize,
				device.RegisterSelectPin(),
				device.EnablePin(),
				nil,
				dataBus[4],
				dataBus[5],
				dataBus[6],
				dataBus[7],
			)
			require.NoError(t, err)

			assert.True(t, device.IsInitialised())
			assert.Equal(t, uint8(4), device.BusSize())
			w, h := device.FontSize()
			assert.Equal(t, test.fontSize.Width(), w)
			assert.Equal(t, test.fontSize.Height(), h)
			if test.rows > 1 {
				assert.Equal(t, uint8(2), device.MemoryLines())
			} else {
				assert.Equal(t, uint8(1), device.MemoryLines())
			}
		})
		t.Run(fmt.Sprintf("8-bit %dx%d %s", test.cols, test.rows, test.fontSize), func(t *testing.T) {

			device := virtual.New(test.fontSize.Width(), test.fontSize.Height(), test.cols, test.rows)

			dataBus := device.DataPins()

			_, err := New(
				test.cols,
				test.rows,
				test.fontSize,
				device.RegisterSelectPin(),
				device.EnablePin(),
				nil,
				dataBus[0],
				dataBus[1],
				dataBus[2],
				dataBus[3],
				dataBus[4],
				dataBus[5],
				dataBus[6],
				dataBus[7],
			)
			require.NoError(t, err)

			assert.True(t, device.IsInitialised())
			assert.Equal(t, uint8(8), device.BusSize())
			w, h := device.FontSize()
			assert.Equal(t, test.fontSize.Width(), w)
			assert.Equal(t, test.fontSize.Height(), h)
			if test.rows > 1 {
				assert.Equal(t, uint8(2), device.MemoryLines())
			} else {
				assert.Equal(t, uint8(1), device.MemoryLines())
			}
		})

	}
}

func Test_BasicWrite(t *testing.T) {

	device := virtual.New(5, 8, 16, 2)

	dataBus := device.DataPins()

	lcd1602, err := New(
		16,
		2,
		FontSize5x8,
		device.RegisterSelectPin(),
		device.EnablePin(),
		nil,
		dataBus[0],
		dataBus[1],
		dataBus[2],
		dataBus[3],
		dataBus[4],
		dataBus[5],
		dataBus[6],
		dataBus[7],
	)
	require.NoError(t, err)

	lcd1602.WriteLine(0, "Hello...")
	lcd1602.WriteLine(1, "       ...world!")

	assert.Equal(t, `Hello...        
       ...world!`, device.String())
}

func Test_WriteWithRWPin(t *testing.T) {

	device := virtual.New(5, 8, 16, 2)

	dataBus := device.DataPins()

	lcd1602, err := New(
		16,
		2,
		FontSize5x8,
		device.RegisterSelectPin(),
		device.EnablePin(),
		device.ReadWritePin(),
		dataBus[0],
		dataBus[1],
		dataBus[2],
		dataBus[3],
		dataBus[4],
		dataBus[5],
		dataBus[6],
		dataBus[7],
	)
	require.NoError(t, err)

	lcd1602.WriteLine(0, "Hello...")
	lcd1602.WriteLine(1, "       ...world!")

	assert.Equal(t, `Hello...        
       ...world!`, device.String())
}
