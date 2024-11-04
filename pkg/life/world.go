package life

import (
	"math/rand"
	"time"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

// Используйте код из предыдущего урока по игре «Жизнь»
func NewWorld(height, width int) *World {
	if height <= 0 || width <= 0 {
		return nil
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  make([][]bool, height),
	}
}

func (w *World) next(x, y int) bool {
	if x < 0 || x >= w.Height || y < 0 || y >= w.Width {
		return false
	}
	return w.Cells[x][y]

}

func (w *World) neighbors(x, y int) int {
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if w.next(i, j) {
				count++
			}
		}
	}
	return count
}

func NextState(oldWorld, newWorld World) {
	// Копируем состояния клеток из старого мира в новый
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			n := oldWorld.neighbors(i, j)
			alive := oldWorld.Cells[i][j]

			if alive && (n < 2 || n > 3) {
				newWorld.Cells[i][j] = false
			} else if !alive && n == 3 {
				newWorld.Cells[i][j] = true
			} else {
				newWorld.Cells[i][j] = alive
			}
		}
	}
}

// RandInit заполняет поля на указанное число процентов
func (w *World) RandInit(percentage int) {
	// Количество живых клеток
	numAlive := percentage * w.Height * w.Width / 100
	// Заполним живыми первые клетки
	w.fillAlive(numAlive)
	// Получаем рандомные числа
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// Рандомно меняем местами
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
