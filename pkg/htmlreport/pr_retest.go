package htmlreport

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/kubevirt/ci-health/pkg/constants"
	"github.com/kubevirt/ci-health/pkg/types"
)

//go:embed pr-retest-report.gohtml
var prRetestReportTemplate string

// PRRetestSection is one of the Open / Merged blocks in the HTML report.
type PRRetestSection struct {
	Source         string
	Title          string
	HeadlineLead   string
	HeadlineLabel  string
	Detail         string
	Total          int
	RetestedPRs    []types.PRRetestSummary
	UnexplainedPRs []types.PRRetestSummary
	CauseFlake     int
	CauseCI        int
	CauseExternal  int
	CauseOther     int
}

// PRRetestReportData is the template input for the PR retest HTML report.
type PRRetestReportData struct {
	GeneratedAt time.Time
	Source      string
	DataDays    int
	EndDate     string
	Open        PRRetestSection
	Merged      PRRetestSection
}

func prGitHubURL(source string, number int) string {
	if source == "" {
		source = constants.DefaultSource
	}
	return fmt.Sprintf("https://github.com/%s/pull/%d", source, number)
}

func causeDisplay(cause string) string {
	switch cause {
	case types.FailureCauseFlake, types.FailureCauseCI, types.FailureCauseExternal:
		return cause
	default:
		return "other"
	}
}

func partitionRetested(prs []types.PRRetestSummary) (explained, unexplained []types.PRRetestSummary, flake, ci, external, other int) {
	for _, pr := range prs {
		if pr.RetestCount <= 0 {
			continue
		}
		if len(pr.Failures) == 0 {
			unexplained = append(unexplained, pr)
			continue
		}
		explained = append(explained, pr)
		for _, failure := range pr.Failures {
			switch failure.Cause {
			case types.FailureCauseFlake:
				flake++
			case types.FailureCauseExternal:
				external++
			case types.FailureCauseCI:
				ci++
			default:
				other++
			}
		}
	}
	return explained, unexplained, flake, ci, external, other
}

func buildSection(source, title, headlineLabel, detail string, prs []types.PRRetestSummary, total int) PRRetestSection {
	explained, unexplained, flake, ci, external, other := partitionRetested(prs)
	retested := len(explained) + len(unexplained)
	if total == 0 {
		total = len(prs)
	}
	if total == 0 {
		total = retested
	}
	return PRRetestSection{
		Source:         source,
		Title:          title,
		HeadlineLead:   fmt.Sprintf("%d / %d", retested, total),
		HeadlineLabel:  headlineLabel,
		Detail:         detail,
		Total:          total,
		RetestedPRs:    explained,
		UnexplainedPRs: unexplained,
		CauseFlake:     flake,
		CauseCI:        ci,
		CauseExternal:  external,
		CauseOther:     other,
	}
}

func buildPRRetestReportData(results *types.Results, generatedAt time.Time) PRRetestReportData {
	return PRRetestReportData{
		GeneratedAt: generatedAt,
		Source:      results.Source,
		DataDays:    results.DataDays,
		EndDate:     results.EndDate,
		Merged: buildSection(
			results.Source,
			"Merged PRs",
			"merged PRs with a /retest in this window",
			"PRs merged in this window that received a /retest. Each lists the required e2e jobs that failed, labeled flake, CI (internal), external, or other.",
			results.PRRetestReport,
			results.MergedPRCount,
		),
		Open: buildSection(
			results.Source,
			"Open PRs",
			"recently updated open PRs with a /retest in this window",
			"Open PRs updated in this window that received a /retest. Each lists the required e2e jobs that failed, labeled flake, CI (internal), external, or other.",
			results.OpenPRRetestReport,
			results.OpenPRCount,
		),
	}
}

// WritePRRetestReport renders the PR-centric retest report next to results.json.
func WritePRRetestReport(results *types.Results, outputDir string) (err error) {
	if results == nil {
		return fmt.Errorf("results are required")
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("could not create report directory: %w", err)
	}

	funcMap := template.FuncMap{
		"prURL":        prGitHubURL,
		"causeDisplay": causeDisplay,
	}
	reportTemplate, err := template.New("prRetests").Funcs(funcMap).Parse(prRetestReportTemplate)
	if err != nil {
		return fmt.Errorf("could not parse pr retest template: %w", err)
	}

	reportPath := filepath.Join(outputDir, constants.PRRetestReportFileName)
	outputFile, err := os.Create(reportPath)
	if err != nil {
		return fmt.Errorf("could not create report file: %w", err)
	}
	defer func() {
		if closeErr := outputFile.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("could not close report file: %w", closeErr)
		}
	}()

	if execErr := reportTemplate.Execute(outputFile, buildPRRetestReportData(results, time.Now().UTC())); execErr != nil {
		return fmt.Errorf("could not execute pr retest template: %w", execErr)
	}
	return err
}

// GeneratePRRetest loads results.json and writes pr-retest-report.html.
func GeneratePRRetest(opt *types.Options) error {
	if opt.ResultsPath == "" {
		return fmt.Errorf("the path to results.json is required")
	}
	body, err := os.ReadFile(opt.ResultsPath)
	if err != nil {
		return fmt.Errorf("failed to read results.json file: %w", err)
	}
	var results types.Results
	if err := json.Unmarshal(body, &results); err != nil {
		return fmt.Errorf("failed to unmarshal results.json: %w", err)
	}
	return WritePRRetestReport(&results, opt.Path)
}
