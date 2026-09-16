package scw

import (
	"context"
	"io"
	"testing"
)

// TestFacadeSignatures pins the exported surface consumers compile against.
// Each assignment fails to compile when upstream renames a symbol or changes a
// signature. It never runs any of the referenced functions.
func TestFacadeSignatures(t *testing.T) {
	t.Helper()
	var (
		_ func() []string                                       = SupportedEcosystems
		_ func(string) bool                                     = IsSupportedEcosystem
		_ func(string) Endpoint                                 = EndpointCurrent
		_ func(string, int64) (*Catalog, error)                 = LoadCatalog
		_ func(io.Writer, io.Writer, string) *Emitter           = OutputNew
		_ func(context.Context, ScanConfig) (ScanResult, error) = ScanRun
	)
	_ = []string{
		SchemaVersion, ScannerName,
		RecordTypePackage, RecordTypeFinding, RecordTypeScanSummary, RecordTypeDiagnostic,
		ScanStatusComplete, ScanStatusPartial, ScanStatusError,
		ProfileBaseline, ProfileProject, ProfileDeep,
		RootKindGlobalPackage, RootKindUserPackage, RootKindProject, RootKindEditorExtension,
		RootKindBrowserExtension, RootKindMCPConfig, RootKindAgentSkill, RootKindHomebrew,
		RootKindDeepHome, RootKindUnknown,
	}
	var (
		_ Record
		_ Finding
		_ ScanSummary
		_ SummaryRoot
		_ Root
	)
	_ = func(em *Emitter, s ScanSummary) error { return em.EmitSummary(s) }
	_ = func(cfg ScanConfig) *Emitter { return cfg.Emitter }
}
