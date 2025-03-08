// Package srp shows how Single Responsibility Principle can be applied in Go.
//
// https://en.wikipedia.org/wiki/Single-responsibility_principle
package srp

// CountLogLevelsFromFile reads a log file, parses logs, and counts log levels.
// All these responsibilities are mixed in one function.
func CountLogLevelsFromFile(f string) (map[string]int, error) {
	// ... read file
	// ... parse logs
	// ... count errors
	return map[string]int{"INFO": 10, "ERROR": 1}, nil
}

// ----------------------------------------------------------------

// LogConsumer is a struct that holds functions for reading, parsing, and counting logs.
// Responsibility is separated into different functions.
type LogConsumer struct {
	FileReader   func(f string) ([]byte, error)
	LogParser    func([]byte) ([]string, error)
	ErrorCounter func([]string) (map[string]int, error)
}

// CountLogLevelsFromFile reads a log file, parses logs, and counts log levels.
// Logic is modular and can be extended easily.
func (lc *LogConsumer) CountLogLevelsFromFile(f string) (map[string]int, error) {
	b, err := lc.FileReader(f)
	if err != nil {
		return nil, err
	}

	logs, err := lc.LogParser(b)
	if err != nil {
		return nil, err
	}

	return lc.ErrorCounter(logs)
}

// ----------------------------------------------------------------

// LogConsumer2 is example of even more modular approach with interfaces.

type Reader interface {
	ReadFile(f string) ([]byte, error)
}

type Parser interface {
	Parse(b []byte) ([]string, error)
}

type Counter interface {
	Count(logs []string) (map[string]int, error)
}

type LogConsumer2 struct {
	FileReader   Reader
	LogParser    Parser
	ErrorCounter Counter
}

func (lc *LogConsumer2) CountLogLevelsFromFile(f string) (map[string]int, error) {
	b, err := lc.FileReader.ReadFile(f)
	if err != nil {
		return nil, err
	}

	logs, err := lc.LogParser.Parse(b)
	if err != nil {
		return nil, err
	}

	return lc.ErrorCounter.Count(logs)
}
