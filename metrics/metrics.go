package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (

	MysqlConn= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_connect_total",
		Help: "mysql连接数",
	},
		[]string{"db_name"},
	)

	MysqlCPUUsage= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_cpu_usage",
		Help: "mysql cpu使用率 (%)",
	},
		[]string{"db_name"},
	)

	MysqlFreeMem= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_free_mem",
		Help: "mysql 可用内存 (Bytes)",
	},
		[]string{"db_name"},
	)

	MysqlWIOPS= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_write_iops",
		Help: "mysql 写入IOPS (次/秒)",
	},
		[]string{"db_name"},
	)

	MysqlRIOPS= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_read_iops",
		Help: "mysql 读取IOPS  (次/秒)",
	},
		[]string{"db_name"},
	)

	MysqlNetworkRec= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_network_rec",
		Help: "mysql网络接收吞吐量 (Bytes/秒)",
	},
		[]string{"db_name"},
	)

	MysqlNetworkTrans= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_network_trans",
		Help: "mysql网络传输吞吐量 (Bytes/秒)",
	},
		[]string{"db_name"},
	)

	MysqlReadLatency= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_read_latency",
		Help: "mysql读取延迟 (微秒)",
	},
		[]string{"db_name"},
	)

	MysqlWriteLatency= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_write_latency",
		Help: "mysql写入延迟 (微秒)",
	},
		[]string{"db_name"},
	)
	MysqlFreeDisk= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_free_disk",
		Help: "mysql空闲磁盘大小 (Bytes)",
	},
		[]string{"db_name"},
	)
	MysqlDiskQueueDepth= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_disk_queue_depth",
		Help: "mysql队列深度",
	},
		[]string{"db_name"},
	)
	MysqlWriteThroughput= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_write_throughput",
		Help: "mysql写入吞吐 (Bytes/秒)",
	},
		[]string{"db_name"},
	)
	MysqlReadThroughput= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_write_throughput",
		Help: "mysql读取吞吐 (Bytes/秒)",
	},
		[]string{"db_name"},
	)
	MysqlSwapUsage= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_swap_usage",
		Help: "mysql交换分区使用 (Bytes)",
	},
		[]string{"db_name"},
	)
	MysqlBinLogDiskUsage= promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mysql_binlog_disk_usage",
		Help: "mysql二进制日志使用大小 (Bytes)",
	},
		[]string{"db_name"},
	)



)
