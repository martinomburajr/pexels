package config

import "testing"

func TestCanonicalBasePath(t *testing.T) {
	type args struct {
		homeDir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"home directory", args{"/home/me"}, "/home/me/.pexels"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanonicalBasePath(tt.args.homeDir); got != tt.want {
				t.Errorf("CanonicalBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanonicalPicturePath(t *testing.T) {
	type args struct {
		homeDir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"home directory", args{"/home/me"}, "/home/me/.pexels/pictures"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanonicalPicturePath(tt.args.homeDir); got != tt.want {
				t.Errorf("CanonicalPicturePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigPath(t *testing.T) {
	type args struct {
		homeDir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"home directory", args{"/home/me"}, "/home/me/.pexels/pexels.config.json"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfigPath(tt.args.homeDir); got != tt.want {
				t.Errorf("ConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestPexelsConfig_Load(t *testing.T) {
	type fields struct {
		APIKEY string
	}
	type args struct {
		configPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"load-api-key - error folder not file", fields{"some-value"}, args{"./testdata"}, true},
		{"load-api-key - readfile ok - jsonunmarshal error", fields{"some-value"}, args{"./testdata/bad-sample-pexels.config.json"}, true},
		{"load-api-key - readfile ok - jsonunmarshal error", fields{"some-value"}, args{"./testdata/bad-file"}, true},
		{"load-api-key - readfile ok - unmarshall ok - apikey in between other keys", fields{"some-value"}, args{"./testdata/fill-sample-pexels.config.json"}, false},
		{"load-api-key - readfile ok - unmarshall ok - duplicate keys - should select last key", fields{"some-value_2"}, args{"./testdata/multiple-keys-pexels.config.json"}, false},
		{"load-api-key - readfile ok - unmarshall ok - empty keys ", fields{""}, args{"./testdata/empty-key-pexels.config.json"}, true},
		{"load-api-key - readfile ok - unmarshall ok - key size is small", fields{"small"}, args{"./testdata/smallkey-pexels.config.json"}, true},
		{"load-api-key - readfile ok - unmarshall ok", fields{"some-value"}, args{"./testdata/sample-pexels.config.json"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PexelsConfig{
				APIKEY: tt.fields.APIKEY,
			}
			if err := p.Load(tt.args.configPath); (err != nil) != tt.wantErr {
				t.Errorf("PexelsConfig.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if p.APIKEY == "" && !tt.wantErr {
				t.Errorf("apikey: %s, wantErr %v", p.APIKEY, tt.wantErr)
			}
		})
	}
}
