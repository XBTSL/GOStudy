package mySort

// 冒泡排序,两两排序比较，每次最大的排到后面
func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

// 选择排序，每次循环选出最小值和第一个交换位置
func SelectionSort(arry []int) []int {
	for i := 0; i < len(arry); i++ {
		minPlace := i
		for j := i + 1; j < len(arry); j++ {
			if arry[minPlace] > arry[j] {
				minPlace = j
			}
		}
		arry[minPlace], arry[i] = arry[i], arry[minPlace]
	}
	return arry
}

// 插入排序的代码实现虽然没有冒泡排序和选择排序那么简单粗暴，但它的原理应该是最容易理解的了，
// 因为只要打过扑克牌的人都应该能够秒懂。插入排序是一种最简单直观的排序算法，它的工作原理是通过构建有序序列，
// 对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
func InsertSort(arry []int) []int {
	for i := 0; i < len(arry)-1; i++ {
		preIndex := i + 1
		for preIndex > 0 && arry[preIndex] < arry[preIndex-1] {
			arry[preIndex-1], arry[preIndex] = arry[preIndex], arry[preIndex-1]
			preIndex--
		}
	}
	return arry
}

// 归并排序  两个链表合并思维
func Merge(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return MergeSort(Merge(left), Merge(right))
}

func MergeSort(a, b []int) []int {
	result := []int{}
	for len(a) != 0 && len(b) != 0 {
		if a[0] > b[0] {
			result = append(result, b[0])
			b = b[1:]
		} else {
			result = append(result, a[0])
			a = a[1:]
		}
	}
	if len(a) == 0 {
		result = append(result, b...)
	}
	if len(b) == 0 {
		result = append(result, a...)
	}
	return result
}

// 快速排序
func Quick(arry []int) []int {
	return QuickSort(arry, 0, len(arry)-1)
}

func QuickSort(arry []int, left, right int) []int {
	if left < right {
		paet := partitionWay(arry, left, right)
		QuickSort(arry, left, paet-1)
		QuickSort(arry, paet+1, right)
	}
	return arry
}

func partitionWay(arry []int, left, right int) int {
	temp := arry[left]
	for right > left {
		for right > left && arry[right] >= temp {
			right--
		}
		arry[left] = arry[right]
		for right > left && arry[left] <= temp {
			left++
		}
		arry[right] = arry[left]
	}
	arry[left] = temp
	return left
}

// 桶排序
func BucketSort(arry []int) []int {
	length := len(arry)
	maxnum, minnum := arry[0], arry[0]
	for i := 1; i < length; i++ {
		if arry[i] > maxnum {
			maxnum = arry[i]
		}
		if arry[i] < minnum {
			minnum = arry[i]
		}
	}
	maxValue := maxnum - minnum
	bucketNum := maxValue/length + 1
	bukets := make([][]int, bucketNum)
	for j := 0; j < bucketNum; j++ {
		bukets[j] = make([]int, 0)
	}
	for i := 0; i < length; i++ {
		id := (arry[i] - minnum) / length
		bukets[id] = append(bukets[id], arry[i])
	}

	arrindex := 0
	for i := 0; i < bucketNum; i++ {
		if len(bukets[i]) == 0 {
			continue
		}
		InsertSort(bukets[i])
		for j := 0; j < len(bukets[i]); j++ {
			arry[arrindex] = bukets[i][j]
			arrindex++
		}
	}
	return arry
}
