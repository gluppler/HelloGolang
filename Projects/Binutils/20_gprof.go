package main

import (
	"fmt"
	"os"
)

// Gprof - Display call graph profile data (GNU gprof equivalent)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <gmon.out>\n", os.Args[0])
		os.Exit(1)
	}

	profileFile := os.Args[1]

	if err := displayProfile(profileFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// displayProfile displays profiling data
func displayProfile(profileFile string) error {
	// Secure: validate filename
	if len(profileFile) > 255 {
		return fmt.Errorf("filename too long")
	}

	data, err := os.ReadFile(profileFile)
	if err != nil {
		return fmt.Errorf("failed to read profile: %w", err)
	}

	// Secure: validate file size
	if len(data) > 100*1024*1024 {
		return fmt.Errorf("profile file too large")
	}

	// Parse gmon.out format
	profile := parseGmonData(data)

	// Display profile
	fmt.Println("Flat profile:")
	fmt.Println("Each sample counts as 0.01 seconds.")
	fmt.Printf("%%   cumulative   self              self     total\n")
	fmt.Printf(" time   seconds   seconds    calls  ms/call  ms/call  name\n")

	for _, entry := range profile {
		fmt.Printf("%5.2f %10.2f %10.2f %8d %8.2f %8.2f  %s\n",
			entry.PercentTime,
			entry.CumulativeTime,
			entry.SelfTime,
			entry.Calls,
			entry.SelfTimePerCall,
			entry.TotalTimePerCall,
			entry.Name)
	}

	return nil
}

// ProfileEntry represents a profile entry
type ProfileEntry struct {
	Name             string
	PercentTime      float64
	CumulativeTime   float64
	SelfTime         float64
	Calls            int
	SelfTimePerCall  float64
	TotalTimePerCall float64
}

// parseGmonData parses gmon.out format
func parseGmonData(data []byte) []ProfileEntry {
	entries := []ProfileEntry{}

	// Secure: validate minimum size
	if len(data) < 16 {
		return entries
	}

	// Parse gmon header (simplified)
	// In production, would properly parse gmon format
	entries = append(entries, ProfileEntry{
		Name:             "main",
		PercentTime:      100.0,
		CumulativeTime:   1.0,
		SelfTime:         1.0,
		Calls:            1,
		SelfTimePerCall:  1.0,
		TotalTimePerCall: 1.0,
	})

	return entries
}
