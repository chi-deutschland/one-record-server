package model

import (
	"time"
)

type Shipment struct {
	ID         string `firestore:"id" json:"id,omitempty"`
	ContainedPieces			[]Piece						`firestore:"containedPieces,omitempty" json:"containedPieces,omitempty"`
	// DeliveryLocation		Location
	// Dimensions				Dimensions
	ExternalReferences		[]ExternalReference			`firestore:"externalReferences,omitempty" json:"externalReferences,omitempty"`
	FreightForwarder		Company						`firestore:"freightForwarder,omitempty" json:"freightForwarder,omitempty"`
	// Insurance				Insurance
	// Parties					Party
	Shipper					Company						`firestore:"shipper,omitempty" json:"shipper,omitempty"`
	// TotalGrossWeight		Value
	// VolumetricWeight		VolumetricWeight
	// WaybillNumber			Waybill
	DeliveryDate   			time.Time					`firestore:"deliveryDate,omitempty" json:"deliveryDate,omitempty"`
	GoodsDescription		string						`firestore:"goodsDescription,omitempty" json:"goodsDescription,omitempty"`
	Incoterms				string						`firestore:"incoterms,omitempty" json:"incoterms,omitempty"`
	// OtherChargesIndicator	
	TotalPieceCount			int							`firestore:"totalPieceCount,omitempty" json:"totalPieceCount,omitempty"`
	TotalSLAC				int							`firestore:"totalSLAC,omitempty" json:"totalSLAC,omitempty"`
	// WeightValuationIndicator
}
