package util

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/kube2pulumi/pkg/pcl2pulumi"
	"github.com/pulumi/kube2pulumi/pkg/yaml2pcl"
)

func RunConversion(dirPath string, filePath string, language string) (string, error) {
	if filePath == "" && dirPath == "" {
		return "", fmt.Errorf("must specify a path for a file or directory\n")
	}
	if filePath != "" && dirPath != "" {
		return "", fmt.Errorf("must specify EITHER a path for a file or directory, not both\n")
	}
	var result string
	var outPath string
	var err error
	// filepath only
	if filePath != "" {
		result, err = yaml2pcl.ConvertFile(filePath)
		outPath, err = pcl2pulumi.Pcl2Pulumi(result, filePath, language)
		if err != nil {
			return "", err
		}
	} else { // dir only
		result, err = yaml2pcl.ConvertDirectory(dirPath)
		outPath, err = pcl2pulumi.Pcl2Pulumi(result, filepath.Join(dirPath, "main"), language)
		if err != nil {
			return "", err
		}
	}
	return outPath, nil
}