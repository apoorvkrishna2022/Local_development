package metrics

func SetGuageVec(packageName, serverName, methodName, env string, val int64) {
	guageVec.WithLabelValues(packageName, serverName, methodName, env).Set(float64(val))
}

func IncGaugeVec(packageName, serverName, methodName, env string) {
	guageVec.WithLabelValues(packageName, serverName, methodName, env).Inc()
}

func DecGaugeVec(packageName, serverName, methodName, env string) {
	guageVec.WithLabelValues(packageName, serverName, methodName, env).Dec()
}

func AddGaugeVec(packageName, serverName, methodName, env string, val float64) {
	guageVec.WithLabelValues(packageName, serverName, methodName, env).Add(val)
}

func SubGaugeVec(packageName, serverName, methodName, env string, val float64) {
	guageVec.WithLabelValues(packageName, serverName, methodName, env).Sub(val)
}
