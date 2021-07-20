//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package api_credential

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_user "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/user"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CreateRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *CreateRequest) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetSpec().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting CreateRequest.spec")
	}

	return nil
}

func (m *CreateRequest) DeepCopy() *CreateRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCreateRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["expiration_days"]; exists {

		vOpts := append(opts, db.WithValidateField("expiration_days"))
		if err := fv(ctx, m.GetExpirationDays(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["expiration_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("expiration_timestamp"))
		if err := fv(ctx, m.GetExpirationTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateRequestValidator = func() *ValidateCreateRequest {
	v := &ValidateCreateRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["spec"] = CustomCreateSpecTypeValidator().Validate

	return v
}()

func CreateRequestValidator() db.Validator {
	return DefaultCreateRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CreateResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateResponse) String() string {
	return "REDACTED"
}

func (m *CreateResponse) GoString() string {
	return "REDACTED"
}

// Redact squashes sensitive info in m (in-place)
func (m *CreateResponse) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.Data = ""
	m.Name = ""

	return nil
}

func (m *CreateResponse) DeepCopy() *CreateResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateResponseValidator().Validate(ctx, m, opts...)
}

type ValidateCreateResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["data"]; exists {

		vOpts := append(opts, db.WithValidateField("data"))
		if err := fv(ctx, m.GetData(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateResponseValidator = func() *ValidateCreateResponse {
	v := &ValidateCreateResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CreateResponseValidator() db.Validator {
	return DefaultCreateResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *CreateServiceCredentialsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateServiceCredentialsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateServiceCredentialsRequest) String() string {
	if m == nil {
		return ""
	}
	copy := m.DeepCopy()
	copy.Password = ""

	return copy.string()
}

func (m *CreateServiceCredentialsRequest) GoString() string {
	copy := m.DeepCopy()
	copy.Password = ""

	return copy.goString()
}

// Redact squashes sensitive info in m (in-place)
func (m *CreateServiceCredentialsRequest) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.Password = ""

	return nil
}

func (m *CreateServiceCredentialsRequest) DeepCopy() *CreateServiceCredentialsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateServiceCredentialsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateServiceCredentialsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateServiceCredentialsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateServiceCredentialsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCreateServiceCredentialsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateServiceCredentialsRequest) TypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(APICredentialType)
		return int32(i)
	}
	// APICredentialType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, APICredentialType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for type")
	}

	return validatorFn, nil
}

func (v *ValidateCreateServiceCredentialsRequest) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for password")
	}

	return validatorFn, nil
}

func (v *ValidateCreateServiceCredentialsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateServiceCredentialsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateServiceCredentialsRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["expiration_days"]; exists {

		vOpts := append(opts, db.WithValidateField("expiration_days"))
		if err := fv(ctx, m.GetExpirationDays(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["expiration_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("expiration_timestamp"))
		if err := fv(ctx, m.GetExpirationTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace_roles"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace_roles"))
		for idx, item := range m.GetNamespaceRoles() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_k8s_name"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_k8s_name"))
		if err := fv(ctx, m.GetVirtualK8SName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_k8s_namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_k8s_namespace"))
		if err := fv(ctx, m.GetVirtualK8SNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateServiceCredentialsRequestValidator = func() *ValidateCreateServiceCredentialsRequest {
	v := &ValidateCreateServiceCredentialsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhType := v.TypeValidationRuleHandler
	rulesType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhType(rulesType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateServiceCredentialsRequest.type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["type"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateServiceCredentialsRequest.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	v.FldValidators["namespace_roles"] = ves_io_schema_user.NamespaceRoleTypeValidator().Validate

	return v
}()

func CreateServiceCredentialsRequestValidator() db.Validator {
	return DefaultCreateServiceCredentialsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CustomCreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CustomCreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CustomCreateSpecType) String() string {
	if m == nil {
		return ""
	}
	copy := m.DeepCopy()
	copy.Password = ""

	return copy.string()
}

func (m *CustomCreateSpecType) GoString() string {
	copy := m.DeepCopy()
	copy.Password = ""

	return copy.goString()
}

// Redact squashes sensitive info in m (in-place)
func (m *CustomCreateSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.Password = ""

	return nil
}

func (m *CustomCreateSpecType) DeepCopy() *CustomCreateSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CustomCreateSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CustomCreateSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CustomCreateSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CustomCreateSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateCustomCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCustomCreateSpecType) TypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(APICredentialType)
		return int32(i)
	}
	// APICredentialType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, APICredentialType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for type")
	}

	return validatorFn, nil
}

func (v *ValidateCustomCreateSpecType) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for password")
	}

	return validatorFn, nil
}

