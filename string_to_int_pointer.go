package BackengHelpersGo

import (
    "strconv"
)

func StrToIntPtr(s string) (*int) {
    i, err := strconv.Atoi(s)
    if err != nil {
        return nil
    }
    return &i
}

func StrToIntWDef(s string, defaultVal int) (int) {
    i, err := strconv.Atoi(s)
    if err != nil {
		return defaultVal
    }
    return i
}

func StrToInt(s string) (int) {
    return StrToIntWDef(s, 0)
}

func StrToBool(s string) (bool) {
    if (s == "true") {
        return true
    } else {
        return false
    }
}

