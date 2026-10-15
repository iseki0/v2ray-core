package main

import (
	"reflect"
	"testing"
)

func TestLegacyArgs(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		want []string
	}{
		{"version", []string{"--version"}, []string{"version"}},
		{"short version", []string{"-version"}, []string{"version"}},
		{"run", []string{"--config", "config.json"}, []string{"run", "-config=config.json"}},
		{"test", []string{"-test", "-config", "config.json"}, []string{"test", "-config=config.json"}},
		{"test after config", []string{"-config=a.json", "--test"}, []string{"test", "-config=a.json"}},
		{"false test", []string{"-test=false", "-c", "a.json"}, []string{"run", "-c=a.json"}},
		{"flag-like filename", []string{"-config", "--test"}, []string{"run", "-config=--test"}},
		{"multiple configs", []string{"-config", "a b.json", "-c=b.json", "-format", "json", "-d", "conf", "-r"}, []string{"run", "-config=a b.json", "-c=b.json", "-format=json", "-d=conf", "-r"}},
		{"modern run", []string{"run", "-c", "a.json"}, []string{"run", "-c", "a.json"}},
		{"modern test", []string{"test", "-c", "a.json"}, []string{"test", "-c", "a.json"}},
		{"modern version", []string{"version"}, []string{"version"}},
		{"help", []string{"--help"}, []string{"--help"}},
		{"no args", nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := legacyArgs(tc.args)
			if err != nil || !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("legacyArgs(%q) = %q, %v; want %q", tc.args, got, err, tc.want)
			}
		})
	}
}

func TestLegacyArgsRejectInvalidInput(t *testing.T) {
	for _, args := range [][]string{{"-config"}, {"--unknown"}, {"-test=invalid"}, {"-config", "a.json", "extra"}} {
		if got, err := legacyArgs(args); err == nil {
			t.Errorf("legacyArgs(%q) = %q; want error", args, got)
		}
	}
}
