/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			router.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			router.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/v2/",
		Index,
	},

	{
		"Call123TestSpecialTags",
		strings.ToUpper("Patch"),
		"/v2/another-fake/dummy",
		Call123TestSpecialTags,
	},

	{
		"FooGet",
		strings.ToUpper("Get"),
		"/v2/foo",
		FooGet,
	},

	{
		"FakeHealthGet",
		strings.ToUpper("Get"),
		"/v2/fake/health",
		FakeHealthGet,
	},

	{
		"FakeOuterBooleanSerialize",
		strings.ToUpper("Post"),
		"/v2/fake/outer/boolean",
		FakeOuterBooleanSerialize,
	},

	{
		"FakeOuterCompositeSerialize",
		strings.ToUpper("Post"),
		"/v2/fake/outer/composite",
		FakeOuterCompositeSerialize,
	},

	{
		"FakeOuterNumberSerialize",
		strings.ToUpper("Post"),
		"/v2/fake/outer/number",
		FakeOuterNumberSerialize,
	},

	{
		"FakeOuterStringSerialize",
		strings.ToUpper("Post"),
		"/v2/fake/outer/string",
		FakeOuterStringSerialize,
	},

	{
		"TestBodyWithFileSchema",
		strings.ToUpper("Put"),
		"/v2/fake/body-with-file-schema",
		TestBodyWithFileSchema,
	},

	{
		"TestBodyWithQueryParams",
		strings.ToUpper("Put"),
		"/v2/fake/body-with-query-params",
		TestBodyWithQueryParams,
	},

	{
		"TestClientModel",
		strings.ToUpper("Patch"),
		"/v2/fake",
		TestClientModel,
	},

	{
		"TestEndpointParameters",
		strings.ToUpper("Post"),
		"/v2/fake",
		TestEndpointParameters,
	},

	{
		"TestEnumParameters",
		strings.ToUpper("Get"),
		"/v2/fake",
		TestEnumParameters,
	},

	{
		"TestGroupParameters",
		strings.ToUpper("Delete"),
		"/v2/fake",
		TestGroupParameters,
	},

	{
		"TestInlineAdditionalProperties",
		strings.ToUpper("Post"),
		"/v2/fake/inline-additionalProperties",
		TestInlineAdditionalProperties,
	},

	{
		"TestJsonFormData",
		strings.ToUpper("Get"),
		"/v2/fake/jsonFormData",
		TestJsonFormData,
	},

	{
		"TestClassname",
		strings.ToUpper("Patch"),
		"/v2/fake_classname_test",
		TestClassname,
	},

	{
		"AddPet",
		strings.ToUpper("Post"),
		"/v2/pet",
		AddPet,
	},

	{
		"DeletePet",
		strings.ToUpper("Delete"),
		"/v2/pet/:petId",
		DeletePet,
	},

	{
		"FindPetsByStatus",
		strings.ToUpper("Get"),
		"/v2/pet/findByStatus",
		FindPetsByStatus,
	},

	{
		"FindPetsByTags",
		strings.ToUpper("Get"),
		"/v2/pet/findByTags",
		FindPetsByTags,
	},

	{
		"GetPetById",
		strings.ToUpper("Get"),
		"/v2/pet/:petId",
		GetPetById,
	},

	{
		"UpdatePet",
		strings.ToUpper("Put"),
		"/v2/pet",
		UpdatePet,
	},

	{
		"UpdatePetWithForm",
		strings.ToUpper("Post"),
		"/v2/pet/:petId",
		UpdatePetWithForm,
	},

	{
		"UploadFile",
		strings.ToUpper("Post"),
		"/v2/pet/:petId/uploadImage",
		UploadFile,
	},

	{
		"UploadFileWithRequiredFile",
		strings.ToUpper("Post"),
		"/v2/fake/:petId/uploadImageWithRequiredFile",
		UploadFileWithRequiredFile,
	},

	{
		"DeleteOrder",
		strings.ToUpper("Delete"),
		"/v2/store/order/:order_id",
		DeleteOrder,
	},

	{
		"GetInventory",
		strings.ToUpper("Get"),
		"/v2/store/inventory",
		GetInventory,
	},

	{
		"GetOrderById",
		strings.ToUpper("Get"),
		"/v2/store/order/:order_id",
		GetOrderById,
	},

	{
		"PlaceOrder",
		strings.ToUpper("Post"),
		"/v2/store/order",
		PlaceOrder,
	},

	{
		"CreateUser",
		strings.ToUpper("Post"),
		"/v2/user",
		CreateUser,
	},

	{
		"CreateUsersWithArrayInput",
		strings.ToUpper("Post"),
		"/v2/user/createWithArray",
		CreateUsersWithArrayInput,
	},

	{
		"CreateUsersWithListInput",
		strings.ToUpper("Post"),
		"/v2/user/createWithList",
		CreateUsersWithListInput,
	},

	{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/v2/user/:username",
		DeleteUser,
	},

	{
		"GetUserByName",
		strings.ToUpper("Get"),
		"/v2/user/:username",
		GetUserByName,
	},

	{
		"LoginUser",
		strings.ToUpper("Get"),
		"/v2/user/login",
		LoginUser,
	},

	{
		"LogoutUser",
		strings.ToUpper("Get"),
		"/v2/user/logout",
		LogoutUser,
	},

	{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/v2/user/:username",
		UpdateUser,
	},
}
