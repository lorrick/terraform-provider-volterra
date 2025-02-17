//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package container_registry

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *CreateSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetPassword().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting CreateSpecType.password")
	}

	return nil
}

func (m *CreateSpecType) DeepCopy() *CreateSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) RegistryValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for registry")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) UserNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for user_name")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for password")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) EmailValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for email")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["registry"]; exists {

		vOpts := append(opts, db.WithValidateField("registry"))
		if err := fv(ctx, m.GetRegistry(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_name"]; exists {

		vOpts := append(opts, db.WithValidateField("user_name"))
		if err := fv(ctx, m.GetUserName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhRegistry := v.RegistryValidationRuleHandler
	rulesRegistry := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.hostname":  "true",
	}
	vFn, err = vrhRegistry(rulesRegistry)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.registry: %s", err)
		panic(errMsg)
	}
	v.FldValidators["registry"] = vFn

	vrhUserName := v.UserNameValidationRuleHandler
	rulesUserName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "128",
	}
	vFn, err = vrhUserName(rulesUserName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.user_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["user_name"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	vrhEmail := v.EmailValidationRuleHandler
	rulesEmail := map[string]string{
		"ves.io.schema.rules.string.email": "true",
	}
	vFn, err = vrhEmail(rulesEmail)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.email: %s", err)
		panic(errMsg)
	}
	v.FldValidators["email"] = vFn

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *GetSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetPassword().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GetSpecType.password")
	}

	return nil
}

func (m *GetSpecType) DeepCopy() *GetSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) RegistryValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for registry")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) UserNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for user_name")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for password")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) EmailValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for email")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["registry"]; exists {

		vOpts := append(opts, db.WithValidateField("registry"))
		if err := fv(ctx, m.GetRegistry(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_name"]; exists {

		vOpts := append(opts, db.WithValidateField("user_name"))
		if err := fv(ctx, m.GetUserName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhRegistry := v.RegistryValidationRuleHandler
	rulesRegistry := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.hostname":  "true",
	}
	vFn, err = vrhRegistry(rulesRegistry)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.registry: %s", err)
		panic(errMsg)
	}
	v.FldValidators["registry"] = vFn

	vrhUserName := v.UserNameValidationRuleHandler
	rulesUserName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "128",
	}
	vFn, err = vrhUserName(rulesUserName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.user_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["user_name"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	vrhEmail := v.EmailValidationRuleHandler
	rulesEmail := map[string]string{
		"ves.io.schema.rules.string.email": "true",
	}
	vFn, err = vrhEmail(rulesEmail)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.email: %s", err)
		panic(errMsg)
	}
	v.FldValidators["email"] = vFn

	return v
}()

func GetSpecTypeValidator() db.Validator {
	return DefaultGetSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *GlobalSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetPassword().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GlobalSpecType.password")
	}

	return nil
}

func (m *GlobalSpecType) DeepCopy() *GlobalSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GlobalSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GlobalSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GlobalSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GlobalSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) RegistryValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for registry")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) UserNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for user_name")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for password")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) EmailValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for email")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GlobalSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GlobalSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["registry"]; exists {

		vOpts := append(opts, db.WithValidateField("registry"))
		if err := fv(ctx, m.GetRegistry(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_name"]; exists {

		vOpts := append(opts, db.WithValidateField("user_name"))
		if err := fv(ctx, m.GetUserName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhRegistry := v.RegistryValidationRuleHandler
	rulesRegistry := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.hostname":  "true",
	}
	vFn, err = vrhRegistry(rulesRegistry)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.registry: %s", err)
		panic(errMsg)
	}
	v.FldValidators["registry"] = vFn

	vrhUserName := v.UserNameValidationRuleHandler
	rulesUserName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "128",
	}
	vFn, err = vrhUserName(rulesUserName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.user_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["user_name"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	vrhEmail := v.EmailValidationRuleHandler
	rulesEmail := map[string]string{
		"ves.io.schema.rules.string.email": "true",
	}
	vFn, err = vrhEmail(rulesEmail)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.email: %s", err)
		panic(errMsg)
	}
	v.FldValidators["email"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *ReplaceSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetPassword().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting ReplaceSpecType.password")
	}

	return nil
}

func (m *ReplaceSpecType) DeepCopy() *ReplaceSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReplaceSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReplaceSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReplaceSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReplaceSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) RegistryValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for registry")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) UserNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for user_name")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for password")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.SecretTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) EmailValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for email")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReplaceSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReplaceSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["registry"]; exists {

		vOpts := append(opts, db.WithValidateField("registry"))
		if err := fv(ctx, m.GetRegistry(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_name"]; exists {

		vOpts := append(opts, db.WithValidateField("user_name"))
		if err := fv(ctx, m.GetUserName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhRegistry := v.RegistryValidationRuleHandler
	rulesRegistry := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.hostname":  "true",
	}
	vFn, err = vrhRegistry(rulesRegistry)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.registry: %s", err)
		panic(errMsg)
	}
	v.FldValidators["registry"] = vFn

	vrhUserName := v.UserNameValidationRuleHandler
	rulesUserName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "128",
	}
	vFn, err = vrhUserName(rulesUserName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.user_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["user_name"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	vrhEmail := v.EmailValidationRuleHandler
	rulesEmail := map[string]string{
		"ves.io.schema.rules.string.email": "true",
	}
	vFn, err = vrhEmail(rulesEmail)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.email: %s", err)
		panic(errMsg)
	}
	v.FldValidators["email"] = vFn

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Email = f.GetEmail()
	m.Password = f.GetPassword()
	m.Registry = f.GetRegistry()
	m.UserName = f.GetUserName()
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Email = m1.Email
	f.Password = m1.Password
	f.Registry = m1.Registry
	f.UserName = m1.UserName
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Email = f.GetEmail()
	m.Password = f.GetPassword()
	m.Registry = f.GetRegistry()
	m.UserName = f.GetUserName()
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Email = m1.Email
	f.Password = m1.Password
	f.Registry = m1.Registry
	f.UserName = m1.UserName
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Email = f.GetEmail()
	m.Password = f.GetPassword()
	m.Registry = f.GetRegistry()
	m.UserName = f.GetUserName()
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Email = m1.Email
	f.Password = m1.Password
	f.Registry = m1.Registry
	f.UserName = m1.UserName
}
