func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n <=1 {
		return n
	}
	type car struct{
		pos int
		speed int
	}
	cars := make([]car,n)
	for i := range n{
		cars[i] = car{pos : position[i],speed : speed[i]}
	}
	sort.Slice(cars, func(i,j int) bool{
		return cars[i].pos > cars[j].pos
	})
	stack := make([]float64,0,n)

	for _,c := range cars{
		time := float64(target - c.pos)/float64(c.speed)
		if len(stack) == 0 || time > stack[len(stack)-1]{
			stack = append(stack,time)
		}
	}
	return len(stack)
}
