// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fasten-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirCoverage struct {
	models.ResourceBase
	// Covered party
	// https://hl7.org/fhir/r4/search.html#reference
	Beneficiary datatypes.JSON `gorm:"column:beneficiary;type:text;serializer:json" json:"beneficiary,omitempty"`
	// Coverage class (eg. plan, group)
	// https://hl7.org/fhir/r4/search.html#token
	ClassType datatypes.JSON `gorm:"column:classType;type:text;serializer:json" json:"classType,omitempty"`
	// Value of the class (eg. Plan number, group number)
	// https://hl7.org/fhir/r4/search.html#string
	ClassValue datatypes.JSON `gorm:"column:classValue;type:text;serializer:json" json:"classValue,omitempty"`
	// Dependent number
	// https://hl7.org/fhir/r4/search.html#string
	Dependent datatypes.JSON `gorm:"column:dependent;type:text;serializer:json" json:"dependent,omitempty"`
	// The primary identifier of the insured and the coverage
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	MetaLastUpdated *time.Time `gorm:"column:metaLastUpdated;type:datetime" json:"metaLastUpdated,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	MetaProfile datatypes.JSON `gorm:"column:metaProfile;type:text;serializer:json" json:"metaProfile,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	MetaTag datatypes.JSON `gorm:"column:metaTag;type:text;serializer:json" json:"metaTag,omitempty"`
	// Tags applied to this resource
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	MetaVersionId string `gorm:"column:metaVersionId;type:text" json:"metaVersionId,omitempty"`
	// The identity of the insurer or party paying for services
	// https://hl7.org/fhir/r4/search.html#reference
	Payor datatypes.JSON `gorm:"column:payor;type:text;serializer:json" json:"payor,omitempty"`
	// Reference to the policyholder
	// https://hl7.org/fhir/r4/search.html#reference
	PolicyHolder datatypes.JSON `gorm:"column:policyHolder;type:text;serializer:json" json:"policyHolder,omitempty"`
	// The status of the Coverage
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Reference to the subscriber
	// https://hl7.org/fhir/r4/search.html#reference
	Subscriber datatypes.JSON `gorm:"column:subscriber;type:text;serializer:json" json:"subscriber,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// The kind of coverage (health plan, auto, Workers Compensation)
	// https://hl7.org/fhir/r4/search.html#token
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirCoverage) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"beneficiary":          "reference",
		"classType":            "token",
		"classValue":           "string",
		"dependent":            "string",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"payor":                "reference",
		"policyHolder":         "reference",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"subscriber":           "reference",
		"text":                 "string",
		"type":                 "token",
	}
	return searchParameters
}
func (s *FhirCoverage) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
	s.ResourceRaw = datatypes.JSON(resourceRaw)
	// unmarshal the raw resource (bytes) into a map
	var resourceRawMap map[string]interface{}
	err := json.Unmarshal(resourceRaw, &resourceRawMap)
	if err != nil {
		return err
	}
	if len(fhirPathJs) == 0 {
		return fmt.Errorf("fhirPathJs script is empty")
	}
	vm := goja.New()
	// setup the global window object
	vm.Set("window", vm.NewObject())
	// set the global FHIR Resource object
	vm.Set("fhirResource", resourceRawMap)
	// compile the fhirpath library
	fhirPathJsProgram, err := goja.Compile("fhirpath.min.js", fhirPathJs, true)
	if err != nil {
		return err
	}
	// compile the searchParametersExtractor library
	searchParametersExtractorJsProgram, err := goja.Compile("searchParameterExtractor.js", searchParameterExtractorJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// add the searchParametersExtractor library in the goja vm
	_, err = vm.RunProgram(searchParametersExtractorJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Beneficiary
	beneficiaryResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Coverage.beneficiary')")
	if err == nil && beneficiaryResult.String() != "undefined" {
		s.Beneficiary = []byte(beneficiaryResult.String())
	}
	// extracting ClassType
	classTypeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Coverage.class.type')")
	if err == nil && classTypeResult.String() != "undefined" {
		s.ClassType = []byte(classTypeResult.String())
	}
	// extracting ClassValue
	classValueResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Coverage.class.value')")
	if err == nil && classValueResult.String() != "undefined" {
		s.ClassValue = []byte(classValueResult.String())
	}
	// extracting Dependent
	dependentResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Coverage.dependent')")
	if err == nil && dependentResult.String() != "undefined" {
		s.Dependent = []byte(dependentResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Coverage.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting MetaLastUpdated
	metaLastUpdatedResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'meta.lastUpdated')")
	if err == nil && metaLastUpdatedResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006-01-02", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006-01", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		}
	}
	// extracting MetaProfile
	metaProfileResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'meta.profile')")
	if err == nil && metaProfileResult.String() != "undefined" {
		s.MetaProfile = []byte(metaProfileResult.String())
	}
	// extracting MetaTag
	metaTagResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'meta.tag')")
	if err == nil && metaTagResult.String() != "undefined" {
		s.MetaTag = []byte(metaTagResult.String())
	}
	// extracting MetaVersionId
	metaVersionIdResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'meta.versionId')")
	if err == nil && metaVersionIdResult.String() != "undefined" {
		s.MetaVersionId = metaVersionIdResult.String()
	}
	// extracting Payor
	payorResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Coverage.payor')")
	if err == nil && payorResult.String() != "undefined" {
		s.Payor = []byte(payorResult.String())
	}
	// extracting PolicyHolder
	policyHolderResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Coverage.policyHolder')")
	if err == nil && policyHolderResult.String() != "undefined" {
		s.PolicyHolder = []byte(policyHolderResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Coverage.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subscriber
	subscriberResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Coverage.subscriber')")
	if err == nil && subscriberResult.String() != "undefined" {
		s.Subscriber = []byte(subscriberResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = []byte(textResult.String())
	}
	// extracting Type
	typeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Coverage.type')")
	if err == nil && typeResult.String() != "undefined" {
		s.Type = []byte(typeResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirCoverage) TableName() string {
	return "fhir_coverage"
}
