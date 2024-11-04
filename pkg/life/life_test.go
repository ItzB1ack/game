package life_test

// Так как это другой пакет, нужно его импортировать
import (
	"testing"

	"github.com/ItzB1ack/game/pkg/life"
)

func TestNewWorld(t *testing.T) {
	type test struct {
		height  int
		width   int
		wantErr bool
	}

	tests := []test{
		{height: 0, width: 4, wantErr: true},
		{height: -1, width: 4, wantErr: true},
		{height: 5, width: 0, wantErr: true},
		{height: 5, width: 6, wantErr: false},
	}

	for _, tt := range tests {
		height := tt.height
		width := tt.width
		world, err := life.NewWorld(height, width)
		if err != nil {
			if tt.wantErr {
				continue
			}
			t.Errorf("Unexpected error: %s", err)
		}

		if world.Height != height {
			t.Errorf("expected height: %d, actual height: %d", height, world.Height)
		}
		if world.Width != width {
			t.Errorf("expected width: %d, actual width: %d", width, world.Width)
		}

		if len(world.Cells) != height {
			t.Errorf("expected height: %d, actual number of rows: %d", height, len(world.Cells))
		}
	}
}
