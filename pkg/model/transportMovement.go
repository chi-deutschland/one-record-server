package model

type TransportMovement struct {
	ID         string `firestore:"id" json:"id,omitempty"`
	ArrivalLocation			Location	`firestore:"arrivalLocation,omitempty" json:"arrivalLocation,omitempty"`
	// Co2CalculationMethod	CO2CalcMethod 
	// Co2Emissions CO2Emissions
	DepartureLocation			Location	`firestore:"departureLocation,omitempty" json:"departureLocation,omitempty"`
	// DistanceCalculated Value
	// DistanceMeasured Value
	ExternalReferences   []ExternalReference		`firestore:"externalReferences,omitempty" json:"externalReferences,omitempty"`
	// FuelAmountCalculated Value
	// FuelAmountMeasured Value
	// MovementTimes MovementTimes
	// Payload Value
	// TransportMeans TransportMeans
	// TransportMeansOperators TransportMeansOperators
	TransportedPieces		[]Piece		`firestore:"transportedPieces,omitempty" json:"transportedPieces,omitempty"`
	// TransportedUlds ULD
	FuelType		string		`firestore:"fuelType,omitempty" json:"fuelType,omitempty"`
	// ModeCode
	// ModeQualifier
	Seal		string		`firestore:"seal,omitempty" json:"seal,omitempty"`
	TransportIdentifier			string	`firestore:"transportIdentifier,omitempty" json:"transportIdentifier,omitempty"`
	UnplannedStop			string	`firestore:"unplannedStop,omitempty" json:"unplannedStop,omitempty"`
}