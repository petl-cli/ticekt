package commands

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// skillsFS holds one markdown file per command group plus a `_global.md` reference,
// baked into the binary at build time. The studio is the source of truth for these
// files — regenerate to pick up edits.
//
//go:embed skills/*.md
var skillsFS embed.FS

// skillFrontmatter is the small machine-readable header at the top of every skill.
// It is hand-parsed (no YAML dependency) so `skills list` stays cheap: only the
// first few lines of each file are read, never the full body.
type skillFrontmatter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	WhenToUse   string `json:"when_to_use"`
}

var skillsCmd = &cobra.Command{
	Use:   "skills",
	Short: "Browse skill files describing each command group for AI agents",
	Long: `Skills are per-group instruction files an agent can load on demand.

Typical agent flow:
  1. discovery-api skills list                  # cheap index — names + when_to_use
  2. discovery-api skills get <group>           # full markdown body for one group
  3. discovery-api <group> <command> --schema   # exact JSON schema for the call

Use ` + "`_global`" + ` for cross-cutting reference (auth, output flags, exit codes).`,
}

var skillsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print a JSON index of all available skills",
	Long: `Returns one line per skill with name, description, and when_to_use.
Reads only the frontmatter of each embedded file — safe to call from agents
that need to survey 20+ groups before committing to one.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := listSkills()
		if err != nil {
			return err
		}
		out, err := json.MarshalIndent(entries, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(out))
		return nil
	},
}

var skillsGetCmd = &cobra.Command{
	Use:   "get <name>",
	Short: "Print the full markdown body of a single skill",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		body, err := readSkill(name)
		if err != nil {
			return err
		}
		fmt.Print(body)
		return nil
	},
}

func init() {
	skillsCmd.AddCommand(skillsListCmd)
	skillsCmd.AddCommand(skillsGetCmd)
	rootCmd.AddCommand(skillsCmd)
}

// listSkills walks the embedded skills directory and returns each file's frontmatter.
func listSkills() ([]skillFrontmatter, error) {
	files, err := fs.ReadDir(skillsFS, "skills")
	if err != nil {
		return nil, fmt.Errorf("reading embedded skills: %w", err)
	}
	out := make([]skillFrontmatter, 0, len(files))
	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".md") {
			continue
		}
		body, err := fs.ReadFile(skillsFS, "skills/"+f.Name())
		if err != nil {
			return nil, err
		}
		fm := parseFrontmatter(body)
		if fm.Name == "" {
			fm.Name = strings.TrimSuffix(f.Name(), ".md")
		}
		out = append(out, fm)
	}
	// Pin the global skill first; alphabetize the rest.
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Name == "_global" {
			return true
		}
		if out[j].Name == "_global" {
			return false
		}
		return out[i].Name < out[j].Name
	})
	return out, nil
}

// readSkill returns the full markdown content for one skill by name.
func readSkill(name string) (string, error) {
	body, err := fs.ReadFile(skillsFS, "skills/"+name+".md")
	if err != nil {
		return "", fmt.Errorf("skill %q not found — run `%s skills list` to see available skills", name, "discovery-api")
	}
	return string(body), nil
}

// parseFrontmatter extracts the leading --- ... --- block as a flat key:value map.
// Tolerates files without frontmatter (returns zero value).
func parseFrontmatter(body []byte) skillFrontmatter {
	s := string(body)
	if !strings.HasPrefix(s, "---\n") {
		return skillFrontmatter{}
	}
	end := strings.Index(s[4:], "\n---")
	if end < 0 {
		return skillFrontmatter{}
	}
	header := s[4 : 4+end]
	var fm skillFrontmatter
	for _, line := range strings.Split(header, "\n") {
		k, v, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		key := strings.TrimSpace(k)
		val := strings.TrimSpace(v)
		switch key {
		case "name":
			fm.Name = val
		case "description":
			fm.Description = val
		case "when_to_use":
			fm.WhenToUse = val
		}
	}
	return fm
}
