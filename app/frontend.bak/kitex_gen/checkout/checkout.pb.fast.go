// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package checkout

import (
	fmt "fmt"
	cart "github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	payment "github.com/baiyutang/gomall/app/frontend/kitex_gen/payment"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Address) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Address[number], err)
}

func (x *Address) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StreetAddress, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.City, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.State, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Country, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.ZipCode, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckoutReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckoutReq[number], err)
}

func (x *CheckoutReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CheckoutReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Firstname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckoutReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Lastname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckoutReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckoutReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v Address
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Address = &v
	return offset, nil
}

func (x *CheckoutReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	var v payment.CreditCardInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CreditCard = &v
	return offset, nil
}

func (x *OrderItem) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OrderItem[number], err)
}

func (x *OrderItem) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v cart.CartItem
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Item = &v
	return offset, nil
}

func (x *OrderItem) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Cost, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *OrderResult) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OrderResult[number], err)
}

func (x *OrderResult) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderResult) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ShippingTrackingId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderResult) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ShippingCost, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *OrderResult) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v Address
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.ShippingAddress = &v
	return offset, nil
}

func (x *OrderResult) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v OrderItem
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Items = append(x.Items, &v)
	return offset, nil
}

func (x *CheckoutRes) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckoutRes[number], err)
}

func (x *CheckoutRes) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v OrderResult
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Order = &v
	return offset, nil
}

func (x *Address) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *Address) fastWriteField1(buf []byte) (offset int) {
	if x.StreetAddress == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetStreetAddress())
	return offset
}

func (x *Address) fastWriteField2(buf []byte) (offset int) {
	if x.City == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCity())
	return offset
}

func (x *Address) fastWriteField3(buf []byte) (offset int) {
	if x.State == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetState())
	return offset
}

func (x *Address) fastWriteField4(buf []byte) (offset int) {
	if x.Country == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetCountry())
	return offset
}

func (x *Address) fastWriteField5(buf []byte) (offset int) {
	if x.ZipCode == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetZipCode())
	return offset
}

func (x *CheckoutReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *CheckoutReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CheckoutReq) fastWriteField2(buf []byte) (offset int) {
	if x.Firstname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFirstname())
	return offset
}

func (x *CheckoutReq) fastWriteField3(buf []byte) (offset int) {
	if x.Lastname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetLastname())
	return offset
}

func (x *CheckoutReq) fastWriteField4(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetEmail())
	return offset
}

func (x *CheckoutReq) fastWriteField5(buf []byte) (offset int) {
	if x.Address == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 5, x.GetAddress())
	return offset
}

func (x *CheckoutReq) fastWriteField6(buf []byte) (offset int) {
	if x.CreditCard == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 6, x.GetCreditCard())
	return offset
}

func (x *OrderItem) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *OrderItem) fastWriteField1(buf []byte) (offset int) {
	if x.Item == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetItem())
	return offset
}

func (x *OrderItem) fastWriteField2(buf []byte) (offset int) {
	if x.Cost == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 2, x.GetCost())
	return offset
}

func (x *OrderResult) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *OrderResult) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *OrderResult) fastWriteField2(buf []byte) (offset int) {
	if x.ShippingTrackingId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetShippingTrackingId())
	return offset
}

func (x *OrderResult) fastWriteField3(buf []byte) (offset int) {
	if x.ShippingCost == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 3, x.GetShippingCost())
	return offset
}

func (x *OrderResult) fastWriteField4(buf []byte) (offset int) {
	if x.ShippingAddress == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetShippingAddress())
	return offset
}

func (x *OrderResult) fastWriteField5(buf []byte) (offset int) {
	if x.Items == nil {
		return offset
	}
	for i := range x.GetItems() {
		offset += fastpb.WriteMessage(buf[offset:], 5, x.GetItems()[i])
	}
	return offset
}

func (x *CheckoutRes) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CheckoutRes) fastWriteField1(buf []byte) (offset int) {
	if x.Order == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetOrder())
	return offset
}

func (x *Address) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *Address) sizeField1() (n int) {
	if x.StreetAddress == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetStreetAddress())
	return n
}

func (x *Address) sizeField2() (n int) {
	if x.City == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetCity())
	return n
}

func (x *Address) sizeField3() (n int) {
	if x.State == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetState())
	return n
}

func (x *Address) sizeField4() (n int) {
	if x.Country == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetCountry())
	return n
}

func (x *Address) sizeField5() (n int) {
	if x.ZipCode == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetZipCode())
	return n
}

func (x *CheckoutReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *CheckoutReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *CheckoutReq) sizeField2() (n int) {
	if x.Firstname == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFirstname())
	return n
}

func (x *CheckoutReq) sizeField3() (n int) {
	if x.Lastname == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetLastname())
	return n
}

func (x *CheckoutReq) sizeField4() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetEmail())
	return n
}

func (x *CheckoutReq) sizeField5() (n int) {
	if x.Address == nil {
		return n
	}
	n += fastpb.SizeMessage(5, x.GetAddress())
	return n
}

func (x *CheckoutReq) sizeField6() (n int) {
	if x.CreditCard == nil {
		return n
	}
	n += fastpb.SizeMessage(6, x.GetCreditCard())
	return n
}

func (x *OrderItem) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *OrderItem) sizeField1() (n int) {
	if x.Item == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetItem())
	return n
}

func (x *OrderItem) sizeField2() (n int) {
	if x.Cost == 0 {
		return n
	}
	n += fastpb.SizeFloat(2, x.GetCost())
	return n
}

func (x *OrderResult) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *OrderResult) sizeField1() (n int) {
	if x.OrderId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetOrderId())
	return n
}

func (x *OrderResult) sizeField2() (n int) {
	if x.ShippingTrackingId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetShippingTrackingId())
	return n
}

func (x *OrderResult) sizeField3() (n int) {
	if x.ShippingCost == 0 {
		return n
	}
	n += fastpb.SizeFloat(3, x.GetShippingCost())
	return n
}

func (x *OrderResult) sizeField4() (n int) {
	if x.ShippingAddress == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetShippingAddress())
	return n
}

func (x *OrderResult) sizeField5() (n int) {
	if x.Items == nil {
		return n
	}
	for i := range x.GetItems() {
		n += fastpb.SizeMessage(5, x.GetItems()[i])
	}
	return n
}

func (x *CheckoutRes) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CheckoutRes) sizeField1() (n int) {
	if x.Order == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetOrder())
	return n
}

var fieldIDToName_Address = map[int32]string{
	1: "StreetAddress",
	2: "City",
	3: "State",
	4: "Country",
	5: "ZipCode",
}

var fieldIDToName_CheckoutReq = map[int32]string{
	1: "UserId",
	2: "Firstname",
	3: "Lastname",
	4: "Email",
	5: "Address",
	6: "CreditCard",
}

var fieldIDToName_OrderItem = map[int32]string{
	1: "Item",
	2: "Cost",
}

var fieldIDToName_OrderResult = map[int32]string{
	1: "OrderId",
	2: "ShippingTrackingId",
	3: "ShippingCost",
	4: "ShippingAddress",
	5: "Items",
}

var fieldIDToName_CheckoutRes = map[int32]string{
	1: "Order",
}

var _ = payment.File_payment_proto
var _ = cart.File_cart_proto