// librariesパッケージのテスト.
//
// use goconvey
package libraries_test

import (
	"github.com/h-tko/revel-base/libraries"
	"github.com/revel/revel"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// config.goのLoadConfigメソッドのテストケース.
//
// 正常系 ...
//
// case1: 存在するconfigを指定した場合に、ただしく、runmode設定されている方のデータが取得されること
//
// case2: runmodeのセクションが存在しない場合に、ただしく、データが取得されること
//
// 異常系 ...
//
// case3: 存在しないconfigを指定した場合、値が取れないこと
func TestLoadConfig(t *testing.T) {
	Convey("LoadConfigテスト", t, func() {
		// runmode設定
		revel.RunMode = "dev"
		revel.ConfPaths = []string{
			"./config/",
		}

		Convey("存在するconfigを指定した場合に、ただしく、runmode設定されている方のデータが取得されること", func() {
			conf, err := libraries.LoadConfig("test.conf")
			So(err, ShouldBeNil)

			test, _ := conf.String("test")
			test2, _ := conf.String("test2")
			So(test, ShouldEqual, "result_dev1")
			So(test2, ShouldEqual, "result_dev2")
		})

		Convey("runmodeのセクションが存在しない場合に、ただしく、データが取得されること", func() {
			conf, err := libraries.LoadConfig("test_norunmode.conf")
			So(err, ShouldBeNil)

			result, _ := conf.String("key")
			So(result, ShouldEqual, "norunmode")
		})

		Convey("存在しないconfigを指定した場合、値が取れないこと", func() {
			conf, _ := libraries.LoadConfig("hogehoge.conf")
			So(len(conf.Options("")), ShouldBeZeroValue)
		})
	})
}
