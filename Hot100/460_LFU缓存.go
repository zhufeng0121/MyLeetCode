package Hot100

type LFUCache struct {
	Key2Value    map[int]int
	Key2Freq     map[int]int
	Freq2KeyList map[int][]int
	minFreq      int
	cap          int
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		Key2Value:    map[int]int{},
		Key2Freq:     map[int]int{},
		Freq2KeyList: map[int][]int{},
		minFreq:      0,
		cap:          capacity,
	}

}

func (this *LFUCache) increaseFreq(key int) {

	freq := this.Key2Freq[key]
	// 更新KF表
	this.Key2Freq[key]++
	// 更新FK表
	list := this.Freq2KeyList[freq]

	newList := []int{}

	for _, v := range list {
		if key != v {
			newList = append(newList, v)
		}
	}

	if len(newList) == 0 {
		delete(this.Freq2KeyList, freq)
		if freq == this.minFreq {
			this.minFreq++
		}
	}
	this.Freq2KeyList[freq] = newList
	this.Freq2KeyList[freq+1] = append(this.Freq2KeyList[freq+1], key)

}

func (this *LFUCache) removeMinFreqKey() {
	keys, _ := this.Freq2KeyList[this.minFreq]

	deleteKey := keys[0]

	keys = keys[1:]
	this.Freq2KeyList[this.minFreq] = keys

	if len(keys) == 0 {
		delete(this.Freq2KeyList, this.minFreq)
	}

	// 更新K-V 表
	delete(this.Key2Value, deleteKey)

	//更新KF表
	delete(this.Key2Freq, deleteKey)

}

func (this *LFUCache) Get(key int) int {
	value, ok := this.Key2Value[key]
	if !ok {
		return -1
	}

	// 增加key对应的Freq
	this.increaseFreq(key)
	return value

}

func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}
	_, ok := this.Key2Value[key]
	if ok {
		// 更新value
		this.Key2Value[key] = value
		this.increaseFreq(key)
		return
	}

	if this.cap == len(this.Key2Value) {
		this.removeMinFreqKey()
	}

	// 插入KV表
	this.Key2Value[key] = value
	// 插入KF表
	this.Key2Freq[key] = 1

	this.Freq2KeyList[1] = append(this.Freq2KeyList[1], key)

	// 插入新key的最小freq为1
	this.minFreq = 1

}
