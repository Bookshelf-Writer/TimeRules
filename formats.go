package TimeRules

import (
	"regexp"
	"strconv"
	"strings"
)

/* Карта форматирования */
var formatMap = map[rune]func(*placeholderFormatsObj) string{
	'y': __year,
	'Y': __year,

	'n': __month,
	'm': __month,
	'F': __monthNameFull,  //##//
	'M': __monthNameShort, //##//
	't': __monthDays,

	'j': __day,
	'd': __day,
	'z': __dayPosInYear,

	'G': __hour,
	'H': __hour,
	'h': __hour,

	'I': __minute,
	'i': __minute,
	'U': __minuteFull, //##//

	'e': __timezoneName, //##//
	'Z': __timezoneOffset,
}

const (
	formatSlise            rune = 's' //Получение фрагмента 	||	 ( ranges[{ Pos от конца строки }:{ Offset размер заполнения }:{ 0 }] )
	formatSseparationBegin rune = 'm' //Целая часть от деления	||	( ranges[{ D делитель  }:{ Pos от конца строки  }:{ Offset размер заполнения }] )
	formatSseparationEnd   rune = 'n' //	Остаток от деления	||	( ranges[{ D делитель  }:{ Pos от конца строки  }:{ Offset размер заполнения }] )
)

//#################################################################################################//

// Буфер для передачи данных на методы форматирования
type placeholderFormatsObj struct {
	rule        *TimeConfigurationRulesObj //	Правила
	dateTime    *DateTimeObj               //	Метка времени
	placeholder *string                    //	Полная строка плейсхолдера
}

// placeholderFormat Создание обьекта для работы с методами форматирования
func placeholderFormat(rule *TimeConfigurationRulesObj, dateTime *DateTimeObj, placeholder *string) (retObj placeholderFormatsObj) {
	retObj.rule = rule
	retObj.dateTime = dateTime
	retObj.placeholder = placeholder

	return retObj
}

// placeholderReplace Функция, принимающая исходную строку и функцию для получения замен плейсхолдеров в формате ${...}
func placeholderReplace(input string, replacer func(string) string) string {
	re := regexp.MustCompile(`\$\{([^}]+)\}`)

	result := re.ReplaceAllStringFunc(input, func(match string) string {
		key := strings.TrimSuffix(strings.TrimPrefix(match, "${"), "}")
		return replacer(key)
	})

	return result
}

/* Вывод форматированной даты */
func (obj *TimeConfigurationRulesObj) Format(dateTime *DateTimeObj, layout string) (text string) {
	return placeholderReplace(layout, func(placeholder string) string {
		firstChar := []rune(placeholder[:1])[0]

		//Поиск совпадения по массиву
		startFunc, status := formatMap[firstChar]
		if status {
			bufObj := placeholderFormat(obj, dateTime, &placeholder)
			return startFunc(&bufObj)
		}

		return placeholder
	})
}

//.//

/* Вывод даты по умолчанию */
func (obj *TimeConfigurationRulesObj) StringDate(date DateObj) string {
	timeObj := DateTimeObj{
		Type:     "",
		Timezone: "",
		Date:     date,
		Time:     obj.Time(0, 0),
	}
	return obj.Format(&timeObj, obj.FormatsDef.Date)
}

/* Вывод времени по умолчанию */
func (obj *TimeConfigurationRulesObj) StringTime(time TimeObj) string {
	timeObj := DateTimeObj{
		Type:     "",
		Timezone: "",
		Date:     obj.Date(0, 0, 0),
		Time:     time,
	}
	return obj.Format(&timeObj, obj.FormatsDef.Time)
}

/* Вывод по умолчанию */
func (obj *TimeConfigurationRulesObj) StringDateTime(dateTime DateTimeObj) string {
	return obj.Format(&dateTime, obj.FormatsDef.Full)
}

//.//

