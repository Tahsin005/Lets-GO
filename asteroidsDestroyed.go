func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Ints(asteroids)

    currentMass := int64(mass)

    for _, asteroid := range asteroids {
        asteroidMass := int64(asteroid)

        if currentMass < asteroidMass {
            return false
        }

        currentMass += asteroidMass
    }

    return true
}
