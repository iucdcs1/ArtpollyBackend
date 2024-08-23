package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type schedule struct {
	ClassID  uint  `json:"id"`
	SUN      []int `json:"sun"`
	MON      []int `json:"mon"`
	TUE      []int `json:"tue"`
	WED      []int `json:"wed"`
	THU      []int `json:"thu"`
	FRI      []int `json:"fri"`
	SAT      []int `json:"sat"`
	Duration int   `json:"duration"`
}

type classInfo struct {
	ClassID  uint `json:"class_id"`
	Hour     int  `json:"hour"`
	Duration int  `json:"duration"`
}

func parseHourToTime(hour int) (time.Time, error) {
	if hour < 1 || hour > 24 {
		return time.Time{}, fmt.Errorf("hour must be between 1 and 24")
	}

	baseDate := time.Date(2024, 8, 17, 0, 0, 0, 0, time.UTC)
	return baseDate.Add(time.Duration(hour-1) * time.Hour), nil
}

func ScheduleSort(scheduleArr []schedule) (map[string][]classInfo, error) {
	daySchedule := map[string][]classInfo{
		"SUN": {},
		"MON": {},
		"TUE": {},
		"WED": {},
		"THU": {},
		"FRI": {},
		"SAT": {},
	}

	processDay := func(day string, hours []int, sch schedule) error {
		for _, hour := range hours {
			parsedTime, err := parseHourToTime(hour)
			if err != nil {
				return err
			}
			daySchedule[day] = append(daySchedule[day], classInfo{
				ClassID:  sch.ClassID,
				Hour:     parsedTime.Hour() + 1, // Store hour in the 1-24 range
				Duration: sch.Duration,
			})
		}
		return nil
	}

	for _, sch := range scheduleArr {
		if err := processDay("SUN", sch.SUN, sch); err != nil {
			return nil, err
		}
		if err := processDay("MON", sch.MON, sch); err != nil {
			return nil, err
		}
		if err := processDay("TUE", sch.TUE, sch); err != nil {
			return nil, err
		}
		if err := processDay("WED", sch.WED, sch); err != nil {
			return nil, err
		}
		if err := processDay("THU", sch.THU, sch); err != nil {
			return nil, err
		}
		if err := processDay("FRI", sch.FRI, sch); err != nil {
			return nil, err
		}
		if err := processDay("SAT", sch.SAT, sch); err != nil {
			return nil, err
		}
	}

	for day := range daySchedule {
		sort.SliceStable(daySchedule[day], func(i, j int) bool {
			if daySchedule[day][i].Hour == daySchedule[day][j].Hour {
				return daySchedule[day][i].Duration < daySchedule[day][j].Duration
			}
			return daySchedule[day][i].Hour < daySchedule[day][j].Hour
		})
	}

	return daySchedule, nil
}

func main() {
	scheduleArr := []schedule{
		{
			ClassID:  1,
			SUN:      []int{3, 5},
			MON:      []int{2, 8},
			Duration: 60,
		},
		{
			ClassID:  2,
			SUN:      []int{2, 6},
			MON:      []int{1, 7},
			Duration: 45,
		},
	}

	sortedSchedules, err := ScheduleSort(scheduleArr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonOutput, err := json.MarshalIndent(sortedSchedules, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonOutput))
}
