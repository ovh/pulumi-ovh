// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Okms.Outputs
{

    [OutputType]
    public sealed class ServiceKeyJWKKey
    {
        /// <summary>
        /// The algorithm intended to be used with the key
        /// </summary>
        public readonly string? Alg;
        /// <summary>
        /// The cryptographic curve used with the key
        /// </summary>
        public readonly string? Crv;
        /// <summary>
        /// The RSA or EC private exponent
        /// </summary>
        public readonly string? D;
        /// <summary>
        /// The RSA private key's first factor CRT exponent
        /// </summary>
        public readonly string? Dp;
        /// <summary>
        /// The RSA private key's second factor CRT exponent
        /// </summary>
        public readonly string? Dq;
        /// <summary>
        /// The exponent value for the RSA public key
        /// </summary>
        public readonly string? E;
        /// <summary>
        /// The value of the symmetric (or other single-valued) key
        /// </summary>
        public readonly string? K;
        /// <summary>
        /// The operation for which the key is intended to be used
        /// </summary>
        public readonly ImmutableArray<string> KeyOps;
        /// <summary>
        /// key ID parameter used to match a specific key
        /// </summary>
        public readonly string? Kid;
        /// <summary>
        /// Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC
        /// </summary>
        public readonly string Kty;
        /// <summary>
        /// The modulus value for the RSA public key
        /// </summary>
        public readonly string? N;
        /// <summary>
        /// The first prime factor of the RSA private key
        /// </summary>
        public readonly string? P;
        /// <summary>
        /// The second prime factor of the RSA private key
        /// </summary>
        public readonly string? Q;
        /// <summary>
        /// The CRT coefficient of the second factor of the RSA private key
        /// </summary>
        public readonly string? Qi;
        /// <summary>
        /// The intended use of the public key
        /// </summary>
        public readonly string? Use;
        /// <summary>
        /// The x coordinate for the Elliptic Curve point
        /// </summary>
        public readonly string? X;
        /// <summary>
        /// The y coordinate for the Elliptic Curve point
        /// </summary>
        public readonly string? Y;

        [OutputConstructor]
        private ServiceKeyJWKKey(
            string? alg,

            string? crv,

            string? d,

            string? dp,

            string? dq,

            string? e,

            string? k,

            ImmutableArray<string> keyOps,

            string? kid,

            string kty,

            string? n,

            string? p,

            string? q,

            string? qi,

            string? use,

            string? x,

            string? y)
        {
            Alg = alg;
            Crv = crv;
            D = d;
            Dp = dp;
            Dq = dq;
            E = e;
            K = k;
            KeyOps = keyOps;
            Kid = kid;
            Kty = kty;
            N = n;
            P = p;
            Q = q;
            Qi = qi;
            Use = use;
            X = x;
            Y = y;
        }
    }
}