# Конспекрт курса Go

Структура материала:

-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#настройка-окружения-и-выбор-редактора">Настройка окружения и выбор редактора</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#базовая-программа">Базовая программа</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#переменные">Переменные</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#комментарии">Комментарии</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#типы-переменных-и-нулевые-значения">Типы переменных и нулевые значения</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#множественное-присвоение">Множественное присвоение</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#область-видимости-и-функции">Область видимости и функции</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#условный-оператор-if-обработка-ошибок-и-множественные-возвращаемые-значения">Условный оператор if, обработка ошибок и множественные возвращаемые значения</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#конструкция-switch-case">Конструкция switch case</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#неограниченное-количество-аргументов-в-функции">Неограниченное количество аргументов в функции</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#анонимные-функции-замыкания">Анонимные функции. Замыкания</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#функция-init">Функция init()</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#ссылки-и-указатели">Ссылки и указатели</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#массивы-и-слайсы">Массивы и слайсы</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#матрицы-цикл-for">Матрицы. Цикл for</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#panic-recover-defer">panic(), recover(), defer</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#мапы">Мапы</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#структуры">Структуры</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#интерфейсы">Интерфейсы</a>
-   <a href="https://github.com/pavloging/go-course-in-4-hours?tab=readme-ov-file#пакеты-и-модули">Пакеты и модули</a>

## Настройка окружения и выбор редактора

Автор показывает, как установить Go и настроить VSCode. Данное видео заменяет вводный урок по настройке.

## Базовая программа

Создание программы Hello World!

```go
package main

import "fmt"

func main() {
	fmt.Println("Hi World")
}
```

Автор подчеркивает:

-   Объяснение команды go run и go build
-   Объяснение структуры программы (package, import, func main())
-   package main - точка входа в наше приложение, и программа начинается с него
-   import "название библиотеки"
-   Показывает документацию на примере fmt
-   func main() - точка входа в наше приложение

## Переменные

Короткий вариант объявления переменной:

```go
message := "Hi World"
```

Полный вариант объявления переменной:

```go
const message string = "Hi World"
var message = "Hi World"
```

Узнать тип переменной:

```go
import "reflect"

message := "Hello World"
reflect.TypeOf(message) // Вернет: string
```

Автор подчеркивает:

-   Если объявили переменную и не используем её, Go выдает ошибку, чтобы не занимать лишнее место.
-   Go сам определяет тип примитивов, и мы можем его не указывать. Например: var message = "Hi World".
-   Отличия var от const: var — переиспользуемая переменная, а const — нет.
-   При объявлении переменной указывается тип переменной, который нельзя поменять. Если мы указываем тип для переменной int, то значение типа string вызовет ошибку.
-   Узнать тип переменной: reflect.TypeOf(message).

## Комментарии

Комментарии — текст, который не воспринимается компилятором.

```go
// Однострочные комментарии
/*
    Многострочные
    Комментарии
*/
```

## Типы переменных и нулевые значения

```go
var message string // Значение переменной по умолчанию: ""
var number int // Значение переменной по умолчанию: 0

var num int8 := 5 // У числа разрядность 8 бит. Это значит, что значение переменной может быть от -128 до 127 и данная переменная занимает 8 бит (1 байт). 0 - по умолчанию

var unum uint := 42 // uint - тип данных положительных чисел включая 0. 0 - по умолчанию

var floatNum float32 := 42.42 // float32, float64 - тип данных чисел с плавающей точкой. 0 - по умолчанию

var b bool = false // bool - тип даннных в качестве значения имеет true/false. false - по умолчанию
b = true

var messByte := []byte("123") // byte - тип данных синоним unit8 и сопоставляется с ASCII
// Числа 1 и 2 и 3 попадают в массив и переобразуются в значение бита, по таблице ASCII
var a byte = 2
fmt.Printf("%c\n", a) // Вывод: ☻
// По итогу byte это значение которое сопоставляется с таблице ASCII и возвращается нам


```

Автор подчеркивает:

-   Если не передавать значение переменной, то для типа string переменной будет присвоена пустая строка "". Для числа — 0.
-   Разрядность числа: int8, int16, int32, int64. Данная разрядность зависит от системы, на которой запускается программа (зависит от процессора). 0 - по умолчанию.
-   uint - тип данных положительных чисел, включая 0. 0 - по умолчанию.
-   float32, float64 - тип данных чисел с плавающей точкой. 0 - по умолчанию.
-   bool - тип данных в качестве значения имеет true/false. false - по умолчанию.
-   byte - тип данных, синоним uint8 и сопоставляется с ASCII.
-   rune - тип данных, синоним int32 и применяется наоборот от byte. То есть с значения, например "a", мы получим его бит - 97.
-   Двойные строки нужны для строк.
-   Одинарные строки нужны для символов.

