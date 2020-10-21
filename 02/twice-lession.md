# Синтаксис

* Числа
  - целые (uint8, int8, int16, uint16 ,...)
  - вещественные (float32, float64)
  - комплексные (complex64, complex124)

* Строки ("")
* Булевые (false)
* Встроенные псевдонимы (byte -> uint8, rune -> int32)

Явные преобразования типов

# Встроенные типы данных
## Значения
- массивы
- структуры

т.е передаются объекты

## Ссылочные
- срезы
- мапы
- каналы
- функции
- интерфейсы

# Переменные
`var var_name type` - когда инициализируем дальше по коду

`var_name` - когда инициализируем здесь и сейчас

# Указатели и ссылки
В go есть указатели на переменную определенного типа, но нет универсального указателя.
Нет адресной арифметики.

Указатели на объекты произвольного типа и адресная арифметика есть в пакете unsafe, но пользоваться им не стоит почти никогда,
как и советует Rob Pike.

## Слайсы
 - указатель на массив
 - текущий размер
 - емкость

## Композиция
 - возможность встроить один структурный тип в другой с сохранением 