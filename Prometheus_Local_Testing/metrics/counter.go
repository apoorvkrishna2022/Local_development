package metrics

func IncCounterVec(packageName, serverName, methodName, env string) {
	counterVec.WithLabelValues(packageName, serverName, methodName, env).Inc()
}

func AddCounterVec(packageName, serverName, methodName, env string, val int) {
	counterVec.WithLabelValues(packageName, serverName, methodName, env).Add(float64(val))
}