## Множественное присвоение

```go

a, b, c := 1, 2, 3

a, b = b, a // Под копотом: a, b = 2, 1

a, _, c = b, a, 2 // _ - это пропуск значения

```

Автор подчеркивает:

-   Множественное присвоение: a, b = b, a // Под копотом: a, b = 2, 1
-   "\_" - это пропуск значения.
-   Переменные сохраняются в оперативной памяти компьютера.

## Область видимости и функции

Пример с областью видимости:

```go
var a, b, c int // Глобальные

func main() {
    a, b, c := 1, 2, 3

    fmt.Println(a, b, c)    // Локальные. Логика: Запрашиваем переменные с локальной области
                            // Если переменной нет, то смотрим на глобальную
                            // То есть, приоритет в локальной области выше

    print()
}

func print() {
    a, b, c = 4, 5, 6 // Изменяем глобальные переменные
    fmt.Println(a, b, c)
}
```

Функции:

```go
func main() {
    var message string
    message = sayHello("Егор", 20) // Получаем значение, которое вернули от функции

    printMessage(message)
}

func printMessage(message string) {
    fmt.Println(message)
}

func sayHello(name string, age int) string { // Типизируем аргументы функции
    return fmt.Sprintf("Привет, %s! Тебе %d лет", name, age) // fmt.Sprintf - возращает строку, которую мы возращаем из функции.
}
```

Автор подчеркивает:

-   Всего два типа переменных по области видимости: локальные и глобальные.
-   Есть приватные и публичные переменные:
    -   Приватные переменные доступны только внутри пакета.
    -   Публичные доступны вне рамок пакета.
-   Функция - переиспользуемый блок кода.
-   Для возврата значения в функции используем return.

## Условный оператор if, обработка ошибок и множественные возвращаемые значения

-   && - логическое И
-   || - логическое ИЛИ
-   ! - логическое НЕ

```go

func main() {
    message, err := enterTheClub(18)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}

func enterTheClub(age int) (string, error) {
    if age >= 18 && age < 45 {
        return "Добро пожаловать в клуб!", nil // return - Функция возвращает значение и заканчивает свое выполнение
    } else if age >= 45 && age < 65 {
        return "Приветствуем вас в клубе!", nil
    } else if age >= 65 {
        return "Это уже слишком много для вас", errors.New("you are too old")
    }

    return "Вам недостаточно возраста для входа в клуб.", errors.New("you are too young")
}

```

Автор подчеркивает:

-   nil - тип данных, обозначает отсутствие значения.
-   error - тип данных, обозначает ошибку.
-   Правила написания ошибок в Go:
    1. Ошибка на английском.
    2. В ошибке кратко отражена суть.
    3. Начало строки начинается с маленькой буквы.
    4. При передаче ошибки в аргумент функции, данный аргумент выносим в конец аргументов.

## Конструкция switch case

switch case - альтернативное использование конструкции else if ...

```go

func main() {
    fmt.Println(prediction("чт"))
}

func prediction(dayOfWeek string) (string, error) {
    // if dayOfWeek == "пн" {
    //   return "Желаю тебе хорошего начала недели"
    // } else if dayOfWeek == "вт" {
    //   return "Хорошего тебе вторника"
    // } else if dayOfWeek == "ср" {
    //   return "Хорошей тебе среды"
    // } else if dayOfWeek == "чт" {
    //   return "Хорошего тебе четверга"
    // }

    switch dayOfWeek {
    case "пн":
        return "Желаю тебе хорошего начала недели!", nil
    case "вт":
        return "Прекрасного тебе вторника!", nil
    case "ср":
        return "Замечательной тебе среды!", nil
    case "чт":
        return "Хорошего тебе четверга!", nil
    default: // default - выполнится тогда, когда никакой case не отработал
        return "Неизвестный день недели", errors.New("invalid day of the week")
    }
}

```

Автор подчеркивает:

-   Когда мы хотим проверить, является ли один элемент равным другому по значению, то мы используем ==.

## Неограниченное количество аргументов в функции

