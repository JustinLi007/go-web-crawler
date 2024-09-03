package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	testcases := map[string]struct {
		input    map[string]int
		expected []page
	}{
		"order count descending": {
			input: map[string]int{
				"url1": 5,
				"url2": 1,
				"url3": 3,
				"url4": 10,
				"url5": 7,
			},
			expected: []page{
				{key: "url4", val: 10},
				{key: "url5", val: 7},
				{key: "url1", val: 5},
				{key: "url3", val: 3},
				{key: "url2", val: 1},
			},
		},
		"alphabetize": {
			input: map[string]int{
				"d": 2,
				"a": 1,
				"e": 3,
				"b": 1,
				"c": 2,
			},
			expected: []page{
				{key: "e", val: 3},
				{key: "c", val: 2},
				{key: "d", val: 2},
				{key: "a", val: 1},
				{key: "b", val: 1},
			},
		},
		"empty map": {
			input:    map[string]int{},
			expected: []page{},
		},
		"nil map": {
			input:    nil,
			expected: []page{},
		},
		"one key": {
			input: map[string]int{
				"url1": 1,
			},
			expected: []page{
				{key: "url1", val: 1},
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			actual := sortPages(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestSortPages2(t *testing.T) {
	testcases := map[string]struct {
		input    map[string]int
		expected []page
	}{
		"order count descending": {
			input: map[string]int{
				"url1": 5,
				"url2": 1,
				"url3": 3,
				"url4": 10,
				"url5": 7,
			},
			expected: []page{
				{key: "url4", val: 10},
				{key: "url5", val: 7},
				{key: "url1", val: 5},
				{key: "url3", val: 3},
				{key: "url2", val: 1},
			},
		},
		"alphabetize": {
			input: map[string]int{
				"d": 2,
				"a": 1,
				"e": 3,
				"b": 1,
				"c": 2,
			},
			expected: []page{
				{key: "e", val: 3},
				{key: "c", val: 2},
				{key: "d", val: 2},
				{key: "a", val: 1},
				{key: "b", val: 1},
			},
		},
		"empty map": {
			input:    map[string]int{},
			expected: []page{},
		},
		"nil map": {
			input:    nil,
			expected: []page{},
		},
		"one key": {
			input: map[string]int{
				"url1": 1,
			},
			expected: []page{
				{key: "url1", val: 1},
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			actual := sortPages2(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
