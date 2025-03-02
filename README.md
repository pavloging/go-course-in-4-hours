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
-   _ - это пропуск значения.
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
    return fmt.Sprintf("Привет, %s! Тебе %d лет", name, age) // Для возврата значения в функции используем return
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
