package main

/*
Нужно два примера объяснить, разобрать с разными версиями GO.
Почему, когда я выставляю runtime.GOMAXPROCS(4), то порядок выполнения может быть любым:

А если выставляю runtime.GOMAXPROCS(1)
То всегда будет начинаться:
Почему, КОЛЯ? 8
*/

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// устанавливаем максимальное кол-во горутин
	runtime.GOMAXPROCS(1)
	// создаем группу
	wg := sync.WaitGroup{}
	wg.Add(9)
	for i := 0; i < 10; i++ {
		i := i
		fmt.Println(i)
		go func(i int) {
			defer wg.Done()
			fmt.Println("                           Почему, КОЛЯ?", i)
		}(i)
	}
	wg.Wait()

	// time.Sleep(1 * time.Second)
	fmt.Println("Паника")
}

/* ОБЪЯСНЕНИЕ

1) все дело в асинхронном выполнении при количестве горутин больше чем одна

2) При GOMAXPROCS больше чем единица вывод итерации будет больше

*/
