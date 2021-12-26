package week3

type target struct {
	X, Y bounds
}

type bounds struct {
	Min, Max int
}

var (
	trenchTarget = target{
		X: bounds{179, 201},
		Y: bounds{-109, -63},
	}

	testTrenchTarget = target{
		X: bounds{20, 30},
		Y: bounds{-10, -5},
	}
)

func trickShot(t target) (int, int) {
	highest := 0
	y0 := 0
	i := 1
	for ; i < 1000000; i++ {
		a := apex(i)
		minDist := a - t.Y.Max
		maxDist := a - t.Y.Min

		for j := i; ; j++ {
			dist := apex(j)
			if dist > maxDist {
				break
			}
			if dist >= minDist && dist <= maxDist {
				highest = a
				y0 = i
				break
			}
		}
	}
	return highest, y0
}

func trickShotEveryAngle() int {
	t := trenchTarget

	totalShots := (t.X.Max - t.X.Min + 1) * (t.Y.Max - t.Y.Min + 1)
	_, maxY := trickShot(t)
	xVelStop := 1
	minX := -1
	for {
		if apex(xVelStop) >= t.X.Min && apex(xVelStop) <= t.X.Max {
			if minX == -1 {
				minX = xVelStop
			}
		}
		if apex(xVelStop) > t.X.Max {
			break
		}
		xVelStop++
	}

	// The furthest down we can shoot and be in the target at time=2
	minY := (t.Y.Min + 1) / 2
	for yVel := minY; yVel <= maxY; yVel++ {
		times := getTimes(yVel, t)
		for xVel := minX; xVel <= (t.X.Max+1)/2; xVel++ {
			for _, time := range times {
				xPos := apex(xVel) - apex(xVel-time)
				if xVel-time < 0 {
					xPos = apex(xVel)
				}
				if xPos >= t.X.Min && xPos <= t.X.Max {
					totalShots++
					break
				}
			}
		}
	}

	return totalShots
}

func getTimes(yVel int, t target) []int {
	currPos := yVel
	var times []int
	for i := 1; currPos >= t.Y.Min; i++ {
		if currPos <= t.Y.Max {
			times = append(times, i)
		}
		currPos = currPos - (-yVel + i)
	}
	return times
}

func apex(y int) int {
	return (y * (y + 1)) / 2
}
