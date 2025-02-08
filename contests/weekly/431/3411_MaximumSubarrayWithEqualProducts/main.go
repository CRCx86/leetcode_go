package main

func main() {

}

func maxLength(nums []int) int {
	var ans int
	n := len(nums)

	for left := 0; left < n; left++ {
		p := nums[left]
		g := nums[left]
		l := nums[left]

		for right := left + 1; right < n; right++ {
			p = prod(p, nums[right])
			g = gcd(g, nums[right])
			l = lcm(l, nums[right])

			if p == g*l {
				ans = _max(ans, right-left+1)
			}
		}
	}

	return ans
}

func prod(i, j int) int {
	return i * j
}

func gcd(i, j int) int {
	if j < i {
		i, j = j, i
	}

	for j != 0 {
		i, j = j, i%j
	}

	return i
}

func lcm(i, j int) int {
	return (i * j) / gcd(i, j)
}

func _max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
