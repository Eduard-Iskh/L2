package compare

import (
	"l210-sort/internal/config"
	"l210-sort/internal/parse"
	"l210-sort/internal/sorted/month"
	"strings"
)

func Compare(a, b config.LineComp, cfg config.Config) int {
	aEl := a.CompElemS
	bEl := b.CompElemS
	if cfg.B {
		aEl = strings.TrimLeft(aEl, " \t")
		bEl = strings.TrimLeft(bEl, " \t")
	}
	switch {
	case cfg.H:
		aElem, okA := parse.ParseSuffix(aEl)
		bElem, okB := parse.ParseSuffix(bEl)
		if okA && okB {
			switch {
			case aElem < bElem:
				return -1
			case aElem > bElem:
				return 1
			default:
				return 0
			}
		} else if okA {
			return 1
		} else if okB {
			return -1
		} else {
			switch {
			case aEl < bEl:
				return -1
			case aEl > bEl:
				return 1
			default:
				return 0
			}
		}
	case cfg.M:
		aElem, okA := month.Month[strings.ToLower(aEl)]
		bElem, okB := month.Month[strings.ToLower(bEl)]
		if okA && okB {
			switch {
			case aElem < bElem:
				return -1
			case aElem > bElem:
				return 1
			default:
				return 0
			}
		} else if okA {
			return 1
		} else if okB {
			return -1
		} else {
			switch {
			case aEl < bEl:
				return -1
			case aEl > bEl:
				return 1
			default:
				return 0
			}
		}
	case cfg.N:

		if a.IsNumber && b.IsNumber {
			switch {
			case a.FCompElemS < b.FCompElemS:
				return -1
			case a.FCompElemS > b.FCompElemS:
				return 1
			default:
				return 0
			}
		} else if a.IsNumber {
			return 1
		} else if b.IsNumber {
			return -1
		} else {
			switch {
			case aEl < bEl:
				return -1
			case aEl > bEl:
				return 1
			default:
				return 0
			}
		}
	default:
		if aEl < bEl {
			return -1
		} else if aEl > bEl {
			return 1
		}
		return 0
	}
}

func Less(a, b config.LineComp, cfg config.Config) bool {
	result := Compare(a, b, cfg)

	if cfg.R {
		return result > 0
	}

	return result < 0
}

// func Less(a, b config.LineComp, cfg config.Config) bool {
// 	var result bool

// 	switch {
// 	case cfg.H:
// 		aElem, okA := parse.ParseSuffix(a.CompElemS)
// 		bElem, okB := parse.ParseSuffix(b.CompElemS)
// 		if okA && okB {
// 			result = aElem < bElem
// 		} else if okA {
// 			result = false
// 		} else if okB {
// 			result = true
// 		} else {
// 			result = a.CompElemS < b.CompElemS
// 		}
// 	case cfg.M:
// 		aElem, okA := month[a.CompElemS]
// 		bElem, okB := month[b.CompElemS]
// 		if okA && okB {
// 			result = aElem < bElem
// 		} else if okA {
// 			result = false
// 		} else if okB {
// 			result = true
// 		} else {
// 			result = a.CompElemS < b.CompElemS
// 		}
// 	case cfg.N:

// 		if a.IsNumber && b.IsNumber {
// 			result = a.FCompElemS < b.FCompElemS
// 		} else if a.IsNumber {
// 			result = false
// 		} else if b.IsNumber {
// 			result = true
// 		} else {
// 			result = a.CompElemS < b.CompElemS
// 		}
// 	default:
// 		result = a.CompElemS < b.CompElemS
// 	}
// 	if cfg.R {
// 		return !result
// 	}
// 	return result
// }
