package utils

type PeriodStats struct {
	WorkingTime     float64 `json:"workingTime"`
	DistractionTime float64 `json:"distractionTime"`
}

type StatsResult struct {
	WorkingTime             int                  `json:"workingTime"` // in seconds
	DistractionTime         int                  `json:"distractionTime"`
	DailyAvgWorkingTime     *int                 `json:"dailyAvgWorkingTime"`
	DailyAvgDistractionTime *int                 `json:"dailyAvgDistractionTime"`
	RealStartDate           string               `json:"realStartDate"`
	RealFinishDate          string               `json:"realFinishDate"`
	NWeeksOfMonth           int                  `json:"nWeeksOfMonth"`
	YearInfo                *map[int]PeriodStats `json:"yearInfo"`
	MonthInfo               *map[int]PeriodStats `json:"monthInfo"`
	WeekInfo                *map[int]PeriodStats `json:"weekInfo"`
	Tags                    *[]string            `json:"tags"`
}
