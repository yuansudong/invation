package invation

// _Base 用于存储进制转换函数
var (
	_Base []byte = []byte{
		'A', 'B', 'C', '6', 'D', 'E', 'F', '2', 'G', 'H', 'J', '3', 'K', 'L', 'M', '4', 'N', 'P', '5', 'Q', 'R', 'S', '7', 'T', '8', 'U', 'V', 'W', '9', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'p', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
	// _BaseMap 用于根据值获取索引,以此来提高索引的效率
	_BaseMap map[byte]int

	// _Binary 多少进制
	_Binary int64

	// _DefaultNum 需要生成多少位的邀请码
	_DefaultNum int
)

const (
	_EmptyString = ""
)

// Init 对其进行初始化
func Init() {
	_BaseMap = make(map[byte]int)
	for index := 0; index < len(_Base); index++ {
		_BaseMap[_Base[index]] = index
	}
	_DefaultNum = 6             //  默认给6位
	_Binary = int64(len(_Base)) //  用于求出这是多少进制的
}
