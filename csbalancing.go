package csbalancing

import (
	"sort"
)

// Entity ...
type Entity struct {
	ID    int
	Score int
}

// CustomerSuccessBalancing ...
func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {
	pair := make(PairList, len(customerSuccess))
	csController := make(map[int]int)
	scoreDistance := 0
	scoreDistanceControl := 100
	picked := map[int]int{}

	for _, ct := range customers {
		for _, cs := range customerSuccess {
			if cs.Score >= ct.Score && !contains(customerSuccessAway, cs.ID) {
				scoreDistance = cs.Score - ct.Score
				if scoreDistance < scoreDistanceControl {

					//Seta a nova distância entre scores
					scoreDistanceControl = scoreDistance

					//Se existe no arr exclui o CS anterior
					if val, ok := picked[ct.ID]; ok {
						csController[val]--
						pair[val] = Pair{val, csController[val]}
					}

					//Adiciona a relação entre o CS e o CT
					csController[cs.ID]++
					pair[cs.ID] = Pair{cs.ID, csController[cs.ID]}
					picked[ct.ID] = cs.ID
				}
			}
		}
		//Reset da distancia
		scoreDistanceControl = 100
	}

	//Sorting
	sort.Sort(pair)

	//Empate
	if pair[len(pair)-1].Value == pair[len(pair)-2].Value {
		return 0
	}

	//Returning o maior pair com o sort já feito
	return pair[len(pair)-1].Key
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//Pair to Sort
type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
