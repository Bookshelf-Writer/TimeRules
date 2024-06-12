![Fork GitHub Release](https://img.shields.io/github/v/release/Bookshelf-Writer/TimeRules)
![Tests](https://github.com/Bookshelf-Writer/TimeRules/actions/workflows/go-test.yml/badge.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/Bookshelf-Writer/TimeRules)](https://goreportcard.com/report/github.com/Bookshelf-Writer/TimeRules)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/Bookshelf-Writer/TimeRules?color=orange)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Bookshelf-Writer/TimeRules?color=green)
![GitHub repo size](https://img.shields.io/github/repo-size/Bookshelf-Writer/TimeRules)

## Модуль поддержки произвольных временных меток
*(привязка к правилам формирования времени)*
 
#### Установка:
```bash
go mod init github.com/Bookshelf-Writer/TimeRules
```

---

#### Внешние модули
- `github.com/mailru/easyjson`  - Внедрен для ускорения marshal/unmarshal структур JSON
- `golang.org/x/crypto`         - Используется для генерации подписей в blake2b

---

#### Особенности:
- Позволяет полноценно работать с датами, пользуясь пользовательскими правилами времени
- Минимальная единица времени **минута**
- Минималистичная структура для описывания правил времени
- Позволяет описать абсолютно уникальный случай времени, не ограниченный стандартным 60m/24h/365d
- Позволяет описать месяцы с разным количеством дней
- Поддерживает пользовательские таймзоны
- Есть шаблонизатор позволяющий описать как именно отдавать дату строкой (дата, время, полная запись)

---

#### Функционал:

1. Работа с правилами времени
   - Чтение правил из файла (только поддерживаемые форматы)
   - Проверка правила на корректность
   - Создание файла правила (только поддерживаемые форматы)
   - Работа с правилами в формате строк и массива байт
   
2. Работа со временем
   - Получение встроенного правила времени (классика с 60m/24h/365d)
   - Работа с датами
   - Проверка времени на соответствие правилу
   - Перевод времени между разными часовыми поясами
   - Возможность перевода между разными временными зонами (правила времени) 
   - Вывод времени строкой согласно форматированию
   - Обработка произвольного форматирования для получения строки из даты
   - _более подробно смотрите в описании методов_

---

#### Поддерживаемые форматы:
- `json`    - Человеко-понятный формат описания правил времени
- `tvr`     - Машинно-понятный формат описания времени. Оптимизированный для хранения и отдачи. Имеется встроенная проверка целостности.

---

#### Физические ограничения:
|                                   | **MIN**                  | **MAX**                 |
|-----------------------------------|--------------------------|-------------------------|
| Минут в часе                      | 1                        | 65534                   |
| Часов в дне                       | 1                        | 65534                   |
| Дней в месяце                     | 1                        | 65534                   |
| Месяцев в году                    | 1                        | 65534                   |
| Предел года                       | -9&times;10<sup>18</sup> | 9&times;10<sup>18</sup> |
| Предел смещения таймзоны          | -9&times;10<sup>18</sup> | 9&times;10<sup>18</sup> |
| Короткое название месяца (длинна) | 1                        | 4                       |
| Полное название месяца (длинна)   | 1                        | 32                      |
| Ключ таймзоны (длинна)            | 1                        | 64                      |


---

---

### Правила времени

#### Структура JSON-файла
```json
{
  "name": "varchar",
  "description": "text",
  "inf": {
    "ver": "1.0.0",
    "creator": "creator name"
  },

  "format": {
    "date": "${Y}/${m}/${d}",
    "time": "${h}:${i}",
    "full": "${Y}/${m}/${d} ${h}:${i} ${Z}"
  },
  "year": {
    "min": -10000,
    "max": 10000
  },
  "maxHour": 10,
  "maxMin": 20,

  "month": [
    {"fullName": "month 1", "shortName": "X", "days": 10},
    {"fullName": "month 2", "shortName": "Y", "days": 20},
    {"fullName": "month 3", "shortName": "Z", "days": 30}
  ],

  "timezones": {
    "name1": 0,
    "name2": 60,
    "name3": -60
  }
}
```

- Все поля обязательные
  - Перечисление в **month** должно быть хотя-бы одно.
  - Указатель в **timezones** должен быть хотя-бы один.
- **name** это название правила. Оно должно быть уникальным. Длинна минимум 3, максимум 32. Правило записи `^[a-zA-Z0-9 \-.:;_+@]*$`
- **description** описывает кратко правило. Правило записи `^[^'"\\]+$`
- **inf** описывает системную информацию.
  - **ver** это номерная версия правила. Не влияет на функционал, необходимо для "визуального" согласования разных редакций одного правила. Правило записи `^[0-9\.]*$`
  - **creator** это имя\логин\ник\контакт создателя\редактора правила. Не влияет на функционал.  Правило записи `^[^'"\\%&,/]+$`
- **format** описывает шаблон формирования строчного представления даты. Подробнее о правилах форматирования ниже.
  - **date** правило описания даты.
  - **time** правило описания времени.
  - **full** правило описания даты-времени.
- **year** допустимые пределы года.
- **maxHour** максимальное число для _часов_. Для 24-часов максимальным будет 23.
- **maxMin** максимальное число для _минут_. Для 60-минутного часа максимальным будет 59.
- **month** массив описания месяца. Каждая запись описывает конкретный месяц. Порядок важен. Количество дней в году насчитывается суммированием **days** в каждом месяце.
  - **fullName** полное название месяца. Правило записи `^[^'"\\%&,/]+$`
  - **shortName** краткое название месяца. Правило записи `^[a-zA-Z0-9]*$`
  - **days** максимальное число для _дней_ в текущем месяце.
- **timezones** описание возможных смещений часовых поясов. Должна быть хотя-бы одна запись. Смещение указывается целым числом в минутах.

---

#### Форматирование

##### Ключи
| **KEY**  |                                                     | **MATH** |
|----------|-----------------------------------------------------|----------|
| `y`      | Числовое представление года                         | ✅        |
| `Y`      | `y`                                                 | ✅        |
| `n`      | Порядковый номер месяца                             | ✅        |
| `m`      | `n`                                                 | ✅        |
| `F`      | Полное наименование месяца                          | ❌        |
| `M`      | Сокращённое наименование месяца                     | ❌        |
| `t`      | Количество дней в указанном месяце                  | ✅        |
| `j`      | День месяца                                         | ✅        |
| `d`      | `j`                                                 | ✅        |
| `z`      | Порядковый номер дня в году                         | ✅        |
| `G`      | Часы                                                | ✅        |
| `H`      | `G`                                                 | ✅        |
| `h`      | `G`                                                 | ✅        |
| `I`      | Минуты                                              | ✅        |
| `i`      | `I`                                                 | ✅        |
| `U`      | Кол-во минут прошедших от точки отчета без смешения | ❌        |
| `e`      | Идентификатор часового пояса                        | ❌        |
| `Z`      | Смещение часового пояса                             | ✅        |


##### Математические операции
| **SYMBOL** |                        | **X1**                   | **X2**                  | **X3**            |
|------------|------------------------|--------------------------|-------------------------|-------------------|
| `s`        | Получение фрагмента    | Позиция от конца строки  | Размер заполнения       | 0                 |
| `m`        | Целая часть от деления | Делитель                 | Позиция от конца строки | Размер заполнения |
| `n`        | Остаток от деления     | Делитель                 | Позиция от конца строки | Размер заполнения |

- В случае если `Позиция от конца строки` больше фактической длинны строки то сначала добавляются недостающие нули
- `Размер заполнения` указывает сколько символов показывать, начиная точки отсчета. Используется с `Позиция от конца строки` для получения фрагмента числа. Примеры ниже.


##### Примеры
| **INPUT**                      | **OUTPUT**  |                                                                       |
|--------------------------------|-------------|-----------------------------------------------------------------------|
| `${Y}`                         | 2024        | Полное числовое представление года                                    |
| `{Ys2}`                        | 24          | Обрезаный год                                                         |
| `${H}`                         | 8           | Часы без ведущего нуля                                                |
| `${Hs2}:${Is2}`                | 08:04       | Формалое представление времени с ведущим нулем                        |
| `${e}`                         | Etc/GMT+1   | Идентификатор часового пояса                                          |
| `${Y}-${M}-${ds:2}`            | 2024-FEB-03 | Формальное представление даты с текстовым месяцев и ведушим нулем дня |
| `{Zm60}`                       | -1          | Формальное представление смешения в часах                             |
| `${m}.${Yn1000::3}.M${Ym1000}` | 113.004.M40 | Формальное представление даты в Warhammer-40K                         |


---

---

### Mirrors

- https://git.bookshelf-writer.fun/Bookshelf-Writer/TimeRules

