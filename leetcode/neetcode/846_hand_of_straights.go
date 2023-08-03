package neetcode

import "sort"

// TODO: research better way
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Slice(hand, func(i, j int) bool {
		return hand[i] < hand[j]
	})

	return isConsecutive(hand, groupSize)
}

func isConsecutive(hand []int, groupSize int) bool {
	var (
		i        int
		groupMap = make(map[int]struct{}, 0)
	)
	for len(hand) > 0 {
		if i == len(hand) {
			return false
		}
		if len(groupMap) == groupSize {
			i = 0
			groupMap = make(map[int]struct{}, 0)
		}
		if _, found := groupMap[hand[i]]; found {
			i++
			continue
		}

		groupMap[hand[i]] = struct{}{}
		_, ok := groupMap[hand[i]-1]
		if !ok && len(groupMap) >= 2 {
			return false
		}
		hand = append(hand[:i], hand[i+1:]...)
	}

	return len(groupMap) == groupSize
}