// sliseTextNumber	Обрезка текстового числа до нужного предела с заполнением
func (obj *placeholderFormatsObj) sliseTextNumber(text string, pos uint16, size uint16) string {
	addN := false
	offset := false

	//Обработка если отрицательное
	if text[:1] == "-" {
		addN = true
		text = text[1:]
	}

	l := len(text)
	s := int(size)
	if l == 0 {
		return "ERROR {sliseTextNumber}"
	}
	if s == 0 {
		s = l
	}

	if int(pos) > l {
		pos = uint16(l)
	}

	//обрезка
	if l > s {
		n := l - int(pos)
		m := n + s

		if m > l {
			m = l
			offset = true
		}
		text = text[n:m]
	} else {
		offset = true
	}

	//Заполнение
	if offset {
		text = stringFill(text, '0', size)
	}

	if addN {
		text = "-" + text
	}

	return text
}

// parsePlaceholderParam	Получение параметров плейсхолдера
func (obj *placeholderFormatsObj) parsePlaceholderParam(placeholder string) (rune, [3]uint16) {
	param := rune(placeholder[1:2][0])
	ranges := [3]uint16{0, 0, 0}

	parseUint16 := func(s string) uint16 {
		num, err := strconv.ParseUint(s, 10, 16)
		if err == nil {
			return uint16(num)
		}

		return 0
	}

	//Разбиение указателей
	numArr := strings.Split(placeholder[2:], ":")

	//Отсечение если разметка превысила ожидания
	if len(numArr) > len(ranges) {
		return param, ranges
	}

	//Перебор указателей плейсхолдера
	for pos, num := range numArr {
		ranges[pos] = parseUint16(num)
	}

	return param, ranges
}

// transformIntFromParam Обрезка получаемого числа по параметрам
func (obj *placeholderFormatsObj) transformIntFromParam(num int64, x uint16, y uint16) string {
	stringNum := stringInt(num)

	//Обреботка только если позиция не нулевая
	if x != 0 && y == 0 && x < uint16(len(stringNum)) { //если размер не задан
		y = x
	}

	//обрезка строки
	stringNum = obj.sliseTextNumber(stringNum, x, y)

	return stringNum
}

func (obj *placeholderFormatsObj) typeFormatsInt(number int64, slise bool, separationBegin bool, separationEnd bool) string {
	bufPlaceholder := *obj.placeholder

	//Обреботка если у форматирования есть параметры
	if len(bufPlaceholder) > 1 {
		param, ranges := obj.parsePlaceholderParam(bufPlaceholder)

		/* Получение фрагмента ( ranges[{ Pos от конца строки }:{ Offset размер заполнения }:{ 0 }] ) */
		if slise && param == formatSlise {
			return obj.transformIntFromParam(number, ranges[0], ranges[1])
		}

		/* Целая часть от деления ( ranges[{ D делитель  }:{ Pos от конца строки  }:{ Offset размер заполнения }] ) */
		if separationBegin && param == formatSseparationBegin {
			if ranges[0] == 0 { //Отсекаем если ноль
				return "0"
			}

			if number < 0 {
				number *= -1
			}

			d := int64(ranges[0])
			if d > number { //отсекаем если 100% не будет целой части
				return "0"
			}

			number = number / d
			return obj.transformIntFromParam(number, ranges[1], ranges[2])
		}

		/* Остаток от деления ( ranges[{ D делитель  }:{ Pos от конца строки  }:{ Offset размер заполнения }] ) */
		if separationEnd && param == formatSseparationEnd {
			if ranges[0] == 0 { //Отсекаем если ноль
				return "0"
			}

			if number < 0 {
				number *= -1
			}

			num := strconv.FormatFloat(float64(number)/float64(ranges[0]), 'f', -1, 64)
			bufArr := strings.Split(num, ".")
			if len(bufArr) != 2 {
				number = 0
			} else {
				number, _ = strconv.ParseInt(bufArr[1], 10, 64)
			}

			return obj.transformIntFromParam(number, ranges[1], ranges[2])
		}
	}

	return stringInt(number)
}

//#################################################################################################//

// __year	[Y] 	Полное числовое представление года					Примеры: -55, 787, 1999, 2003, 10191
func __year(obj *placeholderFormatsObj) string {
	return obj.typeFormatsInt(obj.dateTime.Date.Year, true, true, true)
}

