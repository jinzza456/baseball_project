package BaseballLib

import "github.com/jinzza456/baseball_project/BaseballLib"

type ProgressOfGameCount struct {
	baseCount   int
	foulCount   int
	strikeCount int
	outCount    int
	score       int
	hitCount    int
}

func (c *ProgressOfGameCount) AddBaseCount() {
	c.baseCount++
}

func (c *ProgressOfGameCount) AddFoulCount() {
	c.foulCount++
}

func (c *ProgressOfGameCount) AddStrikeCount() {
	c.strikeCount++
}

func (c *ProgressOfGameCount) AddOutCount() {
	c.outCount++
}

func (c *ProgressOfGameCount) AddScore() {
	c.score++
}

func (c *ProgressOfGameCount) AddHitCount() {
	c.hitCount++
}

func (c *ProgressOfGameCount) ResetBaseCount() {
	c.baseCount = 0
}

func (c *ProgressOfGameCount) ResetFoulCount() {
	c.foulCount = 0
}

func (c *ProgressOfGameCount) ResetStrikeCount() {
	c.strikeCount = 0
}

func (c *ProgressOfGameCount) ResetHitCount() {
	c.hitCount = 0
}

func IfBaseCountIsFour() {
	if BaseballLib.baseCount == 4 {
		BaseballLib.ScoredCall()
		BaseballLib.AddScore()
		BaseballLib.ResetBaseCount()
	}
}

func IfFoulCountIsTwo() {
	if BaseballLib.foulCount == 2 {
		BaseballLib.AddStrikeCount()
		BaseballLib.ResetFoulCount()
	}
}

func IfStrikeCountIsThree() {
	if BaseballLib.strikeCount == 3 {
		BaseballLib.OutCall()
		BaseballLib.AddOutCount()
	}
}

func IfHitCountIsOne() {
	BaseballLib.ResetHitCount()
}
