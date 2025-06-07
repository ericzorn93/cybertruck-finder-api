package models

import "time"

// Type aliases
type Group string
type Condition string
type Type string
type Units string
type SourceMetro string
type Color string

// Enum values
const (
	Autopilot    Group = "AUTOPILOT"
	Connectivity Group = "CONNECTIVITY"
	SuperCharger Group = "SUPER_CHARGER"
)

const (
	New Condition = "new"
)

const (
	Transportation Type = "transportation"
)

const (
	Miles Units = "miles"
)

const (
	BayArea SourceMetro = "Bay Area"
)

const (
	White Color = "WHITE"
	Grey  Color = "GREY"
)

type CyberTruckInventoryResponse struct {
	Results           Results `json:"results"`
	TotalMatchesFound int64   `json:"totalMatchesFound"`
}

type Results struct {
	Exact              []Truck `json:"exact"`
	Approximate        []Truck `json:"approximate"`
	ApproximateOutside []Truck `json:"approximateOutside"`
}

type Truck struct {
	InTransit                   bool                   `json:"inTransit"`
	CashDetails                 CashDetails            `json:"cashDetails"`
	CountryCode                 string                 `json:"countryCode"`
	CountryHasVehicleAtLocation bool                   `json:"countryHasVehicleAtLocation"`
	CountryOfOrigin             string                 `json:"countryOfOrigin"`
	CurrencyCode                string                 `json:"currencyCode"`
	CurrencyCodes               string                 `json:"currencyCodes"`
	Decor                       interface{}            `json:"decor"`
	DamageDisclosure            bool                   `json:"damageDisclosure"`
	DamageDisclosureStatus      *string                `json:"damageDisclosureStatus"`
	DestinationHandlingFee      int64                  `json:"destinationHandlingFee"`
	Discount                    int64                  `json:"discount"`
	FactoryCode                 string                 `json:"factoryCode"`
	FederalIncentives           FederalIncentives      `json:"federalIncentives"`
	FinplatDetails              FinplatDetails         `json:"finplatDetails"`
	FixedAssets                 bool                   `json:"fixedAssets"`
	FlexibleOptionsData         []FlexibleOptionsDatum `json:"flexibleOptionsData"`
	HasDamagePhotos             bool                   `json:"hasDamagePhotos"`
	Hash                        string                 `json:"hash"`
	Interior                    []Color                `json:"interior"`

	// InspectionDocumentGUID         interface{}            `json:"inspectionDocumentGuid"`
	// InventoryPrice                 int64                  `json:"inventoryPrice"`
	// IsChargingConnectorIncluded    bool                   `json:"isChargingConnectorIncluded"`
	// IsDemo                         bool                   `json:"isDemo"`
	// IsFactoryGated                 bool                   `json:"isFactoryGated"`
	// IsFederalCreditEligible        bool                   `json:"isFederalCreditEligible"`
	// IsInTransit                    bool                   `json:"isInTransit"`
	// IsLegacy                       bool                   `json:"isLegacy"`
	// IsPreProdWithDisclaimer        bool                   `json:"isPreProdWithDisclaimer"`
	// Language                       string                 `json:"language"`
	// Languages                      []string               `json:"languages"`
	// LeaseCountries                 []string               `json:"leaseCountries"`
	// LexiconDefaultOptions          []FlexibleOptionsDatum `json:"lexiconDefaultOptions"`
	// ListingType                    string                 `json:"listingType"`
	// ListingTypes                   string                 `json:"listingTypes"`
	// LoanCountries                  []string               `json:"loanCountries"`
	// Model                          string                 `json:"model"`
	// Odometer                       int64                  `json:"odometer"`
	// OdometerType                   string                 `json:"odometerType"`
	// OptionCodeData                 []OptionCodeDatum      `json:"optionCodeData"`
	// OptionCodeList                 string                 `json:"optionCodeList"`
	// OptionCodeListDisplayOnly      interface{}            `json:"optionCodeListDisplayOnly"`
	// OptionCodePricing              []OptionCodePricing    `json:"optionCodePricing"`
	// OrderFee                       interface{}            `json:"orderFee"`
	// OriginalDeliveryDate           interface{}            `json:"originalDeliveryDate"`
	// OriginalInCustomerGarageDate   interface{}            `json:"originalInCustomerGarageDate"`
	// Paint                          interface{}            `json:"paint"`
	// PreviouslyRegistered           bool                   `json:"previouslyRegistered"`
	// Price                          int64                  `json:"price"`
	// PurchasePrice                  int64                  `json:"purchasePrice"`
	// Roof                           interface{}            `json:"roof"`
	// RegistrationCount              int64                  `json:"registrationCount"`
	// Rewards                        []interface{}          `json:"rewards"`
	// SteeringWheel                  interface{}            `json:"steeringWheel"`
	// Trim                           []string               `json:"trim"`
	// TaxScheme                      interface{}            `json:"taxScheme"`
	// ThirdPartyHistoryURL           interface{}            `json:"thirdPartyHistoryUrl"`
	// TitleStatus                    string                 `json:"titleStatus"`
	// TitleSubtype                   []string               `json:"titleSubtype"`
	// TotalPrice                     int64                  `json:"totalPrice"`
	// TradeInType                    interface{}            `json:"tradeInType"`
	// TransportFees                  TransportFees          `json:"transportFees"`
	// TrimCode                       string                 `json:"trimCode"`
	// TrimName                       string                 `json:"trimName"`
	// TrimVariantCode                string                 `json:"trimVariantCode"`
	// Trt                            int64                  `json:"trt"`
	// Vin                            string                 `json:"vin"`
	// VehicleHistory                 interface{}            `json:"vehicleHistory"`
	// Wheels                         []string               `json:"wheels"`
	// WarrantyBatteryExpDate         time.Time              `json:"warrantyBatteryExpDate"`
	// WarrantyBatteryIsExpired       bool                   `json:"warrantyBatteryIsExpired"`
	// WarrantyBatteryMile            int64                  `json:"warrantyBatteryMile"`
	// WarrantyBatteryYear            int64                  `json:"warrantyBatteryYear"`
	// WarrantyData                   WarrantyData           `json:"warrantyData"`
	// WarrantyDriveUnitExpDate       time.Time              `json:"warrantyDriveUnitExpDate"`
	// WarrantyDriveUnitMile          int64                  `json:"warrantyDriveUnitMile"`
	// WarrantyDriveUnitYear          int64                  `json:"warrantyDriveUnitYear"`
	// WarrantyMile                   int64                  `json:"warrantyMile"`
	// WarrantyVehicleExpDate         time.Time              `json:"warrantyVehicleExpDate"`
	// WarrantyVehicleIsExpired       bool                   `json:"warrantyVehicleIsExpired"`
	// WarrantyYear                   int64                  `json:"warrantyYear"`
	// Year                           int64                  `json:"year"`
	// AlternateCurrency              []interface{}          `json:"alternateCurrency"`
	// UsedVehicleLimitedWarrantyMile int64                  `json:"usedVehicleLimitedWarrantyMile"`
	// UsedVehicleLimitedWarrantyYear int64                  `json:"usedVehicleLimitedWarrantyYear"`
	// OdometerTypeShort              string                 `json:"odometerTypeShort"`
	// DeliveryDateDisplay            bool                   `json:"deliveryDateDisplay"`
	// TransportationFee              int64                  `json:"transportationFee"`
	// OptionCodeSpecs                OptionCodeSpecs        `json:"optionCodeSpecs"`
	// CompositorViewsCustom          CompositorViewsCustom  `json:"compositorViewsCustom"`
	// IsRangeStandard                bool                   `json:"isRangeStandard"`
	// HasMarketingOptions            bool                   `json:"hasMarketingOptions"`
	// FirstRegistrationDate          interface{}            `json:"firstRegistrationDate"`
}

