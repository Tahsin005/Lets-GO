func dayOfTheWeek(day int, month int, year int) string {
    months := []int{1, 4, 4, 0, 2, 5, 0, 3, 6, 1, 4, 6}

    yy := year % 100
    total := yy + yy / 4 + day + months[month - 1]
 
    century := year / 100
    
    switch century {
    case 19:
        // +0
    case 20:
        total--
    case 21:
        total += 4
    }

    if isLeap(year) && (month == 1 || month == 2) {
        total--
    }

    days := []string{
        "Saturday",
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
    }

    return days[total % 7]
}

func isLeap(year int) bool {
    return year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)
}
