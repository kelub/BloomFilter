package BloomFilter

import (
	"github.com/spaolacci/murmur3"
)

type BloomFilterer interface{
	Add(value []byte)error
	IsExists(value []byte)(bool ,error)
}

type BloomFilter struct {
	// 哈希函数数量
	numHash  int
	// 过滤器容量
	bloomCap 	uint64
	// 过滤器数组
	blooms	[]bool
}

func InitBlooms(bf *BloomFilter){
	// 从数据存储处获取
	bf.blooms = make([]bool,bf.bloomCap,bf.bloomCap)
}

func NewBloomFilter() *BloomFilter{
	bloomCap := uint64(1000000000)
	b := &BloomFilter{
		bloomCap:bloomCap,
	}
	return b
}

func (bf *BloomFilter)Add(value []byte) error{
	for _,v := range bf.hash(value){
		index := v % bf.bloomCap
		bf.blooms[index] = true
	}
	return nil
}

func (bf *BloomFilter)IsExists(value []byte)(bool ,error){
	for _,v := range bf.hash(value){
		index := v % bf.bloomCap
		if !bf.blooms[index]{
			return false,nil
		}
	}
	return true,nil
}

func (bf *BloomFilter)hash(value []byte) []uint64{
	hashs := make([]uint64, bf.numHash)
	var sead uint32
	for i:=0;i<bf.numHash;i++{
		hash := murmur3.Sum64WithSeed(value,sead)
		hashs[i] = hash
		sead = uint32(hash)
	}
	return hashs
}