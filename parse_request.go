package main

import (
	"encoding/json"
	"fmt"
)

// Structs for nested data
type MemberCode struct {
	MemberID string `json:"member_id"`
}

type Investor struct {
	ClientCode string `json:"client_code"`
}

type Identifier struct {
	IdentifierType   string `json:"identifier_type"`
	IdentifierNumber string `json:"identifier_number"`
	IssueDate        string `json:"issue_date"`
	ExpiryDate       string `json:"expiry_date"`
	FileName         string `json:"file_name"`
	FileSize         int    `json:"file_size"`
	FileBlob         string `json:"file_blob"`
	AdditionalInfo   string `json:"additional_info"`
}

type Person struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	DOB        string `json:"dob"`
	Gender     string `json:"gender"`
}

type Contact struct {
	ContactNumber      string `json:"contact_number"`
	CountryCode        string `json:"country_code"`
	WhoseContactNumber string `json:"whose_contact_number"`
	EmailAddress       string `json:"email_address"`
	WhoseEmailAddress  string `json:"whose_email_address"`
	ContactType        string `json:"contact_type"`
}

type Holder struct {
	HolderRank        string        `json:"holder_rank"`
	OccCode           string        `json:"occ_code"`
	AuthMode          string        `json:"auth_mode"`
	IsPanExempt       bool          `json:"is_pan_exempt"`
	PanExemptCategory string        `json:"pan_exempt_category"`
	Identifier        []Identifier  `json:"identifier"`
	KycType           string        `json:"kyc_type"`
	CkycNumber        string        `json:"ckyc_number"`
	Person            Person        `json:"person"`
	Contact           []Contact     `json:"contact"`
	Nomination        []interface{} `json:"nomination"`
}

type CommAddr struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	AddressLine3 string `json:"address_line_3"`
	PostalCode   string `json:"postalcode"`
}

type Depository struct {
	DepositoryCode string      `json:"depository_code"`
	DpID           string      `json:"dp_id"`
	ClientID       string      `json:"client_id"`
	CmbpID         string      `json:"cmbp_id"`
	BankAccount    string      `json:"bank_account"`
	AccountOwner   string      `json:"account_owner"`
	Identifier     interface{} `json:"identifier"`
	IsDefault      bool        `json:"is_default"`
}

type BankAccount struct {
	IfscCode     string      `json:"ifsc_code"`
	BankAccNum   string      `json:"bank_acc_num"`
	BankAccType  string      `json:"bank_acc_type"`
	AccountOwner string      `json:"account_owner"`
	Identifier   interface{} `json:"identifier"`
}

type TaxResidency struct {
	Country   string `json:"country"`
	TaxIDNo   string `json:"tax_id_no"`
	TaxIDType string `json:"tax_id_type"`
}

type FATCA struct {
	HolderRank             string         `json:"HolderRank"`
	PlaceOfBirth           string         `json:"place_of_birth"`
	CountryOfBirth         string         `json:"country_of_birth"`
	ClientName             string         `json:"client_name"`
	InvestorType           string         `json:"investor_type"`
	DOB                    string         `json:"dob"`
	FatherName             string         `json:"father_name"`
	SpouseName             string         `json:"spouse_name"`
	AddressType            string         `json:"address_type"`
	OccCode                string         `json:"occ_code"`
	OccType                string         `json:"occ_type"`
	TaxStatus              string         `json:"tax_status"`
	ExemptionCode          string         `json:"exemption_code"`
	Identifier             Identifier     `json:"identifier"`
	CorporateServiceSector string         `json:"corporate_service_sector"`
	WealthSource           string         `json:"wealth_source"`
	IncomeSlab             string         `json:"income_slab"`
	NetWorth               int            `json:"net_worth"`
	DateOfNetWorth         string         `json:"date_of_net_worth"`
	PoliticallyExposed     string         `json:"politically_exposed"`
	IsSelfDeclared         bool           `json:"is_self_declared"`
	DataSource             string         `json:"data_source"`
	TaxResidency           []TaxResidency `json:"tax_residency"`
	FFIDrnfe               string         `json:"ffi_drnfe"`
	IsGIINAvail            string         `json:"is_giin_avail"`
	GIINNo                 string         `json:"giin_no"`
	SprName                string         `json:"spr_name"`
	NfeCategory            string         `json:"nfe_category"`
	NfeSubCategory         string         `json:"nfe_sub_category"`
	NatureOfBusiness       string         `json:"nature_of_business"`
	NatureOfRelation       string         `json:"nature_of_relation"`
}