func (v *ValidateCustomCreateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CustomCreateSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CustomCreateSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_k8s_name"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_k8s_name"))
		if err := fv(ctx, m.GetVirtualK8SName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_k8s_namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_k8s_namespace"))
		if err := fv(ctx, m.GetVirtualK8SNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCustomCreateSpecTypeValidator = func() *ValidateCustomCreateSpecType {
	v := &ValidateCustomCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhType := v.TypeValidationRuleHandler
	rulesType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhType(rulesType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CustomCreateSpecType.type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["type"] = vFn

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CustomCreateSpecType.password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["password"] = vFn

	return v
}()

func CustomCreateSpecTypeValidator() db.Validator {
	return DefaultCustomCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *DeleteRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DeleteRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DeleteRequest) DeepCopy() *DeleteRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DeleteRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DeleteRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DeleteRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DeleteRequestValidator().Validate(ctx, m, opts...)
}

type ValidateDeleteRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDeleteRequest) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateDeleteRequest) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateDeleteRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DeleteRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DeleteRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDeleteRequestValidator = func() *ValidateDeleteRequest {
	v := &ValidateDeleteRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for DeleteRequest.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for DeleteRequest.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	return v
}()

func DeleteRequestValidator() db.Validator {
	return DefaultDeleteRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetRequest) DeepCopy() *GetRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetRequest) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateGetRequest) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateGetRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetRequestValidator = func() *ValidateGetRequest {
	v := &ValidateGetRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetRequest.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetRequest.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	return v
}()

func GetRequestValidator() db.Validator {
	return DefaultGetRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetResponse) DeepCopy() *GetResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo

	return drInfos, nil
}

type ValidateGetResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["object"]; exists {

		vOpts := append(opts, db.WithValidateField("object"))
		if err := fv(ctx, m.GetObject(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetResponseValidator = func() *ValidateGetResponse {
	v := &ValidateGetResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["object"] = ObjectValidator().Validate

	return v
}()

func GetResponseValidator() db.Validator {
	return DefaultGetResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ListRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListRequest) DeepCopy() *ListRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListRequestValidator().Validate(ctx, m, opts...)
}

type ValidateListRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListRequestValidator = func() *ValidateListRequest {
	v := &ValidateListRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListRequestValidator() db.Validator {
	return DefaultListRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ListResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListResponse) DeepCopy() *ListResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListResponseValidator().Validate(ctx, m, opts...)
}

type ValidateListResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["items"]; exists {

		vOpts := append(opts, db.WithValidateField("items"))
		for idx, item := range m.GetItems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListResponseValidator = func() *ValidateListResponse {
	v := &ValidateListResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListResponseValidator() db.Validator {
	return DefaultListResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ListResponseItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListResponseItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListResponseItem) DeepCopy() *ListResponseItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListResponseItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListResponseItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListResponseItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListResponseItemValidator().Validate(ctx, m, opts...)
}

type ValidateListResponseItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListResponseItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListResponseItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListResponseItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["active"]; exists {

		vOpts := append(opts, db.WithValidateField("active"))
		if err := fv(ctx, m.GetActive(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["create_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("create_timestamp"))
		if err := fv(ctx, m.GetCreateTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["expiry_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("expiry_timestamp"))
		if err := fv(ctx, m.GetExpiryTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["uid"]; exists {

		vOpts := append(opts, db.WithValidateField("uid"))
		if err := fv(ctx, m.GetUid(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_email"]; exists {

		vOpts := append(opts, db.WithValidateField("user_email"))
		if err := fv(ctx, m.GetUserEmail(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListResponseItemValidator = func() *ValidateListResponseItem {
	v := &ValidateListResponseItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListResponseItemValidator() db.Validator {
	return DefaultListResponseItemValidator
}

// augmented methods on protoc/std generated struct

func (m *RenewRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RenewRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RenewRequest) DeepCopy() *RenewRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RenewRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RenewRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RenewRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RenewRequestValidator().Validate(ctx, m, opts...)
}

type ValidateRenewRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRenewRequest) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateRenewRequest) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateRenewRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RenewRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RenewRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["expiration_days"]; exists {

		vOpts := append(opts, db.WithValidateField("expiration_days"))
		if err := fv(ctx, m.GetExpirationDays(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["expiration_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("expiration_timestamp"))
		if err := fv(ctx, m.GetExpirationTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRenewRequestValidator = func() *ValidateRenewRequest {
	v := &ValidateRenewRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RenewRequest.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RenewRequest.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	return v
}()

func RenewRequestValidator() db.Validator {
	return DefaultRenewRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *StatusResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *StatusResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *StatusResponse) DeepCopy() *StatusResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &StatusResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *StatusResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *StatusResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return StatusResponseValidator().Validate(ctx, m, opts...)
}

type ValidateStatusResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateStatusResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*StatusResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *StatusResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultStatusResponseValidator = func() *ValidateStatusResponse {
	v := &ValidateStatusResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func StatusResponseValidator() db.Validator {
	return DefaultStatusResponseValidator
}

func (m *CreateRequest) FromObject(e db.Entry) {
	f := e.DeepCopy().(*DBObject)
	_ = f

}

func (m *CreateRequest) ToObject(e db.Entry) {
	m1 := m.DeepCopy()
	_ = m1
	f := e.(*DBObject)
	_ = f

}
