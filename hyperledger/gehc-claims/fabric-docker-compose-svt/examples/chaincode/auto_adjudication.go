package main

import (
    "fmt"
    "time"
	"bytes"
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
    DEDUCTIBLE string `json:"deductible"`
    AMOUNT string `json:"amount"`
    COPAY string `json:"copay"`
    ELIGIBILITY_CODE string `json:"eligibitity_code"`
    PROCEDURE_CODE string `json:"procedure_code"`
    STATUS string `json:"status"`
    OWNER string `json:"owner"`
	COMMENTS string `json:"comments"`
	SETTLEMENT_AMOUNT string `json:"settlement_amount"`
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
    AMOUNT string `json:"amount"`
    PROCEDURE_CODE string `json:"procedure_code"`
    DOS string `json:"dos"`


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


func(t * Healthcare) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}


// Invoke - Our entry point for Invocations
// ========================================


func(t * Healthcare) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println(" ")
    fmt.Println("starting invoke, for - " + function)

    if

    function == "init" {
        return t.Init(stub)
    } else if
	
	function == "schedule_appointment" {//schedule_appointment for Insurance holder
        return t.schedule_appointment(stub, args)
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

    function == "getDataByPatientName" {// get claim deatils by subcriber name
        return t.getDataByPatientName(stub, args)
    } else if

    function == "getDataByPatientId" {//get claim deatils by subcriber id
        return t.getDataByPatientId(stub, args)
    } else if

    function == "getDataByStatus" {//get claim deatils by status
        return t.getDataByStatus(stub, args)
    } else if

    function == "getDataByPayersname" {//get claim deatils by payer name 
        return t.getDataByPayersname(stub, args)
    } else if

    function == "getDataByPayersId" {//get claim deatils by payer id
        return t.getDataByPayersId(stub, args)
    } else if

    function == "getallclaim" {//get all claim
        return t.getallclaim(stub)
    } else if

    function == "getDataByOwner" {//get claim deatils by owner
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

// ============================================================
// schedule_appointment - create a appointment for subcriber
// ============================================================
func(t * Healthcare) schedule_appointment(stub shim.ChaincodeStubInterface, args[] string) pb.Response {
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
    deductible:=args[10]//500
    copay :=args[11]//45
    eligibility_code :=args[12]//1
    status:= "Scheduled"
    owner:=provider_fullname
	comments:=args[13]

    
    fmt.Println("starting create_claim")

    
    var claim = & Claim {
	"claim",claim_id ,subscriber_fullname ,subscriber_gender,subscriberid,provider_fullname,providerid,payername,payerid,servicetype_code,dos,"",deductible,"",copay,eligibility_code,"",status,owner,comments,""}
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

func(t * Healthcare) Check_in(stub shim.ChaincodeStubInterface, args[] string) pb.Response {
    var err error
	if len(args) != 3 {
        return shim.Error("Incorrect number of arguments. Expecting 3")
    }
	claim_id :=args[0]//c0
    deductible :=args[1]
	datenow:=args[2]
	
	// ==== Check if claim already exists ====
	var claim Claim
	
    claimAsBytes, err := stub.GetState(claim_id)
    json.Unmarshal(claimAsBytes, & claim)
	
	if err != nil {
		return shim.Error("Failed to get claim: " + err.Error())
	} else if claimAsBytes != nil {
		if claim.STATUS=="CheckIn"{
		return shim.Error("This claim already exists: " + claim_id)
		}
	}
	
	
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
func(t * Healthcare) approve_cliam(stub shim.ChaincodeStubInterface, args[] string) pb.Response {
	
	if len(args) != 3 {
        return shim.Error("Incorrect number of arguments. Expecting 3")
    }
	amount :=args[2]
    Procedure_code :=args[1]
	claim_id:=args[0]

    var claim Claim

	// ==== Check if claim already exists ====
	
    claimAsBytes, err := stub.GetState(claim_id)
    json.Unmarshal(claimAsBytes, & claim)
	
	if err != nil {
		return shim.Error("Failed to get claim: " + err.Error())
	} else if claimAsBytes != nil {
		if claim.STATUS=="Approved"||claim.STATUS=="Pendding"||claim.STATUS=="Decline"{
		return shim.Error("This claim already exists: " + claim_id)
		}
	}
	
    Damount := claim.DEDUCTIBLE
    New_Owner_ID := claim.PAYERNAME
	
    amt, _ := strconv.Atoi(amount)
    damt, _ := strconv.Atoi(Damount)

	//Check Procedure_code
    if (Procedure_code == "1134" || Procedure_code == "1133" || Procedure_code == "1122" || Procedure_code == "1167" || Procedure_code == "1187" || Procedure_code == "1188" || Procedure_code == "1199" || Procedure_code == "1167") {

	//Check amount is less than 1000$
	if (amt < 1000) {
	
	var remitamount int
	var newdeductiable int
	
	//check Deductiable is greter than amount and calculate new deductible and remit amount
	if damt>amt{
	remitamount= 0
	newdeductiable=damt-amt
	}
	
	//check Deductiable is less than amount and calculate new deductible and remit amount
	if damt<amt{
	remitamount1:=amt-damt
	discount:=0.75
	remitamount=int(float64(remitamount1)*discount)
	newdeductiable=0
	}
	
	remitamount1 := strconv.Itoa(remitamount)
	newdeductiable1 := strconv.Itoa(newdeductiable)
	
	
	claim.STATUS = "Approved"
	claim.DEDUCTIBLE = newdeductiable1
	claim.AMOUNT = amount
	claim.OWNER = New_Owner_ID
	claim.PROCEDURE_CODE=Procedure_code
	claim.COMMENTS="You receive refound within next 3-4 working days"
	claim.SETTLEMENT_AMOUNT=remitamount1
	
	jsonAsBytes,_ := json.Marshal(claim)
	err = stub.PutState(args[0], jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	t := time.Now()
	
	// crate remit receipt
	claimid := "R" + args[0]
	sender_acc_no := "s" + t.Format("20060102150405")
	receiver_acc_no := "r" + t.Format("20060102150405")
	checkno := "c" + t.Format("20060102150405")
	payeeid := claim.PROVIDERID
	subcriberid := claim.SUBSCRIBERID
	procedure_code := Procedure_code
	dos := claim.DOS

	var remit = & REMIT {
	"remit", claimid, sender_acc_no, receiver_acc_no, checkno, New_Owner_ID, payeeid, subcriberid, remitamount1, procedure_code, dos}
	remitAsBytes,_:= json.Marshal(remit)
	err = stub.PutState(claimid, remitAsBytes)
	
	if err != nil {
		fmt.Println("Could not store claim")
		return shim.Error(err.Error())
	}
	}else{
	claim.STATUS = "Pendding"
	claim.DEDUCTIBLE = strconv.Itoa(amt)
	claim.OWNER = New_Owner_ID
	claim.PROCEDURE_CODE=Procedure_code
	claim.COMMENTS="Amount grater than 1000$"
	
	jsonAsBytes, _ := json.Marshal(claim)

	err = stub.PutState(args[0], jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	}
    }else {

	claim.STATUS = "Decline"
	claim.DEDUCTIBLE = strconv.Itoa(amt)
	claim.OWNER = New_Owner_ID
	claim.PROCEDURE_CODE=Procedure_code
	claim.COMMENTS="Invalid service_code/procedure_code"
	jsonAsBytes, _ := json.Marshal(claim)

	err = stub.PutState(args[0], jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
    }

    fmt.Println("- end init_hospital")
    return shim.Success(nil)

}

// ============================================================
// getRemitData -remit details by claim id
// ============================================================
func(t * Healthcare) getRemitData(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

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
func(t * Healthcare) getDataByClaimId(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

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
func(t * Healthcare) getDataByOwner(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

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

        if claim.OWNER == id[0] {
            getdata.Claim = append(getdata.Claim, claim)
        }

    }
    result, _ := json.Marshal(getdata)
    return shim.Success(result)

}


// ============================================================
// getDataByIDstatus -Claim details by id and status
// ============================================================

func(t * Healthcare) getDataByIDstatus(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

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
func(t * Healthcare) getDataByPatientId(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

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
// ============================================================
// getDataByPayersname -Claim details by subcriber name
// ============================================================

func(t * Healthcare) getDataByPayersname(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

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
func(t * Healthcare) getDataByPayersId(stub shim.ChaincodeStubInterface, id[] string) pb.Response {
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
func(t * Healthcare) getDataByPatientName(stub shim.ChaincodeStubInterface, id[] string) pb.Response {
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

func(t * Healthcare) getDataByStatus(stub shim.ChaincodeStubInterface, id[] string) pb.Response {

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

func(t * Healthcare) getallclaim(stub shim.ChaincodeStubInterface) pb.Response {

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

