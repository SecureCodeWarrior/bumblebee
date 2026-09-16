// Package scw is the SecureCodeWarrior consumption facade for bumblebee.
//
// It contains only type aliases, constant re-exports, and function re-exports
// of symbols from bumblebee's internal packages so an external module can run
// the inventory scanner as a library. It carries no logic: every behavior
// lives either upstream or in the consuming module. Upstream renames surface
// here as compile errors.
package scw

import (
	"github.com/perplexityai/bumblebee/internal/endpoint"
	"github.com/perplexityai/bumblebee/internal/exposure"
	"github.com/perplexityai/bumblebee/internal/model"
	"github.com/perplexityai/bumblebee/internal/output"
	"github.com/perplexityai/bumblebee/internal/scanner"
)

type (
	Record      = model.Record
	Finding     = model.Finding
	ScanSummary = model.ScanSummary
	SummaryRoot = model.SummaryRoot
	Endpoint    = model.Endpoint
	Catalog     = exposure.Catalog
	Emitter     = output.Emitter
	ScanConfig  = scanner.Config
	ScanResult  = scanner.Result
	Root        = scanner.Root
)

const (
	SchemaVersion = model.SchemaVersion
	ScannerName   = model.ScannerName

	RecordTypePackage     = model.RecordTypePackage
	RecordTypeFinding     = model.RecordTypeFinding
	RecordTypeScanSummary = model.RecordTypeScanSummary
	RecordTypeDiagnostic  = model.RecordTypeDiagnostic

	ScanStatusComplete = model.ScanStatusComplete
	ScanStatusPartial  = model.ScanStatusPartial
	ScanStatusError    = model.ScanStatusError

	ProfileBaseline = model.ProfileBaseline
	ProfileProject  = model.ProfileProject
	ProfileDeep     = model.ProfileDeep

	RootKindGlobalPackage    = model.RootKindGlobalPackage
	RootKindUserPackage      = model.RootKindUserPackage
	RootKindProject          = model.RootKindProject
	RootKindEditorExtension  = model.RootKindEditorExtension
	RootKindBrowserExtension = model.RootKindBrowserExtension
	RootKindMCPConfig        = model.RootKindMCPConfig
	RootKindAgentSkill       = model.RootKindAgentSkill
	RootKindHomebrew         = model.RootKindHomebrew
	RootKindDeepHome         = model.RootKindDeepHome
	RootKindUnknown          = model.RootKindUnknown
)

var (
	SupportedEcosystems  = model.SupportedEcosystems
	IsSupportedEcosystem = model.IsSupportedEcosystem
	EndpointCurrent      = endpoint.Current
	LoadCatalog          = exposure.Load
	OutputNew            = output.New
	ScanRun              = scanner.Run
)
