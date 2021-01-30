/*
* Author: yuansudong
* Date: 2021-1-21
* Description: 邀请码相关
*
 */

package invation

import "container/list"

// Decode 将邀请码转换为一个数字
func Decode(str string) (
	int64, //  数字
	bool, // [true,false] true代表转码成功,false代表转码失败
) {
	if str == _EmptyString {
		return 0, false
	}
	store := []byte(str)
	length := len(store)
	value := int64(0)
	for index := length; index > 0; index-- {
		pos := index - 1
		num, isExists := _BaseMap[store[pos]]
		if !isExists {
			return value, false
		}
		value += int64(num) * _GetDividend(length-index, _Binary)
	}
	return value, true
}

// Encode 对数字进行编码
func Encode(num int64) string {
	store := make([]byte, 0, _DefaultNum)
	l := list.New()
	_Encode(num, l, 1)
	// 查看位数填充,序列集合上抽像的0值
	for index := l.Len(); index < _DefaultNum; index++ {
		store = append(store, _Base[0])
	}
	for curr := l.Front(); curr != nil; curr = curr.Next() {
		store = append(store, curr.Value.(byte))
	}
	return string(store)
}

// _Encode 是Encode 的辅助函数
func _Encode(encodeNum int64, l *list.List, index int) {
	// merchant 用于求商
	merchant := encodeNum / _Binary
	// remainder 用于求余
	remainder := encodeNum % _Binary
	// 余数是上一位的的取值,商数是当前位的取值
	l.PushFront(_Base[int(remainder)])
	// 检查商是否是需要进1, 如果需要进1, 需要进入下一轮
	if merchant >= _Binary {
		_Encode(merchant, l, index+1)
	} else {
		if merchant != 0 {
			l.PushFront(_Base[int(merchant)])
		}
	}
}

// _GetDividend 用于获取被除数
func _GetDividend(c int, bin int64) int64 {
	sum := int64(1)
	for index := 0; index < c; index++ {
		sum *= bin
	}
	return sum
}
