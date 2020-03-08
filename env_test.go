package env

import (
	"os"
	"strconv"
	"testing"
)

func TestImportantVar(t *testing.T) {
	type args struct {
		envVarName   string
		defaultValue string
	}
	tests := []struct {
		name      string
		args      args
		shouldSet bool
		want      string
	}{
		{
			name: "should return default value",
			args: args{
				envVarName:   "IMPORTANT_VAR",
				defaultValue: "foobar",
			},
			want: "foobar",
		},
		{
			name: "should return os file",
			args: args{
				envVarName:   "IMPORTANT_VAR",
				defaultValue: "foobar",
			},
			want:      "candyland",
			shouldSet: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, tt.want)
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			if got := ImportantVar(tt.args.envVarName, tt.args.defaultValue); got != tt.want {
				t.Errorf("ImportantVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMandatoryVar(t *testing.T) {
	type args struct {
		envVarName string
	}
	tests := []struct {
		name        string
		args        args
		shouldSet   bool
		shouldPanic bool
		want        string
	}{
		{
			name: "should panic env is missing",
			args: args{
				envVarName: "MANDATORY_VAR",
			},
			shouldPanic: true,
			want:        "",
		},
		{
			name: "should retrieve value",
			args: args{
				envVarName: "MANDATORY_VAR",
			},
			shouldSet:   true,
			shouldPanic: false,
			want:        "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, tt.want)
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			defer func(t *testing.T, shouldPanic bool) {
				if r := recover(); r == nil && shouldPanic {
					t.Fatal("MandatoryVar should panic", r)
				}
			}(t, tt.shouldPanic)

			if got := MandatoryVar(tt.args.envVarName); got != tt.want {
				t.Errorf("MandatoryVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMandatoryVarAsBool(t *testing.T) {
	type args struct {
		envVarName string
	}
	tests := []struct {
		name        string
		args        args
		shouldSet   bool
		shouldPanic bool
		want        bool
	}{
		{
			name: "should panic env is missing",
			args: args{
				envVarName: "MANDATORY_BOOL",
			},
			shouldPanic: true,
		},
		{
			name: "should retrieve value",
			args: args{
				envVarName: "MANDATORY_BOOL",
			},
			shouldSet:   true,
			shouldPanic: false,
			want:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, strconv.FormatBool(tt.want))
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			defer func(t *testing.T, shouldPanic bool) {
				if r := recover(); r == nil && shouldPanic {
					t.Fatal("MandatoryVarAsBool should panic", r)
				}
			}(t, tt.shouldPanic)

			if got := MandatoryVarAsBool(tt.args.envVarName); got != tt.want {
				t.Errorf("MandatoryVarAsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMandatoryVarAsInt(t *testing.T) {
	type args struct {
		envVarName string
	}
	tests := []struct {
		name        string
		args        args
		shouldSet   bool
		shouldPanic bool
		want        int
	}{
		{
			name: "should panic env is missing",
			args: args{
				envVarName: "MANDATORY_INT",
			},
			shouldPanic: true,
		},
		{
			name: "should retrieve value",
			args: args{
				envVarName: "MANDATORY_INT",
			},
			shouldSet:   true,
			shouldPanic: false,
			want:        10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, strconv.Itoa(tt.want))
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			defer func(t *testing.T, shouldPanic bool) {
				if r := recover(); r == nil && shouldPanic {
					t.Fatal("MandatoryVarAsInt should panic", r)
				}
			}(t, tt.shouldPanic)

			if got := MandatoryVarAsInt(tt.args.envVarName); got != tt.want {
				t.Errorf("MandatoryVarAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMandatoryVarAsInt64(t *testing.T) {
	type args struct {
		envVarName string
	}
	tests := []struct {
		name        string
		args        args
		shouldSet   bool
		shouldPanic bool
		want        int64
	}{
		{
			name: "should panic env is missing",
			args: args{
				envVarName: "MANDATORY_INT64",
			},
			shouldPanic: true,
		},
		{
			name: "should retrieve value",
			args: args{
				envVarName: "MANDATORY_INT64",
			},
			shouldSet:   true,
			shouldPanic: false,
			want:        10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, strconv.Itoa(int(tt.want)))
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			defer func(t *testing.T, shouldPanic bool) {
				if r := recover(); r == nil && shouldPanic {
					t.Fatal("MandatoryVarAsInt64 should panic", r)
				}
			}(t, tt.shouldPanic)

			if got := MandatoryVarAsInt64(tt.args.envVarName); got != tt.want {
				t.Errorf("MandatoryVarAsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVar(t *testing.T) {
	type args struct {
		envVarName   string
		defaultValue string
	}
	tests := []struct {
		name      string
		args      args
		shouldSet bool
		want      string
	}{
		{
			name: "should return default value",
			args: args{
				envVarName:   "VAR",
				defaultValue: "foobar",
			},
			want: "foobar",
		},
		{
			name: "should return os file",
			args: args{
				envVarName:   "VAR",
				defaultValue: "foobar",
			},
			want:      "candyland",
			shouldSet: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, tt.want)
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			if got := Var(tt.args.envVarName, tt.args.defaultValue); got != tt.want {
				t.Errorf("Var() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVarAsBool(t *testing.T) {
	type args struct {
		envVarName   string
		defaultValue bool
	}
	tests := []struct {
		name      string
		args      args
		shouldSet bool
		want      bool
	}{
		{
			name: "should return default value",
			args: args{
				envVarName:   "BOOL",
				defaultValue: true,
			},
			want: true,
		},
		{
			name: "should return os file",
			args: args{
				envVarName:   "BOOL",
				defaultValue: false,
			},
			want:      true,
			shouldSet: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, strconv.FormatBool(tt.want))
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			if got := VarAsBool(tt.args.envVarName, tt.args.defaultValue); got != tt.want {
				t.Errorf("VarAsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVarAsInt(t *testing.T) {
	type args struct {
		envVarName   string
		defaultValue int
	}
	tests := []struct {
		name      string
		args      args
		shouldSet bool
		want      int
	}{
		{
			name: "should return default value",
			args: args{
				envVarName:   "INT",
				defaultValue: 10,
			},
			want: 10,
		},
		{
			name: "should return os file",
			args: args{
				envVarName:   "INT",
				defaultValue: 20,
			},
			want:      10,
			shouldSet: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, strconv.Itoa(tt.want))
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			if got := VarAsInt(tt.args.envVarName, tt.args.defaultValue); got != tt.want {
				t.Errorf("VarAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVarAsInt64(t *testing.T) {
	type args struct {
		envVarName   string
		defaultValue int64
	}
	tests := []struct {
		name      string
		args      args
		shouldSet bool
		want      int64
	}{
		{
			name: "should return default value",
			args: args{
				envVarName:   "INT64",
				defaultValue: 10,
			},
			want: 10,
		},
		{
			name: "should return os file",
			args: args{
				envVarName:   "INT64",
				defaultValue: 50,
			},
			want:      20,
			shouldSet: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, strconv.Itoa(int(tt.want)))
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			if got := VarAsInt64(tt.args.envVarName, tt.args.defaultValue); got != tt.want {
				t.Errorf("VarAsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVar(t *testing.T) {
	type args struct {
		envVarName   string
		defaultValue string
	}
	tests := []struct {
		name      string
		args      args
		shouldSet bool
		want      string
	}{
		{
			name: "should return default value",
			args: args{
				envVarName:   "IMPORTANT_VAR",
				defaultValue: "foobar",
			},
			want: "foobar",
		},
		{
			name: "should return os file",
			args: args{
				envVarName:   "IMPORTANT_VAR",
				defaultValue: "foobar",
			},
			want:      "candyland",
			shouldSet: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSet {
				os.Setenv(tt.args.envVarName, tt.want)
				t.Cleanup(func() {
					os.Unsetenv(tt.args.envVarName)
				})
			}

			if got := getVar(tt.args.envVarName, tt.args.defaultValue, false); got != tt.want {
				t.Errorf("getVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
