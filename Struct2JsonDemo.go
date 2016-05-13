package main

import (
	"encoding/json"
	"database/sql"
	"fmt"
)

type NullInt struct {
	sql.NullInt64
}

func NewNullInt32(i int32, v bool) NullInt {
	return NullInt{sql.NullInt64{Int64: int64(i), Valid: v}}
}

func NewNullInt(i int64, v bool) *NullInt {
	return &NullInt{sql.NullInt64{Int64: int64(i), Valid: v}}
}

func NewValidNullInt(i int64) *NullInt {
	return &NullInt{sql.NullInt64{Int64: int64(i), Valid: true}}
}

func NewValidNullIntType(i int) NullInt {
	return NullInt{sql.NullInt64{Int64: int64(i), Valid: true}}
}

func (i NullInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Int64)
}

type NullFloat struct {
	sql.NullFloat64
}

func NewNullFloat32(i float32) NullFloat {
	return NullFloat{sql.NullFloat64{Float64: float64(i), Valid: i != 0}}
}

func NewNullFloat64(i float64) NullFloat {
	return NullFloat{sql.NullFloat64{Float64: float64(i), Valid: i != 0}}
}

func (n NullFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Float64)
}

type NullBool struct {
	sql.NullBool
}

func NewNullBool(b bool, v bool) *NullBool {
	return &NullBool{sql.NullBool{Bool: b, Valid: v}}
}

func NewValidNullBool(b bool) *NullBool {
	return &NullBool{sql.NullBool{Bool: b, Valid: true}}
}

func (i *NullBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Bool)
}

type NullString struct {
	sql.NullString
}

func NewNullString(s string, v bool) *NullString {
	return &NullString{sql.NullString{String: s, Valid: v}}
}

func NewValidNullString(s string) *NullString {
	return &NullString{sql.NullString{String: s, Valid: true}}
}

func (i *NullString) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String)
}

type UserModel struct {
	ID           *NullInt `json:"ID,omitempty"`
	UUID         *NullString `json:"UUID,omitempty"`
	TusoID       *NullString `json:"TusoID,omitempty"`
	Email        *NullString `json:"Email,omitempty"`
	MobileNumber *NullString `json:"MobileNumber,omitempty"`
	Password     *NullString `json:"Password,omitempty"`
	Salt         *NullString `json:"Salt,omitempty"`
	Token        *NullString `json:"Token,omitempty"`
	DeviceToken  *NullString `json:"DeviceToken,omitempty"`
	Nickname     *NullString `json:"Nickname,omitempty"`
	RealName     *NullString `json:"RealName,omitempty"`
	Gender       *NullString `json:"Gender,omitempty"`
	Birthday     *NullString `json:"Birthday,omitempty"`
	Location     *NullString `json:"Location,omitempty"`
	Secrets      *NullString `json:"Secrets,omitempty"`
	AvatarURL    *NullString `json:"AvatarURL,omitempty"`
	IsGoodMan    *NullBool `json:"IsGoodMan,omitempty" `
	Model1       *Model1 `json:"Model1,omitempty"`
	Model2       *Model2 `json:"Model2,omitempty"`
	Model3       *Model3 `json:"Model3,omitempty"`
}

type Model1 struct {
	ID   *NullInt `json:"ID,omitempty"`
	UUID *NullString `json:"UUID,omitempty"`
	Name *NullString `json:"Name,omitempty"`
}

type Model2 struct {
	ID   *NullInt `json:"ID,omitempty"`
	UUID *NullString `json:"UUID,omitempty"`
	Name *NullString `json:"Name,omitempty"`
}

type Model3 struct {
	ID   *NullInt `json:"ID,omitempty"`
	UUID *NullString `json:"UUID,omitempty"`
	Name *NullString `json:"Name,omitempty"`
}

func main() {

	u := &UserModel{
		ID:NewValidNullInt(120),
		UUID:NewValidNullString("This is a UUID"),
		TusoID:NewValidNullString("This is a TusoID"),
		Email:NewValidNullString("This is a Email"),
		MobileNumber:NewValidNullString("This is a MobileNumber"),
		Password:NewValidNullString("This is a Password"),
		Salt:NewValidNullString("This is a Salt"),
		Token:NewValidNullString("This is a Token"),
		DeviceToken:NewValidNullString("This is a DeviceToken"),
		Nickname:NewValidNullString("This is a Nickname"),
		RealName:NewValidNullString("This is a RealName"),
		Gender:NewValidNullString("This is a Gender"),
		Birthday:NewValidNullString("This is a Birthday"),
		Location:NewValidNullString("This is a Location"),
		Secrets:NewValidNullString("This is a Secrets"),
		AvatarURL:NewValidNullString("This is a AvatarURL"),
	}

	u1 := UserModel{}
	u1.ID = NewValidNullInt(0)
	u1.UUID = u.UUID
	u1.TusoID = u.TusoID
	u1.Email = NewValidNullString("huxu@123.com")

	u1.IsGoodMan = NewValidNullBool(false)

	u1.Model1 = &Model1{ID:NewValidNullInt(12)}

	result, _ := json.Marshal(&u1)

	fmt.Println(string(result))

	//u2 :=&UserModel{}

	//json.Unmarshal(result,u2)

	//fmt.Println(u2.ID,u2.UUID)


}
