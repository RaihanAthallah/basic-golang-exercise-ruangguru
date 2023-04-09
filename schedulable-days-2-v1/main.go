package main

func SchedulableDays(villager [][]int) []int {
	result := []int{}
	if len(villager) == 1 {
		for i := 0; i < len(villager[0]); i++ {
			result = append(result, villager[0][i])
		}
	} else if len(villager) == 0 {
		return result
	} else {
		for _, v := range villager[0] {
			for j := 1; j < len(villager)-1; j++ {
				for k := 0; k < len(villager[j]); k++ {
					if v == villager[j][k] {
						result = append(result, villager[j][k])
					}
				}

			}
		}
	}

	return result // TODO: replace this
}
