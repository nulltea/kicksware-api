using System;
using Core.Entities.Users;
using Core.Gateway;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Http;

namespace Web.Auth
{
	public class MiddlewareAuthOptions: AuthenticationSchemeOptions
	{
		public string Challenge { get; set; } = MiddlewareAuthDefaults.AuthenticationScheme;

		public RequestParams RequestParams { get; set; }

		public TimeSpan RequestTimeout { get; set; } = TimeSpan.FromMinutes(1);

		public Func<User, AuthClaims> ClaimSelector { get; set; }

		public bool SaveTokens { get; set; }

		public PathString LoginPath { get; set; }

		public PathString LogoutPath { get; set; }

		public PathString UnauthorisedPath { get; set; }

		public string ReturnUrlParameter { get; set; }

		public CookieBuilder Cookie { get; set; }

		public CookieSecurePolicy CookieSecure { get; set; } = CookieSecurePolicy.None;

		public TimeSpan CookieExpireTimeSpan { get; set; }

		public MiddlewareAuthOptions() { }
	}
}