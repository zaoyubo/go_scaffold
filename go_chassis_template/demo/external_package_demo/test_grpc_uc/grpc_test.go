package test_grpc_uc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
)

const (
	//address = "10.103.102.73:8080"
	address = "127.0.0.1:8080"
)

func TestGrpcInvoker(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Println(time.Now())
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		fmt.Println(ctx.Deadline())
		defer cancel()
		reqStr := "{\"req_header\":{\"request_id\":\"-1f375645-f064-4155-96a2-b98c91daa238-vs24E\",\"account\":\"LFS\",\"token\":\"token\",\"timestamp\":1656408207,\"caller_ip\":\"\"},\"business_party\":1,\"business_situation\":1,\"number\":\"BR226412212986FT\",\"number_type\":1,\"charge_request_time\":1656408207,\"charge_product_id\":91003,\"level_order_actual_weight\":0,\"level_order_volumetric_weight\":0,\"level_order_height\":0,\"level_order_width\":0,\"level_order_length\":0,\"level_cod_amount\":0,\"level_cogs_amount\":10.010000228881836,\"level_wms_flag\":0,\"level_dg_flag\":1,\"level_pickup_location_ids\":[11007634,11016127],\"level_delivery_location_ids\":[11008793,11017299],\"level_pickup_postcode\":\"1001000\",\"level_deliver_postcode\":\"6009999\",\"level_pickup_longitude\":\"\",\"level_pickup_latitude\":\"\",\"level_deliver_longitude\":\"-40.271915\",\"level_deliver_latitude\":\"-20.259600\",\"level_fallback_flag\":0,\"level_pickup_tag\":1,\"level_estimated_shipping_fee\":8.279999732971191,\"level_three_pl_asf_shopee\":0,\"level_third_party_pushed_shipping_fee\":6000,\"level_third_party_charged_weight\":900,\"sku_infos\":[{\"item_id\":12003093381,\"category_id\":31754,\"weight\":100,\"quantity\":1,\"length\":5,\"width\":20,\"height\":5,\"item_price_usd\":258.8900146484375,\"item_price\":1390}],\"invoice_number\":\"1\",\"invoice_serial_number\":\"2\",\"invoice_access_key\":\"35220224213924000134550020000000011012227262\",\"invoice_issue_time\":1644238466,\"invoice_total_value\":10.010000228881836,\"invoice_seller_registered_name\":\"\",\"invoice_seller_legal_entity_name\":\"VBF TEcnologia e Internet\",\"invoice_buyer_tax_number\":\"82093040700\",\"invoice_seller_tax_number\":\"24213924000134\",\"invoice_seller_state_registration\":\"260200062112\",\"mile_fm_line_id\":\"LBR45\",\"mile_lm_line_id\":\"LBR25\",\"mile_xd_line_id\":\"SBR25\",\"mile_type\":1,\"object_buyer_name\":\"Cleide Occhi\",\"object_buyer_phone\":\"5527999019077\",\"object_seller_name\":\"Sheik Importações\",\"object_seller_address\":\"R Vitor Hugo, 233, BOX 17\",\"object_buyer_address\":\"R Domingos P Lemos, 60, Ed. Rio de Conta apt 301\",\"object_shop_type_id_group\":[0],\"buyer_paid_shipping_fee\":20,\"direction\":1,\"xd_inbound_id\":\"SBR25-AP001\",\"xd_outbound_id\":\"SBR25-AP001\",\"business_status\":\"1001\",\"country\":\"BR\",\"seller_rts_fee_flag\":0}"
		req := &ChargeOrderReq{}
		err := json.Unmarshal([]byte(reqStr), req)
		if err != nil {
			panic(err)
		}

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		c := NewChargeOrderApiServiceClient(conn)

		r, err := c.CreateChargeOrder(ctx, req)
		if err != nil {
			fmt.Printf("could not greet: %v, r:%v\n", err, r)
		}
	}
}
