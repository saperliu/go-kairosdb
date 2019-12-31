package client

import (
	"fmt"
	"github.com/saperliu/go-kairosdb/builder"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_Kairosdb(t *testing.T) {
	kairosdbClent := NewHttpClient("http://65.3.1.220:8080")
	//增加数据
	metricBuilder := builder.NewMetricBuilder()
	var metric builder.Metric;
	timeTemp := time.Now().UnixNano()/1e6

	//dateTemp := time.Unix(0, timeTemp*1e6)

	metric = metricBuilder.AddMetric(strings.ToUpper("TEST"))
	metric.AddTag("rstatus", "0")
	//boolNum :=util.IsNumeric(storePointVo.Value)
	v, err := strconv.ParseFloat("123", 64)
	if (err != nil) {
		metric.AddDataPoint(timeTemp, v);
	} else {
		metric.AddDataPoint(timeTemp, "200");
	}

	//if (storePointVo.OrgId != "") {
	//	//logger.Info("----- OrgId  %v     ", storePointVo.OrgId)
	//	metric.AddTag(KAIROSDB_TAG_ORG, storePointVo.OrgId);
	//}
	//if (storePointVo.SiteId != "") {
	//	//logger.Info("----- SiteId  %v  ", storePointVo.SiteId)
	//	metric.AddTag(KAIROSDB_TAG_SITE, storePointVo.SiteId);
	//}
	//if (storePointVo.PositionType != "") {
	//	//logger.Info("----- PositionType  %v  ", storePointVo.PositionType)
	//	metric.AddTag(KAIROSDB_TAG_POSITION_TYPE, storePointVo.PositionType);
	//}

	//mb := kairosdbClent.builderPointMetric(storePointVo)
	fmt.Printf("-----metricBuilder result ---  %v      %v \n", metricBuilder, err)
	respone, err := kairosdbClent.PushMetrics(metricBuilder)

	fmt.Printf("-----save result ---  %v      %v\n", respone, err)
}
