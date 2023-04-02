package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// TwoDecimalFloat64 ..
func TwoDecimalFloat64(value float64) float64 {
	return float64(int64(value*100)) / 100
}

// ParseParamToInt ..
func ParseParamToInt(param string) int {
	if param != "" {
		result, err := strconv.Atoi(param)
		if err != nil {
			return 0
		}

		return result
	}

	return 0
}

// ParseParamToArrayInt ..
func ParseParamToArrayInt(param string) []int {
	arrayInt := []int{}
	if err := json.Unmarshal([]byte(param), &arrayInt); err != nil {
		return arrayInt
	}

	return arrayInt
}

// ParseParamToFloat64 ..
func ParseParamToFloat64(param string) float64 {
	if param != "" {
		cost, err := strconv.ParseFloat(param, 2)
		if err != nil {
			return 0
		}

		return cost
	}

	return 0
}

// ParseParamToBoolean ..
func ParseParamToBoolean(param string) bool {
	if param != "" {
		boolean, err := strconv.ParseBool(param)
		if err != nil {
			return false
		}

		return boolean
	}

	return false
}

// ParseParamToTimeString ..
func ParseParamToTimeString(timeParam string, hhmmssFormat string) string {
	if timeParam != "" {
		searchTimeParam, err := time.Parse("02/01/2006 15:04:05-07", fmt.Sprintf("%s %s", timeParam, hhmmssFormat))
		if err != nil {
			return ""
		}
		return searchTimeParam.Format("2006-01-02 15:04:05-07")
	}

	return ""
}

// NotNullValidator : Using for validate interface type ..
func NotNullValidator(values []interface{}) (interface{}, error) {
	if len(values) == 0 {
		return nil, nil
	}

	for _, value := range values {
		switch value.(type) {
		case int:
			if value.(int) == 0 {
				return value.(int), fmt.Errorf("null or empty: int")
			}
		case float64:
			if value.(float64) == 0 {
				return value.(float64), fmt.Errorf("null or empty: float64")
			}
		case string:
			if value.(string) == "" {
				return value.(string), fmt.Errorf("null or empty: string")
			}
		case bool:
			return nil, nil
		case []int:
			validateInts := value.([]int)
			if len(validateInts) == 0 {
				return value.(string), fmt.Errorf("null or empty: []int")
			}

			for _, validateInt := range validateInts {
				if validateInt == 0 {
					return validateInt, fmt.Errorf("null or empty: int")
				}
			}
		default:
			return value, fmt.Errorf("null or empty: unknown")
		}
	}

	return nil, nil
}

func ParseParamToTime(timeParam string) time.Time {
	if timeParam != "" {
		timeParamRes, err := time.Parse("02/01/2006 15:04:05-07", fmt.Sprintf("%s", timeParam))

		if err != nil {
			return time.Time{}
		}

		return timeParamRes
	}

	return time.Time{}
}

func ParseParamToTimeV2(format string, timeParam string) time.Time {
	if timeParam != "" {
		timeParamRes, err := time.Parse(format, fmt.Sprintf("%s", timeParam))

		if err != nil {
			return time.Time{}
		}

		return timeParamRes
	}

	return time.Time{}
}

func ParseParamsToTime(dateParam string, timeParam string) time.Time {
	if timeParam != "" {
		timeParamRes, err := time.Parse("02/01/2006 15:04:05-07", fmt.Sprintf("%s %s", dateParam, timeParam))
		if err != nil {
			return time.Time{}
		}
		return timeParamRes
	}

	return time.Time{}
}

// Used for check is time.Time is in range or not
func IsDateInRange(dateToCheck time.Time, startDate time.Time, endDate time.Time) bool {
	return dateToCheck.After(startDate) && dateToCheck.Before(endDate)
}
