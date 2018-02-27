package ch

// TODO(js) Why aren't there tools to generate all this from JSON sample files?

// These are all hand generated from the online documentation, current 2016-09-29
// Not that they have not been extensively tested, and may contain mistakes.

type ErrorResource struct {
	Errors []struct {
		Error        string
		ErrorValues  []map[string]string `json:"error_values"`
		Location     string
		LocationType string
		Type         string
	} `json:"errors"`
}

type CompanyProfileResource struct {
	Accounts struct {
		AccountingReferenceDate struct {
			Day   string // int?
			Month string // int?
		} `json:"accounting_reference_date"`
		LastAccounts struct {
			MadeUpTo string `json:"made_up_to"` // date
			Type     string
		} `json:"last_accounts"`
		NextDue      string `json:"next_due"`        // date
		NextMadeUpTo string `json:"next_made_up_to"` // date
		Overdue      bool
	}
	AnnualReturn struct {
		LastMadeUpTo string `json:"last_made_up_to"` // date
		NextDue      string `json:"next_due"`        // date
		NextMadeUpTo string `json:"next_made_up_to"` // date
		Overdue      bool
	} `json:"annual_return"`
	BranchCompanyDetails struct {
		BusinessActivity    string `json:"business_activity"`
		ParentCompanyName   string `json:"parent_company_name"`
		ParentCompanyNumber string `json:"parent_company_number"`
	} `json:"branch_company_details"`
	CanFile               bool   `json:"can_file"`
	CompanyName           string `json:"company_name"`
	CompanyNumber         string `json:"company_number"`
	CompanyStatus         string `json:"company_status"`
	CompanyStatusDetail   string `json:"company_status_detail"`
	ConfirmationStatement struct {
		LastMadeUpTo string `json:"last_made_up_to"` // date
		NextDue      string `json:"next_due"`        // date
		NextMadeUpTo string `json:"next_made_up_to"` // date
		Overdue      bool
	} `json:"confirmation_statement"`
	DateOfCessation       string `json:"date_of_cessation"` // date
	DateOfCreation        string `json:"date_of_creation"`  // date
	ETag                  string
	ForeignCompanyDetails struct {
		AccountingRequirement struct {
			ForeignAccountType        string `json:"foreign_account_type"`
			TermsOfAccountPublication string `json:"terms_of_account_publication"`
		} `json:"accounting_requirement"`
		Accounts struct {
			AccountPeriodFrom struct {
				Day   int
				Month int
			} `json:"account_period_from"`
			AccountPeriodTo struct {
				Day   int
				Month int
			} `json:"account_period_to"`
			MustFileWithin struct {
				Months int
			} `json:"must_file_within"`
		}
		BusinessActivity            string `json:"business_activity"`
		CompanyType                 string `json:"company_type"`
		GovernedBy                  string `json:"governed_by"`
		IsACreditFinanceInstitution bool   `json:"is_a_credit_finance_institution"`
		OriginatingRegistry         struct {
			Country string
			Name    string
		} `json:"originating_registry"`
		RegistrationNumber string `json:"registration_number"`
	} `json:"foreign_company_details"`
	HasBeenLiquidated          bool `json:"has_been_liquidated"`
	HasCharges                 bool `json:"has_charges"`
	HasInsolvencyHistory       bool `json:"has_insolvency_history"`
	IsCommunityInterestCompany bool `json:"is_community_interest_company"`
	Jurisdiction               string
	LastFullMembersListDate    string `json:"last_full_members_list_date"` // date
	Links                      struct {
		Charges                                 string
		FilingHistory                           string `json:"filing_history"`
		Insolvency                              string
		Officers                                string
		Overseas                                string
		PersonsWithSignificantControl           string `json:"persons_with_significant_control"`
		PersonsWithSignificantControlStatements string `json:"persons_with_significant_control_statements"`
		Registers                               string
		Self                                    string
		UKEstablishments                        string `json:"uk-establishments"`
	}
	PreviousCompanyNames []struct {
		CeasedOn      string `json:"ceased_on"`      // date
		EffectiveFrom string `json:"effective_from"` // date
		Name          string
	} `json:"previous_company_names"`
	RegisteredOfficeAddress struct {
		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		CareOf       string `json:"care_of"`
		Country      string
		Locality     string
		POBox        string `json:"po_box"`
		PostalCode   string `json:"postal_code"`
		Premises     string
		Region       string
	} `json:"registered_office_address"`
	RegisteredOfficeIsInDispute          bool     `json:"registered_office_is_in_dispute"`
	SICCodes                             []string `json:"sic_codes"`
	Type                                 string
	UndeliverableRegisteredOfficeAddress bool `json:"undeliverable_registered_office_address"`
}

