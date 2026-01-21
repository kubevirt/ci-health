package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	cifailures "github.com/kubevirt/ci-health/pkg/ci-failures"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const outputDirectory = "output/tmp"

var rootCmd = &cobra.Command{
	Use:   "ci-failures",
	Short: "A CLI tool to process CI failures.",
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate reports.",
}

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Generate YAML failure reports.",
	RunE:  generateYAML,
}

var mdCmd = &cobra.Command{
	Use:   "md",
	Short: "Generate Markdown summary report from YAML files.",
	RunE:  generateMarkdown,
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate YAML files and full Markdown summary report.",
	RunE:  generateReport,
}

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
	if err = os.MkdirAll(outputDirectory, 0755); err != nil {
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
	log.Printf("generating output in directory %s", outputDirectory)
	ciFailureJobsFile, err := os.Create(filepath.Join(outputDirectory, "ci-failure-jobs.txt"))
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

type ErrorFrequency struct {
	Content string
	Hash    string
	Jobs    []string
	Count   int
}

func generateMarkdown(_ *cobra.Command, _ []string) error {
	files, err := filepath.Glob("output/tmp/errors-*/*.yaml")
	if err != nil {
		return fmt.Errorf("failed to find yaml files: %v", err)
	}

	errorOccurrences := make(map[string][]string)
	sigFailures := make(map[string]int)
	failedJobURLs := make(map[string]struct{})

	for _, file := range files {
		// Extract SIGs from directory path
		dir := filepath.Dir(file)
		dirName := filepath.Base(dir)
		if strings.HasPrefix(dirName, "errors-") {
			trimmed := strings.TrimPrefix(dirName, "errors-")
			lastDash := strings.LastIndex(trimmed, "-")
			if lastDash != -1 {
				sigsStr := trimmed[:lastDash]
				sigs := strings.Split(sigsStr, "-")
				for _, sig := range sigs {
					if sig != "" {
						sigFailures["sig-"+sig]++
					}
				}
			}
		}

		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %v", file, err)
		}

		var jobErrors cifailures.JobBuildErrors
		err = yaml.Unmarshal(data, &jobErrors)
		if err != nil {
			return fmt.Errorf("failed to unmarshal yaml from %s: %v", file, err)
		}

		for _, buildError := range jobErrors.BuildErrors {
			failedJobURLs[buildError.JobURL] = struct{}{}
			for _, snippet := range buildError.BuildLogErrorSnippets {
				trimmedContent := strings.TrimSpace(snippet.Context)
				errorOccurrences[trimmedContent] = append(errorOccurrences[trimmedContent], buildError.JobURL)
			}
		}
	}

	var sortedErrors []ErrorFrequency
	for content, jobs := range errorOccurrences {
		hash := sha256.Sum256([]byte(content))
		sortedErrors = append(sortedErrors, ErrorFrequency{
			Content: content,
			Hash:    hex.EncodeToString(hash[:])[:8],
			Jobs:    jobs,
			Count:   len(jobs),
		})
	}

	sort.Slice(sortedErrors, func(i, j int) bool {
		return sortedErrors[i].Count > sortedErrors[j].Count
	})

	var mdContent strings.Builder
	mdContent.WriteString("# CI Failure Summary\n\n")

	mdContent.WriteString("## Top 10 Failing Error Snippets\n\n")
	limit := 10
	if len(sortedErrors) < limit {
		limit = len(sortedErrors)
	}
	for i := 0; i < limit; i++ {
		err := sortedErrors[i]
		mdContent.WriteString(fmt.Sprintf("### Error: `%s`\n\n", err.Hash))
		mdContent.WriteString(fmt.Sprintf("**Occurrences:** %d\n\n", err.Count))
		mdContent.WriteString("```\n")
		mdContent.WriteString(err.Content)
		mdContent.WriteString("\n```\n\n")
		mdContent.WriteString("**Failed Jobs:**\n")
		// unique jobs
		jobMap := make(map[string]struct{})
		for _, job := range err.Jobs {
			jobMap[job] = struct{}{}
		}
		for job := range jobMap {
			mdContent.WriteString(fmt.Sprintf("- %s\n", job))
		}
		mdContent.WriteString("\n")
	}

	mdContent.WriteString("## Failures per SIG\n\n")
	for sig, count := range sigFailures {
		mdContent.WriteString(fmt.Sprintf("- **%s:** %d failures\n", sig, count))
	}
	mdContent.WriteString("\n")

	mdContent.WriteString("## All Failed Jobs\n\n")
	for url := range failedJobURLs {
		mdContent.WriteString(fmt.Sprintf("- %s\n", url))
	}

	err = os.WriteFile("output/tmp/summary.md", []byte(mdContent.String()), 0644)
	if err != nil {
		log.Fatalf("failed to write summary.md: %v", err)
	}
	fmt.Println("Generated Markdown summary: output/tmp/summary.md")
	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(yamlCmd)
	generateCmd.AddCommand(mdCmd)
	generateCmd.AddCommand(reportCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
