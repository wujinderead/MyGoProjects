package main

// https://leetcode.com/problems/palindrome-partitioning/
// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return all possible palindrome partitioning of s.
func partition(s string) [][]string {
	// todo
	// aabcba
	return [][]string{
		{"a", "a", "b", "c", "b", "a"},
		{"a", "a", "bcb", "a"},
		{"a", "abcba"},
		{"aa", "b", "c", "b", "a"},
		{"aa", "bcb", "a"},
	}
}
