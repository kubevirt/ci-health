package cost

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	//go:embed cost-report.gohtml
	costReportTemplate string
)

// ReportData wraps UsageReport with additional fields for the HTML template.
type ReportData struct {
	*UsageReport
	GeneratedAt time.Time
}

func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"gib": func(bytes float64) float64 {
			return bytes / (1024 * 1024 * 1024)
		},
		"pctClass": func(pct float64) template.HTML {
			class := "pct-low"
			if pct >= 5.0 {
				class = "pct-high"
			} else if pct >= 2.0 {
				class = "pct-medium"
			}
			return template.HTML(fmt.Sprintf(`<span class="%s">%.2f%%</span>`, class, pct))
		},
		"deref": func(p *float64) float64 {
			if p == nil {
				return 0
			}
			return *p
		},
	}
}

// GenerateHTMLReport renders the usage report as an HTML file.
func GenerateHTMLReport(report *UsageReport, outputDir string) error {
	tmpl, err := template.New("costReport").Funcs(templateFuncs()).Parse(costReportTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	reportPath := filepath.Join(outputDir, "cost-report.html")
	f, err := os.Create(reportPath)
	if err != nil {
		return fmt.Errorf("creating report file: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.WithError(err).Warn("failed closing report file")
		}
	}()

	data := &ReportData{
		UsageReport: report,
		GeneratedAt: time.Now().UTC(),
	}

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	log.Infof("wrote HTML report to %s", reportPath)
	return nil
}

// GenerateReport writes both JSON and HTML output.
func GenerateReport(report *UsageReport, basePath, source string) error {
	outDir := filepath.Join(basePath, strings.ReplaceAll(source, "/", string(filepath.Separator)))

	if err := writeJSON(report, outDir); err != nil {
		return err
	}
	if err := GenerateHTMLReport(report, outDir); err != nil {
		return err
	}

	return nil
}

func writeJSON(report *UsageReport, outDir string) error {
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	resultsPath := filepath.Join(outDir, "cost-results.json")
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling results: %w", err)
	}
	if err := os.WriteFile(resultsPath, data, 0644); err != nil {
		return fmt.Errorf("writing results: %w", err)
	}

	log.Infof("wrote JSON results to %s", resultsPath)
	return nil
}
