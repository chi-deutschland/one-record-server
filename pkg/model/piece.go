package model

type Piece struct {
	ID							string							`firestore:"id" json:"id,omitempty"`
	// LogisticsObject
	// ContainedItems				[]Item
	ContainedPieces				[]Piece							`firestore:"containedPieces,omitempty" json:"containedPieces,omitempty"`
	// CustomsInfo					CustomsInfo
	// Dimensions					Dimensions
	// ExternalReferences			ExternalReference				`firestore:"externalRefernces,omitempty" json:"externalRefernces,omitempty"`
	// GrossWeight					Value
	// HandlingInstructions		HandlingInstructions
	// OtherIdentifiers			[]OtherIdentifier
	OtherParty					Company							`firestore:"otherParty,omitempty" json:"otherParty,omitempty"`
	// PackagingType				PackagingType
	// Parties						[]Party
	// ProductionCountry			ProductionCountry
	// SecurityDeclaration			SecurityDeclaration
	// SecurityStatus				SecurityDeclaration
	// ServiceRequest				ServiceRequest
	// Shipment					Shipment						`firestore:"shipment,omitempty" json:"shipment,omitempty"`
	Shipper						Company							`firestore:"shipper,omitempty" json:"shipper,omitempty"`
	// SpecialHandling				SpecialHandling
	// TransportMovements			[]TransportMovement				`firestore:"transportMovements,omitempty" json:"transportMovements,omitempty"`
	// TransportSegments			[]TransportSegment
	// UldReference				ULD
	// VolumetricWeight			VolumetricWeight
	Coload						bool							`firestore:"coload,omitempty" json:"coload,omitempty"`
	DeclaredValueForCarriage	string							`firestore:"declaredValueForCarriage,omitempty" json:"declaredValueForCarriage,omitempty"`
	GoodsDescription			string							`firestore:"goodsDescription,omitempty" json:"goodsDescription,omitempty"`
	// LoadType					
	NvdForCarriage				bool							`firestore:"nvdForCarriage,omitempty" json:"nvdForCarriage,omitempty"`
	NvdForCustoms				bool							`firestore:"nvdForCustoms,omitempty" json:"nvdForCustoms,omitempty"`
	// PackageMarkCoded			
	PackagedeIdentifier			string							`firestore:"packagedeIdentifier,omitempty" json:"packagedeIdentifier,omitempty"`
	ShippingMarks				string							`firestore:"shippingMarks,omitempty" json:"shippingMarks,omitempty"`
	Slac						int								`firestore:"slac,omitempty" json:"slac,omitempty"`
	Stackable					bool							`firestore:"stackable,omitempty" json:"stackable,omitempty"`
	Turnable					bool							`firestore:"turnable,omitempty" json:"turnable,omitempty"`
	Upid						string							`firestore:"upid,omitempty" json:"upid,omitempty"`
}
