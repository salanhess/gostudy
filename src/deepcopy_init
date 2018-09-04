package main

// refer to https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.5.md
import (
	jb "basic/json"
	//jboa "basic/json_oa"
	"basic"
	jl "basic/listfilterconvert"
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"reflect"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	School string
}

type Empoyee struct {
	Human
	ltd string
}

type OpenapiUrlpara struct {
	restmethod string
	objectype  string
	endstr     string //end string for openapi url ,can be used for filter
}

func (s Human) sayHi() {
	fmt.Printf("Human %s age is %d and tel is %s\n", s.name, s.age, s.phone)
}

func callrefstruct(urlpara *OpenapiUrlpara) {
	fmt.Printf("urlpara restmethod is %s\n", urlpara.restmethod)
}

func (s Empoyee) sayHi() {
	fmt.Printf("Empoyee %s work at %s,age is %d and tel is %s\n", s.name, s.ltd, s.age, s.phone)
}
func returnPointerSlice(s []int) []*int {
	var slice = make([]*int, 0)
	for _, v := range s {
		var tmp = v + 10
		slice = append(slice, &tmp)
	}
	return slice
}
func Try_init() {
	urlpara := OpenapiUrlpara{"GET", "disks", ""}
	callrefstruct(&urlpara)

	//	var CommonParams = &Human{age: 10}
	var openapi_params = &Student{School: "2zhong", Human: Human{age: 10}}
	var openapi = Student{Human: Human{age: 10}, School: "2zhong" + "lihai"}
	fmt.Println(*openapi_params, openapi)

	// var via module
	var newpara = basic.CreateVolumeParams{CommonParams: basic.CommonParams{Region: "south", TenantId: "zbs"}, Name: "bigbear"}
	fmt.Println(newpara)
}

func Try_slice() {
	var slice = []int{1, 2, 3, 4, 5, 6}
	var p = returnPointerSlice(slice)
	fmt.Println(p)

	createDisksData := new(jb.CreateDisksData)
	//	result := "{4foakip57v map[diskIds:[vol-sim1hk9b3x]]}"

	//	result := jb.RespData{RequestId: "4foakip57v", Result: "{\"diskIds\": [\"vol-n6j1mo3n2x\",\"vol-vapxv1713k\"]}"}

	sub := new(jb.CreateDisksData)
	sub.DiskIds = append(sub.DiskIds, "vol-n6j1mo3n2x", "vol-vapxv1713k")

	fmt.Println("============", sub)
	result := jb.RespData{RequestId: "4foakip57v", Result: *sub}

	jb.FormatResponse(result.Result, createDisksData) //not sure whether can get disk info,need unmarsal more
	//	fmt.Println(createDisksData)

	fmt.Println(result.Result, createDisksData)
	for i, v := range createDisksData.DiskIds {
		fmt.Printf("createDisksData.DiskIds[%d]=%s\n", i, v)
	}
}

type SharedContext struct {
	VolumeMap map[string]Empoyee
}

func NewSharedContextObject() *SharedContext {
	s := &SharedContext{}
	s.VolumeMap = make(map[string]Empoyee)
	return s
}

func Try_structinit() {
	//	var openapi = Student{Human: Human{age: 10}, School: "2zhong" + "lihai"}
	em := Empoyee{Human: Human{name: "cc", age: 12}, ltd: "jd"}
	em.Human.name = "bb"
	em.Human.age = 33
	fmt.Println(em)

	var context *SharedContext = NewSharedContextObject()
	context.VolumeMap["key1"] = em
	fmt.Println(context)
}

func Try_structreflect() {
	//	var openapi = Student{Human: Human{age: 10}, School: "2zhong" + "lihai"}
	em := Empoyee{Human: Human{name: "cc", age: 12}, ltd: "jd"}
	em.Human.name = "bb"
	em.Human.age = 33
	//	fmt.Println(em)

	var secret interface{} = em
	value := reflect.ValueOf(secret)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
}

func Md5_crypto() {

	str := "abc123"

	//方法一
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	fmt.Println(md5str1)
}

type CreateVolumeOptions struct {
	VolumeName     *string `json:"name"`
	Size           *uint64 `json:"diskSizeGB"`
	VolumeTypeName *string `json:"diskType"`
	AzName         *string `json:"az"`
	Description    *string `json:"description"`
	SnapshotId     *string `json:"snapshotId"`
}

type CreateBatchVolumesOptions struct {
	VolumeOpt   *CreateVolumeOptions `json:"diskSpec"`
	MaxCount    *int                 `json:"maxCount"`
	ClientToken *string              `json:"clientToken"`
}

type SnapSuite struct {
	CreateVolMap map[string]*CreateBatchVolumesOptions
}

var (
	OpenapiMaxCount   = 2
	clientToken       = "xxx"
	volName           = "vol1"
	OpenapiAz         = "OpenapiAz1"
	volDes            = "volDes"
	OpenapiDiskType   = "OpenapiDiskType"
	OpenapiDiskSizeGB = uint64(10)
	TenantId          = "zbs-admin"
	OpenapiRegionId   = "cn-south-1"
	OpenapiXjcloudPin = "xxx=1131"
)