type RegisteredOfficeAddressResource struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	// CareOf       string `json:"care_of"`
	Country    string
	ETag       string
	Locality   string
	POBox      string `json:"po_box"`
	PostalCode string `json:"postal_code"`
	Premises   string
	Region     string
}

type OfficerListResource struct {
	ActiveCount   int `json:"active_count"`
	ETag          string
	InactiveCount int `json:"inactive_count"`
	Items         []struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			CareOf       string `json:"care_of"`
			Country      string
			Locality     string
			POBox        string `json:"po_box"`
			PostalCode   string `json:"postal_code"`
			Premises     string
			Region       string
		}
		AppointedOn        string `json:"appointed_on"` // date
		CountryOfResidence string `json:"country_of_residence"`
		DateOfBirth        struct {
			Day   int
			Month int
			Year  int
		} `json:"date_of_birth"`
		FormerNames []struct {
			Forenames string
			Surname   string
		} `json:"former_names"`
		Identification struct {
			IdentificationType string `json:"identification_type"`
			LegalAuthority     string `json:"legal_authority"`
			LegalForm          string `json:"legal_form"`
			PlaceRegistered    string `json:"place_registered"`
			RegistrationNumber string `json:"registration_number"`
		}
		Links struct { // TODO(js) Code smell?
			Officer struct {
				Appointments string
			}
		}
		Name        string
		Nationality string
		Occupation  string
		OfficerRole string `json:"officer_role"`
		ResignedOn  string `json:"resigned_on"` // date
	}
	ItemsPerPage int `json:"items_per_page"`
	Kind         string
	Links        struct {
		Self string
	}
	ResignedCount int `json:"resigned_count"`
	StartIndex    int `json:"start_index"`
	TotalResults  int `json:"total_results"`
}

type FilingHistoryListResource struct {
	ETag                string
	FilingHistoryStatus string `json:"filing_history_status"`
	Items               []FilingHistoryItemResource
	ItemsPerPage        int `json:"items_per_page"`
	Kind                string
	StartIndex          int `json:"start_index"`
	TotalCount          int `json:"total_count"`
}

type FilingHistoryItemResource struct {
	Annotations []struct {
		Annotation  string
		Date        string // date
		Description string
	}
	AssociatedFilings []struct {
		Date        string // date
		Description string
		Type        string
	} `json:"associated_filings"`
	Barcode     string
	Category    string
	Date        string // date
	Description string
	Links       struct {
		DocumentMetadata string `json:"document_metadata"`
		Self             string
	}
	Pages       int
	PaperFiled  bool `json:"paper_filed"`
	Resolutions []struct {
		Category    string
		Description string
		DocumentID  string `json:"document_id"`
		ReceiveDate string `json:"receive_date"` // date
		Subcategory string
		Type        string
	}
	Subcategory   string
	TransactionID string `json:"transaction_id"`
	Type          string
}

type CompanyInsolvencyResource struct {
	Cases []struct {
		Dates []struct {
			Date string // date
			Type string
		}
		Links struct {
			Charge string
		}
		Notes         []string
		Practitioners []struct {
			Address struct {
				AddressLine1 string `json:"address_line_1"`
				AddressLine2 string `json:"address_line_2"`
				Country      string
				Locality     string
				PostalCode   string `json:"postal_code"`
				Region       string
			}
		}
		Type string
	}
	ETag   string
	Status interface{} // TODO(js) In original docs this is `"status" : [ null ]` - what does that even mean?
}

type ChargeListResource struct {
	ETag               string
	Items              []ChargeDetailsResource
	PartSatisfiedCount int `json:"part_satisfied_count"`
	SatisfiedCount     int `json:"satisfied_count"`
	TotalCount         int `json:"total_count"`
	UnfilteredCount    int `json:"unfiltered_count"`
}