type AOFIdentifier struct {
	IdentifierType   string `json:"identifier_type"`
	IdentifierNumber string `json:"identifier_number"`
	IssueDate        string `json:"issue_date"`
	ExpiryDate       string `json:"expiry_date"`
	FileName         string `json:"file_name"`
	FileSize         int    `json:"file_size"`
	FileBlob         string `json:"file_blob"`
	AdditionalInfo   string `json:"additional_info"`
}

type AOF struct {
	IsAOFSubmitted bool          `json:"is_aof_submitted"`
	Identifier     AOFIdentifier `json:"identifier"`
}

type Data struct {
	MemberCode         MemberCode    `json:"member_code"`
	Investor           Investor      `json:"investor"`
	IsMultiUCC         bool          `json:"is_multi_ucc"`
	ParentClientCode   string        `json:"parent_client_code"`
	HoldingNature      string        `json:"holding_nature"`
	TaxStatus          string        `json:"tax_status"`
	TaxCode            string        `json:"tax_code"`
	RdmpIDCWPayMode    string        `json:"rdmp_idcw_pay_mode"`
	IsClientPhysical   bool          `json:"is_client_physical"`
	IsClientDemat      bool          `json:"is_client_demat"`
	IsNominationOpted  bool          `json:"is_nomination_opted"`
	NominationAuthMode string        `json:"nomination_auth_mode"`
	CommMode           string        `json:"comm_mode"`
	Onboarding         string        `json:"onboarding"`
	Holder             []Holder      `json:"holder"`
	CommAddr           CommAddr      `json:"comm_addr"`
	ForeignAddr        interface{}   `json:"foreign_addr"`
	Depository         []Depository  `json:"depository"`
	BankAccount        []BankAccount `json:"bank_account"`
	FATCA              []FATCA       `json:"fatca"`
	UBO                interface{}   `json:"ubo"`
	NPO                interface{}   `json:"npo"`
	PMSClient          string        `json:"pms_client"`
	PMSCode            string        `json:"pms_code"`
	AOF                AOF           `json:"aof"`
}

// Root struct for main object
type Root struct {
	Data Data `json:"data"`
}

