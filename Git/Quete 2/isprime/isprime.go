package main

func IsPrime(nb int) bool {
	if nb >= 2 {
		if nb%2 == 0 && nb != 2 {
			return false
		}
		for i := 3; i*i <= nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