Для получения неограниченного количества элементов в функции до указания типа вставляем ... Например, ...int, как продемонстрировано ниже:

```go
func main() {
	fmt.Println(findMin(1, 2, 9329, -1, 0, 42))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
	}

	return min
}

```

Автор подчеркивает:

-   Встроенная функция len() возвращает количество элементов.
-   for - цикл, который позволяет пройтись по каждому элементу массива.

## Анонимные функции. Замыкания

Что такое анонимные функции?

Анонимные функции в Go — это функции, которые не имеют имени. Они могут быть определены и вызваны на месте, что делает их удобными для использования в качестве аргументов для других функций или для создания замыканий.

```go
func main() {
    // Определяем и вызываем анонимную функцию
    func() {
        fmt.Println("Привет из анонимной функции!")
    }() // Вызов функции

    // Пример с передачей анонимной функции в другую функцию
    sum := func(a int, b int) int {
        return a + b
    }
    fmt.Println(sum(3, 5)) // Вывод: 8
}

```

Что такое замыкания?

Замыкания - это функция, которая определена внутри другой функции и использует переменные и аргументы внешней функции.

```go
func main() {
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func increment() func() int {
	count := 0
	// Возвращаем новую функцию, которая увеличивает счетчик на 1 и возвращает его текущее значение
	return func() int {
		count++
		return count
	}
}
```

## Функция init()

Функция init() срабатывает перед вызовом функции main() и перед инициализацией пакетов

```go
var message string

func init() {
	message = "new value"
}

func main() {
	fmt.Println(message)
}

```

## Ссылки и указатели

-   Указатели в Go — это переменные, которые хранят адреса других переменных.
    Они позволяют работать с данными по ссылке, что позволяет изменять значения переменных без необходимости копирования их значений.
    Указатели обозначаются символом \* перед типом, а для получения адреса переменной используется символ &.
    Указатель является отдельным типом данных.

-   "&" - ссылка (reference):
    Этот оператор используется для получения адреса переменной,
    позволяя нам передать ссылку на её значение.
    Например, если у нас есть переменная, мы можем использовать & перед её именем,
    чтобы получить адрес, по которому она хранится в памяти.

-   "_" - разыменование (dereference):
    Этот оператор позволяет нам получить доступ к значению,
    на которое указывает указатель.
    Используя _ перед указателем, мы можем изменять или читать значение,
    находящееся по этому адресу.

```go

package main

import "fmt"

func main() {
	message := "Скоро я стану Go программистом!"
	printMessage(&message) // Мы передаем адрес переменной message в функцию printMessage с помощью оператора &. Это означает, что в функцию будет передан указатель на строку, а не сама строка.

	fmt.Println(message)
}

func printMessage(message *string) {
	*message += " И буду курить бамбук!" // В этой функции message является указателем на строку. Используя оператор *, мы можем изменять значение, на которое указывает указатель.
	fmt.Println(message)
}

```

## Массивы и слайсы

-   Массив является хранилищем данных одного типа с фиксированной длиной, которая задается при инициализации массива.  
    Синтаксис:

```go
var arr = [3]string{"1", "2", "3"}
```

1. [3] - количество элементов.
2. string - тип.
3. {} - значения, которые будут в данном массиве.  
   Также есть возможность задавать количество элементов автоматически. Для этого нужно передать вместо числа троеточие:  
   `go[...]string{"1", "2", "3"}`

К массиву можно обращаться по индексу для получения одного элемента (индекс начинается с нуля).  
Получим первый элемент массива:

```go
var arr = [3]int{100, 32, 105}
fmt.Println(arr[0]) // Вывод: 100

// Получаем индекс последнего элемента
lastIndex := len(arr) - 1
fmt.Println(arr[lastIndex]) // Вывод: 105
```

-   Слайс является тем же массивом, но с рядом отличий:

1. Слайс хранит ссылку на массив, а не сам массив. Это означает, что слайсы могут динамически изменять свой размер и могут ссылаться на одну и ту же часть памяти.
2. Слайсы имеют длину и емкость, которые можно изменять. Длина слайса — это количество элементов в слайсе, а емкость — это количество элементов, которые могут быть размещены в слайсе, начиная с его первого элемента.
3. Слайсы могут быть легко расширены с помощью функции append(), что позволяет добавлять новые элементы в слайс без необходимости создания нового массива.
4. Слайсы могут быть созданы из массивов, других слайсов или с помощью литералов слайсов, что делает их более гибкими в использовании.
5. Слайсы поддерживают более удобные операции по работе с подмножествами данных, позволяя создавать новые слайсы из существующих слайсов с указанием диапазонов.

