package day15

import (
	"advent-of-code-2024/util"
	"fmt"
)

func SolvePart1(useRealInput bool) (int, error) {
	w, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	for _, dir := range w.movement {
		w.move(dir)
	}
	score := w.sumGpsCoordinates()

	return score, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	w, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	w = w.widen()

	for _, dir := range w.movement {
		w.moveWide(dir)
	}
	score := w.sumGpsCoordinates()

	return score, nil
}

func (w *warehouse) move(dir util.Direction) {
	startPos := w.botLocation
	moveVec := dir.ToVec()
	p := startPos
	for {
		p = p.Plus(moveVec)
		curItem := w.get(p)
		if curItem == WALL {
			return // Cannot move anything
		} else if curItem == BOX {
			continue // keep looking
		} else if curItem == EMPTY {
			break
		} else {
			panic("Invalid item found")
		}
	}

	// Found empty spot, now move backwards
	oppositeMoveVec := moveVec.Times(-1)
	for {
		predecessor := p.Plus(oppositeMoveVec)
		if predecessor == startPos {
			w.botLocation = p
			return
		}

		// Move predecessor box into current place
		w.set(p, w.get(predecessor))
		w.set(predecessor, EMPTY)
		p = predecessor
	}
}

func (w *warehouse) moveWide(dir util.Direction) {
	startPos := w.botLocation
	moveVec := dir.ToVec()
	movingHorizontally := dir == util.LEFT || dir == util.RIGHT

	// First, we find which locations will be affected
	itemsToMove := [][]util.Vec{{startPos}}
rows:
	for i := 0; true; i++ {
		// Collect all locations that will be chagned by pushing previous row
		previousRow := itemsToMove[i]
		nextRow := make([]util.Vec, 0)
		for _, p := range previousRow {
			// empty items don't push anything forward
			previousItem := w.get(p)
			if p != startPos && previousItem == EMPTY {
				continue
			}

			// p contains a box, now see what it would move into
			p = p.Plus(moveVec)
			nextItem := w.get(p)
			if nextItem == WALL {
				return // Cannot move anything
			} else if movingHorizontally && (nextItem == LBOX || nextItem == RBOX) {
				nextRow = append(nextRow, p)
			} else if nextItem == LBOX {
				nextRow = append(nextRow, p, p.PlusDir(util.RIGHT))
			} else if nextItem == RBOX {
				nextRow = append(nextRow, p, p.PlusDir(util.LEFT))
			} else if nextItem == EMPTY {
				nextRow = append(nextRow, p)
			} else {
				panic("Invalid item found while moving wide")
			}
		}

		// We can get duplicate positions since boxes can be 2 wide
		uniqueLocations := make(map[util.Vec]bool, 0)
		for _, v := range nextRow {
			uniqueLocations[v] = true
		}
		keys := make([]util.Vec, 0, len(uniqueLocations))
		for v := range uniqueLocations {
			keys = append(keys, v)
		}
		nextRow = keys

		// nextRow only contains boxes or empty spaces
		for _, p := range nextRow {
			if w.get(p) != EMPTY {
				// at least one box in next row, so continue to next next row
				itemsToMove = append(itemsToMove, nextRow)
				continue rows
			}
		}

		// nextRow is completely empty space, so start to move
		break
	}

	// Working back from the end, move everything
	for i := len(itemsToMove) - 1; i > 0; i-- {
		currentRow := itemsToMove[i]
		for _, p := range currentRow {
			if w.get(p) != EMPTY {
				w.set(p.Plus(moveVec), w.get(p))
				w.set(p, EMPTY)
			}
		}
	}

	w.botLocation = startPos.Plus(moveVec)
}

func (w warehouse) sumGpsCoordinates() int {
	sum := 0
	for y, row := range w.warehouseMap {
		for x, item := range row {
			if item == BOX || item == LBOX {
				sum += 100*y + x
			}
		}
	}
	return sum
}

type item int

const (
	EMPTY item = iota
	WALL
	BOX
	LBOX
	RBOX
)

type warehouse struct {
	botLocation  util.Vec
	warehouseMap [][]item
	movement     []util.Direction
}

func (w warehouse) get(position util.Vec) item {
	return w.warehouseMap[position.Y][position.X]
}

func (w *warehouse) set(position util.Vec, newItem item) {
	w.warehouseMap[position.Y][position.X] = newItem
}

// func (w warehouse) print() {
// 	for y, row := range w.warehouseMap {
// 		for x, item := range row {
// 			switch {
// 			case x == w.botLocation.X && y == w.botLocation.Y:
// 				fmt.Print("@")
// 			case item == BOX:
// 				fmt.Print("O")
// 			case item == LBOX:
// 				fmt.Print("[")
// 			case item == RBOX:
// 				fmt.Print("]")
// 			case item == WALL:
// 				fmt.Print("#")
// 			default:
// 				fmt.Print(".")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

func (w warehouse) widen() *warehouse {
	wideMap := make([][]item, len(w.warehouseMap))

	for y, row := range w.warehouseMap {
		wideMap[y] = make([]item, 2*len(row))
		for x, item := range row {
			l, r := widen(item)
			wideMap[y][2*x] = l
			wideMap[y][2*x+1] = r
		}
	}

	return &warehouse{util.Vec{X: 2 * w.botLocation.X, Y: w.botLocation.Y}, wideMap, w.movement}

}

func widen(i item) (item, item) {
	switch i {
	case EMPTY:
		return EMPTY, EMPTY
	case BOX:
		return LBOX, RBOX
	case WALL:
		return WALL, WALL
	default:
		panic("Received invalid item to widen")
	}
}

func parseInput(useRealInput bool) (*warehouse, error) {
	data, err := util.ReadInputMulti(15, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 2 {
		return nil, fmt.Errorf("expected two sections of input")
	}

	warehouseMap := make([][]item, len(data[0]))
	var startPosition util.Vec
	for y, row := range data[0] {
		warehouseMap[y] = make([]item, len(row))
		for x, r := range row {
			var p item
			if r == rune('#') {
				p = WALL
			} else if r == rune('.') {
				p = EMPTY
			} else if r == rune('O') {
				p = BOX
			} else if r == rune('@') {
				startPosition = util.Vec{X: x, Y: y}
				p = EMPTY
			} else {
				panic("invalid warehouse item input!")
			}
			warehouseMap[y][x] = p
		}
	}

	movement := make([]util.Direction, 0)
	for _, row := range data[1] {
		for _, instruction := range row {
			var dir util.Direction
			switch instruction {
			case '<':
				dir = util.LEFT
			case 'v':
				dir = util.DOWN
			case '>':
				dir = util.RIGHT
			case '^':
				dir = util.UP
			default:
				panic("invalid movement input received")
			}
			movement = append(movement, dir)
		}
	}

	return &warehouse{startPosition, warehouseMap, movement}, nil
}
