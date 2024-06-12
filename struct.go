package TimeRules

import "encoding/gob"

// SystemInfoObj	Системная информация по объекту
type SystemInfoObj struct {
	Ver     string `json:"ver"`     //	Версия файла. Для отслеживания изменений по файлу
	Creator string `json:"creator"` //	Имя\логин\иное создателя\редактора
}

type FormatsInfoObj struct {
	Date string `json:"date"` //	Представление даты
	Time string `json:"time"` //	Представление времени
	Full string `json:"full"` //	Представление временного объекта
}

// YearLimitsObj	Лимит времени в годах
type YearLimitsObj struct {
	Min int64 `json:"min"` //	Минимально допустимый год
	Max int64 `json:"max"` //	Максимально допустимый год
}

// MonthObj	Обьект месяца
type MonthObj struct {
	FullName  string `json:"fullName"`  //	Полноное название месяца
	ShortName string `json:"shortName"` //	Сокращенное название месяца латиницей
	Days      uint16 `json:"days"`      //	Количество дней в месяце
}

/* Базовая стукрута конфигурации времени */
type TimeConfigurationRulesObj struct {
	Name        string `json:"name"`        //< Уникальное название правила
	Description string `json:"description"` //< Описание правила

	INF        SystemInfoObj  `json:"inf"`    //<	Системная информация по файлу
	FormatsDef FormatsInfoObj `json:"format"` //<	Форматирование строчного представления даты по умолчанию

	Year    YearLimitsObj `json:"year"`    //< Годовые ограничения
	MaxHour uint16        `json:"maxHour"` //< Максимальное количество часов в дне
	MaxMin  uint16        `json:"maxMin"`  //< Максимальное количество минут в часе

	Month      []MonthObj       `json:"month"`      //< Массив месяцев. Порядок важен
	DaysInYear uint64           `json:"daysInYear"` //<	Дней в году (автоматически генерируется из месяцев)
	Timezones  map[string]int64 `json:"timezones"`  //< Уникальные часовые зоны со смещением в минутах
}

//.//

func init() {
	gob.Register(TimeConfigurationRulesObj{})
}