Пример создания слайса:

```go
func main() {
    messages := []string{"1", "2", "3"}
    printMessages(messages)
    fmt.Println(messages)
}

func printMessages(messages []string) error {
    if len(messages) == 0 {
        return errors.New("empty array")
    }
    messages[1] = "5"

    fmt.Println(messages)

    return nil
}
```

Альтернатива создания слайса через make и указание capacity:

```go
a := make([]string, 5) // len=5, cap=5
b := make([]string, 5, 15) // len=5, cap=15
```

Автор подчеркивает:

-   Индекс в массиве начинается с нуля.
-   Массивы с разной длиной - это разные типы.
-   Слайс является ссылкой на массив с возможностью добавления/удаления/перезаписи элементов.
-   Когда мы достигаем лимита capacity в слайсе, то значение capacity умножается на 2. Но при большом значении коэффициент умножения снижается.

## Матрицы. Цикл for

Цикл for — это конструкция, которая позволяет выполнять код несколько раз, основываясь на заданном условии.

### Инициализация матрицы

Чтобы создать матрицу в Go, мы можем использовать следующий код:

matrix := make([][]int, 10) // Инициализация матрицы размером 10x10

### Заполнение матрицы с помощью циклов

Чтобы заполнить матрицу значениями, мы можем использовать вложенные циклы for. Вот пример:

```go
for x := 0; x < 10; x++ {
for y := 0; y < 10; y++ {
matrix[y] = make([]int, 10) // Инициализация строки матрицы
matrix[y][x] = x // Заполнение матрицы
}
fmt.Println(matrix[x]) // Вывод строки матрицы
}
```

### Разные способы использования цикла for

Цикл for в Go поддерживает несколько форм записи. Например, вы можете использовать бесконечный цикл без указания итерации:

```go
for x := 0; true; {
    x++
    fmt.Println(x)
}
```

Другие примеры бесконечных циклов:

```go
// Пример 2
x := 0
for true {
    x++
    fmt.Println(x)
}

// Пример 3
a := 0
for {
    a++
    fmt.Println(a)
}
```

### Полное заполнение матрицы

Вот как можно заполнить матрицу последовательными числами:

```go
matrix := make([][]int, 10) // Инициализация матрицы
counter := 0

for x := 0; x < 10; x++ {
    matrix[x] = make([]int, 10) // Инициализация строки матрицы
    for y := 0; y < 10; y++ {
        counter++ // Увеличиваем счетчик
        matrix[x][y] = counter // Заполняем матрицу
    }
    fmt.Println(matrix[x]) // Вывод строки матрицы
}
```

### Альтернатива обычному for с индексами

Вы также можете использовать конструкцию range для перебора элементов в срезах. Например:

```go
messages := []string{
    "message 1",
    "message 2",
    "message 3",
    "message 4",
}

for i, message := range messages { // Аналог for i := 0; i < len(messages); i++
    fmt.Println(messages[i]) // Выводим элемент по индексу
    fmt.Println(message) // Выводим текущий элемент
}
```

### Оператор break

Если вам нужно выйти из цикла раньше времени, вы можете использовать оператор break. Вот пример:

```go
counter := 0
for {
    if counter == 100 {
        break // Выход из цикла, когда счетчик достигает 100
    }
    counter++
    fmt.Println(counter) // Выводим значение счетчика
}
```

## panic(), recover(), defer

### Что такое panic()?

Это ошибка которая проявляется в runtime, если вы пытаетесь получить доступ к элементу массива или среза за пределами его границ. Например слайс состоит из 3-х элементов, а мы хотим к 4-му элементу добавить новое значение - будет паника!
Например:

```go
messages := []string{"message 1", "message 2", "message 3"}
messages[3] = "message 4" // Это вызовет панику!
```

Так же панику можно вызвать сомостоятельно:

```go
panic("Текст нашей паники") // Вызывает панику с указанным сообщением
```

### Что такое defer?

defer — это ключевое слово в Go, которое откладывает выполнение функции до тех пор, пока не завершится окружающая функция. Это полезно, когда нужно выполнить определенные действия в конце, например, закрыть файл или освободить ресурсы.
Например:

