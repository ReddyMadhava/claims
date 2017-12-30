/*
Copyright TCS Ltd 2017 All Rights Reserved.
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/


package main

import (
    "fmt"
    "time"
	"bytes"
	//"os"
	//"strings"
    "strconv"
    "encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)


type Healthcare struct {}


type Claim struct {

    ObjectType string `json:"docType"`
    CLAIM_ID string `json:"claim_ID"`
    SUBSCRIBER_FULLNAME string `json:"subcriber_name"`
    SUBSCRIBER_GENDER string `json:"subcriber_gender"`
    SUBSCRIBERID string `json:"subcriber_id"`
    PROVIDER_FULLNAME string `json:"provider_name"`
    PROVIDERID string `json:"provider_id"`
    PAYERNAME string `json:"payers_name"`
    PAYERID string `json:"payers_id"`
    SERVICETYPE_CODE string `json:"service_code"`
    DOS string `json:"appointment_date"`
    VISIT_TIME string `json:"visit_time"`
    DEDUCTIBLE int `json:"deductible"`
    AMOUNT int `json:"amount"`
    COPAY string `json:"copay"`
    ELIGIBILITY_CODE string `json:"eligibitity_code"`
    PROCEDURE_CODE int `json:"procedure_code"`
    STATUS string `json:"status"`
    OWNER string `json:"owner"`
	COMMENTS string `json:"comments"`
	SETTLEMENT_AMOUNT int `json:"settlement_amount"`
}

type REMIT struct {

    ObjectType string `json:"docType"`
    CLAIMID string `json:"docType"`
    SENDER_BANK_ACCT_NO string `json:"sender bank acc no"`
    RECEIVER_BANK_ACCT_NO string `json:"receiver bank acc no"`
    CHECK_NUM string `json:"check no"`
    PAYERID string `json:"payer id"`
    PAYEE_ID string `json:"payee id"`
    SUBSCRIBERID string `json:"subcriber id"`
    AMOUNT int `json:"amount"`
    PROCEDURE_CODE int `json:"procedure_code"`
    DOS string `json:"dos"`


}
type Procedure_Details struct {
	ObjectType string `json:"docType"`
    Procedure_code  int `json:"procedurecode"`
	Procedure_id string `json:"procedureid"`
    Limit   int `json:"limit"`
	Discount int `json:"discount"`
	Discription string `json:"description"`
}

// ===================================================================================
// Main
// ===================================================================================

func main() {
    err := shim.Start(new(Healthcare))
    if err != nil {
        fmt.Printf("Error starting Healthcare Chaincode- %s", err)
    }
}

// Init initializes chaincode
// ===========================


func(t *Healthcare) Init(stub shim.ChaincodeStubInterface) pb.Response {
    var err error
	 
    fmt.Println("starting init")
   
    var procedure = & Procedure_Details {"Procedure",1134,"PC1",1000,75,"sample procedure code"}
    procedureAsBytes1, _:= json.Marshal(procedure)
    err = stub.PutState("1134", procedureAsBytes1)
	
	var procedure1 = & Procedure_Details {"Procedure",1133,"PC2",1000,65,"sample procedure code"}
    procedureAsBytes2, _:= json.Marshal(procedure1)
    err = stub.PutState("1133", procedureAsBytes2)
	
	var procedure2 = & Procedure_Details {"Procedure",1122,"PC3",1000,55,"sample procedure code"}
    procedureAsBytes3, _:= json.Marshal(procedure2)
    err = stub.PutState("1122", procedureAsBytes3)
	
	var procedure3 = & Procedure_Details {"Procedure",1167,"PC4",1000,80,"sample procedure code"}
    procedureAsBytes4, _:= json.Marshal(procedure3)
    err = stub.PutState("1167", procedureAsBytes4)
	
	var procedure4 = & Procedure_Details {"Procedure",1187,"PC5",1000,65,"sample procedure code"}
    procedureAsBytes5, _:= json.Marshal(procedure4)
    err = stub.PutState("1187", procedureAsBytes5)
	
	var procedure5 = & Procedure_Details {"Procedure",1188,"PC6",1000,60,"sample procedure code"}
    procedureAsBytes6, _:= json.Marshal(procedure5)
    err = stub.PutState("1188", procedureAsBytes6)
	
	
	var procedure6 = & Procedure_Details {"Procedure",1199,"PC7",1000,50,"sample procedure code"}
    procedureAsBytes7, _:= json.Marshal(procedure6)
    err = stub.PutState("1199", procedureAsBytes7)
	
    if err != nil {
        fmt.Println("Could not store procedure")
        return shim.Error(err.Error())
    }

    fmt.Println("- end init_claim")
    return shim.Success(nil)
}


// Invoke - Our entry point for Invocations
// ========================================


func(t *Healthcare) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println(" ")
    fmt.Println("starting invoke, for - " + function)

    if function == "init" {
        return t.Init(stub)
    } else if
	
	function == "schedule_appointment" {//schedule_appointment for Insurance holder
        return t.schedule_appointment(stub, args)
    }else if
	
	function == "add_procedure" {//schedule_appointment for Insurance holder
        return t.add_procedure(stub, args)
    }else if
	
	function == "getDataByProcedureId" {//schedule_appointment for Insurance holder
        return t.getDataByProcedureId(stub, args)
    }else if
	
    function == "Check_in" {//update chein details to claim
        return t.Check_in(stub, args)
    } else if

    function == "getDataByClaimId" {//get claim details with claim ID 
        return t.getDataByClaimId(stub, args)
    } else if

    function == "approve_cliam" {//approve_cliam for Insurance holder
        return t.approve_cliam(stub, args)
    } else if

    function == "getDataByPatientName" {// get claim details by subcriber name
        return t.getDataByPatientName(stub, args)
    } else if

    function == "getDataByPatientId" {//get claim details by subcriber id
        return t.getDataByPatientId(stub, args)
    } else if

    function == "getDataByStatus" {//get claim details by status
        return t.getDataByStatus(stub, args)
    } else if

    function == "getDataByPayersname" {//get claim ckdetails by payer name 
        return t.getDataByPayersname(stub, args)
    } else if

    function == "getDataByPayersId" {//get claim details by payer id
        return t.getDataByPayersId(stub, args)
    } else if

    function == "getallclaim" {//get all claim
        return t.getallclaim(stub)
    } else if

    function == "getDataByOwner" {//get claim details by owner
        return t.getDataByOwner(stub, args)
    } else if

    function == "getRemitData" {//get remit details
        return t.getRemitData(stub, args)
    } else if

    function == "getDataByIDstatus" {//get claim deatils by claim id and status
        return t.getDataByIDstatus(stub, args)
    }else if

    function == "getHistoryForClaim" {//get History For Claim
        return t.getHistoryForClaim(stub, args)
    }

    fmt.Println("Received unknown invoke function name - " + function)
    return shim.Error("Received unknown invoke function name - '" + function +"'")
}
func(t *Healthcare) add_procedure(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	
	if len(args) != 5{
        return shim.Error("Incorrect number of arguments. Expecting 5")
    }
	
	Procedure_code,_ :=strconv.Atoi(args[0])
	Procedure_id:=args[1]
	Limit,_:=strconv.Atoi(args[2])
	Discount,_:=strconv.Atoi(args[3])
	Discription:=args[4]
	// ==== Check if procedure already exists ====
	
	procedureAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get procedure: " + err.Error())
	} else if procedureAsBytes != nil {
		fmt.Println("This procedure already exists: " + args[0])
		return shim.Error("This procedure already exists: " + args[0])
	}
	
    
    fmt.Println("starting add procedure")
    
    var procedure = & Procedure_Details {"Procedure",Procedure_code,Procedure_id,Limit,Discount,Discription}
    procedureAsBytes1, _:= json.Marshal(procedure)
    err = stub.PutState(args[0], procedureAsBytes1)
    if err != nil {
        fmt.Println("Could not store procedure")
        return shim.Error(err.Error())
    }

    fmt.Println("- end init_claim")
    return shim.Success(nil)

}
// ============================================================
// schedule_appointment - create a appointment for subcriber
// ============================================================
func(t *Healthcare) schedule_appointment(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 14 {
        return shim.Error("Incorrect number of arguments. Expecting 14")
    }
	
	claim_id :=args[0]//c0
	
	// ==== Check if claim already exists ====
	
	claimAsBytes, err := stub.GetState(claim_id)
	if err != nil {
		return shim.Error("Failed to get claim: " + err.Error())
	} else if claimAsBytes != nil {
		fmt.Println("This claim already exists: " + claim_id)
		return shim.Error("This claim already exists: " + claim_id)
	}
	
    subscriber_fullname :=args[1]//Pranali karekar
    subscriber_gender :=args[2]//Female
    subscriberid :=args[3]//PK101
    provider_fullname :=args[4]//Fortis
    providerid :=args[5]//HF1
    payername :=args[6]//ICICI
    payerid := args[7]//IC1
    servicetype_code :=args[8]//30
    dos :=args[9]//2017-05-28
    deductible, _ :=strconv.Atoi(args[10])//500
    copay :=args[11]//45
    eligibility_code :=args[12]//1
    status:= "Scheduled"
    owner:=provider_fullname
	comments:=args[13]

    
    fmt.Println("starting create_claim")

    
    var claim = & Claim {
	"claim",claim_id ,subscriber_fullname ,subscriber_gender,subscriberid,provider_fullname,providerid,payername,payerid,servicetype_code,dos,"",deductible,0,copay,eligibility_code,0,status,owner,comments,0}
    claimAsBytes1, _:= json.Marshal(claim)
    err = stub.PutState(args[0], claimAsBytes1)
    if err != nil {
        fmt.Println("Could not store claim")
        return shim.Error(err.Error())
    }

    fmt.Println("- end init_claim")
    return shim.Success(nil)

}
// ============================================================
// Check_in - update checkin deatails to claim- time ,date ,deductible,status
// ============================================================

func(t *Healthcare) Check_in(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
	if len(args) != 3 {
        return shim.Error("Incorrect number of arguments. Expecting 3")
    }
	claim_id :=args[0]//c0
    deductible, _  :=strconv.Atoi(args[1])
	datenow:=args[2]
	
	// ==== Check if claim already exists ====
	var claim Claim
	
    claimAsBytes, err := stub.GetState(claim_id)
    
	
	if err != nil {
		return shim.Error("Failed to get claim: " + err.Error())
	} else if claimAsBytes != nil {
		if claim.STATUS=="CheckIn"{
		return shim.Error("This claim already exists: " + claim_id)
		}
	}
	
	json.Unmarshal(claimAsBytes, & claim)
    fmt.Println("starting checkin_claim")

	//update current date and time 
    timenow := time.Now().Format("3:04PM")
	//datenow := time.Now().Format(time.RFC3339)

    claim.VISIT_TIME = timenow
	claim.DEDUCTIBLE = deductible
	claim.DOS=datenow
	
	//update status
    claim.STATUS = "CheckIn"
	
	
    jsonAsBytes, _:= json.Marshal(claim)
    err = stub.PutState(args[0], jsonAsBytes)

    if err != nil {
        fmt.Println("Could not Update Claim")
        return shim.Error(err.Error())
    }

    fmt.Println("- end init_claim")
    return shim.Success(nil)
}
// ============================================================
// approve_cliam - approve claim if amount is less than 1000$ else upadte status to pendding
// ============================================================
func(t *Healthcare) approve_cliam(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	if len(args) != 3 {
        return shim.Error("Incorrect number of arguments. Expecting 3")
    }
	amt, _  :=strconv.Atoi(args[2])
    Procedure_code, _  :=strconv.Atoi(args[1])
	claim_id:=args[0]

    var claim Claim

	// ==== Check if claim already exists ====
	
    claimAsBytes, err := stub.GetState(claim_id)
		
	if err != nil {
		return shim.Error("Failed to get claim: " + err.Error())
	} else if claimAsBytes != nil {
		json.Unmarshal(claimAsBytes, & claim)
		if claim.STATUS=="Approved"||claim.STATUS=="Pending"||claim.STATUS=="Decline"{
		return shim.Error("This claim already exists: " + claim_id)
		}
	}
	
    damt := claim.DEDUCTIBLE
    New_Owner_ID := claim.PAYERNAME
	
	// ==== Check product details ====
	
	var procedure Procedure_Details
	productAsBytes, err := stub.GetState(args[1])
    
	if err != nil {
	return shim.Error("Failed to get procedure: " + err.Error())
	}
	
	if productAsBytes == nil {
	claim.STATUS = "Decline"
	claim.DEDUCTIBLE = damt
	claim.OWNER = New_Owner_ID
	claim.PROCEDURE_CODE=Procedure_code
	claim.COMMENTS="Error getting procedure code"
	claim.AMOUNT = amt
	jsonAsBytes, _ := json.Marshal(claim)

	err = stub.PutState(args[0], jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	}
	
	if productAsBytes != nil{
	json.Unmarshal(productAsBytes, & procedure)
	
	limit:=procedure.Limit
	discount:=procedure.Discount
	
	if amt <= limit{
	
	var remitamount int
	var newdeductiable int
	
	//check Deductiable is greter than amount and calculate new deductible and remit amount
	if amt<=damt{
	remitamount= 0
	newdeductiable=damt-amt
	}
	
	//check Deductiable is less than amount and calculate new deductible and remit amount
	if damt<amt{
	remitamount1:=amt-damt
	finaldis:=float64(discount)/float64(100)
	remitamount=int(float64(remitamount1)*finaldis)
	newdeductiable=0
	}
	
	
	//remitamount1 := strconv.Itoa(remitamount)
	//newdeductiable1 := strconv.Itoa(newdeductiable)
	
	
	claim.STATUS = "Approved"
	claim.DEDUCTIBLE = newdeductiable
	claim.AMOUNT = amt
	claim.OWNER = New_Owner_ID
	claim.PROCEDURE_CODE=Procedure_code
	claim.COMMENTS="You receive refound within next 3-4 working days"
	claim.SETTLEMENT_AMOUNT=remitamount
	
	jsonAsBytes,_ := json.Marshal(claim)
	err = stub.PutState(args[0], jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	t := time.Now()
	
	// create remit receipt
	claimid := "R" + args[0]
	sender_acc_no := "s" + t.Format("20060102150405")
	receiver_acc_no := "r" + t.Format("20060102150405")
	checkno := "c" + t.Format("20060102150405")
	payeeid := claim.PROVIDERID
	subcriberid := claim.SUBSCRIBERID
	procedure_code := Procedure_code
	dos := claim.DOS

	var remit = & REMIT {
	"remit", claimid, sender_acc_no, receiver_acc_no, checkno, New_Owner_ID, payeeid, subcriberid, remitamount, procedure_code, dos}
	remitAsBytes,_:= json.Marshal(remit)
	err = stub.PutState(claimid, remitAsBytes)
	
	if err != nil {
		fmt.Println("Could not store claim")
		return shim.Error(err.Error())
	}
	}else{
	claim.STATUS = "Pending"
	claim.DEDUCTIBLE = damt
	claim.OWNER = New_Owner_ID
	claim.PROCEDURE_CODE=Procedure_code
	claim.COMMENTS="Amount grater than 1000$"
	claim.AMOUNT = amt
	
	jsonAsBytes, _ := json.Marshal(claim)

	err = stub.PutState(args[0], jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	}
    }
	
    fmt.Println("- end init_hospital")
    return shim.Success(nil)

}

// ============================================================
// getRemitData -remit details by claim id
// ============================================================
func(t *Healthcare) getRemitData(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

    if len(id) != 1 {
        return shim.Error("Please provide Claim Id")
    }
	Rid:="R"+id[0]
    claimAsBytes, err := stub.GetState(Rid)
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(claimAsBytes)
}
// ============================================================
// getDataByClaimId -Claim details by id C0
// ============================================================
func(t *Healthcare) getDataByClaimId(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

    if len(id) != 1 {
        return shim.Error("Please provide Claim Id")
    }
    claimAsBytes, err := stub.GetState(id[0])
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(claimAsBytes)
}

// ============================================================
// getDataByOwner -Claim details by Owner
// ============================================================
func(t *Healthcare) getDataByOwner(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

    type getData struct {
        Claim[] Claim `json:"claim"`
    }
    var getdata getData
	
	if id[0]=="Admin"{
    resultsIterator, err := stub.GetStateByRange("C0", "C9999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }

    defer resultsIterator.Close()

    for resultsIterator.HasNext() {

	aKeyValue, err := resultsIterator.Next()
	if err != nil {
		return shim.Error(err.Error())
	}

	queryValAsByte := aKeyValue.Value
	var claim Claim
	json.Unmarshal(queryValAsByte, & claim)
	if claim.STATUS=="Approved" || claim.STATUS=="Pending" || claim.STATUS=="Decline" {
		getdata.Claim = append(getdata.Claim, claim)
	}

    }
    result, _ := json.Marshal(getdata)
    return shim.Success(result)
	}else{
	
    resultsIterator, err := stub.GetStateByRange("C0", "C9999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    for resultsIterator.HasNext() {
        aKeyValue, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }

        queryValAsByte := aKeyValue.Value
        var claim Claim
        json.Unmarshal(queryValAsByte, & claim)

        if claim.OWNER == id[0] {
            getdata.Claim = append(getdata.Claim, claim)
        }

    
	}
    
    
}
result, _ := json.Marshal(getdata)
return shim.Success(result)
}


// ============================================================
// getDataByIDstatus -Claim details by id and status
// ============================================================

func(t *Healthcare) getDataByIDstatus(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

    type getData struct {
        Claim[] Claim `json:"claim"`
    }
    var getdata getData
    resultsIterator, err := stub.GetStateByRange("C0", "C9999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    for resultsIterator.HasNext() {

        aKeyValue, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }

        queryValAsByte := aKeyValue.Value

        var claim Claim

        json.Unmarshal(queryValAsByte, & claim)

        if (claim.STATUS == id[1]) && (claim.CLAIM_ID == id[0]) {
            getdata.Claim = append(getdata.Claim, claim)
        }

    }
    result, _ := json.Marshal(getdata)
    return shim.Success(result)

}

// ============================================================
// getDataByPatientId -Claim details by subcriber id
// ============================================================
func(t *Healthcare) getDataByPatientId(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

    type getData struct {
        Claim[] Claim `json:"claim"`
    }
    var getdata getData
    resultsIterator, err := stub.GetStateByRange("C0", "C9999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    for resultsIterator.HasNext() {
        aKeyValue, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        queryValAsByte := aKeyValue.Value
        var claim Claim
        json.Unmarshal(queryValAsByte, & claim)
        if claim.SUBSCRIBERID == id[0] {
            getdata.Claim = append(getdata.Claim, claim)
        }

    }
    result, _ := json.Marshal(getdata)
    return shim.Success(result)

}

func(t *Healthcare) getDataByProcedureId(stub shim.ChaincodeStubInterface, id[] string) pb.Response {
	
	type getData struct {
        Procedure_Details[] Procedure_Details `json:"procedure"`
    }
    var getdata getData
    resultsIterator, err := stub.GetStateByRange("1111", "999999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }
	pcode, _  :=strconv.Atoi(id[0])
    defer resultsIterator.Close()
    for resultsIterator.HasNext() {
        aKeyValue, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        queryValAsByte := aKeyValue.Value
        var procedure Procedure_Details

        json.Unmarshal(queryValAsByte, & procedure)
        if procedure.Procedure_code  == pcode {
            getdata.Procedure_Details = append(getdata.Procedure_Details, procedure)
        }

    }
    result, _ := json.Marshal(getdata)
    return shim.Success(result)

}
// ============================================================
// getDataByPayersname -Claim details by subcriber name
// ============================================================

func(t *Healthcare) getDataByPayersname(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

    type getData struct {
        Claim[] Claim `json:"claim"`
    }
    var getdata getData
    resultsIterator, err := stub.GetStateByRange("C0", "C9999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    for resultsIterator.HasNext() {
        aKeyValue, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        queryValAsByte := aKeyValue.Value
        var claim Claim
        json.Unmarshal(queryValAsByte, & claim)
        if claim.PAYERNAME == id[0] {
            getdata.Claim = append(getdata.Claim, claim)
        }

    }
    result, _ := json.Marshal(getdata)
    return shim.Success(result)
}

// ============================================================
// getDataByPayersId -Claim details by Provider id
// ============================================================
func(t *Healthcare) getDataByPayersId(stub shim.ChaincodeStubInterface, id[] string) pb.Response {
    type getData struct {
        Claim[] Claim `json:"claim"`
    }
    var getdata getData
    resultsIterator, err := stub.GetStateByRange("C0", "C9999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    for resultsIterator.HasNext() {
        aKeyValue, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        queryValAsByte := aKeyValue.Value
        var claim Claim
        json.Unmarshal(queryValAsByte, & claim)
        if claim.PAYERID == id[0] {
            getdata.Claim = append(getdata.Claim, claim)
        }
    }
    result, _ := json.Marshal(getdata)
    return shim.Success(result)

}

// ============================================================
// getDataByPayersId -Claim details by Provider name
// ============================================================
func(t *Healthcare) getDataByPatientName(stub shim.ChaincodeStubInterface, id[] string) pb.Response {
    type getData struct {
        Claim[] Claim `json:"claim"`
    }

    var getdata getData

    resultsIterator, err := stub.GetStateByRange("C0", "C9999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }

    defer resultsIterator.Close()
    for resultsIterator.HasNext() {

        aKeyValue, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }

        queryValAsByte := aKeyValue.Value
        var claim Claim
        json.Unmarshal(queryValAsByte, & claim)

        if claim.SUBSCRIBER_FULLNAME == id[0] {
            getdata.Claim = append(getdata.Claim, claim)
        }

    }
    result, _ := json.Marshal(getdata)
    return shim.Success(result)
        //return shim.Error("Please Provide a valid Patient Id")
}

// ============================================================
// getDataByPayersId -Claim details by status
// ============================================================

func(t *Healthcare) getDataByStatus(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

    type getData struct {
        Claim[] Claim `json:"claim"`
    }

    var getdata getData
    resultsIterator, err := stub.GetStateByRange("C0", "C9999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }

    defer resultsIterator.Close()
    for resultsIterator.HasNext() {
        aKeyValue, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }

        queryValAsByte := aKeyValue.Value

        var claim Claim
        json.Unmarshal(queryValAsByte, & claim)
        if claim.STATUS == id[0] {
            getdata.Claim = append(getdata.Claim, claim)
        }

    }

    result, _ := json.Marshal(getdata)
    return shim.Success(result)
        //return shim.Error("Please Provide a valid Patient Id")
}

// ============================================================
// getallclaim -Claim details
// ============================================================

func(t *Healthcare) getallclaim(stub shim.ChaincodeStubInterface) pb.Response {

    type getData struct {
        Claim[] Claim `json:"claim"`
    }

    var getdata getData

    resultsIterator, err := stub.GetStateByRange("C0", "C9999999999999999999999999999")
    if err != nil {
        return shim.Error(err.Error())
    }

    defer resultsIterator.Close()

    for resultsIterator.HasNext() {

        aKeyValue, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }

        queryValAsByte := aKeyValue.Value
        var claim Claim
        json.Unmarshal(queryValAsByte, & claim)
        getdata.Claim = append(getdata.Claim, claim)
    }
    result, _ := json.Marshal(getdata)
    return shim.Success(result)
        //return shim.Error("Please Provide a valid Patient Id")
}

// ============================================================
// getHistoryForClaim -Claim detail history
// ============================================================


func (t *Healthcare) getHistoryForClaim(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ClaimID := args[0]

	fmt.Printf("- start getHistoryForClaim: %s\n", ClaimID)

	resultsIterator, err := stub.GetHistoryForKey(ClaimID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForClaim returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


