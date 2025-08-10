package main

import "fmt"

// https://leetcode.com/problems/reordered-power-of-2/description/?envType=daily-question&envId=2025-08-10

func reorderedPowerOf2(n int) bool {
    if (n & (n - 1)) == 0 {
        return true
    }
	all2power := make(map[int][]string)
	p := 1
	for i := 0; i < 31; i++ {
		p = (1 << i)
		str := strconv.Itoa(p)
		l := len(str)
		all2power[l] = append(all2power[l], str)
	}

	nstr := strconv.Itoa(n)
	nl := len(nstr)

	nchar := strings.Split(nstr, "")
	sort.Strings(nchar)
	sortedstr := strings.Join(nchar, "")

	for _, s := range all2power[nl] {
		ns := strings.Split(s, "")
		sort.Strings(ns)
		st := strings.Join(ns, "")
		if st == sortedstr {
			return true
		}
	}

	return false
}

func main () {

}