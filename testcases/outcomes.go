package testcases

var (
	BasicOperationsResult       = make(map[int]float64)
	CombinedAdditionSubtraction = make(map[int]float64)
)

var BasicOperations []string = []string{
	"89+2",
	"36-2",
	"12*2",
	"88/4",
	"12.5+663.2",
	"365.25-22.2",
	"66.1*1.25",
	"61.58/23.8"}

var CombinedAdditionSubtractionOperations []string = []string{
	"3+5-2",
	"10-4+2",
	"1+2-3",
	"100-50+25-10",
	"-5+10-3",
	"0+0-0",

	"7.5-3.2+2.3",
	"15.5-10.2+3.1",
	"10.7-3.5+2.40",
}

func InitCombinedAdditionSubtractionOperationsMap() {
	CombinedAdditionSubtraction[0] = 6
	CombinedAdditionSubtraction[1] = 8
	CombinedAdditionSubtraction[2] = 0
	CombinedAdditionSubtraction[3] = 65
	CombinedAdditionSubtraction[4] = 2
	CombinedAdditionSubtraction[5] = 0
	CombinedAdditionSubtraction[6] = 6.6
	CombinedAdditionSubtraction[7] = 8.4
	CombinedAdditionSubtraction[8] = 9.6
	CombinedAdditionSubtraction[9] = 9.7
	CombinedAdditionSubtraction[10] = 78.3
	CombinedAdditionSubtraction[11] = 5.7
	CombinedAdditionSubtraction[12] = 5.7
	CombinedAdditionSubtraction[13] = 7.5
}

func InitBasicOperationsResultMap() {
	BasicOperationsResult[0] = 91
	BasicOperationsResult[1] = 34
	BasicOperationsResult[2] = 24
	BasicOperationsResult[3] = 22
	BasicOperationsResult[4] = 675.7
	BasicOperationsResult[5] = 343.05
	BasicOperationsResult[6] = 82.625
	BasicOperationsResult[7] = 2.587394957983193

}

type Operation struct {
	Expression string
	Result     float64
}
