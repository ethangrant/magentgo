package magentgo

import (
	"context"
	"fmt"
)

type ProductService struct {
	client *Client
}

type ProductsResponse struct {
	Products []ProductResponse `json:"items"`
}

type ProductResponse struct {
	ErrorResponse
	ID                  int    `json:"id"`
	Sku                 string `json:"sku"`
	Name                string `json:"name"`
	AttributeSetID      int    `json:"attribute_set_id"`
	Price               float32    `json:"price"`
	Status              int    `json:"status"`
	Visibility          int    `json:"visibility"`
	TypeID              string `json:"type_id"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	Weight              int    `json:"weight"`
	ExtensionAttributes struct {
		WebsiteIds []int `json:"website_ids"`
		StockItem  struct {
			ItemID                         int  `json:"item_id"`
			ProductID                      int  `json:"product_id"`
			StockID                        int  `json:"stock_id"`
			Qty                            int  `json:"qty"`
			IsInStock                      bool `json:"is_in_stock"`
			IsQtyDecimal                   bool `json:"is_qty_decimal"`
			ShowDefaultNotificationMessage bool `json:"show_default_notification_message"`
			UseConfigMinQty                bool `json:"use_config_min_qty"`
			MinQty                         int  `json:"min_qty"`
			UseConfigMinSaleQty            int  `json:"use_config_min_sale_qty"`
			MinSaleQty                     int  `json:"min_sale_qty"`
			UseConfigMaxSaleQty            bool `json:"use_config_max_sale_qty"`
			MaxSaleQty                     int  `json:"max_sale_qty"`
			UseConfigBackorders            bool `json:"use_config_backorders"`
			Backorders                     int  `json:"backorders"`
			UseConfigNotifyStockQty        bool `json:"use_config_notify_stock_qty"`
			NotifyStockQty                 int  `json:"notify_stock_qty"`
			UseConfigQtyIncrements         bool `json:"use_config_qty_increments"`
			QtyIncrements                  int  `json:"qty_increments"`
			UseConfigEnableQtyInc          bool `json:"use_config_enable_qty_inc"`
			EnableQtyIncrements            bool `json:"enable_qty_increments"`
			UseConfigManageStock           bool `json:"use_config_manage_stock"`
			ManageStock                    bool `json:"manage_stock"`
			LowStockDate                   any  `json:"low_stock_date"`
			IsDecimalDivided               bool `json:"is_decimal_divided"`
			StockStatusChangedAuto         int  `json:"stock_status_changed_auto"`
		} `json:"stock_item"`
		ConfigurableProductOptions []any `json:"configurable_product_options"`
		ConfigurableProductLinks   []any `json:"configurable_product_links"`
	} `json:"extension_attributes"`
	ProductLinks        []any `json:"product_links"`
	Options             []any `json:"options"`
	MediaGalleryEntries []any `json:"media_gallery_entries"`
	TierPrices          []any `json:"tier_prices"`
	CustomAttributes    []struct {
		AttributeCode string `json:"attribute_code"`
		Value         any    `json:"value"`
	} `json:"custom_attributes"`
	SearchCriteria struct {
		FilterGroups []any `json:"filter_groups"`
		PageSize     int   `json:"page_size"`
		CurrentPage  int   `json:"current_page"`
	} `json:"search_criteria"`
	TotalCount int `json:"total_count"`
}

// request single product by sku
func (p *ProductService) GetBySku(sku string, ctx context.Context) (ProductResponse, error) {
	productResponse := &ProductResponse{}
	_, err := p.client.call(fmt.Sprintf("products/%s", sku), "GET", nil, productResponse, ctx)
	if err != nil {
		return *productResponse, err
	}

	return *productResponse, nil
}

// request single product by sku
func (p *ProductService) GetById(id int, ctx context.Context) (ProductResponse, error) {
	productResponse := &ProductResponse{}
	_, err := p.client.call(fmt.Sprintf("products/id/%d", id), "GET", nil, productResponse, ctx)
	if err != nil {
		return *productResponse, err
	}

	return *productResponse, nil
}

// get list of products based on search criteria
func (p *ProductService) GetProducts(searchCriteria string, ctx context.Context) (ProductsResponse, error) {
	productsResponse := &ProductsResponse{}
	_, err := p.client.call(fmt.Sprintf("products?%s", searchCriteria), "GET", nil, productsResponse, ctx)
	if err != nil {
		return *productsResponse, err;
	}

	return *productsResponse, nil
}

// TODO: test this
func (p *ProductService) UpdateProduct(sku string, product ProductResponse, ctx context.Context) (ProductResponse, error) {
	productResponse := &ProductResponse{}
	_, err := p.client.call(fmt.Sprintf("products/%s", sku), "PUT", product, productResponse, ctx)
	if err != nil {
		return *productResponse, err
	}

	return *productResponse, nil
}

// TODO: test this
func (p *ProductService) DeleteProduct(sku string, ctx context.Context) ([]byte, error) {
	res, err := p.client.call(fmt.Sprintf("products/%s", sku), "DELETE", nil, nil, ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}