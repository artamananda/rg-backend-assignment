package main

import (
	"fmt"
	"strconv"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	var res string

	typ1, _ := time.(string)
	typ2, _ := time.([]int)
	typ3, _ := time.(map[string]int)
	typ4, _ := time.(Time)

	if len(typ1) == 5 {
		if typ1[0] >= '1' && typ1[1] > '2' {
			strTmp := string(typ1[0]) + string(typ1[1])
			iTmp, _ := strconv.Atoi(strTmp)
			iTmp -= 12
			if iTmp < 10 {
				res += "0" + fmt.Sprintf("%d", iTmp) + ":" + string(typ1[3]) + string(typ1[4]) + " PM"
			} else {
				res += fmt.Sprintf("%d", iTmp) + ":" + string(typ1[3]) + string(typ1[4]) + " PM"
			}
		} else if typ1[0] == '1' && typ1[1] == '2' {
			res += typ1 + " PM"
		} else {
			strTmp := string(typ1[0]) + string(typ1[1])
			iTmp, _ := strconv.Atoi(strTmp)
			if iTmp < 10 {
				res += "0" + fmt.Sprintf("%d", iTmp) + ":" + string(typ1[3]) + string(typ1[4]) + " AM"
			} else {
				res += fmt.Sprintf("%d", iTmp) + ":" + string(typ1[3]) + string(typ1[4]) + " AM"
			}
		}
	} else if len(typ2) == 2 {
		hours := typ2[0]
		minutes := typ2[1]
		if hours >= 0 && hours <= 23 && minutes >= 0 && minutes <= 59 {
			if hours > 12 {
				if hours-12 < 10 {
					res += "0"
				}
				if minutes < 10 {
					res += fmt.Sprintf("%d", hours-12) + ":0" + fmt.Sprintf("%d", minutes) + " PM"
				} else {
					res += fmt.Sprintf("%d", hours-12) + ":" + fmt.Sprintf("%d", minutes) + " PM"
				}

			} else if hours == 12 {
				if minutes < 10 {
					res += fmt.Sprintf("%d", hours) + ":0" + fmt.Sprintf("%d", minutes) + " PM"
				} else {
					res += fmt.Sprintf("%d", hours) + ":" + fmt.Sprintf("%d", minutes) + " PM"
				}
			} else {
				if hours < 10 {
					res += "0"
				}
				if minutes < 10 {
					res += fmt.Sprintf("%d", hours) + ":0" + fmt.Sprintf("%d", minutes) + " AM"
				} else {
					res += fmt.Sprintf("%d", hours) + ":" + fmt.Sprintf("%d", minutes) + " AM"
				}
			}
		}
	} else if len(typ3) == 2 {
		hours, ok := typ3["hour"]
		minutes, ok2 := typ3["minute"]
		if ok && ok2 && hours >= 0 && hours <= 23 && minutes >= 0 && minutes <= 59 {
			if hours > 12 {
				if hours-12 < 10 {
					res += "0"
				}
				if minutes < 10 {
					res += fmt.Sprintf("%d", hours-12) + ":0" + fmt.Sprintf("%d", minutes) + " PM"
				} else {
					res += fmt.Sprintf("%d", hours-12) + ":" + fmt.Sprintf("%d", minutes) + " PM"
				}

			} else if hours == 12 {
				if minutes < 10 {
					res += fmt.Sprintf("%d", hours) + ":0" + fmt.Sprintf("%d", minutes) + " PM"
				} else {
					res += fmt.Sprintf("%d", hours) + ":" + fmt.Sprintf("%d", minutes) + " PM"
				}
			} else {
				if hours < 10 {
					res += "0"
				}
				if minutes < 10 {
					res += fmt.Sprintf("%d", hours) + ":0" + fmt.Sprintf("%d", minutes) + " AM"
				} else {
					res += fmt.Sprintf("%d", hours) + ":" + fmt.Sprintf("%d", minutes) + " AM"
				}
			}
		} else {
			return "Invalid input"
		}
	} else if typ4.Hour > 0 && typ4.Hour <= 23 && typ4.Minute >= 0 && typ4.Minute <= 59 {
		hours := typ4.Hour
		minutes := typ4.Minute
		if hours >= 0 && hours <= 23 && minutes >= 0 && minutes <= 59 {
			if hours > 12 {
				if hours-12 < 10 {
					res += "0"
				}
				if minutes < 10 {
					res += fmt.Sprintf("%d", hours-12) + ":0" + fmt.Sprintf("%d", minutes) + " PM"
				} else {
					res += fmt.Sprintf("%d", hours-12) + ":" + fmt.Sprintf("%d", minutes) + " PM"
				}

			} else if hours == 12 {
				if minutes < 10 {
					res += fmt.Sprintf("%d", hours) + ":0" + fmt.Sprintf("%d", minutes) + " PM"
				} else {
					res += fmt.Sprintf("%d", hours) + ":" + fmt.Sprintf("%d", minutes) + " PM"
				}
			} else {
				if hours < 10 {
					res += "0"
				}
				if minutes < 10 {
					res += fmt.Sprintf("%d", hours) + ":0" + fmt.Sprintf("%d", minutes) + " AM"
				} else {
					res += fmt.Sprintf("%d", hours) + ":" + fmt.Sprintf("%d", minutes) + " AM"
				}
			}
		}

	} else {
		res = "Invalid input"
	}

	return res // TODO: replace this
}
func main() {
	fmt.Println(ChangeToStandartTime("2300"))
	fmt.Println(ChangeToStandartTime([]int{23}))
	fmt.Println(ChangeToStandartTime(map[string]int{"minute": 11, "second": 11}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
