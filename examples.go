//Применение каналов для сигнализации о завершении работы горутин:
//В этом подходе используется дополнительный канал для сигнализации о завершении работы горутин.
//Когда все горутины завершают свою работу, они отправляют сигнал в этот канал.
//Вот пример:

//package main
//
//import (
//	"context"
//	"fmt"
//	"sync"
//	"time"
//)
//
//func worker(id int, tasks <-chan int, results chan<- int, done chan<- struct{}) {
//	defer func() {
//		done <- struct{}{}
//	}()
//	for task := range tasks {
//		results <- task * task
//	}
//}
//
//func main() {
//	numWorkers := 4
//	tasks := make(chan int, 10)
//	results := make(chan int, 10)
//	done := make(chan struct{}, numWorkers)
//	var wg sync.WaitGroup
//
//	for i := 0; i < numWorkers; i++ {
//		wg.Add(1)
//		go func(id int) {
//			defer wg.Done()
//			worker(id, tasks, results, done)
//		}(i)
//	}
//
//	// Заполнение канала задач
//	for i := 0; i < 10; i++ {
//		tasks <- i
//	}
//	close(tasks)
//
//	// Ожидание завершения работы всех горутин
//	go func() {
//		wg.Wait()
//		close(results)
//	}()
//
//	// Ожидание завершения работы всех горутин
//	for i := 0; i < numWorkers; i++ {
//		<-done
//	}
//
//	// Получение результатов
//	for result := range results {
//		fmt.Println(result)
//	}
//}

//Использование пакета sync/atomic для безопасного доступа к общим данным:
//В этом подходе используется пакет sync/atomic для безопасной работы с общими данными из нескольких горутин.
//Вот пример:

//package main
//
//import (
//"fmt"
//"sync"
//"sync/atomic"
//)
//
//func main() {
//	var sum int64
//	numWorkers := 4
//	var wg sync.WaitGroup
//
//	for i := 0; i < numWorkers; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for i := 0; i < 1000; i++ {
//				atomic.AddInt64(&sum, int64(i))
//			}
//		}()
//	}
//
//	wg.Wait()
//
//	fmt.Println("Sum:", sum)
//}

//package main
//
//import (
//"context"
//"fmt"
//"time"
//)
//
//func main() {
//	forever := make(chan struct{})
//	ctx, cancel := context.WithCancel(context.Background())
//
//	go func(ctx context.Context) {
//		for {
//			select {
//			case <-ctx.Done():  // if cancel() execute
//				forever <- struct{}{}
//				return
//			default:
//				fmt.Println("for loop")
//			}
//
//			time.Sleep(500 * time.Millisecond)
//		}
//	}(ctx)
//
//	go func() {
//		time.Sleep(3 * time.Second)
//		cancel()
//	}()
//
//	<-forever
//	fmt.Println("finish")
//}

//package main
//
//import (
//	"fmt"
//	"os"
//	"os/signal"
//	"syscall"
//	"time"
//)
//
//func worker() {
//	for {
//		fmt.Println("Working...")
//		time.Sleep(1 * time.Second)
//	}
//}
//
//func main() {
//	sig := make(chan os.Signal, 1)
//	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
//
//	go worker()
//
//	// Ожидаем сигнал остановки
//	<-sig
//	fmt.Println("Stopping...")
//	os.Exit(0)
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var (
//	done  bool
//	mutex sync.Mutex
//)
//
//func worker() {
//	for {
//		mutex.Lock()
//		if done {
//			mutex.Unlock()
//			fmt.Println("Stopping...")
//			return
//		}
//		mutex.Unlock()
//
//		fmt.Println("Working...")
//		time.Sleep(1 * time.Second)
//	}
//}
//
//func main() {
//	go worker()
//
//	time.Sleep(3 * time.Second)
//
//	mutex.Lock()
//	done = true // Установить флаг, чтобы остановить горутину
//	mutex.Unlock()
//
//	time.Sleep(1 * time.Second)
//	fmt.Println("Main goroutine exits")
//}

//package main
//
//import "fmt"
//
//func setBitToValue(num int64, pos uint, bitValue int64) int64 {
//	mask := int64(1 << pos) // Создаем маску с единицей в позиции i
//	mask = ^mask            // Инвертируем биты маски, чтобы установить i-й бит в 0
//
//	if bitValue == 1 {
//		return num | mask // Устанавливаем i-й бит в 1 с помощью побитовой операции ИЛИ
//	} else {
//		return num & mask // Устанавливаем i-й бит в 0 с помощью побитовой операции И
//	}
//}
//
//func main() {
//	var num int64 = -5   // Начальное число: 0101 в двоичном формате
//	pos := uint(1)       // Устанавливаем 2-й бит (с нуля)
//	bitValue := int64(0) // Устанавливаем бит в 0
//
//	result := setBitToValue(num, pos, bitValue)
//	fmt.Printf("Число после установки i-го бита в %d: %d\n", bitValue, result) // Ожидаемый результат: 1 (0001 в двоичном формате)
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func setBitToValue(num int64, pos uint, bitValue int64) int64 {
//	mask := int64(1 << pos) // Создаем маску с единицей в позиции i
//	b := strconv.FormatInt(mask, 2)
//	fmt.Println(b)
//	mask = ^mask // Инвертируем биты маски, чтобы установить i-й бит в 0
//
//	if bitValue == 1 {
//		return num | mask // Устанавливаем i-й бит в 1 с помощью побитовой операции ИЛИ
//	} else {
//		return num & mask // Устанавливаем i-й бит в 0 с помощью побитовой операции И
//	}
//}
//
//func main() {
//	var num int64 = 10   // Начальное число: 0101 в двоичном формате
//	pos := uint(1)       // Устанавливаем 2-й бит (с нуля)
//	bitValue := int64(0) // Устанавливаем бит в 0
//
//	result := setBitToValue(num, pos, bitValue)
//	fmt.Printf("Число после установки i-го бита в %d: %d\n", bitValue, result) // Ожидаемый результат: 1 (0001 в двоичном формате)
//}

