package utils

import "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"

func CalcReviewStat(reviews []*ent.Review) (count int, rating float32) {
	// calculate stats
	var accRating float32 = 0
	for _, r := range reviews {
		accRating += float32(*r.Score)
	}
	// check if a course have never been rate yed
	if len(reviews) <= 0 {
		count = 0
		rating = 0.0
		// prevent zero division
	} else if accRating == 0.0 {
		count = len(reviews)
		rating = 0.0
		// start counting
	} else {
		count = len(reviews)
		rating = accRating / float32(count)
	}
	return
}
