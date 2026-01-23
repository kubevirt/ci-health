package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	cifailures "github.com/kubevirt/ci-health/pkg/ci-failures"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	tmpOutputPath = "output/tmp"
	mdOutputPath  = "output/kubevirt/kubevirt/ci-failures"
)

var (
	//go:embed ci-failures.gomd
	ciFailuresTemplateContent string

	ciFailuresTemplate *template.Template

	rootCmd = &cobra.Command{
		Use:   "ci-failures",
		Short: "A CLI tool to process CI failures.",
	}

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate reports.",
	}

	yamlCmd = &cobra.Command{
		Use:   "yaml",
		Short: "Generate YAML failure reports.",
		RunE:  generateYAML,
	}

	mdCmd = &cobra.Command{
		Use:   "md",
		Short: "Generate Markdown summary report from YAML files.",
		RunE:  generateMarkdown,
	}

	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Generate YAML files and full Markdown summary report.",
		RunE:  generateReport,
	}
)

func generateReport(cmd *cobra.Command, args []string) error {
	err := generateYAML(cmd, args)
	if err != nil {
		return fmt.Errorf("failed generating ci-failure YAML files: %v", err)
	}
	err = generateMarkdown(cmd, args)
	if err != nil {
		return fmt.Errorf("failed generating ci failure markdown report: %v", err)
	}
	return nil
}

func generateYAML(_ *cobra.Command, _ []string) error {
	ciFailureJobURLs, err := cifailures.ShowCIFailureJobs()
	if err != nil {
		return fmt.Errorf("failed retrieving ci failure job urls: %v", err)
	}
	if err = os.MkdirAll(tmpOutputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}
	err = writeCIFailureJobsURLFile(ciFailureJobURLs)
	if err != nil {
		return fmt.Errorf("failed to write ci failure jobs file: %w", err)
	}
	_, err = cifailures.DownloadBuildLogs(ciFailureJobURLs)
	if err != nil {
		return fmt.Errorf("failed downloading build logs: %v", err)
	}
	_, err = cifailures.ExtractErrors(ciFailureJobURLs)
	if err != nil {
		return fmt.Errorf("failed extracting errors from build logs: %v", err)
	}
	return nil
}

func writeCIFailureJobsURLFile(ciFailureJobURLs []string) error {
	log.Printf("generating output in directory %s", tmpOutputPath)
	ciFailureJobsFile, err := os.Create(filepath.Join(tmpOutputPath, "ci-failure-jobs.txt"))
	if err != nil {
		return fmt.Errorf("failed creating ci failure jobs file: %v", err)
	}
	defer ciFailureJobsFile.Close()
	_, err = io.WriteString(ciFailureJobsFile, strings.Join(ciFailureJobURLs, "\n"))
	if err != nil {
		return fmt.Errorf("failed writing ci failure jobs file: %v", err)
	}
	return nil
}

func generateMarkdown(_ *cobra.Command, _ []string) error {
	if err := os.MkdirAll(mdOutputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	files, err := filepath.Glob(filepath.Join(tmpOutputPath, "/errors-*/*.yaml"))
	if err != nil {
		return fmt.Errorf("failed to find yaml files: %v", err)
	}

	templateData := TemplateData{
		Date:              time.Now().Format(time.DateTime),
		Org:               "kubevirt",
		Repo:              "kubevirt",
		FailuresPerBranch: Failures{CategoryName: "per branch"},
		FailuresPerDay:    Failures{CategoryName: "per day"},
		FailuresPerSIG:    Failures{CategoryName: "per SIG"},
	}

	jobBuildErrorsByJobName := map[string]*cifailures.JobBuildErrors{}

	// Deserialize JobBuildErrors files
	for _, fileName := range files {
		file, err := os.Open(fileName)
		var jobBuildErrors *cifailures.JobBuildErrors
		err = yaml.NewDecoder(file).Decode(&jobBuildErrors)
		if err != nil {
			return fmt.Errorf("failed to decode file %q: %v", fileName, err)
		}
		jobBuildErrorsByJobName[jobBuildErrors.JobName] = jobBuildErrors
	}

	for _, jobBuildErrorsForJob := range jobBuildErrorsByJobName {

		templateData.FailuresPerSIG.AddAll(jobBuildErrorsForJob.SIG(), jobBuildErrorsForJob)
		templateData.FailuresPerBranch.AddAll(jobBuildErrorsForJob.Branch(), jobBuildErrorsForJob)

		for _, jobBuildError := range jobBuildErrorsForJob.BuildErrors {
			day := jobBuildError.Started.Format(time.DateOnly)
			templateData.FailuresPerDay.AddError(day, jobBuildError)
		}
	}

	file, err := os.Create("output/kubevirt/kubevirt/ci-failures/summary.md")
	if err != nil {
		log.Fatalf("failed to create summary.md: %v", err)
	}
	err = ciFailuresTemplate.Execute(file, &templateData)
	if err != nil {
		log.Fatalf("failed to create summary.md when executing template: %v", err)
	}
	log.Printf("Generated Markdown summary: %s", file.Name())
	return nil
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(yamlCmd)
	generateCmd.AddCommand(mdCmd)
	generateCmd.AddCommand(reportCmd)

	var err error
	ciFailuresTemplate, err = template.New("summary.md").Parse(ciFailuresTemplateContent)
	if err != nil {
		log.Fatalf("failed to create summary.md when executing template: %v", err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
