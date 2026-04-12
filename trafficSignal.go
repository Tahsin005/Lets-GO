func trafficSignal(timer int) string {
    if timer > 30 && timer <= 90 {
        return "Red"
    } else if timer == 30 {
        return "Orange"
    } else if timer == 0 {
        return "Green"
    }

    return "Invalid"
}