type ChargeDetailsResource struct {
	AcquiredOn           string `json:"acquired_on"` // date
	AssetsCeasedReleased string `json:"assets_ceased_released"`
	ChargeCode           string `json:"charge_code"`
	ChargeNumber         int    `json:"charge_number"`
	Classification       struct {
		Description string
		Type        string
	}
	CoveringInstrumentDate string `json:"covering_instrument_date"` // date
	CreatedOn              string `json:"created_on"`               // date
	DeliveredOn            string `json:"delivered_on"`             // date
	ETag                   string
	ID                     string
	InsolvencyCases        []struct {
		CaseNumber int `json:"case_number"`
		Links      struct {
			Case string
		}
		TransactionID int `json:"transaction_id"`
	} `json:"insolvency_cases"`
	Links struct {
		Self string
	}
	MoreThanFourPersonsEntitled bool `json:"more_than_four_persons_entitled"`
	Particulars                 struct {
		ChargorActingAsBareTrustee bool `json:"chargor_acting_as_bare_trustee"`
		ContainsFixedCharge        bool `json:"contains_fixed_charge"`
		ContainsFloatingCharge     bool `json:"contains_floating_charge"`
		ContainsNegativePledge     bool `json:"contains_negative_pledge"`
		Description                string
		FloatingChargeCoversAll    bool `json:"floating_charge_covers_all"`
		Type                       string
	}
	PersonsEntitled []struct {
		Name string
	} `json:"persons_entitled"`
	ResolvedOn          string `json:"resolved_on"`  // date
	SatisfiedOn         string `json:"satisfied_on"` // date
	ScottishAlterations struct {
		HasAlterationsToOrder        bool `json:"has_alterations_to_order"`
		HasAlterationsToProhibitions bool `json:"has_alterations_to_prohibitions"`
		HasRestrictingProvisions     bool `json:"has_restricting_provisions"`
	} `json:"scottish_alterations"`
	SecuredDetails struct {
		Description string
		Type        string
	} `json:"secured_details"`
	Status       string
	Transactions []struct {
		DeliveredOn          string `json:"delivered_on"` // date
		FilingType           string `json:"filing_type"`
		InsolvencyCaseNumber int    `json:"insolvency_case_number"`
		Links                struct {
			Filing         string
			InsolvencyCase string `json:"insolvency_case"`
		}
		TransactionID int `json:"transaction_id"`
	}
}

type AppointmentListResource struct {
	DateOfBirth struct {
		Month int
		Year  int
	} `json:"date_of_birth"`
	ETag               string
	IsCorporateOfficer bool `json:"is_corporate_officer"`
	Items              []struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			CareOf       string `json:"care_of"`
			Country      string
			Locality     string
			POBox        string `json:"po_box"`
			PostalCode   string `json:"postal_code"`
			Premises     string
			Region       string
		}
		AppointedBefore string `json:"appointed_before"` // date
		AppointedOn     string `json:"appointed_on"`     // date
		AppointedTo     struct {
			CompanyName   string `json:"company_name"`
			CompanyNumber string `json:"company_number"`
			CompanyStatus string `json:"company_status"`
		} `json:"appointed_to"`
		CountryOfResidence string `json:"country_of_residence"`
		FormerNames        []struct {
			Forenames string
			Surname   string
		} `json:"former_names"`
		Identification struct {
			IdentificationType string `json:"identification_type"`
			LegalAuthority     string `json:"legal_authority"`
			LegalForm          string `json:"legal_form"`
			PlaceRegistered    string `json:"place_registered"`
			RegistrationNumber string `json:"registration_number"`
		}
		IsPre1992Appointment bool `json:"is_pre_1992_appointment"`
		Links                struct {
			Company string
		}
		Name         string
		NameElements struct {
			Forename       string
			Honours        string
			OtherForenames string `json:"other_forenames"`
			Surname        string
			Title          string
		} `json:"name_elements"`
		Nationality string
		Occupation  string
		OfficerRole string `json:"officer_role"`
		ResignedOn  string `json:"resigned_on"` // date
	}
	ItemsPerPage int `json:"items_per_page"`
	Kind         string
	Links        struct {
		Self string
	}
	Name         string
	StartIndex   int `json:"start_index"`
	TotalResults int `json:"total_results"`
}

