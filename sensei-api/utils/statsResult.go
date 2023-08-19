package utils

type ActivityStatElement struct {
	ActivityName string `json:"activityName"`
	ActivityIcon string `json:"activityIcon"`
	NTimes       int    `json:"nTimes"`
}

type StatsResult struct {
	ScheduledDans       int                    `json:"scheduledDans"`
	CompletedDans       int                    `json:"completedDans"`
	CompletedPercentage float32                `json:"completedPercentage"`
	DailyAvgScheduled   interface{}            `json:"dailyAvgScheduled"`
	DailyAvgCompleted   interface{}            `json:"dailyAvgCompleted"`
	RealStartDate       string                 `json:"realStartDate"`
	RealFinishDate      string                 `json:"realFinishDate"`
	NWeeksOfMonth       int                    `json:"nWeeksOfMonth"`
	ActivityInfo        *[]ActivityStatElement `json:"activityInfo"`
}
