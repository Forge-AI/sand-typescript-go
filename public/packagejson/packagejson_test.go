package packagejson_test

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/Forge-AI/sand-typescript-go/public/packagejson"
	"github.com/Forge-AI/sand-typescript-go/public/parser"
	"github.com/Forge-AI/sand-typescript-go/public/repo"
	"github.com/Forge-AI/sand-typescript-go/public/testutil/filefixture"
	"github.com/Forge-AI/sand-typescript-go/public/tspath"

	json2 "github.com/go-json-experiment/json"
)

var packageJsonFixtures = []filefixture.Fixture{
	filefixture.FromFile("package.json", filepath.Join(repo.RootPath, "package.json")),
	filefixture.FromFile("date-fns.json", filepath.Join(repo.TestDataPath, "fixtures", "packagejson", "date-fns.json")),
}

func BenchmarkPackageJSON(b *testing.B) {
	for _, f := range packageJsonFixtures {
		f.SkipIfNotExist(b)
		content := []byte(f.ReadFile(b))
		b.Run("UnmarshalJSON", func(b *testing.B) {
			b.Run(f.Name(), func(b *testing.B) {
				for b.Loop() {
					var p packagejson.Fields
					if err := json.Unmarshal(content, &p); err != nil {
						b.Fatal(err)
					}
				}
			})
		})

		b.Run("UnmarshalJSONV2", func(b *testing.B) {
			b.Run(f.Name(), func(b *testing.B) {
				for b.Loop() {
					var p packagejson.Fields
					if err := json2.Unmarshal(content, &p); err != nil {
						b.Fatal(err)
					}
				}
			})
		})

		b.Run("ParseJSONText", func(b *testing.B) {
			b.Run(f.Name(), func(b *testing.B) {
				fileName := "/" + f.Name()
				for b.Loop() {
					parser.ParseJSONText(fileName, tspath.Path(fileName), string(content))
				}
			})
		})
	}
}
