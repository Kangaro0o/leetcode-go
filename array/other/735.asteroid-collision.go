package other

// asteroidCollision 735. 行星碰撞
func asteroidCollision(asteroids []int) []int {
	if len(asteroids) == 0 {
		return nil
	}
	var stack []int
	stack = append(stack, asteroids[0])
	var top int // 保存栈顶元素
	for i := 1; i < len(asteroids); i++ {
		if len(stack) > 0 {
			top = stack[len(stack)-1]
		} else {
			top = -1
		}
		if top < 0 || asteroids[i] > 0 { // 此时不会发生碰撞，添加到栈顶
			stack = append(stack, asteroids[i])
			continue
		}
		// 拿到栈顶元素
		for top > 0 && asteroids[i] < 0 && len(stack) > 0 { // 此时发生碰撞
			// 1. 质量相同
			if top+asteroids[i] == 0 {
				stack = stack[:len(stack)-1]
				break
			}
			// 2. 栈顶 > 新元素
			if top+asteroids[i] > 0 {
				break
			}
			// 3. 栈顶 < 新元素
			if top+asteroids[i] < 0 {
				stack = stack[:len(stack)-1]
				top = -1
				if len(stack) > 0 {
					top = stack[len(stack)-1]
				}
				if top < 0 {
					stack = append(stack, asteroids[i])
				}
			}
		}
	}
	return stack
}
