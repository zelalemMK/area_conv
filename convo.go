package areaconv

func KMSquaredToAcre(k KMeterSquared) Acre {
	return Acre(k * 247.105)
}

func AcreToKMSquared(a Acre) KMeterSquared {
	return KMeterSquared(a / 247.105)
}

func MToKSquared(m MeterSquared) KMeterSquared {
	return KMeterSquared(m / 1000000)
}

func KToMSquared(k KMeterSquared) MeterSquared {
	return MeterSquared(k * 1000000)
}

func AcreToMSquared(a Acre) MeterSquared {
	return MeterSquared(a * 0.000247105)
}
func MSquareToAcre(m MeterSquared) Acre {
	return Acre(m * 4046.85642)
}
