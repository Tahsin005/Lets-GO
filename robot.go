type Robot struct {
    steps int;
    width, height int;
    moved bool;
}

func Constructor(width int, height int) Robot {
    return Robot{0, width, height, false}
}


func (this *Robot) Step(num int)  {
    this.moved = true;
    this.steps = (this.steps + num) % (2 * this.width + 2 * this.height - 4)
}


func (this *Robot) GetPos() []int {
    if this.steps == 0 {
        return []int{0, 0}
    }

    switch this.GetDir() {
    case "East":
        return []int{this.steps, 0}
    case "North":
        return []int{this.width - 1, this.steps - (this.width - 1)}
    case "West":
        return []int{this.width - 1 - (this.steps - (this.width + this.height - 2)), this.height - 1}
    case "South":
        return []int{0, this.height - 1 - (this.steps - (2 * this.width + this.height - 3))}
    }
    return []int{-1, -1}
}

func (this *Robot) GetDir() string {
    steps := this.steps
    if steps == 0 {
        if this.moved {
            return "South"
        } else {
            return "East"
        }
    }

    if steps < this.width {
        return "East"
    }
    steps -= this.width - 1
    if steps < this.height {
        return "North"
    }
    steps -= this.height - 1
    if steps < this.width {
        return "West"
    }
    return "South"
}


/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