// __month	[n] 	Порядковый номер месяца без ведущего нуля 			От 1 до 12
func __month(obj *placeholderFormatsObj) string {
	return obj.typeFormatsInt(int64(obj.dateTime.Date.Month), true, true, true)
}

// __monthNameFull	[`F`] 	Полное наименование месяца, 						например, January или March
func __monthNameFull(obj *placeholderFormatsObj) string {
	return obj.rule.MonthFull(&obj.dateTime.Date)
}

// __monthNameShort	[M] 	Сокращённое наименование месяца, 3 символа 			От Jan до Dec
func __monthNameShort(obj *placeholderFormatsObj) string {
	return obj.rule.MonthShort(&obj.dateTime.Date)
}

// __monthDays	[t] 	Количество дней в указанном месяце 					От 28 до 31
func __monthDays(obj *placeholderFormatsObj) string {
	return obj.typeFormatsInt(int64(obj.rule.Month[obj.dateTime.Date.Month-1].Days), true, true, true)
}

// __day	[j] 	День месяца без ведущего нуля 						От 1 до 31
func __day(obj *placeholderFormatsObj) string {
	return obj.typeFormatsInt(int64(obj.dateTime.Date.Day), true, true, true)
}

// __dayPosInYear	[z] 	Порядковый номер дня в году (начиная с 0) 			От 0 до 365
func __dayPosInYear(obj *placeholderFormatsObj) string {
	sum := uint64(0)
	for i := uint16(0); i < obj.dateTime.Date.Month-1; i++ {
		sum += uint64(obj.rule.Month[i].Days)
	}
	return obj.typeFormatsInt(int64(sum+uint64(obj.dateTime.Date.Day)-1), true, true, true)
}

//.//

// __hour	[G] 	Часы без ведущего нуля 								От 0 до 23
func __hour(obj *placeholderFormatsObj) string {
	return obj.typeFormatsInt(int64(obj.dateTime.Time.Hour), true, true, true)
}

// __minute	[I] 	Минуты без ведущего нуля 							От 0 до 59
func __minute(obj *placeholderFormatsObj) string {
	return obj.typeFormatsInt(int64(obj.dateTime.Time.Min), true, true, true)
}

// __minuteFull	[U] 	Кол-во минут прошедших от точки отчета без смешения	Примеры: -55, 787, 10191
func __minuteFull(obj *placeholderFormatsObj) string {
	bufDate, err := obj.rule.LocationSet(*obj.dateTime, "")
	if err != nil {
		return "ERROR: " + err.Error()
	}
	buf := obj.rule.DateTimeToMinutes(bufDate.Date, bufDate.Time)
	return buf.String()
}

//.//

// __timezoneName	[e] 	Идентификатор часового пояса 						Примеры: UTC, GMT, Atlantic/Azores
func __timezoneName(obj *placeholderFormatsObj) string {
	return obj.dateTime.Timezone
}

// __timezoneOffset	[Z] 	Смещение часового пояса в минутах					Примеры: -55, 787, 10191
func __timezoneOffset(obj *placeholderFormatsObj) string {
	if len(obj.dateTime.Timezone) > 0 {
		offset := obj.typeFormatsInt(obj.rule.Timezones[obj.dateTime.Timezone], true, true, true)
		if len(offset) > 0 && offset[0] != '-' {
			offset = "+" + offset
		}
		return offset
	} else {
		return "0"
	}
}

//#################################################################################################//

/* Сокращенное название месяца */
func (obj *TimeConfigurationRulesObj) MonthShort(date *DateObj) string {
	if date.Month > uint16(len(obj.Month)) {
		return stringUint(uint64(date.Month))
	}
	return obj.Month[date.Month-1].ShortName
}

/* Полное название месяца человеко-понятно */
func (obj *TimeConfigurationRulesObj) MonthFull(date *DateObj) string {
	if date.Month > uint16(len(obj.Month)) {
		return stringUint(uint64(date.Month))
	}
	return obj.Month[date.Month-1].FullName
}
