package api

func init() {
	Swagger.Add("cds", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/cds/cds.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/beta/content/credentials": {
      "post": {
        "summary": "SubmitCredentials",
        "description": "Submit a Chef Cloud Credentials to enable content\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ncontent:credentials:add\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "SubmitCredentials",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cds.response.Credentials"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cds.request.Credentials"
            }
          }
        ],
        "tags": [
          "Cds"
        ]
      }
    },
    "/api/beta/content/enabled": {
      "get": {
        "summary": "IsContentEnabled",
        "description": "Check if the content is enable for this Automate instance. \n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ncontent:credentials:enabled\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "IsContentEnabled",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cds.response.ContentEnabled"
            }
          }
        },
        "tags": [
          "Cds"
        ]
      }
    },
    "/api/beta/content/install": {
      "post": {
        "summary": "InstallContentItem",
        "description": "Installs a content item from its ID\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ncontent:items:install\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "InstallContentItem",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cds.response.InstallContentItem"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cds.request.InstallContentItem"
            }
          }
        ],
        "tags": [
          "Cds"
        ]
      }
    },
    "/api/beta/content/items": {
      "get": {
        "summary": "ListContentItems",
        "description": "Returns a list of metadata for each CDS content. Provides a description and current \nstate of each content item.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ncontent:items:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ListContentItems",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.cds.response.ContentItems"
            }
          }
        },
        "tags": [
          "Cds"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.cds.request.Credentials": {
      "type": "object",
      "properties": {
        "client_id": {
          "type": "string"
        },
        "client_secret": {
          "type": "string"
        },
        "tenant_specific_url": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.cds.request.InstallContentItem": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "request_user": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.cds.response.ContentEnabled": {
      "type": "object",
      "properties": {
        "is_content_enabled": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "chef.automate.api.cds.response.ContentItem": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "platforms": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "can_be_installed": {
          "type": "boolean",
          "format": "boolean"
        },
        "filename": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.cds.response.ContentItems": {
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.cds.response.ContentItem"
          }
        }
      }
    },
    "chef.automate.api.cds.response.Credentials": {
      "type": "object"
    },
    "chef.automate.api.cds.response.InstallContentItem": {
      "type": "object"
    },
    "chef.automate.api.common.ExportData": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string",
          "format": "byte",
          "description": "Exported reports in JSON or CSV."
        }
      }
    },
    "google.protobuf.Any": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string",
          "description": "A URL/resource name that uniquely identifies the type of the serialized\nprotocol buffer message. This string must contain at least\none \"/\" character. The last segment of the URL's path must represent\nthe fully qualified name of the type (as in\n` + "`" + `path/google.protobuf.Duration` + "`" + `). The name should be in a canonical form\n(e.g., leading \".\" is not accepted).\n\nIn practice, teams usually precompile into the binary all types that they\nexpect it to use in the context of Any. However, for URLs which use the\nscheme ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, or no scheme, one can optionally set up a type\nserver that maps type URLs to message definitions as follows:\n\n* If no scheme is provided, ` + "`" + `https` + "`" + ` is assumed.\n* An HTTP GET on the URL must yield a [google.protobuf.Type][]\n  value in binary format, or produce an error.\n* Applications are allowed to cache lookup results based on the\n  URL, or have them precompiled into a binary to avoid any\n  lookup. Therefore, binary compatibility needs to be preserved\n  on changes to types. (Use versioned type names to manage\n  breaking changes.)\n\nNote: this functionality is not currently available in the official\nprotobuf release, and it is not used for type URLs beginning with\ntype.googleapis.com.\n\nSchemes other than ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` (or the empty scheme) might be\nused with implementation specific semantics."
        },
        "value": {
          "type": "string",
          "format": "byte",
          "description": "Must be a valid serialized protocol buffer of the above specified type."
        }
      },
      "description": "` + "`" + `Any` + "`" + ` contains an arbitrary serialized protocol buffer message along with a\nURL that describes the type of the serialized message.\n\nProtobuf library provides support to pack/unpack Any values in the form\nof utility functions or additional generated methods of the Any type.\n\nExample 1: Pack and unpack a message in C++.\n\n    Foo foo = ...;\n    Any any;\n    any.PackFrom(foo);\n    ...\n    if (any.UnpackTo(\u0026foo)) {\n      ...\n    }\n\nExample 2: Pack and unpack a message in Java.\n\n    Foo foo = ...;\n    Any any = Any.pack(foo);\n    ...\n    if (any.is(Foo.class)) {\n      foo = any.unpack(Foo.class);\n    }\n\n Example 3: Pack and unpack a message in Python.\n\n    foo = Foo(...)\n    any = Any()\n    any.Pack(foo)\n    ...\n    if any.Is(Foo.DESCRIPTOR):\n      any.Unpack(foo)\n      ...\n\n Example 4: Pack and unpack a message in Go\n\n     foo := \u0026pb.Foo{...}\n     any, err := ptypes.MarshalAny(foo)\n     ...\n     foo := \u0026pb.Foo{}\n     if err := ptypes.UnmarshalAny(any, foo); err != nil {\n       ...\n     }\n\nThe pack methods provided by protobuf library will by default use\n'type.googleapis.com/full.type.name' as the type URL and the unpack\nmethods only use the fully qualified type name after the last '/'\nin the type URL, for example \"foo.bar.com/x/y.z\" will yield type\nname \"y.z\".\n\n\nJSON\n====\nThe JSON representation of an ` + "`" + `Any` + "`" + ` value uses the regular\nrepresentation of the deserialized, embedded message, with an\nadditional field ` + "`" + `@type` + "`" + ` which contains the type URL. Example:\n\n    package google.profile;\n    message Person {\n      string first_name = 1;\n      string last_name = 2;\n    }\n\n    {\n      \"@type\": \"type.googleapis.com/google.profile.Person\",\n      \"firstName\": \u003cstring\u003e,\n      \"lastName\": \u003cstring\u003e\n    }\n\nIf the embedded message type is well-known and has a custom JSON\nrepresentation, that representation will be embedded adding a field\n` + "`" + `value` + "`" + ` which holds the custom JSON in addition to the ` + "`" + `@type` + "`" + `\nfield. Example (for message [google.protobuf.Duration][]):\n\n    {\n      \"@type\": \"type.googleapis.com/google.protobuf.Duration\",\n      \"value\": \"1.212s\"\n    }"
    },
    "grpc.gateway.runtime.StreamError": {
      "type": "object",
      "properties": {
        "grpc_code": {
          "type": "integer",
          "format": "int32"
        },
        "http_code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "http_status": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/google.protobuf.Any"
          }
        }
      }
    }
  },
  "x-stream-definitions": {
    "chef.automate.api.common.ExportData": {
      "type": "object",
      "properties": {
        "result": {
          "$ref": "#/definitions/chef.automate.api.common.ExportData"
        },
        "error": {
          "$ref": "#/definitions/grpc.gateway.runtime.StreamError"
        }
      },
      "title": "Stream result of chef.automate.api.common.ExportData"
    }
  }
}
`)
}
