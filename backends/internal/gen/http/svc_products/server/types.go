// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-products HTTP server types
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package server

import (
	svcproducts "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_products"
	svcproductsviews "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_products/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateProductRequestBody is the type of the "svc-products" service
// "createProduct" endpoint HTTP request body.
type CreateProductRequestBody struct {
	// Title's product
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Product description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Last part of the url which use to idepntify the user
	Handle *string `form:"handle,omitempty" json:"handle,omitempty" xml:"handle,omitempty"`
	// Product's status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Product tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Vendor's product
	VendorID *string `form:"vendorId,omitempty" json:"vendorId,omitempty" xml:"vendorId,omitempty"`
	// Product variants
	Variants []*ProductVariantInputRequestBody `form:"variants,omitempty" json:"variants,omitempty" xml:"variants,omitempty"`
	Medias   []*ProductMediaInputRequestBody   `form:"medias,omitempty" json:"medias,omitempty" xml:"medias,omitempty"`
}

// UpdateProductByIDRequestBody is the type of the "svc-products" service
// "updateProductById" endpoint HTTP request body.
type UpdateProductByIDRequestBody struct {
	// Title's product
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Product description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Last part of the url which use to idepntify the user
	Handle *string `form:"handle,omitempty" json:"handle,omitempty" xml:"handle,omitempty"`
	// Product's status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Product tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Vendor's product
	VendorID *string `form:"vendorId,omitempty" json:"vendorId,omitempty" xml:"vendorId,omitempty"`
	// Product variants
	Variants []*ProductVariantInputRequestBodyRequestBody `form:"variants,omitempty" json:"variants,omitempty" xml:"variants,omitempty"`
	Medias   []*ProductMediaInputRequestBodyRequestBody   `form:"medias,omitempty" json:"medias,omitempty" xml:"medias,omitempty"`
}

// ListProductResponseBody is the type of the "svc-products" service
// "listProduct" endpoint HTTP response body.
type ListProductResponseBody struct {
	// Data
	Data ProductResponseBodyCollection `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// Pagination information
	PageInfo *PageInfoResponseBody `form:"pageInfo,omitempty" json:"pageInfo,omitempty" xml:"pageInfo,omitempty"`
}

// GetProductByIDResponseBody is the type of the "svc-products" service
// "getProductById" endpoint HTTP response body.
type GetProductByIDResponseBody struct {
	// Key ID
	ID int `form:"id" json:"id" xml:"id"`
	// Title
	Title string `form:"title" json:"title" xml:"title"`
	// Product description
	Description string `form:"description" json:"description" xml:"description"`
	// Handle
	Handle string              `form:"handle" json:"handle" xml:"handle"`
	Vendor *VendorResponseBody `form:"vendor,omitempty" json:"vendor,omitempty" xml:"vendor,omitempty"`
	// Product tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// The product's status on ecommerce site
	Status   string                        `form:"status" json:"status" xml:"status"`
	Variants []*ProductVariantResponseBody `form:"variants,omitempty" json:"variants,omitempty" xml:"variants,omitempty"`
	Medias   []*ProductMediaResponseBody   `form:"medias,omitempty" json:"medias,omitempty" xml:"medias,omitempty"`
}

// CreateProductResponseBody is the type of the "svc-products" service
// "createProduct" endpoint HTTP response body.
type CreateProductResponseBody struct {
	// Key ID
	ID int `form:"id" json:"id" xml:"id"`
	// Title
	Title string `form:"title" json:"title" xml:"title"`
	// Product description
	Description string `form:"description" json:"description" xml:"description"`
	// Handle
	Handle string              `form:"handle" json:"handle" xml:"handle"`
	Vendor *VendorResponseBody `form:"vendor,omitempty" json:"vendor,omitempty" xml:"vendor,omitempty"`
	// Product tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// The product's status on ecommerce site
	Status   string                        `form:"status" json:"status" xml:"status"`
	Variants []*ProductVariantResponseBody `form:"variants,omitempty" json:"variants,omitempty" xml:"variants,omitempty"`
	Medias   []*ProductMediaResponseBody   `form:"medias,omitempty" json:"medias,omitempty" xml:"medias,omitempty"`
}

// UpdateProductByIDResponseBody is the type of the "svc-products" service
// "updateProductById" endpoint HTTP response body.
type UpdateProductByIDResponseBody struct {
	// Key ID
	ID int `form:"id" json:"id" xml:"id"`
	// Title
	Title string `form:"title" json:"title" xml:"title"`
	// Product description
	Description string `form:"description" json:"description" xml:"description"`
	// Handle
	Handle string              `form:"handle" json:"handle" xml:"handle"`
	Vendor *VendorResponseBody `form:"vendor,omitempty" json:"vendor,omitempty" xml:"vendor,omitempty"`
	// Product tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// The product's status on ecommerce site
	Status   string                        `form:"status" json:"status" xml:"status"`
	Variants []*ProductVariantResponseBody `form:"variants,omitempty" json:"variants,omitempty" xml:"variants,omitempty"`
	Medias   []*ProductMediaResponseBody   `form:"medias,omitempty" json:"medias,omitempty" xml:"medias,omitempty"`
}

// ListProductBadRequestResponseBody is the type of the "svc-products" service
// "listProduct" endpoint HTTP response body for the "BadRequest" error.
type ListProductBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetProductByIDNotFoundResponseBody is the type of the "svc-products" service
// "getProductById" endpoint HTTP response body for the "NotFound" error.
type GetProductByIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateProductConflictResponseBody is the type of the "svc-products" service
// "createProduct" endpoint HTTP response body for the "Conflict" error.
type CreateProductConflictResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateProductByIDConflictResponseBody is the type of the "svc-products"
// service "updateProductById" endpoint HTTP response body for the "Conflict"
// error.
type UpdateProductByIDConflictResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteProductByIDNotFoundResponseBody is the type of the "svc-products"
// service "deleteProductById" endpoint HTTP response body for the "NotFound"
// error.
type DeleteProductByIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ProductResponseBodyCollection is used to define fields on response body
// types.
type ProductResponseBodyCollection []*ProductResponseBody

// ProductResponseBody is used to define fields on response body types.
type ProductResponseBody struct {
	// Key ID
	ID int `form:"id" json:"id" xml:"id"`
	// Title
	Title string `form:"title" json:"title" xml:"title"`
	// Product description
	Description string `form:"description" json:"description" xml:"description"`
	// Handle
	Handle string              `form:"handle" json:"handle" xml:"handle"`
	Vendor *VendorResponseBody `form:"vendor,omitempty" json:"vendor,omitempty" xml:"vendor,omitempty"`
	// Product tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// The product's status on ecommerce site
	Status   string                        `form:"status" json:"status" xml:"status"`
	Variants []*ProductVariantResponseBody `form:"variants,omitempty" json:"variants,omitempty" xml:"variants,omitempty"`
	Medias   []*ProductMediaResponseBody   `form:"medias,omitempty" json:"medias,omitempty" xml:"medias,omitempty"`
}

// VendorResponseBody is used to define fields on response body types.
type VendorResponseBody struct {
	// Key ID
	ID   *int   `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name string `form:"name" json:"name" xml:"name"`
}

// ProductVariantResponseBody is used to define fields on response body types.
type ProductVariantResponseBody struct {
	// Key ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Color variant option
	ColorName *string `form:"colorName,omitempty" json:"colorName,omitempty" xml:"colorName,omitempty"`
	// Color in HEX value that would be used on the variant picker
	ColorHex *string `form:"colorHex,omitempty" json:"colorHex,omitempty" xml:"colorHex,omitempty"`
	// Price on cents
	Price *int `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// Product where the variant belongs to
	ProductID *string `form:"productId,omitempty" json:"productId,omitempty" xml:"productId,omitempty"`
	// ProductMedia which would be focus when a variant is picked by the user
	FeatureMediaID *string `form:"featureMediaId,omitempty" json:"featureMediaId,omitempty" xml:"featureMediaId,omitempty"`
}

// ProductMediaResponseBody is used to define fields on response body types.
type ProductMediaResponseBody struct {
	// Key ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// URL to the media
	URL       *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	MediaType *string `form:"mediaType,omitempty" json:"mediaType,omitempty" xml:"mediaType,omitempty"`
}

// PageInfoResponseBody is used to define fields on response body types.
type PageInfoResponseBody struct {
	// The starting cursor for pagination
	StartCursor string `form:"startCursor" json:"startCursor" xml:"startCursor"`
	// The ending cursor for pagination
	EndCursor string `form:"endCursor" json:"endCursor" xml:"endCursor"`
	// Indicates if there are more results available
	HasMore bool `form:"hasMore" json:"hasMore" xml:"hasMore"`
	// Total number of resources available
	TotalResource int `form:"totalResource" json:"totalResource" xml:"totalResource"`
}

// ProductVariantInputRequestBody is used to define fields on request body
// types.
type ProductVariantInputRequestBody struct {
	// Color variant option
	ColorName *string `form:"colorName,omitempty" json:"colorName,omitempty" xml:"colorName,omitempty"`
	// Color in HEX value that would be used on the variant picker
	ColorHex *string `form:"colorHex,omitempty" json:"colorHex,omitempty" xml:"colorHex,omitempty"`
	// Price on cents
	Price *int `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// ProductMediaInputRequestBody is used to define fields on request body types.
type ProductMediaInputRequestBody struct {
	// ID of the media record where the resource has being upload
	MediaID *string `form:"mediaId,omitempty" json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// Position on the images of the product
	SortNumber *int `form:"sortNumber,omitempty" json:"sortNumber,omitempty" xml:"sortNumber,omitempty"`
	// Alt text that would show in case the image does not render
	Alt *string `form:"alt,omitempty" json:"alt,omitempty" xml:"alt,omitempty"`
}

// ProductVariantInputRequestBodyRequestBody is used to define fields on
// request body types.
type ProductVariantInputRequestBodyRequestBody struct {
	// Color variant option
	ColorName *string `form:"colorName,omitempty" json:"colorName,omitempty" xml:"colorName,omitempty"`
	// Color in HEX value that would be used on the variant picker
	ColorHex *string `form:"colorHex,omitempty" json:"colorHex,omitempty" xml:"colorHex,omitempty"`
	// Price on cents
	Price *int `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// ProductMediaInputRequestBodyRequestBody is used to define fields on request
// body types.
type ProductMediaInputRequestBodyRequestBody struct {
	// ID of the media record where the resource has being upload
	MediaID *string `form:"mediaId,omitempty" json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// Position on the images of the product
	SortNumber *int `form:"sortNumber,omitempty" json:"sortNumber,omitempty" xml:"sortNumber,omitempty"`
	// Alt text that would show in case the image does not render
	Alt *string `form:"alt,omitempty" json:"alt,omitempty" xml:"alt,omitempty"`
}

// NewListProductResponseBody builds the HTTP response body from the result of
// the "listProduct" endpoint of the "svc-products" service.
func NewListProductResponseBody(res *svcproductsviews.ProductsListView) *ListProductResponseBody {
	body := &ListProductResponseBody{}
	if res.Data != nil {
		body.Data = make([]*ProductResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalSvcproductsviewsProductViewToProductResponseBody(val)
		}
	}
	if res.PageInfo != nil {
		body.PageInfo = marshalSvcproductsviewsPageInfoViewToPageInfoResponseBody(res.PageInfo)
	}
	return body
}

// NewGetProductByIDResponseBody builds the HTTP response body from the result
// of the "getProductById" endpoint of the "svc-products" service.
func NewGetProductByIDResponseBody(res *svcproductsviews.ProductView) *GetProductByIDResponseBody {
	body := &GetProductByIDResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Description: *res.Description,
		Handle:      *res.Handle,
		Status:      string(*res.Status),
	}
	if res.Vendor != nil {
		body.Vendor = marshalSvcproductsviewsVendorViewToVendorResponseBody(res.Vendor)
	}
	if res.Tags != nil {
		body.Tags = make([]string, len(res.Tags))
		for i, val := range res.Tags {
			body.Tags[i] = val
		}
	}
	if res.Variants != nil {
		body.Variants = make([]*ProductVariantResponseBody, len(res.Variants))
		for i, val := range res.Variants {
			body.Variants[i] = marshalSvcproductsviewsProductVariantViewToProductVariantResponseBody(val)
		}
	}
	if res.Medias != nil {
		body.Medias = make([]*ProductMediaResponseBody, len(res.Medias))
		for i, val := range res.Medias {
			body.Medias[i] = marshalSvcproductsviewsProductMediaViewToProductMediaResponseBody(val)
		}
	}
	return body
}

// NewCreateProductResponseBody builds the HTTP response body from the result
// of the "createProduct" endpoint of the "svc-products" service.
func NewCreateProductResponseBody(res *svcproductsviews.ProductView) *CreateProductResponseBody {
	body := &CreateProductResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Description: *res.Description,
		Handle:      *res.Handle,
		Status:      string(*res.Status),
	}
	if res.Vendor != nil {
		body.Vendor = marshalSvcproductsviewsVendorViewToVendorResponseBody(res.Vendor)
	}
	if res.Tags != nil {
		body.Tags = make([]string, len(res.Tags))
		for i, val := range res.Tags {
			body.Tags[i] = val
		}
	}
	if res.Variants != nil {
		body.Variants = make([]*ProductVariantResponseBody, len(res.Variants))
		for i, val := range res.Variants {
			body.Variants[i] = marshalSvcproductsviewsProductVariantViewToProductVariantResponseBody(val)
		}
	}
	if res.Medias != nil {
		body.Medias = make([]*ProductMediaResponseBody, len(res.Medias))
		for i, val := range res.Medias {
			body.Medias[i] = marshalSvcproductsviewsProductMediaViewToProductMediaResponseBody(val)
		}
	}
	return body
}

// NewUpdateProductByIDResponseBody builds the HTTP response body from the
// result of the "updateProductById" endpoint of the "svc-products" service.
func NewUpdateProductByIDResponseBody(res *svcproductsviews.ProductView) *UpdateProductByIDResponseBody {
	body := &UpdateProductByIDResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Description: *res.Description,
		Handle:      *res.Handle,
		Status:      string(*res.Status),
	}
	if res.Vendor != nil {
		body.Vendor = marshalSvcproductsviewsVendorViewToVendorResponseBody(res.Vendor)
	}
	if res.Tags != nil {
		body.Tags = make([]string, len(res.Tags))
		for i, val := range res.Tags {
			body.Tags[i] = val
		}
	}
	if res.Variants != nil {
		body.Variants = make([]*ProductVariantResponseBody, len(res.Variants))
		for i, val := range res.Variants {
			body.Variants[i] = marshalSvcproductsviewsProductVariantViewToProductVariantResponseBody(val)
		}
	}
	if res.Medias != nil {
		body.Medias = make([]*ProductMediaResponseBody, len(res.Medias))
		for i, val := range res.Medias {
			body.Medias[i] = marshalSvcproductsviewsProductMediaViewToProductMediaResponseBody(val)
		}
	}
	return body
}

// NewListProductBadRequestResponseBody builds the HTTP response body from the
// result of the "listProduct" endpoint of the "svc-products" service.
func NewListProductBadRequestResponseBody(res *goa.ServiceError) *ListProductBadRequestResponseBody {
	body := &ListProductBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetProductByIDNotFoundResponseBody builds the HTTP response body from the
// result of the "getProductById" endpoint of the "svc-products" service.
func NewGetProductByIDNotFoundResponseBody(res *goa.ServiceError) *GetProductByIDNotFoundResponseBody {
	body := &GetProductByIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateProductConflictResponseBody builds the HTTP response body from the
// result of the "createProduct" endpoint of the "svc-products" service.
func NewCreateProductConflictResponseBody(res *goa.ServiceError) *CreateProductConflictResponseBody {
	body := &CreateProductConflictResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateProductByIDConflictResponseBody builds the HTTP response body from
// the result of the "updateProductById" endpoint of the "svc-products" service.
func NewUpdateProductByIDConflictResponseBody(res *goa.ServiceError) *UpdateProductByIDConflictResponseBody {
	body := &UpdateProductByIDConflictResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteProductByIDNotFoundResponseBody builds the HTTP response body from
// the result of the "deleteProductById" endpoint of the "svc-products" service.
func NewDeleteProductByIDNotFoundResponseBody(res *goa.ServiceError) *DeleteProductByIDNotFoundResponseBody {
	body := &DeleteProductByIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListProductPayload builds a svc-products service listProduct endpoint
// payload.
func NewListProductPayload(pageSize int, after *int) *svcproducts.ListProductPayload {
	v := &svcproducts.ListProductPayload{}
	v.PageSize = pageSize
	v.After = after

	return v
}

// NewGetProductByIDPayload builds a svc-products service getProductById
// endpoint payload.
func NewGetProductByIDPayload(productID int) *svcproducts.GetProductByIDPayload {
	v := &svcproducts.GetProductByIDPayload{}
	v.ProductID = productID

	return v
}

// NewCreateProductProductInput builds a svc-products service createProduct
// endpoint payload.
func NewCreateProductProductInput(body *CreateProductRequestBody) *svcproducts.ProductInput {
	v := &svcproducts.ProductInput{
		Title:       *body.Title,
		Description: *body.Description,
		Handle:      body.Handle,
		VendorID:    *body.VendorID,
	}
	if body.Status != nil {
		status := svcproducts.ProductStatus(*body.Status)
		v.Status = &status
	}
	v.Tags = make([]string, len(body.Tags))
	for i, val := range body.Tags {
		v.Tags[i] = val
	}
	if body.Variants != nil {
		v.Variants = make([]*svcproducts.ProductVariantInput, len(body.Variants))
		for i, val := range body.Variants {
			v.Variants[i] = unmarshalProductVariantInputRequestBodyToSvcproductsProductVariantInput(val)
		}
	}
	if body.Medias != nil {
		v.Medias = make([]*svcproducts.ProductMediaInput, len(body.Medias))
		for i, val := range body.Medias {
			v.Medias[i] = unmarshalProductMediaInputRequestBodyToSvcproductsProductMediaInput(val)
		}
	}

	return v
}

// NewUpdateProductByIDPayload builds a svc-products service updateProductById
// endpoint payload.
func NewUpdateProductByIDPayload(body *UpdateProductByIDRequestBody, productID int) *svcproducts.UpdateProductByIDPayload {
	v := &svcproducts.ProductInput{
		Title:       *body.Title,
		Description: *body.Description,
		Handle:      body.Handle,
		VendorID:    *body.VendorID,
	}
	if body.Status != nil {
		status := svcproducts.ProductStatus(*body.Status)
		v.Status = &status
	}
	v.Tags = make([]string, len(body.Tags))
	for i, val := range body.Tags {
		v.Tags[i] = val
	}
	if body.Variants != nil {
		v.Variants = make([]*svcproducts.ProductVariantInput, len(body.Variants))
		for i, val := range body.Variants {
			v.Variants[i] = unmarshalProductVariantInputRequestBodyRequestBodyToSvcproductsProductVariantInput(val)
		}
	}
	if body.Medias != nil {
		v.Medias = make([]*svcproducts.ProductMediaInput, len(body.Medias))
		for i, val := range body.Medias {
			v.Medias[i] = unmarshalProductMediaInputRequestBodyRequestBodyToSvcproductsProductMediaInput(val)
		}
	}
	res := &svcproducts.UpdateProductByIDPayload{
		Payload: v,
	}
	res.ProductID = productID

	return res
}

// NewDeleteProductByIDPayload builds a svc-products service deleteProductById
// endpoint payload.
func NewDeleteProductByIDPayload(productID int) *svcproducts.DeleteProductByIDPayload {
	v := &svcproducts.DeleteProductByIDPayload{}
	v.ProductID = productID

	return v
}

// ValidateCreateProductRequestBody runs the validations defined on
// CreateProductRequestBody
func ValidateCreateProductRequestBody(body *CreateProductRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "body"))
	}
	if body.VendorID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vendorId", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "ACTIVE" || *body.Status == "DRAFT") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"ACTIVE", "DRAFT"}))
		}
	}
	for _, e := range body.Variants {
		if e != nil {
			if err2 := ValidateProductVariantInputRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range body.Medias {
		if e != nil {
			if err2 := ValidateProductMediaInputRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateUpdateProductByIDRequestBody runs the validations defined on
// UpdateProductByIdRequestBody
func ValidateUpdateProductByIDRequestBody(body *UpdateProductByIDRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "body"))
	}
	if body.VendorID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vendorId", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "ACTIVE" || *body.Status == "DRAFT") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"ACTIVE", "DRAFT"}))
		}
	}
	for _, e := range body.Variants {
		if e != nil {
			if err2 := ValidateProductVariantInputRequestBodyRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range body.Medias {
		if e != nil {
			if err2 := ValidateProductMediaInputRequestBodyRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateProductVariantInputRequestBody runs the validations defined on
// ProductVariantInputRequestBody
func ValidateProductVariantInputRequestBody(body *ProductVariantInputRequestBody) (err error) {
	if body.ColorName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("colorName", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	return
}

// ValidateProductMediaInputRequestBody runs the validations defined on
// ProductMediaInputRequestBody
func ValidateProductMediaInputRequestBody(body *ProductMediaInputRequestBody) (err error) {
	if body.MediaID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mediaId", "body"))
	}
	if body.SortNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sortNumber", "body"))
	}
	return
}

// ValidateProductVariantInputRequestBodyRequestBody runs the validations
// defined on ProductVariantInputRequestBodyRequestBody
func ValidateProductVariantInputRequestBodyRequestBody(body *ProductVariantInputRequestBodyRequestBody) (err error) {
	if body.ColorName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("colorName", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	return
}

// ValidateProductMediaInputRequestBodyRequestBody runs the validations defined
// on ProductMediaInputRequestBodyRequestBody
func ValidateProductMediaInputRequestBodyRequestBody(body *ProductMediaInputRequestBodyRequestBody) (err error) {
	if body.MediaID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mediaId", "body"))
	}
	if body.SortNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sortNumber", "body"))
	}
	return
}