```go
defer fmt.Println("1") // Эта строка будет выполнена в самом конце
fmt.Println("2") // Выводит 2
fmt.Println("3") // Выводит 3
// Вывод будет: 2 3 1
```

### Что такое recover()?

recover() — это функция в Go, которая используется для обработки паник в программе. Если паника происходит, recover() позволяет восстановить выполнение программы и предотвратить её аварийное завершение. Это особенно полезно, когда вы хотите обработать ошибку и продолжить работу программы.
Например:

```go
package main

import "fmt"

func main() {
    defer handlePanic() // Отложенный вызов функции handlePanic

    messages := []string{
        "message 1",
        "message 2",
        "message 3",
        "message 4",
    }

    // Попытка доступа к несуществующему индексу
    messages[4] = "message 5" // Это вызовет панику

    fmt.Println(messages) // Эта строка не будет выполнена
}

// Функция для обработки паники
func handlePanic() {
    if r := recover(); r != nil { // Если произошла паника, recover() вернет значение
        fmt.Println("Произошла паника:", r) // Выводим сообщение о панике
    }

    fmt.Println("handlePanic() выполнилась успешно") // Сообщение о том, что функция выполнена
}

```

Автор подчеркивает:

-   При возникновении паники (исполнении panic()) приложение завершает работу с кодом состояния 2.
-   Использование defer может добавить небольшую задержку (около 50 миллисекунд) к выполнению программы.
-   handlePanic() редко используется в проде (рабочих системах), но может быть полезен для отладки и обработки ошибок.

## Мапы

Map (или карта) — это тип данных, который представляет собой словарь. Он работает по принципу "ключ-значение", где каждый ключ должен быть уникальным (не может повторяться).

Вот пример объявления и инициализации map:

```go
users := map[string]int{
    "Alex":  10,
    "Ivan":  25,
    "Artem": 42,
}
fmt.Println(users["Alex"]) // Вывод: 10
```

Если вы попытаетесь обратиться к ключу, которого нет в map, будет возвращено значение по умолчанию. Например:

```go
fmt.Println(users["Vova"]) // Вывод: 0
```

Также в map можно проверить, существует ли элемент. Для этого нужно использовать вторую переменную, которую обычно называют exists или ok. Если значение существует - переменная будет равна true, если нет — false. Пример:

```go
age, exists := users["Ivan"]
fmt.Println(age, exists) // Вывод: 25 true
```

### Итерация по map

Вы можете проходить по элементам map так же, как и по срезам, используя цикл for:

for key, value := range users {
fmt.Println(key, value)
}

### Перезапись значения

Чтобы изменить значение по существующему ключу, просто присвойте новое значение:

```go
users["Alex"] = 52 // Теперь значение для "Alex" равно 52
```

### Удаление значения по ключу

Чтобы удалить элемент из map, используйте функцию delete:

```go
delete(users, "Alex") // Удаляет значение по ключу "Alex"
```

### Добавление значений после инициализации

Если вы хотите использовать map, вам нужно ее инициализировать с помощью функции make. Если вы просто объявите переменную, она будет равна nil, и это приведет к ошибке:

-   Без make (это вызовет панику):

```go
var users map[string]int // users == nil
// users["Alice"] = 1 // Это вызовет панику!
```

-   С make (это работает корректно):

```go
users := make(map[string]int) // Теперь users готова к использованию
users["Alice"] = 1 // Теперь это работает
```

### Определение количества элементов в map

Чтобы узнать, сколько элементов в map, используйте функцию len:

users := map[string]int{"one": 1, "two": 2}
fmt.Println(len(users)) // Вывод: 2

## Структуры

Структура в Go — это пользовательский тип данных, который позволяет хранить данные в формате "ключ-значение" и методы. С помощью структур вы можете группировать связанные данные и добавлять методы для работы с ними. Это позволяет создавать множество объектов на основе одной структуры и переиспользовать их в коде.

В кратце, структуры нужны для создания кастомных типов, которые можно переиспользовать.

### Не переиспользуемая структура

Пример создания структуры без возможности повторного использования:

```go
user := struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}{"Vova", 23, "Male", 75, 188}

	fmt.Printf("%+v", user)
```

### Переиспользуемая структура

Пример создания структуры, которую можно использовать многократно:

