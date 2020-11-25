package git_dailylog

import (
	"errors"
	"github.com/k0kubun/pp"
	"github.com/uniplaces/carbon"
	"os"
	"strconv"
	"strings"
	"time"
)

type BetweenTimestamps struct {
	From *carbon.Carbon
	To   *carbon.Carbon
}

func (s *BetweenTimestamps) IsBetween(t time.Time) bool {
	return s.From.Timestamp() <= t.Unix() && t.Unix() <= s.To.Timestamp()
}

type TimeDurationParser struct {
	t string
}

func (s *TimeDurationParser) getDiff(sep string) int {
	num := "1"
	if strings.Split(s.t, sep)[0] != "" {
		num = strings.Split(s.t, sep)[0]
	}
	i, err := strconv.Atoi(num)
	if err != nil {
		pp.Println(err.Error())
		os.Exit(1)
	}
	return i
}

func (s *TimeDurationParser) IsToday() bool {
	return strings.Contains(s.t, "today")
}

func (s *TimeDurationParser) GetToday() BetweenTimestamps {
	from, err := carbon.Create(carbon.Now().Year(), carbon.Now().Month(), carbon.Now().Day(), 0, 0, 0, 0, "Local")
	if err != nil {
		pp.Println(err.Error())
		os.Exit(1)
	}
	to := from.AddDays(1).SubSeconds(1)
	return BetweenTimestamps{
		From: from,
		To:   to,
	}
}

func (s *TimeDurationParser) IsYesterday() bool {
	return strings.Contains(s.t, "yesterday")
}

func (s *TimeDurationParser) GetYesterday() BetweenTimestamps {
	today := s.GetToday()
	return BetweenTimestamps{
		From: today.From.SubDays(1),
		To:   today.To,
	}
}

func (s *TimeDurationParser) IsDay() bool {
	return strings.Contains(s.t, "day")
}

func (s *TimeDurationParser) GetDay() BetweenTimestamps {
	today := s.GetToday()
	return BetweenTimestamps{
		From: today.From.SubDays(s.getDiff("day")),
		To:   today.To,
	}
}

func (s *TimeDurationParser) IsWeek() bool {
	return strings.Contains(s.t, "week")
}

func (s *TimeDurationParser) GetWeek() BetweenTimestamps {
	today := s.GetToday()
	return BetweenTimestamps{
		From: today.From.SubWeeks(s.getDiff("week")),
		To:   today.To,
	}
}

func (s *TimeDurationParser) IsYear() bool {
	return strings.Contains(s.t, "year")
}

func (s *TimeDurationParser) GetYear() BetweenTimestamps {
	today := s.GetToday()
	return BetweenTimestamps{
		From: today.From.SubYears(s.getDiff("year")),
		To:   today.To,
	}
}

func (s *TimeDurationParser) Parse() BetweenTimestamps {
	if s.IsToday() {
		return s.GetToday()
	}
	if s.IsYesterday() {
		return s.GetYesterday()
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
	pp.Println(errors.New("DateParseError").Error())
	os.Exit(1)
	return BetweenTimestamps{}
}