type CashDetails struct {
	Cash Cash `json:"cash"`
}

type Cash struct {
	InventoryPriceWithoutDiscounts interface{} `json:"inventoryPriceWithoutDiscounts"`
}

type CompositorViews struct {
	FrontView    string `json:"frontView"`
	SideView     string `json:"sideView"`
	InteriorView string `json:"interiorView"`
}

type CompositorViewsCustom struct {
	IsProductWithCustomViews bool         `json:"isProductWithCustomViews"`
	ExternalZoom             ExternalZoom `json:"externalZoom"`
	ExternalCrop             ExternalCrop `json:"externalCrop"`
}

type ExternalCrop struct {
	Order  string `json:"order"`
	Search string `json:"search"`
}

type ExternalZoom struct {
	Order  int64 `json:"order"`
	Search int64 `json:"search"`
}

type FederalIncentives struct {
	IsTaxIncentiveEligible bool  `json:"isTaxIncentiveEligible"`
	PriceAfterTaxIncentive int64 `json:"priceAfterTaxIncentive"`
	TaxIncentiveAmount     int64 `json:"taxIncentiveAmount"`
}

type FinplatDetails struct {
	AutoLeaseOperationalLeaseCTPrivate AutoLCTPrivate `json:"autoLeaseOperationalLeaseCTPrivate"`
	AutoLoanLoanCTPrivate              AutoLCTPrivate `json:"autoLoanLoanCTPrivate"`
}

type AutoLCTPrivate struct {
	Calculated Calculated `json:"calculated"`
}

type Calculated struct {
	Inputs  Inputs  `json:"inputs"`
	Outputs Outputs `json:"outputs"`
}

