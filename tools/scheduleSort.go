package tools

import (
	"artpollybackend/models"
	"fmt"
	"sort"
	"strconv"
	"strings"
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

func IntSliceToString(arr []int, separator string) string {
	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = strconv.Itoa(num)
	}
	return strings.Join(strArr, separator)
}

func StringToIntSlice(s string, separator string) ([]int, error) {
	strArr := strings.Split(s, separator)
	intArr := make([]int, len(strArr))

	for i, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intArr[i] = num
	}

	return intArr, nil
}

func parseHourToTime(hour int) (time.Time, error) {
	if hour < 1 || hour > 24 {
		return time.Time{}, fmt.Errorf("hour must be between 1 and 24")
	}

	baseDate := time.Date(2024, 8, 17, 0, 0, 0, 0, time.UTC)
	return baseDate.Add(time.Duration(hour-1) * time.Hour), nil
}

func ScheduleSort(scheduleArrModels []models.Schedule) (map[string][]classInfo, error) {
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

	for _, schModel := range scheduleArrModels {
		sunHours, err := StringToIntSlice(schModel.SUN, ",")
		if err != nil {
			return nil, fmt.Errorf("error parsing SUN hours: %w", err)
		}

		monHours, err := StringToIntSlice(schModel.MON, ",")
		if err != nil {
			return nil, fmt.Errorf("error parsing MON hours: %w", err)
		}

		tueHours, err := StringToIntSlice(schModel.TUE, ",")
		if err != nil {
			return nil, fmt.Errorf("error parsing TUE hours: %w", err)
		}

		wedHours, err := StringToIntSlice(schModel.WED, ",")
		if err != nil {
			return nil, fmt.Errorf("error parsing WED hours: %w", err)
		}

		thuHours, err := StringToIntSlice(schModel.THU, ",")
		if err != nil {
			return nil, fmt.Errorf("error parsing THU hours: %w", err)
		}

		friHours, err := StringToIntSlice(schModel.FRI, ",")
		if err != nil {
			return nil, fmt.Errorf("error parsing FRI hours: %w", err)
		}

		satHours, err := StringToIntSlice(schModel.SAT, ",")
		if err != nil {
			return nil, fmt.Errorf("error parsing SAT hours: %w", err)
		}

		sch := schedule{
			ClassID:  schModel.Class.ID,
			SUN:      sunHours,
			MON:      monHours,
			TUE:      tueHours,
			WED:      wedHours,
			THU:      thuHours,
			FRI:      friHours,
			SAT:      satHours,
			Duration: schModel.Duration,
		}

		for day, hours := range map[string][]int{
			"SUN": sch.SUN,
			"MON": sch.MON,
			"TUE": sch.TUE,
			"WED": sch.WED,
			"THU": sch.THU,
			"FRI": sch.FRI,
			"SAT": sch.SAT,
		} {
			if err := processDay(day, hours, sch); err != nil {
				return nil, err
			}
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
