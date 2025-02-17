//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package waf

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

var (
	_ = fmt.Sprintf("dummy for fmt import use")
)

// Create CustomAPI GRPC Client satisfying server.CustomClient
type CustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomAPIGrpcClient) doRPCVirtualHostWafStatus(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &VirtualHostWafStatusReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.waf.VirtualHostWafStatusReq", yamlReq)
	}
	rsp, err := c.grpcClient.VirtualHostWafStatus(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}
	if cco.YAMLReq == "" {
		return nil, fmt.Errorf("Error, empty request body")
	}
	ctx = client.AddHdrsToCtx(cco.Headers, ctx)

	rsp, err := rpcFn(ctx, cco.YAMLReq, cco.GrpcCallOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using GRPC")
	}
	if cco.OutCallResponse != nil {
		cco.OutCallResponse.ProtoMsg = rsp
	}
	return rsp, nil
}

func NewCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["VirtualHostWafStatus"] = ccl.doRPCVirtualHostWafStatus

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPI REST Client satisfying server.CustomClient
type CustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomAPIRestClient) doRPCVirtualHostWafStatus(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &VirtualHostWafStatusReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.waf.VirtualHostWafStatusReq: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post":
		jsn, err := req.ToJSON()
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		newReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP POST request for custom API")
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &VirtualHostWafStatusRsp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.waf.VirtualHostWafStatusRsp", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}

	rsp, err := rpcFn(ctx, cco)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using Rest")
	}
	return rsp, nil
}

func NewCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["VirtualHostWafStatus"] = ccl.doRPCVirtualHostWafStatus

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPIInprocClient

// INPROC Client (satisfying CustomAPIClient interface)
type CustomAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomAPIInprocClient) VirtualHostWafStatus(ctx context.Context, in *VirtualHostWafStatusReq, opts ...grpc.CallOption) (*VirtualHostWafStatusRsp, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.waf.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPISrv", ah)
	}

	var (
		rsp *VirtualHostWafStatusRsp
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.waf.VirtualHostWafStatusReq", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.VirtualHostWafStatus' operation on 'waf'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.waf.CustomAPI.VirtualHostWafStatus"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.VirtualHostWafStatus(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.waf.VirtualHostWafStatusRsp", rsp)...)

	return rsp, nil
}

func NewCustomAPIInprocClient(svc svcfw.Service) CustomAPIClient {
	return &CustomAPIInprocClient{svc: svc}
}

// RegisterGwCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomAPIHandlerClient(ctx, mux, NewCustomAPIInprocClient(s))
}

var CustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "WAF object",
        "description": "WAF Object to configure a WAF instance which can be associated with virtual_hosts or routes.",
        "version": "version not set"
    },
    "schemes": [
        "http",
        "https"
    ],
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "tags": null,
    "paths": {
        "/public/namespaces/{namespace}/waf/virtual_host/status/{name}": {
            "get": {
                "summary": "Virtual Host WAF Status",
                "description": "Virtual Host WAF Status API is used to get information about the exact configuration, including\na list of waf instances and their respective waf-rules that are currently configured for various\nroutes on a given virtual_host identified by (Namespace, Name).",
                "operationId": "ves.io.schema.waf.CustomAPI.VirtualHostWafStatus",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/wafVirtualHostWafStatusRsp"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "description": "Namespace\n\nx-example: \"blogging-app-namespace-1\"\nNamespace of the virtual host",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-example: \"greatblogs-vhost\"\nName of the virtual host for which waf status is requested",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Virtual Host Name"
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-waf-CustomAPI-VirtualHostWafStatus"
                },
                "x-ves-proto-rpc": "ves.io.schema.waf.CustomAPI.VirtualHostWafStatus"
            },
            "x-displayname": "WAF",
            "x-ves-proto-service": "ves.io.schema.waf.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "schemaObjectRefType": {
            "type": "object",
            "description": "This type establishes a 'direct reference' from one object(the referrer) to another(the referred). \nSuch a reference is in form of tenant/namespace/name for public API and Uid for private API\nThis type of reference is called direct because the relation is explicit and concrete (as opposed\nto selector reference which builds a group based on labels of selectee objects)",
            "title": "ObjectRefType",
            "x-displayname": "Object reference",
            "x-ves-proto-message": "ves.io.schema.ObjectRefType",
            "properties": {
                "kind": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then kind will hold the referred object's kind (e.g. \"route\")\n\nExample: - \"virtual_site\"-",
                    "title": "kind",
                    "x-displayname": "Kind",
                    "x-ves-example": "virtual_site"
                },
                "name": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then name will hold the referred object's(e.g. route's) name.\n\nExample: - \"contactus-route\"-",
                    "title": "name",
                    "x-displayname": "Name",
                    "x-ves-example": "contactus-route"
                },
                "namespace": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then namespace will hold the referred object's(e.g. route's) namespace.\n\nExample: - \"ns1\"-",
                    "title": "namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1"
                },
                "tenant": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then tenant will hold the referred object's(e.g. route's) tenant.\n\nExample: - \"acmecorp\"-",
                    "title": "tenant",
                    "x-displayname": "Tenant",
                    "x-ves-example": "acmecorp"
                },
                "uid": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then uid will hold the referred object's(e.g. route's) uid.\n\nExample: - \"d15f1fad-4d37-48c0-8706-df1824d76d31\"-",
                    "title": "uid",
                    "x-displayname": "UID",
                    "x-ves-example": "d15f1fad-4d37-48c0-8706-df1824d76d31"
                }
            }
        },
        "schemaWafModeType": {
            "type": "string",
            "description": "The mode of operation for Web Application Firewall\n\nBlock on detection\nOnly raise alert on detection",
            "title": "WafModeType",
            "enum": [
                "BLOCK",
                "ALERT_ONLY"
            ],
            "default": "BLOCK",
            "x-displayname": "WAF Mode",
            "x-ves-proto-enum": "ves.io.schema.WafModeType"
        },
        "wafVirtualHostWafStatusRsp": {
            "type": "object",
            "description": "Response is a list of detailed rule configurations currently enabled for the given virtual_host.",
            "title": "Virtual Host WAF Status Response",
            "x-displayname": "Virtual Host WAF Rules Status Response",
            "x-ves-proto-message": "ves.io.schema.waf.VirtualHostWafStatusRsp",
            "properties": {
                "waf_status": {
                    "type": "array",
                    "description": " Detailed configuration of all WAF instances under this virtual host",
                    "title": "Virtual Host WAF Status",
                    "items": {
                        "$ref": "#/definitions/wafWafStatus"
                    },
                    "x-displayname": "Virtual Host WAF Status"
                }
            }
        },
        "wafWafStatus": {
            "type": "object",
            "description": "A list of detailed rule configurations currently enabled for a WAF instance.",
            "title": "WAF Status",
            "x-displayname": "WAF Status",
            "x-ves-proto-message": "ves.io.schema.waf.WafStatus",
            "properties": {
                "waf_ref": {
                    "type": "array",
                    "description": " WAF instance associated with this virtual host.",
                    "title": "WAF Instance",
                    "items": {
                        "$ref": "#/definitions/schemaObjectRefType"
                    },
                    "x-displayname": "WAF Instance"
                },
                "waf_rules_ref": {
                    "type": "array",
                    "description": " WAF Rules instance associated with this virtual host.",
                    "title": "WAF Rules Instance",
                    "items": {
                        "$ref": "#/definitions/schemaObjectRefType"
                    },
                    "x-displayname": "WAF Rules Instance"
                },
                "waf_rules_status": {
                    "type": "array",
                    "description": " WAF rules configured for this waf instance",
                    "title": "WAF Rules Status",
                    "items": {
                        "$ref": "#/definitions/waf_rulesWafRulesStatus"
                    },
                    "x-displayname": "WAF Rules Status"
                }
            }
        },
        "waf_rulesRuleModeType": {
            "type": "string",
            "description": "Specify whether rule is to be included or excluded.\n\nExclude\nInclude",
            "title": "RuleModeType",
            "enum": [
                "EXCLUDE",
                "INCLUDE"
            ],
            "default": "EXCLUDE",
            "x-displayname": "Rule Mode",
            "x-ves-proto-enum": "ves.io.schema.waf_rules.RuleModeType"
        },
        "waf_rulesRules": {
            "type": "object",
            "description": "Every WAF rule will have these properties associated",
            "title": "Rule",
            "x-displayname": "Rule",
            "x-ves-proto-message": "ves.io.schema.waf_rules.Rules",
            "properties": {
                "description": {
                    "type": "string",
                    "description": " This is the brief description of the rule.\n\nExample: - \"IE XSS Filters - Attack Detected.\"-",
                    "title": "description",
                    "x-displayname": "Description",
                    "x-ves-example": "IE XSS Filters - Attack Detected."
                },
                "id": {
                    "type": "integer",
                    "description": " WAF rule ID which is a unique ID with in waf_rules object.\n Generated alerts will have the id displayed in alert.\n\nExample: - \"941210\"-",
                    "title": "id",
                    "format": "int64",
                    "x-displayname": "ID",
                    "x-ves-example": "941210"
                },
                "mode": {
                    "description": " Whether the Rule is excluded or included",
                    "title": "mode",
                    "$ref": "#/definitions/waf_rulesRuleModeType",
                    "x-displayname": "Mode"
                },
                "severity": {
                    "description": " Severity of the rule.",
                    "title": "severity",
                    "$ref": "#/definitions/waf_rulesSeverityType",
                    "x-displayname": "Severity"
                },
                "tags": {
                    "type": "array",
                    "description": "Tags are a set of string labels associated with a rule. For eg a particular rule may be under\n\"attack-sqli\" and \"attack-protocol\" tags. This is used by user to find out Tags associated with a\nRule.\n\nExample: - \"[\"attack-sqli\", \"attack-protocol\"]\"-",
                    "title": "tags",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Tags",
                    "x-ves-example": "[\"attack-sqli\", \"attack-protocol\"]"
                }
            }
        },
        "waf_rulesSeverityType": {
            "type": "string",
            "description": "Rule severity as defined in the rule\n\nEmergency Level\nAlert Level\nCritical Level\nError Level\nWarning Level\nNotice Level\nInfo Level\nDebug Level",
            "title": "SeverityType",
            "enum": [
                "EMERGENCY",
                "ALERT",
                "CRITICAL",
                "ERROR",
                "WARNING",
                "NOTICE",
                "INFO",
                "DEBUG"
            ],
            "default": "EMERGENCY",
            "x-displayname": "Severity",
            "x-ves-proto-enum": "ves.io.schema.waf_rules.SeverityType"
        },
        "waf_rulesWafRulesStatus": {
            "type": "object",
            "description": "Detailed information about the current configuration for a \nwaf-rules object",
            "title": "WAF Rules Status Response",
            "x-displayname": "WAF Rules Status Response",
            "x-ves-proto-message": "ves.io.schema.waf_rules.WafRulesStatus",
            "properties": {
                "anomaly_score_threshold": {
                    "type": "integer",
                    "description": " When A WAF rule hits on inspection of request/response, it causes that http transaction's\n anomaly score to go up by an amount determined by the rule. Anomaly Score Threshold\n is used to set the maximum value of the per transaction anomaly score beyond which \n WAF will alert or block the transaction depending on WAF mode.\n\nExample: - \"4\"-",
                    "title": "anomaly_score_threshold",
                    "format": "int64",
                    "x-displayname": "Anomaly Score Threshold",
                    "x-ves-example": "4"
                },
                "mode": {
                    "description": " Is the WAF instance configured to be in Blocking or Alert-Only mode",
                    "title": "WAF Mode",
                    "$ref": "#/definitions/schemaWafModeType",
                    "x-displayname": "Mode"
                },
                "name": {
                    "type": "string",
                    "description": " Name of the waf-rules object\n\nExample: - \"greatblogs-waf-rules\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "greatblogs-waf-rules"
                },
                "paranoia_level": {
                    "type": "integer",
                    "description": " Paranoia level is used to tune the sensitivity at which WAF is supposed to alert/block.\n High paranoia level means that WAF will alert/block requests that look even slightly suspicious.\n Low paranoia level means that WAF will alert/block only those requests that highly suspicious.\n\nExample: - \"2\"-",
                    "title": "paranoia_level",
                    "format": "int64",
                    "x-displayname": "Paranoia Level",
                    "x-ves-example": "2"
                },
                "rules": {
                    "type": "array",
                    "description": " List of all rules including information about whether the rule is included or excluded",
                    "title": "rules",
                    "items": {
                        "$ref": "#/definitions/waf_rulesRules"
                    },
                    "x-displayname": "WAF Rules"
                }
            }
        }
    },
    "x-displayname": "WAF",
    "x-ves-proto-file": "ves.io/schema/waf/public_customapi.proto"
}`
