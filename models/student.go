package models

import (
	"crudgolang/db"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Student
type Student struct {
	ID          string `json="id"`
	Name        string `json="name"`
	Age         int    `json="age"`
	Grade       int    `json="grade"`
	AccountCode string `json="account_code"`
}

// // Data
// type Data struct {
// 	ID          string `json:"id"`
// 	Name        string `json:"name"`
// 	AccountCode string `json:"account_code" db:"account_code"`
// }

// Data
type Data struct {
	ID                    *string `json:"id"`
	Fare                  *string `json:"fare"`
	Extra                 *string `json:"extra"`
	DiscountAmt           *string `json:"discount_amt" db:"discount_amt"`
	PaidAmount            *string `json:"paid_amount" db:"paid_amount"`
	PromoCode             *string `json:"promo_code" db:"promo_code"`
	PaymentToken          *string `json:"payment_token" db:"payment_token"`
	TransactionTime       *string `json:"transaction_time" db:"transaction_time"`
	Identifier            *string `json:"identifier" db:"identifier"`
	PaymentType           *string `json:"payment_type" db:"payment_type"`
	VehicleId             *string `json:"vehicle_id" db:"vehicle_id"`
	VehicleName           *string `json:"vehicle_name" db:"vehicle_name"`
	ServiceType           *string `json:"service_type" db:"service_type"`
	DriverID              *string `json:"driver_id" db:"driver_id"`
	PickUpSuburb          *string `json:"pick_up_suburb" db:"pick_up_suburb"`
	PickUpArea            *string `json:"pick_up_area" db:"pick_up_area"`
	DestinationArea       *string `json:"destination_area" db:"destination_area"`
	DSestinationSuburb    *string `json:"destination_suburb" db:"destination_suburb"`
	PickUpLatitude        *string `json:"pick_up_latitude" db:"pick_up_latitude"`
	PickUpLng             *string `json:"pick_up_lng" db:"pick_up_lng"`
	PaymentProfileID      *string `json:"payment_profile_id" db:"payment_profile_id"`
	State                 *string `json:"state"`
	ReleasedAt            *string `json:"released_at" db:"released_at"`
	CompletedAt           *string `json:"completed_at" db:"completed_at"`
	CreatedAt             *string `json:"created_at" db:"created_at"`
	Updated_at            *string `json:"updated_at" db:"updated_at"`
	CcIdentifier          *string `json:"cc_identifier" db:"cc_identifier"`
	AccountID             *string `json:"account_id" db:"account_id"`
	SapSentAt             *string `json:"sap_sent_at" db:"sap_sent_at"`
	SapState              *string `json:"sap_state" db:"sap_state"`
	MsakuState            *string `json:"msaku_state" db:"msaku_state"`
	CvNumber              *string `json:"cv_number" db:"cv_number"`
	ValidityPeriod        *string `json:"validity_period" db:"validity_period"`
	ItopID                *string `json:"itop_id" db:"itop_id"`
	OrderID               *string `json:"order_id" db:"order_id"`
	PickupAdress          *string `json:"pickup_adress" db:"pickup_adress"`
	PickedUp              *string `json:"picked_up_at" db:"picked_up_at"`
	TripPurpose           *string `json:"trip_purpose" db:"trip_purpose"`
	MsakuTransactionID    *string `json:"msaku_transaction_id" db:"msaku_transaction_id"`
	TripPurposedriverName *string `json:"trip_purposedriver_name" db:"trip_purposedriver_name"`
	ExternalOrderID       *string `json:"external_order_id" db:"external_order_id"`
	RouteImage            *string `json:"route_image" db:"route_image"`
	DepartmentName        *string `json:"department_name" db:"department_name"`
	AccountCode           *string `json:"account_code" db:"account_code"`
	UserName              *string `json:"user_name" db:"user_name"`
	InvoiceNumber         *string `json:"invoice_number,omitempty" db:"invoice_number"`
	PostingDate           *string `json:"posting_date" db:"posting_date"`
	Distance              *string `json:"distance"`
	OtherInformation      *string `json:"other_information" db:"other_information"`
	PickUpLat             *string `json:"pick_up_lat" db:"pick_up_lat"`
	DestinationLat        *string `json:"destination_lat" db:"destination_lat"`
	DestinationLng        *string `json:"destination_lng" db:"destination_lng"`
	MsakuResponse         *string `json:"msaku_response" db:"msaku_response"`
	PickupAddress         *string `json:"pickup_address" db:"pickup_address"`
	DropoffAddress        *string `json:"dropoff_address" db:"dropoff_address"`
	Tips                  *string `json:"tips" db:"tips"`
	DriverName            *string `json:"driver_name" db:"driver_name"`
}

func GetStudent(u *Student, id string) int {
	query := "SELECT * FROM tb_student WHERE id= '" + id + "'"
	err := db.Db.Get(u, query)
	log.Println(query)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound
	}
	log.Println("u: ", u)
	return http.StatusOK
}

