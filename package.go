package gobrew

// Package represents a package installed by Homebrew. It was designed to encapsulate Homebrew's JSON output.
type Package struct {
	Name              string `json:"name"`
	FullName          string `json:"full_name"`
	Tap               string `json:"tap"`
	Oldname           any    `json:"oldname"`
	Oldnames          []any  `json:"oldnames"`
	Aliases           []any  `json:"aliases"`
	VersionedFormulae []any  `json:"versioned_formulae"`
	Desc              string `json:"desc"`
	License           string `json:"license"`
	Homepage          string `json:"homepage"`
	Versions          struct {
		Stable string `json:"stable"`
		Head   string `json:"head"`
		Bottle bool   `json:"bottle"`
	} `json:"versions"`
	Urls struct {
		Stable struct {
			URL      string `json:"url"`
			Tag      any    `json:"tag"`
			Revision any    `json:"revision"`
			Using    any    `json:"using"`
			Checksum string `json:"checksum"`
		} `json:"stable"`
		Head struct {
			URL    string `json:"url"`
			Branch string `json:"branch"`
			Using  any    `json:"using"`
		} `json:"head"`
	} `json:"urls,omitempty"`
	Revision      int `json:"revision"`
	VersionScheme int `json:"version_scheme"`
	Bottle        struct {
		Stable struct {
			Rebuild int    `json:"rebuild"`
			RootURL string `json:"root_url"`
			Files   struct {
				Arm64Sonoma struct {
					Cellar string `json:"cellar"`
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"arm64_sonoma"`
				Arm64Ventura struct {
					Cellar string `json:"cellar"`
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"arm64_ventura"`
				Arm64Monterey struct {
					Cellar string `json:"cellar"`
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"arm64_monterey"`
				Sonoma struct {
					Cellar string `json:"cellar"`
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"sonoma"`
				Ventura struct {
					Cellar string `json:"cellar"`
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"ventura"`
				Monterey struct {
					Cellar string `json:"cellar"`
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"monterey"`
				X8664Linux struct {
					Cellar string `json:"cellar"`
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"x86_64_linux"`
			} `json:"files"`
		} `json:"stable"`
	} `json:"bottle,omitempty"`
	PourBottleOnlyIf        any      `json:"pour_bottle_only_if"`
	KegOnly                 bool     `json:"keg_only"`
	KegOnlyReason           any      `json:"keg_only_reason"`
	Options                 []any    `json:"options"`
	BuildDependencies       []string `json:"build_dependencies"`
	Dependencies            []any    `json:"dependencies"`
	TestDependencies        []any    `json:"test_dependencies"`
	RecommendedDependencies []any    `json:"recommended_dependencies"`
	OptionalDependencies    []any    `json:"optional_dependencies"`
	UsesFromMacos           []any    `json:"uses_from_macos"`
	UsesFromMacosBounds     []any    `json:"uses_from_macos_bounds"`
	Requirements            []any    `json:"requirements"`
	ConflictsWith           []any    `json:"conflicts_with"`
	ConflictsWithReasons    []any    `json:"conflicts_with_reasons"`
	LinkOverwrite           []any    `json:"link_overwrite"`
	Caveats                 any      `json:"caveats"`
	Installed               []struct {
		Version               string `json:"version"`
		UsedOptions           []any  `json:"used_options"`
		BuiltAsBottle         bool   `json:"built_as_bottle"`
		PouredFromBottle      bool   `json:"poured_from_bottle"`
		Time                  int    `json:"time"`
		RuntimeDependencies   []any  `json:"runtime_dependencies"`
		InstalledAsDependency bool   `json:"installed_as_dependency"`
		InstalledOnRequest    bool   `json:"installed_on_request"`
	} `json:"installed"`
	LinkedKeg          string `json:"linked_keg"`
	Pinned             bool   `json:"pinned"`
	Outdated           bool   `json:"outdated"`
	Deprecated         bool   `json:"deprecated"`
	DeprecationDate    any    `json:"deprecation_date"`
	DeprecationReason  any    `json:"deprecation_reason"`
	Disabled           bool   `json:"disabled"`
	DisableDate        any    `json:"disable_date"`
	DisableReason      any    `json:"disable_reason"`
	PostInstallDefined bool   `json:"post_install_defined"`
	Service            any    `json:"service"`
	TapGitHead         string `json:"tap_git_head"`
	RubySourcePath     string `json:"ruby_source_path"`
	RubySourceChecksum struct {
		Sha256 string `json:"sha256"`
	} `json:"ruby_source_checksum"`
}
