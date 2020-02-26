package main
import (
	"encoding/json"
	"fmt"
	//  "strconv"
	//  "strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)
//milk数据结构体
type MilkInfo struct{
	MilkID string `json:MilkID`                             	//牛奶溯源编码
	MilkProInfo InitProInfo `json:MilkProInfo`                  //牛奶生产信息
	MilkIngInfo []IngInfo `json:MilkIngInfo`                 	 //牛奶信息
	MilkLogInfo LogInfo `json:MilkLogInfo`
	MilkCowBreed CowBreed`json:MilkCowBreed`            		//奶牛养殖信息
	MilkEntWare EntWare`json:MilkEntWare`						//入库信息
	MilkOutWare OutWare`json:MilkOutWare`						//出库信息
}
type MilkAllInfo struct{
	MilkID string `json:MilkId`
	MilkProInfo InitProInfo `json:MilkProInfo`
	MilkIngInfo []IngInfo `json:MilkIngInfo`
	MilkLogInfo []LogInfo `json:MilkLogInfo`
	MilkCowBreed []CowBreed`json:MilkCowBreed`
	MilkEntWare [] EntWare`json:MilkEntWare`
	MilkOutWare []OutWare`json:MilkOutWare`
}
//初始化牛奶生产信息
type InitProInfo struct{
	MilkID string `json:MilkId`								//牛奶溯源编码
	MilkBat string `json:MilkBat`							//牛奶产品批次号
	MilkName string `json:MilkName`                         //牛奶名称
	MilkSpec string `json:MilkSpec`                         //牛奶规格
	MilkMFGDate string `json:MilkMFGDate`                   //牛奶出产日期
	MilkEXPDate string `json:MilkEXPDate`                   //牛奶保质期
	MilkLOT string `json:MilkLOT`                           //牛奶批次号  保存方法
	MilkQSID string `json:MilkQSID`                         //牛奶生产许可证编号
	MilkMFRSName string `json:MilkMFRSName`                 //牛奶生产商名称
	MilkProPrice string `json:MilkProPrice`                 //牛奶生产价格
	MilkProPlace string `json:MilkProPlace`                 //牛奶生产所在地

}
//入库信息
type EntWare struct {
	MilkID string  `json:MilkId`				//奶制品溯源编号
	EntWareNum string`json:EntWareNum`			//入库批次号
	EntWareName string`json:EntWareName`		//仓库名称
	EntMilkName string`json:EntMilkName`		//奶制品名称
	EntWareAm string `json:EntWareAm`			//入库量
	EntWareTime string`json:EntWareTime` 		//入库日期
	EntWarePerson string `json:EntWarePerson `	//操作人

}
//出库信息
type OutWare struct {
	MilkID string  `json:MilkId`				//奶制品批次号
	OutWareNum string`json:OutWareNum`			//出库批次号
	OutWareName string`json:OutWareName`		//仓库名称
	OutMilkName string`json:OutMilkName `		//奶制品名称
	OutWareAm string `json:OutWareAm`			//出库量
	OutWareLast string`json:OutWareLast`		//仓库剩余量
	OutWarePerson string `json:OutWarePerson`	//操作人
}
//奶牛养殖信息
type CowBreed struct {
	MilkID string `json:MilkID `   			//牛奶ID
	Company string `json:Company`           //奶牛养殖公司名称
	CowPlace string	`json:CowPlace`		    //养殖地方
	ConPerson string`json:ConPerson`		//公司法人
	Number string	`json:Number`		    //法人电话
	CowScale string	`json:CowScale`		    //养殖规模
	CowEpPre string  `json:CowEpPre`        //防疫合格证书编号
	ProLicNum string `json:ProLicNum` 		//工商营业执照编号
	CowBreeNum string `json:CowBreeNum`     //兽畜养殖代码证
	PollFreeNum string `json:PollFreeNum`   //无公害认证编号
	GreeNum string `json:GreeNum`           //绿色产品证书编号
	CowHeaHash string`json:CowHeaHash`      //奶牛健康体检报告
}
//牛奶配料信息
type IngInfo struct{
	IngID string `json:IngID`                               //牛奶配料ID
	IngName string `json:IngName`                           //牛奶配料名称
}
//物流信息
type LogInfo struct {
	LogDepartureTm string `json:LogDepartureTm` //出发时间
	LogArrivalTm   string `json:LogArrivalTm`   //到达时间
	LogMission     string `json:LogMission`     //处理业务(储存or运输)
	LogDeparturePl string `json:LogDeparturePl` //出发地
	LogDest        string `json:LogDest`        //目的地
	LogToSeller    string `json:LogToSeller`    //销售商
	LogStorageTm   string `json:LogStorageTm`   //存储时间
	LogMOT         string `json:LogMOT`         //运送方式
	LogCopName     string `json:LogCopName`     //物流公司名称
	LogCost        string `json:LogCost`        //运送价格
	MilkNum        string `json:MilkNum`        //物流运送号

	//费用
}
type MilkChainCode struct{
}
func (a *MilkChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (a *MilkChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn,args := stub.GetFunctionAndParameters()
	if fn == "InitProInfo"{
		return a.InitProInfo(stub,args)//如果传入的功能为InitProInfo就解析其功能
	} else if fn == "addIngInfo"{
		return a.addIngInfo(stub,args)
	} else if fn == "getMilkInfo"{
		return a.getMilkInfo(stub,args)
	}else if fn == "addLogInfo"{
		return a.addLogInfo(stub,args)
	}else if fn=="addCowBreed"{
		return a.addCowBreed(stub,args)
	}else if fn=="addMilkEntWare"{
		return a.addMilkEntWare(stub,args)
	}else if fn=="addMilkOutWare"{
		return a.addMilkOutWare(stub,args)
	} else if fn=="getMilkEntWare" {
		return a.getMilkEntWare(stub,args)
	} else if fn=="getMilkOutWare" {
		return a.getMilkOutWare(stub,args)
	} else if fn=="getCowBreed"{
		return a.getCowBreed(stub,args)
	}else if fn == "getProInfo"{
		return a.getProInfo(stub,args)
	}else if fn=="delProInfo"{
		return a.delProInfo(stub,args)
	} else if fn == "getLogInfo"{
		return a.getLogInfo(stub,args)
	}else if fn == "getIngInfo"{
		return a.getIngInfo(stub,args)
	}else if fn == "getLogInfo_l"{
		return a.getLogInfo_l(stub,args)
	}else if fn=="updateLogInfo"{
		return a.updateLogInfo(stub,args)

	}


	return shim.Error("Recevied unkown function invocation")
}

func (a *MilkChainCode) InitProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var MilkInfos MilkInfo

	if len(args)!=11{//如果参数没有10个，报错
		return shim.Error("Incorrect number of arguments.")
	}
	MilkInfos.MilkID = args[0]
	if MilkInfos.MilkID == ""{
		return shim.Error("MilkID can not be empty.")
	}
	MilkInfos.MilkProInfo.MilkBat=args[1]
	MilkInfos.MilkProInfo.MilkName = args[2]
	MilkInfos.MilkProInfo.MilkSpec = args[3]
	MilkInfos.MilkProInfo.MilkMFGDate = args[4]
	MilkInfos.MilkProInfo.MilkEXPDate = args[5]
	MilkInfos.MilkProInfo.MilkLOT = args[6]
	MilkInfos.MilkProInfo.MilkQSID = args[7]
	MilkInfos.MilkProInfo.MilkMFRSName = args[8]
	MilkInfos.MilkProInfo.MilkProPrice = args[9]
	MilkInfos.MilkProInfo.MilkProPlace = args[10]
	ProInfosJSONasBytes,err := json.Marshal(MilkInfos)
	if err != nil{
		return shim.Error(err.Error())
	}
//放入世界状态里面
	err = stub.PutState(MilkInfos.MilkID,ProInfosJSONasBytes)
	if err != nil{
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func(a *MilkChainCode) addIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

	var MilkInfos MilkInfo
	var IngInfoitem IngInfo

	if  (len(args)-1)%2 != 0 || len(args) == 1{
		return shim.Error("Incorrect number of arguments")
	}

	MilkID := args[0]
	for i :=1;i < len(args);{
		IngInfoitem.IngID = args[i]
		IngInfoitem.IngName = args[i+1]
		MilkInfos.MilkIngInfo = append(MilkInfos.MilkIngInfo,IngInfoitem)
		i = i+2
	}
	MilkInfos.MilkID = MilkID
	/*  MilkInfos.MilkIngInfo = MilkIngInfo*/
	IngInfoJsonAsBytes,err := json.Marshal(MilkInfos)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(MilkInfos.MilkID,IngInfoJsonAsBytes)
	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}
func(a *MilkChainCode)addCowBreed(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	var err error
	var MilkInfos MilkInfo


	if len(args)!=12{//如果参数没有11个，报错
		return shim.Error("Incorrect number of arguments.")
	}
	MilkInfos.MilkID = args[0]
	if MilkInfos.MilkID == ""{
		return shim.Error("MilkID can not be empty.")
	}
	MilkInfos.MilkCowBreed.CowPlace= args[1]
	MilkInfos.MilkCowBreed.Company=args[2]
	MilkInfos.MilkCowBreed.ConPerson = args[3]
	MilkInfos.MilkCowBreed.Number = args[4]
	MilkInfos.MilkCowBreed.CowScale= args[5]
	MilkInfos.MilkCowBreed.CowEpPre = args[6]
	MilkInfos.MilkCowBreed.ProLicNum = args[7]
	MilkInfos.MilkCowBreed.CowBreeNum = args[8]
	MilkInfos.MilkCowBreed.PollFreeNum = args[9]
	MilkInfos.MilkCowBreed.GreeNum=args[10]
	MilkInfos.MilkCowBreed.CowHeaHash=args[11]
	CowBreedsJSONasBytes,err := json.Marshal(MilkInfos)
	if err != nil{
		return shim.Error(err.Error())
	}

	err = stub.PutState(MilkInfos.MilkID,CowBreedsJSONasBytes)
	if err != nil{
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
func(a *MilkChainCode)addMilkEntWare(stub shim.ChaincodeStubInterface,args []string)pb.Response {
	var err error
	var MilkInfos MilkInfo
	if len(args)!=7{//如果参数没有8个，报错
		return shim.Error("Incorrect number of arguments.")
	}
	MilkInfos.MilkID = args[0]
	if MilkInfos.MilkID == ""{
		return shim.Error("MilkID can not be empty.")
	}
	MilkInfos.MilkEntWare.EntWareNum=args[1]
	MilkInfos.MilkEntWare.EntWareName=args[2]
	MilkInfos.MilkEntWare.EntMilkName=args[3]
	MilkInfos.MilkEntWare.EntWareAm=args[4]
	MilkInfos.MilkEntWare.EntWareTime=args[5]
	MilkInfos.MilkEntWare.EntWarePerson=args[6]
	EntWareJSONasBytes,err := json.Marshal(MilkInfos)
	if err != nil{
		return shim.Error(err.Error())
	}
	err = stub.PutState(MilkInfos.MilkID,EntWareJSONasBytes)
	if err != nil{
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
func(a *MilkChainCode)addMilkOutWare(stub shim.ChaincodeStubInterface,args []string)pb.Response {
	var err error
	var MilkInfos MilkInfo
	if len(args)!=7{//如果参数没有8个，报错
		return shim.Error("Incorrect number of arguments.")
	}
	MilkInfos.MilkID = args[0]
	if MilkInfos.MilkID == ""{
		return shim.Error("MilkID can not be empty.")
	}
	MilkInfos.MilkOutWare.OutWareNum=args[1]
	MilkInfos.MilkOutWare.OutWareName=args[2]
	MilkInfos.MilkOutWare.OutMilkName=args[3]
	MilkInfos.MilkOutWare.OutWareAm=args[4]
	MilkInfos.MilkOutWare.OutWareLast=args[5]
	MilkInfos.MilkOutWare.OutWarePerson=args[6]
	OutWareJSONasBytes,err := json.Marshal(MilkInfos)
	if err != nil{
		return shim.Error(err.Error())
	}
	err = stub.PutState(MilkInfos.MilkID,OutWareJSONasBytes)
	if err != nil{
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func(a *MilkChainCode)getCowBreed(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	var CowBreeds []CowBreed
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}
	// 根据MilkID查询牛奶状态

	MilkID := args[0]
	resultsIterator,err := stub.GetHistoryForKey(MilkID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext(){
		var  MilkInfos MilkInfo
		response,err :=resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value,&MilkInfos)
		if MilkInfos.MilkCowBreed.CowPlace!= ""{
			CowBreeds=append(CowBreeds,MilkInfos.MilkCowBreed)
			continue
		}
	}
	jsonsAsBytes,err := json.Marshal(CowBreeds)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}
func(a *MilkChainCode)getMilkEntWare(stub shim.ChaincodeStubInterface,args []string)pb.Response {
	var EntWares []EntWare
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}
	// 根据MilkID查询牛奶状态

	MilkID := args[0]
	resultsIterator,err := stub.GetHistoryForKey(MilkID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext(){
		var  MilkInfos MilkInfo
		response,err :=resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value,&MilkInfos)
		if MilkInfos.MilkEntWare.EntMilkName!= ""{
			EntWares=append(EntWares,MilkInfos.MilkEntWare)
			continue
		}
	}
	jsonsAsBytes,err := json.Marshal(EntWares)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonsAsBytes)
}
func(a *MilkChainCode)getMilkOutWare(stub shim.ChaincodeStubInterface,args []string)pb.Response {
	var OutWares []OutWare
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}
	// 根据MilkID查询牛奶状态

	MilkID := args[0]
	resultsIterator,err := stub.GetHistoryForKey(MilkID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext(){
		var  MilkInfos MilkInfo
		response,err :=resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value,&MilkInfos)
		if MilkInfos.MilkOutWare.OutMilkName!= ""{
			OutWares=append(OutWares,MilkInfos.MilkOutWare)
			continue
		}
	}
	jsonsAsBytes,err := json.Marshal(OutWares)
	if err != nil {
		return shim.Error(err.Error())
	}


	return shim.Success(jsonsAsBytes)
}

func(a *MilkChainCode) addLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

	var err error
	var MilkInfos MilkInfo

	if len(args)!=12{
		return shim.Error("Incorrect number of arguments.")
	}
	MilkInfos.MilkID = args[0]
	if MilkInfos.MilkID == ""{
		return shim.Error("MilkID can not be empty.")
	}
	MilkInfos.MilkLogInfo.LogDepartureTm = args[1]
	MilkInfos.MilkLogInfo.LogArrivalTm = args[2]
	MilkInfos.MilkLogInfo.LogMission = args[3]
	MilkInfos.MilkLogInfo.LogDeparturePl = args[4]
	MilkInfos.MilkLogInfo.LogDest = args[5]
	MilkInfos.MilkLogInfo.LogToSeller = args[6]
	MilkInfos.MilkLogInfo.LogStorageTm = args[7]
	MilkInfos.MilkLogInfo.LogMOT = args[8]
	MilkInfos.MilkLogInfo.LogCopName = args[9]
	MilkInfos.MilkLogInfo.LogCost = args[10]
	MilkInfos.MilkLogInfo.MilkNum = args[11]

	LogInfosJSONasBytes,err := json.Marshal(MilkInfos)
	if err != nil{
		return shim.Error(err.Error())
	}
	err = stub.PutState(MilkInfos.MilkID,LogInfosJSONasBytes)
	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}


func(a *MilkChainCode) getMilkInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}
	MilkID := args[0]
	resultsIterator,err := stub.GetHistoryForKey(MilkID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var milkAllinfo MilkAllInfo

	for resultsIterator.HasNext() {
		var MilkInfos MilkInfo
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value, &MilkInfos)
		if MilkInfos.MilkProInfo.MilkName != "" {
			milkAllinfo.MilkProInfo = MilkInfos.MilkProInfo
		} else if MilkInfos.MilkIngInfo != nil {
			milkAllinfo.MilkIngInfo = MilkInfos.MilkIngInfo
		} else if MilkInfos.MilkLogInfo.LogMission != "" {
			milkAllinfo.MilkLogInfo = append(milkAllinfo.MilkLogInfo, MilkInfos.MilkLogInfo)
		} else if MilkInfos.MilkCowBreed.CowPlace != "" {
			milkAllinfo.MilkCowBreed = append(milkAllinfo.MilkCowBreed, MilkInfos.MilkCowBreed)
		} else if MilkInfos.MilkEntWare.EntMilkName != ""{
			milkAllinfo.MilkEntWare=append(milkAllinfo.MilkEntWare,MilkInfos.MilkEntWare)
		}else if MilkInfos.MilkOutWare.OutMilkName!=""{
			milkAllinfo.MilkOutWare=append(milkAllinfo.MilkOutWare,MilkInfos.MilkOutWare)
		}

	}

	jsonsAsBytes,err := json.Marshal(milkAllinfo)
	if err != nil{
		return shim.Error(err.Error())
	}

	return shim.Success(jsonsAsBytes)
}
func(a *MilkChainCode) getProInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}
	MilkID := args[0]
	resultsIterator,err := stub.GetHistoryForKey(MilkID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	var milkProInfo InitProInfo

	for resultsIterator.HasNext(){
		var MilkInfos MilkInfo
		response,err :=resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value,&MilkInfos)
		if MilkInfos.MilkProInfo.MilkName != ""{
			milkProInfo = MilkInfos.MilkProInfo
			continue
		}
	}
	jsonsAsBytes,err := json.Marshal(milkProInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}

