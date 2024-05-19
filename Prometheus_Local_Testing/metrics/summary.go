package metrics

func ObserveSummaryVec(packageName, serverName, methodName, code string, val int64) {
	summaryVec.WithLabelValues(packageName, serverName, methodName, code, env).Observe(float64(val))
}
