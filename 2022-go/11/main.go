package main

import (
	"fmt"
)

func main() {
	//input, err := os.ReadFile("./2022-go/11/sample.txt")
	//input, err := os.ReadFile("./2022-go/11/input.txt")
	//if err != nil {
	//	panic(err)
	//}

	// monkeys are hand written, hope I didn't screw it up
	monkeys := inputMonkeys()
	//monkeys := sampleMonkeys()
	// Part 1
	//const numRounds = 20
	// Part 2
	const numRounds = 10000
	//commonMod := 96577
	commonMod := 9699690

	// process all rounds
	for i := 0; i < numRounds; i++ {
		switch i {
		case 1, 20, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000:
			fmt.Printf("*** After %d rounds ***\n", i)
			for j, m := range monkeys {
				fmt.Printf("Monkey %d: %d\n", j, m.countInspections)
			}
		}
		// each monkey takes its turn
		for _, m := range monkeys {
			//inspect each item
			for k, v := range m.items {
				m.countInspections++
				v.value = m.op(v.value) % commonMod
				// Part 1 only
				//v.value = v.value / 3
				var throwTarget int
				if v.value%m.throwDivisor == 0 {
					throwTarget = m.throwTrue
				} else {
					throwTarget = m.throwFalse
				}
				monkeys[throwTarget].items[k] = v
				delete(m.items, k)
			}
		}
	}

	// calculate monkey business
	var scoreA, scoreB int
	fmt.Println("***Monkey Business***")
	for i, m := range monkeys {
		fmt.Printf("Monkey %d: %d\n", i, m.countInspections)
		switch {
		case m.countInspections > scoreA:
			scoreB = scoreA
			scoreA = m.countInspections
		case m.countInspections > scoreB:
			scoreB = m.countInspections
		}
	}

	total := scoreA * scoreB
	fmt.Printf("A: %d, B: %d\nMonkey Business: %d\n", scoreA, scoreB, total)
}

type monkey struct {
	items                               map[int]*myItem
	op                                  func(int) int
	throwDivisor, throwTrue, throwFalse int
	countInspections                    int
}

type myItem struct {
	id    int
	value int
}

// there's only 8... hardcode it
func inputMonkeys() []*monkey {
	items := inputItems()
	monkeys := make([]*monkey, 8, 8)
	monkeys[0] = &monkey{
		items: map[int]*myItem{
			items[0].id: items[0],
			items[1].id: items[1],
			items[2].id: items[2],
			items[3].id: items[3],
			items[4].id: items[4],
		},
		op:           func(old int) int { return old * 11 },
		throwDivisor: 19,
		throwTrue:    6,
		throwFalse:   7,
	}
	monkeys[1] = &monkey{
		items: map[int]*myItem{
			items[5].id: items[5],
			items[6].id: items[6],
			items[7].id: items[7],
		},
		//op:               func(old *big.Int) { old.Add(old, big.NewInt(8)) },
		op:           func(old int) int { return old + 8 },
		throwDivisor: 2,
		throwTrue:    6,
		throwFalse:   0,
	}
	monkeys[2] = &monkey{
		items: map[int]*myItem{
			items[8].id:  items[8],
			items[9].id:  items[9],
			items[10].id: items[10],
			items[11].id: items[11],
			items[12].id: items[12],
			items[13].id: items[13],
			items[14].id: items[14],
		},
		//op:               func(old *big.Int) { old.Add(old, big.NewInt(1)) },
		op:           func(old int) int { return old + 1 },
		throwDivisor: 3,
		throwTrue:    5,
		throwFalse:   3,
	}
	monkeys[3] = &monkey{
		items: map[int]*myItem{
			items[15].id: items[15],
		},
		//op:               func(old *big.Int) { old.Mul(old, big.NewInt(7)) },
		op:           func(old int) int { return old * 7 },
		throwDivisor: 17,
		throwTrue:    5,
		throwFalse:   4,
	}
	monkeys[4] = &monkey{
		items: map[int]*myItem{
			items[16].id: items[16],
			items[17].id: items[17],
			items[18].id: items[18],
			items[19].id: items[19],
			items[20].id: items[20],
			items[21].id: items[21],
			items[22].id: items[22],
			items[23].id: items[23],
		},
		//op:               func(old *big.Int) { old.Add(old, big.NewInt(4)) },
		op:           func(old int) int { return old + 4 },
		throwDivisor: 13,
		throwTrue:    0,
		throwFalse:   1,
	}
	monkeys[5] = &monkey{
		items: map[int]*myItem{
			items[24].id: items[24],
			items[25].id: items[25],
			items[26].id: items[26],
			items[27].id: items[27],
			items[28].id: items[28],
			items[29].id: items[29],
		},
		//op:               func(old *big.Int) { old.Add(old, big.NewInt(7)) },
		op:           func(old int) int { return old + 7 },
		throwDivisor: 7,
		throwTrue:    1,
		throwFalse:   4,
	}
	monkeys[6] = &monkey{
		items: map[int]*myItem{
			items[30].id: items[30],
			items[31].id: items[31],
			items[32].id: items[32],
			items[33].id: items[33],
		},
		//op:               func(old *big.Int) { old.Mul(old, old) },
		op:           func(old int) int { return old * old },
		throwDivisor: 5,
		throwTrue:    7,
		throwFalse:   2,
	}
	monkeys[7] = &monkey{
		items: map[int]*myItem{
			items[34].id: items[34],
			items[35].id: items[35],
		},
		//op:               func(old *big.Int) { old.Add(old, big.NewInt(6)) },
		op:           func(old int) int { return old + 6 },
		throwDivisor: 11,
		throwTrue:    2,
		throwFalse:   3,
	}
	return monkeys
}

