package strings

import "strings"

// Delete the content after the last sub string "flag" in "s"
// If there is no "flag" in "s", returns ""
func RemoveTailAfterLastFlag(s, flag string) string {
	lastFlagIndex := strings.LastIndex(s, flag)
	if lastFlagIndex == -1 {
		return ""
	}
	result := s[:lastFlagIndex+1]
	return result
}

// Return a sub string between str1 and str2 from s
func SubStringBetween(s, str1, str2 string) string {
	index1 := strings.Index(s, str1)
	index2 := strings.Index(s, str2)
	if index1 == -1 || index2 == -1 {
		return ""
	}
	if index1+len(str1) < index2 {
		return s[index1+len(str1): index2]
	}
	if index2+len(str2) < index1 {
		return s[index2+len(str2): index1]
	}
	return ""
}
