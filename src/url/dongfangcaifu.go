package url

import (
	"fmt"
	"errors"
)

func Url(curpage int, gupiaocode string, rt string) (url string, err error) {
	baseUrl := "http://hqdigi2.eastmoney.com/EM_Quote2010NumericApplication/CompatiblePage.aspx?Type=%s&stk=%s&Reference=xml&limit=0&page=%d"
	
	if len(gupiaocode) <= 0 || curpage <= 0 {
		fmt.Println("the parameter value errors")
		return "", errors.New("the parameter value errors")
	}
	
	var vtype string
	var market int
	switch gupiaocode[0] {
		case '0':
			vtype = "OB"
			market = 2
		case '6':
			vtype = "OB"
			market = 1
		case '3':
			vtype = "FS"
			market = 3
		default:
			fmt.Println("gupiaocode errors")
			return "", errors.New("gupiaocode errors")
	}

	gupiaocode = fmt.Sprintf("%s%d", gupiaocode, int(market))
	url = fmt.Sprintf(baseUrl, vtype, string(gupiaocode), int(curpage))
	if len(rt) > 0 {
		url = fmt.Sprintf("%s&rt=%s", url, rt) 
	}
	return url, nil
}