func(a *MilkChainCode)delProInfo(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	MilkID := args[0]
	if len(args)<1{
		return shim.Error("Incorrect number of arguents.Expecting 1(milkID")
	}
	milkKey,err:=stub.CreateCompositeKey("MilkID",[]string{"milkid",MilkID})

	if err!=nil{
		return shim.Error(err.Error())
	}
	milkAsBytes,err:=stub.GetState(milkKey)
	if err!=nil{
		return shim.Error("Fail to get milkid"+err.Error())
	}else if milkAsBytes==nil{
		return shim.Error("MilkID does not exit")
	}
	//删除MilkID
	err=stub.DelState(milkKey)
	if err!=nil{
		return shim.Error("Failed delect school:"+MilkID+err.Error())
	}
	//删除牛奶ID下所有的信息
	queryString:=fmt.Sprintf("{\"selector\":{\"milk_id\":\"%s\"}}",MilkID)
	resultsIterator,err:=stub.GetQueryResult(queryString)
	if err!=nil{
		return shim.Error("Rich query failed")
	}
	defer resultsIterator.Close()
	for i:=0;resultsIterator.HasNext();i++{
		responseRange,err:=resultsIterator.Next()
		if err!=nil{
			return shim.Error(err.Error())
		}
		err=stub.DelState(responseRange.Key)
		if err!=nil{
			return shim.Error("Failed to delect milkid:"+responseRange.Key+err.Error())
		}
	}
	/*milkidAsBytes,err:=stub.GetState(MilkID)
	if err!=nil{
		return shim.Error("Failed to get MilkID:"+err.Error())
	}else if milkidAsBytes==nil{
		return shim.Error("Milk does not exit")
	}
err=stub.DelState(MilkID)
if err!=nil{
	return shim.Error("Failed to delete milk:"+MilkID+err.Error())
}*/
return shim.Success(nil)

}
func(a *MilkChainCode) getIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

	if len(args) !=1{
		return shim.Error("Incorrect number of arguments.")
	}
	MilkID := args[0]
	resultsIterator,err := stub.GetHistoryForKey(MilkID)

	if err != nil{
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var milkIngInfo []IngInfo
	for resultsIterator.HasNext(){
		var MilkInfos MilkInfo
		response,err := resultsIterator.Next()
		if err != nil{
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value,&MilkInfos)
		if MilkInfos.MilkIngInfo != nil{
			milkIngInfo = MilkInfos.MilkIngInfo
			continue
		}
	}
	jsonsAsBytes,err := json.Marshal(milkIngInfo)
	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}