func GetData(u *Data, id string) int {
	query := "select id, name from accounts where id= " + id
	err := db.Db.Get(u, query)
	log.Println(query)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound
	}
	log.Println("u: ", u)
	return http.StatusOK
}

func GetAllStudent(c *[]Student, params map[string]string) (uint64, error) {
	query := "SELECT * FROM tb_student"
	var condition string
	// Combine where clause
	clause := false
	for key, value := range params {
		if (key != "orderBy") && (key != "orderType") {
			if clause == false {
				condition += " WHERE"
			} else {
				condition += " AND"
			}
			condition += " tb_student." + key + " = '" + value + "'"
			clause = true
		}
	}
	// Check order by
	var present bool
	var orderBy, orderType string
	if orderBy, present = params["orderBy"]; present == true {
		condition += " ORDER BY tb_student." + orderBy
		if orderType, present = params["orderType"]; present == true {
			condition += " " + orderType
		}
	}
	query += condition

	// Main query
	log.Println(query)
	err := db.Db.Select(c, query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// if pagination == false {
	total := uint64(len(*c))
	// }

	log.Println(total)
	return total, nil
}

// func GetAllData(c *[]Data) error {
// 	query := "select id, name, account_code from accounts"

// 	// Main query
// 	log.Println(query)
// 	err := db.Db.Select(c, query)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	return nil
// }

func GetAllData(c *[]Data) error {
	query := `
	SELECT distinct
		payment_transactions.*,
		departments.name AS department_name,
		accounts.account_code AS account_code,
		users.full_name AS user_name,
		info.invoice_number AS invoice_number,
		info.posting_date AS posting_date,
		info.distance AS distance,
		users.other_information AS other_information
		FROM "payment_transactions"
		LEFT JOIN voucher_profiles vp ON vp.id = payment_transactions.cc_identifier::integer
		LEFT JOIN users ON users.id = vp.user_id
		LEFT JOIN departments ON departments.id = vp.department_id
		LEFT JOIN accounts ON accounts.id = payment_transactions.account_id
		LEFT JOIN payment_transaction_infos info ON info.payment_transaction_id = payment_transactions.id
	WHERE
		"payment_transactions"."state" IN (2, 3) AND
		"payment_transactions"."payment_type" IN ('ecv', 'edc') AND
		("payment_transactions"."completed_at" BETWEEN '2017-07-01 00:00:00.000000' AND '2021-05-28 00:00:00.000000')
	ORDER BY "payment_transactions"."picked_up_at" ASC, "payment_transactions"."completed_at" ASC`

	// Main query
	log.Println(query)
	err := db.Db.Select(c, query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateStudent(params map[string]string) int {
	query := "INSERT INTO tb_student("
	// Get params
	var fields, values string
	i := 0
	for key, value := range params {
		fields += "`" + key + "`"
		values += "'" + value + "'"
		if (len(params) - 1) > i {
			fields += ", "
			values += ", "
		}
		i++
	}
	// Combile params to build query
	query += fields + ") VALUES(" + values + ")"
	log.Println(query)

	tx, err := db.Db.Begin()
	if err != nil {
		log.Println(err)
		return http.StatusBadGateway
	}
	_, err = tx.Exec(query)
	tx.Commit()
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func UpdateStudent(params map[string]string) int {
	query := "UPDATE tb_student SET "
	// Get params
	i := 0
	for key, value := range params {
		if key != "id" {
			query += "`" + key + "`" + " = '" + value + "'"
			if (len(params) - 2) > i {
				query += ", "
			}
			i++
		}
	}
	query += " WHERE id = '" + params["id"] + "'"
	log.Println(query)

	tx, err := db.Db.Begin()
	if err != nil {
		log.Println(err)
		return http.StatusBadGateway
	}
	_, err = tx.Exec(query)
	tx.Commit()
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func DeleteStudent(id string) int {
	query := "DELETE FROM tb_student WHERE id = '" + id + "'"
	tx, err := db.Db.Begin()
	if err != nil {
		log.Println(err)
		return http.StatusBadGateway
	}
	var ret sql.Result
	log.Println(query)
	ret, err = tx.Exec(query)
	row, _ := ret.RowsAffected()
	if row > 0 {
		tx.Commit()
	} else {
		return http.StatusNotFound
	}
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	return http.StatusOK
}