type NaturalDisqualificationResource struct {
	DateOfBirth       string `json:"date_of_birth"`
	Disqualifications []struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			Country      string
			Locality     string
			PostalCode   string `json:"postal_code"`
			Premises     string
			Region       string
		}
		CaseIdentifier       string   `json:"case_identifier"`
		CompanyNames         []string `json:"company_names"`
		CourtName            string   `json:"court_name"`
		DisqualificationType string   `json:"disqualification_type"`
		DisqualifiedFrom     string   `json:"disqualified_from"`  // date
		DisqualifiedUntil    string   `json:"disqualified_until"` // date
		HeardOn              string   `json:"heard_on"`           // date
		LastVariation        struct {
			CaseIdentifier string `json:"case_identifier"`
			CourtName      string `json:"court_name"`
			VariedOn       string `json:"varied_on"` // date
		} `json:"last_variation"`
		Reason struct {
			Act           string
			Article       string
			DescriptionID string `json:"description_identifier"`
			Section       string
		}
		UndertakenOn string `json:"undertaken_on"` // date
	}
	ETag     string
	Forename string
	Honours  string
	Kind     string
	Links    struct {
		Self string
	}
	Nationality      string
	OtherForenames   string `json:"other_forenames"`
	PermissionsToAct []struct {
		CompanyNames []string `json:"company_names"`
		CourtName    string   `json:"court_name"`
		ExpiresOn    string   `json:"expires_on"` // date
		GrantedOn    string   `json:"granted_on"` // date

	} `json:"permissions_to_act"`
	Surname string
	Title   string
}

type CorporateDisqualificationResource struct {
	CompanyNumber         string `json:"company_number"`
	CountryOfRegistration string `json:"country_of_registration"`
	Disqualifications     []struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			Country      string
			Locality     string
			PostalCode   string `json:"postal_code"`
			Premises     string
			Region       string
		}
		CaseIdentifier       string   `json:"case_identifier"`
		CompanyNames         []string `json:"company_names"`
		CourtName            string   `json:"court_name"`
		DisqualificationType string   `json:"disqualification_type"`
		DisqualifiedFrom     string   `json:"disqualified_from"`  // date
		DisqualifiedUntil    string   `json:"disqualified_until"` // date
		HeardOn              string   `json:"heard_on"`           // date
		LastVariation        struct {
			CaseIdentifier string `json:"case_identifier"`
			CourtName      string `json:"court_name"`
			VariedOn       string `json:"varied_on"` // date
		} `json:"last_variation"`
		Reason struct {
			Act           string
			Article       string
			DescriptionID string `json:"description_identifier"`
			Section       string
		}
		UndertakenOn string `json:"undertaken_on"` // date
	}
	ETag  string
	Kind  string
	Links struct {
		Self string
	}
	Name             string
	PermissionsToAct []struct {
		CompanyNames []string `json:"company_names"`
		CourtName    string   `json:"court_name"`
		ExpiresOn    string   `json:"expires_on"` // date
		GrantedOn    string   `json:"granted_on"` // date

	} `json:"permissions_to_act"`
}

type CompanyUKEstablishmentsResource struct {
	ETag  string
	Items []struct {
		CompanyName   string `json:"company_name"`
		CompanyNumber string `json:"company_number"`
		CompanyStatus string `json:"company_status"`
		Links         struct {
			Company string
		}
		Locality string
	}
	Kind  string
	Links struct {
		Self string
	}
}