```go
type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func main() {
	user1 := User{"Vova", 23, "Male", 75, 188}
	user2 := User{"Petya", 27, "Male", 81, 198}

    // Получение значения струкутры
    fmt.Println(user1.name, user1.age)
	fmt.Println(user2.name)

	fmt.Printf("%+v", user1)
	fmt.Printf("%+v", user2)
}
```

## Конструктор

Конструктор — это функция, которая создает и инициализирует объект определенного типа. Это помогает упростить создание объектов и сделать код более читаемым.

```go
type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

// Хорошей практикой является называть конструктор с префиксом "New"
func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    age,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vova", "Male", 23, 75, 188) // Инициализация переменной с помощью конструктора
	fmt.Printf("%+v", user1)
}

```

### Пример конструктора с указателем

Иногда полезно возвращать указатель на объект, чтобы избежать копирования больших структур:

```go
type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
    return &DumbDatabase{m: make(map[string]string)}
}
```

### Методы в структре

В структурах можно определять методы. Методы могут быть двух типов:

-   Value receivers (Значение): работают с копией объекта.
-   Pointer receivers (Ссылка): работают с оригинальным объектом.

Пример с Value receivers:

```go

package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

// Данный метод работает по значению (с копией объекта)
func (u User) printUserInfo() {
	u.name = "Новое имя" // Значение name не будет изменено, так как этот метод (Value receivers) работает по значению, а не по ссылке
	fmt.Println(u.name, u.sex, u.sex, u.weight, u.height)
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    age,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vova", "Male", 23, 75, 188)
	user2 := User{"Petya", 27, "Male", 81, 198}

	user1.printUserInfo()
	user2.printUserInfo()

	fmt.Println(user1) // Значение name будет старое - Vova
}

```

Пример с Pointer receivers:

```go

package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

// Метод работает по ссылке (с оригинальным объектом)
func (u User) printUserInfo() {
	u.name = "Новое имя" // Значение name будет изменено, так как этот метод (Pointer receivers) работает по ссылке
	fmt.Println(u.name, u.sex, u.sex, u.weight, u.height)
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    age,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vova", "Male", 23, 75, 188)
	user2 := User{"Petya", 27, "Male", 81, 198}

	user1.printUserInfo()
	user2.printUserInfo()

	fmt.Println(user1) // Значение name новое - "Новое имя"
}

```

### Дополнительный пример с Value receivers и Pointer receivers

```go

package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

// Метод с Pointer receiver - изменяет оригинальный объект
func (u *User) setName(name string) {
    u.name = name
}

// Метод с Value receiver - не изменяет оригинальный объект
func (u User) getName() string {
	return u.name
}

// Конструктор для создания нового пользователя
func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    age,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vova", "Male", 23, 75, 188)
	user2 := User{"Petya", 27, "Male", 81, 198}

    // Изменяем имя пользователя user1
    user1.setName("Serega")
	fmt.Println(user1.getName()) // Вывод: Serega
	fmt.Println(user2.getName()) // Вывод: Petya

}
```

В этом примере метод setName использует Pointer receiver, что позволяет изменять имя пользователя user1. Метод getName, использует Value receiver и возвращает текущее имя без изменения оригинального объекта.

### Создание нового типа данных для использования методов

```go
type Age int

// Метод для проверки, является ли возраст взрослым
func (a Age) isAdult bool {
    return a >= 18
}
```

### Приведение типов

В Go можно использовать приведение типов, чтобы преобразовать значения между различными типами. Например, мы можем создать структуру User, которая будет использовать тип Age:

```go
package main

import "fmt"

type Age int // Создаем новый тип Age

// Метод для проверки, является ли возраст взрослым
func (a Age) isAdult() bool {
    return a >= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

// Конструктор для создания нового пользователя
func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    Age(age), // Используем приведение типа
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vova", "Male", 11, 75, 188)
    fmt.Println(user1.age.isAdult())
}
```

Автор упомянул:

-   Структуры в Go не являются классами в традиционном смысле объектно-ориентированного программирования (ООП), как это реализовано в языках, таких как Java или C++.
-   Структуры бывают двух типов: Pointer receivers и Value receivers

## Интерфейсы

Интерфейс в Go — это тип, который определяет набор методов, которые должны быть реализованы другими типами. Интерфейсы позволяют создавать абстракции и обеспечивают гибкость в работе с различными типами.

### Чем интерфейс отличается от структуры?

1. Содержимое:

    - Структура: Содержит поля (данные) и может иметь методы для работы с этими данными.
    - Интерфейс: Содержит только методы и не имеет полей. Он описывает поведение, которое должны реализовать другие типы.

