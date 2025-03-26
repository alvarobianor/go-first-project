package highOrder

import (
	"fmt"
	"math/rand"
	"time"
)

func orderArray(arr *[6]uint8) {
	temp := make([]uint8, len(arr))
	mergeSort(arr[:], temp, 0, len(arr)-1)
}

func mergeSort(arr []uint8, temp []uint8, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(arr, temp, left, mid)
		mergeSort(arr, temp, mid+1, right)
		merge(arr, temp, left, mid, right)
	}
}

func merge(arr []uint8, temp []uint8, left int, mid int, right int) {
	for i := left; i <= right; i++ {
		temp[i] = arr[i]
	}

	i, j, k := left, mid+1, left

	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}

	for i <= mid {
		arr[k] = temp[i]
		i++
		k++
	}
}

func generateGame(game *[6]uint8) {
	i := 0
	for i < 6 {
		number := rand.Intn(61)
		if number == 0 {
			continue
		}
		hasIqual := false
		for _, value := range game {
			if value == uint8(number) {
				hasIqual = true
				break
			}
		}
		if !hasIqual {
			game[i] = uint8(number)
			i++
		}

	}
	orderArray(game)
}

func showGame(game *[6]uint8) {
	generateGame(game)
}

func CreateStats() {
	begin := time.Now()

	count := 0
	gamesWin := [][6]uint8{}
	gamesAlmost5 := [][6]uint8{}
	gamesAlmost4 := [][6]uint8{}

	numberOfGames := 10000000

	for range numberOfGames {
		var game [6]uint8 = [6]uint8{}
		showGame(&game)

		var sortedGame [6]uint8 = [6]uint8{}
		showGame(&sortedGame)

		for i, num := range sortedGame {
			if game[i] == num {
				count++
			}
		}

		if count == 6 {
			gamesWin = append(gamesWin, game)
		} else if count == 5 {
			gamesAlmost5 = append(gamesAlmost5, game)
		} else if count == 4 {
			gamesAlmost4 = append(gamesAlmost4, game)
		}
		count = 0

	}
	ends := time.Since(begin)

	fmt.Println("Games Win -> \n", gamesWin)
	fmt.Println("Games almost with 5 -> \n", gamesAlmost5)
	fmt.Println("Games almost with 4 -> \n", gamesAlmost4)

	maxWin := len(gamesWin)
	almost5 := len(gamesAlmost5)
	almost4 := len(gamesAlmost4)

	fmt.Println("Total of games win some win-> ", maxWin+almost5+almost4)
	fmt.Println("Total of games win max win -> ", maxWin, " with a chance of -> ", float64(maxWin)/float64(numberOfGames)*100, "%")
	fmt.Println("Total of games win almost 5 -> ", almost5, " with a chance of -> ", float64(almost5)/float64(numberOfGames)*100, "%")
	fmt.Println("Total of games win almost 4 -> ", almost4, " with a chance of -> ", float64(almost4)/float64(numberOfGames)*100, "%")
	fmt.Println("time -> ", ends)
}
