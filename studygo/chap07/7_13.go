func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	}else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	}else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	}else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	}else if s, ok := x.(string); ok {
		return sqlQuoteString(s)
	}else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

// 用switch改写上面程序
func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(s)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}