2. Назначение:

    - Структура: Используется для создания объектов с конкретным состоянием (например, информация о человеке).
    - Интерфейс: Определяет набор методов, которые различные типы могут реализовать, позволяя работать с ними через общий интерфейс.

3. Имплементация:
    - Структура: Является конкретным типом с заданными данными.
    - Интерфейс: Не требует явной реализации. Если тип реализует все методы интерфейса, он автоматически считается его реализацией.

Таким образом, структура фокусируется на данных, а интерфейс — на поведении.

```go
package main

import (
	"fmt"
	"math"
)

// Определяем интерфейс Shape, который требует реализации метода Area()
type Shape interface {
  Area() float32
}

// Определяем структуру Square, представляющую квадрат
type Square struct {
  sideLength float32
}

// Реализуем метод Area() для структуры Square
func (s Square) Area() float32 {
  return s.sideLength * s.sideLength
}

// Определяем структуру Circle, представляющую круг
type Circle struct {
  radius float32
}

// Реализуем метод Area() для структуры Circle
func (c Circle) Area() float32 {
  return c.radius * c.radius * math.Pi
}


func main() {
    // Создаем экземпляры Square и Circle
	square := Square{5}
	circle := Circle{8}

	printShapeArea(square)
	printShapeArea(circle)

	// Ошибки не будет в строчках ниже, так как все Имплеменирует пустой интерфейс
	printInterface(square)
	printInterface(circle)
	printInterface("qwerty")
	printInterface(123)
	printInterface(false)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

// Функция для работы с пустым интерфейсом
func printInterface(i interface{}) {
  // TypeSwitch позволяет определить тип переменной интерфейса
  switch value := i.(type) {
  case int:
    fmt.Println("int", value) // Если i - целое число
  case bool:
    fmt.Println("bool", value) // Если i - булевое значение
  default:
    fmt.Println("unknown type", value) // Если i - не число и не булевое
  }
  fmt.Println(i)
}

```

### Приведение типов (работа с конкретным типом)

```go
func main() {
	printInterface("qwerty") // 6
	printInterface(123) // "interface is not string"
}

func printInterface(i interface{}) {
    str, ok := i.(string) // Пробуем привести i к типу string
    if !ok {
        fmt.Println("interface is not string") // Если не строка, выводим сообщение
        return
    }
    fmt.Println(len(str)) // Если ok, печатаем длину строки
}
```

### Композиция интерфейса

```go
package main

import (
	"fmt"
	"math"
)

// Определяем интерфейс Shape, который включает в себя два других интерфейса
type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}
// Интерфейс для вычисления площади
type ShapeWithArea interface {
  Area() float32
}

// Интерфейс для вычисления периметра
type ShapeWithPerimeter interface {
  Perimeter() float32
}

// Определяем структуру Square, представляющую квадрат
type Square struct {
  sideLength float32
}

// Реализуем метод Area() для структуры Square
func (s Square) Area() float32 {
  return s.sideLength * s.sideLength
}

// Реализуем метод Perimeter() для структуры Square
func (s Square) Perimeter() float32 {
  return s.sideLength * 4
}

// Определяем структуру Circle, представляющую круг
type Circle struct {
  radius float32
}

// Реализуем метод Area() для структуры Circle
func (c Circle) Area() float32 {
  return c.radius * c.radius * math.Pi
}

// Реализуем метод Perimeter() для структуры Circle
func (c Circle) Perimeter() float32 {
  return 2 * c.radius * math.Pi
}

func main() {
	square := Square{5}
	circle := Circle{8}

	printShapeArea(square)
    // Ошибка, так как Circle не реализует метод Perimeter()
    // printShapeArea(circle) // Закомментировано, так как Circle не соответствует интерфейсу Shape
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

```

Автор упомянул:

- Интерфейсы позволяют создавать гибкие и расширяемые системы, где различные типы могут реализовывать общий набор методов. Мы также изучили, как интерфейсы могут быть скомпонованы для создания более сложных абстракций.
- В Go тип автоматически считается реализацией интерфейса, если он реализует все его методы. Это позволяет использовать подход "утиной типизации": если что-то ведет себя как утка, то это, вероятно, и есть утка.
- Использование интерфейсов и композиции интерфейсов позволяет создавать более чистый и поддерживаемый код, что является важным аспектом разработки на Go.