package eth

import (
	"context"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/h2non/gock.v1"
)

func TestCase(t *testing.T) {
	Convey("try to test ethereum client", t, func() {
		//
		v := viper.New()
		l := zap.S()

		// Defaults:
		const base = "http://localhost:8545"
		v.SetDefault("eth.address", base)

		cfg, err := NewDefaultConfig(v)
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)

		cli, err := NewClient(cfg, l)
		So(err, ShouldBeNil)
		So(cli, ShouldNotBeNil)

		ctx := context.Background()

		cli.cli.Transport = gock.DefaultTransport

		Convey("getBlockByNumber should be ok", func() {
			gock.
				New(base).
				Post("/").
				Reply(http.StatusOK).
				AddHeader("Content-Type", MIMEApplicationJSON).
				File("fixtures/blockByNumber.json")

			bl, err := cli.GetBlockByNumber(ctx, 0)
			So(err, ShouldBeNil)
			So(bl, ShouldNotBeNil)
			So(bl.Transactions, ShouldHaveLength, 1)
			So(bl.Transactions[0], ShouldNotBeNil)
		})

		Convey("blockNumber should be ok", func() {
			gock.
				New(base).
				Post("/").
				Reply(http.StatusOK).
				AddHeader("Content-Type", MIMEApplicationJSON).
				File("fixtures/blockNumber.json")

			bl, err := cli.GetLastBlock(ctx)
			So(err, ShouldBeNil)
			So(bl, ShouldEqual, int64(1207))
		})
	})
}
