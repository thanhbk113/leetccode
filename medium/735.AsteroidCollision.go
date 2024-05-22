package main

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for i := 0; i < len(asteroids); i++ {
		asteroid := asteroids[i]
		if len(stack) == 0 || stack[len(stack)-1] < 0 || asteroid > 0 {
			stack = append(stack, asteroid)
		} else {
			top := stack[len(stack)-1]

			if -asteroid == top {
				stack = stack[:len(stack)-1]
			} else if -asteroid > top {
				stack = stack[:len(stack)-1]
				i--
			}
		}
	}

	return stack
}