func(a *MilkChainCode) getLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

	var LogInfos []LogInfo

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}

	MilkID := args[0]
	resultsIterator,err :=stub.GetHistoryForKey(MilkID)
	if err != nil{
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()


	for resultsIterator.HasNext(){
		var MilkInfos MilkInfo
		response,err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value,&MilkInfos)
		if MilkInfos.MilkLogInfo.LogMission != ""{
			LogInfos = append(LogInfos,MilkInfos.MilkLogInfo)
		}
	}
	jsonsAsBytes,err := json.Marshal(LogInfos)
	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}

func GetMilkInfo(stub shim.ChaincodeStubInterface, MilkID string) (MilkInfo, bool) {
	var Milk MilkInfo
	b,err:=stub.GetState(MilkID)//获取状态
	if err!=nil{
		return Milk,false
	}
	if b==nil{
		return Milk,false
	}
	err=json.Unmarshal(b,&Milk)
	if err!=nil{
		return Milk,false
	}
	return Milk,true
}


func PutMilk(stub shim.ChaincodeStubInterface,Milk MilkInfo)([]byte,bool){
	b,err:=json.Marshal(Milk)
	if err!=nil{
		return nil,false
	}
	err=stub.PutState(Milk.MilkID,b)//重新更新到世界状态中
	if err!=nil {
		return nil,false
	}
	return b,true
}
func(a *MilkChainCode)updateLogInfo(stub shim.ChaincodeStubInterface,args [] string) pb.Response{

	var err error
	var MilkInfos MilkInfo

	if len(args)!=12{
		return shim.Error("Incorrect number of arguments.")
	}
	MilkInfos.MilkID = args[0]
	if MilkInfos.MilkID == ""{
		return shim.Error("MilkID can not be empty.")
	}
	//var loginfo LogInfo
	//据物流单号查询信息
	result,b1:=GetMilkInfo(stub,MilkInfos.MilkID)
	if !b1{
		return shim.Error("根据货物编号查询信息发生错误")
	}
	result.MilkLogInfo.LogDepartureTm=MilkInfos.MilkLogInfo.LogDepartureTm
	result.MilkLogInfo.LogArrivalTm=MilkInfos.MilkLogInfo.LogArrivalTm
	result.MilkLogInfo.LogMission=MilkInfos.MilkLogInfo.LogMission
	result.MilkLogInfo.LogDeparturePl=MilkInfos.MilkLogInfo.LogDeparturePl
	result.MilkLogInfo.LogDest=MilkInfos.MilkLogInfo.LogDest
	result.MilkLogInfo.LogToSeller=MilkInfos.MilkLogInfo.LogToSeller
	result.MilkLogInfo.LogStorageTm=MilkInfos.MilkLogInfo.LogStorageTm
	result.MilkLogInfo.LogMOT=MilkInfos.MilkLogInfo.LogMOT
	result.MilkLogInfo.LogCopName=MilkInfos.MilkLogInfo.LogCopName
	result.MilkLogInfo.LogCost=MilkInfos.MilkLogInfo.LogCost
	result.MilkLogInfo.MilkNum=MilkInfos.MilkLogInfo.MilkNum

	LogInfosJSONasBytes,err := json.Marshal(MilkInfos)
	if err != nil{
		return shim.Error(err.Error())
	}
	err = stub.PutState(MilkInfos.MilkID,LogInfosJSONasBytes)
	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}



func(a *MilkChainCode) getLogInfo_l(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	var Loginfo LogInfo

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.")
	}

	FoodID := args[0]
	resultsIterator,err :=stub.GetHistoryForKey(FoodID)
	if err != nil{
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()


	for resultsIterator.HasNext(){
		var MilkInfos MilkInfo
		response,err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value,&MilkInfos)
		if MilkInfos.MilkLogInfo.LogMission != ""{
			Loginfo =MilkInfos.MilkLogInfo
			continue
		}
	}
	jsonsAsBytes,err := json.Marshal(Loginfo)
	if err != nil{
		return shim.Error(err.Error ())
	}
	return shim.Success(jsonsAsBytes)
}


func main(){
	err := shim.Start(new(MilkChainCode))
	if err != nil {
		fmt.Printf("Error starting Food chaincode: %s ",err)
	}
}

