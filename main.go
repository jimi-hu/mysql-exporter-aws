package main



import (
	"fmt"
	"github.com/HarveyJie/cloud-api-go/amazon/monitor"
	"github.com/HarveyJie/cloud-api-go/amazon/resources"
	"net/http"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/HarveyJie/mysql-exporter-aws/metrics"

)


func mysqlMetrics(){
	go func() {
		for {
            allDB ,_:= resources.ListRDS()
            startTime:=time.Now()
            endTime:=startTime.Add(-time.Second*60)
            for _,v := range allDB.DBInstances {
				if(*v.DBInstanceStatus !="available"){
					continue
				}

				result,err:=monitor.RdsConnFilterDBName(startTime,endTime,*v.DBInstanceIdentifier)
				if(err!=nil){
					fmt.Println(err)
				}else {
					if(len(result.Datapoints) > 0 ) {
						metrics.MysqlConn.WithLabelValues(*v.DBInstanceIdentifier).Set(*result.Datapoints[0].Average)
					}
				}

				result,err=monitor.RdsCPUUsagFilterDBName(startTime,endTime,*v.DBInstanceIdentifier)
				if(err!=nil){
					fmt.Println(err)
				}else {
					if(len(result.Datapoints) > 0 ) {
						metrics.MysqlCPUUsage.WithLabelValues(*v.DBInstanceIdentifier).Set(*result.Datapoints[0].Average)
					}
				}


				result,err=monitor.RdsFreeMemFilterDBName(startTime,endTime,*v.DBInstanceIdentifier)
				if(err!=nil){
					fmt.Println(err)
				}else {
					if(len(result.Datapoints) > 0 ) {
						metrics.MysqlFreeMem.WithLabelValues(*v.DBInstanceIdentifier).Set(*result.Datapoints[0].Average)
					}
				}

				result,err=monitor.RdsWIOPSFilterDBName(startTime,endTime,*v.DBInstanceIdentifier)
				if(err!=nil){
					fmt.Println(err)
				}else {
					if(len(result.Datapoints) > 0 ) {
						metrics.MysqlWIOPS.WithLabelValues(*v.DBInstanceIdentifier).Set(*result.Datapoints[0].Average)
					}
				}

				result,err=monitor.RdsRIOPSFilterDBName(startTime,endTime,*v.DBInstanceIdentifier)
				if(err!=nil){
					fmt.Println(err)
				}else {
					if(len(result.Datapoints) > 0 ) {
						metrics.MysqlRIOPS.WithLabelValues(*v.DBInstanceIdentifier).Set(*result.Datapoints[0].Average)
					}
				}

				result,err=monitor.RdsNetworkRecFilterDBName(startTime,endTime,*v.DBInstanceIdentifier)
				if(err!=nil){
					fmt.Println(err)
				}else {
					if(len(result.Datapoints) > 0 ) {
						metrics.MysqlNetworkRec.WithLabelValues(*v.DBInstanceIdentifier).Set(*result.Datapoints[0].Average)
					}
				}

				result,err=monitor.RdsNetworkTransFilterDBName(startTime,endTime,*v.DBInstanceIdentifier)
				if(err!=nil){
					fmt.Println(err)
				}else {
					if(len(result.Datapoints) > 0 ) {
						metrics.MysqlNetworkTrans.WithLabelValues(*v.DBInstanceIdentifier).Set(*result.Datapoints[0].Average)
					}
				}

				result,err=monitor.RdsReadLatencyFilterDBName(startTime,endTime,*v.DBInstanceIdentifier)
				if(err!=nil){
					fmt.Println(err)
				}else {
					if(len(result.Datapoints) > 0 ) {
						metrics.MysqlReadLatency.WithLabelValues(*v.DBInstanceIdentifier).Set(*result.Datapoints[0].Average)
					}
				}

				result,err=monitor.RdsWriteLatencyFilterDBName(startTime,endTime,*v.DBInstanceIdentifier)
				if(err!=nil){
					fmt.Println(err)
				}else {
					if(len(result.Datapoints) > 0 ) {
						metrics.MysqlWriteLatency.WithLabelValues(*v.DBInstanceIdentifier).Set(*result.Datapoints[0].Average)
					}
				}
			}


			time.Sleep(2*time.Minute)
		}

	}()
}


func main() {
	mysqlMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}