type PSCListResource struct {
	ActiveCount int `json:"active_count"`
	CeasedCount int `json:"ceased_count"`
	ETag        string
	Items       []struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			CareOf       string `json:"care_of"`
			Country      string
			Locality     string
			POBox        string `json:"po_box"`
			PostalCode   string `json:"postal_code"`
			Premises     string
			Region       string
		}
		CeasedOn           string `json:"ceased_on"` // date
		CountryOfResidence string `json:"country_of_residence"`
		DateOfBirth        struct {
			Day   int
			Month int
			Year  int
		} `json:"date_of_birth"`
		ETag string
		// Kind string // diff
		Links struct {
			Self      string
			Statement string
		}
		Name         string
		NameElements struct {
			Forename       string
			OtherForenames string `json:"other_forenames"`
			Surname        string
			Title          string
		} `json:"name_elements"`
		Nationality      string
		NaturesOfControl []string `json:"natures_of_control"`
		NotifiedOn       string   `json:"notified_on"` // date
	}
	ItemsPerPage int `json:"items_per_page"`
	Kind         string
	Links        struct {
		PersonsWithSignificantControlStatementsList string `json:"persons_with_significant_control_statements_list"`
		Self                                        string
	}
	StartIndex   int `json:"start_index"`
	TotalResults int `json:"total_results"`
}

type PSCIndividualResource struct {
	Address struct {
		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		CareOf       string `json:"care_of"`
		Country      string
		Locality     string
		POBox        string `json:"po_box"`
		PostalCode   string `json:"postal_code"`
		Premises     string
		Region       string
	}
	CeasedOn           string `json:"ceased_on"` // date
	CountryOfResidence string `json:"country_of_residence"`
	DateOfBirth        struct {
		Day   int
		Month int
		Year  int
	} `json:"date_of_birth"`
	ETag  string
	Kind  string // diff
	Links struct {
		Self      string
		Statement string
	}
	Name         string
	NameElements struct {
		Forename       string
		OtherForenames string `json:"other_forenames"`
		Surname        string
		Title          string
	} `json:"name_elements"`
	Nationality      string
	NaturesOfControl []string `json:"natures_of_control"`
	NotifiedOn       string   `json:"notified_on"` // date
}

type PSCCorporateEntityResource struct {
	Address struct {
		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		CareOf       string `json:"care_of"`
		Country      string
		Locality     string
		POBox        string `json:"po_box"`
		PostalCode   string `json:"postal_code"`
		Premises     string
		Region       string
	}
	CeasedOn       string `json:"ceased_on"` // date
	ETag           string
	Identification struct {
		CountryRegistered  string `json:"country_registered"`
		LegalAuthority     string `json:"legal_authority"`
		LegalForm          string `json:"legal_form"`
		PlaceRegistered    string `json:"place_registered"`
		RegistrationNumber string `json:"registration_number"`
	}
	Kind  string
	Links struct {
		Self      string
		Statement string
	}
	Name             string
	NaturesOfControl []string `json:"natures_of_control"`
	NotifiedOn       string   `json:"notified_on"` // date
}

type PSCLegalPersonResource struct {
	Address struct {
		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		CareOf       string `json:"care_of"`
		Country      string
		Locality     string
		POBox        string `json:"po_box"`
		PostalCode   string `json:"postal_code"`
		Premises     string
		Region       string
	}
	CeasedOn       string `json:"ceased_on"` // date
	ETag           string
	Identification struct {
		LegalAuthority string `json:"legal_authority"`
		LegalForm      string `json:"legal_form"`
	}
	Kind  string
	Links struct {
		Self      string
		Statement string
	}
	Name             string
	NaturesOfControl []string `json:"natures_of_control"`
	NotifiedOn       string   `json:"notified_on"` // date
}

type PSCStatementListResource struct {
	ActiveCount  int `json:"active_count"`
	CeasedCount  int `json:"ceased_count"`
	ETag         string
	Items        []PSCStatementResource
	ItemsPerPage int `json:"items_per_page"`
	Kind         string
	Links        struct {
		PersonsWithSignificantControlList string `json:"persons_with_significant_control_list"`
		Self                              string
	}
	StartIndex   int `json:"start_index"`
	TotalResults int `json:"total_results"`
}

type PSCStatementResource struct {
	CeasedOn      string `json:"ceased_on"` // date
	ETag          string
	Kind          string
	LinkedPSCName string `json:"linked_psc_name"`
	Links         struct {
		PersonWithSignificantControl string `json:"person_with_significant_control"`
		Self                         string
	}
	NotifiedOn                         string `json:"notified_on"` // date
	RestrictionsNoticeWithdrawalReason string `json:"restrictions_notice_withdrawal_reason"`
	Statement                          string
}

type PSCSuperSecureResource struct {
	Ceased      bool
	Description string
	ETag        string
	Kind        string
	Links       struct {
		Self string
	}
}