var CreateVolkey1 = &CreateBatchVolumesOptions{
	MaxCount:    &OpenapiMaxCount,
	ClientToken: &clientToken,
	VolumeOpt:   &CreateVolumeOptions{VolumeName: &volName, AzName: &OpenapiAz, Description: &volDes, VolumeTypeName: &OpenapiDiskType, Size: &OpenapiDiskSizeGB},
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func CopyCreateBatchVolumesOptions(src *CreateBatchVolumesOptions) (dst *CreateBatchVolumesOptions) {
	dst = new(CreateBatchVolumesOptions)
	_ = deepCopy(dst, src)
	return
}

func main() {
	//	Try_slice()
	//basic.Four_method_initmap()
	//	jl.Listfilterconvert()
	jl.Listfilterconvert2()
	Try_structinit()
	//	Try_structreflect()

	var volmap = make(map[string]int)
	volmap["vol1"] = 1
	fmt.Printf("volmap is %v", volmap)
	fmt.Println(volmap["vol1"])
	var snaplist []string
	snaplist = make([]string, 1)
	snaplist[0] = "abc"
	fmt.Println(snaplist)
	fmt.Println("===============")
	Md5_crypto()
	fmt.Println("===============")
	s := SnapSuite{CreateVolMap: map[string]*CreateBatchVolumesOptions{"vol_oa1": CreateVolkey1}}
	for i, key := range s.CreateVolMap {
		fmt.Printf("[Checking before copy val====]CreateVolMap[%s]=%v, MaxCount=%d,VolumeOpt.VolumeTypeName=%v\n", i, key, *key.MaxCount, *key.VolumeOpt.VolumeTypeName)
	}
	invalidMaxCount := 666
	invalidDisktype := "ssdXX"
	//	var CreateVolkey2 = &CreateBatchVolumesOptions{
	//		MaxCount:    &OpenapiMaxCount,
	//		ClientToken: &clientToken,
	//		VolumeOpt:   &CreateVolumeOptions{VolumeName: &volName, AzName: &OpenapiAz, Description: &volDes, VolumeTypeName: &OpenapiDiskType, Size: &OpenapiDiskSizeGB},
	//	}

	//	var CreateVolkey3 = &CreateBatchVolumesOptions{
	//		MaxCount:    &OpenapiMaxCount,
	//		ClientToken: &clientToken,
	//		VolumeOpt:   &CreateVolumeOptions{VolumeName: &volName, AzName: &OpenapiAz, Description: &volDes, VolumeTypeName: &OpenapiDiskType, Size: &OpenapiDiskSizeGB},
	//	}
	//	//	var CreateVolkey2 = new(CreateBatchVolumesOptions)
	//	//	var CreateVolkey3 = new(CreateBatchVolumesOptions)

	//	s.CreateVolMap["case2"] = CreateVolkey2
	//	s.CreateVolMap["case2"].MaxCount = &invalidMaxCount
	//	s.CreateVolMap["case3"] = CreateVolkey3
	//	s.CreateVolMap["case3"].VolumeOpt.VolumeTypeName = &invalidDisktype

	//	for i, key := range s.CreateVolMap {
	//		fmt.Printf("[Checking .....         ====]CreateVolMap[%s]=%v, MaxCount=%d,VolumeOpt.VolumeTypeName=%v\n", i, key, *key.MaxCount, *key.VolumeOpt.VolumeTypeName)
	//	}

	//	var CreateVolkey2 = new(CreateBatchVolumesOptions)
	//	err1 := deepCopy(CreateVolkey2, *CreateVolkey1)
	//	fmt.Println("***********", err1)

	//	var CreateVolkey2 *CreateBatchVolumesOptions
	//	CreateVolkey2 = CopyCreateBatchVolumesOptions(CreateVolkey1)

	//	var CreateVolkey3 = new(CreateBatchVolumesOptions)
	//	err2 := deepCopy(CreateVolkey3, *CreateVolkey1)
	//	fmt.Println("***********", err2)

	fmt.Println("----------------------------------------")
	s.CreateVolMap["case2"] = CopyCreateBatchVolumesOptions(CreateVolkey1)
	s.CreateVolMap["case2"].MaxCount = &invalidMaxCount
	s.CreateVolMap["case3"] = CopyCreateBatchVolumesOptions(CreateVolkey1)
	s.CreateVolMap["case3"].VolumeOpt.VolumeTypeName = &invalidDisktype
	for i, key := range s.CreateVolMap {
		fmt.Printf("[Checking .....         ====]CreateVolMap[%s]=%v, MaxCount=%d,VolumeOpt.VolumeTypeName=%v\n", i, key, *key.MaxCount, *key.VolumeOpt.VolumeTypeName)
	}
	//basic.Hello_struct()
	//	age10 := 10
	//	namec := "carol"
	//	p1 := CreateBatchVolumesOptions{MaxCount: &age10, ClientToken: &namec}
	//	p2 := p1
	//	age33 := 33
	//	p2.MaxCount = &age33
	//	fmt.Printf("CreateBatchVolumesOptions1 MaxCount: %d, ClientToken: %s\n", *p1.MaxCount, *p1.ClientToken)
	//	fmt.Printf("CreateBatchVolumesOptions2 MaxCount: %d, ClientToken: %s\n", *p2.MaxCount, *p2.ClientToken)

}
