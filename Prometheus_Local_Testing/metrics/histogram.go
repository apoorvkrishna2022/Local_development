package metrics

func ObserveHistogramVec(packageName, serverName, methodName, code string, val int64) {
	histogramVec.WithLabelValues(packageName, serverName, methodName, code, env).Observe(float64(val))
}
