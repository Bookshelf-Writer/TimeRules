package TimeRules

import "strconv"

// DateObj	Группа даты
type DateObj struct {
	Year  int64  `json:"year"`  //< Год
	Month uint16 `json:"month"` //< Месяц
	Day   uint16 `json:"day"`   //< День
}

// TimeObj	Группа времени
type TimeObj struct {
	Hour uint16 `json:"hour"` //< Часы
	Min  uint16 `json:"min"`  //< Минуты
}

/* Структура универсальной точки времени */
type DateTimeObj struct {
	Type string `json:"type"` //< Тип календаря

	Date     DateObj `json:"date"`     //<	Дата
	Time     TimeObj `json:"time"`     //<	Время
	Timezone string  `json:"timezone"` //<	Таймзона для расчета смещения
}

//#################################################################################################//

func (obj DateObj) String() string {
	return strconv.FormatInt(obj.Year, 10) + "-" + strconv.FormatUint(uint64(obj.Month), 10) + "-" + strconv.FormatUint(uint64(obj.Day), 10)
}

func (obj TimeObj) String() string {
	return strconv.FormatUint(uint64(obj.Hour), 10) + ":" + strconv.FormatUint(uint64(obj.Min), 10)
}

func (obj DateTimeObj) String() string {
	return obj.Date.String() + " " + obj.Time.String() + " " + obj.Timezone
}
