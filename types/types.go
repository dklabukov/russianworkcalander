package types

type Calendar struct {
	Year                 string `json:"Год/Месяц"`
	Jan                  string `json:"Январь"`
	Feb                  string `json:"Февраль"`
	Mar                  string `json:"Март"`
	Apr                  string `json:"Апрель"`
	May                  string `json:"Май"`
	Jun                  string `json:"Июнь"`
	Jul                  string `json:"Июль"`
	Aug                  string `json:"Август"`
	Sep                  string `json:"Сентябрь"`
	Oct                  string `json:"Октябрь"`
	Nov                  string `json:"Ноябрь"`
	Dec                  string `json:"Декабрь"`
	TotalWorkDays        string `json:"Всего рабочих дней"`
	TotalWeekendDays     string `json:"Всего праздничных и выходных дней"`
	TotalWorkHoursWith40 string `json:"Количество рабочих часов при 40-часовой рабочей неделе"`
	TotalWorkHoursWith36 string `json:"Количество рабочих часов при 36-часовой рабочей неделе"`
	TotalWorkHoursWith24 string `json:"Количество рабочих часов при 24-часовой рабочей неделе"`
}