type CompanyRegisterResource struct {
	CompanyNumber string `json:"company_number"`
	ETag          string
	Kind          string
	Links         struct {
		Self string
	}
	Registers struct {
		Directors struct {
			Items []struct {
				Links struct {
					Filing string
				}
				MovedOn         string `json:"moved_on"` // date
				RegisterMovedTo string `json:"register_moved_to"`
			}
			Links struct {
				DirectorsRegister string `json:"directors_register"`
			}
			RegisterType string `json:"register_type"`
		}
		LLPMembers struct {
			Items []struct {
				Links struct {
					Filing string
				}
				MovedOn         string `json:"moved_on"` // date
				RegisterMovedTo string `json:"register_moved_to"`
			}
			Links struct {
				LLPMembers string `json:"llp_members"`
			}
			RegisterType string `json:"register_type"`
		} `json:"llp_members"`
		LLPUsualResidentialAddress struct {
			Items []struct {
				Links struct {
					Filing string
				}
				MovedOn         string `json:"moved_on"` // date
				RegisterMovedTo string `json:"register_moved_to"`
			}
			Links struct {
				LLPUsualResidentialAddress string `json:"llp_usual_residential_address"`
			}
			RegisterType string `json:"register_type"`
		} `json:"llp_usual_residential_address"`
		Members struct {
			Items []struct {
				Links struct {
					Filing string
				}
				MovedOn         string `json:"moved_on"` // date
				RegisterMovedTo string `json:"register_moved_to"`
			}
			Links struct {
				Members string
			}
			RegisterType string `json:"register_type"`
		}
		PersonsWithSignificantControl struct {
			Items []struct {
				Links struct {
					Filing string
				}
				MovedOn         string `json:"moved_on"` // date
				RegisterMovedTo string `json:"register_moved_to"`
			}
			Links struct {
				PersonsWithSignificantControlRegister string `json:"persons_with_significant_control_register"`
			}
			RegisterType string `json:"register_type"`
		} `json:"persons_with_significant_control"`
		Secretaries struct {
			Items []struct {
				Links struct {
					Filing string
				}
				MovedOn         string `json:"moved_on"` // date
				RegisterMovedTo string `json:"register_moved_to"`
			}
			Links struct {
				SecretariesRegister string `json:"secretaries_register"`
			}
			RegisterType string `json:"register_type"`
		}
		UsualResidentialAddress struct {
			Items []struct {
				Links struct {
					Filing string
				}
				MovedOn         string `json:"moved_on"` // date
				RegisterMovedTo string `json:"register_moved_to"`
			}
			Links struct {
				UsualResidentialAddress string `json:"usual_residential_address"`
			}
			RegisterType string `json:"register_type"`
		} `json:"usual_residential_address"`
	}
}

type CompanyExemptionsResource struct {
	Exemptions struct {
		DisclosureTransparencyRulesChapterFiveApplies struct {
			ExemptionType string `json:"exemption_type"`
			Items         []struct {
				ExemptFrom string `json:"exempt_from"` // date
				ExemptTo   string `json:"exempt_to"`   // date
			}
		} `json:"disclosure_transparency_rules_chapter_five_applies"`
		PSCExemptAsSharesAdmittedOnMarket struct {
			ExemptionType string `json:"exemption_type"`
			Items         []struct {
				ExemptFrom string `json:"exempt_from"` // date
				ExemptTo   string `json:"exempt_to"`   // date
			}
		} `json:"psc_exempt_as_shares_admitted_on_market"`
		PSCExemptAsTradingOnRegulatedMarket struct {
			ExemptionType string `json:"exemption_type"`
			Items         []struct {
				ExemptFrom string `json:"exempt_from"` // date
				ExemptTo   string `json:"exempt_to"`   // date
			}
		} `json:"psc_exempt_as_trading_on_regulated_market"`
	}
	Kind  string
	Links struct {
		Self string
	}
}

type DocumentMetadataResource struct {
	CreatedAt string `json:"created_at"` // date
	ETag      string
	ID        string
	Links     struct {
		Document string
		Self     string
	}
	Pages     int // number
	Resources map[string]ResourceMetadata
	UpdatedAt string `json:"updated_at"` // date
}

