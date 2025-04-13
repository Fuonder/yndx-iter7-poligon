package antipatterns

// структура, содержащая данные о балансе
//var Balance struct {
//	Amount int
//	// защищено мьютексом
//	sync.Mutex
//}
//
//// функция, вызываемая потоками «доходов»
//func Deposit(i int) {
//	// закрываем замок мьютекса
//	Balance.Lock()
//	// увеличиваем баланс
//	Balance.Amount += i
//	// открываем замок
//	Balance.Unlock()
//}
//
//// функция, вызываемая потоками «расходов»
//func Withdraw(i int) {
//	// закрываем замок мьютекса
//	Balance.Lock()
//	// уменьшаем баланс
//	Balance.Amount -= i
//	// открываем замок
//	Balance.Unlock()
//}

func Example2_bad() {
	// логика кошелька
	// запросы к балансу могут поступать конкурентно
	// из разных потоков
	//go Deposit(salary)
	//// ...
	//go Deposit(bonus)
	//// ...
	//go Deposit(interest)
	//// ...
	//go Withdraw(rent)
	// ... etc
}

// переменная, которую нужно менять из нескольких потоков
var Balance int

// канал коммуникации между горутинами
var ch chan int

// эту функцию можно безопасно вызывать из разных горутин конкурентно,
// она не изменяет состояние переменной в памяти, а только пишет в канал
func Deposit(i int) {
	ch <- i
}

// эта функция тоже
func Withdraw(i int) {
	ch <- -i
}

// а это единственная функция с доступом к Balance,
// она принимает данные из канала от многих источников
// и непосредственно изменяет состояние памяти
func Account() {
	for {
		Balance += <-ch
	}
}

func Example2() {
	// ...
	go Account()
	// ...
	//go Deposit(salary)
	//// ...
	//go Deposit(bonus)
	//// ...
	//go Withdraw(rent)
	// ...
}
