func convertTime(current string, correct string) int {
    c := 0

    t1 := strings.Split(current, ":")
    t2 := strings.Split(correct, ":")

    h1, _ := strconv.Atoi(t1[0])
    h2, _ := strconv.Atoi(t2[0])

    m1, _ := strconv.Atoi(t1[1])
    m2, _ := strconv.Atoi(t2[1])

    minutes := (h2 - h1) * 60 + (m2 - m1)

    total := 0

    for total < minutes {
        if total + 60 <= minutes {
            total += 60
        } else if total + 15 <= minutes {
            total += 15
        } else if total + 5 <= minutes {
            total += 5
        } else {
            total += 1
        }

        c++
    }

    return c
}
