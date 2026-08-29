package cli

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FormatBindReport renders a BindResult as either human-readable console
// text (format == "console" or "") or JSON (format == "json").
func FormatBindReport(result BindResult, format string) (string, error) {
	switch format {
	case "", "console":
		return formatBindConsole(result), nil
	case "json":
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return "", fmt.Errorf("marshal bind report: %w", err)
		}
		return string(data) + "\n", nil
	default:
		return "", fmt.Errorf("unknown format %q (want console or json)", format)
	}
}

func formatBindConsole(result BindResult) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Protocol: %s\n", result.ProtocolID)

	if len(result.Bindings) == 0 {
		b.WriteString("No bindings in this architecture reference this protocol.\n")
		return b.String()
	}

	for _, check := range result.Bindings {
		status := "✅ OK"
		if !check.OK() {
			status = "❌ FAILED"
		}
		fmt.Fprintf(&b, "\nBinding %q: %s\n", check.BindingID, status)
		for _, e := range check.UnknownEntities {
			fmt.Fprintf(&b, "  [error] participant key %q is not an entity in this protocol\n", e)
		}
		for _, n := range check.UnresolvedNodes {
			fmt.Fprintf(&b, "  [error] participant value %q does not resolve to a node\n", n)
		}
		for _, e := range check.UnboundEntities {
			fmt.Fprintf(&b, "  [info] entity %q has no participant in this binding\n", e)
		}
	}

	return b.String()
}