func inputItems() []*myItem {
	return []*myItem{
		{
			id:    0,
			value: 74,
		},
		{
			id:    1,
			value: 73,
		},
		{
			id:    2,
			value: 57,
		},
		{
			id:    3,
			value: 77,
		},
		{
			id:    4,
			value: 74,
		},
		{
			id:    5,
			value: 99,
		},
		{
			id:    6,
			value: 77,
		},
		{
			id:    7,
			value: 79,
		},
		{
			id:    8,
			value: 64,
		},
		{
			id:    9,
			value: 67,
		},
		{
			id:    10,
			value: 50,
		},
		{
			id:    11,
			value: 96,
		},
		{
			id:    12,
			value: 89,
		},
		{
			id:    13,
			value: 82,
		},
		{
			id:    14,
			value: 82,
		},
		{
			id:    15,
			value: 88,
		},
		{
			id:    16,
			value: 80,
		},
		{
			id:    17,
			value: 66,
		},
		{
			id:    18,
			value: 98,
		},
		{
			id:    19,
			value: 83,
		},
		{
			id:    20,
			value: 70,
		},
		{
			id:    21,
			value: 63,
		},
		{
			id:    22,
			value: 57,
		},
		{
			id:    23,
			value: 66,
		},
		{
			id:    24,
			value: 81,
		},
		{
			id:    25,
			value: 93,
		},
		{
			id:    26,
			value: 90,
		},
		{
			id:    27,
			value: 61,
		},
		{
			id:    28,
			value: 62,
		},
		{
			id:    29,
			value: 64,
		},
		{
			id:    30,
			value: 69,
		},
		{
			id:    31,
			value: 97,
		},
		{
			id:    32,
			value: 88,
		},
		{
			id:    33,
			value: 93,
		},
		{
			id:    34,
			value: 59,
		},
		{
			id:    35,
			value: 80,
		},
	}
}

// there's only 4... hardcode it
func sampleMonkeys() []*monkey {
	items := sampleItems()
	monkeys := make([]*monkey, 4, 4)
	monkeys[0] = &monkey{
		items: map[int]*myItem{
			items[0].id: items[0],
			items[1].id: items[1],
		},
		//op:               func(old *big.Int) { old.Mul(old, big.NewInt(19)) },
		op:           func(old int) int { return old * 19 },
		throwDivisor: 23,
		throwTrue:    2,
		throwFalse:   3,
	}
	monkeys[1] = &monkey{
		items: map[int]*myItem{
			items[2].id: items[2],
			items[3].id: items[3],
			items[4].id: items[4],
			items[5].id: items[5],
		},
		//op:               func(old *big.Int) { old.Add(old, big.NewInt(6)) },
		op:           func(old int) int { return old + 6 },
		throwDivisor: 19,
		throwTrue:    2,
		throwFalse:   0,
	}
	monkeys[2] = &monkey{
		items: map[int]*myItem{
			items[6].id: items[6],
			items[7].id: items[7],
			items[8].id: items[8],
		},
		//op:               func(old *big.Int) { old.Mul(old, old) },
		op:           func(old int) int { return old * old },
		throwDivisor: 13,
		throwTrue:    1,
		throwFalse:   3,
	}
	monkeys[3] = &monkey{
		items: map[int]*myItem{
			items[9].id: items[9],
		},
		//op:               func(old *big.Int) { old.Add(old, big.NewInt(3)) },
		op:           func(old int) int { return old + 3 },
		throwDivisor: 17,
		throwTrue:    0,
		throwFalse:   1,
	}
	return monkeys
}

func sampleItems() []*myItem {
	return []*myItem{
		{
			id:    0,
			value: 79,
		},
		{
			id:    1,
			value: 98,
		},
		{
			id:    2,
			value: 54,
		},
		{
			id:    3,
			value: 65,
		},
		{
			id:    4,
			value: 75,
		},
		{
			id:    5,
			value: 74,
		},
		{
			id:    6,
			value: 79,
		},
		{
			id:    7,
			value: 60,
		},
		{
			id:    8,
			value: 97,
		},
		{
			id:    9,
			value: 74,
		},
	}
}
