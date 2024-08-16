package webhook

import (
	"reflect"
	"testing"
)

func TestGetConf(t *testing.T) {
	testCases := []struct {
		name          string
		inputPath     string
		expectedConf  SparkAppConfig
		expectedError bool
	}{
		{
			name:          "valid_config_file",
			inputPath:     "testdata/valid_config.yaml",
			expectedConf:  SparkAppConfig{FeatureList: FeatureList{Toleration: &FeatureMode{Enabled: true}}, SparkPatchValue: SparkPatchValue{}},
			expectedError: false,
		},
		{
			name:          "invalid_config_file",
			inputPath:     "testdata/invalid_config.yaml",
			expectedConf:  SparkAppConfig{},
			expectedError: true,
		},
		{
			name:          "empty_config_file",
			inputPath:     "testdata/empty_config.yaml",
			expectedConf:  SparkAppConfig{},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			conf, err := GetConf(tc.inputPath)
			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, Got error: %v", tc.expectedError, err)
			}
			if tc.expectedError {
				return
			}
			if !reflect.DeepEqual(conf, tc.expectedConf) {
				t.Errorf("Expected config: %v, Got config: %v", tc.expectedConf, conf)
			}
		})
	}
}

//
//func TestGetConf_invalid_file_path(t *testing.T) {
//	_, err := os.Stat("testdata/invalid_file_path.yaml")
//	if !os.IsNotExist(err) {
//		t.Errorf("Expected file not to exist, Got error: %v", err)
//	}
//	_, err = GetConf("testdata/invalid_file_path.yaml")
//	if err == nil {
//		t.Errorf("Expected error when file does not exist, Got no error")
//	}
//}
//
//func TestGetConf_empty_file_content(t *testing.T) {
//	tmpFile, err := ioutil.TempFile("", "testdata/empty_config.yaml")
//	if err != nil {
//		t.Fatalf("Error creating temporary file: %v", err)
//	}
//	defer os.Remove(tmpFile.Name())
//
//	err = tmpFile.Close()
//	if err != nil {
//		t.Fatalf("Error closing temporary file: %v", err)
//	}
//
//	conf, err := GetConf(tmpFile.Name())
//	if err == nil {
//		t.Errorf("Expected error when file is empty, Got no error")
//	}
//}
//
//func TestGetConf_invalid_yaml_content(t *testing.T) {
//	tmpFile, err := ioutil.TempFile("", "testdata/invalid_config.yaml")
//	if err != nil {
//		t.Fatalf("Error creating temporary file: %v", err)
//	}
//	defer os.Remove(tmpFile.Name())
//
//	_, err = tmpFile.Write([]byte("invalid_yaml_content"))
//	if err != nil {
//		t.Fatalf("Error writing to temporary file: %v", err)
//	}
//
//	err = tmpFile.Close()
//	if err != nil {
//		t.Fatalf("Error closing temporary file: %v", err)
//	}
//
//	conf, err := GetConf(tmpFile.Name())
//	if err == nil {
//		t.Errorf("Expected error when file contains invalid YAML, Got no error")
//	}
//}
