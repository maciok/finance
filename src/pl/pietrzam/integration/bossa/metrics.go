package bossa

type BossaMetric int

const (
	GPW_ALL          BossaMetric = iota
	FOREIGN_INDICED  BossaMetric = iota
	CURRENCIES       BossaMetric = iota
	INVESTMENT_FUNDS BossaMetric = iota
)

func (metric BossaMetric) String() string {

	if metric < GPW_ALL || metric > INVESTMENT_FUNDS {
		return "Unknown"
	}

	return metricNames[metric]
}

var metricNames = [...]string{
	"GPW market quotes",
	"Foreign market indices",
	"NBP currencies exchange rate",
	"Investment funds net value",
}

var metricToUri = map[BossaMetric]string{
	GPW_ALL:          "/metastock/mstock/sesjaall/sesjaall.prn",
	FOREIGN_INDICED:  "/indzagr/mstock/sesjazgr/sesjazgr.prn",
	CURRENCIES:       "/waluty/mstock/sesjanbp/sesjanbp.prn",
	INVESTMENT_FUNDS: "/fundinwest/mstock/sesjafun/sesjafun.prn",
}
