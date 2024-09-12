// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Me.outputs;

import com.ovhcloud.ovh.Me.outputs.GetMeCurrency;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMeResult {
    /**
     * @return The resource URN of the account, to be used when writing IAM policies
     * 
     */
    private String AccountURN;
    /**
     * @return Postal address of the account
     * 
     */
    private String address;
    /**
     * @return Area of the account
     * 
     */
    private String area;
    /**
     * @return City of birth
     * 
     */
    private String birthCity;
    /**
     * @return Birth date
     * 
     */
    private String birthDay;
    /**
     * @return City of the account
     * 
     */
    private String city;
    /**
     * @return This is the national identification number of the company that possess this account
     * 
     */
    private String companyNationalIdentificationNumber;
    /**
     * @return Type of corporation
     * 
     */
    private String corporationType;
    /**
     * @return Country of the account
     * 
     */
    private String country;
    private List<GetMeCurrency> currencies;
    /**
     * @return The customer code of this account (a numerical value used for identification when contacting support via phone call)
     * 
     */
    private String customerCode;
    /**
     * @return Email address
     * 
     */
    private String email;
    /**
     * @return Fax number
     * 
     */
    private String fax;
    /**
     * @return First name
     * 
     */
    private String firstname;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Italian SDI
     * 
     */
    private String italianSdi;
    /**
     * @return Preferred language for this account
     * 
     */
    private String language;
    /**
     * @return Legal form of the account
     * 
     */
    private String legalform;
    /**
     * @return Name of the account holder
     * 
     */
    private String name;
    /**
     * @return National Identification Number of this account
     * 
     */
    private String nationalIdentificationNumber;
    /**
     * @return Nic handle / customer identifier
     * 
     */
    private String nichandle;
    /**
     * @return Name of the organisation for this account
     * 
     */
    private String organisation;
    /**
     * @return OVHcloud subsidiary
     * 
     */
    private String ovhCompany;
    /**
     * @return OVHcloud subsidiary
     * 
     */
    private String ovhSubsidiary;
    /**
     * @return Phone number
     * 
     */
    private String phone;
    /**
     * @return Country code of the phone number
     * 
     */
    private String phoneCountry;
    /**
     * @return Gender of the account holder
     * 
     */
    private String sex;
    /**
     * @return Backup email address
     * 
     */
    private String spareEmail;
    /**
     * @return State of the postal address
     * 
     */
    private String state;
    /**
     * @return VAT number
     * 
     */
    private String vat;
    /**
     * @return Zipcode of the address
     * 
     */
    private String zip;

    private GetMeResult() {}
    /**
     * @return The resource URN of the account, to be used when writing IAM policies
     * 
     */
    public String AccountURN() {
        return this.AccountURN;
    }
    /**
     * @return Postal address of the account
     * 
     */
    public String address() {
        return this.address;
    }
    /**
     * @return Area of the account
     * 
     */
    public String area() {
        return this.area;
    }
    /**
     * @return City of birth
     * 
     */
    public String birthCity() {
        return this.birthCity;
    }
    /**
     * @return Birth date
     * 
     */
    public String birthDay() {
        return this.birthDay;
    }
    /**
     * @return City of the account
     * 
     */
    public String city() {
        return this.city;
    }
    /**
     * @return This is the national identification number of the company that possess this account
     * 
     */
    public String companyNationalIdentificationNumber() {
        return this.companyNationalIdentificationNumber;
    }
    /**
     * @return Type of corporation
     * 
     */
    public String corporationType() {
        return this.corporationType;
    }
    /**
     * @return Country of the account
     * 
     */
    public String country() {
        return this.country;
    }
    public List<GetMeCurrency> currencies() {
        return this.currencies;
    }
    /**
     * @return The customer code of this account (a numerical value used for identification when contacting support via phone call)
     * 
     */
    public String customerCode() {
        return this.customerCode;
    }
    /**
     * @return Email address
     * 
     */
    public String email() {
        return this.email;
    }
    /**
     * @return Fax number
     * 
     */
    public String fax() {
        return this.fax;
    }
    /**
     * @return First name
     * 
     */
    public String firstname() {
        return this.firstname;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Italian SDI
     * 
     */
    public String italianSdi() {
        return this.italianSdi;
    }
    /**
     * @return Preferred language for this account
     * 
     */
    public String language() {
        return this.language;
    }
    /**
     * @return Legal form of the account
     * 
     */
    public String legalform() {
        return this.legalform;
    }
    /**
     * @return Name of the account holder
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return National Identification Number of this account
     * 
     */
    public String nationalIdentificationNumber() {
        return this.nationalIdentificationNumber;
    }
    /**
     * @return Nic handle / customer identifier
     * 
     */
    public String nichandle() {
        return this.nichandle;
    }
    /**
     * @return Name of the organisation for this account
     * 
     */
    public String organisation() {
        return this.organisation;
    }
    /**
     * @return OVHcloud subsidiary
     * 
     */
    public String ovhCompany() {
        return this.ovhCompany;
    }
    /**
     * @return OVHcloud subsidiary
     * 
     */
    public String ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    /**
     * @return Phone number
     * 
     */
    public String phone() {
        return this.phone;
    }
    /**
     * @return Country code of the phone number
     * 
     */
    public String phoneCountry() {
        return this.phoneCountry;
    }
    /**
     * @return Gender of the account holder
     * 
     */
    public String sex() {
        return this.sex;
    }
    /**
     * @return Backup email address
     * 
     */
    public String spareEmail() {
        return this.spareEmail;
    }
    /**
     * @return State of the postal address
     * 
     */
    public String state() {
        return this.state;
    }
    /**
     * @return VAT number
     * 
     */
    public String vat() {
        return this.vat;
    }
    /**
     * @return Zipcode of the address
     * 
     */
    public String zip() {
        return this.zip;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String AccountURN;
        private String address;
        private String area;
        private String birthCity;
        private String birthDay;
        private String city;
        private String companyNationalIdentificationNumber;
        private String corporationType;
        private String country;
        private List<GetMeCurrency> currencies;
        private String customerCode;
        private String email;
        private String fax;
        private String firstname;
        private String id;
        private String italianSdi;
        private String language;
        private String legalform;
        private String name;
        private String nationalIdentificationNumber;
        private String nichandle;
        private String organisation;
        private String ovhCompany;
        private String ovhSubsidiary;
        private String phone;
        private String phoneCountry;
        private String sex;
        private String spareEmail;
        private String state;
        private String vat;
        private String zip;
        public Builder() {}
        public Builder(GetMeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.AccountURN = defaults.AccountURN;
    	      this.address = defaults.address;
    	      this.area = defaults.area;
    	      this.birthCity = defaults.birthCity;
    	      this.birthDay = defaults.birthDay;
    	      this.city = defaults.city;
    	      this.companyNationalIdentificationNumber = defaults.companyNationalIdentificationNumber;
    	      this.corporationType = defaults.corporationType;
    	      this.country = defaults.country;
    	      this.currencies = defaults.currencies;
    	      this.customerCode = defaults.customerCode;
    	      this.email = defaults.email;
    	      this.fax = defaults.fax;
    	      this.firstname = defaults.firstname;
    	      this.id = defaults.id;
    	      this.italianSdi = defaults.italianSdi;
    	      this.language = defaults.language;
    	      this.legalform = defaults.legalform;
    	      this.name = defaults.name;
    	      this.nationalIdentificationNumber = defaults.nationalIdentificationNumber;
    	      this.nichandle = defaults.nichandle;
    	      this.organisation = defaults.organisation;
    	      this.ovhCompany = defaults.ovhCompany;
    	      this.ovhSubsidiary = defaults.ovhSubsidiary;
    	      this.phone = defaults.phone;
    	      this.phoneCountry = defaults.phoneCountry;
    	      this.sex = defaults.sex;
    	      this.spareEmail = defaults.spareEmail;
    	      this.state = defaults.state;
    	      this.vat = defaults.vat;
    	      this.zip = defaults.zip;
        }

        @CustomType.Setter
        public Builder AccountURN(String AccountURN) {
            if (AccountURN == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "AccountURN");
            }
            this.AccountURN = AccountURN;
            return this;
        }
        @CustomType.Setter
        public Builder address(String address) {
            if (address == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "address");
            }
            this.address = address;
            return this;
        }
        @CustomType.Setter
        public Builder area(String area) {
            if (area == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "area");
            }
            this.area = area;
            return this;
        }
        @CustomType.Setter
        public Builder birthCity(String birthCity) {
            if (birthCity == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "birthCity");
            }
            this.birthCity = birthCity;
            return this;
        }
        @CustomType.Setter
        public Builder birthDay(String birthDay) {
            if (birthDay == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "birthDay");
            }
            this.birthDay = birthDay;
            return this;
        }
        @CustomType.Setter
        public Builder city(String city) {
            if (city == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "city");
            }
            this.city = city;
            return this;
        }
        @CustomType.Setter
        public Builder companyNationalIdentificationNumber(String companyNationalIdentificationNumber) {
            if (companyNationalIdentificationNumber == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "companyNationalIdentificationNumber");
            }
            this.companyNationalIdentificationNumber = companyNationalIdentificationNumber;
            return this;
        }
        @CustomType.Setter
        public Builder corporationType(String corporationType) {
            if (corporationType == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "corporationType");
            }
            this.corporationType = corporationType;
            return this;
        }
        @CustomType.Setter
        public Builder country(String country) {
            if (country == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "country");
            }
            this.country = country;
            return this;
        }
        @CustomType.Setter
        public Builder currencies(List<GetMeCurrency> currencies) {
            if (currencies == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "currencies");
            }
            this.currencies = currencies;
            return this;
        }
        public Builder currencies(GetMeCurrency... currencies) {
            return currencies(List.of(currencies));
        }
        @CustomType.Setter
        public Builder customerCode(String customerCode) {
            if (customerCode == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "customerCode");
            }
            this.customerCode = customerCode;
            return this;
        }
        @CustomType.Setter
        public Builder email(String email) {
            if (email == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "email");
            }
            this.email = email;
            return this;
        }
        @CustomType.Setter
        public Builder fax(String fax) {
            if (fax == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "fax");
            }
            this.fax = fax;
            return this;
        }
        @CustomType.Setter
        public Builder firstname(String firstname) {
            if (firstname == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "firstname");
            }
            this.firstname = firstname;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder italianSdi(String italianSdi) {
            if (italianSdi == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "italianSdi");
            }
            this.italianSdi = italianSdi;
            return this;
        }
        @CustomType.Setter
        public Builder language(String language) {
            if (language == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "language");
            }
            this.language = language;
            return this;
        }
        @CustomType.Setter
        public Builder legalform(String legalform) {
            if (legalform == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "legalform");
            }
            this.legalform = legalform;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nationalIdentificationNumber(String nationalIdentificationNumber) {
            if (nationalIdentificationNumber == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "nationalIdentificationNumber");
            }
            this.nationalIdentificationNumber = nationalIdentificationNumber;
            return this;
        }
        @CustomType.Setter
        public Builder nichandle(String nichandle) {
            if (nichandle == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "nichandle");
            }
            this.nichandle = nichandle;
            return this;
        }
        @CustomType.Setter
        public Builder organisation(String organisation) {
            if (organisation == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "organisation");
            }
            this.organisation = organisation;
            return this;
        }
        @CustomType.Setter
        public Builder ovhCompany(String ovhCompany) {
            if (ovhCompany == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "ovhCompany");
            }
            this.ovhCompany = ovhCompany;
            return this;
        }
        @CustomType.Setter
        public Builder ovhSubsidiary(String ovhSubsidiary) {
            if (ovhSubsidiary == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "ovhSubsidiary");
            }
            this.ovhSubsidiary = ovhSubsidiary;
            return this;
        }
        @CustomType.Setter
        public Builder phone(String phone) {
            if (phone == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "phone");
            }
            this.phone = phone;
            return this;
        }
        @CustomType.Setter
        public Builder phoneCountry(String phoneCountry) {
            if (phoneCountry == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "phoneCountry");
            }
            this.phoneCountry = phoneCountry;
            return this;
        }
        @CustomType.Setter
        public Builder sex(String sex) {
            if (sex == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "sex");
            }
            this.sex = sex;
            return this;
        }
        @CustomType.Setter
        public Builder spareEmail(String spareEmail) {
            if (spareEmail == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "spareEmail");
            }
            this.spareEmail = spareEmail;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder vat(String vat) {
            if (vat == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "vat");
            }
            this.vat = vat;
            return this;
        }
        @CustomType.Setter
        public Builder zip(String zip) {
            if (zip == null) {
              throw new MissingRequiredPropertyException("GetMeResult", "zip");
            }
            this.zip = zip;
            return this;
        }
        public GetMeResult build() {
            final var _resultValue = new GetMeResult();
            _resultValue.AccountURN = AccountURN;
            _resultValue.address = address;
            _resultValue.area = area;
            _resultValue.birthCity = birthCity;
            _resultValue.birthDay = birthDay;
            _resultValue.city = city;
            _resultValue.companyNationalIdentificationNumber = companyNationalIdentificationNumber;
            _resultValue.corporationType = corporationType;
            _resultValue.country = country;
            _resultValue.currencies = currencies;
            _resultValue.customerCode = customerCode;
            _resultValue.email = email;
            _resultValue.fax = fax;
            _resultValue.firstname = firstname;
            _resultValue.id = id;
            _resultValue.italianSdi = italianSdi;
            _resultValue.language = language;
            _resultValue.legalform = legalform;
            _resultValue.name = name;
            _resultValue.nationalIdentificationNumber = nationalIdentificationNumber;
            _resultValue.nichandle = nichandle;
            _resultValue.organisation = organisation;
            _resultValue.ovhCompany = ovhCompany;
            _resultValue.ovhSubsidiary = ovhSubsidiary;
            _resultValue.phone = phone;
            _resultValue.phoneCountry = phoneCountry;
            _resultValue.sex = sex;
            _resultValue.spareEmail = spareEmail;
            _resultValue.state = state;
            _resultValue.vat = vat;
            _resultValue.zip = zip;
            return _resultValue;
        }
    }
}