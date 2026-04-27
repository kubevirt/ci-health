package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	cifailures "github.com/kubevirt/ci-health/pkg/ci-failures"
	"github.com/kubevirt/ci-health/pkg/sigretests"
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

	analyzeBuildCmd = &cobra.Command{
		Use:   "analyze-build [prow-job-url]",
		Short: "Analyze a single build failure given a Prow job URL and output YAML.",
		Args:  cobra.ExactArgs(1),
		RunE:  analyzeBuild,
	}

	analyzePRCmd = &cobra.Command{
		Use:   "analyze-pr [github-pr-url]",
		Short: "Analyze all failed builds for a GitHub PR.",
		Long: `Analyze all failed builds for a GitHub PR.

Accepts a GitHub pull request URL, e.g.:
  https://github.com/kubevirt/kubevirt/pull/17287

Only repos served by prow.ci.kubevirt.io are supported.`,
		Args: cobra.ExactArgs(1),
		RunE: analyzePR,
	}

	analyzeK8sCmd = &cobra.Command{
		Use:   "analyze-k8s [prow-job-url]",
		Short: "Analyze Kubernetes cluster state from k8s-reporter artifacts.",
		Long: `Analyze Kubernetes cluster state from k8s-reporter artifacts for a Prow job.

Downloads pods, nodes, events, and KubeVirt object dumps from the k8s-reporter
artifacts directory in GCS, then runs failure detectors to identify issues like
CrashLoopBackOff, OOMKilled, NotReady nodes, warning events, and failed VMI
migrations.

Accepts a Prow job URL, e.g.:
  https://prow.ci.kubevirt.io/view/gs/kubevirt-prow/pr-logs/pull/kubevirt_kubevirt/17287/pull-kubevirt-e2e-k8s-1.31-sig-compute/2045831803410845696`,
		Args: cobra.ExactArgs(1),
		RunE: analyzeK8s,
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
		Date:                time.Now().Format(time.DateTime),
		Org:                 "kubevirt",
		Repo:                "kubevirt",
		FailuresPerBranch:   Failures{CategoryName: "per branch", Anchor: "per-branch"},
		FailuresPerDay:      Failures{CategoryName: "per day", Anchor: "per-day"},
		FailuresPerSIG:      Failures{CategoryName: "per SIG", Anchor: "per-sig"},
		FailuresPerCategory: Failures{CategoryName: "per error category", Anchor: "per-error-category"},
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

			templateData.FailuresPerCategory.AddError(jobBuildError.Category, jobBuildError)
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

func analyzeBuild(_ *cobra.Command, args []string) error {
	prowJobURL := args[0]

	jobBuildErrors, err := cifailures.AnalyzeBuild(prowJobURL)
	if err != nil {
		return fmt.Errorf("failed to analyze build: %v", err)
	}

	if err = os.MkdirAll(tmpOutputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	outputPath := filepath.Join(tmpOutputPath, fmt.Sprintf("%s.yaml", jobBuildErrors.JobName))
	if err = cifailures.WriteJobBuildErrorsYAML(outputPath, jobBuildErrors); err != nil {
		return fmt.Errorf("failed to write YAML output: %v", err)
	}

	log.Printf("wrote analysis to %s", outputPath)

	k8sResult, k8sErr := cifailures.AnalyzeK8s(prowJobURL)
	if k8sErr != nil {
		log.WithError(k8sErr).Warn("k8s artifact analysis failed, skipping")
	} else {
		k8sOutputPath := filepath.Join(tmpOutputPath, fmt.Sprintf("k8s-analysis-%d.yaml", k8sResult.BuildID))
		if writeErr := cifailures.WriteK8sAnalysisResultYAML(k8sOutputPath, k8sResult); writeErr != nil {
			log.WithError(writeErr).Warn("failed to write k8s analysis YAML")
		} else {
			log.Infof("wrote k8s analysis to %s (%d findings)", k8sOutputPath, k8sResult.Summary.TotalFindings)
		}
	}

	return nil
}

var allowedOrgs = map[string]bool{
	"kubevirt": true,
}

func parsePRURL(prURL string) (org, repo, prNumber string, err error) {
	const prefix = "https://github.com/"
	if !strings.HasPrefix(prURL, prefix) {
		return "", "", "", fmt.Errorf("expected a GitHub PR URL (https://github.com/org/repo/pull/N), got: %s", prURL)
	}
	parts := strings.Split(strings.TrimPrefix(prURL, prefix), "/")
	if len(parts) < 4 || parts[2] != "pull" {
		return "", "", "", fmt.Errorf("invalid GitHub PR URL format: %s", prURL)
	}
	org = parts[0]
	repo = parts[1]
	prNumber = parts[3]
	if _, parseErr := strconv.Atoi(prNumber); parseErr != nil {
		return "", "", "", fmt.Errorf("invalid PR number in URL: %s", prNumber)
	}
	if !allowedOrgs[org] {
		return "", "", "", fmt.Errorf("org %q is not served by prow.ci.kubevirt.io (allowed: %v)", org, allowedOrgs)
	}
	return org, repo, prNumber, nil
}

func analyzePR(_ *cobra.Command, args []string) error {
	prOrg, prRepo, prNumber, err := parsePRURL(args[0])
	if err != nil {
		return err
	}

	failedURLs, err := sigretests.ListPRFailures(prNumber, prOrg, prRepo)
	if err != nil {
		return fmt.Errorf("failed to list PR failures: %v", err)
	}

	if len(failedURLs) == 0 {
		log.Info("no failed builds found for this PR")
		return nil
	}

	log.Infof("found %d failed build(s)", len(failedURLs))

	if err = os.MkdirAll(tmpOutputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	for _, url := range failedURLs {
		jobBuildErrors, err := cifailures.AnalyzeBuild(url)
		if err != nil {
			log.WithError(err).Warnf("failed to analyze build %s, skipping", url)
			continue
		}

		outputPath := filepath.Join(tmpOutputPath, fmt.Sprintf("%s.yaml", jobBuildErrors.JobName))
		if err = cifailures.WriteJobBuildErrorsYAML(outputPath, jobBuildErrors); err != nil {
			log.WithError(err).Warnf("failed to write YAML for %s", jobBuildErrors.JobName)
			continue
		}

		log.Infof("wrote analysis to %s", outputPath)

		k8sResult, k8sErr := cifailures.AnalyzeK8s(url)
		if k8sErr != nil {
			log.WithError(k8sErr).Warnf("k8s artifact analysis failed for %s, skipping", url)
		} else {
			k8sOutputPath := filepath.Join(tmpOutputPath, fmt.Sprintf("k8s-analysis-%d.yaml", k8sResult.BuildID))
			if writeErr := cifailures.WriteK8sAnalysisResultYAML(k8sOutputPath, k8sResult); writeErr != nil {
				log.WithError(writeErr).Warnf("failed to write k8s analysis YAML for %s", url)
			} else {
				log.Infof("wrote k8s analysis to %s (%d findings)", k8sOutputPath, k8sResult.Summary.TotalFindings)
			}
		}
	}

	return nil
}

func analyzeK8s(_ *cobra.Command, args []string) error {
	prowJobURL := args[0]

	result, err := cifailures.AnalyzeK8s(prowJobURL)
	if err != nil {
		return fmt.Errorf("failed to analyze k8s artifacts: %v", err)
	}

	if err = os.MkdirAll(tmpOutputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	outputPath := filepath.Join(tmpOutputPath, fmt.Sprintf("k8s-analysis-%d.yaml", result.BuildID))
	if err = cifailures.WriteK8sAnalysisResultYAML(outputPath, result); err != nil {
		return fmt.Errorf("failed to write YAML output: %v", err)
	}

	log.Infof("wrote k8s analysis to %s (%d findings)", outputPath, result.Summary.TotalFindings)
	return nil
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(analyzeBuildCmd)
	rootCmd.AddCommand(analyzePRCmd)
	rootCmd.AddCommand(analyzeK8sCmd)
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
