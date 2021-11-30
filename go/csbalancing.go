package csbalancing

import (
	"sort"
)

// Entity struct
type Entity struct {
	ID    int
	Score int
}

// EntityCountable struct
type EntityCountable struct {
	Entity
	Count int
}

// CustomerSuccessBalancing function
func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {
	csMap := make(map[int]int, len(customerSuccess)) // map[ID]Score

	// Fill a CustomerSucess map - O(n)
	// An attempt to reduce time complexity increasing the space one
	for _, cs := range customerSuccess {
		csMap[cs.ID] = cs.Score
	}

	// Delete Away Customer Sucess from the map - O(t)
	for _, csa := range customerSuccessAway {
		delete(csMap, csa)
	}

	if len(csMap) == 0 || len(customers) == 0 {
		return 0
	}

	// Build a sorted slice from the CS map - O(n log n)
	var sortedCS []EntityCountable
	for k, v := range csMap {
		sortedCS = append(sortedCS, EntityCountable{
			Entity: Entity{k, v},
			Count:  0,
		})
	}
	sort.Slice(sortedCS, func(i, j int) bool {
		return sortedCS[i].Score < sortedCS[j].Score
	})

	// Sort customers slice in place - O(m log m)
	sort.Slice(customers, func(i, j int) bool {
		return customers[i].Score < customers[j].Score
	})

	cIdx := 0
	csIdx := 0
	attendMostCostumers := 0

	// Traverse customers and CS doing the count
	// This shows the use of indexes instead of range
	// O(n m)
	for cIdx < len(customers) && csIdx < len(sortedCS) {
		if customers[cIdx].Score <= sortedCS[csIdx].Score {
			sortedCS[csIdx].Count++
			cIdx++
			continue
		}
		csIdx++
	}

	maxCostumers := 0

	// Traverse CS discovering the right ID (most customers ID or 0 if there's more than one)
	// O(n)
	for _, cs := range sortedCS {
		if cs.Count == 0 {
			continue
		}

		if cs.Count == maxCostumers {
			return 0
		}

		if cs.Count > maxCostumers {
			maxCostumers = cs.Count
			attendMostCostumers = cs.ID
		}
	}

	return attendMostCostumers
}
