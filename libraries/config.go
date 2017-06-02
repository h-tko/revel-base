// 全部で共通しそうな処理をまとめるpackage.
//
// ただし、業務に依存しない処理のみを書くこと
package libraries

import (
	"github.com/revel/config"
	"github.com/revel/revel"
)

// Configファイルのローダー.
//
// confディレクトリ直下を指定可能
func LoadConfig(confFileName string) (*config.Context, error) {
	conf, err := config.LoadContext(confFileName, revel.ConfPaths)

	if err != nil {
		return nil, err
	}

	if conf.HasSection(revel.RunMode) {
		conf.SetSection(revel.RunMode)
	}

	return conf, nil
}
