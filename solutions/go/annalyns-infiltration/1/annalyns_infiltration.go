package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

func CanSpy(k, a, p bool) bool {
    if (k || a || p) {
        return true
    }
    return false
}

func CanSignalPrisoner(a, p bool) bool {
    if (!a && p) {
        return true
    }
    return false
}

func CanFreePrisoner(k, a, p, d bool) bool {
    if d && !a {
        return true
    }
    if !k && !a && p {
        return true
    }
    return false
}

