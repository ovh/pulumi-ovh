// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOkmsServiceKeyJwkKey {
    /**
     * @return The algorithm intended to be used with the key
     * 
     */
    private String alg;
    /**
     * @return The cryptographic curve used with the key
     * 
     */
    private String crv;
    /**
     * @return The exponent value for the RSA public key
     * 
     */
    private String e;
    /**
     * @return The operation for which the key is intended to be used
     * 
     */
    private List<String> keyOps;
    /**
     * @return key ID parameter used to match a specific key
     * 
     */
    private String kid;
    /**
     * @return Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC
     * 
     */
    private String kty;
    /**
     * @return The modulus value for the RSA public key
     * 
     */
    private String n;
    /**
     * @return The intended use of the public key
     * 
     */
    private String use;
    /**
     * @return The x coordinate for the Elliptic Curve point
     * 
     */
    private String x;
    /**
     * @return The y coordinate for the Elliptic Curve point
     * 
     */
    private String y;

    private GetOkmsServiceKeyJwkKey() {}
    /**
     * @return The algorithm intended to be used with the key
     * 
     */
    public String alg() {
        return this.alg;
    }
    /**
     * @return The cryptographic curve used with the key
     * 
     */
    public String crv() {
        return this.crv;
    }
    /**
     * @return The exponent value for the RSA public key
     * 
     */
    public String e() {
        return this.e;
    }
    /**
     * @return The operation for which the key is intended to be used
     * 
     */
    public List<String> keyOps() {
        return this.keyOps;
    }
    /**
     * @return key ID parameter used to match a specific key
     * 
     */
    public String kid() {
        return this.kid;
    }
    /**
     * @return Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC
     * 
     */
    public String kty() {
        return this.kty;
    }
    /**
     * @return The modulus value for the RSA public key
     * 
     */
    public String n() {
        return this.n;
    }
    /**
     * @return The intended use of the public key
     * 
     */
    public String use() {
        return this.use;
    }
    /**
     * @return The x coordinate for the Elliptic Curve point
     * 
     */
    public String x() {
        return this.x;
    }
    /**
     * @return The y coordinate for the Elliptic Curve point
     * 
     */
    public String y() {
        return this.y;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOkmsServiceKeyJwkKey defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String alg;
        private String crv;
        private String e;
        private List<String> keyOps;
        private String kid;
        private String kty;
        private String n;
        private String use;
        private String x;
        private String y;
        public Builder() {}
        public Builder(GetOkmsServiceKeyJwkKey defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alg = defaults.alg;
    	      this.crv = defaults.crv;
    	      this.e = defaults.e;
    	      this.keyOps = defaults.keyOps;
    	      this.kid = defaults.kid;
    	      this.kty = defaults.kty;
    	      this.n = defaults.n;
    	      this.use = defaults.use;
    	      this.x = defaults.x;
    	      this.y = defaults.y;
        }

        @CustomType.Setter
        public Builder alg(String alg) {
            if (alg == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "alg");
            }
            this.alg = alg;
            return this;
        }
        @CustomType.Setter
        public Builder crv(String crv) {
            if (crv == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "crv");
            }
            this.crv = crv;
            return this;
        }
        @CustomType.Setter
        public Builder e(String e) {
            if (e == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "e");
            }
            this.e = e;
            return this;
        }
        @CustomType.Setter
        public Builder keyOps(List<String> keyOps) {
            if (keyOps == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "keyOps");
            }
            this.keyOps = keyOps;
            return this;
        }
        public Builder keyOps(String... keyOps) {
            return keyOps(List.of(keyOps));
        }
        @CustomType.Setter
        public Builder kid(String kid) {
            if (kid == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "kid");
            }
            this.kid = kid;
            return this;
        }
        @CustomType.Setter
        public Builder kty(String kty) {
            if (kty == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "kty");
            }
            this.kty = kty;
            return this;
        }
        @CustomType.Setter
        public Builder n(String n) {
            if (n == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "n");
            }
            this.n = n;
            return this;
        }
        @CustomType.Setter
        public Builder use(String use) {
            if (use == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "use");
            }
            this.use = use;
            return this;
        }
        @CustomType.Setter
        public Builder x(String x) {
            if (x == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "x");
            }
            this.x = x;
            return this;
        }
        @CustomType.Setter
        public Builder y(String y) {
            if (y == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkKey", "y");
            }
            this.y = y;
            return this;
        }
        public GetOkmsServiceKeyJwkKey build() {
            final var _resultValue = new GetOkmsServiceKeyJwkKey();
            _resultValue.alg = alg;
            _resultValue.crv = crv;
            _resultValue.e = e;
            _resultValue.keyOps = keyOps;
            _resultValue.kid = kid;
            _resultValue.kty = kty;
            _resultValue.n = n;
            _resultValue.use = use;
            _resultValue.x = x;
            _resultValue.y = y;
            return _resultValue;
        }
    }
}