type ResourceMetadata struct {
	ContentLength int    `json:"content_length"` // number
	CreatedAt     string `json:"created_at"`     // date
	UpdatedAt     string `json:"updated_at"`     // date
}

// Search structures.

type kind struct {
	Kind string
}

// TODO(js) Looks like SearchResource's Items actually include a mixture of types :/ - look into this.

type SearchResource struct {
	ETag  string
	Items []struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			CareOf       string `json:"care_of"`
			Country      string
			Locality     string
			POBox        string `json:"po_box"`
			PostalCode   string `json:"postal_code"`
			Region       string
		}
		AddressSnippet string `json:"address_snippet"`
		Description    string
		DescriptionIDs []string `json:"description_identifiers"`
		Kind           string
		Links          struct {
			Self string
		}
		Matches struct {
			AddressSnippet []int `json:"address_snippet"`
			Snippet        []int
			Title          []int
		}
		Snippet string
		Title   string
	}
	ItemsPerPage int `json:"items_per_page"`
	Kind         string
	StartIndex   int `json:"start_index"`
	TotalResults int `json:"total_results"`
	// PageNumber   int `json:"page_number"` // Undocumented.
}

type CompanySearchResource struct {
	ETag  string
	Items []struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			CareOf       string `json:"care_of"`
			Country      string
			Locality     string
			POBox        string `json:"po_box"`
			PostalCode   string `json:"postal_code"`
			Region       string
		}
		AddressSnippet  string `json:"address_snippet"`
		CompanyNumber   string `json:"company_number"`
		CompanyStatus   string `json:"company_status"`    // enum
		CompanyType     string `json:"company_type"`      // enum
		DateOfCessation string `json:"date_of_cessation"` // date
		DateOfCreation  string `json:"date_of_creation"`  // date
		Description     string
		DescriptionIDs  []string `json:"description_identifier"` // Singular
		Kind            string
		Links           struct {
			Self string
		}
		Matches struct {
			AddressSnippet []int `json:"address_snippet"`
			Snippet        []int
			Title          []int
		}
		Snippet string
		Title   string
	}
	ItemsPerPage int `json:"items_per_page"`
	Kind         string
	StartIndex   int `json:"start_index"`
	TotalResults int `json:"total_results"`
}

type OfficerSearchResource struct {
	Items []struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			CareOf       string `json:"care_of"`
			Country      string
			Locality     string
			POBox        string `json:"po_box"`
			PostalCode   string `json:"postal_code"`
			Premises     string // diff?
			Region       string
		}
		AddressSnippet   string `json:"address_snippet"`
		AppointmentCount int    `json:"appointment_count"`
		DateOfBirth      struct {
			Month int
			Year  int
		} `json:"date_of_birth"`
		Description    string
		DescriptionIDs []string `json:"description_identifiers"`
		Kind           string
		Links          struct {
			Self string
		}
		Matches struct {
			AddressSnippet []int `json:"address_snippet"`
			Snippet        []int
			Title          []int
		}
		Snippet string
		Title   string
	}
	ItemsPerPage int `json:"items_per_page"`
	Kind         string
	StartIndex   int `json:"start_index"`
	TotalResults int `json:"total_results"`
}

// TODO(js) Disqualified Officers are untested - data plz?

type DisqualifiedOfficerSearchResource struct {
	Items []struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			// CareOf       string `json:"care_of"`
			Country  string
			Locality string
			// POBox        string `json:"po_box"`
			PostalCode string `json:"postal_code"`
			Premises   string // diff?
			Region     string
		}
		AddressSnippet string `json:"address_snippet"`
		// AppointmentCount int    `json:"appointment_count"`
		DateOfBirth    string `json:"date_of_birth"` // date
		Description    string
		DescriptionIDs []string `json:"description_identifiers"`
		Kind           string
		Links          struct {
			Self string
		}
		Matches struct {
			AddressSnippet []int `json:"address_snippet"`
			Snippet        []int
			Title          []int
		}
		Snippet string
		Title   string
	}
	ItemsPerPage int `json:"items_per_page"`
	Kind         string
	StartIndex   int `json:"start_index"`
	TotalResults int `json:"total_results"`
}
