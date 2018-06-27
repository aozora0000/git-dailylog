package command

import (
	"strings"
	"github.com/uniplaces/carbon"
	"strconv"
		"errors"
)

type TimeDurationParser struct {
	t string
}

type BetweenTimestamps struct {
	From *carbon.Carbon
	To   *carbon.Carbon
}

func (s *TimeDurationParser) getDiff(sep string) int {
	i, err := strconv.Atoi(strings.Split(s.t, sep)[0])
	if err != nil {
		panic(err)
	}
	return i
}

func (s *TimeDurationParser) IsToday() bool {
	return strings.Contains(s.t, "today")
}

func (s *TimeDurationParser) GetToday() BetweenTimestamps {
	from, err := carbon.Create(carbon.Now().Year(), carbon.Now().Month(), carbon.Now().Day(), 0, 0, 0, 0, "Local")
	if err != nil {
		panic(err)
	}
	to := from.AddDays(1).SubSeconds(1)
	return BetweenTimestamps{
		From: from,
		To:   to,
	}
}

func (s *TimeDurationParser) IsDay() bool {
	return strings.Contains(s.t, "day")
}

func (s *TimeDurationParser) GetDay() BetweenTimestamps {
	today := s.GetToday()
	return BetweenTimestamps{
		From: today.From.SubDays(s.getDiff("day")),
		To: today.To,
	}
}

func (s *TimeDurationParser) IsWeek() bool {
	return strings.Contains(s.t, "week")
}

func (s *TimeDurationParser) GetWeek() BetweenTimestamps {
	today := s.GetToday()
	return BetweenTimestamps{
		From: today.From.SubWeeks(s.getDiff("week")),
		To: today.To,
	}
}

func (s *TimeDurationParser) IsYear() bool {
	return strings.Contains(s.t, "year")
}

func (s *TimeDurationParser) GetYear() BetweenTimestamps {
	today := s.GetToday()
	return BetweenTimestamps{
		From: today.From.SubYears(s.getDiff("year")),
		To: today.To,
	}
}

func (s *TimeDurationParser) Parse() BetweenTimestamps {
	if s.IsToday() {
		return s.GetToday()
	}
	if s.IsDay() {
		return s.GetDay()
	}
	if s.IsWeek() {
		return s.GetWeek()
	}
	if s.IsYear() {
		return s.GetYear()
	}
	panic(errors.New("DateParseError"))
}