type Inputs struct {
	CashDownPayment int64   `json:"cashDownPayment"`
	DistanceAllowed *int64  `json:"distanceAllowed,omitempty"`
	InterestRate    float64 `json:"interestRate"`
	MonthlyPayment  int64   `json:"monthlyPayment"`
	TermLength      int64   `json:"termLength"`
}

type Outputs struct {
	MonthlyPayment int64 `json:"monthlyPayment"`
}

type FlexibleOptionsDatum struct {
	Code         string  `json:"code"`
	Description  string  `json:"description"`
	Group        *Group  `json:"group,omitempty"`
	LongName     string  `json:"longName"`
	Name         string  `json:"name"`
	Price        *int64  `json:"price,omitempty"`
	LexiconGroup *string `json:"lexiconGroup,omitempty"`
}

type OptionCodeDatum struct {
	AccelerationUnitLong    *string `json:"accelerationUnitLong,omitempty"`
	AccelerationUnitShort   *string `json:"accelerationUnitShort,omitempty"`
	AccelerationValue       *string `json:"accelerationValue,omitempty"`
	Code                    string  `json:"code"`
	Group                   string  `json:"group"`
	Price                   *int64  `json:"price,omitempty"`
	UnitLong                *string `json:"unitLong,omitempty"`
	UnitShort               *string `json:"unitShort,omitempty"`
	Value                   *string `json:"value,omitempty"`
	TopSpeedLabel           *string `json:"topSpeedLabel,omitempty"`
	RangeLabelSource        *string `json:"rangeLabelSource,omitempty"`
	RangeSource             *string `json:"rangeSource,omitempty"`
	RangeSourceInventoryNew *string `json:"rangeSourceInventoryNew,omitempty"`
	Description             *string `json:"description,omitempty"`
	LongName                *string `json:"longName,omitempty"`
	Name                    *string `json:"name,omitempty"`
}

type OptionCodePricing struct {
	Code  string `json:"code"`
	Group string `json:"group"`
	Price int64  `json:"price"`
}

type OptionCodeSpecs struct {
	CSpecs    CCallouts `json:"cSpecs"`
	CDesign   CCallouts `json:"cDesign"`
	CCallouts CCallouts `json:"cCallouts"`
	COpts     CCallouts `json:"cOpts"`
}

type CCallouts struct {
	Code    string                 `json:"code"`
	Name    string                 `json:"name"`
	Options []FlexibleOptionsDatum `json:"options"`
}

type TransportFees struct {
	ExemptVRL    []int64       `json:"exemptVRL"`
	Fees         []Fee         `json:"fees"`
	MetroFees    []MetroFee    `json:"metroFees"`
	TrtToTrtFees []TrtToTrtFee `json:"trtToTrtFees"`
}

type Fee struct {
	Amount    int64     `json:"amount"`
	Condition Condition `json:"condition"`
	Query     Query     `json:"query"`
	Type      Type      `json:"type"`
	Units     Units     `json:"units"`
}

type Query struct {
	Max int64   `json:"max"`
	Min float64 `json:"min"`
}

type MetroFee struct {
	Amount           int64       `json:"amount"`
	DestinationMetro string      `json:"destinationMetro"`
	SourceMetro      SourceMetro `json:"sourceMetro"`
}

type TrtToTrtFee struct {
	Amount        int64   `json:"amount"`
	DestinationID []int64 `json:"destinationId"`
	SourceID      []int64 `json:"sourceId"`
}

type WarrantyData struct {
	UsedVehicleLimitedWarrantyMile int64     `json:"usedVehicleLimitedWarrantyMile"`
	UsedVehicleLimitedWarrantyYear int64     `json:"usedVehicleLimitedWarrantyYear"`
	WarrantyBatteryExpDate         time.Time `json:"warrantyBatteryExpDate"`
	WarrantyBatteryIsExpired       bool      `json:"warrantyBatteryIsExpired"`
	WarrantyBatteryMile            int64     `json:"warrantyBatteryMile"`
	WarrantyBatteryYear            int64     `json:"warrantyBatteryYear"`
	WarrantyDriveUnitExpDate       time.Time `json:"warrantyDriveUnitExpDate"`
	WarrantyDriveUnitMile          int64     `json:"warrantyDriveUnitMile"`
	WarrantyDriveUnitYear          int64     `json:"warrantyDriveUnitYear"`
	WarrantyMile                   int64     `json:"warrantyMile"`
	WarrantyVehicleExpDate         time.Time `json:"warrantyVehicleExpDate"`
	WarrantyVehicleIsExpired       bool      `json:"warrantyVehicleIsExpired"`
	WarrantyYear                   int64     `json:"warrantyYear"`
}
