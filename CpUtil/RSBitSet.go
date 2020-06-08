package CpUtil

import (
	mapset "github.com/deckarep/golang-set"
	"math"
)

type RSBitSet struct {
	NumLevel, NumBit, lastLimits, top int
	Words                             [][]uint64
	Limit, Index, Levels              []int
	mask                              []uint64
}

type FDERSBitSet struct {
	RSBitSet
	CurrentWord []uint64
	in          map[int]interface{}
	inn         mapset.Set
}

func NewRSBitSet(numTuples, numVars int) *RSBitSet {
	s := new(RSBitSet)
	s.NumLevel = numVars + 1
	s.NumBit = int(math.Ceil(float64(numTuples / BITSIZE)))
	s.lastLimits = numTuples % BITSIZE
	// 初始化Words
	s.Words = make([][]uint64, s.NumLevel)
	for i := 0; i < s.NumBit; i++ {
		s.Words[0][i] = ALLONE64
	}
	if s.lastLimits != 0 {
		s.Words[0][s.NumBit-1] <<= BITSIZE - s.lastLimits
	}
	for i := 1; i < s.NumLevel; i++ {
		s.Words[i] = make([]uint64, s.NumBit)
	}

	//初始化limit, index, mask
	// rint
	s.Limit = make([]int, s.NumLevel)
	s.Limit[0] = s.NumBit - 1
	for i := 1; i < s.NumLevel; i++ {
		s.Limit[i] = -1
	}

	// array of int,  index.length = p
	s.Index = make([]int, s.NumBit)
	for i := 0; i < s.NumBit; i++ {
		s.Index[i] = i
	}
	//  val map = Array.range(0, numBit)
	// array of long, mask.length = p

	s.mask = make([]uint64, s.NumBit)
	s.top = 0

	return s
}

func (s *RSBitSet) NewLevel(level int) {
	// 记录的最顶层非level，说明需要新加一层
	if s.Levels[s.top] != level {
		// 记录新层和旧层
		preTop := s.top
		s.top++
		//oriLevel := s.Levels[preTop]
		newLevel := level
		s.Levels[s.top] = newLevel
		s.Limit[s.top] = s.Limit[preTop]

		for i := s.Limit[newLevel]; i >= 0; i-- {
			offset := s.Index[i]
			s.Words[s.top][offset] = s.Words[preTop][offset]
		}
	}
}

func (s *RSBitSet) BackLevel(level int) {
	for ; s.Levels[s.top] >= level; s.top-- {
		s.Limit[s.top] = INDEXOVERFLOW
	}
}

func (s *RSBitSet) IsEmpty() bool {
	return s.Limit[s.top] == INDEXOVERFLOW
}

func (s *RSBitSet) CurrentLevel() int {
	return s.Levels[s.top]
}

func (s *RSBitSet) ClearMask() {
	for i := s.Limit[s.top]; i >= 0; i-- {
		offset := s.Index[i]
		s.mask[offset] = 0
	}
}

func (s *RSBitSet) ReverseMask() {
	for i := s.Limit[s.top]; i >= 0; i-- {
		offset := s.Index[i]
		s.mask[offset] = ^s.mask[offset]
	}
}

func (s *RSBitSet) AddToMask(m []uint64) {
	for i := s.Limit[s.top]; i >= 0; i-- {
		offset := s.Index[i]
		s.mask[offset] |= m[offset]
	}
}

func (s *RSBitSet) IntersectWithMask() bool {
	changed := false
	var w, currentWords uint64

	for i := s.Limit[s.top]; i >= 0; i-- {
		offset := s.Index[i]
		currentWords = s.Words[s.top][offset]
		w = currentWords & s.mask[offset]
		if w != currentWords {
			s.Words[s.top][offset] = w
			changed = true
			if w == 0 {
				j := s.Limit[s.top]
				s.Index[i] = s.Index[j]
				s.Index[j] = offset
				s.Limit[s.top]--
			}
		}
	}

	return changed
}

func (s *RSBitSet) IntersectIndex(m []uint64) int {
	for i := s.Limit[s.top]; i >= 0; i-- {
		offset := s.Index[i]
		if s.Words[s.top][offset]&m[offset] != 0 {
			return offset
		}
	}
	return INDEXOVERFLOW
}

func (s *RSBitSet) TopLevel() int {
	return s.Levels[s.top]
}
