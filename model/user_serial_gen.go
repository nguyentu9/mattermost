// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *User) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 34 {
		err = msgp.ArrayError{Wanted: 34, Got: zb0001}
		return
	}
	z.Id, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	z.CreateAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "CreateAt")
		return
	}
	z.UpdateAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "UpdateAt")
		return
	}
	z.DeleteAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "DeleteAt")
		return
	}
	z.Username, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Username")
		return
	}
	z.Password, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Password")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "AuthData")
			return
		}
		z.AuthData = nil
	} else {
		if z.AuthData == nil {
			z.AuthData = new(string)
		}
		*z.AuthData, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "AuthData")
			return
		}
	}
	z.AuthService, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "AuthService")
		return
	}
	z.Email, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Email")
		return
	}
	z.EmailVerified, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "EmailVerified")
		return
	}
	z.Nickname, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Nickname")
		return
	}
	z.FirstName, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "FirstName")
		return
	}
	z.LastName, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "LastName")
		return
	}
	z.Position, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Position")
		return
	}
	z.Roles, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	z.AllowMarketing, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "AllowMarketing")
		return
	}
	err = z.Props.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	err = z.NotifyProps.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "NotifyProps")
		return
	}
	z.LastPasswordUpdate, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "LastPasswordUpdate")
		return
	}
	z.LastPictureUpdate, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "LastPictureUpdate")
		return
	}
	z.FailedAttempts, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "FailedAttempts")
		return
	}
	z.Locale, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Locale")
		return
	}
	err = z.Timezone.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Timezone")
		return
	}
	z.MfaActive, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "MfaActive")
		return
	}
	z.MfaSecret, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "MfaSecret")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "RemoteId")
			return
		}
		z.RemoteId = nil
	} else {
		if z.RemoteId == nil {
			z.RemoteId = new(string)
		}
		*z.RemoteId, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "RemoteId")
			return
		}
	}
	z.LastActivityAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "LastActivityAt")
		return
	}
	z.IsBot, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "IsBot")
		return
	}
	z.BotDescription, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "BotDescription")
		return
	}
	z.BotLastIconUpdate, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "BotLastIconUpdate")
		return
	}
	z.TermsOfServiceId, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceId")
		return
	}
	z.TermsOfServiceCreateAt, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceCreateAt")
		return
	}
	z.DisableWelcomeEmail, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "DisableWelcomeEmail")
		return
	}
	z.Bio, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Bio")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *User) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 34
	err = en.Append(0xdc, 0x0, 0x22)
	if err != nil {
		return
	}
	err = en.WriteString(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	err = en.WriteInt64(z.CreateAt)
	if err != nil {
		err = msgp.WrapError(err, "CreateAt")
		return
	}
	err = en.WriteInt64(z.UpdateAt)
	if err != nil {
		err = msgp.WrapError(err, "UpdateAt")
		return
	}
	err = en.WriteInt64(z.DeleteAt)
	if err != nil {
		err = msgp.WrapError(err, "DeleteAt")
		return
	}
	err = en.WriteString(z.Username)
	if err != nil {
		err = msgp.WrapError(err, "Username")
		return
	}
	err = en.WriteString(z.Password)
	if err != nil {
		err = msgp.WrapError(err, "Password")
		return
	}
	if z.AuthData == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteString(*z.AuthData)
		if err != nil {
			err = msgp.WrapError(err, "AuthData")
			return
		}
	}
	err = en.WriteString(z.AuthService)
	if err != nil {
		err = msgp.WrapError(err, "AuthService")
		return
	}
	err = en.WriteString(z.Email)
	if err != nil {
		err = msgp.WrapError(err, "Email")
		return
	}
	err = en.WriteBool(z.EmailVerified)
	if err != nil {
		err = msgp.WrapError(err, "EmailVerified")
		return
	}
	err = en.WriteString(z.Nickname)
	if err != nil {
		err = msgp.WrapError(err, "Nickname")
		return
	}
	err = en.WriteString(z.FirstName)
	if err != nil {
		err = msgp.WrapError(err, "FirstName")
		return
	}
	err = en.WriteString(z.LastName)
	if err != nil {
		err = msgp.WrapError(err, "LastName")
		return
	}
	err = en.WriteString(z.Position)
	if err != nil {
		err = msgp.WrapError(err, "Position")
		return
	}
	err = en.WriteString(z.Roles)
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	err = en.WriteBool(z.AllowMarketing)
	if err != nil {
		err = msgp.WrapError(err, "AllowMarketing")
		return
	}
	err = z.Props.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	err = z.NotifyProps.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "NotifyProps")
		return
	}
	err = en.WriteInt64(z.LastPasswordUpdate)
	if err != nil {
		err = msgp.WrapError(err, "LastPasswordUpdate")
		return
	}
	err = en.WriteInt64(z.LastPictureUpdate)
	if err != nil {
		err = msgp.WrapError(err, "LastPictureUpdate")
		return
	}
	err = en.WriteInt(z.FailedAttempts)
	if err != nil {
		err = msgp.WrapError(err, "FailedAttempts")
		return
	}
	err = en.WriteString(z.Locale)
	if err != nil {
		err = msgp.WrapError(err, "Locale")
		return
	}
	err = z.Timezone.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Timezone")
		return
	}
	err = en.WriteBool(z.MfaActive)
	if err != nil {
		err = msgp.WrapError(err, "MfaActive")
		return
	}
	err = en.WriteString(z.MfaSecret)
	if err != nil {
		err = msgp.WrapError(err, "MfaSecret")
		return
	}
	if z.RemoteId == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteString(*z.RemoteId)
		if err != nil {
			err = msgp.WrapError(err, "RemoteId")
			return
		}
	}
	err = en.WriteInt64(z.LastActivityAt)
	if err != nil {
		err = msgp.WrapError(err, "LastActivityAt")
		return
	}
	err = en.WriteBool(z.IsBot)
	if err != nil {
		err = msgp.WrapError(err, "IsBot")
		return
	}
	err = en.WriteString(z.BotDescription)
	if err != nil {
		err = msgp.WrapError(err, "BotDescription")
		return
	}
	err = en.WriteInt64(z.BotLastIconUpdate)
	if err != nil {
		err = msgp.WrapError(err, "BotLastIconUpdate")
		return
	}
	err = en.WriteString(z.TermsOfServiceId)
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceId")
		return
	}
	err = en.WriteInt64(z.TermsOfServiceCreateAt)
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceCreateAt")
		return
	}
	err = en.WriteBool(z.DisableWelcomeEmail)
	if err != nil {
		err = msgp.WrapError(err, "DisableWelcomeEmail")
		return
	}
	err = en.WriteString(z.Bio)
	if err != nil {
		err = msgp.WrapError(err, "Bio")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *User) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 34
	o = append(o, 0xdc, 0x0, 0x22)
	o = msgp.AppendString(o, z.Id)
	o = msgp.AppendInt64(o, z.CreateAt)
	o = msgp.AppendInt64(o, z.UpdateAt)
	o = msgp.AppendInt64(o, z.DeleteAt)
	o = msgp.AppendString(o, z.Username)
	o = msgp.AppendString(o, z.Password)
	if z.AuthData == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendString(o, *z.AuthData)
	}
	o = msgp.AppendString(o, z.AuthService)
	o = msgp.AppendString(o, z.Email)
	o = msgp.AppendBool(o, z.EmailVerified)
	o = msgp.AppendString(o, z.Nickname)
	o = msgp.AppendString(o, z.FirstName)
	o = msgp.AppendString(o, z.LastName)
	o = msgp.AppendString(o, z.Position)
	o = msgp.AppendString(o, z.Roles)
	o = msgp.AppendBool(o, z.AllowMarketing)
	o, err = z.Props.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	o, err = z.NotifyProps.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "NotifyProps")
		return
	}
	o = msgp.AppendInt64(o, z.LastPasswordUpdate)
	o = msgp.AppendInt64(o, z.LastPictureUpdate)
	o = msgp.AppendInt(o, z.FailedAttempts)
	o = msgp.AppendString(o, z.Locale)
	o, err = z.Timezone.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Timezone")
		return
	}
	o = msgp.AppendBool(o, z.MfaActive)
	o = msgp.AppendString(o, z.MfaSecret)
	if z.RemoteId == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendString(o, *z.RemoteId)
	}
	o = msgp.AppendInt64(o, z.LastActivityAt)
	o = msgp.AppendBool(o, z.IsBot)
	o = msgp.AppendString(o, z.BotDescription)
	o = msgp.AppendInt64(o, z.BotLastIconUpdate)
	o = msgp.AppendString(o, z.TermsOfServiceId)
	o = msgp.AppendInt64(o, z.TermsOfServiceCreateAt)
	o = msgp.AppendBool(o, z.DisableWelcomeEmail)
	o = msgp.AppendString(o, z.Bio)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *User) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 34 {
		err = msgp.ArrayError{Wanted: 34, Got: zb0001}
		return
	}
	z.Id, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	z.CreateAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "CreateAt")
		return
	}
	z.UpdateAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "UpdateAt")
		return
	}
	z.DeleteAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DeleteAt")
		return
	}
	z.Username, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Username")
		return
	}
	z.Password, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Password")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.AuthData = nil
	} else {
		if z.AuthData == nil {
			z.AuthData = new(string)
		}
		*z.AuthData, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "AuthData")
			return
		}
	}
	z.AuthService, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AuthService")
		return
	}
	z.Email, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Email")
		return
	}
	z.EmailVerified, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "EmailVerified")
		return
	}
	z.Nickname, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Nickname")
		return
	}
	z.FirstName, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "FirstName")
		return
	}
	z.LastName, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastName")
		return
	}
	z.Position, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Position")
		return
	}
	z.Roles, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Roles")
		return
	}
	z.AllowMarketing, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AllowMarketing")
		return
	}
	bts, err = z.Props.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Props")
		return
	}
	bts, err = z.NotifyProps.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "NotifyProps")
		return
	}
	z.LastPasswordUpdate, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastPasswordUpdate")
		return
	}
	z.LastPictureUpdate, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastPictureUpdate")
		return
	}
	z.FailedAttempts, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "FailedAttempts")
		return
	}
	z.Locale, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Locale")
		return
	}
	bts, err = z.Timezone.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Timezone")
		return
	}
	z.MfaActive, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MfaActive")
		return
	}
	z.MfaSecret, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MfaSecret")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.RemoteId = nil
	} else {
		if z.RemoteId == nil {
			z.RemoteId = new(string)
		}
		*z.RemoteId, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "RemoteId")
			return
		}
	}
	z.LastActivityAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LastActivityAt")
		return
	}
	z.IsBot, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "IsBot")
		return
	}
	z.BotDescription, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "BotDescription")
		return
	}
	z.BotLastIconUpdate, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "BotLastIconUpdate")
		return
	}
	z.TermsOfServiceId, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceId")
		return
	}
	z.TermsOfServiceCreateAt, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "TermsOfServiceCreateAt")
		return
	}
	z.DisableWelcomeEmail, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "DisableWelcomeEmail")
		return
	}
	z.Bio, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Bio")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *User) Msgsize() (s int) {
	s = 3 + msgp.StringPrefixSize + len(z.Id) + msgp.Int64Size + msgp.Int64Size + msgp.Int64Size + msgp.StringPrefixSize + len(z.Username) + msgp.StringPrefixSize + len(z.Password)
	if z.AuthData == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.AuthData)
	}
	s += msgp.StringPrefixSize + len(z.AuthService) + msgp.StringPrefixSize + len(z.Email) + msgp.BoolSize + msgp.StringPrefixSize + len(z.Nickname) + msgp.StringPrefixSize + len(z.FirstName) + msgp.StringPrefixSize + len(z.LastName) + msgp.StringPrefixSize + len(z.Position) + msgp.StringPrefixSize + len(z.Roles) + msgp.BoolSize + z.Props.Msgsize() + z.NotifyProps.Msgsize() + msgp.Int64Size + msgp.Int64Size + msgp.IntSize + msgp.StringPrefixSize + len(z.Locale) + z.Timezone.Msgsize() + msgp.BoolSize + msgp.StringPrefixSize + len(z.MfaSecret)
	if z.RemoteId == nil {
		s += msgp.NilSize
	} else {
		s += msgp.StringPrefixSize + len(*z.RemoteId)
	}
	s += msgp.Int64Size + msgp.BoolSize + msgp.StringPrefixSize + len(z.BotDescription) + msgp.Int64Size + msgp.StringPrefixSize + len(z.TermsOfServiceId) + msgp.Int64Size + msgp.BoolSize + msgp.StringPrefixSize + len(z.Bio)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserMap) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0003 uint32
	zb0003, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if (*z) == nil {
		(*z) = make(UserMap, zb0003)
	} else if len((*z)) > 0 {
		for key := range *z {
			delete((*z), key)
		}
	}
	for zb0003 > 0 {
		zb0003--
		var zb0001 string
		var zb0002 *User
		zb0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
			zb0002 = nil
		} else {
			if zb0002 == nil {
				zb0002 = new(User)
			}
			err = zb0002.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
		}
		(*z)[zb0001] = zb0002
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z UserMap) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0004, zb0005 := range z {
		err = en.WriteString(zb0004)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0005 == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = zb0005.EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, zb0004)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z UserMap) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, uint32(len(z)))
	for zb0004, zb0005 := range z {
		o = msgp.AppendString(o, zb0004)
		if zb0005 == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zb0005.MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, zb0004)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserMap) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if (*z) == nil {
		(*z) = make(UserMap, zb0003)
	} else if len((*z)) > 0 {
		for key := range *z {
			delete((*z), key)
		}
	}
	for zb0003 > 0 {
		var zb0001 string
		var zb0002 *User
		zb0003--
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			zb0002 = nil
		} else {
			if zb0002 == nil {
				zb0002 = new(User)
			}
			bts, err = zb0002.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
		}
		(*z)[zb0001] = zb0002
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z UserMap) Msgsize() (s int) {
	s = msgp.MapHeaderSize
	if z != nil {
		for zb0004, zb0005 := range z {
			_ = zb0005
			s += msgp.StringPrefixSize + len(zb0004)
			if zb0005 == nil {
				s += msgp.NilSize
			} else {
				s += zb0005.Msgsize()
			}
		}
	}
	return
}