func main() {
	// Example JSON
	jsonData := `{
		"data": {
			"member_code": {
				"member_id": "0103"
			},
			"investor": {
				"client_code": ""
			},
			"is_multi_ucc": true,
			"parent_client_code": "8509000093",
			"holding_nature": "AS",
			"tax_status": "INDIVIDUAL",
			"tax_code": "01",
			"rdmp_idcw_pay_mode": "02",
			"is_client_physical": true,
			"is_client_demat": true,
			"is_nomination_opted": false,
			"nomination_auth_mode": "",
			"comm_mode": "P",
			"onboarding": "P",
			"holder": [
				{
					"holder_rank": "1",
					"occ_code": "99",
					"auth_mode": "B",
					"is_pan_exempt": false,
					"pan_exempt_category": "",
					"identifier": [
						{
							"identifier_type": "pan",
							"identifier_number": "FHPAN2461F",
							"issue_date": "",
							"expiry_date": "",
							"file_name": "",
							"file_size": 0,
							"file_blob": "",
							"additional_info": ""
						}
					],
					"kyc_type": "K",
					"ckyc_number": "",
					"person": {
						"first_name": "RAM",
						"middle_name": "SHYAM",
						"last_name": "SHARMA",
						"dob": "01-Jan-1995",
						"gender": ""
					},
					"contact": [
						{
							"contact_number": "9930262443",
							"country_code": "91",
							"whose_contact_number": "SE",
							"email_address": "11vikas.mehra@icicisecurities.com",
							"whose_email_address": "SP",
							"contact_type": "OT"
						}
					],
					"nomination": []
				},
				{
					"holder_rank": "2",
					"occ_code": "99",
					"auth_mode": "B",
					"is_pan_exempt": false,
					"pan_exempt_category": "",
					"identifier": [],
					"kyc_type": "K",
					"ckyc_number": "",
					"person": {
						"first_name": "RAJESH",
						"middle_name": "",
						"last_name": "KUMAR",
						"dob": "01-Jan-2001",
						"gender": "M"
					},
					"contact": [
						{
							"contact_number": "",
							"country_code": "",
							"whose_contact_number": "",
							"email_address": "jyoti.kocharekar@icicidirect.com",
							"whose_email_address": "SE",
							"contact_type": "OT"
						}
					],
					"nomination": []
				},
				{
					"holder_rank": "3",
					"occ_code": "99",
					"auth_mode": "B",
					"is_pan_exempt": false,
					"pan_exempt_category": "",
					"identifier": [
						{
							"identifier_type": "pan",
							"identifier_number": "FHPAN2461K",
							"issue_date": "",
							"expiry_date": "",
							"file_name": "",
							"file_size": 0,
							"file_blob": "",
							"additional_info": ""
						}
					],
					"kyc_type": "K",
					"ckyc_number": "",
					"person": {
						"first_name": "RAMCHANDER",
						"middle_name": "",
						"last_name": "SAHU",
						"dob": "12-Sep-1991",
						"gender": "M"
					},
					"contact": [
						{
							"contact_number": "9989221121",
							"country_code": "91",
							"whose_contact_number": "",
							"email_address": "bhagyaraj@email.com",
							"whose_email_address": "",
							"contact_type": "OF"
						}
					],
					"nomination": []
				}
			],
			"comm_addr": {
				"address_line_1": "HNO 9-9 RAM NAGAR",
				"address_line_2": "KRISHNA COLONY",
				"address_line_3": "HYDERABAD TELANGANA ",
				"postalcode": "410210"
			},
			"foreign_addr": null,
			"depository": [
				{
					"depository_code": "NSDL",
					"dp_id": "IN303028",
					"client_id": "50024607",
					"cmbp_id": "IN123456",
					"bank_account": "001705001404",
					"account_owner": "SELF",
					"identifier": null,
					"is_default": false
				},
				{
					"depository_code": "NSDL",
					"dp_id": "IN303028",
					"client_id": "50028495",
					"cmbp_id": "IN123456",
					"bank_account": "001705001404",
					"account_owner": "SELF",
					"identifier": null,
					"is_default": false
				}
			],
			"bank_account": [
				{
					"ifsc_code": "ICIC0000001",
					"bank_acc_num": "001705001404",
					"bank_acc_type": "SB",
					"account_owner": "SELF",
					"identifier": null
				}
			],
			"fatca": [
				{
					"HolderRank": "1",
					"place_of_birth": "LONDON",
					"country_of_birth": "INDIA",
					"client_name": "RAM SHYAM SHARMA",
					"investor_type": "Individual",
					"dob": "01-Jan-1995",
					"father_name": "SHYAM  SHARMA",
					"spouse_name": "",
					"address_type": "5",
					"occ_code": "08",
					"occ_type": "O",
					"tax_status": "INDIVIDUAL",
					"exemption_code": "",
					"identifier": {
						"identifier_type": "pan",
						"identifier_number": "FHPAN2461F",
						"issue_date": "",
						"expiry_date": "",
						"file_name": "",
						"file_size": 0,
						"file_blob": "",
						"additional_info": ""
					},
					"corporate_service_sector": "",
					"wealth_source": "1",
					"income_slab": "32",
					"net_worth": 0,
					"date_of_net_worth": "",
					"politically_exposed": "N",
					"is_self_declared": true,
					"data_source": "E",
					"tax_residency": [
						{
							"country": "UK",
							"tax_id_no": "9879879879",
							"tax_id_type": "C"
						}
					],
					"ffi_drnfe": "",
					"is_giin_avail": "",
					"giin_no": "",
					"spr_name": "",
					"nfe_category": "",
					"nfe_sub_category": "",
					"nature_of_business": "",
					"nature_of_relation": ""
				},
				{
					"HolderRank": "2",
					"place_of_birth": "LONDON",
					"country_of_birth": "INDIA",
					"client_name": "RAJESH  KUMAR",
					"investor_type": "Individual",
					"dob": "01-Jan-2001",
					"father_name": "SHYAM  SHARMA",
					"spouse_name": "",
					"address_type": "5",
					"occ_code": "08",
					"occ_type": "O",
					"tax_status": "INDIVIDUAL",
					"exemption_code": "",
					"identifier": {
						"identifier_type": "pan",
						"identifier_number": "FHPAN2461A",
						"issue_date": "",
						"expiry_date": "",
						"file_name": "",
						"file_size": 0,
						"file_blob": "",
						"additional_info": ""
					},
					"corporate_service_sector": "",
					"wealth_source": "1",
					"income_slab": "",
					"net_worth": 0,
					"date_of_net_worth": "",
					"politically_exposed": "N",
					"is_self_declared": true,
					"data_source": "E",
					"tax_residency": [],
					"ffi_drnfe": "",
					"is_giin_avail": "",
					"giin_no": "",
					"spr_name": "",
					"nfe_category": "",
					"nfe_sub_category": "",
					"nature_of_business": "",
					"nature_of_relation": ""
				},
				{
					"HolderRank": "3",
					"place_of_birth": "LONDON",
					"country_of_birth": "INDIA",
					"client_name": "RAMCHANDER  SAHU",
					"investor_type": "Individual",
					"dob": "12-Sep-1991",
					"father_name": "SHYAM  SHARMA",
					"spouse_name": "",
					"address_type": "5",
					"occ_code": "08",
					"occ_type": "O",
					"tax_status": "INDIVIDUAL",
					"exemption_code": "",
					"identifier": {
						"identifier_type": "pan",
						"identifier_number": "FHPAN2461K",
						"issue_date": "",
						"expiry_date": "",
						"file_name": "",
						"file_size": 0,
						"file_blob": "",
						"additional_info": ""
					},
					"corporate_service_sector": "",
					"wealth_source": "1",
					"income_slab": "33",
					"net_worth": 0,
					"date_of_net_worth": "",
					"politically_exposed": "N",
					"is_self_declared": true,
					"data_source": "E",
					"tax_residency": [
						{
							"country": "UK",
							"tax_id_no": "1234567891",
							"tax_id_type": "C"
						}
					],
					"ffi_drnfe": "",
					"is_giin_avail": "",
					"giin_no": "",
					"spr_name": "",
					"nfe_category": "",
					"nfe_sub_category": "",
					"nature_of_business": "",
					"nature_of_relation": ""
				}
			],
			"ubo": null,
			"npo": null,
			"pms_client": "N",
			"pms_code": "",
			"aof": {
				"is_aof_submitted": false,
				"identifier": {
					"identifier_type": null,
					"identifier_number": null,
					"issue_date": null,
					"expiry_date": null,
					"file_name": null,
					"file_size": 0,
					"file_blob": null,
					"additional_info": null
				}
			}
		}
	}`

	// Unmarshal the JSON into the Root struct
	var root Root
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Accessing the data
	fmt.Printf("member_code: %s\n", root.Data.MemberCode.MemberID)
	fmt.Printf("client_code: %s\n", root.Data.Investor.ClientCode)
	fmt.Printf("is_multi_ucc: %t\n", root.Data.IsMultiUCC)
	fmt.Printf("parent_client_code: %s\n", root.Data.ParentClientCode)
	fmt.Printf("primary_holder_name: %s\n", root.Data.Holder[0].Person.FirstName)
}
